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

// GetUserRoleID is the resolver for the GetUserRoleID field.
func (r *mutationResolver) GetUserRoleID(ctx context.Context, id *string) (*model.UserRole, error) {
	panic(fmt.Errorf("not implemented"))
}

// AddNewUserRole is the resolver for the AddNewUserRole field.
func (r *mutationResolver) AddNewUserRole(ctx context.Context, input *model.NewUserRole) (*model.UserRole, error) {
	panic(fmt.Errorf("not implemented"))
}

// ModifyUserRole is the resolver for the ModifyUserRole field.
func (r *mutationResolver) ModifyUserRole(ctx context.Context, id *string, input *model.NewUserRole) (*model.UserRole, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeleUserRole is the resolver for the DeleUserRole field.
func (r *mutationResolver) DeleUserRole(ctx context.Context, id *string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetUserMetaCollection is the resolver for the GetUserMetaCollection field.
func (r *queryResolver) GetUserMetaCollection(ctx context.Context) (*model.MetaUserCollection, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetAllUserRole is the resolver for the GetAllUserRole field.
func (r *queryResolver) GetAllUserRole(ctx context.Context) ([]*model.UserRole, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
