package main

import (
	pb "comments/protobuf/generated"
	"context"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	goredis "github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

type comment struct {
	ID     string
	showId string
	userId string
	text   string
}

type commentsServer struct {
	pb.UnimplementedCommentServer
	commentsPerShow map[string][]*comment
	Redis           *goredis.Client
}

func (s *commentsServer) AddComment(ctx context.Context, req *pb.AddCommentRequest) (*pb.AddedComment, error) {
	comment := &comment{ID: uuid.NewString(), showId: req.ShowId, userId: req.UserId, text: req.Text}
	s.commentsPerShow[req.ShowId] = append(s.commentsPerShow[req.ShowId], comment)

	addedComment := &pb.AddedComment{Id: comment.ID, ShowId: comment.showId, UserId: comment.userId, Text: comment.text}

	bytes, err := proto.Marshal(addedComment)
	if err == nil {
		s.Redis.Publish(ctx, "comment.added@show."+req.ShowId, bytes)
	}

	return addedComment, nil
}

func (s *commentsServer) OnCommentAdded(req *pb.OnCommentAddedRequest, stream pb.Comment_OnCommentAddedServer) error {
	ctx := stream.Context()
	sub := s.Redis.Subscribe(ctx, "comment.added@show."+req.ShowId)
	_, err := sub.Receive(ctx)
	if err != nil {
		return err
	}

	subscribedChannel := sub.Channel()

	for {
		select {
		case message := <-subscribedChannel:
			var addedComment pb.AddedComment
			err = proto.Unmarshal([]byte(message.Payload), &addedComment)
			if err != nil {
				fmt.Print("err 1")
				sub.Close()
				return err
			}
			err = stream.Send(&addedComment)
			if err != nil {
				fmt.Print("err 2")
				sub.Close()
				return err
			}
		case <-ctx.Done():
			fmt.Print("context closed")
			sub.Close()
			return nil
		}
	}
}

func newServer() *commentsServer {
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	DB, err := strconv.Atoi(os.Getenv("REDIS_DB"))

	if err != nil {
		DB = 0
	}

	redis := goredis.NewClient(&goredis.Options{Addr: redisAddr, ReadTimeout: time.Second * 60, DB: DB})

	s := &commentsServer{
		commentsPerShow: make(map[string][]*comment),
		Redis:           redis,
	}
	return s
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	netListener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return
	}
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterCommentServer(grpcServer, newServer())
	fmt.Printf("Listening on port %s\n", port)
	grpcServer.Serve(netListener)
}
