package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/deyr02/bnzlcrm/graph/generated"
	"github.com/deyr02/bnzlcrm/graph/model"
	"github.com/deyr02/bnzlcrm/repositories/database"
	"github.com/deyr02/bnzlcrm/repositories/user/userCollection"
	user "github.com/deyr02/bnzlcrm/repositories/user/userMeta"
	userrole "github.com/deyr02/bnzlcrm/repositories/userRole"
)

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///---------------------------------------- Repositories --------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
var client = database.CreateConnection()
var userRepository user.User_Meta_Repository = user.New_User_Meta_repository(client)
var UserRoleRepository userrole.User_Role_Repository = userrole.New_User_Role_repository(client)
var userCollectionRepository userCollection.User_Repository = userCollection.New_User_Repository(client)

///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///----------------------------------------- Mutations ----------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------

///--------------------------------------------------------------------------------------------------------
///------------------------------------Meta User Collection------------------------------------------------
///--------------------------------------------------------------------------------------------------------

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

///--------------------------------------------------------------------------------------------------------
///--------------------------------------- User Collection ------------------------------------------------
///--------------------------------------------------------------------------------------------------------

// AddNewUser is the resolver for the AddNewUser field.
func (r *mutationResolver) AddNewUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	return userCollectionRepository.AddNewUser(ctx, input)
}

// DeleteUser is the resolver for the DeleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id *string) (string, error) {
	return userCollectionRepository.DeleteUser(ctx, *id)
}

// ModifyUser is the resolver for the ModifyUser field.
func (r *mutationResolver) ModifyUser(ctx context.Context, id *string, input *model.NewUser) (*model.User, error) {
	return userCollectionRepository.ModifyUser(ctx, *id, input)
}

// Login is the resolver for the Login field.
func (r *mutationResolver) Login(ctx context.Context, input *model.Login) (*model.UserDto, error) {
	return userCollectionRepository.Login(ctx, input)
}

///--------------------------------------------------------------------------------------------------------
///------------------------------------ User Role collection ----------------------------------------------
///--------------------------------------------------------------------------------------------------------

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

///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///------------------------------------------ Queries -----------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------
///--------------------------------------------------------------------------------------------------------

///--------------------------------------------------------------------------------------------------------
///------------------------------------ User Meta collection ----------------------------------------------
///--------------------------------------------------------------------------------------------------------

// GetUserMetaCollection is the resolver for the GetUserMetaCollection field.
func (r *queryResolver) GetUserMetaCollection(ctx context.Context) (*model.MetaUserCollection, error) {
	return userRepository.GetUser_MetaCollection(ctx)
}

///--------------------------------------------------------------------------------------------------------
///------------------------------------ User Role collection ----------------------------------------------
///--------------------------------------------------------------------------------------------------------
// GetAllUserRole is the resolver for the GetAllUserRole field.
func (r *queryResolver) GetAllUserRole(ctx context.Context) ([]*model.UserRole, error) {
	return UserRoleRepository.GetAllUserRole(ctx)
}

// GetUserRoleByID is the resolver for the GetUserRoleByID field.
func (r *queryResolver) GetUserRoleByID(ctx context.Context, id *string) (*model.UserRole, error) {
	return UserRoleRepository.GetUserRoleByID(ctx, *id)
}

///--------------------------------------------------------------------------------------------------------
///------------------------------------ Usercollection ----------------------------------------------
///--------------------------------------------------------------------------------------------------------
// GetAllUser is the resolver for the GetAllUser field.
func (r *queryResolver) GetAllUser(ctx context.Context) ([]*model.User, error) {
	return userCollectionRepository.GetAllUser(ctx)
}

// GetUserByID is the resolver for the GetUserByID field.
func (r *queryResolver) GetUserByID(ctx context.Context, id *string) (*model.User, error) {
	return userCollectionRepository.GetUserByID(ctx, *id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
