package manager

import (
	"fmt"
	"github.com/systems-deployment/uwb-css490/diskmon2/collect"
	"io"
	"net/http"
	"strings"
	"testing"
)

type mockClient struct {
	index     int
	responses []string
}

type ReadCloser struct {
	reader io.Reader
}

var Done chan bool

func init() {
	Done = make(chan bool)
	collect.Client = &mockClient{
		responses: []string{
			"/dev/disk0s2    250G   114G   136G    46% 27795469 33273971   46%   /\n",
			"/dev/disk0s2    250G   114G   136G    80% 27795469 33273971   46%   /\n",
			"/dev/disk0s2    250G   114G   136G    92% 27795469 33273971   46%   /\n",
			"/dev/disk0s2    250G   114G   136G    96% 27795469 33273971   46%   /\n",
			"/dev/disk0s2    250G   114G   136G    92% 27795469 33273971   46%   /\n",
			"/dev/disk0s2    250G   114G   136G    80% 27795469 33273971   46%   /\n",
			"/dev/disk0s2    250G   114G   136G    90% 27795469 33273971   46%   /\n",
		},
	}
}

func (this *mockClient) Get(url string) (*http.Response, error) {
	if this.index >= len(this.responses) {
		Done <- true
		return nil, fmt.Errorf("out of responses")
	}
	response := &http.Response{
		StatusCode: 200,
		Body: &ReadCloser{
			reader: strings.NewReader(this.responses[this.index]),
		},
	}
	this.index++
	return response, nil
}

func (this *ReadCloser) Read(p []byte) (int, error) {
	return this.reader.Read(p)
}

func (this *ReadCloser) Close() error {
	return nil
}

func TestMonitor(t *testing.T) {
	go Monitor([]string{"/dev/disk0s2"})
	<-Done
}
