package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/kens1n/vitechmssample/graph/generated"
	"github.com/kens1n/vitechmssample/postgres"
	"github.com/kens1n/vitechmssample/tracing"
)

type guidChannelResponse struct {
	guid string
	err  error
}

type hashChannelResponse struct {
	hash string
	err  error
}

func (r *queryResolver) Hashcode(ctx context.Context, code string) (string, error) {

	//guid, _ := r.GuidService.GuidGenerate(code)
	//hash, _ := r.HashService.HashCalc(code)
	span := tracing.StartSpanWithRootSpanInContext(ctx, "Hashcode")
	defer span.Finish()

	ch1 := make(chan guidChannelResponse, 1)
	ch2 := make(chan hashChannelResponse, 1)

	go getGuid(code, r, ch1, ctx)
	go getHash(code, r, ch2, ctx)

	guidResponse := <-ch1
	if guidResponse.err != nil {
		return "", guidResponse.err
	}
	hashResponse := <-ch2
	if hashResponse.err != nil {
		return "", hashResponse.err
	}

	return getHdataDataByGuidAndHash(
		r.HdataRepo,
		ctx,
		guidResponse.guid,
		hashResponse.hash,
	)
}

func getGuid(code string, r *queryResolver, ch chan guidChannelResponse, ctx context.Context) {
	span := tracing.StartSpanWithRootSpanInContext(ctx, "getGuid")
	defer span.Finish()
	defer close(ch)
	guid, err := r.GuidService.GuidGenerate(code)

	ch <- guidChannelResponse{
		guid,
		err,
	}
}

func getHash(code string, r *queryResolver, ch chan hashChannelResponse, ctx context.Context) {
	span := tracing.StartSpanWithRootSpanInContext(ctx, "getHash")
	defer span.Finish()
	defer close(ch)
	hash, err := r.HashService.HashCalc(code)

	ch <- hashChannelResponse{
		hash,
		err,
	}
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
