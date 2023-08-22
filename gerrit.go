package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gocolly/colly/v2"
)

// https://blog.logrocket.com/building-web-scraper-go-colly/

const (
	host = "pj1-gerrit"
	repo = "pj1/AOSP-manifest-release"

	// host = "pj2-gerrit"
	// repo = "pj2/AOSP-manifest-release"

	// host = "pj3-gerrit"
	// repo = "pj3/AOSP-manifest-release"

	loginpage    = "https://" + host + "/login/plugins/gitiles/" + repo
	releasesPage = "https://" + host + "/plugins/gitiles/" + repo
	logout       = "https://" + host + "/logout"
)

func main() {
	// main instance that has c.visit(url)
	c := colly.NewCollector()

	// set time out
	c.SetRequestTimeout(120 * time.Second)

	// Login
	err := c.Post(loginpage, map[string]string{"username": "your_username", "password": "your_password"})
	if err != nil {
		log.Fatal(err)
	}

	// Handle error while making request
	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e)
	})

	c.OnHTML(".RefList", func(e *colly.HTMLElement) {

		e.ForEach(".RefList-title", func(i int, h *colly.HTMLElement) {

			if h.Text == "Tags" {
				e.ForEach(".RefList-item", func(i int, h *colly.HTMLElement) {

					fmt.Println(h.Text)
				})
			}
		})
	})

	c.Visit(releasesPage)

}
