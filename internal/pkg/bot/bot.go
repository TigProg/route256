package bot

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	configPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/config"
	commandPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/bot/command"
)

type Interface interface {
	Run() error
	RegisterHandler(cmd commandPkg.Interface)
}

func MustNew() Interface {
	bot, err := tgbotapi.NewBotAPI(configPkg.ApiKey)
	if err != nil {
		log.Panic(errors.Wrap(err, "init telegram bot"))
	}
	bot.Debug = configPkg.TelegramBotApiDebug
	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &commander{
		bot:   bot,
		route: make(map[string]commandPkg.Interface),
	}
}

type commander struct {
	bot   *tgbotapi.BotAPI
	route map[string]commandPkg.Interface
}

// RegisterHandler - not thread-safe
func (c *commander) RegisterHandler(cmd commandPkg.Interface) {
	c.route[cmd.Name()] = cmd
}

func (c *commander) Run() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = configPkg.TelegramBotApiTimeout
	updates := c.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		if cmdName := update.Message.Command(); cmdName != "" {
			if cmd, ok := c.route[cmdName]; ok {
				cmdArgs := update.Message.CommandArguments()
				log.Printf("Run [%s] command with args: <%s>", cmdName, cmdArgs)
				msg.Text = cmd.Process(cmdArgs)
			} else {
				log.Printf("Get unknown command: <%s>", cmdName)
				msg.Text = "Unknown command, try /help"
			}
		} else {
			log.Printf("Get plain text: [%s]", update.Message.Text)
			msg.Text = fmt.Sprintf("you send <%v>", update.Message.Text)
		}

		if _, err := c.bot.Send(msg); err != nil {
			return errors.Wrap(err, "send tg message")
		}
	}
	return nil
}
