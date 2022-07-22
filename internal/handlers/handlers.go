package handlers

import (
	"fmt"
	"log"
	"strings"

	"github.com/pkg/errors"

	"gitlab.ozon.dev/tigprog/homeword-1/internal/commander"
	"gitlab.ozon.dev/tigprog/homeword-1/internal/storage"
	"gitlab.ozon.dev/tigprog/homeword-1/internal/tools"
)

const (
	helpCmd   = "help"
	listCmd   = "list"
	getCmd    = "get"
	addCmd    = "add"
	updateCmd = "update"
	deleteCmd = "delete"
)

var BadArgument = errors.New("bad argument")

func listFunc(s string) string {
	data := storage.List()
	res := make([]string, 0, len(data))
	for _, v := range data {
		res = append(res, v.String())
	}
	if len(res) == 0 {
		return "<empty list>"
	}
	return strings.Join(res, "\n")
}

func getFunc(data string) string {
	log.Printf("get command param: <%s>", data)
	params := strings.Split(data, " ")
	if len(params) != 1 {
		return errors.Wrapf(BadArgument, "%d items: <%v>", len(params), params).Error()
	}

	id, err := tools.StringToUint(params[0])
	if err != nil {
		return err.Error()
	}
	bb, err := storage.Get(id)
	if err != nil {
		return err.Error()
	}
	return bb.String()
}

func helpFunc(s string) string {
	return "/help - list commands\n" +
		"/list - list of bus bookings\n" +
		"/get <id> - get bus booking by id\n" +
		"/add <route> <date> <seat> - add new bus booking with route, date and seat\n" +
		"/update <id> <field> <new_value> - update field to new_value for bus by id\n" +
		"/delete <id> - delete bus booking by id"
}

func addFunc(data string) string {
	log.Printf("add command param: <%s>", data)
	params := strings.Split(data, " ")
	if len(params) != 3 {
		return errors.Wrapf(BadArgument, "%d items: <%v>", len(params), params).Error()
	}

	seat, err := tools.StringToUint(params[2])
	if err != nil {
		return errors.Wrap(err, params[2]).Error()
	}
	bb, err := storage.NewBusBooking(params[0], params[1], seat)
	if err != nil {
		return err.Error()
	}

	err = storage.Add(bb)
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("bus booking %v added", bb)
}

func updateFunc(data string) string {
	log.Printf("update command param: <%s>", data)
	params := strings.Split(data, " ")
	if len(params) != 3 {
		return errors.Wrapf(BadArgument, "%d items: <%v>", len(params), params).Error()
	}

	id, err := tools.StringToUint(params[0])
	if err != nil {
		return err.Error()
	}
	err = storage.Update(id, params[1], params[2])
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("bus booking <%d> updated", id)
}

func deleteFunc(data string) string {
	log.Printf("delete command param: <%s>", data)
	params := strings.Split(data, " ")
	if len(params) != 1 {
		return errors.Wrapf(BadArgument, "%d items: <%v>", len(params), params).Error()
	}

	id, err := tools.StringToUint(params[0])
	if err != nil {
		return err.Error()
	}
	err = storage.Delete(id)
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("bus booking <%d> deleted", id)
}

func AddHandlers(c *commander.Commander) {
	c.RegisterHandler(helpCmd, helpFunc)
	c.RegisterHandler(listCmd, listFunc)
	c.RegisterHandler(getCmd, getFunc)
	c.RegisterHandler(addCmd, addFunc)
	c.RegisterHandler(updateCmd, updateFunc)
	c.RegisterHandler(deleteCmd, deleteFunc)
}
