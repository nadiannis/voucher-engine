package main

import "time"

type Operator string

const (
	Equal              Operator = "eq"
	NotEqual           Operator = "neq"
	GreaterThan        Operator = "gt"
	GreaterThanOrEqual Operator = "gte"
	LessThan           Operator = "lt"
	LessThanOrEqual    Operator = "lte"
	In                 Operator = "in"
	NotIn              Operator = "nin"
)

type ConditionType string

const (
	MinPurchaseType       ConditionType = "min_purchase"
	DateValidityType      ConditionType = "date_validity"
	MerchantExclusionType ConditionType = "merchant_exclusion"
	PaymentMethodType     ConditionType = "payment_method"
	ProductQuantityType   ConditionType = "product_quantity"
)

type Condition interface {
	GetType() ConditionType
	Evaluate(ctx *EvaluationContext) bool
}

type MinPurchaseCondition struct {
	MinAmount float64
}

func (c MinPurchaseCondition) GetType() ConditionType {
	return MinPurchaseType
}

func (c MinPurchaseCondition) Evaluate(ctx *EvaluationContext) bool {
	return calculateTotalPurchase(ctx.Facts["cart"].(Cart)) >= c.MinAmount
}

type DateValidityCondition struct {
	StartDate time.Time
	EndDate   time.Time
}

func (c DateValidityCondition) GetType() ConditionType {
	return DateValidityType
}

func (c DateValidityCondition) Evaluate(ctx *EvaluationContext) bool {
	cart := ctx.Facts["cart"].(Cart)

	if cart.CreatedAt.Equal(c.StartDate) || cart.CreatedAt.Equal(c.EndDate) {
		return true
	}

	return cart.CreatedAt.After(c.StartDate) && cart.CreatedAt.Before(c.EndDate)
}

type MerchantExclusionCondition struct {
	ExcludedMerchants []int64
}

func (c MerchantExclusionCondition) GetType() ConditionType {
	return MerchantExclusionType
}

func (c MerchantExclusionCondition) Evaluate(ctx *EvaluationContext) bool {
	for _, merchantID := range c.ExcludedMerchants {
		if ctx.Facts["cart"].(Cart).Merchant.ID == merchantID {
			return false
		}
	}

	return true
}

type PaymentMethodCondition struct {
	AllowedPaymentMethods []string
}

func (c PaymentMethodCondition) GetType() ConditionType {
	return PaymentMethodType
}

func (c PaymentMethodCondition) Evaluate(ctx *EvaluationContext) bool {
	for _, paymentMethod := range c.AllowedPaymentMethods {
		if ctx.Facts["cart"].(Cart).PaymentMethod == paymentMethod {
			return true
		}
	}

	return false
}

type ProductQuantityCondition struct {
	ProductID int64
	Operator  Operator
	Quantity  int
}

func (c ProductQuantityCondition) GetType() ConditionType {
	return ProductQuantityType
}

func (c ProductQuantityCondition) Evaluate(ctx *EvaluationContext) bool {
	cart := ctx.Facts["cart"].(Cart)

	for _, item := range cart.Items {
		if item.Product.ID == c.ProductID {
			return compareValues(c.Operator, item.Quantity, c.Quantity)
		}
	}

	return false
}
