package getfeed

import "arsse/pkg/entities"

type RSSFeedPresenter interface {
	ShowFeeds([]entities.RSSFeed) error
}
