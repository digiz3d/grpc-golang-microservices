package main

import (
	pb "comments/protobuf/generated"
	"context"
	"net"
	"os"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type comment struct {
	ID     string
	showId int64
	userId int64
	text   string
}

type commentsServer struct {
	pb.UnimplementedCommentServer
	commentsPerShow map[int64][]*comment
}

func (s *commentsServer) AddComment(ctx context.Context, req *pb.AddCommentRequest) (*pb.AddCommentReply, error) {
	comment := &comment{ID: uuid.NewString(), showId: req.ShowId, userId: req.UserId, text: req.Text}
	s.commentsPerShow[req.ShowId] = append(s.commentsPerShow[req.ShowId], comment)
	return &pb.AddCommentReply{Id: comment.ID, ShowId: comment.showId, UserId: comment.userId, Text: comment.text}, nil
}

func newServer() *commentsServer {
	s := &commentsServer{commentsPerShow: make(map[int64][]*comment)}
	return s
}

func main() {
	port := os.Getenv("PORT")
	netListener, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		return
	}
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterCommentServer(grpcServer, newServer())
	grpcServer.Serve(netListener)
}
