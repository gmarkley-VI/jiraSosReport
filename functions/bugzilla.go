package functions

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"time"
)

func ReturnBugs(team string) (string, string) {
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Make request
	teamurl := "https://shiftzilla.dpp.openshift.com/team_" + team + ".html"
	response, err := client.Get(teamurl)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	// Find All and Current Version bug counts
	all_bugs := document.Find(`body > div > div:nth-child(3) > div:nth-child(1) > table > tbody > tr:nth-child(2) > td:nth-child(2)`).Text()
	current_bugs := document.Find(`body > div > div:nth-child(3) > div:nth-child(1) > table > tbody > tr:nth-child(2) > td:nth-child(3)`).Text()

	return all_bugs, current_bugs
}
