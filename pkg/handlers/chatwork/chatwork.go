package chatwork

import (
	"fmt"
	"log"
	"net/url"
	"os"

	chatwork "github.com/yoppi/go-chatwork"

	"github.com/bitnami-labs/kubewatch/config"
	"github.com/bitnami-labs/kubewatch/pkg/event"
)

var chatworkErrMsg = `
%s

You need to set both chatwork token and room for chatwork notify,
using "--token/-t", "--room/-r", and "--url/-u" or using environment variables:

export KW_CHATWORK_TOKEN=chatwork_token
export KW_CHATWORK_ROOM=chatwork_room
export KW_CHATWORK_URL=chatwork_url (defaults to https://api.chatwork.com/v2)

Command line flags will override environment variables

`

// Chatwork handler implements handler.Handler interface,
// Notify event to chatwork room
type Chatwork struct {
	Token string
	Room  string
	Url   string
}

// Init prepares chatwork configuration
func (s *Chatwork) Init(c *config.Config) error {
	baseUrl := c.Handler.Chatwork.Url
	room := c.Handler.Chatwork.Room
	token := c.Handler.Chatwork.Token

	if token == "" {
		token = os.Getenv("KW_CHATWORK_TOKEN")
	}

	if room == "" {
		room = os.Getenv("KW_CHATWORK_ROOM")
	}

	if baseUrl == "" {
		baseUrl = os.Getenv("KW_CHATWORK_URL")
	}

	s.Token = token
	s.Room = room
	s.Url = baseUrl

	return checkMissingChatworkVars(s)
}

// Handle handles the notification.
func (s *Chatwork) Handle(e event.Event) {
	client := chatwork.NewClient(s.Token)
	if s.Url != "" {
		_, err := url.Parse(s.Url)
		if err != nil {
			panic(err)
		}
		client.BaseUrl = s.Url
	}

	notification := fmt.Sprintf("%s From K8S Cluster By KubeWatch", e.Message())
	result := client.PostRoomMessage(s.Room, notification)

	log.Println(string(result))

	log.Printf("Message successfully sent to room %s", s.Room)
}

func checkMissingChatworkVars(s *Chatwork) error {
	if s.Token == "" || s.Room == "" {
		return fmt.Errorf(chatworkErrMsg, "Missing chatwork token or room")
	}

	return nil
}
