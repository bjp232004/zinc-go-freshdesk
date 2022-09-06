package lib

import (
	"log"
	"os"

	"github.com/bjp232004/zinc-go-freshdesk/handlers"
	"github.com/joho/godotenv"
)

func CreateTicket() {
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

	// Create Ticket
	ticket := &handlers.Ticket{
		Priority:    1,
		RequesterID: 88000825587,
		ResponderID: 88000825587,
		Source:      1,
		Status:      2,
		Subject:     "Ticket Title3",
		Name:        "Ticket Title1",
		Type:        "Incident",
		Email:       "fakeemail@noreply.com",
		Description: "this is a sample ticket",
	}

	if _, err := client.Tickets().Create(ticket); err != nil {
		log.Fatalf("failed to create ticket: %s", err)
	}
}
