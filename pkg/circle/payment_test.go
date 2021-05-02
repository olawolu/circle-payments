package circle

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/olawolu/circle-payments/pkg/payments"
)

func TestClient_CreatePaymentCall(t *testing.T) {
	type fields struct {
		APIKey    string
		PublicKey string
		PublicID  string
		URL       string
		Client    http.Client
	}
	type args struct {
		id   string
		card payments.PaymentData
		md   payments.MetaData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *payments.PaymentResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"Create payment test",
			fields{
				"QVBJX0tFWTowNGNlZGE4NTQ2MzJkNDliYjdiNDViMDU4ZjQxNTJjODplMjEwN2JiZDAzNzdmMjM1ZWY3OTBkMzM0MjE1YjFjNw==",
				"1234566",
				"key1",
				"https://api-sandbox.circle.com/v1",
				http.Client{},
			},
			args{
				"fc988ed5-c129-4f70-a064-e5beb7eb8e32",
				payments.PaymentData{
					Amount:      "3.14",
					Description: "Payment",
				},
				payments.MetaData{
					Email:       "satoshi@circle.com",
					PhoneNumber: "+14155555555",
					SessionID:   "DE6FA86F60BB47B379307F851E238617",
					IPAddress:   "244.28.239.130",
				},
			},
			&payments.PaymentResponse{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ReqClient{
				APIKey:    tt.fields.APIKey,
				PublicKey: tt.fields.PublicKey,
				PublicID:  tt.fields.PublicID,
				URL:       tt.fields.URL,
				Client:    tt.fields.Client,
			}
			got, err := c.CreatePaymentCall(tt.args.id, tt.args.card, tt.args.md)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreatePaymentCall() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.CreatePaymentCall() = %v, want %v", got, tt.want)
			}
		})
	}
}
