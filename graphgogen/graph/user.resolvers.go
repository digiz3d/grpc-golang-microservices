package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/digiz3d/graphgogen/graph/generated"
	"github.com/digiz3d/graphgogen/graph/model"
	"github.com/google/uuid"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.CreateUserPayload, error) {
	newUser := &model.User{ID: uuid.NewString(), Username: "ok"}
	r.UsersRepository[newUser.ID] = newUser
	return &model.CreateUserPayload{User: r.UsersRepository[newUser.ID]}, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	user := r.UsersRepository[id]
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (r *userResolver) Shows(ctx context.Context, obj *model.User) ([]*model.Show, error) {
	shows := make([]*model.Show, 0)

	for _, currentShow := range r.ShowsRepository {
		if currentShow.UserID == obj.ID {
			shows = append(shows, currentShow)
		}
	}

	return shows, nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
