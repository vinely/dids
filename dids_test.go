package dids

import (
	"fmt"
	"testing"
)

func TestParseScheme(t *testing.T) {
	type args struct {
		uri string
	}
	tests := []struct {
		name    string
		args    args
		want    *DIDScheme
		wantErr bool
	}{
		{"1", args{"did:ont:TFt7y1hc396kVSemHfcATBEH3NdU9LYffi"}, &DIDScheme{Method: "ont", ID: "TFt7y1hc396kVSemHfcATBEH3NdU9LYffi"}, false},
		{"2", args{"url:ont:xcjvkef"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseScheme(tt.args.uri)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			fmt.Printf("Parse() = %v, want %v\n", got, tt.want)

			if got != nil {
				fmt.Println(got.String())
			}

		})
	}
}
