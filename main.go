package main

import (
	"time"
)

func main() {
	cart1 := Cart{
		ID: 1,
		User: User{
			ID:        1,
			Username:  "nadiannis",
			Birthdate: ParseDate("2010-10-05"),
			Location: Location{
				City: "Bekasi",
			},
			VoucherUsed: make(map[string]int),
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
		Discount: Discount{
			Type:      "percentage",
			Value:     50,
			MaxAmount: 200000,
		},
		MinPurchaseAmount: 80000,
		StartDate:         ParseDateTime("2025-01-01 00:00:00"),
		ExpiryDate:        ParseDateTime("2025-07-31 23:59:59"),
		UsageLimit:        3,
		ExcludedMerchants: []int64{10, 2, 4},
		PaymentMethods:    []string{"bank_transfer", "virtual_account", "ewallet"},
	}

	engine := NewRuleEngine()

	output := engine.CheckEligibility(cart1, voucher1)
	PrintJSON(output)

	output = engine.CalculateDiscount(cart1, voucher1)
	PrintJSON(output)
}
