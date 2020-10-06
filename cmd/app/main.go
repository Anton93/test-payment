package main

import (
	"fmt"
	"github.com/Anton93/test-payment/app"
	"github.com/Anton93/test-payment/middleware"
	"github.com/Anton93/test-payment/payment"
	"github.com/Anton93/test-payment/storage"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/http"

	api "github.com/Anton93/test-payment/payment/apple.pay.integration"
	gpi "github.com/Anton93/test-payment/payment/google.pay.integration"
	ppi "github.com/Anton93/test-payment/payment/paypal.integration"
	spi "github.com/Anton93/test-payment/payment/stripe.integration"
)

type Config struct {
	ServerPort int `envconfig:"API_PORT" default:"8080"`
}

func main() {

	config := Config{}
	envconfig.MustProcess("APP", &config)

	mux := http.NewServeMux()
	mmux := http.NewServeMux()



	iapi, _ := api.NewPayService()
	igpi, _ := gpi.NewPayService()
	ippi, _ := ppi.NewPayService()
	ispi, _ := spi.NewPayService()

	paymentService,  _ := payment.Create(storage.NewMockStorage(), iapi, igpi, ippi, ispi)

	application := &app.App{
		PaymentService: paymentService,
	}

	// TODO: gorilla router
	//r := mux.NewRouter()
	//r.HandleFunc("/payment-buttons", application.GetWebPaymentButtons).
	//	Methods("POST").
	//	Schemes("https")


	mux.HandleFunc(app.ApiPrefix + "/payment-buttons", application.GetWebPaymentButtons)

	mmux.Handle(app.ApiPrefix +"/", middleware.Handle(mux))

	log.Printf("Listening on :%d...\n", config.ServerPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.ServerPort), mmux)
	if err != nil {
		log.Panic(err)
	}
}
