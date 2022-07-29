package commander

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/tigprog/bus_booking/internal/config"
)

type CmdHandler func(string) string

type Interface interface {
	Run() error
	RegisterHandler(cmd string, f CmdHandler)
}

func MustNew() Interface {
	bot, err := tgbotapi.NewBotAPI(config.ApiKey)
	if err != nil {
		log.Panic(errors.Wrap(err, "init tgbot"))
	}
	bot.Debug = config.TelegramBotApiDebug
	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &commander{
		bot:   bot,
		route: make(map[string]CmdHandler),
	}
}

type commander struct {
	bot   *tgbotapi.BotAPI
	route map[string]CmdHandler
}

func (c *commander) RegisterHandler(cmd string, f CmdHandler) {
	c.route[cmd] = f
}

func (c *commander) Run() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = config.TelegramBotApiTimeout
	updates := c.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		if cmd := update.Message.Command(); cmd != "" {
			if f, ok := c.route[cmd]; ok {
				msg.Text = f(update.Message.CommandArguments())
			} else {
				msg.Text = "Unknown command"
			}
		} else {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			msg.Text = fmt.Sprintf("you send <%v>", update.Message.Text)
		}

		_, err := c.bot.Send(msg)
		if err != nil {
			return errors.Wrap(err, "send tg message")
		}
	}
	return nil
}
