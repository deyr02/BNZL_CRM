package userrole

import (
	"context"

	"github.com/deyr02/bnzlcrm/graph/model"
	"github.com/deyr02/bnzlcrm/repositories/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	USER_ROLE = "user_role"
)

type Database struct {
	client *mongo.Client
}
type User_Role_Repository interface {
	GetAllUserRole(ctx context.Context, _id string) ([]*model.UserRole, error)
	GetUserRoleID(ctx context.Context, _id string) (*model.UserRole, error)
	AddNewUserRole(ctx context.Context, input model.NewUserRole) (*model.UserRole, error)
	ModifyUserRole(ctx context.Context, _id string, input *model.NewUserRole) (*model.UserRole, error)
	DeleUserRole(ctx context.Context, _id string) (string, error)
}

func New_User_Role_repository(client *mongo.Client) User_Role_Repository {
	_client := client
	collection := _client.Database(database.DATABASE_NAME).Collection(USER_ROLE)
	_count, _ := collection.CountDocuments(nil, bson.D{})

	if _count == 0 {
		var adminRole model.UserRole
		var itsupportRole model.UserRole

		adminRole.RoleID = "1111111111111111111"
		adminRole.RoleName = AdminUserRole.RoleName
		adminRole.SystemRole = true
		adminRole.Operations = AdminUserRole.Operations

		_, err := collection.InsertOne(nil, adminRole)
		if err != nil {
			panic(err)
		}

		itsupportRole.RoleID = "1111111111111111112"
		itsupportRole.RoleName = ITSupportUserRole.RoleName
		itsupportRole.SystemRole = false
		itsupportRole.Operations = ITSupportUserRole.Operations

		_, err_1 := collection.InsertOne(nil, itsupportRole)
		if err_1 != nil {
			panic(err)
		}
	}
	return &Database{
		client: _client,
	}
}

func (db *Database) GetAllUserRole(ctx context.Context, _id string) ([]*model.UserRole, error) {

	return nil, nil
}

func (db *Database) GetUserRoleID(ctx context.Context, _id string) (*model.UserRole, error) {

	return nil, nil
}

func (db *Database) AddNewUserRole(ctx context.Context, input model.NewUserRole) (*model.UserRole, error) {

	return nil, nil
}

func (db *Database) ModifyUserRole(ctx context.Context, _id string, input *model.NewUserRole) (*model.UserRole, error) {

	return nil, nil
}

func (db *Database) DeleUserRole(ctx context.Context, _id string) (string, error) {

	return "", nil
}
