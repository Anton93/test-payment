package payment

import (
	"context"
	"errors"
	"github.com/Anton93/test-payment/model"
)

type _Service struct {
	exts    []model.ExternalServiceIntegration
	storage model.Storage
}

func Create(storage model.Storage, exts ...model.ExternalServiceIntegration) (model.PaymentService, error) {
	if len(exts) == 0 {
		return nil, errors.New("not provided services")
	}

	s := &_Service{
		exts:    exts,
		storage: storage,
	}

	return s, nil
}

func (s *_Service) GetWebPaymentButtons(ctx context.Context, request *model.GetWebPaymentButtonsRequest) (response *model.GetWebPaymentButtonsResponse, err error) {
	if request == nil {
		err = errors.New("empty request")
		return
	}

	ids, err := s.storage.GetProductIDPaymentExt(ctx, request.ProductID)
	if err != nil  {
		return nil, err
	}

	response = &model.GetWebPaymentButtonsResponse{}
	for _, e := range s.exts {
		key := e.GetServiceKey(ctx)
		id, ok := ids[key]
		if !ok {
			id = request.ProductID
		}
		button, rErr := e.GetButtonUrl(ctx, id)
		if rErr != nil {
			err = rErr
			return
		}
		response.Button = append(response.Button, button)
	}

	return
}
