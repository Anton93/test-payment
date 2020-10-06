package storage

import (
	"context"
	"errors"
	"github.com/Anton93/test-payment/model"

	api "github.com/Anton93/test-payment/payment/apple.pay.integration"
	gpi "github.com/Anton93/test-payment/payment/google.pay.integration"
	ppi "github.com/Anton93/test-payment/payment/paypal.integration"
	spi "github.com/Anton93/test-payment/payment/stripe.integration"
)

type _StorageMock struct {
	ids map[string]map[string]string
}

func NewMockStorage() model.Storage {
	s := &_StorageMock{}

	s.ids = map[string]map[string]string{
		"1": {
			gpi.PaymentServiceName: "google_pay_1",
			ppi.PaymentServiceName: "paypal_1",
			api.PaymentServiceName: "apple_pay_1",
			spi.PaymentServiceName: "stripe_1",
		},
		"2": {
			gpi.PaymentServiceName: "google_pay_2",
			ppi.PaymentServiceName: "paypal_2",
			api.PaymentServiceName: "apple_pay_2",
		},
		"3": {},
	}

	return s
}

func (s *_StorageMock) GetProductIDPaymentExt(_ context.Context, productId string) (map[string]string, error) {
	ids, ok := s.ids[productId]
	if !ok {
		return nil, errors.New("product not found")
	}
	return ids, nil
}
