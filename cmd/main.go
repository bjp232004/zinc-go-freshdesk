package main

import (
	"os"

	"github.com/bjp232004/zinc-go-freshdesk/cmd/lib"
)

func main() {
	funcname := os.Args[1]

	if funcname == "create" {
		lib.CreateTicket()
	}

	if funcname == "list" {
		lib.ListTickets()
	}
}
