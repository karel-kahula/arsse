package getfeed

import "cursse/pkg/entities"

type RSSFeedFetcher interface {
	Fetch(feeds []string) ([]entities.RSSFeed, error)
}
