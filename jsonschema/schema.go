package jsonschema

import (
	"encoding/json"
	"errors"
)

type (
	Boolean     bool
	Int32       int32
	Int64       int64
	Float       float32
	Double      float64
	String      string
	Object      interface{}
	StringArray []string
	ObjectArray []Object
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
	Reference            *String
	Id                   *String
	Title                *String
	Schema               *String
	Description          *String
	Pattern              *String
	MultipleOf           *Double
	Maximum              *Double
	ExclusiveMaximum     *Boolean
	Minimum              *Double
	ExclusiveMinimum     *Boolean
	MaxLength            *Int64
	MinLength            *Int64
	MaxItems             *Int64
	MinItems             *Int64
	UniqueItems          *Boolean
	MaxProperties        *Int64
	MinProperties        *Int64
	Required             *StringArray
	Default              *Object
	Type                 *AnyOfSchemaType
	AdditionalItems      *AnyOfSchemaBoolean
	AdditionalProperties *AnyOfSchemaBoolean
	Items                *AnyOfSchemaSchemaArray
	Definitions          *SchemaDict
	Properties           *SchemaDict
	PatternProperties    *SchemaDict
	AllOf                *SchemaArray
	AnyOf                *SchemaArray
	OneOf                *SchemaArray
	Not                  *Schema
	Enum                 *ObjectArray
	Dependencies         *AnyOfSchemaSchemaArrayDict
}

type AnyOfSchemaSchemaArray struct {
	*Schema
	*SchemaArray
}

type (
	SchemaArray                []Schema
	SchemaDict                 map[string]Schema
	AnyOfSchemaSchemaArrayDict map[string]AnyOfSchemaSchemaArray
)

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
		if value, ok := temp["$ref"]; ok {
			var temp String
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.Reference = &temp
			}
		}
		if value, ok := temp["id"]; ok {
			var temp String
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.Id = &temp
			}
		}
		if value, ok := temp["title"]; ok {
			var temp String
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.Title = &temp
			}
		}
		if value, ok := temp["$schema"]; ok {
			var temp String
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.Schema = &temp
			}
		}
		if value, ok := temp["description"]; ok {
			var temp String
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.Description = &temp
			}
		}
		if value, ok := temp["pattern"]; ok {
			var temp String
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.Pattern = &temp
			}
		}
		if value, ok := temp["required"]; ok {
			var temp StringArray
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.Required = &temp
			}
		}
		if value, ok := temp["type"]; ok {
			var temp AnyOfSchemaType
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.Type = &temp
			}
		}
		if value, ok := temp["default"]; ok {
			var temp Object
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.Default = &temp
			}
		}
		if value, ok := temp["items"]; ok {
			var temp AnyOfSchemaSchemaArray
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.Items = &temp
			}
		}
		if value, ok := temp["additionalProperties"]; ok {
			var temp AnyOfSchemaBoolean
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.AdditionalProperties = &temp
			}
		}
		if value, ok := temp["additionalItems"]; ok {
			var temp AnyOfSchemaBoolean
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.AdditionalItems = &temp
			}
		}
		if value, ok := temp["definitions"]; ok {
			var temp SchemaDict
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.Definitions = &temp
			}
		}
		if value, ok := temp["properties"]; ok {
			var temp SchemaDict
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.Properties = &temp
			}
		}
		if value, ok := temp["patternProperties"]; ok {
			var temp SchemaDict
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.PatternProperties = &temp
			}
		}
		if value, ok := temp["multipleOf"]; ok {
			var temp Double
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.MultipleOf = &temp
			}
		}
		if value, ok := temp["multipleOf"]; ok {
			var temp Double
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.MultipleOf = &temp
			}
		}
		if value, ok := temp["maximum"]; ok {
			var temp Double
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.Maximum = &temp
			}
		}
		if value, ok := temp["exclusiveMaximum"]; ok {
			var temp Boolean
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.ExclusiveMaximum = &temp
			}
		}
		if value, ok := temp["minimum"]; ok {
			var temp Double
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.Minimum = &temp
			}
		}
		if value, ok := temp["exclusiveMinimum"]; ok {
			var temp Boolean
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.ExclusiveMinimum = &temp
			}
		}
		if value, ok := temp["maxLength"]; ok {
			var temp Int64
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.MaxLength = &temp
			}
		}
		if value, ok := temp["minLength"]; ok {
			var temp Int64
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.MinLength = &temp
			}
		}
		if value, ok := temp["maxItems"]; ok {
			var temp Int64
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.MaxItems = &temp
			}
		}
		if value, ok := temp["minItems"]; ok {
			var temp Int64
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.MinItems = &temp
			}
		}
		if value, ok := temp["uniqueItems"]; ok {
			var temp Boolean
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.UniqueItems = &temp
			}
		}
		if value, ok := temp["maxProperties"]; ok {
			var temp Int64
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.MaxProperties = &temp
			}
		}
		if value, ok := temp["minProperties"]; ok {
			var temp Int64
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.MinProperties = &temp
			}
		}
		if value, ok := temp["dependencies"]; ok {
			var temp AnyOfSchemaSchemaArrayDict
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.Dependencies = &temp
			}
		}
		if value, ok := temp["enum"]; ok {
			var temp ObjectArray
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.Enum = &temp
			}
		}
		if value, ok := temp["allOf"]; ok {
			var temp SchemaArray
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.AllOf = &temp
			}
		}
		if value, ok := temp["anyOf"]; ok {
			var temp SchemaArray
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.AnyOf = &temp
			}
		}
		if value, ok := temp["oneOf"]; ok {
			var temp SchemaArray
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.OneOf = &temp
			}
		}
		if value, ok := temp["not"]; ok {
			var temp Schema
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.Not = &temp
			}
		}
	}
	return err
}
