package getfeed

import "arsse/pkg/entities"

type RSSFeedFetcher interface {
	Fetch(feeds []string) ([]entities.RSSFeed, error)
}
