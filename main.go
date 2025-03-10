package main

import (
	"log"
	"time"
)

func main() {
	cart1 := Cart{
		ID: 1,
		User: User{
			ID:        1,
			Username:  "nadiannis",
			Birthdate: parseDate("2010-10-05"),
			Location: Location{
				City: "Bekasi",
			},
		},
		Merchant: Merchant{
			ID:   1,
			Name: "Tomoro Coffee",
		},
		Items: []CartItem{
			{
				Product: Product{
					ID:    1,
					Name:  "Product 1",
					Price: 40000,
					Stock: 100,
				},
				Quantity: 1,
			},
			{
				Product: Product{
					ID:    2,
					Name:  "Product 2",
					Price: 34000,
					Stock: 80,
				},
				Quantity: 2,
			},
		},
		PaymentMethod: "ewallet",
		CreatedAt:     time.Now(),
	}

	voucher1 := Voucher{
		ID:   1,
		Code: "NEWYEAR123",
		Rule: Rule{
			Conditions: []Condition{
				MinPurchaseCondition{
					MinAmount: 80000,
				},
				DateValidityCondition{
					StartDate: parseDateTime("2025-01-01 00:00:00"),
					EndDate:   parseDateTime("2025-07-31 23:59:59"),
				},
				MerchantExclusionCondition{
					ExcludedMerchants: []int64{10, 2, 4},
				},
				PaymentMethodCondition{
					AllowedPaymentMethods: []string{"bank_transfer", "virtual_account", "ewallet"},
				},
			},
			Action: PercentageDiscountAction{
				Amount:    50,
				MaxAmount: 200000,
			},
		},
	}

	voucher2 := Voucher{
		ID:   1,
		Code: "BUY2GET1",
		Rule: Rule{
			Conditions: []Condition{
				DateValidityCondition{
					StartDate: parseDateTime("2025-02-01 00:00:00"),
					EndDate:   parseDateTime("2025-08-31 23:59:59"),
				},
				MerchantExclusionCondition{
					ExcludedMerchants: []int64{3, 7},
				},
				PaymentMethodCondition{
					AllowedPaymentMethods: []string{"virtual_account", "ewallet"},
				},
				ProductQuantityCondition{
					ProductID: 2,
					Operator:  GreaterThanOrEqual,
					Quantity:  2,
				},
			},
			Action: FreeItemAction{
				ProductID: 2,
				Quantity:  1,
			},
		},
	}

	engine := NewRuleEngine()
	engine.RegisterVoucher(voucher1)
	engine.RegisterVoucher(voucher2)

	output, err := engine.ApplyVoucher(cart1, "NEWYEAR123")
	if err != nil {
		log.Println(err)
		printJSON(output)
		return
	}

	printJSON(output)

	output, err = engine.ApplyVoucher(cart1, "BUY2GET1")
	if err != nil {
		log.Println(err)
		printJSON(output)
		return
	}

	printJSON(output)
}
