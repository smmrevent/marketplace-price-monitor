package main

import (
	"context"
	"fmt"

	"github.com/chromedp/chromedp"
)

// Парс фанпея, все лоты с конкретной страницы
func ParseFunpay(ctx context.Context, ch chan []Lot) {
	var lots []Lot

	err := chromedp.Run(ctx,
		chromedp.Navigate("https://funpay.com/lots/4432/"),
		chromedp.Evaluate(`
    const lots = document.querySelectorAll(".tc-item");
    const result = [];

    for (const lot of lots) {
        const service = lot.querySelector(".tc-desc-text")?.innerText.trim() || "";
        const seller = lot.querySelector(".media-user-name")?.innerText.trim() || "";
        const amount = lot.querySelector(".tc-amount")?.innerText.trim() || "";
        const type = lot.getAttribute("data-f-type") || "";
        const price = lot.querySelector(".tc-price")?.innerText.trim() || "";
		const market = "Funpay";

        result.push({
            service: service,
            seller: seller,
            amount: amount,
            type: type,
            price: price,
			marketplace: market
        });
    }
	
    result`, &lots),
	)
	if err != nil {
		fmt.Println(err)
	}

	ch <- lots
}
