package main

import (
	"context"
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/riyan10dec/tax-calc/business-service/product/repository"
	gcontext "github.com/riyan10dec/tax-calc/context"
	"github.com/riyan10dec/tax-calc/resolvers"
	"github.com/riyan10dec/tax-calc/schema"

	"github.com/rs/cors"
)

func main() {

	config := gcontext.LoadConfig(".")

	db, err := gcontext.OpenDB(config)
	if err != nil {
		log.Fatalf("Unable to connect to db: %s \n", err)
	}
	ctx := context.Background()

	br := repository.NewSqlProductRepository(db)
	ctx = context.WithValue(ctx, "ProductRepository", br)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST", "GET", "OPTIONS"},
		AllowedHeaders: []string{"X-Requested-With",
			"Authorization",
			"Content-Type",
			"X-Auth-Token",
			"Origin",
			"Accept"},
		Debug: true,
	})
	Handle(ctx, c)

	log.Panic(http.ListenAndServe(":8080", nil))
}
func Handle(ctx context.Context, c *cors.Cors) {
	// Default Schema
	graphql.MaxDepth(5)
	graphql.UseStringDescriptions()

	resolver1 := &resolvers.Resolver{}
	graphqlSchema := graphql.MustParseSchema(schema.GetRootSchema(), resolver1)
	http.Handle("/query",
		c.Handler(
			AddContext(ctx, &relay.Handler{Schema: graphqlSchema}),
		),
	)

}

func AddContext(ctx context.Context, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
