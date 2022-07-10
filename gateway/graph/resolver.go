package graph

import (
	"gateway/graph/model"

	commentsService "services/comments/pb/generated"

	goredis "github.com/go-redis/redis/v9"
)

type Resolver struct {
	CommentsService commentsService.CommentClient
	Redis           *goredis.Client
	ShowsRepository map[string]*model.Show
	UsersRepository map[string]*model.User
}
