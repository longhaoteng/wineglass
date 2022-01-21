package rsa

import (
	"testing"
)

func TestRsaEncryptAndDecrypt(t *testing.T) {
	r, err := Gen()
	if err != nil {
		t.Errorf("GenRsaKey() error = %v", err)
		return
	}
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"1", args{s: "test"}, "test", false},
		{"2", args{s: "hello"}, "hello", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.Encrypt(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("RsaEncrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			want, err := r.Decrypt(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("RsaDecrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if want != tt.want {
				t.Errorf("RsaDecrypt() got = %v, want %v", want, tt.want)
			}
		})
	}
}
