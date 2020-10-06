package model

import "context"

type Storage interface {
	GetProductIDPaymentExt(ctx context.Context, productId string) (map[string]string, error)
}
