package circle

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/google/uuid"
	"github.com/olawolu/circle-payments/pkg/payments"
)

type createPaymentResponse struct {
	Data      payments.PaymentResponse `json:"data"`
	ErrorCode int                      `json:"code"`
	Message   string                   `json:"message"`
}

type paymentPayload struct {
	IdempotencyKey string          `json:"idempotencyKey"`
	Verification   string          `json:"verification"`
	Description    string          `json:"description"`
	Source         payments.Source `json:"source"`
	Amount         payments.Bill   `json:"amount"`
	payments.MetaData
}

func (c *ReqClient) CreatePaymentCall(id string, card payments.PaymentData, md payments.MetaData) (*payments.PaymentResponse, error) {
	var payload paymentPayload
	var response createPaymentResponse

	// payments.CopyIdenticalFields(card, &payload)
	payload.IdempotencyKey = uuid.NewString()
	payload.Source.ID = id
	payload.Description = card.Description
	payload.Verification = "none"
	payload.Source.Type = "card"
	payload.Amount.Amount = card.Amount
	payload.Amount.Currency = "USD"
	payload.MetaData = md

	body, err := json.Marshal(payload)
	if err != nil {
		log.Println("client.CreatePaymentCall() json: ", err)
	}

	log.Println("requestBody: ", string(body))

	resp, err := c.makeRequest("/payments", "POST", bytes.NewBuffer(body))
	if err != nil {
		log.Println("client.CreatePaymentCall() resp: ", err)
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("client.CreatePaymentCall() read: ", err)
	}
	// log.Println("ReqClient.CreatePaymentCall() response: ",string(responseBody))

	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		log.Fatal("client.CreatePaymentCall() decode:", err)
	}
	fmt.Println(&response)

	return &response.Data, nil
}

func (c *ReqClient) GetPaymentCall(paymentId string) (*payments.PaymentResponse, error) {
	var response createPaymentResponse

	resp, err := c.makeRequest("/payments"+"/"+paymentId, "GET", nil)
	if err != nil {
		log.Println("client.CreatePaymentCall() resp: ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("client.GetPaymentCall() read: ", err)
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal("client.GetPaymentCall() decode: ", err)
	}

	return &response.Data, nil
}
