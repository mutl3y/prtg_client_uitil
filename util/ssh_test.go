package util

import (
	"github.com/appleboy/easyssh-proxy"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// dockerfile included for integration testing /util/tests/Dockerfile
var (
	dest = SshStruct{
		User:     "prtgUtil",
		Server:   "localhost",
		Port:     "22",
		Password: "integrationTesting",
		Timeout:  time.Minute,
	}

	proxy = SshStruct{
		User:     "prtgUtil",
		Server:   "linuxserver",
		Port:     "1422",
		Password: "integrationTesting",
		Timeout:  time.Minute,
	}
)

func Test_main(t *testing.T) {

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

func Test_getUname(t *testing.T) {

	tests := []struct {
		name    string
		wantErr bool
	}{
		{"", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ssh := NewCon(dest, proxy)
			plat, err := ssh.getUname()
			if (err != nil) != tt.wantErr {
				t.Errorf("getUname() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Logf("%+v", plat)

		})
	}
}

func TestConn_Deploy(t *testing.T) {
	ssh := NewCon(dest, proxy)
	_ = ssh.Deploy("/releases/")
}

func Test_conn_CreateUsers(t *testing.T) {

	type args struct {
		tuser, tpass, juser, jpass string
		dest, proxy                SshStruct
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"pass", args{"root", "integrationTesting", "", "", proxy, easyssh.DefaultConfig{}}, false},
		{"fail", args{"root", "integrationTesting", "", "", dest, easyssh.DefaultConfig{}}, true},
		{"fail", args{"root", "integrationTestingfail", "", "", proxy, easyssh.DefaultConfig{}}, true},

		{"pass", args{"root", "integrationTesting", "root", "integrationTesting", dest, proxy}, false},
		{"fail", args{"root", "integrationTesting", "root", "integrationTestingfail", dest, proxy}, true},
		{"fail", args{"root", "integrationTesting", "root", "integrationTestingfail", dest, proxy}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ssh := NewCon(tt.args.dest, tt.args.proxy)
			if err := ssh.CreateUsers(tt.args.tuser, tt.args.tpass, tt.args.juser, tt.args.jpass); (err != nil) != tt.wantErr {
				t.Errorf("CreateUsers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_conn_createUser(t *testing.T) {
	ssh := NewCon(dest, proxy)

	type args struct {
		usr    string
		passwd string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"pass", args{usr: "prtgUtil", passwd: "integrationTesting"}, false},
		{"fail", args{usr: "prtgUtil", passwd: ""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ssh.createUser(tt.args.usr, tt.args.passwd); (err != nil) != tt.wantErr {
				t.Errorf("createUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_conn_Deploy(t *testing.T) {
	pwd, _ := os.Getwd()
	releases := filepath.Dir(pwd) + string(os.PathSeparator) + "releases"
	type args struct {
		dir         string
		dest, proxy SshStruct
	}

	var tests = []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"pass", args{releases, proxy, easyssh.DefaultConfig{}}, false},
		{"pass", args{releases, dest, proxy}, false},
		{"pass", args{"../releases", dest, proxy}, false},
		{"fail", args{"releases", dest, proxy}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ssh := NewCon(tt.args.dest, tt.args.proxy)
			if err := ssh.Deploy(tt.args.dir); (err != nil) != tt.wantErr {
				t.Errorf("Deploy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
