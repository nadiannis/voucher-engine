package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

func CalculateTotalPurchase(cart Cart) float64 {
	var total float64

	for _, item := range cart.Items {
		total += float64(item.Quantity) * item.Product.Price
	}

	return total
}

func GetFieldValue(obj any, fieldName string) reflect.Value {
	value := reflect.ValueOf(obj)
	actualObj := reflect.Indirect(value)

	if actualObj.Kind() != reflect.Struct {
		panic("object is not a struct")
	}

	fieldValue := actualObj.FieldByName(fieldName)
	if !fieldValue.IsValid() {
		panic(fmt.Sprintf("field '%s' is not found", fieldName))
	}

	return fieldValue
}

func PrintJSON(data any) error {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	fmt.Println(string(jsonBytes))
	return nil
}

func ParseDate(date string) time.Time {
	format := "2006-01-02"
	t, _ := time.Parse(format, date)
	return t
}

func ParseDateTime(datetime string) time.Time {
	format := "2006-01-02 15:04:05"
	t, _ := time.Parse(format, datetime)
	return t
}
