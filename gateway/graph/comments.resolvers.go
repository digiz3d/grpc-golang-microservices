package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gateway/graph/model"
	"io"
	pbComments "services/comments/pb/generated"
)

func (r *mutationResolver) AddComment(ctx context.Context, input model.AddCommentInput) (*model.AddCommentPayload, error) {
	result, err := r.CommentsService.AddComment(ctx, &pbComments.AddCommentRequest{
		Text: input.Text, ShowId: input.ShowID, UserId: input.UserID,
	})

	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return &model.AddCommentPayload{Comment: &model.Comment{ID: result.Id, UserID: result.UserId, ShowID: result.ShowId, Text: result.Text}}, nil
}

func (r *subscriptionResolver) OnCommentAdded(ctx context.Context, input model.OnCommentAddedInput) (<-chan *model.OnCommentAddedPayload, error) {
	channel := make(chan *model.OnCommentAddedPayload, 1)

	stream, err := r.CommentsService.OnCommentAdded(context.Background(), &pbComments.OnCommentAddedRequest{ShowId: input.ShowID})

	if err != nil {
		return nil, err
	}

	go func() {
		for {
			addedComment, err := stream.Recv()
			if err == io.EOF {
				fmt.Printf("End of file ... %v", err)
				// close(channel)
				break
			}
			if err != nil {
				fmt.Printf("Some other error... %vx", err)
			}

			comment := model.Comment{ID: addedComment.Id, UserID: addedComment.UserId, ShowID: addedComment.ShowId, Text: addedComment.Text}
			channel <- &model.OnCommentAddedPayload{Comment: &comment}
		}
	}()

	return channel, nil
}
