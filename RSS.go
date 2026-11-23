package main

import (
	"strings"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

var pubDateLayouts = []string{
	"Mon, 2 Jan 2006 15:04:05 -0700",
	"Mon, 2 Jan 2006 15:04:05 MST",
	"2 Jan 2006 15:04:05 -0700",
	time.RFC3339,
}

func parsePubDate(s string) (time.Time, bool) {
	s = strings.TrimSpace(s)

	for _, layout := range pubDateLayouts {
		if pubDate, err := time.Parse(layout, s); err == nil {
			return pubDate, true
		}
	}

	return time.Time{}, false
}
