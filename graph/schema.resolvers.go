package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/deyr02/bnzlcrm/graph/generated"
	"github.com/deyr02/bnzlcrm/graph/model"
)

// AddNewElementMetaUser is the resolver for the AddNewElement_Meta_User field.
func (r *mutationResolver) AddNewElementMetaUser(ctx context.Context, input *model.NewCustomFieldElement) (*model.MetaUserCollection, error) {
	panic(fmt.Errorf("not implemented"))
}

// ModifyElementMetaUser is the resolver for the ModifyElement_Meta_user field.
func (r *mutationResolver) ModifyElementMetaUser(ctx context.Context, id *string, input *model.NewCustomFieldElement) (*model.MetaUserCollection, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeleteElementMetaUser is the resolver for the DeleteElement_Meta_user field.
func (r *mutationResolver) DeleteElementMetaUser(ctx context.Context, id *string) (*model.MetaUserCollection, error) {
	panic(fmt.Errorf("not implemented"))
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

// GetUserCollectionMeta is the resolver for the GetUserCollection_Meta field.
func (r *queryResolver) GetUserCollectionMeta(ctx context.Context) (*model.MetaUserCollection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
