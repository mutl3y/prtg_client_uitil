package sensor

import (
	"fmt"
	"github.com/PaesslerAG/go-prtg-sensor-api"
	"github.com/beevik/ntp"
	"time"
)

func init() { *show = 1 }

func NtpCheck(addr string, timeout time.Duration) (time.Duration, *ntp.Response, error) {
	start := time.Now()
	response, err := ntp.QueryWithOptions(addr, ntp.QueryOptions{
		Timeout: timeout,
		TTL:     50,
	})
	elapsed := time.Since(start)
	if err != nil {
		return elapsed, nil, fmt.Errorf("%v", err)
	}

	err = response.Validate()
	if err != nil {
		return elapsed, response, fmt.Errorf("response data is not suitable for synchronization purposes  %v", err)
	}
	return elapsed, response, nil
}

func PrtgNtp(a string, timeout, drift time.Duration) error {

	// Create empty response and log start time
	r := &prtg.SensorResponse{}
	dur, res, err := NtpCheck(a, timeout)
	if err != nil {
		r.SensorResult.Error = 1
		r.SensorResult.Text = fmt.Sprintf("%v", err)
		fmt.Println(r.String())
		return err
	}
	r.AddChannel(prtg.SensorChannel{
		Name:      "response time",
		Value:     dur.Truncate(time.Millisecond).Seconds() * 1000,
		Float:     1,
		ShowChart: show,
		ShowTable: show,
		Unit:      prtg.UnitTimeResponse,
	})
	r.AddChannel(prtg.SensorChannel{
		Name:      "offset",
		Value:     res.ClockOffset.Truncate(time.Millisecond).Seconds() * 1000,
		Float:     1,
		ShowChart: show,
		ShowTable: show,
		Unit:      prtg.UnitTimeResponse,
	})
	r.AddChannel(prtg.SensorChannel{
		Name:      "rtt",
		Value:     res.RTT.Truncate(time.Millisecond).Seconds() * 1000,
		Float:     1,
		ShowChart: show,
		ShowTable: show,
		Unit:      prtg.UnitTimeResponse,
	})
	r.AddChannel(prtg.SensorChannel{
		Name:      "precision",
		Value:     (res.Precision * 1000).Truncate(time.Microsecond).Seconds(),
		Float:     1,
		ShowChart: show,
		ShowTable: show,
		Unit:      prtg.UnitTimeResponse,
	})

	if drift > 0 {
		if res.ClockOffset >= drift {
			r.SensorResult.Error = 1
			txt := fmt.Sprintf("offset returned %v greater than allowed drift %v", res.ClockOffset, drift)
			r.SensorResult.Text = txt
			fmt.Println(r.String())
			return fmt.Errorf(txt)
		}
	}

	fmt.Println(r.String())
	return nil
}
