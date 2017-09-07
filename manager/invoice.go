package manager

import (
	"fmt"
	"time"
)

type Invoice struct {
	ID                 uint64
	Client             string
	Emitted, Delivered time.Time
	Services           []Service
	Comment            string
	Currency           rune
	PaymentDays        uint8
	Quote              bool
}

func (i Invoice) Total() float64 {
	var t float64

	for _, s := range i.Services {
		t += s.Amount()
	}

	return t
}

type Service struct {
	Description string
	UnitCost    float64
	Unit        string
	Quantity    float64
}

func (s Service) format(currency rune) (desc, qty, uc, amount string) {
	desc = s.Description
	qty = fmt.Sprintf("%g %s", s.Quantity, s.Unit)
	uc = formatMoney(s.UnitCost, currency)
	amount = formatMoney(s.UnitCost*s.Quantity, currency)
	return
}

func (s Service) Amount() float64 {
	return s.UnitCost * s.Quantity
}

func formatMoney(value float64, currency rune) string {
	return fmt.Sprintf("%.2f %c", value, currency)
}
