package taker

import (
	"fmt"
	"testing"
)

func TestDoDNSCNAME(t *testing.T) {
	hosts := LoadHosts("p0c.txt", t)
	for _, host := range hosts {
		cname, err := FindCNAME(host)
		if err != nil {
			t.Fatal(err)
		}
		if cname != "" {
			fmt.Printf("%s: %s\n", host, cname)
		}
	}
}
