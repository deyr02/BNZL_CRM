package userCollection

import (
	"context"

	"github.com/deyr02/bnzlcrm/graph/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	USER_COLLECTION = "user"
)

type Database struct {
	client *mongo.Client
}

type User_Repository interface {
	GetAllUser(ctx context.Context) ([]*model.User, error)
	GetUserByID(ctx context.Context, _id string) (*model.User, error)
	AddNewUser(ctx context.Context, input *model.NewUser) (*model.User, error)
	ModifyUser(ctx context.Context, _id string, input *model.NewUser) (*model.User, error)
	DeleteUser(ctx context.Context, _id string) (string, error)
	Login(ctx context.Context, input *model.Login) (*model.UserDto, error)
}

func New_User_Repository(client *mongo.Client) User_Repository {
	_client := client
	return &Database{
		client: _client,
	}
}

func (db *Database) GetAllUser(ctx context.Context) ([]*model.User, error) {
	return nil, nil
}
func (db *Database) GetUserByID(ctx context.Context, _id string) (*model.User, error) {
	return nil, nil
}
func (db *Database) AddNewUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	return nil, nil
}
func (db *Database) ModifyUser(ctx context.Context, _id string, input *model.NewUser) (*model.User, error) {
	return nil, nil
}
func (db *Database) DeleteUser(ctx context.Context, _id string) (string, error) {
	return "", nil
}

func (db *Database) Login(ctx context.Context, input *model.Login) (*model.UserDto, error) {
	return nil, nil
}
