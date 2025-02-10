package main

import (
	"context"
	"fmt"
	"time"

	chromedpundetected "github.com/Davincible/chromedp-undetected"
	"github.com/chromedp/chromedp"
)

func main() {

	cfg := chromedpundetected.NewConfig(
		chromedpundetected.WithTimeout(20*time.Second),
		chromedpundetected.WithHeadless(),
	)

	ctx, cancel, err := chromedpundetected.New(cfg)
	defer cancel()

	if err != nil {
		fmt.Println(err)
	}

	chromedp.Run(ctx,
		chromedp.Navigate("https://nowsecure.nl"),
		chromedp.WaitVisible(`//div[@class="hystericalbg"]`),
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Println("Started")
			return nil
		}),
	)

}
