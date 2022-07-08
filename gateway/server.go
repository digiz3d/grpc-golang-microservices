package main

//go:generate go run github.com/99designs/gqlgen generate

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"gateway/graph"
	"gateway/graph/generated"
	"gateway/graph/model"

	pbComments "services/comments/protobuf/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	goredis "github.com/go-redis/redis/v9"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const defaultPort = "3000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	DB, err := strconv.Atoi(os.Getenv("REDIS_DB"))

	if err != nil {
		DB = 0
	}

	redis := goredis.NewClient(&goredis.Options{Addr: redisAddr, ReadTimeout: time.Second * 60, DB: DB})

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	connCommentService, err := grpc.Dial(os.Getenv("SERVICE_COMMENTS_ADDR"), opts...)

	if err != nil {
		log.Fatalf("failed to dial: %v", err)
		return
	}
	log.Printf("ok dialed")

	commentServiceClient := pbComments.NewCommentClient(connCommentService)

	srv := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		CommentsService: commentServiceClient,
		Redis:           redis,
		ShowsRepository: make(map[string]*model.Show),
		UsersRepository: make(map[string]*model.User),
	}}))

	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	srv.Use(extension.Introspection{})

	http.Handle("/playground", playground.Handler("Playground", "/graphql"))
	http.Handle("/graphql", c.Handler(srv))

	log.Printf("connect to http://localhost:%s/playground for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
