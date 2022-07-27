package main

import (
	"fmt"
	"math/rand"
)

var googleDomains = map[string]string{}

type SearchResults struct {
	ResultRank  int
	ResultURL   string
	ResultTitle string
	ResultDesc  string
}

var userAgents = []string{}

func randomUserAgent() {
	randNum := rand.Int() % len(userAgents)
	return userAgents[randNum]
}

func main() {
	res, err := GoogleScraper("Ichthoth")
	if err == nil {
		for _, res := range res {
			fmt.Println(res)
		}
	}
}
