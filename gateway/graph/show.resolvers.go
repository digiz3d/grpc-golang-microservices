package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gateway/graph/generated"
	"gateway/graph/model"
	protogen "gateway/pb/generated"
	"log"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

func (r *mutationResolver) CreateShow(ctx context.Context, input model.CreateShowInput) (*model.CreateShowPayload, error) {
	foundUser := r.UsersRepository[input.UserID]

	if foundUser == nil {
		return nil, fmt.Errorf("user not found")
	}

	show := &model.Show{ID: uuid.NewString(), Name: input.Name, Description: input.Description, UserID: foundUser.ID}

	r.ShowsRepository[show.ID] = show
	createShowPayload := &model.CreateShowPayload{Show: show}

	event := &protogen.ShowCreatedEvent{Id: show.ID}
	bytes, err := proto.Marshal(event)
	if err != nil {
		fmt.Println("mince alors")
	}
	r.Redis.Publish(ctx, SHOW_CREATED, bytes)

	return createShowPayload, nil
}

func (r *queryResolver) Show(ctx context.Context, id string) (*model.Show, error) {
	show := r.ShowsRepository[id]
	if show == nil {
		return nil, fmt.Errorf("show not found")
	}
	return show, nil
}

func (r *showResolver) User(ctx context.Context, obj *model.Show) (*model.User, error) {
	user := r.UsersRepository[obj.UserID]
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (r *subscriptionResolver) OnCreateShow(ctx context.Context) (<-chan *model.CreateShowPayload, error) {
	channel := make(chan *model.CreateShowPayload, 1)

	go func() {
		sub := r.Redis.Subscribe(ctx, SHOW_CREATED)
		_, err := sub.Receive(ctx)
		if err != nil {
			return
		}
		ch := sub.Channel()
		for {
			select {
			case message := <-ch:
				var event protogen.ShowCreatedEvent
				err := proto.Unmarshal([]byte(message.Payload), &event)
				if err != nil {
					log.Println(err)
					return
				}
				channel <- &model.CreateShowPayload{Show: &model.Show{ID: event.Id}}
			case <-ctx.Done():
				sub.Close()
				return
			}
		}
	}()

	return channel, nil
}

// Show returns generated.ShowResolver implementation.
func (r *Resolver) Show() generated.ShowResolver { return &showResolver{r} }

type showResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
const SHOW_CREATED = "SHOW_CREATED"
