package app

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Anton93/test-payment/model"
	"io"
	"net/http"
	"log"
)

const ApiPrefix = "/v1/api"

type App struct {
	model.PaymentService
}

type Error struct {
	Description string `json:"description"`
}

type CommonResponse struct {
	Success bool             `json:"success"`
	Data    json.RawMessage `json:"data"`
	Error   *Error           `json:"error,omitempty"`
}

func (a *App) GetWebPaymentButtons(writer http.ResponseWriter, request *http.Request) {
	reqData := model.GetWebPaymentButtonsRequest{}

	err := json.NewDecoder(request.Body).Decode(&reqData)
	if err != nil {
		if err == io.EOF {
			err = errors.New("empty request")
		}
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	respData, err := a.PaymentService.GetWebPaymentButtons(context.Background(), &reqData)

	a.PackResponse(writer, respData, err)

}

func (a *App) PackResponse(writer http.ResponseWriter, data interface{}, err error) {
	cr := &CommonResponse{}
	rawData, _ := json.Marshal(data)
	if err != nil {
		cr.Success = false
		cr.Error = &Error{
			Description: err.Error(),
		}
	} else {
		cr.Success = true
		cr.Data = (json.RawMessage)(rawData)
	}
	err = json.NewEncoder(writer).Encode(cr)
	if err != nil {
		log.Println(err)
	}
}
