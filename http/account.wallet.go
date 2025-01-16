package http

import (
	accPages "beltech/frontend/pages/account"
	"beltech/types"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
	"golang.org/x/exp/rand"
)

func (a *App) walletPageHandler(w http.ResponseWriter, r *http.Request) {

	accountEmail := r.Context().Value(contextKeyEmail).(string)
	user, err := a.getAccountFromCache(context.Background(), accountEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}

	if err := accPages.WalletPage(account(user), coinsToDisplay()).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) walletMainTagPartial(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.WalletMainOpeningTag().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) walletBalance(w http.ResponseWriter, r *http.Request) {

	accountEmail := r.Context().Value(contextKeyEmail).(string)
	user, err := a.getAccountFromCache(context.Background(), accountEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.WalletBalance(account(user).CurrBalance).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) walletOptionsPartial(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.WalletOptionsPartial().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) walletcoinData(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.WalletCoinData(coinsToDisplay()).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) walletswiper(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.WalletSwiper().Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) walletLivecoinData(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.WalletLiveCoinData(coinsToDisplay()).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func coinsToDisplay() []types.TokenData {

	arr := make([]types.TokenData, 0, 7)
	arr = append(arr, liveCryptoData.bitcoin, liveCryptoData.solana)

	remaining := []types.TokenData{liveCryptoData.raydium, liveCryptoData.ai16z,
		liveCryptoData.bonk1, liveCryptoData.dogwifhat, liveCryptoData.usdCoin,
		liveCryptoData.jupiterAg, liveCryptoData.pudgyPenguins}

	rand.Shuffle(len(remaining), func(i, j int) {
		remaining[i], remaining[j] = remaining[j], remaining[i]
	})

	arr = append(arr, remaining[:4]...)

	return arr
}

var coins = []string{"bitcoin", "ethereum", "solana", "usd-coin", "bonk1",
	"pudgy-penguins", "raydium", "dogwifhat", "jupiter-ag", "pyth-network", "ai16z"}

type tokenData struct {
	bitcoin, ethereum, solana, usdCoin, bonk1, pudgyPenguins, raydium, dogwifhat, jupiterAg, pythNetwork, ai16z types.TokenData
	rwLock                                                                                                      sync.RWMutex
}

var liveCryptoData tokenData

func (a *App) ScrapeCryptoData() context.CancelFunc {

	getCoinData()

	context, cancel := context.WithCancel(context.Background())
	ticker := time.NewTicker(30 * time.Second)
	go foreverRunning(ticker, context)

	return cancel
}

func foreverRunning(ticker *time.Ticker, closing context.Context) {
loop:
	for {
		select {
		case <-closing.Done():
			ticker.Stop()
			break loop

		case <-ticker.C:
			// var wg sync.WaitGroup
			// wg.Add(1)

			// go func(wg *sync.WaitGroup) {
			// 	defer wg.Done()

			// 	if err := getCoinData(); err != nil {
			// 		log.Println(err)
			// 	}
			// }(&wg)

			// wg.Wait()

			method2("https://coinmarketcap.com/?page=1")

			// Drain extra ticks if tasks took too long
			for len(ticker.C) > 0 {
				<-ticker.C
			}
		}
	}
}

func method2(url string) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var htmlContent string

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(2*time.Second),
		chromedp.KeyEvent(kb.End),
		chromedp.Sleep(2*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			if err := chromedp.OuterHTML("html", &htmlContent).Do(ctx); err != nil {
				fmt.Println("failed to get the full HTML content: ", err)
			}

			if err := scraping2(strings.NewReader(htmlContent)); err != nil {
				fmt.Println("failed to scrape the page: ", err)
			}
			return nil
		}),
	)

	if err != nil {
		fmt.Printf("chromedp error: %v\n\n", err)
	} else {
		fmt.Println("chromedp success")
	}
}

func scraping2(body io.Reader) error {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return fmt.Errorf("failed to parse the HTML document: err %v", err)
	}

	var rtData types.TokenData

	for _, v := range coins {

		identifier := fmt.Sprintf("a[href='/currencies/%s/']", v)

		hold := doc.Find(identifier).Closest("td")
		hold2 := hold.NextAll().Eq(3)

		rtData.TokenSymbol = doc.Find(identifier).Children().First().Find("p").Eq(1).Text()
		rtData.TokenPrice = hold.Next().Text()
		rtData.PercentUpOrDown = hold2.Text()
		rtData.TokenName = doc.Find(identifier).Children().First().Find("p").First().Text()
		rtData.TokenLogo, _ = doc.Find(identifier).Children().First().Find("img").Attr("src")
		rtData.ChartUrl, _ = hold.NextAll().Last().Find("img").Attr("src")
		rtData.UpOrDown, _ = hold2.Find("span>span").Attr("class")

		// fmt.Printf("%v : %+v\n\n\n", v, rtData)
		writeResult2(rtData, v)
	}
	return nil
}

func writeResult2(data types.TokenData, v string) {
	switch v {
	case "bitcoin":
		liveCryptoData.bitcoin = data
	case "ethereum":
		liveCryptoData.ethereum = data
	case "solana":
		liveCryptoData.solana = data
	case "usd-coin":
		liveCryptoData.usdCoin = data
	case "bonk1":
		liveCryptoData.bonk1 = data
	case "pudgy-penguins":
		liveCryptoData.pudgyPenguins = data
	case "raydium":
		liveCryptoData.raydium = data
	case "dogwifhat":
		liveCryptoData.dogwifhat = data
	case "jupiter-ag":
		liveCryptoData.jupiterAg = data
	case "pyth-network":
		liveCryptoData.pythNetwork = data
	case "ai16z":
		liveCryptoData.ai16z = data
	}
}

func getCoinData() error {

	res, err := http.Get("https://coinmarketcap.com/")
	if err != nil {
		return fmt.Errorf("failed to connect to the target page: err %v", err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return fmt.Errorf("HTTP Error %d: %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return fmt.Errorf("failed to parse the HTML document: err %v", err)
	}

	scraping(doc)

	return nil
}

func scraping(doc *goquery.Document) {
	var (
		rtData types.TokenData
		exist  bool
	)

	for _, v := range []string{"bitcoin", "ethereum", "solana", "usd-coin", "bonk1",
		"pudgy-penguins", "raydium", "dogwifhat", "jupiter-ag", "pyth-network", "ai16z"} {

		identifier := fmt.Sprintf("a[href='/currencies/%s/']", v)

		hold := doc.Find(identifier).Closest("td")
		hold2 := hold.NextAll().Eq(3)

		rtData.TokenPrice = hold.Next().Text()
		rtData.PercentUpOrDown = hold2.Text()
		rtData.TokenName = doc.Find(identifier).Children().First().Find("p").First().Text()
		rtData.TokenSymbol = doc.Find(identifier).Children().First().Find("p").Eq(1).Text()

		if rtData.TokenLogo, exist = doc.Find(identifier).Children().First().Find("img").Attr("src"); !exist {
			rtData.TokenLogo = "not-found"
			continue
		}
		if rtData.ChartUrl, exist = hold.NextAll().Last().Find("img").Attr("src"); !exist {
			rtData.ChartUrl = "not-found"
			continue
		}
		if rtData.UpOrDown, exist = hold2.Find("span>span").Attr("class"); !exist {
			rtData.UpOrDown = "not-found"
			continue
		}

		// fmt.Printf("%+v\n", rtData)
		liveCryptoData.writeResult(rtData)
	}
}

func (td *tokenData) writeResult(data types.TokenData) {
	td.rwLock.Lock()
	defer td.rwLock.Unlock()

	switch data.TokenName {
	case "bitcoin":
		liveCryptoData.bitcoin.TokenName = data.TokenName
		liveCryptoData.bitcoin.TokenPrice = data.TokenPrice
		liveCryptoData.bitcoin.UpOrDown = data.UpOrDown
		liveCryptoData.bitcoin.PercentUpOrDown = data.PercentUpOrDown
		liveCryptoData.bitcoin.ChartUrl = data.ChartUrl
		liveCryptoData.bitcoin.TokenSymbol = "BTC"
		liveCryptoData.bitcoin.TokenLogo = "/public/assets/images/bitcoin.png"

	case "ethereum":
		liveCryptoData.ethereum.TokenName = data.TokenName
		liveCryptoData.ethereum.TokenPrice = data.TokenPrice
		liveCryptoData.ethereum.UpOrDown = data.UpOrDown
		liveCryptoData.ethereum.PercentUpOrDown = data.PercentUpOrDown
		liveCryptoData.ethereum.ChartUrl = data.ChartUrl
		liveCryptoData.ethereum.TokenSymbol = "ETH"
		liveCryptoData.ethereum.TokenLogo = "/public/assets/images/ethereum.png"

	case "solana":
		liveCryptoData.solana.TokenName = data.TokenName
		liveCryptoData.solana.TokenPrice = data.TokenPrice
		liveCryptoData.solana.UpOrDown = data.UpOrDown
		liveCryptoData.solana.PercentUpOrDown = data.PercentUpOrDown
		liveCryptoData.solana.ChartUrl = data.ChartUrl
		liveCryptoData.solana.TokenSymbol = "SOL"
		liveCryptoData.solana.TokenLogo = "/public/assets/images/solana.png"

	case "usd-coin":
		liveCryptoData.usdCoin.TokenName = "Circle USDC"
		liveCryptoData.usdCoin.TokenPrice = data.TokenPrice
		liveCryptoData.usdCoin.UpOrDown = data.UpOrDown
		liveCryptoData.usdCoin.PercentUpOrDown = data.PercentUpOrDown
		liveCryptoData.usdCoin.ChartUrl = data.ChartUrl
		liveCryptoData.usdCoin.TokenSymbol = "USDC"
		liveCryptoData.usdCoin.TokenLogo = "/public/assets/images/usdc.png"
	}
}
