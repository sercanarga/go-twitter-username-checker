package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/robfig/cron"
	"log"
	"os"
	"os/signal"
	"time"
	"twitter-profile-checker/helpers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	read, err := helpers.ReadAccountsJson(os.Getenv("ACCOUNTS_JSON"))
	if err != nil {
		log.Fatal(err)
	}

	cron := cron.New()
	cron.AddFunc("@every 10m", func() {
		for _, v := range read {
			request, err := helpers.GetProfile(helpers.ProfileStruct{
				APIUrl:      os.Getenv("API_URL"),
				ScreenName:  v,
				BearerToken: os.Getenv("AUTH_TOKEN"),
			})

			// send error
			if err != nil {
				helpers.SendErrorComplex(err)
				return
			}

			// send down account
			if request != 200 {
				helpers.SendDownAccount(v)
			} else {
				fmt.Println("Account is up: ", v)
			}

			time.Sleep(1 * time.Second)
		}
	})
	cron.Start()

	fmt.Println("tracker is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)
	<-sc
	os.Exit(0)

}
