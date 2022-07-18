package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/deyr02/bnzlcrm/graph/generated"
	"github.com/deyr02/bnzlcrm/graph/model"
	"github.com/deyr02/bnzlcrm/repositories/user"
)

var userRepository user.User_Meta_Repository = user.New_User_Meta_repository()

// AddNewElementMetaUser is the resolver for the AddNewElement_Meta_User field.
func (r *mutationResolver) AddNewElementMetaUser(ctx context.Context, input *model.NewCustomFieldElement) (*model.MetaUserCollection, error) {
	return userRepository.AddNewElement_Meta_User(ctx, input)
}

// ModifyElementMetaUser is the resolver for the ModifyElement_Meta_user field.
func (r *mutationResolver) ModifyElementMetaUser(ctx context.Context, id *string, input *model.NewCustomFieldElement) (*model.MetaUserCollection, error) {
	return userRepository.ModifyElement_Meta_User(ctx, *id, input)
}

// DeleteElementMetaUser is the resolver for the DeleteElement_Meta_user field.
func (r *mutationResolver) DeleteElementMetaUser(ctx context.Context, id *string) (*model.MetaUserCollection, error) {
	return userRepository.DeleteElement_Meta_User(ctx, *id)
}

// GetAllUser is the resolver for the GetAllUser field.
func (r *mutationResolver) GetAllUser(ctx context.Context) ([]*model.UserCollection, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetUserByID is the resolver for the GetUserByID field.
func (r *mutationResolver) GetUserByID(ctx context.Context, id *string) (*model.UserCollection, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeleteUserByID is the resolver for the DeleteUserByID field.
func (r *mutationResolver) DeleteUserByID(ctx context.Context, id *string) (*model.UserCollection, error) {
	panic(fmt.Errorf("not implemented"))
}

// ModifyUserByID is the resolver for the ModifyUserByID field.
func (r *mutationResolver) ModifyUserByID(ctx context.Context, id *string) (*model.UserCollection, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetUserMetaCollection is the resolver for the GetUserMetaCollection field.
func (r *queryResolver) GetUserMetaCollection(ctx context.Context) (*model.MetaUserCollection, error) {
	return userRepository.GetUser_MetaCollection(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
