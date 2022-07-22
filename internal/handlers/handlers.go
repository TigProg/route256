package handlers

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/tigprog/homeword-1/internal/commander"

	"gitlab.ozon.dev/tigprog/homeword-1/internal/storage"
)

const (
	helpCmd   = "help"
	listCmd   = "list"
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

func helpFunc(s string) string {
	return "/help - list commands\n" +
		"/list - list data\n" +
		"/add <route> <date> <seat> - add new bus booking with route, date and seat\n" +
		"/update <id> <field> <new_value> - update field to new_value for bus by id\n" +
		"/delete <id> - delete bus booking by id"
}

func addFunc(data string) string {
	log.Printf("add command param: <data>")
	params := strings.Split(data, " ")
	if len(params) != 3 {
		return errors.Wrapf(BadArgument, "%d items: <%v>", len(params), params).Error()
	}
	seatUint64, err := strconv.ParseUint(params[2], 10, 32)
	if err != nil {
		return err.Error()
	}
	bb, err := storage.NewBusBooking(params[0], params[1], uint(seatUint64))
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
	log.Printf("update command param: <data>")
	params := strings.Split(data, " ")
	if len(params) != 3 {
		return errors.Wrapf(BadArgument, "%d items: <%v>", len(params), params).Error()
	}
	idUint64, err := strconv.ParseUint(params[0], 10, 32)
	if err != nil {
		return err.Error()
	}
	err = storage.Update(uint(idUint64), params[1], params[2])
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("bus booking <%d> updated", idUint64)
}

func deleteFunc(data string) string {
	log.Printf("delete command param: <data>")
	params := strings.Split(data, " ")
	if len(params) != 1 {
		return errors.Wrapf(BadArgument, "%d items: <%v>", len(params), params).Error()
	}
	idUint64, err := strconv.ParseUint(params[0], 10, 32)
	if err != nil {
		return err.Error()
	}
	err = storage.Delete(uint(idUint64))
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("bus booking <%d> deleted", idUint64)
}

func AddHandlers(c *commander.Commander) {
	c.RegisterHandler(helpCmd, helpFunc)
	c.RegisterHandler(listCmd, listFunc)
	c.RegisterHandler(addCmd, addFunc)
	c.RegisterHandler(updateCmd, updateFunc)
	c.RegisterHandler(deleteCmd, deleteFunc)
}
