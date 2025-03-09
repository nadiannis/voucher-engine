package main

type ActionType string

const (
	PercentageDiscountType ActionType = "percentage_discount"
	FixedDiscountType      ActionType = "fixed_discount"
)

type Action interface {
	GetType() ActionType
	Apply(ctx *EvaluationContext) float64
}

type PercentageDiscountAction struct {
	Amount    float64
	MaxAmount float64
}

func (a PercentageDiscountAction) GetType() ActionType {
	return PercentageDiscountType
}

func (a PercentageDiscountAction) Apply(ctx *EvaluationContext) float64 {
	totalPurchase := calculateTotalPurchase(ctx.Facts["cart"].(Cart))
	discount := totalPurchase * a.Amount / 100

	if a.MaxAmount > 0 && discount > a.MaxAmount {
		discount = a.MaxAmount
	}

	return discount
}

type FixedDiscountAction struct {
	Amount    float64
	MaxAmount float64
}

func (a FixedDiscountAction) GetType() ActionType {
	return FixedDiscountType
}

func (a FixedDiscountAction) Apply(ctx *EvaluationContext) float64 {
	discount := a.Amount

	if a.MaxAmount > 0 && discount > a.MaxAmount {
		discount = a.MaxAmount
	}

	totalPurchase := calculateTotalPurchase(ctx.Facts["cart"].(Cart))
	if discount > totalPurchase {
		discount = totalPurchase
	}

	return discount
}
