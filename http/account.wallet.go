package http

import (
	accountPages "beltech/frontend/pages/account"
	"beltech/types"
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func (a *App) walletPageHandler(w http.ResponseWriter, r *http.Request) {

	userEmail := r.Context().Value(contextKeyEmail).(string)

	user, err := a.userFromCache(context.Background(), userEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accountPages.WalletPage(account(*user), []types.TokenData{liveCryptoData.bitcoin, liveCryptoData.ethereum, liveCryptoData.solana, liveCryptoData.usdCoin}).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) coinDataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accountPages.CoinData([]types.TokenData{liveCryptoData.bitcoin, liveCryptoData.ethereum, liveCryptoData.solana, liveCryptoData.usdCoin}).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

type tokenData struct {
	bitcoin, ethereum, solana, usdCoin types.TokenData
	rwLock                             sync.RWMutex
}

var liveCryptoData tokenData

func (a *App) ScrapeCryptoData() chan struct{} {

	ticker := time.NewTicker(10 * time.Second)
	closing := make(chan struct{})
	go foreverRunning(ticker, closing)

	return closing
}

func foreverRunning(ticker *time.Ticker, closing chan struct{}) {
loop:
	for {
		select {
		case <-closing:
			ticker.Stop()
			close(closing)
			break loop

		case <-ticker.C:
			var wg sync.WaitGroup
			wg.Add(1)

			go func(wg *sync.WaitGroup) {
				defer wg.Done()

				if err := getCoinData(); err != nil {
					log.Println(err)
				}
			}(&wg)

			wg.Wait()

			// Drain extra ticks if tasks took too long
			for len(ticker.C) > 0 {
				<-ticker.C
			}
		}
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

	// parse the HTML document
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

	for _, v := range [4]string{"bitcoin", "ethereum", "solana", "usd-coin"} {

		rtData.TokenName = v
		identifier := fmt.Sprintf("a[href='/currencies/%s/']", v)

		hold := doc.Find(identifier).Closest("td")
		rtData.TokenPrice = hold.Next().Text()
		if rtData.ChartUrl, exist = hold.NextAll().Last().Find("img").Attr("src"); !exist {
			rtData.ChartUrl = "not-found"
		}

		hold2 := hold.NextAll().Eq(3)
		rtData.PercentUpOrDown = hold2.Text()
		if rtData.UpOrDown, exist = hold2.Find("span>span").Attr("class"); !exist {
			rtData.UpOrDown = "not-found"
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
