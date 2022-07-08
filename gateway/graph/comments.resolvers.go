package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gateway/graph/model"
	pbComments "services/comments/protobuf/generated"
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
