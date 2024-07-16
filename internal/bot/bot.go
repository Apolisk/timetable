package bot

import (
	"github.com/Apolisk/bus"
	"github.com/Apolisk/bus/internal/database"

	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/layout"
	"gopkg.in/telebot.v3/middleware"
)

type Bot struct {
	*tele.Bot
	*layout.Layout
	db *database.DB
}

func New(path string, boot bus.Bootstrap) (*Bot, error) {
	lt, err := layout.New(path)
	if err != nil {
		return nil, err
	}

	b, err := tele.NewBot(lt.Settings())
	if err != nil {
		return nil, err
	}

	if cmds := lt.Commands(); cmds != nil {
		if err := b.SetCommands(cmds); err != nil {
			return nil, err
		}
	}

	return &Bot{
		Bot:    b,
		Layout: lt,
		db:     boot.DB,
	}, nil
}

func (b *Bot) Start() {
	// Middlewares
	b.Use(middleware.Logger())
	b.Use(middleware.AutoRespond())
	b.Use(b.Middleware("en"))

	// Handlers
	b.Handle("/start", b.onStart)
	b.Handle(tele.OnQuery, b.onQuery)
	b.Bot.Start()
}
