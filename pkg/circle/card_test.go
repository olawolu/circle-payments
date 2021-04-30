package circle

import (
	"testing"
)

func Test_GetPublicKey(t *testing.T) {
	tests := []struct {
		name    string
		want    *publicKeyResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"key test",
			&publicKeyResponse{
				struct {
					KeyID     string "json:\"keyId\""
					PublicKey string "json:\"publicKey\""
				}{
					"random key id",
					"random key",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := GetPublicKey()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPublicKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.Data.KeyID == "" {
				t.Errorf("GetPublicKey().Data.KeyID = %v, want %v", got.Data.KeyID, tt.want.Data.KeyID)
			}
			if got.Data.PublicKey == "" {
				t.Errorf("GetPublicKey().Data.PublicKey = %v, want %v", got.Data.PublicKey, tt.want.Data.PublicKey)
			}
		})
	}
}
