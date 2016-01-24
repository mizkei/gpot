package gpot

import (
	"fmt"
	"golang.org/x/oauth2"
	plusd "google.golang.org/api/plusdomains/v1"
	"log"
)

func init() {
	log.SetFlags(log.Llongfile)
}

// BotPlus is simple bot.
type BotPlus struct {
	service *plusd.Service
}

// Post post comment
func (b *BotPlus) Post(msg string) error {
	circlesvc := plusd.NewCirclesService(b.service)
	call := circlesvc.List("me")
	feed, err := call.Do()
	if err != nil {
		log.Fatal(err)
		return err
	}
	bt, err := feed.MarshalJSON()
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("res: ", string(bt))
	return nil
}

// EventTicker returns channel that send empty
// when activity posted.
func (b *BotPlus) EventTicker() <-chan struct{} {
	ch := make(chan struct{})

	return ch
}

// NewBotPlus generate BotPlus.
func NewBotPlus(config oauth2.Config) *BotPlus {
	url := config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("visit the URL:\n %v\n", url)

	fmt.Print("input code: ")
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}

	tkn, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatal(err)
	}

	client := config.Client(oauth2.NoContext, tkn)
	service, err := plusd.New(client)
	if err != nil {
		log.Fatal(err)
	}

	return &BotPlus{
		service: service,
	}
}
