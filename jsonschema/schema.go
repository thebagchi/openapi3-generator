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
	Object      interface{}
	StringArray []string
)

type SimpleTypes String

const (
	SimpleTypeArray   = SimpleTypes("array")
	SimpleTypeBoolean = SimpleTypes("boolean")
	SimpleTypeInteger = SimpleTypes("integer")
	SimpleTypeNull    = SimpleTypes("null")
	SimpleTypeNumber  = SimpleTypes("number")
	SimpleTypeObject  = SimpleTypes("object")
	SimpleTypeString  = SimpleTypes("string")
)

func (o *SimpleTypes) UnmarshalJSON(data []byte) error {
	var temp String
	err := json.Unmarshal(data, &temp)
	if nil == err {
		switch SimpleTypes(temp) {
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
			*o = SimpleTypes(temp)
		default:
			return errors.New("cannot unmarshal SimpleType")
		}
	}
	return err
}

type SimpleTypesArray []SimpleTypes

type AnyOfSchemaType struct {
	*SimpleTypes
	*SimpleTypesArray
}

func (o *AnyOfSchemaType) UnmarshalJSON(data []byte) error {
	{
		var temp SimpleTypes
		err := json.Unmarshal(data, &temp)
		if nil == err {
			o.SimpleTypes = &temp
		}
	}
	{
		var temp SimpleTypesArray
		err := json.Unmarshal(data, &temp)
		if nil == err {
			o.SimpleTypesArray = &temp
		}
	}
	return nil
}

type AnyOfSchemaBoolean struct {
	*Boolean
	*Schema
}

func (o *AnyOfSchemaBoolean) UnmarshalJSON(data []byte) error {
	{
		var temp Boolean
		err := json.Unmarshal(data, &temp)
		if nil == err {
			o.Boolean = &temp
		}
	}
	{
		var temp Schema
		err := json.Unmarshal(data, &temp)
		if nil == err {
			o.Schema = &temp
		}
	}
	return nil
}

type Schema struct {
	Id                   *String
	Title                *String
	Schema               *String
	Description          *String
	Required             *StringArray
	Default              *Object
	Type                 *AnyOfSchemaType
	AdditionalItems      *AnyOfSchemaBoolean
	AdditionalProperties *AnyOfSchemaBoolean
}

type SchemaArray []Schema

type AnyOfSchemaSchemaArray struct {
	*Schema
	*SchemaArray
}

func (o *AnyOfSchemaSchemaArray) UnmarshalJSON(data []byte) error {
	{
		var temp Schema
		err := json.Unmarshal(data, &temp)
		if nil == err {
			o.Schema = &temp
		}
	}
	{
		var temp SchemaArray
		err := json.Unmarshal(data, &temp)
		if nil == err {
			o.SchemaArray = &temp
		}
	}
	return nil
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
			var temp AnyOfSchemaType
			err := json.Unmarshal(value, &temp)
			if nil == err {
				o.Type = &temp
			}
		}
		if value, ok := temp["default"]; ok {
			var temp Object
			err := json.Unmarshal(value, &temp)
			if nil == err {
				o.Default = &temp
			}
		}
		if value, ok := temp["additionalProperties"]; ok {
			var temp AnyOfSchemaBoolean
			err := json.Unmarshal(value, &temp)
			if nil == err {
				o.AdditionalProperties = &temp
			}
		}
		if value, ok := temp["additionalItems"]; ok {
			var temp AnyOfSchemaBoolean
			err := json.Unmarshal(value, &temp)
			if nil == err {
				o.AdditionalItems = &temp
			}
		}
	}
	return err
}
