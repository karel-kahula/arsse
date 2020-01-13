package testdoubles

import "cursse/pkg/entities"

type RSSFeedPresenterStub struct{}

func (f *RSSFeedPresenterStub) ShowFeeds([]entities.RSSFeed) error { return nil }
