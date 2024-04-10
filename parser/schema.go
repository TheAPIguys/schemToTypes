package parser

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/atotto/clipboard"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

// *
// Schema is an interface that defines the methods that a schema should implement to generate types for the schema
// *
type Schema interface {
	SetKey(key string)
	HasKey() bool
	GetRef(exportType TypeOption) string
	IsRequired() bool
	GetType(exportType TypeOption) string
	AddProperty(exportType TypeOption) string
	GenerateTypes(exportType TypeOption) string
}

// *
// StringToTitle converts a string to title case.
// *
func StringToTitle(s string) string {
	// return capitalized key value example name -> Name
	return cases.Title(language.Und, cases.NoLower).String(s)

}

// *
// StringToLower converts a string to lower case.
// *
func StringToLower(s string) string {
	return strings.ToLower(s)
}

// *
// TypeOption is an enum for the type of the export type
// *

type TypeOption string

const (
	Golang     TypeOption = "golang"
	TypeScript TypeOption = "typescript"
)

type YamlType struct {
	Key         *string              `json:"key"`
	Type        string               `json:"type"`
	Properties  *map[string]YamlType `json:"properties"`
	Format      *string              `json:"format"`
	Items       *YamlType            `json:"items"`
	Example     *string              `json:"example"`
	Ref         *string              `json:"$ref"`
	Definitions *map[string]YamlType `json:"definitions"`
}

func (y *YamlType) SetKey(key string) {
	y.Key = &key
}

func (y *YamlType) HasKey() bool {
	return y.Key != nil
}

func (y *YamlType) GetRef(exportType TypeOption) string {
	if y.Ref != nil {
		var strArray = strings.Split(*y.Ref, "/")
		return StringToTitle(strArray[len(strArray)-1])
	}
	return y.GetType(exportType)
}

func (y *YamlType) IsRequired() bool {

	isRequired := true
	if y.Format != nil && *y.Format == "nullable" {
		isRequired = false
	}
	return isRequired
}

func (y *YamlType) GetType(exportType TypeOption) string {
	var t string = ""
	switch exportType {
	case Golang:
		if !y.IsRequired() {
			t = "*"
		}
		switch y.Type {
		case "string":
			return t + "string"
		case "integer":
			return t + "int"
		case "number":
			return t + "float64"
		case "boolean":
			return t + "bool"
		case "array":
			t = t + "[]"
			return t + y.Items.GetRef(exportType)
		default:
			return "interface{}"
		}
	case TypeScript:
		// TO DO
		return ""
	}
	return ""
}
func (y *YamlType) AddProperty(exportType TypeOption) string {
	var text = ""
	var k = *y.Key
	if y.Key == nil {
		k = "Key"
	}
	switch exportType {
	case Golang:
		text = fmt.Sprintf("\t %s \t %s \t `json:\"%s\"`", StringToTitle(k), y.GetType(exportType), StringToLower(*y.Key))
	case TypeScript:
		text = fmt.Sprintf("\t%s: \t %s", StringToTitle(*y.Key), y.GetType(exportType))
	}
	return text

}

func (y *YamlType) GenerateTypes(exportTypeOption TypeOption) string {
	var text = ""

	switch exportTypeOption {
	case Golang:
		text = fmt.Sprintf("type %s struct {", StringToTitle(*y.Key))
		text = text + "\n"
		if y.Properties != nil {
			for key, value := range *y.Properties {
				value.SetKey(key)
				text = text + value.AddProperty(exportTypeOption) + "\n"
			}
		}
		text = text + "}" + "\n"
		if y.Ref != nil {
			for key, value := range *y.Definitions {
				value.SetKey(key)
				text = text + value.GenerateTypes(exportTypeOption)
			}
		}
	case TypeScript:
		text = fmt.Sprintf("export type %s = {", StringToTitle(*y.Key))
		text = text + "\n"
		if y.Properties != nil {
			for key, value := range *y.Properties {
				value.SetKey(key)
				text = text + value.AddProperty(exportTypeOption) + "\n"
			}
		}
		text = text + "}" + "\n"

	}
	return text
}

type JsonType struct {
	Key        *string              `json:"key"`
	Type       string               `json:"type"`
	Properties *map[string]JsonType `json:"properties"`
	Required   *[]string            `json:"required"`
	Items      *JsonType            `json:"items"`
	Defs       *map[string]JsonType `json:"$defs"`
}

func (j *JsonType) SetKey(key string) {
	j.Key = &key
}

func (j *JsonType) HasKey() bool {
	return j.Key != nil
}

func (j *JsonType) GetRef(exportType TypeOption) string {
	if j.Key != nil {
		return StringToTitle(*j.Key)
	}
	return j.GetType(exportType)
}

func (j *JsonType) IsRequired() bool {
	var required bool = false

	if j.Required != nil {
		for _, value := range *j.Required {
			if j.Key != nil && *j.Key == value {
				required = true
			}
		}
	}
	return required
}

func (j *JsonType) GetType(exportType TypeOption) string {
	var t string = ""
	switch exportType {
	case Golang:
		if !j.IsRequired() {
			t = "*"
		}
		switch j.Type {
		case "string":
			return t + "string"
		case "integer":
			return t + "int"
		case "number":
			return t + "float64"
		case "boolean":
			return t + "bool"
		case "array":
			t = t + "[]"
			return t + j.Items.GetRef(exportType)
		case "object":
			if j.Key != nil {
				return t + StringToTitle(*j.Key)
			}
			return t + "interface{}"
		default:
			return "interface{}"
		}
	case TypeScript:
		// TO DO
		return ""
	}
	return ""
}
func (j *JsonType) AddProperty(exportType TypeOption) string {
	var text = ""
	switch exportType {
	case Golang:
		text = fmt.Sprintf("\t %s \t %s \t `json:\"%s\"`", StringToTitle(*j.Key), j.GetType(exportType), StringToLower(*j.Key))
	case TypeScript:
		text = fmt.Sprintf("\t%s: \t %s", StringToTitle(*j.Key), j.GetType(exportType))
	}
	return text

}

func (j *JsonType) GenerateTypes(exportTypeOption TypeOption) string {
	var text string
	if j.Defs == nil {
		j.Defs = &map[string]JsonType{}
	}
	switch exportTypeOption {
	case Golang:
		text = fmt.Sprintf("type %s struct {", StringToTitle(*j.Key))
		text += "\n"
		if j.Properties != nil {
			for key, value := range *j.Properties {
				value.SetKey(key)
				text += value.AddProperty(exportTypeOption) + "\n"
				if value.Items != nil {
					(*j.Defs)[key] = *value.Items
					recursiveAddDefs(j.Defs, value.Items)
				}
			}
		}
		text += "}" + "\n"

		if j.Defs != nil {
			for key, value := range *j.Defs {
				value.SetKey(key)
				text += value.GenerateTypes(exportTypeOption)
			}
		}
	case TypeScript:
		text = fmt.Sprintf("export type %s = {", StringToTitle(*j.Key))
		text += "\n"
		if j.Properties != nil {
			for key, value := range *j.Properties {
				value.SetKey(key)
				text += value.AddProperty(exportTypeOption) + "\n"
			}
		}
		text += "}" + "\n"
	}

	return text
}

// *
// recursiveAddDefs is a helper function to recursively add definitions to the definitions map, defs is a map of definitions
// and item is the current item being processed
// *
func recursiveAddDefs(defs *map[string]JsonType, item *JsonType) {
	if item.Properties != nil {
		for k, v := range *item.Properties {
			v.SetKey(k)
			if v.Type == "object" {
				(*defs)[k] = v
				recursiveAddDefs(defs, &v)
			}
		}
	}
}

// *
// ProcessRequest processes the request data and returns the generated types as a string
// requestType can be either "yaml" or "json"
// exportType can be either Golang or TypeScript
// *
func ProcessRequest(requestData []byte, requestType string, exportType TypeOption, name string) (string, error) {
	var schema Schema
	var err error

	// Parse the request data based on the request type
	switch requestType {
	case "yml":
		var yamlSchema YamlType
		err = yaml.Unmarshal(requestData, &yamlSchema)
		schema = &yamlSchema
	case "json":
		var jsonSchema JsonType
		err = json.Unmarshal(requestData, &jsonSchema)
		schema = &jsonSchema
	default:
		return "", fmt.Errorf("unsupported request type")
	}

	if err != nil {
		return "", err
	}
	schema.SetKey(name)
	// Generate types based on the export type
	return schema.GenerateTypes(exportType), nil
}

func SendToClipboard(text string) {

	errClip := clipboard.WriteAll(text)
	if errClip != nil {
		log.Fatalf("error: %v", errClip)
		return
	}
	fmt.Println("The struct has been generated successfully! and its saved in your clipboard. Paste it in your code editor.")
}
