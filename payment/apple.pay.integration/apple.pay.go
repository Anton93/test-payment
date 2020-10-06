package integration

import (
	"context"
	"github.com/Anton93/test-payment/model"
)

const PaymentServiceID = 1
const PaymentServiceName = "apple_pay"

type ApplePayApi struct {
	
}

func (a *ApplePayApi) GetButtonUrl(ctx context.Context, productIdExt string) (*model.WebPaymentButtonExtService, error) {
	panic("implement me")
}

