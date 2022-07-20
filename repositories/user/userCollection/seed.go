package userCollection

import (
	"github.com/deyr02/bnzlcrm/graph/model"
)

var adminUser = &model.NewUser{
	RoleID:   "1111111111111111111",
	UserName: "bnzl_admin",
	Email:    "admin@bnzl.com",
	Password: "Password",
	Properties: []*model.NewElementValue{
		&model.NewElementValue{
			Key:      "FirstName",
			DataType: model.DataTypeString,
			Value:    "Admin",
		},
		&model.NewElementValue{
			Key:      "LastName",
			DataType: model.DataTypeString,
			Value:    "Admin",
		},
		&model.NewElementValue{
			Key:      "email",
			DataType: model.DataTypeString,
			Value:    "admin@bnzl.com",
		},
	},
}
