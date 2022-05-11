package filesender

import (
	"log"
	"time"
)

type sender struct {
	config *Configuration
}

func NewSender(config *Configuration) *sender {
	return &sender{
		config: config,
	}
}

func (s *sender) Start(folder string) {
	ticker := time.NewTicker(time.Second * time.Duration(s.config.TickerDuration))
	defer ticker.Stop()

	for {
		<-ticker.C

		s.Schedule(folder)
	}
}

func check(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
