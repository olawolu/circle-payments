package circle

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/olawolu/circle-payments/pkg/payments"
)

func TestClient_GetPublicKey(t *testing.T) {
	type fields struct {
		APIKey    string
		PublicKey string
		PublicID  string
		URL       string
		Client    http.Client
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"key test",
			fields{
				"QVBJX0tFWTowNGNlZGE4NTQ2MzJkNDliYjdiNDViMDU4ZjQxNTJjODplMjEwN2JiZDAzNzdmMjM1ZWY3OTBkMzM0MjE1YjFjNw==",
				"",
				"",
				"https://api-sandbox.circle.com/v1",
				http.Client{},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ReqClient{
				APIKey:    tt.fields.APIKey,
				PublicKey: tt.fields.PublicKey,
				PublicID:  tt.fields.PublicID,
				URL:       tt.fields.URL,
				Client:    tt.fields.Client,
			}
			if err := c.GetPublicKey(); (err != nil) != tt.wantErr {
				t.Errorf("Client.GetPublicKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_CreateCardCall(t *testing.T) {
	type fields struct {
		APIKey    string
		PublicKey string
		PublicID  string
		URL       string
		Client    http.Client
	}
	type args struct {
		card payments.CardData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"key test",
			fields{
				"QVBJX0tFWTowNGNlZGE4NTQ2MzJkNDliYjdiNDViMDU4ZjQxNTJjODplMjEwN2JiZDAzNzdmMjM1ZWY3OTBkMzM0MjE1YjFjNw==",
				"",
				"",
				"https://api-sandbox.circle.com/v1",
				http.Client{},
			},
			args{
				payments.CardData{
					ExpiryMonth: 9,
					ExpiryYear:  2021,
					CardDetails: payments.CardDetails{
						CardNumber: "12345678990",
						CVV:        "123",
					},
					BillingDetails: payments.BillingDetails{
						Name:         "Satoshi Nakamoto",
						City:         "City",
						Country:      "Country",
						AddressLine1: "line1",
						AddressLine2: "line2",
						District:     "District",
						PostalCode:   "1234",
					},
					MetaData: payments.MetaData{
						Email:       "satoshi@circle.com",
						PhoneNumber: "+14155555555",
						SessionID:   "DE6FA86F60BB47B379307F851E238617",
						IPAddress:   "244.28.239.130",
					},
				},
			},
			"12345675443",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ReqClient{
				APIKey:    tt.fields.APIKey,
				PublicKey: tt.fields.PublicKey,
				PublicID:  tt.fields.PublicID,
				URL:       tt.fields.URL,
				Client:    tt.fields.Client,
			}
			got, err := c.CreateCardCall(tt.args.card)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateCardCall() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("Client.CreateCardCall() = %v, want %v", got, tt.want)
			}
		})
	}
}
