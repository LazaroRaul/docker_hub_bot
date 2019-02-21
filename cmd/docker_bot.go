package main


import (
	"fmt"
	"log"
	"time"

	"github.com/yanzay/tbot"
)
func main() {
	token := "TOKEN"
	// Create new telegram bot server using token
	bot, err := tbot.NewServer(token)
	if err != nil {
		log.Fatal(err)
	}

	// Fill whitelist with users allowed to interact with bot
	//var whitelist []string
	// bot.AddMiddleware(tbot.NewAuth(whitelist))

	// Handle Connections with Github and Docker Hub Api
	bot.HandleFunc("/connect", Connect)

	// List Integrations with Github
	bot.HandleFunc("/listIntegrations", ListIntegrations)

	// Handle Delete specific integration
	bot.HandleFunc("/delIntegration", DelIntegration)

	// Handle status in Docker Hub of current repositories
	bot.HandleFunc("/status", HandleStatus)

	// Display Help
	bot.HandleFunc("/help", HelpHandler)

	fmt.Println("Server is Running!")
	// Start listening for messages
	err = bot.ListenAndServe()

	log.Fatal(err)
}

// Handle status in Docker Hub of current repositories
func HandleStatus(message *tbot.Message) {
	message.Reply("We should display status in Docker Hub of current repositories")
}

// Display Help
func HelpHandler(message *tbot.Message) {
	message.Reply("We should display some help")

}

// Handle new connections
func Connect(message *tbot.Message)  {
	// Implement Connections
	message.Replyf("Hello, %s!", message.From)
	message.Reply("We should connect now with github and docker hub")
	// Something with ConnectGithub
	// Something with ConnectDockerHub
}

// Handle Github Connections
func ConnectGithub(message *tbot.Message) {
	time.Sleep(1 * time.Second)
	message.Reply("We should connect now with github")
}
// Handle Docker Hub Connections
func ConnectDockerHub(message *tbot.Message) {
	time.Sleep(1 * time.Second)
	message.Reply("We should connect now with docker hub")

}

// Handle all current integrations
func ListIntegrations(message *tbot.Message) {
	// Should list all the integrations with github and docker hub
	message.Reply("We should list all integrations")
}

// Handle Delete specific integration
func DelIntegration(message *tbot.Message) {
	// Should list all the integrations with github and docker hub
	message.Reply("We should delete an integration with github")

}
