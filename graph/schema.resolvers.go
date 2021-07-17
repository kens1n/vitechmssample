package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/kens1n/vitechmssample/graph/generated"
	"github.com/kens1n/vitechmssample/postgres"
	"github.com/kens1n/vitechmssample/tracing"
)

func (r *queryResolver) Hashcode(ctx context.Context, code string) (string, error) {

	//guid, _ := r.GuidService.GuidGenerate(code)
	//hash, _ := r.HashService.HashCalc(code)
	span := tracing.StartSpanWithRootSpanInContext(ctx, "Hashcode")
	defer span.Finish()

	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	go getGuid(code, r, ch1, ctx)
	go getHash(code, r, ch2, ctx)

	guid := <-ch1
	hash := <-ch2

	return getHdataDataByGuidAndHash(
		r.HdataRepo,
		ctx,
		guid,
		hash,
	)
}

func getGuid(code string, r *queryResolver, ch chan string, ctx context.Context) {
	span := tracing.StartSpanWithRootSpanInContext(ctx, "getGuid")
	defer span.Finish()
	defer close(ch)
	guid, _ := r.GuidService.GuidGenerate(code)

	ch <- guid
}

func getHash(code string, r *queryResolver, ch chan string, ctx context.Context) {
	span := tracing.StartSpanWithRootSpanInContext(ctx, "getHash")
	defer span.Finish()
	defer close(ch)
	hash, _ := r.HashService.HashCalc(code)

	ch <- hash
}

func getHdataDataByGuidAndHash(hdataRepo postgres.HdataRepo, ctx context.Context, guid string, hash string) (string, error) {
	span := tracing.StartSpanWithRootSpanInContext(ctx, "getHdataDataByGuidAndHash")
	defer span.Finish()

	return hdataRepo.GetHdataDataByGuidAndHash(
		guid,
		hash,
	)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
