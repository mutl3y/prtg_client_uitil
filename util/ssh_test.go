package util

import (
	"testing"
	"time"
)

func Test_main(t *testing.T) {
	dest := SshStruct{
		User:     "prtgUtil",
		Server:   "localhost",
		Key:      "",
		KeyPath:  "",
		Port:     "22",
		Password: ",.password",
		Timeout:  0,
	}

	proxy := SshStruct{
		User:     "prtgUtil",
		Server:   "linuxserver",
		Port:     "22",
		Password: ",.password",
	}

	type args struct {
		dest, proxy SshStruct
		timeout     time.Duration
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"1", args{dest: dest, proxy: proxy, timeout: 60 * time.Second}, false},
		{"2", args{dest: dest, proxy: SshStruct{}, timeout: 60 * time.Second}, true},
		{"3", args{dest: dest, proxy: proxy, timeout: 1 * time.Millisecond}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ss := NewCon(tt.args.dest, tt.args.proxy)
			err := ss.Remote("ping", tt.args.timeout)
			if err != nil && !tt.wantErr {
				t.Errorf("failed %v", err)
			}
		})
	}
}
