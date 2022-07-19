package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/deyr02/bnzlcrm/graph/generated"
	"github.com/deyr02/bnzlcrm/graph/model"
	"github.com/deyr02/bnzlcrm/repositories/database"
	"github.com/deyr02/bnzlcrm/repositories/user"
	userrole "github.com/deyr02/bnzlcrm/repositories/userRole"
)

var client = database.CreateConnection()
var userRepository user.User_Meta_Repository = user.New_User_Meta_repository(client)
var UserRoleRepository userrole.User_Role_Repository = userrole.New_User_Role_repository(client)

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

// AddNewUserRole is the resolver for the AddNewUserRole field.
func (r *mutationResolver) AddNewUserRole(ctx context.Context, input *model.NewUserRole) (*model.UserRole, error) {
	return UserRoleRepository.AddNewUserRole(ctx, *input)
}

// ModifyUserRole is the resolver for the ModifyUserRole field.
func (r *mutationResolver) ModifyUserRole(ctx context.Context, id *string, input *model.NewUserRole) (*model.UserRole, error) {
	return UserRoleRepository.ModifyUserRole(ctx, *id, input)
}

// DeleUserRole is the resolver for the DeleUserRole field.
func (r *mutationResolver) DeleUserRole(ctx context.Context, id *string) (string, error) {
	return UserRoleRepository.DeleUserRole(ctx, *id)
}

// GetUserMetaCollection is the resolver for the GetUserMetaCollection field.
func (r *queryResolver) GetUserMetaCollection(ctx context.Context) (*model.MetaUserCollection, error) {
	return userRepository.GetUser_MetaCollection(ctx)
}

// GetAllUserRole is the resolver for the GetAllUserRole field.
func (r *queryResolver) GetAllUserRole(ctx context.Context) ([]*model.UserRole, error) {
	return UserRoleRepository.GetAllUserRole(ctx)
}

// GetUserRoleByID is the resolver for the GetUserRoleByID field.
func (r *queryResolver) GetUserRoleByID(ctx context.Context, id *string) (*model.UserRole, error) {
	return UserRoleRepository.GetUserRoleByID(ctx, *id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
