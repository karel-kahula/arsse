package getfeed

import "cursse/pkg/entities"

type RSSFeedPresenter interface {
	ShowFeeds([]entities.RSSFeed) error
}
