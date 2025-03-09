package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

func calculateTotalPurchase(cart Cart) float64 {
	var total float64

	for _, item := range cart.Items {
		total += float64(item.Quantity) * item.Product.Price
	}

	return total
}

func printJSON(data any) error {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	fmt.Println(string(jsonBytes))
	return nil
}

func parseDate(date string) time.Time {
	format := "2006-01-02"
	t, _ := time.Parse(format, date)
	return t
}

func parseDateTime(datetime string) time.Time {
	format := "2006-01-02 15:04:05"
	t, _ := time.Parse(format, datetime)
	return t
}

func compareValues(operator Operator, a, b any) bool {
	switch operator {
	case Equal:
		return compareEqual(a, b)
	case NotEqual:
		return !compareEqual(a, b)
	case GreaterThan:
		return compareGreaterThan(a, b)
	case GreaterThanOrEqual:
		return compareGreaterThanOrEqual(a, b)
	case LessThan:
		return compareLessThan(a, b)
	case LessThanOrEqual:
		return compareLessThanOrEqual(a, b)
	case In:
		return compareIn(a, b)
	case NotIn:
		return !compareIn(a, b)
	default:
		fmt.Println("unsupported condition operator")
		return false
	}
}

func compareEqual(a, b any) bool {
	return fmt.Sprint(a) == fmt.Sprint(b)
}

func compareGreaterThan(a, b any) bool {
	aFloat, aOk := toFloat64(a)
	bFloat, bOk := toFloat64(b)

	if aOk && bOk {
		return aFloat > bFloat
	}

	aTime, aOk := a.(time.Time)
	bTime, bOk := b.(time.Time)

	if aOk && bOk {
		return aTime.After(bTime)
	}

	fmt.Println("cannot compare the values")
	return false
}

func compareGreaterThanOrEqual(a, b any) bool {
	if compareEqual(a, b) {
		return true
	}

	return compareGreaterThan(a, b)
}

func compareLessThan(a, b any) bool {
	aFloat, aOk := toFloat64(a)
	bFloat, bOk := toFloat64(b)

	if aOk && bOk {
		return aFloat < bFloat
	}

	aTime, aOk := a.(time.Time)
	bTime, bOk := b.(time.Time)

	if aOk && bOk {
		return aTime.Before(bTime)
	}

	fmt.Println("cannot compare the values")
	return false
}

func compareLessThanOrEqual(a, b any) bool {
	if compareEqual(a, b) {
		return true
	}

	return compareLessThan(a, b)
}

func compareIn(value, collection any) bool {
	c := reflect.ValueOf(collection)
	if c.Kind() == reflect.Slice || c.Kind() == reflect.Array {
		for i := 0; i < c.Len(); i++ {
			if compareEqual(value, c.Index(i).Interface()) {
				return true
			}
		}
	}

	return false
}

func toFloat64(v any) (float64, bool) {
	switch value := v.(type) {
	case int:
		return float64(value), true
	case int32:
		return float64(value), true
	case int64:
		return float64(value), true
	case float32:
		return float64(value), true
	case float64:
		return value, true
	default:
		return 0, false
	}
}
