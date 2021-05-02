package circle

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/google/uuid"
	"github.com/olawolu/circle-payments/pkg/payments"
)

// make a request to circle's public key api and cache the response
var (
	paymentURL   = "https://api-sandbox.circle.com/v1/payments"
	cardURL      = "https://api-sandbox.circle.com/v1/cards"
	publicKeyURL = "https://api-sandbox.circle.com/v1/encryption/public"
	API_KEY      = "QVBJX0tFWTplOTMyMDRmZDBmNDhjZWNlNzI0NWQ4MTgwYmFkZTQ1YTowYjZhMzcwZGY1ODdlMjhhMWVkNGY0MzBjNjdmNDIyMw=="
)


type publicKeyResponse struct {
	Data struct {
		KeyID     string `json:"keyId"`
		PublicKey string `json:"publicKey"`
	} `json:"data"`
}

type createCardResponse struct {
	Data []struct {
		ID string `json:"id"`
	} `json:"data"`
	ErrorCode int    `json:"code"`
	Message   string `json:"message"`
}

type cardPayload struct {
	IdempotencyKey          string `json:"idempotencyKey"`
	KeyID                   string `json:"keyId"`
	ExpMonth                int32  `json:"expMonth"`
	ExpYear                 int32  `json:"expYear"`
	EncryptedData           string `json:"encryptedData"`
	payments.BillingDetails `json:"billingDetails"`
	payments.MetaData       `json:"metaData"`
}

func (c *ReqClient) CreateCardCall(card payments.CardData) (string, error) {
	var payload cardPayload
	var response createCardResponse

	// payments.CopyIdenticalFields(card, &payload)

	err := c.GetPublicKey()
	if err != nil {
		log.Printf("%v: could not get public key", err)
	}
	payload.IdempotencyKey = uuid.NewString()
	payload.KeyID = c.PublicID
	payload.EncryptedData = payments.Encrypt(card.CardDetails, c.PublicKey)
	payload.BillingDetails = card.BillingDetails
	payload.ExpMonth = card.ExpiryMonth
	payload.ExpYear = card.ExpiryYear
	payload.MetaData = card.MetaData

	// log.Println("CreateCard payload: ", payload)
	body, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
	}

	log.Println("requestBody: ", string(body))

	resp, err := c.makeRequest("/cards", "POST", bytes.NewBuffer(body))
	if err != nil {
		log.Println("ReqClientCall.CreateCardCall().makeRequest()", err)
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	// log.Println("ReqClient.CreateCardCall() response: ", string(responseBody))
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		log.Fatal("decode:", err)
	}

	if len(response.Data) != 0 {
		return response.Data[0].ID, nil

	}
	return response.Message, nil
}
