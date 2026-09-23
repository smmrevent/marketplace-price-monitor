package main

import (
	"context"
	"fmt"

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

	var funpayLots []Lot
	var playerokLots []PlayerokLot

	funpChan := make(chan []Lot)
	playerokChan := make(chan []PlayerokLot)

	go ParsePlayerok(playerokCtx, playerokChan)
	go ParseFunpay(funpayCtx, funpChan)

	for i := 0; i < 2; i++ {
		select {
		case funpayLots = <-funpChan:
			fmt.Println(funpayLots, "Парс фанпея прошел")
		case playerokLots = <-playerokChan:
			fmt.Println(playerokLots, "Парс плеерка прошел")
		}
		fmt.Printf("%d/2", i+1)
	}
}
