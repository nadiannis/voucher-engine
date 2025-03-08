package main

import (
	"fmt"
	"strings"
)

type RuleEngine struct {
	eligibilityRules []Rule
}

func NewRuleEngine() *RuleEngine {
	return &RuleEngine{
		eligibilityRules: GenerateEligibilityRules(),
	}
}

func (e *RuleEngine) CheckEligibility(cart Cart, voucher Voucher) Output {
	facts := map[string]any{
		"cart":          cart,
		"voucher":       voucher,
		"totalPurchase": calculateTotalPurchase(cart),
	}

	for _, rule := range e.eligibilityRules {
		if !evaluateRule(rule, facts) {
			return Output{
				Type:  "eligibility",
				Value: false,
			}
		}
	}

	return Output{
		Type:  "eligibility",
		Value: true,
	}
}

func (e *RuleEngine) CalculateDiscount(cart Cart, voucher Voucher) Output {
	output := Output{
		Type:  "discount",
		Value: 0.0,
	}

	if eligibility := e.CheckEligibility(cart, voucher); !eligibility.Value.(bool) {
		fmt.Println("this purchase is not eligible to use the voucher")
		return output
	}

	totalPurchase := calculateTotalPurchase(cart)
	var discount float64

	switch voucher.Discount.Type {
	case "fixed":
		discount = voucher.Discount.Value

		if discount > voucher.Discount.MaxAmount {
			discount = voucher.Discount.MaxAmount
		}

		if discount > totalPurchase {
			discount = totalPurchase
		}
	case "percentage":
		discount = totalPurchase * voucher.Discount.Value / 100

		if discount > voucher.Discount.MaxAmount {
			discount = voucher.Discount.MaxAmount
		}
	default:
		fmt.Println("unsupported discount type")
	}

	output.Value = discount
	return output
}

func evaluateRule(rule Rule, facts map[string]any) bool {
	for _, condition := range rule.Conditions {
		if !match(condition, facts) {
			return false
		}
	}

	return true
}

func match(condition Condition, facts map[string]any) bool {
	fieldPath := strings.Split(condition.Field, ".")
	operand1 := getNestedValue(facts, fieldPath)

	var operand2 any
	if v, ok := condition.Value.(string); ok && strings.Contains(v, ".") {
		refPath := strings.Split(v, ".")
		operand2 = getNestedValue(facts, refPath)
	} else {
		operand2 = condition.Value
	}

	switch condition.Operator {
	case Equal:
		return compareEqual(operand1, operand2)
	case NotEqual:
		return !compareEqual(operand1, operand2)
	case GreaterThan:
		return compareGreaterThan(operand1, operand2)
	case GreaterThanOrEqual:
		return compareGreaterThanOrEqual(operand1, operand2)
	case LessThan:
		return compareLessThan(operand1, operand2)
	case LessThanOrEqual:
		return compareLessThanOrEqual(operand1, operand2)
	case In:
		return compareIn(operand1, operand2)
	case NotIn:
		return !compareIn(operand1, operand2)
	default:
		fmt.Println("unsupported condition operator")
		return false
	}
}
