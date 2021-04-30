package payments

import (
	"reflect"
	"testing"
)

func Test_Encrypt(t *testing.T) {
	type args struct {
		data CardDetails
		key  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"encryption test",
			args{
				CardDetails{
					"4111111111111111",
					"123",
				},
				"109208-381920023-2",
			},
			"encrypted string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Encrypt(tt.args.data, tt.args.key)

			t.Logf("Encrypt() = %v", f)
			expectedType := reflect.String
			returnType := reflect.TypeOf(f)

			if got := returnType.Kind(); got != expectedType {
				t.Errorf("Encrypt() = %v, want %v", got, tt.want)
			}

			// if got := encrypt(tt.args.data, tt.args.key); got != tt.want {
			// 	t.Errorf("encrypt() = %v, want %v", got, tt.want)
			// }
		})
	}
}
