package bot

import (
	"time"

	"github.com/Apolisk/bus/internal/database"
	tele "gopkg.in/telebot.v3"
)

func (b Bot) onStart(c tele.Context) error {
	user := c.Sender()
	file, err := b.FileByID("BAACAgIAAxkBAAEalWZmoL57PB048KNOfY3t7pERZ3ES8QACjkwAAkSACUk_tVaWf9nQyzUE")
	if err != nil {
		return err
	}
	video := &tele.Video{
		Width:   270,
		Height:  480,
		File:    file,
		Caption: b.Text(c, "start", user),
	}

	if !b.db.HasUser(user.ID) {
		err := b.db.InsertUser(database.User{
			ID:        user.ID,
			FirstName: user.FirstName,
			CreatedAt: time.Now(),
		})
		if err != nil {
			return err
		}
	}

	return c.Send(
		video,
		tele.NoPreview,
	)
}
