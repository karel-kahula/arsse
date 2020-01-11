package testdoubles

import "arsse/pkg/entities"

type RSSFeedFetcherStub struct{}

func (f *RSSFeedFetcherStub) Fetch(feedList []string) ([]entities.RSSFeed, error) {
	return []entities.RSSFeed{}, nil
}
