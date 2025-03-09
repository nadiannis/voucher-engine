package main

import "time"

type Location struct {
	City string
}

type User struct {
	ID        int64
	Username  string
	Birthdate time.Time
	Location  Location
}

type Merchant struct {
	ID   int64
	Name string
}

type Product struct {
	ID    int64
	Name  string
	Price float64
	Stock int
}

type CartItem struct {
	Product  Product
	Quantity int
}

type Cart struct {
	ID            int64
	User          User
	Merchant      Merchant
	Items         []CartItem
	PaymentMethod string
	CreatedAt     time.Time
}

type Discount struct {
	Type      string
	Value     float64
	MaxAmount float64
}

type Voucher struct {
	ID   int64
	Code string
	Rule Rule
}
