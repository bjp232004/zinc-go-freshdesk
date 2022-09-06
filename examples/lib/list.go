package lib

import (
	"fmt"
	"log"
	"os"

	"github.com/bjp232004/zinc-go-freshdesk/handlers"
	"github.com/joho/godotenv"
)

func ListTickets() {
	err := godotenv.Load("./examples/local.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	var domain = os.Getenv("FRESHDESK_DOMAIN")
	var apikey = os.Getenv("FRESHDESK_APIKEY")

	client, err := handlers.FreshDeskClient(domain, apikey)
	if err != nil {
		log.Fatalf("Could not create client: %s", err)
	}

	// // List tickets
	resp, err := client.Tickets().ListAll()
	if err != nil {
		log.Fatalf("failed to create ticket: %s", err)
	}

	for _, value := range resp {
		fmt.Print(value.ID, value.Subject)
		fmt.Print("\n")
	}
}
