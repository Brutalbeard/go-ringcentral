package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/caarlos0/env"
	"github.com/grokify/gotilla/config"
	"github.com/grokify/gotilla/fmt/fmtutil"
	om "github.com/grokify/oauth2more"
	"github.com/jessevdk/go-flags"

	rc "github.com/brutalbeard/go-ringcentral/client"
	ru "github.com/brutalbeard/go-ringcentral/clientutil"
)

type Options struct {
	SubscriptionId string `short:"s" long:"subscriptionId" description:"SubscriptionId to delete" required:"true"`
}

type RingCentralConfig struct {
	TokenJSON  string `env:"RINGCENTRAL_TOKEN_JSON"`
	ServerURL  string `env:"RINGCENTRAL_SERVER_URL"`
	WebhookURL string `env:"RINGCENTRAL_WEBHOOK_URL"`
}

// This code takes a bot token and creates a permanent webhook.
func main() {
	opts := &Options{}
	_, err := flags.Parse(opts)
	if err != nil {
		panic(err)
	}

	err = config.LoadDotEnvSkipEmpty(os.Getenv("ENV_PATH"), "./.env")
	if err != nil {
		log.Fatal(err)
	}

	appCfg := RingCentralConfig{}
	err = env.Parse(&appCfg)
	if err != nil {
		log.Fatal(err)
	}

	httpClient, err := om.NewClientTokenJSON(
		context.Background(), []byte(appCfg.TokenJSON))
	if err != nil {
		log.Fatal(err)
	}

	apiClient, err := ru.NewApiClientHttpClientBaseURL(
		httpClient, appCfg.ServerURL)
	if err != nil {
		log.Fatal(err)
	}

	//msi := map[string]interface{}{}
	groups, resp, err := apiClient.GlipApi.LoadGroupList(context.Background(),
		&rc.LoadGroupListOpts{})
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode >= 300 {
		log.Fatal(fmt.Sprintf("API Status %v", resp.StatusCode))
	}

	fmt.Printf("Status: %v\n", resp.StatusCode)

	fmtutil.PrintJSON(groups)

	fmt.Println("DONE")
}
