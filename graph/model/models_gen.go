// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type CustomFieldElement struct {
	FieldID        string    `json:"FieldID"`
	FieldName      string    `json:"FieldName"`
	DataType       DataType  `json:"DataType"`
	FieldType      FieldType `json:"FieldType"`
	IsRequired     bool      `json:"IsRequired"`
	Visibility     bool      `json:"Visibility"`
	SystemFieled   bool      `json:"SystemFieled"`
	MaxValue       *int      `json:"MaxValue"`
	MinValue       *int      `json:"MinValue"`
	DefaultValue   string    `json:"DefaultValue"`
	PossibleValues []*string `json:"PossibleValues"`
	FieldOrder     int       `json:"FieldOrder"`
}

type ElementValue struct {
	Key      string   `json:"key"`
	DataType DataType `json:"DataType"`
	Value    string   `json:"value"`
}

type MetaUserCollection struct {
	Fields []*CustomFieldElement `json:"Fields"`
}

type NewCollection struct {
	Fields []*NewCustomFieldElement `json:"Fields"`
}

type NewCustomFieldElement struct {
	FieldName      string    `json:"FieldName"`
	DataType       DataType  `json:"DataType"`
	FieldType      FieldType `json:"FieldType"`
	IsRequired     bool      `json:"IsRequired"`
	Visibility     bool      `json:"Visibility"`
	SystemFieled   bool      `json:"SystemFieled"`
	MaxValue       *int      `json:"MaxValue"`
	MinValue       *int      `json:"MinValue"`
	DefaultValue   string    `json:"DefaultValue"`
	PossibleValues []*string `json:"PossibleValues"`
	FieldOrder     int       `json:"FieldOrder"`
}

type NewElementValue struct {
	Key      string   `json:"key"`
	DataType DataType `json:"DataType"`
	Value    string   `json:"value"`
}

type NewUserCollection struct {
	Data []*NewElementValue `json:"data"`
}

type NewUserRole struct {
	RoleName   string      `json:"RoleName"`
	Operations []Operation `json:"Operations"`
	SystemRole *bool       `json:"SystemRole"`
}

type UserCollection struct {
	RecordID string          `json:"RecordID"`
	Data     []*ElementValue `json:"data"`
}

type UserRole struct {
	Role       string      `json:"Role"`
	RoleName   string      `json:"RoleName"`
	Operations []Operation `json:"Operations"`
	SystemRole bool        `json:"SystemRole"`
}

type DataType string

const (
	DataTypeBoolean DataType = "Boolean"
	DataTypeDate    DataType = "Date"
	DataTypeDouble  DataType = "Double"
	DataTypeInteger DataType = "Integer"
	DataTypeString  DataType = "String"
)

var AllDataType = []DataType{
	DataTypeBoolean,
	DataTypeDate,
	DataTypeDouble,
	DataTypeInteger,
	DataTypeString,
}

func (e DataType) IsValid() bool {
	switch e {
	case DataTypeBoolean, DataTypeDate, DataTypeDouble, DataTypeInteger, DataTypeString:
		return true
	}
	return false
}

func (e DataType) String() string {
	return string(e)
}

func (e *DataType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DataType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DataType", str)
	}
	return nil
}

func (e DataType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type FieldType string

const (
	FieldTypeTextBox        FieldType = "TextBox"
	FieldTypeRadioButtons   FieldType = "RadioButtons"
	FieldTypeCheckBox       FieldType = "CheckBox"
	FieldTypeCheckBoxes     FieldType = "CheckBoxes"
	FieldTypeDropDownList   FieldType = "DropDownList"
	FieldTypeInputFile      FieldType = "InputFile"
	FieldTypeDatePicker     FieldType = "DatePicker"
	FieldTypeDateTimePicker FieldType = "DateTimePicker"
)

var AllFieldType = []FieldType{
	FieldTypeTextBox,
	FieldTypeRadioButtons,
	FieldTypeCheckBox,
	FieldTypeCheckBoxes,
	FieldTypeDropDownList,
	FieldTypeInputFile,
	FieldTypeDatePicker,
	FieldTypeDateTimePicker,
}

func (e FieldType) IsValid() bool {
	switch e {
	case FieldTypeTextBox, FieldTypeRadioButtons, FieldTypeCheckBox, FieldTypeCheckBoxes, FieldTypeDropDownList, FieldTypeInputFile, FieldTypeDatePicker, FieldTypeDateTimePicker:
		return true
	}
	return false
}

func (e FieldType) String() string {
	return string(e)
}

func (e *FieldType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FieldType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FieldType", str)
	}
	return nil
}

func (e FieldType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Operation string

const (
	OperationGet  Operation = "GET"
	OperationPost Operation = "POST"
	OperationPut  Operation = "PUT"
	OperationDel  Operation = "DEL"
)

var AllOperation = []Operation{
	OperationGet,
	OperationPost,
	OperationPut,
	OperationDel,
}

func (e Operation) IsValid() bool {
	switch e {
	case OperationGet, OperationPost, OperationPut, OperationDel:
		return true
	}
	return false
}

func (e Operation) String() string {
	return string(e)
}

func (e *Operation) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Operation(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid operation", str)
	}
	return nil
}

func (e Operation) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
