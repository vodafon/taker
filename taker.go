package taker

import (
	"net/url"
	"strings"

	"golang.org/x/net/publicsuffix"
)

func AvailableDomain(host string) (bool, string, error) {
	tld, err := TLD(host)
	if err != nil {
		return false, "", err
	}
	if tld == "" {
		return false, "", nil
	}
	ans, err := FindSOAAnsver(tld)
	if err != nil {
		return false, "", err
	}
	if len(ans) == 0 {
		return true, tld, nil
	}
	return false, "", nil
}

func AvailableDomainCNAME(host string) (bool, string, error) {
	uri, err := url.Parse(hostToURL(host))
	if err != nil {
		return false, "", err
	}
	cname, err := FindCNAME(uri.Host)
	if err != nil {
		return false, "", err
	}
	if cname == "" {
		return false, "", nil
	}
	cname = strings.TrimSuffix(cname, ".")
	return AvailableDomain(cname)
}

func AvailableDomainLink(host string) (bool, string, error) {
	uri, err := url.Parse(hostToURL(host))
	if err != nil {
		return false, "", err
	}
	return AvailableDomain(uri.Host)
}

func TLD(host string) (string, error) {
	tld, icann := publicsuffix.PublicSuffix(host)
	if icann {
		return publicsuffix.EffectiveTLDPlusOne(host)
	}
	return tld, nil
}

func hostToURL(link string) string {
	if strings.HasPrefix(link, "http") {
		return link
	}
	if strings.HasPrefix(link, "//") {
		return "https:" + link
	}
	return "https://" + link
}
