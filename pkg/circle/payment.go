package circle

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/olawolu/circle-payments/pkg/payments"
)

type createPaymentResponse struct {
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}

type paymentPayload struct {
	IdempotencyKey string `json:"idempotencyKey"`
	Verification   string `json:"verification"`
	Description    string `json:"description"`
	Source         source `json:"source"`
	Amount         bill   `json:"amount"`
	payments.MetaData
}

type source struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type bill struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

func CreatePaymentCall(id string, card payments.PaymentData) (string, error) {
	var payload paymentPayload
	var response createPaymentResponse

	client := &http.Client{}
	bearer := fmt.Sprintf("Bearer %v", API_KEY)

	copyIdenticalFields(card, &payload)
	payload.IdempotencyKey = uuid.NewString()
	payload.Source.ID = id
	payload.Source.Type = "card"
	payload.Amount.Amount = card.Amount
	payload.Amount.Currency = "USD"

	body, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
	}

	req, err := http.NewRequest("POST", paymentURL, bytes.NewBuffer(body))
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Authorization", bearer)
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		log.Fatal("decode:", err)
	}
	return response.Data.ID, nil
}

