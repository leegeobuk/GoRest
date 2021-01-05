package server

import (
	"context"
	"fmt"

	"github.com/leegeobuk/GoServer/GoServer/pb/currency"
)

// Currency struct
type Currency struct {
}

// NewCurrency is a constructor of Currency
func NewCurrency() *Currency {
	return &Currency{}
}

// GetRate return rate
func (c *Currency) GetRate(ctx context.Context, req *currency.RateRequest) (*currency.RateResponse, error) {
	fmt.Println("Handle GetRate", "base:", req.GetBase(), "dst:", req.GetDestination())
	return &currency.RateResponse{Rate: 0.5}, nil
}
