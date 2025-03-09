package main

type Rule struct {
	Conditions []Condition
	Action     Action
}

type Output struct {
	Type       ActionType `json:"type"`
	IsEligible bool       `json:"is_eligible"`
	Discount   float64    `json:"discount"`
}

type EvaluationContext struct {
	Facts map[string]any
}
