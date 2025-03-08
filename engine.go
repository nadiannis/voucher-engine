package main

import (
	"fmt"
	"slices"
)

type RuleEngine struct {
}

func NewRuleEngine() *RuleEngine {
	return &RuleEngine{}
}

func (e *RuleEngine) CheckEligibility(cart Cart, voucher Voucher) Output {
	output := Output{
		Type:  "eligibility",
		Value: false,
	}

	totalPurchase := CalculateTotalPurchase(cart)
	if totalPurchase < voucher.MinPurchaseAmount {
		return output
	}

	if cart.CreatedAt.Before(voucher.StartDate) || cart.CreatedAt.After(voucher.ExpiryDate) {
		return output
	}

	if cart.User.VoucherUsed[voucher.Code] >= voucher.UsageLimit {
		return output
	}

	if slices.Contains(voucher.ExcludedMerchants, cart.Merchant.ID) {
		return output
	}

	if !slices.Contains(voucher.PaymentMethods, cart.PaymentMethod) {
		return output
	}

	output.Value = true
	return output
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

	totalPurchase := CalculateTotalPurchase(cart)
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
