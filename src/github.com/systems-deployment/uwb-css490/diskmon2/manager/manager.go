package manager

import (
	"fmt"
	"github.com/systems-deployment/uwb-css490/diskmon2/alert"
	"github.com/systems-deployment/uwb-css490/diskmon2/collect"
	"os"
	"time"
)

const (
	maxConsecutiveErrors = 10
	remote               = "localhost:8080"
	//interval = 5 * time.Minute
	//interval = 10 * time.Second
	interval = 5 * time.Second
)

const (
	thresholdYellow = 90.0
	thresholdRed    = 95.0
)

var (
	alerts = make(map[string]alert.Alert)
)

func statsFromList(disks []string) map[string]float64 {
	stats := make(map[string]float64)
	for _, d := range disks {
		stats[d] = 0.0
	}
	return stats
}

func Monitor(disks []string) error {
	errorCount := 0
	ticker := time.Tick(interval)
	for {
		stats := statsFromList(disks)
		err := collect.Get(stats, remote)
		if err != nil {
			errorCount++
			if errorCount >= maxConsecutiveErrors {
				err := fmt.Errorf("Too many errors\n")
				return err
			}
		} else {
			errorCount = 0
			for disk, utilization := range stats {
				fmt.Fprintf(os.Stderr, "utilization: %s\t%f\n", disk, utilization)
				if utilization > thresholdRed {
					if thisAlert := alerts[disk]; thisAlert != nil {
						if thisAlert.Level() != alert.Red {
							thisAlert.Reset(alert.Red, fmt.Sprintf("%s: over %5.1f% (%.1f%%)",
								disk, thresholdRed, utilization))
						}
					} else {
						alerts[disk] = alert.New(alert.Red, fmt.Sprintf("%s: over %.1f%% (%.1f%%)",
							disk, thresholdRed, utilization))
					}
				} else if utilization > thresholdYellow {
					if thisAlert := alerts[disk]; thisAlert != nil {
						if thisAlert.Level() != alert.Yellow {
							thisAlert.Reset(alert.Yellow, fmt.Sprintf("%s: over %.1f%% (%.1f%%)",
								disk, thresholdYellow, utilization))
						}
					} else {
						alerts[disk] = alert.New(alert.Yellow, fmt.Sprintf("%s: over %.1f%% (%.1f%%)",
							disk, thresholdYellow, utilization))
					}
				} else {
					if thisAlert := alerts[disk]; thisAlert != nil {
						thisAlert.Reset(alert.Clear, "")
						alerts[disk] = nil
					}
				}
			}
		}
		<-ticker
	}
}
