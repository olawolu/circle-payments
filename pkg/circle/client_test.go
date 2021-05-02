package circle

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_makeRequest(t *testing.T) {
	type fields struct {
		APIKey    string
		PublicKey string
		URL       string
		Client    http.Client
	}
	type args struct {
		endpoint string
		rtype    string
		body     io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"Test Request",
			fields{
				"QVBJX0tFWTowNGNlZGE4NTQ2MzJkNDliYjdiNDViMDU4ZjQxNTJjODplMjEwN2JiZDAzNzdmMjM1ZWY3OTBkMzM0MjE1YjFjNw==",
				"1234566",
				"https://api-sandbox.circle.com/v1",
				http.Client{},
			},
			args{
				"/encryption/public",
				"GET",
				nil,
			},
			httptest.NewRecorder().Result(),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ReqClient{
				APIKey:    tt.fields.APIKey,
				PublicKey: tt.fields.PublicKey,
				URL:       tt.fields.URL,
				Client:    tt.fields.Client,
			}
			got, err := c.makeRequest(tt.args.endpoint, tt.args.rtype, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.makeRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("Client.makeRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
