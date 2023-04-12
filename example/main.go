package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
	"github.com/yvv4git/ascendex"
)

func main() {
	app := &cli.App{
		Name:    "Ascendex client",
		Usage:   `The client for ascendex spot`,
		Version: "v0.0.1",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "symbol",
				Value:       "BTC_USDT",
				Usage:       "ascendex BTC_USDT",
				DefaultText: "BTC_USDT",
			},
			&cli.DurationFlag{
				Name:  "timeout",
				Value: time.Second * 10,
				Usage: "ascendex BTC_USDT 10",
			},
		},
		Action: func(cCtx *cli.Context) error {
			symbol := cCtx.String("symbol")
			timeout := cCtx.Duration("timeout")
			log.Printf("Symbol: %s  Timeout: %v", symbol, timeout)
			return run(symbol, timeout)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("something wrong: %v", err)
	}
}

func run(symbol string, timeout time.Duration) error {
	urlCreator := ascendex.NewAscendexURL()
	client := ascendex.NewClient(urlCreator)

	if err := client.Connection(); err != nil {
		return fmt.Errorf("error on init connection: %w", err)
	}

	if err := client.SubscribeToChannel(symbol); err != nil {
		return fmt.Errorf("error on subscribe to channel: %w", err)
	}

	go func() {
		for {
			client.WriteMessagesToChannel()
			time.Sleep(timeout)
		}
	}()

	ch := make(chan ascendex.BestOrderBook, 30)
	go client.ReadMessagesFromChannel(ch)

	for bestOrderBook := range ch {
		log.Printf(
			"Ask[price=%v amount=%v] Bid[price=%v amount=%v]",
			bestOrderBook.Ask.Price, bestOrderBook.Ask.Amount, bestOrderBook.Bid.Price, bestOrderBook.Bid.Amount)
	}

	return nil
}
