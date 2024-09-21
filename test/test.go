package main

import (
	"fmt"
	"net/http"
	"time"
)

func ping(domain string) bool {
	client := http.Client{
		Timeout: time.Second * 1,
	}

	// send a head request to avoid downloading the entire content
	resp, err := client.Head(domain)
	if err != nil {
		fmt.Println("error", err)
		return false
	}

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header)
	fmt.Println(resp.Cookies())
	fmt.Println(resp.Proto)
	fmt.Println(resp.Uncompressed)
	fmt.Println(resp.ProtoMinor)
	fmt.Println(resp.ProtoMajor)
	fmt.Println(resp.Body)
	fmt.Println(resp.ContentLength)
	// if the status code is 200 (ok) domain is reachable
	if resp.StatusCode == http.StatusOK {
		fmt.Println("domain found")
		return true
	}

	fmt.Println("failed with status code", resp.StatusCode)
	return false
}

func main() {
	domain := "https://learn.zone01kisumu.ke" // the domain to check
	if ping(domain) {
		fmt.Println("internet is available")
	} else {
		fmt.Println("intenet error")
	}
}

// https://www.perplexity.ai/search/body-body-background-color-fdf-_dnZSgcfQp68c2MJKuo54Q
