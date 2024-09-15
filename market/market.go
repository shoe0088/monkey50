package market

import "time"

type Market struct {
	Date        time.Time    // 日付
	Open         float64    // 始値
	Close      float64 // 終値
	Low  float64  // 安値
	High float64  // 高値
	RSI  float64
}
