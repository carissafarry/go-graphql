package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"github.com/vektah/gqlparser/v2/ast"
	"go-graphql/internal/transport/graphql/resolvers"
	"go-graphql/internal/transport/graphql/graph"

	// domain Usecase
	"go-graphql/internal/domain/post"
	"go-graphql/internal/domain/user"

	// infra
	"go-graphql/internal/infra/db/postgres"
	"go-graphql/internal/infra/security"
	redisinfra "go-graphql/internal/infra/cache/redis"
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
	// REDIS
	// =========================
	host := os.Getenv("REDIS_HOST")
	if host == "" {
		host = "redis"
	}

	redisPort := os.Getenv("REDIS_PORT")
	if redisPort == "" {
		redisPort = "6379"
	}

	addr := host + ":" + redisPort

	log.Printf("REDIS ADDR = %s", addr)
	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr, // e.g. localhost:6379
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("failed to connect redis: %v", err)
	}


	// =========================
	// REPOSITORIES
	// =========================
	userRepo := postgres.NewUserRepo(userDB)
	postRepo := postgres.NewPostRepo(postDB)


	// =========================
	// INFRA IMPLEMENTATIONS
	// =========================
	pendingUserStore := redisinfra.NewPendingUserStore(redisClient) 
	otpStore := redisinfra.NewOTPStore(redisClient) 
	otpGenerator := security.NewOTPGenerator()


	// =========================
	// USECASES
	// =========================
	userUsecase := user.NewUsecase(
		userRepo,
		pendingUserStore,
		otpStore,
		otpGenerator,
	)
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
