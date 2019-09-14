package sensor

import (
	"testing"
	"time"
)

func TestPing(t *testing.T) {
	type args struct {
		addr              string
		count, size       int
		timeout, interval time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "ip", args: args{addr: "192.168.0.1", count: 3, size: 32, timeout: time.Second, interval: time.Millisecond * 20}, wantErr: false},
		{name: "google.com", args: args{addr: "google.com", count: 5, size: 32, timeout: time.Second, interval: time.Millisecond * 20}, wantErr: false},
		{name: "dead ip", args: args{addr: "192.168.32.123", count: 5, size: 32, timeout: time.Second, interval: time.Millisecond * 20}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Ping(tt.args.addr, tt.args.count, tt.args.size, tt.args.timeout, tt.args.interval)
			if (err != nil) != tt.wantErr {
				t.Errorf("Ping() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("%v %v", tt.args.addr, got)

		})
	}
}

func TestPrtg(t *testing.T) {
	//Debug = true
	type args struct {
		addr              []string
		count, size       int
		timeout, interval time.Duration
		op                string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "ip", args: args{addr: []string{"192.168.0.1", " 8.8.8.8", "8.8.4.4"}, count: 3, size: 32, timeout: time.Second, interval: time.Millisecond * 20}, wantErr: false},
		{name: "google.com", args: args{addr: []string{"google.com"}, count: 5, size: 32, timeout: time.Second, interval: time.Millisecond * 250}, wantErr: false},
		{name: "dead ip", args: args{addr: []string{"google.com", "192.168.32.123"}, count: 5, size: 32, timeout: time.Second, interval: time.Millisecond * 20}, wantErr: true},

		{name: "loss ip", args: args{addr: []string{"192.168.0.1"}, count: 3, size: 32, timeout: time.Second, interval: time.Millisecond * 20, op: "loss"}, wantErr: false},
		{name: "loss google.com", args: args{addr: []string{"google.com"}, count: 5, size: 32, timeout: time.Second, interval: time.Millisecond * 250, op: "loss"}, wantErr: false},
		{name: "loss dead ip", args: args{addr: []string{"google.com", "192.168.32.123"}, count: 5, size: 32, timeout: time.Second, interval: time.Millisecond * 20, op: "loss"}, wantErr: true},

		{name: "everything", args: args{addr: []string{"8.8.8.8"}, count: 3, size: 32, timeout: time.Second, interval: time.Millisecond * 20, op: "everything"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PrtgPing(tt.args.addr, tt.args.count, tt.args.size, tt.args.timeout, tt.args.interval, tt.args.op)
			if (err != nil) != tt.wantErr {
				t.Errorf("Ping() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
