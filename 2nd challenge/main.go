package main

import (
	"CodeChallenge/convertor"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Started main!")
	inpJson := `{
		"number_1": {
		  "N": "1.50"
		},
		"string_1": {
		  "S": "784498 "
		},
		"string_2": {
		  "S": "2014-07-16T20:55:46Z"
		},
		"map_1": {
		  "M": {
			"bool_1": {
			  "BOOL": "truthy"
			},
			"null_1": {
			  "NULL ": "true"
			},
			"list_1": {
			  "L": [
				{
				  "S": ""
				},
				{
				  "N": "011"
				},
				{
				  "N": "5215s"
				},
				{
				  "BOOL": "f"
				},
				{
				  "NULL": "0"
				}
			  ]
			}
		  }
		},
		"list_2": {
		  "L": "noop"
		},
		"list_3": {
		  "L": [
			"noop"
		  ]
		},
		"": {
		  "S": "noop"
		}
	  }`
	inpStruct := convertor.ConvertJsonToStruct(inpJson)
	outputStruct, err := convertor.Transformer(inpStruct)
	if err != nil {
		log.Fatal(err.Error())
	}
	outputJson, err := convertor.ConvertStructToJson(outputStruct)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("outputJson:", outputJson)
}
