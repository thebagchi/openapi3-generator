package jsonschema

import (
	"encoding/json"
	"errors"
)

type (
	Boolean     bool
	Int32       int32
	Int64       int64
	String      string
	StringArray []string
)

type SimpleType String

const (
	SimpleTypeArray   = SimpleType("array")
	SimpleTypeBoolean = SimpleType("boolean")
	SimpleTypeInteger = SimpleType("integer")
	SimpleTypeNull    = SimpleType("null")
	SimpleTypeNumber  = SimpleType("number")
	SimpleTypeObject  = SimpleType("object")
	SimpleTypeString  = SimpleType("string")
)

func (o *SimpleType) UnmarshalJSON(data []byte) error {
	var temp String
	err := json.Unmarshal(data, &temp)
	if nil == err {
		switch SimpleType(temp) {
		case SimpleTypeArray:
			fallthrough
		case SimpleTypeBoolean:
			fallthrough
		case SimpleTypeNull:
			fallthrough
		case SimpleTypeNumber:
			fallthrough
		case SimpleTypeInteger:
			fallthrough
		case SimpleTypeObject:
			fallthrough
		case SimpleTypeString:
			*o = SimpleType(temp)
		default:
			return errors.New("cannot unmarshal SimpleType")
		}
	}
	return err
}

type Schema struct {
	Id          *String
	Title       *String
	Schema      *String
	Description *String
	Required    *StringArray
	Type        *SimpleType
}

func (o *Schema) MarshalJSON() ([]byte, error) {
	return []byte(""), nil
}

func (o *Schema) UnmarshalJSON(data []byte) error {
	var temp map[string]json.RawMessage
	err := json.Unmarshal(data, &temp)
	if nil == err {
		if value, ok := temp["id"]; ok {
			var temp String
			err := json.Unmarshal(value, &temp)
			if nil == err {
				o.Id = &temp
			}
		}
		if value, ok := temp["title"]; ok {
			var temp String
			err := json.Unmarshal(value, &temp)
			if nil == err {
				o.Title = &temp
			}
		}
		if value, ok := temp["$schema"]; ok {
			var temp String
			err := json.Unmarshal(value, &temp)
			if nil == err {
				o.Schema = &temp
			}
		}
		if value, ok := temp["description"]; ok {
			var temp String
			err := json.Unmarshal(value, &temp)
			if nil == err {
				o.Description = &temp
			}
		}
		if value, ok := temp["required"]; ok {
			var temp StringArray
			err := json.Unmarshal(value, &temp)
			if nil == err {
				o.Required = &temp
			}
		}
		if value, ok := temp["type"]; ok {
			var temp SimpleType
			err := json.Unmarshal(value, &temp)
			if nil == err {
				o.Type = &temp
			}
		}
	}
	return err
}
