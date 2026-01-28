package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/vektah/gqlparser/v2/ast"
	"go-graphql/cmd/gateway/resolvers"
	"go-graphql/graph"

	// domain Usecase
	"go-graphql/internal/domain/post"
	"go-graphql/internal/domain/user"

	// infra
	"go-graphql/internal/infra/postgres"
)

const defaultPort = "8080"

func main() {
	// =========================
	// PORT
	// =========================
	_ = godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}


	// =========================
	// DATABASE CONNECTIONS
	// =========================
	userDBCfg := postgres.LoadConfig("USER_DB")

	userDB, err := postgres.NewPostgres(userDBCfg)
	if err != nil {
		log.Fatal(err)
	}
	defer userDB.Close()
	
	postDBCfg := postgres.LoadConfig("POST_DB")
	postDB, err := postgres.NewPostgres(postDBCfg)
	if err != nil {
		log.Fatal(err)
	}
	defer postDB.Close()


	// =========================
	// REPOSITORIES
	// =========================
	userRepo := postgres.NewUserRepo(userDB)
	postRepo := postgres.NewPostRepo(postDB)


	// =========================
	// USECASES
	// =========================
	userUsecase := user.NewUsecase(userRepo)
	postUsecase := post.NewUsecase(postRepo)


	// =========================
	// RESOLVER (DI)
	// =========================
	resolver := &resolvers.Resolver{
		UserUsecase: userUsecase,
		PostUsecase: postUsecase,
	}


	// =========================
	// GRAPHQL SERVER
	// =========================
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
