package main

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/chromedp"
)

// Парс Плеерока, все лоты с главной(еще не закончил)
func ParsePlayerok(ctx context.Context, ch chan []PlayerokLot) {

	var lots []PlayerokLot

	err := chromedp.Run(ctx,
		chromedp.Navigate("https://playerok.com/"),
		chromedp.WaitVisible("div[data-id]"),
		chromedp.Sleep(5*time.Second),
		chromedp.Evaluate(`
    const lots = document.querySelectorAll("div[data-id]");
    const result = [];

    for (const lot of lots) {
        const links = lot.querySelectorAll('a[href^="/products/"]');

        const name = links[1]?.innerText.trim() || "";
        const price = lot.querySelector(".text-text-accent-blue")?.innerText.trim() || "";
		const marketplace = "Playerok";

        result.push({
            name: name,
            price: price,
			marketplace: marketplace
        });
    }

    result
`, &lots),
	)

	if err != nil {
		fmt.Println(err)
	}

	ch <- lots
}
