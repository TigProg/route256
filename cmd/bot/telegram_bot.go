package main

import (
	"log"

	botPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/bot"
	bbPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking"

	cmdAddPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/bot/command/add"
	cmdChangeDateSeatPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/bot/command/change_date_seat"
	cmdChangeSeatPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/bot/command/change_seat"
	cmdDeletePkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/bot/command/delete"
	cmdGetPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/bot/command/get"
	cmdHelpPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/bot/command/help"
	cmdListPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/bot/command/list"
)

func initBot(bb bbPkg.Interface) botPkg.Interface {
	bot := botPkg.MustNew()

	commandAdd := cmdAddPkg.New(bb)
	bot.RegisterHandler(commandAdd)

	commandList := cmdListPkg.New(bb)
	bot.RegisterHandler(commandList)

	commandGet := cmdGetPkg.New(bb)
	bot.RegisterHandler(commandGet)

	commandChangeSeat := cmdChangeSeatPkg.New(bb)
	bot.RegisterHandler(commandChangeSeat)

	commandChangeDateSeat := cmdChangeDateSeatPkg.New(bb)
	bot.RegisterHandler(commandChangeDateSeat)

	commandDelete := cmdDeletePkg.New(bb)
	bot.RegisterHandler(commandDelete)

	commandHelp := cmdHelpPkg.New(map[string]string{
		commandAdd.Name():            commandAdd.Description(),
		commandList.Name():           commandList.Description(),
		commandGet.Name():            commandGet.Description(),
		commandChangeSeat.Name():     commandChangeSeat.Description(),
		commandChangeDateSeat.Name(): commandChangeDateSeat.Description(),
		commandDelete.Name():         commandDelete.Description(),
	})
	bot.RegisterHandler(commandHelp)

	return bot
}

func runBot(bb bbPkg.Interface) {
	bot := initBot(bb)

	if err := bot.Run(); err != nil {
		log.Panic(err)
	}
}
