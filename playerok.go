package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/chromedp/chromedp"
)

// Парс Плеерока, все лоты с главной(еще не закончил)
func ParsePlayerok(ctx context.Context, wg *sync.WaitGroup) {

	var lots []PlayerokLot

	defer wg.Done()

	err := chromedp.Run(ctx,
		chromedp.Navigate("https://playerok.com/"),
		chromedp.WaitVisible("div[data-id]"),
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

	for _, lot := range lots {
		fmt.Println("========== ЛОТ ==========")
		fmt.Println("Название:", lot.Name)
		fmt.Println("Цена:", lot.Price)
		fmt.Println("Маркетплейс:", lot.Marketplace)
	}
}
