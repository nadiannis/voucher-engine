package main

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

type Condition struct {
	Field    string
	Operator Operator
	Value    any
}

type Rule struct {
	ID         int64
	Name       string
	Conditions []Condition
}

type Output struct {
	Type  string `json:"type"`
	Value any    `json:"value"`
}

func GenerateEligibilityRules() []Rule {
	return []Rule{
		{
			ID:   1,
			Name: "Minimum Purchase Amount",
			Conditions: []Condition{
				{
					Field:    "totalPurchase",
					Operator: GreaterThanOrEqual,
					Value:    "voucher.MinPurchaseAmount",
				},
			},
		},
		{
			ID:   2,
			Name: "Date Validity",
			Conditions: []Condition{
				{
					Field:    "cart.CreatedAt",
					Operator: GreaterThanOrEqual,
					Value:    "voucher.StartDate",
				},
				{
					Field:    "cart.CreatedAt",
					Operator: LessThanOrEqual,
					Value:    "voucher.ExpiryDate",
				},
			},
		},
		{
			ID:   3,
			Name: "Merchant Exclusion",
			Conditions: []Condition{
				{
					Field:    "cart.Merchant.ID",
					Operator: NotIn,
					Value:    "voucher.ExcludedMerchants",
				},
			},
		},
		{
			ID:   4,
			Name: "Payment Method",
			Conditions: []Condition{
				{
					Field:    "cart.PaymentMethod",
					Operator: In,
					Value:    "voucher.PaymentMethods",
				},
			},
		},
	}
}
