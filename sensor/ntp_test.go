package sensor

import (
	"testing"
	"time"
)

func TestNtpCheck(t *testing.T) {
	type args struct {
		addr           string
		timeout, drift time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"", args{addr: "1.beevik-ntp.pool.ntp.org", timeout: 5000 * time.Millisecond}, false},
		{"", args{addr: "time.google.com", timeout: 5000 * time.Millisecond}, false},
		{"", args{addr: "0.beevik-ntp.pool.nt1p.org", timeout: 5000 * time.Millisecond}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := NtpCheck(tt.args.addr, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("NtpCheck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestPrtgNtp(t *testing.T) {
	type args struct {
		a              string
		timeout, drift time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"", args{a: "1.beevik-ntp1.pool.ntp.org", timeout: 5000 * time.Millisecond}, true},
		{"", args{a: "time.google.com", timeout: 5000 * time.Millisecond, drift: time.Second}, false},
		{"", args{a: "time.google.com", timeout: 5000 * time.Millisecond, drift: time.Nanosecond}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PrtgNtp(tt.args.a, tt.args.timeout, tt.args.drift); (err != nil) != tt.wantErr {
				t.Errorf("PrtgNtp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
