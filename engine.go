package main

import (
	"errors"
	"fmt"
)

type RuleEngine struct {
	vouchers map[string]Voucher
}

func NewRuleEngine() *RuleEngine {
	return &RuleEngine{
		vouchers: make(map[string]Voucher),
	}
}

func (e *RuleEngine) RegisterVoucher(voucher Voucher) {
	e.vouchers[voucher.Code] = voucher
}

func (e *RuleEngine) ApplyVoucher(cart Cart, voucherCode string) (output *Output, err error) {
	voucher, exists := e.vouchers[voucherCode]
	if !exists {
		return nil, errors.New("voucher not found")
	}

	ctx := &EvaluationContext{
		Facts: map[string]any{
			"cart":          cart,
			"voucher":       voucher,
			"totalPurchase": calculateTotalPurchase(cart),
		},
	}

	output = &Output{
		Type:       voucher.Rule.Action.GetType(),
		IsEligible: false,
		Discount:   0.0,
	}

	for _, condition := range voucher.Rule.Conditions {
		if !condition.Evaluate(ctx) {
			return output, fmt.Errorf("cart does not satisfy '%v' rule condition", condition.GetType())
		}
	}

	discount := voucher.Rule.Action.Apply(ctx)

	output.IsEligible = true
	output.Discount = discount

	return output, nil
}
