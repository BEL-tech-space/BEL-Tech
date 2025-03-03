package frontend

import (
	"fmt"
	"math"
	"strings"

	"github.com/google/uuid"

	"time"
)

type WasteSorted struct {
	AccountEmail                  string
	Glass, Metal, Plastics, Paper []WasteTypelog
}

type WasteTypelog struct {
	WasteID           uuid.UUID
	AgentID           uuid.UUID
	Claimed           bool
	WasteWeighed      float32
	WastePointsEarned float32
	LocationAddr      interface{}
	TxnName           string
	CreatedAt         time.Time
}

type TokenData struct {
	TokenName       string `json:"token_name"`
	TokenPrice      string `json:"token_price"`
	PercentUpOrDown string `json:"percent_up_or_down"`
	UpOrDown        string `json:"up_or_down"`
	ChartUrl        string `json:"chart_url"`
	TokenSymbol     string `json:"token_symbol"`
	TokenLogo       string `json:"token_logo"`
}

func chartUrlGreen(url string) string {
	return fmt.Sprintf(`
	<div
		style="
			background-color: #2FBE85;
			width: 7.5rem;
			height: 3rem;
			mask: url('%s') no-repeat center / contain;
			-webkit-mask: url('%s') no-repeat center / contain;"
	></div>`, url, url)
}

func chartUrlRed(url string) string {
	return fmt.Sprintf(`
	<div
		style="
			background-color: #F6455D;
			width: 7.5rem;
			height: 3rem;
			mask: url('%s') no-repeat center / contain;
			-webkit-mask: url('%s') no-repeat center / contain;"
	></div>`, url, url)
}

func formatTotalWasteWeighed(f float64) string {
	switch {
	case f < 10_000:
		return fmt.Sprintf("%.1f", f)
	case f < 100_000:
		return fmt.Sprintf("%.1fk", math.Round(f/1_000*10)/10)
	case f < 1_000_000:
		return fmt.Sprintf("%.1fk", math.Round(f/1_000*10)/10)
	default:
		return fmt.Sprintf("%.1fM", math.Round(f/1_000_000*10)/10)
	}
}

func formatFloatWithComma[T float32 | float64](f T, dp int) string {
	if dp < 0 {
		dp = 1
	}

	factor := math.Pow(10, float64(dp))
	ff := math.Round(float64(f)*factor) / factor

	formatStr := fmt.Sprintf("%%.%df", dp)
	formatted := fmt.Sprintf(formatStr, ff)

	parts := strings.Split(formatted, ".")

	wholeWithCommas := addCommasToNumber(parts[0])

	if dp > 0 {
		return wholeWithCommas + "." + parts[1]
	}
	return wholeWithCommas
}

// addCommasToNumber inserts commas into a whole number string
func addCommasToNumber(num string) string {
	n := len(num)
	if n <= 3 {
		return num
	}

	var result []string
	start := n % 3
	if start > 0 {
		result = append(result, num[:start])
	}

	for i := start; i < n; i += 3 {
		if i > 0 {
			result = append(result, ",")
		}
		result = append(result, num[i:i+3])
	}

	return strings.Join(result, "")
}
