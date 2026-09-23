package main

import (
	"context"
	"sync"

	"github.com/chromedp/chromedp"
)

type Lot struct {
	Service     string
	Seller      string
	Amount      string
	Type        string
	Price       string
	Marketplace string
}

type PlayerokLot struct {
	Name        string
	Price       string
	Marketplace string
}

func main() {
	//Настройки для создания Аллокатора
	opts := append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", false))

	allocCtx, allocCancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer allocCancel()

	funpayCtx, funpayCancel := chromedp.NewContext(allocCtx)
	defer funpayCancel()

	playerokCtx, playerokCancel := chromedp.NewContext(allocCtx)
	defer playerokCancel()

	var wg sync.WaitGroup

	wg.Add(2)
	go ParsePlayerok(playerokCtx, &wg)
	go ParseFunpay(funpayCtx, &wg)
	wg.Wait()

}
