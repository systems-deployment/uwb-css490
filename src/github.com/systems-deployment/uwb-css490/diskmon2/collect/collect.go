package collect

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

type Getter interface {
	Get(url string) (*http.Response, error)
}

var Client Getter = &http.Client{}

func Utilization(reader io.Reader, disks map[string]float64) error {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		for disk := range disks {
			if strings.HasPrefix(line, disk) {
				fields := strings.Fields(line)
				utilizationStr := fields[4]
				thisUtilization, err := strconv.ParseFloat(utilizationStr[:len(utilizationStr)-1], 64)
				if err != nil {
					return fmt.Errorf("parse error processing line: \"%s\"\n\t%s\"", line, err)
				}
				disks[disk] = thisUtilization
			}
		}
	}
	return nil
}

func Get(utilization map[string]float64, source string) error {
	request := fmt.Sprintf("http://%s/stats", source)
	start := time.Now()
	fmt.Fprintf(os.Stderr, "%s: request %s\n", start, request)
	response, err := Client.Get(request)
	now := time.Now()
	fmt.Fprintf(os.Stderr, "%s: elapsed time %s\n", now, now.Sub(start))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: request stats from client: %s\n", now, err)
		return err
	}
	//defer response.Body.Close()
	if response.StatusCode != 200 {
		err = fmt.Errorf("received status %s", response.Status)
		fmt.Fprintf(os.Stderr, "%s\t%s\n", now, err)
		return err
	}
	err = Utilization(response.Body, utilization)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\tparsing data: %s\n", time.Now(), err)
		return err
	}
	return nil
}
