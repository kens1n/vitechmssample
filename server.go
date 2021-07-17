package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
	"github.com/kens1n/vitechmssample/external_services"
	"github.com/kens1n/vitechmssample/graph"
	"github.com/kens1n/vitechmssample/graph/generated"
	"github.com/kens1n/vitechmssample/postgres"
	"github.com/kens1n/vitechmssample/tracing"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
	"log"
	"net/http"
	"os"
	"time"
)

const defaultPort = "8080"

type LogRecord struct {
	Time        string `json:"time"`
	Url         string `json:"url"`
	Request     string `json:"request"`
	ElapsedTime string `json:"elapsed_time"`
}

func logWithOperationContext(oc *graphql.OperationContext) {
	logRecord := LogRecord{
		Time:        time.Now().Format(time.RFC3339),
		Request:     oc.RawQuery,
		ElapsedTime: time.Now().Sub(oc.Stats.OperationStart).String(),
	}
	jsLogRecord, _ := json.Marshal(logRecord)
	log.Println(string(jsLogRecord))
}

func main() {
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	closer, err := cfg.InitGlobalTracer(
		"jaeger-vitechmssample",
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	if err != nil {
		log.Printf("Could not initialize jaeger tracer: %s", err.Error())
		return
	}
	defer closer.Close()

	godotenv.Load(".env")

	db := pg.Connect(&pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Database: os.Getenv("DB_NAME"),
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
	})

	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		HdataRepo:   postgres.HdataRepo{DB: db},
		GuidService: external_services.GuidServiceLocal{},
		HashService: external_services.HashServiceLocal{},
	}}))

	srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		oc := graphql.GetOperationContext(ctx)

		// Костыль потому что генерируется много левых запросов скорее всего графическим интерфейсом
		if oc.OperationName == "" {
			ctx = tracing.StartSpanForContext(ctx)
			defer opentracing.SpanFromContext(ctx).Finish()
			defer logWithOperationContext(oc)
		}

		return next(ctx)
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Printf("connect to http://localhost:16686/search for jaeger")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
