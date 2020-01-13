package getfeed

import (
	"testing"

	"cursse/pkg/testdoubles"
)

func TestNewUseCase(t *testing.T) {
	if u := NewUseCase(
		[]string{},
		&testdoubles.RSSFeedFetcherStub{},
		&testdoubles.RSSFeedPresenterStub{},
	); u == nil {
		t.Error("NewUseCase returned nil")
	}
}

func TestGetFeeds(t *testing.T) {
	u := NewUseCase(
		[]string{},
		&testdoubles.RSSFeedFetcherStub{},
		&testdoubles.RSSFeedPresenterStub{},
	)
	if err := u.GetFeeds(); err != nil {
		t.Errorf("GetFeeds returned error: %v", err)
	}
}
