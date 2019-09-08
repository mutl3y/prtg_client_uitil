package sensor

import (
	"testing"
	"time"
)

func TestPing(t *testing.T) {
	type args struct {
		addr    string
		count   int
		timeout time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "ip", args: args{addr: "192.168.0.1", count: 3, timeout: time.Second}, wantErr: false},
		{name: "google.com", args: args{addr: "google.com", count: 5, timeout: time.Second}, wantErr: false},
		{name: "dead ip", args: args{addr: "192.168.32.123", count: 5, timeout: time.Second}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Ping(tt.args.addr, tt.args.count, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("Ping() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("%v %v", tt.args.addr, got)

		})
	}
}

func TestPrtg(t *testing.T) {
	type args struct {
		addr    []string
		count   int
		timeout time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "ip", args: args{addr: []string{"192.168.0.1"}, count: 3, timeout: time.Second}, wantErr: false},
		{name: "google.com", args: args{addr: []string{"google.com"}, count: 5, timeout: time.Second}, wantErr: false},
		{name: "dead ip", args: args{addr: []string{"google.com", "192.168.32.123"}, count: 5, timeout: time.Second}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PrtgPing(tt.args.addr, tt.args.count, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("Ping() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
