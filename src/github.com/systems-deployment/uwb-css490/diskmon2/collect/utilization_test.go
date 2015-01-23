package collect

import (
	//"fmt"
	"strings"
	"testing"
)

func check(t *testing.T, m map[string]float64, disk string, expected float64) {
	actual, ok := m[disk]
	if !ok {
		t.Errorf("missing %s", disk)
	}
	if actual != expected {
		t.Errorf("disk %s: expected %f, got %f", disk, expected, actual)
	}
}

func TestUtilization(t *testing.T) {
	in := strings.NewReader(`
/dev/disk0s2    250G   114G   136G    46% 27795469 33273971   46%   /
/dev/disk0s3    250G   114G   136G    99% 27795469 33273971   88%   /mnt
`)
	out := map[string]float64{
		"/dev/disk0s2": 0.0,
		"/dev/disk0s3": 0.0,
		"/dev/disk0s4": 0.0,
	}
	err := Utilization(in, out)
	if err != nil {
		t.Errorf("parse error %s", err)
	}
	if len(out) != 3 {
		t.Errorf("expected 3 values, received %d", len(out))
	}
	check(t, out, "/dev/disk0s2", 46.0)
	check(t, out, "/dev/disk0s3", 99.0)
	check(t, out, "/dev/disk0s4", 0.0)
}
