package sensor

import (
	"fmt"
	"github.com/PaesslerAG/go-prtg-sensor-api"
	"github.com/sparrc/go-ping"
	"time"
)

func Ping(addr string, count int, timeout time.Duration) (*ping.Statistics, error) {
	pinger, err := ping.NewPinger(addr)
	if err != nil {
		return &ping.Statistics{}, err
	}

	//	if runtime.GOOS == "windows" {
	pinger.SetPrivileged(true)
	//	}

	pinger.Timeout = timeout
	pinger.Count = count
	pinger.Run() // blocks until finished

	s := pinger.Statistics() // get send/receive/rtt stats
	if s.PacketLoss == float64(100) {
		return s, fmt.Errorf("%v 100%% packet loss", s.IPAddr)
	}

	return s, nil
}

func PrtgPing(addr []string, count int, timeout time.Duration) error {
	r := new(prtg.SensorResponse)

	for _, v := range addr {
		s, err := Ping(v, count, timeout)
		if err != nil {
			r.SensorResult.Error = 1
			r.SensorResult.Text = fmt.Sprintf("%v", err)
			fmt.Println(r.String())
			return err
		}

		r.AddChannel(prtg.SensorChannel{
			Name:      fmt.Sprintf("%v", v),
			Value:     s.AvgRtt.Truncate(time.Millisecond).Seconds() * 1000,
			Float:     1,
			ShowChart: show,
			ShowTable: show,
			Unit:      prtg.UnitTimeResponse,
		})
	}
	fmt.Println(r.String())
	return nil
}

func PrtgPacketLoss(addr []string, count int, timeout time.Duration) error {
	r := new(prtg.SensorResponse)

	for _, v := range addr {
		s, err := Ping(v, count, timeout)
		if err != nil {
			r.SensorResult.Error = 1
			r.SensorResult.Text = fmt.Sprintf("%v", err)
			fmt.Println(r.String())
			return err
		}

		r.AddChannel(prtg.SensorChannel{
			Name:      fmt.Sprintf("%v", v),
			Value:     s.PacketLoss,
			Float:     1,
			ShowChart: show,
			ShowTable: show,
			Unit:      prtg.UnitPercent,
		})
	}
	fmt.Println(r.String())
	return nil
}
