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
	StringArray []*String
	ObjectArray []*Object
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

type SimpleTypesArray []*SimpleTypes

type AnyOfSchemaType struct {
	*SimpleTypes
	*SimpleTypesArray
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
	SchemaArray                []*Schema
	SchemaDict                 map[string]*Schema
	AnyOfSchemaSchemaArrayDict map[string]*AnyOfSchemaSchemaArray
)

type AnyOfSchemaSchemaArray struct {
	*Schema
	*SchemaArray
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

func (o *Schema) MarshalJSON() ([]byte, error) {
	var temp = make(map[string]json.RawMessage)
	if nil != o.Reference {
		data, err := json.Marshal(o.Reference)
		if nil == err {
			temp["$ref"] = data
		}
	}
	if nil != o.Id {
		data, err := json.Marshal(o.Id)
		if nil == err {
			temp["id"] = data
		}
	}
	if nil != o.Title {
		data, err := json.Marshal(o.Title)
		if nil == err {
			temp["title"] = data
		}
	}
	if nil != o.Schema {
		data, err := json.Marshal(o.Schema)
		if nil == err {
			temp["$schema"] = data
		}
	}
	if nil != o.Description {
		data, err := json.Marshal(o.Description)
		if nil == err {
			temp["description"] = data
		}
	}
	if nil != o.Pattern {
		data, err := json.Marshal(o.Pattern)
		if nil == err {
			temp["pattern"] = data
		}
	}
	if nil != o.Required {
		data, err := json.Marshal(o.Required)
		if nil == err {
			temp["required"] = data
		}
	}
	if nil != o.Type {
		data, err := json.Marshal(o.Type)
		if nil == err {
			temp["type"] = data
		}
	}
	if nil != o.Default {
		data, err := json.Marshal(o.Default)
		if nil == err {
			temp["default"] = data
		}
	}
	if nil != o.Items {
		data, err := json.Marshal(o.Items)
		if nil == err {
			temp["items"] = data
		}
	}
	if nil != o.AdditionalProperties {
		data, err := json.Marshal(o.AdditionalProperties)
		if nil == err {
			temp["additionalProperties"] = data
		}
	}
	if nil != o.AdditionalItems {
		data, err := json.Marshal(o.AdditionalItems)
		if nil == err {
			temp["additionalItems"] = data
		}
	}
	if nil != o.Definitions {
		data, err := json.Marshal(o.Definitions)
		if nil == err {
			temp["definitions"] = data
		}
	}
	if nil != o.Properties {
		data, err := json.Marshal(o.Properties)
		if nil == err {
			temp["properties"] = data
		}
	}
	if nil != o.PatternProperties {
		data, err := json.Marshal(o.PatternProperties)
		if nil == err {
			temp["patternProperties"] = data
		}
	}
	if nil != o.MultipleOf {
		data, err := json.Marshal(o.MultipleOf)
		if nil == err {
			temp["multipleOf"] = data
		}
	}
	if nil != o.Maximum {
		data, err := json.Marshal(o.Maximum)
		if nil == err {
			temp["maximum"] = data
		}
	}
	if nil != o.ExclusiveMaximum {
		data, err := json.Marshal(o.ExclusiveMaximum)
		if nil == err {
			temp["exclusiveMaximum"] = data
		}
	}
	if nil != o.ExclusiveMinimum {
		data, err := json.Marshal(o.ExclusiveMinimum)
		if nil == err {
			temp["exclusiveMinimum"] = data
		}
	}
	if nil != o.Minimum {
		data, err := json.Marshal(o.Minimum)
		if nil == err {
			temp["minimum"] = data
		}
	}
	if nil != o.ExclusiveMinimum {
		data, err := json.Marshal(o.ExclusiveMinimum)
		if nil == err {
			temp["exclusiveMinimum"] = data
		}
	}
	if nil != o.MaxLength {
		data, err := json.Marshal(o.MaxLength)
		if nil == err {
			temp["maxLength"] = data
		}
	}
	if nil != o.MinLength {
		data, err := json.Marshal(o.MinLength)
		if nil == err {
			temp["minLength"] = data
		}
	}
	if nil != o.MaxItems {
		data, err := json.Marshal(o.MaxItems)
		if nil == err {
			temp["maxItems"] = data
		}
	}
	if nil != o.MinItems {
		data, err := json.Marshal(o.MinItems)
		if nil == err {
			temp["minItems"] = data
		}
	}
	if nil != o.UniqueItems {
		data, err := json.Marshal(o.UniqueItems)
		if nil == err {
			temp["uniqueItems"] = data
		}
	}
	if nil != o.MaxProperties {
		data, err := json.Marshal(o.MaxProperties)
		if nil == err {
			temp["maxProperties"] = data
		}
	}
	if nil != o.MinProperties {
		data, err := json.Marshal(o.MinProperties)
		if nil == err {
			temp["minProperties"] = data
		}
	}
	if nil != o.Dependencies {
		data, err := json.Marshal(o.Dependencies)
		if nil == err {
			temp["dependencies"] = data
		}
	}
	if nil != o.Enum {
		data, err := json.Marshal(o.Enum)
		if nil == err {
			temp["enum"] = data
		}
	}
	if nil != o.AllOf {
		data, err := json.Marshal(o.AllOf)
		if nil == err {
			temp["allOf"] = data
		}
	}
	if nil != o.AnyOf {
		data, err := json.Marshal(o.AnyOf)
		if nil == err {
			temp["anyOf"] = data
		}
	}
	if nil != o.OneOf {
		data, err := json.Marshal(o.OneOf)
		if nil == err {
			temp["oneOf"] = data
		}
	}
	if nil != o.Not {
		data, err := json.Marshal(o.Not)
		if nil == err {
			temp["not"] = data
		}
	}
	if len(temp) > 0 {
		return json.Marshal(&temp)
	}
	return nil, errors.New("failed marshalling Schema")
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
		} else {
			o.Reference = nil
		}
		if value, ok := temp["id"]; ok {
			var temp String
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.Id = &temp
			}
		} else {
			o.Id = nil
		}
		if value, ok := temp["title"]; ok {
			var temp String
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.Title = &temp
			}
		} else {
			o.Title = nil
		}
		if value, ok := temp["$schema"]; ok {
			var temp String
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.Schema = &temp
			}
		} else {
			o.Schema = nil
		}
		if value, ok := temp["description"]; ok {
			var temp String
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.Description = &temp
			}
		} else {
			o.Description = nil
		}
		if value, ok := temp["pattern"]; ok {
			var temp String
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.Pattern = &temp
			}
		} else {
			o.Pattern = nil
		}
		if value, ok := temp["required"]; ok {
			var temp StringArray
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.Required = &temp
			}
		} else {
			o.Required = nil
		}
		if value, ok := temp["type"]; ok {
			var temp AnyOfSchemaType
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.Type = &temp
			}
		} else {
			o.Type = nil
		}
		if value, ok := temp["default"]; ok {
			var temp Object
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.Default = &temp
			}
		} else {
			o.Default = nil
		}
		if value, ok := temp["items"]; ok {
			var temp AnyOfSchemaSchemaArray
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.Items = &temp
			}
		} else {
			o.Items = nil
		}
		if value, ok := temp["additionalProperties"]; ok {
			var temp AnyOfSchemaBoolean
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.AdditionalProperties = &temp
			}
		} else {
			o.AdditionalProperties = nil
		}
		if value, ok := temp["additionalItems"]; ok {
			var temp AnyOfSchemaBoolean
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.AdditionalItems = &temp
			}
		} else {
			o.AdditionalItems = nil
		}
		if value, ok := temp["definitions"]; ok {
			var temp SchemaDict
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.Definitions = &temp
			}
		} else {
			o.Definitions = nil
		}
		if value, ok := temp["properties"]; ok {
			var temp SchemaDict
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.Properties = &temp
			}
		} else {
			o.Properties = nil
		}
		if value, ok := temp["patternProperties"]; ok {
			var temp SchemaDict
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.PatternProperties = &temp
			}
		} else {
			o.PatternProperties = nil
		}
		if value, ok := temp["multipleOf"]; ok {
			var temp Double
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.MultipleOf = &temp
			}
		} else {
			o.MultipleOf = nil
		}
		if value, ok := temp["maximum"]; ok {
			var temp Double
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.Maximum = &temp
			}
		} else {
			o.Maximum = nil
		}
		if value, ok := temp["exclusiveMaximum"]; ok {
			var temp Boolean
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.ExclusiveMaximum = &temp
			}
		} else {
			o.ExclusiveMaximum = nil
		}
		if value, ok := temp["minimum"]; ok {
			var temp Double
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.Minimum = &temp
			}
		} else {
			o.Minimum = nil
		}
		if value, ok := temp["exclusiveMinimum"]; ok {
			var temp Boolean
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.ExclusiveMinimum = &temp
			}
		} else {
			o.ExclusiveMinimum = nil
		}
		if value, ok := temp["maxLength"]; ok {
			var temp Int64
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.MaxLength = &temp
			}
		} else {
			o.MaxLength = nil
		}
		if value, ok := temp["minLength"]; ok {
			var temp Int64
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.MinLength = &temp
			}
		} else {
			o.MinLength = nil
		}
		if value, ok := temp["maxItems"]; ok {
			var temp Int64
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.MaxItems = &temp
			}
		} else {
			o.MaxItems = nil
		}
		if value, ok := temp["minItems"]; ok {
			var temp Int64
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.MinItems = &temp
			}
		} else {
			o.MinItems = nil
		}
		if value, ok := temp["uniqueItems"]; ok {
			var temp Boolean
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.UniqueItems = &temp
			}
		} else {
			o.UniqueItems = nil
		}
		if value, ok := temp["maxProperties"]; ok {
			var temp Int64
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.MaxProperties = &temp
			}
		} else {
			o.MaxProperties = nil
		}
		if value, ok := temp["minProperties"]; ok {
			var temp Int64
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.MinProperties = &temp
			}
		} else {
			o.MinProperties = nil
		}
		if value, ok := temp["dependencies"]; ok {
			var temp AnyOfSchemaSchemaArrayDict
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.Dependencies = &temp
			}
		} else {
			o.Dependencies = nil
		}
		if value, ok := temp["enum"]; ok {
			var temp ObjectArray
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.Enum = &temp
			}
		} else {
			o.Enum = nil
		}
		if value, ok := temp["allOf"]; ok {
			var temp SchemaArray
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.AllOf = &temp
			}
		} else {
			o.AllOf = nil
		}
		if value, ok := temp["anyOf"]; ok {
			var temp SchemaArray
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.AnyOf = &temp
			}
		} else {
			o.AnyOf = nil
		}
		if value, ok := temp["oneOf"]; ok {
			var temp SchemaArray
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.OneOf = &temp
			}
		} else {
			o.OneOf = nil
		}
		if value, ok := temp["not"]; ok {
			var temp Schema
			err = json.Unmarshal(value, &temp)
			if nil == err {
				o.Not = &temp
			}
		} else {
			o.Not = nil
		}
	}
	return err
}
