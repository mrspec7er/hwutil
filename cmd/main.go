package main

import (
	"fmt"
	"time"

	"github.com/mrspec7er/hwutil/internal/monitor"
)

func main() {
	s := monitor.Service{}

	for {
		time.Sleep(time.Second * 3)
		func() {
			s.MemoryUtilization()

			s.CpuUtilization()

			s.DiskUtilization()

			fmt.Println()
		}()
	}

}
