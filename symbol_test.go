package ascendex

import (
	"testing"
)

func TestSymbolChecker(t *testing.T) {
	type args struct {
		symbol string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "CASE-1",
			args: args{
				symbol: "BTC_USD",
			},
			want:    "BTC/USD",
			wantErr: false,
		},
		{
			name: "CASE-2",
			args: args{
				symbol: "BTCUSD",
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := SymbolPrepare(tt.args.symbol)
			t.Logf("err: %v", err)
			t.Logf("result: %v", result)
			if (err != nil) != tt.wantErr {
				t.Errorf("SymbolPrepare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if result != tt.want {
				t.Errorf("SymbolPrepare() got = %v, want %v", result, tt.want)
			}
		})
	}
}
