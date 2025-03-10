package main

type ActionType string

const (
	PercentageDiscountType ActionType = "percentage_discount"
	FixedDiscountType      ActionType = "fixed_discount"
	FreeItemType           ActionType = "free_item"
)

type Action interface {
	GetType() ActionType
	Apply(ctx *EvaluationContext) any
}

type PercentageDiscountAction struct {
	Amount    float64
	MaxAmount float64
}

func (a PercentageDiscountAction) GetType() ActionType {
	return PercentageDiscountType
}

func (a PercentageDiscountAction) Apply(ctx *EvaluationContext) any {
	totalPurchase := calculateTotalPurchase(ctx.Facts["cart"].(Cart))
	discount := totalPurchase * a.Amount / 100

	if a.MaxAmount > 0 && discount > a.MaxAmount {
		discount = a.MaxAmount
	}

	return discount
}

type FixedDiscountAction struct {
	Amount float64
}

func (a FixedDiscountAction) GetType() ActionType {
	return FixedDiscountType
}

func (a FixedDiscountAction) Apply(ctx *EvaluationContext) any {
	discount := a.Amount

	totalPurchase := calculateTotalPurchase(ctx.Facts["cart"].(Cart))
	if discount > totalPurchase {
		discount = totalPurchase
	}

	return discount
}

type FreeItemAction struct {
	ProductID int64
	Quantity  int
}

func (a FreeItemAction) GetType() ActionType {
	return FreeItemType
}

func (a FreeItemAction) Apply(ctx *EvaluationContext) any {
	return map[string]any{
		"product_id": a.ProductID,
		"quantity":   a.Quantity,
	}
}
