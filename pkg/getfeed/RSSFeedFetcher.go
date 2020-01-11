package getfeed

import "arsse/pkg/entities"

type RSSFeedFetcher interface {
	Fetch(feedList []string) ([]entities.RSSFeed, error)
}
