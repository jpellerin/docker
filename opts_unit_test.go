package docker

import (
	"testing"
)

func TestValidateIP4(t *testing.T) {
	if ret, err := ValidateIp4Address(`1.2.3.4`); err != nil || !ret {
		t.Fatalf("ValidateIp4Address(`1.2.3.4`) got %s %s", ret, err)
	}


}
