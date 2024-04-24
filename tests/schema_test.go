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
        "$schema": "http://json-schema.org/draft-07/schema#",
        "title": "salesOrderResponse",
        "type": "object",
        "properties": {
          "status": {
            "type": "string"
          },
                    "message": {},
                    "totalResult": {
                      "type": "number"
                    },
                    "salesOrders": {
                      "type": "array",
                      "items": {
                        "type": "object",
                        "properties": {
                          "id": {
                            "type": "number"
                          },
                          "invoiceDate": {
                            "type": "number"
                          },
                          "invoiceDateAsText": {
                            "type": "string"
                          },
                          "customerId": {
                            "type": "number"
                          },
                          "customerName": {
                            "type": "string"
                          },
                          "sendTo": {
                            "type": "object",
                            "properties": {
                              "id": {},
                              "primeName": {
                                "type": "string"
                              },
                              "givenName": {},
                              "phone": {
                                "type": "string"
                              },
                              "email": {},
                              "address": {
                                "type": "object",
                                "properties": {
                                  "street1": {
                                    "type": "string"
                                  },
                                  "street2": {
                                    "type": "string"
                                  },
                                  "city": {
                                    "type": "string"
                                  },
                                  "state": {
                                    "type": "string"
                                  },
                                  "postalCode": {
                                    "type": "string"
                                  },
                                  "country": {
                                    "type": "string"
                                  }
                                },
                                "required": [
                                  "street1",
                                  "street2"
                                ]
                              },
                              "isOrganization": {}
                            },
                            "required": [
                              "id",
                              "primeName",
                              "givenName",
                              "email",
                              "address",
                              "isOrganization"
                            ]
                          },
                          "salesType": {
                            "type": "string"
                          },
                          "salesPriceListId": {
                            "type": "number"
                          },
                          "salesPriceListName": {
                            "type": "string"
                          },
                          "priceDetails": {
                            "type": "object",
                            "properties": {
                              "countryCurrencyCode": {
                                "type": "string"
                              },
                              "taxPolicy": {
                                "type": "string"
                              }
                            },
                            "required": [
                              "countryCurrencyCode",
                              "taxPolicy"
                            ]
                          },
                          "salesOrderStatus": {
                            "type": "string"
                          },
                          "salesOrderItems": {
                            "type": "array",
                            "items": {
                              "type": "object",
                              "properties": {
                                "id": {
                                  "type": "number"
                                },
                                "itemId": {
                                  "type": "number"
                                },
                                "itemName": {
                                  "type": "string"
                                },
                                "unitPrice": {
                                  "type": "number"
                                },
                                "quantity": {
                                  "type": "number"
                                },
                                "unitOfMeasure": {
                                  "type": "string"
                                },
                                "discountPct": {
                                  "type": "number"
                                },
                                "adjustment": {
                                  "type": "number"
                                },
                                "taxAmount": {
                                  "type": "number"
                                },
                                "lineTotal": {
                                  "type": "number"
                                },
                                "accountId": {
                                  "type": "number"
                                },
                                "accountCode": {
                                  "type": "string"
                                },
                                "taxRateId": {
                                  "type": "number"
                                },
                                "taxRateName": {
                                  "type": "string"
                                },
                                "sku": {},
                                "externalReference": {}
                              },
                              "required": [
                                "id",
                                "itemId",
                                "itemName",
                                "unitPrice",
                                "quantity",
                                "unitOfMeasure",
                                "taxAmount",
                                "lineTotal",
                                "accountId",
                                "accountCode",
                                "taxRateId",
                                "taxRateName",
                                "sku",
                                "externalReference"
                              ]
                            }
                          },
                          "code": {
                            "type": "string"
                          },
                          "description": {},
                          "reference": {
                            "type": "string"
                          },
                          "orderDate": {
                            "type": "number"
                          },
                          "orderDateAsText": {
                            "type": "string"
                          },
                          "wineryId": {
                            "type": "number"
                          },
                          "wineryName": {
                            "type": "string"
                          },
                          "fulfillment": {
                            "type": "string"
                          },
                          "fulfillmentDate": {
                            "type": "number"
                          },
                          "fulfillmentDateAsText": {
                            "type": "string"
                          },
                          "salesRegionId": {
                            "type": "number"
                          },
                          "salesRegionCode": {
                            "type": "string"
                          },
                          "notes": {
                            "type": "string"
                          },
                          "customerPickup": {
                            "type": "boolean"
                          },
                          "disableAccountsSync": {
                            "type": "boolean"
                          },
                          "subTotal": {
                            "type": "number"
                          },
                          "taxBreakdown": {
                            "type": "array",
                            "items": {
                              "type": "object",
                              "properties": {
                                "name": {
                                  "type": "string"
                                },
                                "amount": {
                                  "type": "number"
                                },
                                "ratePct": {
                                  "type": "number"
                                },
                                "inclusive": {
                                  "type": "boolean"
                                }
                              },
                              "required": [
                                "name",
                                "amount",
                                "ratePct",
                                "inclusive"
                              ]
                            }
                          },
                          "total": {
                            "type": "number"
                          },
                          "acctReference": {
                            "type": "string"
                          },
                          "posSaleReference": {
                            "type": "array",
                            "items": {}
                          },
                          "storageAreaId": {},
                          "storageAreaCode": {},
                          "externalTransactionId": {},
                          "ignoreStockError": {
                            "type": "boolean"
                          },
                          "posSaleDate": {},
                          "unlinkedItemsPresent": {
                            "type": "boolean"
                          }
                        },
                        "required": [
                          "id",
                          "invoiceDateAsText",
                          "customerId",
                          "customerName",
                          "sendTo",
                          "salesType",
                          "salesPriceListId",
                          "salesPriceListName",
                          "priceDetails",
                          "salesOrderStatus",
                          "salesOrderItems",
                          "code",
                          "description",
                          "orderDate",
                          "orderDateAsText",
                          "wineryId",
                          "wineryName",
                          "fulfillment",
                          "fulfillmentDateAsText",
                          "salesRegionId",
                          "salesRegionCode",
                          "notes",
                          "customerPickup",
                          "disableAccountsSync",
                          "subTotal",
                          "taxBreakdown",
                          "total",
                          "posSaleReference",
                          "storageAreaId",
                          "storageAreaCode",
                          "externalTransactionId",
                          "ignoreStockError",
                          "posSaleDate",
                          "unlinkedItemsPresent"
                        ]
                      }
                    }
                  },
                  "required": [
                    "status",
                    "message",
                    "totalResult",
                    "salesOrders"
                  ]
                }`),
			requestType: "json",
			exportType:  parser.Golang, // replace with the actual Golang TypeOption
			name:        "salesOrderResponse",
			expected: `type SalesOrderResponse struct {
              Status                string        ` + "`json:\"status\"`" + `
              Message               interface{}   ` + "`json:\"message\"`" + `
              TotalResult           float64       ` + "`json:\"totalResult\"`" + `
              SalesOrders           []SalesOrder  ` + "`json:\"salesOrders\"`" + `
            }
            type SalesOrder struct {
              Id                    float64       ` + "`json:\"id\"`" + `
              InvoiceDate           float64       ` + "`json:\"invoiceDate\"`" + `
              InvoiceDateAsText     string        ` + "`json:\"invoiceDateAsText\"`" + `
              CustomerId            float64       ` + "`json:\"customerId\"`" + `
              CustomerName          string        ` + "`json:\"customerName\"`" + `
              SendTo                SendTo        ` + "`json:\"sendTo\"`" + `
              SalesType             string        ` + "`json:\"salesType\"`" + `
              SalesPriceListId      float64       ` + "`json:\"salesPriceListId\"`" + `
              SalesPriceListName    string        ` + "`json:\"salesPriceListName\"`" + `
              PriceDetails          PriceDetails  ` + "`json:\"priceDetails\"`" + `
              SalesOrderStatus      string        ` + "`json:\"salesOrderStatus\"`" + `
              SalesOrderItems       []SalesOrderItem ` + "`json:\"salesOrderItems\"`" + `
              Code                  string        ` + "`json:\"code\"`" + `
              Description           interface{}   ` + "`json:\"description\"`" + `
              Reference             string        ` + "`json:\"reference\"`" + `
              OrderDate             float64       ` + "`json:\"orderDate\"`" + `
              OrderDateAsText       string        ` + "`json:\"orderDateAsText\"`" + `
              WineryId              float64       ` + "`json:\"wineryId\"`" + `
              WineryName            string        ` + "`json:\"wineryName\"`" + `
              Fulfillment           string        ` + "`json:\"fulfillment\"`" + `
              FulfillmentDate       *float64      ` + "`json:\"fulfillmentDate\"`" + `
              FulfillmentDateAsText string        ` + "`json:\"fulfillmentDateAsText\"`" + `
              SalesRegionId         float64       ` + "`json:\"salesRegionId\"`" + `
              SalesRegionCode       string        ` + "`json:\"salesRegionCode\"`" + `
              Notes                 string        ` + "`json:\"notes\"`" + `
              CustomerPickup        bool          ` + "`json:\"customerPickup\"`" + `
              DisableAccountsSync   bool          ` + "`json:\"disableAccountsSync\"`" + `
              SubTotal              float64       ` + "`json:\"subTotal\"`" + `
              TaxBreakdown          []TaxBreakdown ` + "`json:\"taxBreakdown\"`" + `
              Total                 float64       ` + "`json:\"total\"`" + `
              AcctReference         string        ` + "`json:\"acctReference\"`" + `
              PosSaleReference      []interface{} ` + "`json:\"posSaleReference\"`" + `
              StorageAreaId         interface{}   ` + "`json:\"storageAreaId\"`" + `
              StorageAreaCode       interface{}   ` + "`json:\"storageAreaCode\"`" + `
              ExternalTransactionId interface{}   ` + "`json:\"externalTransactionId\"`" + `
              IgnoreStockError      bool          ` + "`json:\"ignoreStockError\"`" + `
              PosSaleDate           interface{}   ` + "`json:\"posSaleDate\"`" + `
              UnlinkedItemsPresent  bool          ` + "`json:\"unlinkedItemsPresent\"`" + `
            }
            type SendTo struct {
              Id              interface{} ` + "`json:\"id\"`" + `
              PrimeName       string      ` + "`json:\"primeName\"`" + `
              GivenName       interface{} ` + "`json:\"givenName\"`" + `
              Phone           string      ` + "`json:\"phone\"`" + `
              Email           interface{} ` + "`json:\"email\"`" + `
              Address         Address     ` + "`json:\"address\"`" + `
              IsOrganization  interface{} ` + "`json:\"isOrganization\"`" + `
            }
            type Address struct {
              Street1     string ` + "`json:\"street1\"`" + `
              Street2     string ` + "`json:\"street2\"`" + `
              City        string ` + "`json:\"city\"`" + `
              State       string ` + "`json:\"state\"`" + `
              PostalCode  string ` + "`json:\"postalCode\"`" + `
              Country     string ` + "`json:\"country\"`" + `
            }
            type PriceDetails struct {
              CountryCurrencyCode string ` + "`json:\"countryCurrencyCode\"`" + `
              TaxPolicy           string ` + "`json:\"taxPolicy\"`" + `
            }
            type SalesOrderItem struct {
              Id                float64 ` + "`json:\"id\"`" + `
              ItemId            float64 ` + "`json:\"itemId\"`" + `
              ItemName          string  ` + "`json:\"itemName\"`" + `
              UnitPrice         float64 ` + "`json:\"unitPrice\"`" + `
              Quantity          float64 ` + "`json:\"quantity\"`" + `
              UnitOfMeasure     string  ` + "`json:\"unitOfMeasure\"`" + `
              DiscountPct       float64 ` + "`json:\"discountPct\"`" + `
              Adjustment        float64 ` + "`json:\"adjustment\"`" + `
              TaxAmount         float64 ` + "`json:\"taxAmount\"`" + `
              LineTotal         float64 ` + "`json:\"lineTotal\"`" + `
              AccountId         float64 ` + "`json:\"accountId\"`" + `
              AccountCode       string  ` + "`json:\"accountCode\"`" + `
              TaxRateId         float64 ` + "`json:\"taxRateId\"`" + `
              TaxRateName       string  ` + "`json:\"taxRateName\"`" + `
              Sku               interface{} ` + "`json:\"sku\"`" + `
              ExternalReference interface{} ` + "`json:\"externalReference\"`" + `
            }
            type TaxBreakdown struct {
              Name      string  ` + "`json:\"name\"`" + `
              Amount    float64 ` + "`json:\"amount\"`" + `
              RatePct   float64 ` + "`json:\"ratePct\"`" + `
              Inclusive bool    ` + "`json:\"inclusive\"`" + `
            }
    SalesRegionCode       string        ` + "`json:\"salesregioncode\"`" + `
    OrderDateAsText       string        ` + "`json:\"orderdateastext\"`" + `
    SalesPriceListId      float64       ` + "`json:\"salespricelistid\"`" + `
    WineryId              float64       ` + "`json:\"wineryid\"`" + `
    WineryName            string        ` + "`json:\"wineryname\"`" + `
    Notes                 string        ` + "`json:\"notes\"`" + `
    AcctReference         *string       ` + "`json:\"acctreference\"`" + `
    PosSaleReference      []interface{} ` + "`json:\"possalereference\"`" + `
  }`,
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
