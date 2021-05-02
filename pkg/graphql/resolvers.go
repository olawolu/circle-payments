package graphql

import (
	"log"

	"github.com/olawolu/circle-payments/pkg/circle"
	"github.com/olawolu/circle-payments/pkg/db"
	"github.com/olawolu/circle-payments/pkg/payments"
)

type cardPayment struct {
	Amount string
	payments.CardData
}

// type cardDetails struct {
// 	CardNumber  string
// 	CVV         string
// 	ExpiryMonth string
// 	ExpiryYear  string
// }

type RootResolver struct {
	circle.ReqClient
	database db.Mongo
}

type PaymentResolver struct {
	payment *payments.PaymentResponse
}

func (r *RootResolver) GetPayment(args struct{ PayID string }) (*PaymentResolver, error) {
	return nil, nil
}

func (r *RootResolver) CreatePayment(args struct{ Details payments.PaymentRequest }) (*PaymentResolver, error) {
	// create card -> create payment
	//
	// create card

	var paymentDetail payments.PaymentData
	log.Println("args:", args)
	paymentDetail.Amount = args.Details.Amount
	paymentDetail.Description = args.Details.Description

	cardId, err := r.CreateCardCall(args.Details.CardData)
	if err != nil {
		log.Println("graphql.RootResolver.CreateCardCall() ", err)
		return nil, err
	}

	// create payment
	response, err := r.CreatePaymentCall(cardId, paymentDetail, args.Details.MetaData)
	if err != nil {
		log.Println("graphql.RootResolver.CreatePaymentCall() ", err)
		return nil, err
	}

	err = r.database.Add(response)
	if err != nil {
		log.Println("graphql.RootResolver.CreatePaymentCall() database add operation failed ", err)
	}

	return &PaymentResolver{response}, nil
}

func (pr *PaymentResolver) PaymentID() string {
	return pr.payment.PaymentID
}

func (pr *PaymentResolver) CardID() string {
	return pr.payment.Source.ID
}

func (pr *PaymentResolver) Email() string {
	return pr.payment.Metadata.Email
}

func (pr *PaymentResolver) Amount() string {
	return pr.payment.Amount.Amount
}

func (pr *PaymentResolver) Currency() string {
	return pr.payment.Amount.Currency
}

func (pr *PaymentResolver) Status() string {
	return pr.payment.Status
}

func (pr *PaymentResolver) CreateDate() string {
	return pr.payment.CreateDate
}

func (pr *PaymentResolver) UpdateDate() string {
	return pr.payment.UpdateDate
}
