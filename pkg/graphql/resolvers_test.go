package graphql

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/olawolu/circle-payments/pkg/circle"
	"github.com/olawolu/circle-payments/pkg/payments"
)

func TestRootResolver_CreatePayment(t *testing.T) {
	type fields struct {
		ReqClient circle.ReqClient
	}
	type args struct {
		args struct{ Details payments.PaymentRequest }
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *PaymentResolver
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"Resolver Test",
			fields{
				ReqClient: circle.ReqClient{
					"QVBJX0tFWTowNGNlZGE4NTQ2MzJkNDliYjdiNDViMDU4ZjQxNTJjODplMjEwN2JiZDAzNzdmMjM1ZWY3OTBkMzM0MjE1YjFjNw==",
					"",
					"",
					"https://api-sandbox.circle.com/v1",
					*http.DefaultClient,
				},
			},
			args{
				args: struct{Details payments.PaymentRequest}{
					payments.PaymentRequest{},
				},
			},
			&PaymentResolver{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RootResolver{
				ReqClient: tt.fields.ReqClient,
			}
			got, err := r.CreatePayment(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("RootResolver.CreatePayment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RootResolver.CreatePayment() = %v, want %v", got, tt.want)
			}
		})
	}
}
