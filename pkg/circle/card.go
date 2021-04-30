package circle

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"

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

// type Data struct {
// 	KeyID     string `json:"keyId"`
// 	PublicKey string `json:"publicKey"`
// }

type publicKeyResponse struct {
	Data struct {
		KeyID     string `json:"keyId"`
		PublicKey string `json:"publicKey"`
	} `json:"data"`
}

type createCardResponse struct {
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}

type cardPayload struct {
	IdempotencyKey string `json:"idempotencyKey"`
	KeyID          string `json:"keyId"`
	ExpMonth       int32  `json:"expMonth"`
	ExpYear        int32  `json:"expYear"`
	EncryptedData  string `json:"encryptedData"`
	payments.BillingDetails
	payments.MetaData
}

func GetPublicKey() (*publicKeyResponse, error) {
	var data publicKeyResponse
	client := &http.Client{}
	bearer := fmt.Sprintf("Bearer %v", API_KEY)

	req, err := http.NewRequest("GET", publicKeyURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", bearer)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal("decode:", err)
	}

	fmt.Println(&data)
	return &data, nil
}

func CreateCardCall(card payments.CardData) (string, error) {
	var payload cardPayload
	var data createCardResponse

	client := &http.Client{}
	bearer := fmt.Sprintf("Bearer %v", API_KEY)

	copyIdenticalFields(card, &payload)

	sensitiveData := card.CardDetails

	key, err := GetPublicKey()
	if err != nil {
		log.Println(err)
	}
	payload.IdempotencyKey = uuid.NewString()
	payload.KeyID = key.Data.KeyID
	payload.EncryptedData = payments.Encrypt(sensitiveData, key.Data.PublicKey)

	body, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
	}

	req, err := http.NewRequest("POST", cardURL, bytes.NewBuffer(body))
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

	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		log.Fatal("decode:", err)
	}
	return data.Data.ID, nil
}

func copyIdenticalFields(a, b interface{}) {
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b).Elem()

	at := av.Type()
	for i := 0; i < at.NumField(); i++ {
		name := at.Field(i).Name

		bf := bv.FieldByName(name)
		if bf.IsValid() {
			bf.Set(av.Field(i))
		}
	}
}
