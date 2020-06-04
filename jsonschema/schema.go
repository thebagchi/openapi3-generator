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

func (o *SimpleTypes) MarshalJSON() ([]byte, error) {
	switch *o {
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
		temp := string(*o)
		return json.Marshal(temp)
	default:
	}
	return nil, errors.New("failed marshalling SimpleTypes")
}

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
			return errors.New("failed unmarshalling SimpleTypes")
		}
	}
	return err
}

type SimpleTypesArray []SimpleTypes

type AnyOfSchemaType struct {
	*SimpleTypes
	*SimpleTypesArray
}

func (o *AnyOfSchemaType) IsSimpleType() bool {
	if nil != o.SimpleTypes {
		return true
	}
	return false
}

func (o *AnyOfSchemaType) IsSimpleTypesArray() bool {
	if nil != o.SimpleTypesArray {
		return true
	}
	return false
}

func (o *AnyOfSchemaType) MarshalJSON() ([]byte, error) {
	if nil != o.SimpleTypes {
		return json.Marshal(o.SimpleTypes)
	}
	if nil != o.SimpleTypesArray {
		return json.Marshal(o.SimpleTypesArray)
	}
	return nil, errors.New("failed marshalling AnyOfSchemaType")
}

func (o *AnyOfSchemaType) UnmarshalJSON(data []byte) error {
	errored := true
	{
		var temp SimpleTypes
		err := json.Unmarshal(data, &temp)
		if nil == err {
			o.SimpleTypes = &temp
			errored = false
		}
	}
	{
		var temp SimpleTypesArray
		err := json.Unmarshal(data, &temp)
		if nil == err {
			o.SimpleTypesArray = &temp
			errored = false
		}
	}
	if errored {
		return errors.New("failed unmarshalling AnyOfSchemaType")
	}
	return nil
}

type AnyOfSchemaBoolean struct {
	*Boolean
	*Schema
}

func (o *AnyOfSchemaBoolean) IsBoolean() bool {
	if nil != o.Boolean {
		return true
	}
	return false
}

func (o *AnyOfSchemaBoolean) IsSchema() bool {
	if nil != o.Schema {
		return true
	}
	return false
}

func (o *AnyOfSchemaBoolean) MarshalJSON() ([]byte, error) {
	if nil != o.Boolean {
		return json.Marshal(o.Boolean)
	}
	if nil != o.Schema {
		return json.Marshal(o.Schema)
	}
	return nil, errors.New("failed marshalling AnyOfSchemaBoolean")
}

func (o *AnyOfSchemaBoolean) UnmarshalJSON(data []byte) error {
	errored := true
	{
		var temp Boolean
		err := json.Unmarshal(data, &temp)
		if nil == err {
			o.Boolean = &temp
			errored = false
		}
	}
	{
		var temp Schema
		err := json.Unmarshal(data, &temp)
		if nil == err {
			o.Schema = &temp
			errored = false
		}
	}
	if errored {
		return errors.New("failed unmarshalling AnyOfSchemaBoolean")
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

type (
	SchemaArray                []Schema
	SchemaDict                 map[string]Schema
	AnyOfSchemaSchemaArrayDict map[string]AnyOfSchemaSchemaArray
)

type AnyOfSchemaSchemaArray struct {
	*Schema
	*SchemaArray
}

func (o *AnyOfSchemaSchemaArray) IsSchema() bool {
	if nil != o.Schema {
		return true
	}
	return false
}

func (o *AnyOfSchemaSchemaArray) IsSchemaArray() bool {
	if nil != o.SchemaArray {
		return true
	}
	return false
}

func (o *AnyOfSchemaSchemaArray) MarshalJSON() ([]byte, error) {
	if nil != o.Schema {
		return json.Marshal(o.Schema)
	}
	if nil != o.SchemaArray {
		return json.Marshal(o.SchemaArray)
	}
	return nil, errors.New("failed marshalling AnyOfSchemaSchemaArray")
}

func (o *AnyOfSchemaSchemaArray) UnmarshalJSON(data []byte) error {
	errored := true
	{
		var temp Schema
		err := json.Unmarshal(data, &temp)
		if nil == err {
			o.Schema = &temp
			errored = false
		}
	}
	{
		var temp SchemaArray
		err := json.Unmarshal(data, &temp)
		if nil == err {
			o.SchemaArray = &temp
			errored = false
		}
	}
	if errored {
		return errors.New("failed unmarshalling AnyOfSchemaBoolean")
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
