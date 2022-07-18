package user

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"github.com/deyr02/bnzlcrm/graph/model"
	"github.com/deyr02/bnzlcrm/repositories/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	META_USER = "meta_user"
)

type Database struct {
	client *mongo.Client
}
type User_Meta_Repository interface {
	GetUser_MetaCollection(ctx context.Context) *model.MetaUserCollection
	AddNewElement_Meta_User(ctx context.Context, newCustomFieldlement *model.NewCustomFieldElement) *model.MetaUserCollection
	ModifyElement_Meta_User(ctx context.Context, _id string, newCustomFieldElement *model.NewCustomFieldElement) *model.MetaUserCollection
	DeleteElement_Meta_User(ctx context.Context, _id string) *model.MetaUserCollection
}

func New_User_Meta_repository() User_Meta_Repository {
	_client := database.CreateConnection()
	collection := _client.Database(database.DATABASE_NAME).Collection(META_USER)
	_count, _ := collection.CountDocuments(nil, bson.D{})

	if _count == 0 {
		var newcollection model.MetaUserCollection

		for _, element := range getSeedData() {
			ele := &model.CustomFieldElement{
				FieldID:        strconv.Itoa(rand.Int()),
				FieldName:      element.FieldName,
				DataType:       element.DataType,
				FieldType:      element.FieldType,
				IsRequired:     element.IsRequired,
				Visibility:     element.Visibility,
				SystemFieled:   element.SystemFieled,
				MaxValue:       element.MaxValue,
				MinValue:       element.MinValue,
				DefaultValue:   element.DefaultValue,
				PossibleValues: element.PossibleValues,
				FieldOrder:     element.FieldOrder,
			}
			fmt.Println(ele)
			newcollection.Fields = append(newcollection.Fields, ele)
		}
		_, err := collection.InsertOne(nil, newcollection)
		if err != nil {
			panic(err)
		}
	}
	return &Database{
		client: _client,
	}
}

func (db *Database) GetUser_MetaCollection(ctx context.Context) *model.MetaUserCollection {
	collection := db.client.Database(database.DATABASE_NAME).Collection(META_USER)
	cursor := collection.FindOne(ctx, bson.D{})

	var meta_user *model.MetaUserCollection
	err := cursor.Decode(&meta_user)
	if err != nil {
		log.Fatal(err)
	}
	return meta_user
}
func (db *Database) AddNewElement_Meta_User(ctx context.Context, newCustomFieldElement *model.NewCustomFieldElement) *model.MetaUserCollection {
	return nil
}
func (db *Database) ModifyElement_Meta_User(ctx context.Context, _id string, newCustomFieldElement *model.NewCustomFieldElement) *model.MetaUserCollection {
	return nil
}
func (db *Database) DeleteElement_Meta_User(ctx context.Context, _id string) *model.MetaUserCollection {
	return nil
}
