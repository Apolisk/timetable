package bot

import (
	"strings"

	timteable "github.com/Apolisk/bus/pkg"
	tele "gopkg.in/telebot.v3"
)

func (b Bot) onQuery(c tele.Context) error {
	query := c.Query().Text

	results, err := timteable.TimeTable()
	if err != nil {
		return err
	}

	var filteredResults []tele.Result
	for _, result := range results {
		switch r := result.(type) {
		case *tele.ArticleResult:
			if strings.Contains(strings.ToLower(r.Title), strings.ToLower(query)) {
				filteredResults = append(filteredResults, result)
			}
		}
	}

	return c.Answer(&tele.QueryResponse{
		Results:   filteredResults,
		CacheTime: 60,
	})
}
