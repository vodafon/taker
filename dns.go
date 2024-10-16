package taker

import (
	"math/rand"
	"strings"
	"time"

	"github.com/miekg/dns"
)

var (
	resolvers = []string{"8.8.8.8", "8.8.4.4", "1.1.1.1", "1.0.0.1", "9.9.9.10", "64.6.64.6"}
	timeout   = 10 * time.Second
)

func FindCNAME(host string) (string, error) {
	cname := ""
	c := dns.Client{Timeout: timeout}
	m := dns.Msg{}
	dnsServer := randomServer()
	m.SetQuestion(host+".", dns.TypeCNAME)
	r, _, err := c.Exchange(&m, dnsServer)
	if err != nil {
		return "", err
	}
	for _, ans := range r.Answer {
		if t, ok := ans.(*dns.CNAME); ok {
			cname = t.Target
		}
	}
	return cname, nil
}

func FindCNAMEs(host string) ([]string, error) {
	cnames := []string{}
	c := dns.Client{Timeout: timeout}
	m := dns.Msg{}
	dnsServer := randomServer()
	m.SetQuestion(host+".", dns.TypeA)
	r, _, err := c.Exchange(&m, dnsServer)
	if err != nil {
		return nil, err
	}
	for _, ans := range r.Answer {
		if t, ok := ans.(*dns.CNAME); ok {
			cnames = append(cnames, strings.TrimSuffix(t.Target, "."))
		}
	}
	return cnames, nil
}

func FindSOAAnsver(host string) ([]dns.RR, error) {
	c := dns.Client{Timeout: timeout}
	m := dns.Msg{}
	dnsServer := randomServer()
	m.SetQuestion(host+".", dns.TypeSOA)
	r, _, err := c.Exchange(&m, dnsServer)
	if err != nil {
		return nil, err
	}
	return r.Answer, nil
}

func randomServer() string {
	return resolvers[rand.Intn(len(resolvers))] + ":53"
}
