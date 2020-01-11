package getfeed

import "fmt"

// UseCase handles polling feed urls for their contents
type UseCase struct {
	feedList  []string
	fetcher   RSSFeedFetcher
	presenter RSSFeedPresenter
}

func NewUseCase(feedList []string, fetcher RSSFeedFetcher, presenter RSSFeedPresenter) *UseCase {
	return &UseCase{
		feedList:  feedList,
		fetcher:   fetcher,
		presenter: presenter,
	}
}

func (u *UseCase) GetFeeds() error {
	feeds, err := u.fetcher.Fetch(u.feedList)
	if err != nil {
		return fmt.Errorf("error in fetcher: %w", err)
	}

	if err := u.presenter.ShowFeeds(feeds); err != nil {
		return fmt.Errorf("error in presenter: %v", err)
	}
	return nil
}
