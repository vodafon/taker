package taker

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestAvailableDomain(t *testing.T) {
	hosts := LoadHosts("p0c.txt", t)
	for _, host := range hosts {
		available, tld, err := AvailableDomainCNAME(host)
		if err != nil {
			t.Fatal(err)
		}
		if available {
			fmt.Printf("domain %s available for %s\n", tld, host)
		}
	}
}

func TestAvailableDomainLink(t *testing.T) {
	links := LoadHosts("links.txt", t)
	for _, link := range links {
		available, tld, err := AvailableDomainLink(link)
		if err != nil {
			t.Fatal(err)
		}
		if available {
			fmt.Printf("domain %s available for %s\n", tld, link)
		}
	}
}

func LoadHosts(name string, t *testing.T) []string {
	file, err := os.Open("./testdata/" + name)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	urls := []string{}
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	return urls
}
