package graph

import (
	"github.com/digiz3d/graphgogen/graph/model"
	goredis "github.com/go-redis/redis/v9"
)

type Resolver struct {
	ShowsRepository map[string]*model.Show
	UsersRepository map[string]*model.User
	Redis           *goredis.Client
}
