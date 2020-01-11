package entities

type RSSFeed struct {
	url      string
	contents string
}

func NewRSSFeed(url string) *RSSFeed {
	return &RSSFeed{url: url}
}

func (r *RSSFeed) GetContents() string { return r.contents }
func (r *RSSFeed) GetURL() string      { return r.url }
