package tests

import (
	"schemToTypes/parser"
	"testing"
)

func TestProcessRequest(t *testing.T) {
	testCases := []struct {
		testName    string
		requestData []byte
		requestType string
		exportType  parser.TypeOption
		name        string
		expected    string
		expectError bool
	}{
		{
			testName: "Test JSON to Golang",
			requestData: []byte(`{
                "$id": "https://example.com/complex-object.schema.json",
                "$schema": "https://json-schema.org/draft/2020-12/schema",
                "title": "Complex Object",
                "type": "object",
                "properties": {
                  "name": {
                    "type": "string"
                  },
                  "age": {
                    "type": "integer",
                    "minimum": 0
                  },
                  "address": {
                    "type": "object",
                    "properties": {
                      "street": {
                        "type": "string"
                      },
                      "city": {
                        "type": "string"
                      },
                      "state": {
                        "type": "string"
                      },
                      "postalCode": {
                        "type": "string",
                        "pattern": "\\d{5}"
                      }
                    },
                    "required": ["street", "city", "state", "postalCode"]
                  },
                  "hobbies": {
                    "type": "array",
                    "items": {
                      "type": "string"
                    }
                  }
                },
                "required": ["name", "age"]
              }
              `),
			requestType: "json",
			exportType:  parser.Golang, // replace with the actual Golang TypeOption
			name:        "Person",
			expected: `type Person struct {
                name: string
                age: number
                Address: Address
                hobbies: string[]
            }
            export type Address = {
                street: string
                city: string
                state: string
                postalCode: string
            }
            
            export type Hobbies = {
            }
            
            
            `,
			expectError: false,
		},
		// Add more test cases here
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := parser.ProcessRequest(tc.requestData, tc.requestType, tc.exportType, tc.name)
			if (err != nil) != tc.expectError {
				t.Fatalf("ProcessRequest() error = %v, expectError %v", err, tc.expectError)
				return
			}
			if result != tc.expected {
				t.Errorf("ProcessRequest() = %v, want %v", result, tc.expected)
			}
		})
	}
}
