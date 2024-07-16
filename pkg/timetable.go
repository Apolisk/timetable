package timteable

import (
	"net/url"

	"github.com/gocolly/colly"
	tele "gopkg.in/telebot.v3"
)

const baseURL = "https://obukhivtrans.com.ua/assets/"

func links() ([]string, error) {
	c := colly.NewCollector()
	var links []string
	c.OnHTML("table", func(e *colly.HTMLElement) {
		links = e.ChildAttrs("a", "href")
	})

	err := c.Visit(baseURL)
	if err != nil {
		return nil, err
	}
	return links, nil
}

func TimeTable() (results []tele.Result, err error) {
	links, err := links()
	if err != nil {
		return nil, err
	}

	for _, link := range links[5:] {
		data, _ := url.PathUnescape(link)
		filetype := data[len(data)-3:]
		title := data[:len(data)-4]
		if filetype == "jpg" {
			article := &tele.ArticleResult{
				Title:    title,
				Text:     title,
				ThumbURL: baseURL + data,
				URL:      baseURL + data,
				HideURL:  true,
				ResultBase: tele.ResultBase{
					Content: &tele.InputTextMessageContent{
						Text: title,
						PreviewOptions: &tele.PreviewOptions{
							Disabled:   false,
							URL:        baseURL + data,
							SmallMedia: false,
							LargeMedia: false,
							AboveText:  false,
						},
					},
				},
			}
			results = append(results, article)
		}
	}
	return results, nil
}
