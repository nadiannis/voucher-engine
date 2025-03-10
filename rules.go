package main

type Rule struct {
	Conditions []Condition
	Action     Action
}

type Output struct {
	Type       ActionType `json:"type"`
	IsEligible bool       `json:"is_eligible"`
	Value      any        `json:"value"`
}

type EvaluationContext struct {
	Facts map[string]any
}
