package util

import (
	"fmt"
	"github.com/PaesslerAG/go-prtg-sensor-api"
	"github.com/appleboy/easyssh-proxy"
	"time"
)

type conn struct{ easyssh.MakeConfig }

type SshStruct = struct {
	User     string
	Server   string
	Key      string
	KeyPath  string
	Port     string
	Password string
	Timeout  time.Duration
}

func NewCon(dest, proxy SshStruct) *conn {
	c := easyssh.MakeConfig{}
	c.User = dest.User
	c.Server = dest.Server
	c.Key = dest.Key
	c.KeyPath = dest.KeyPath
	c.Port = dest.Port
	c.Password = dest.Password
	c.Timeout = dest.Timeout
	c.Proxy = proxy
	mc := conn{c}
	return &mc
}

func (ssh *conn) Remote(command string, timeout time.Duration) error {
	dir := "/var/prtg/scriptsxml/prtg_client_util "
	stdoutChan, stderrChan, doneChan, errChan, err := ssh.Stream(dir+command, timeout)
	// Handle errors
	if err != nil {
		FailRemote(fmt.Errorf("%v %v", err, dir))
		return fmt.Errorf("Can't run remote command: %v", err.Error())
	} else {
		// read from the output channel until the done signal is passed
		isTimeout := true
	loop:
		for {
			select {
			case isTimeout = <-doneChan:
				break loop
			case outline := <-stdoutChan:
				fmt.Println(outline)
			case errline := <-stderrChan:
				err = fmt.Errorf(errline)
			case err = <-errChan:
			}
		}

		// get exit code or command error.
		if err != nil {
			if err.Error() == "Process exited with status 127" {
				err = fmt.Errorf("command not found")
			}
			FailRemote(fmt.Errorf("%v %v", err, dir))
			return err
		}

		// command time out
		if !isTimeout {
			err := fmt.Errorf("error: command timeout")
			FailRemote(err)
			return err
		}
	}

	return nil
}

func FailRemote(err error) {
	r := prtg.SensorResponse{}
	r.SensorResult.Error = 1
	r.SensorResult.Text = fmt.Sprintf("%v", err)

	fmt.Println(r.String())
}
