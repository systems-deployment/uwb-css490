package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	maxConsecutiveErrors = 10
	remote               = "localhost:8080"
	//interval = 5 * time.Minute
	interval = 10 * time.Second
)

const (
	thresholdYellow = 90.0
	thresholdRed    = 95.0
)

var (
	disks = []string{"/dev/sda5"}
)

func Utilization(reader io.Reader) map[string]float64 {
	utilization := make(map[string]float64)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		for _, disk := range disks {
			if strings.HasPrefix(line, disk) {
				fields := strings.Fields(line)
				thisDisk := fields[0]
				utilizationStr := fields[4]
				thisUtilization, err := strconv.ParseFloat(utilizationStr[:len(utilizationStr)-1], 64)
				if err != nil {
					fmt.Fprintf(os.Stderr, "unable to parse line: \"%s\"\n\t%s\n", line, err)
					continue
				}
				utilization[thisDisk] = thisUtilization
			}
		}
	}
	return utilization
}

func main() {
	errorCount := 0
	client := http.Client{}
	ticker := time.Tick(interval)
	for {
		start := time.Now()
		fmt.Fprintf(os.Stderr, "%s: request\n", start)
		response, err := client.Get(fmt.Sprintf("http://%s/stats", remote))
		now := time.Now()
		fmt.Fprintf(os.Stderr, "%s: response received (%s elapsed) \n", now, now.Sub(start))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: request stats from client: %s\n", now, err)
			errorCount++
			if errorCount++; errorCount > maxConsecutiveErrors {
				fmt.Fprintf(os.Stderr, "Too many errors.  Terminating\n")
				os.Exit(1)
			}
		} else if response.StatusCode != 200 {
			fmt.Fprintf(os.Stderr, "%s: %s", now, response.Status)
		} else {
			//io.Copy(os.Stdout, response.Body)
			utilization := Utilization(response.Body)
			//fmt.Fprintf(os.Stdout, "%v\n", status)
			for theDisk, theUtilization := range utilization {
				if theUtilization > thresholdRed {
					fmt.Fprintf(os.Stdout, "RED   ")
				} else if theUtilization > thresholdYellow {
					fmt.Fprintf(os.Stdout, "YELLOW")
				} else {
					fmt.Fprintf(os.Stdout, "      ")
				}
				fmt.Fprintf(os.Stdout, "\t%15s\t%f\n", theDisk, theUtilization)
			}
		}
		<-ticker
	}
}
