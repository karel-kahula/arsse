package web

import (
	"cursse/pkg/entities"
)

type rssFeedFetcherImpl struct {
	feedURLs []string
}

func NewRSSFeedFetcher() *rssFeedFetcherImpl {
	return &rssFeedFetcherImpl{}
}

func (f *rssFeedFetcherImpl) Fetch(feeds []string) ([]entities.RSSFeed, error) {
	return []entities.RSSFeed{}, nil
}
