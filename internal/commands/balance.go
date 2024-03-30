package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/everFinance/goar"
	"github.com/everFinance/goar/types"
	"github.com/liteseed/aogo"
	"github.com/urfave/cli/v2"
)

var Balance = &cli.Command{
	Name:  "balance",
	Usage: "Check the balance of the wallet",
	Flags: []cli.Flag{
		&cli.PathFlag{Name: "config", Aliases: []string{"c"}, Value: "./config.json", Usage: "path to config value"},
	},
	Action: balance,
}

func balance(context *cli.Context) error {

	var data = "Balance"
	var tags = []types.Tag{{Name: "Action", Value: "Balance"}}

	configPath := context.Path("config")
	configData, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalln(err)
	}

	var config Config

	err = json.Unmarshal(configData, &config)
	if err != nil {
		log.Fatalln(err)
	}

	ao, err := aogo.New()
	if err != nil {
		log.Fatal(err)
	}

	signer, err := goar.NewSignerFromPath(config.Signer)
	if err != nil {
		log.Fatal(err)
	}

	itemSigner, err := goar.NewItemSigner(signer)
	if err != nil {
		log.Fatal(err)
	}

	messageId, err := ao.SendMessage(config.Process, data, tags, "", itemSigner)
	if err != nil {
		log.Fatal(err)
	}

	result, err := ao.ReadResult(config.Process, messageId)
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Printf("Balance: %s BUN\n", result.Messages[0]["Data"])
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
