package main

import "log"

import (
	"arsse/pkg/getfeed"
	"arsse/pkg/web"
	//"arse/pkg/ui/stdout"
)

func main() {
	feeds := []string{}
	feedFetcher := web.NewRSSFeedFetcher()
	//feedPresenter := stdout.NewRSSFeedPresenter()
	getFeed := getfeed.NewUseCase(feeds, feedFetcher, nil)
	if err := getFeed.GetFeeds(); err != nil {
		log.Fatalf("Error in getFeed: %v", err)
	}
}
