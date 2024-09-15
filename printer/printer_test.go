package printer

import (
	"testing"
	"time"

	"monkey50/market"
)

func TestImportCSV(t *testing.T) {
	// テスト用のCSVファイルパス
	filePath := "test.csv"

	// Marketリストを読み込む関数をモックする
	markets, err := ImportCSV(filePath)
	if err != nil {
		t.Fatalf("CSV読み込み中にエラーが発生しました: %v", err)
	}

	// 結果を検証する
	expectedMarkets := []market.Market{
		{
			Date:   parseDate("2023-09-12"),
			Open:   100.5,
			Close:  105.3,
			High:   106.0,
			Low:    99.8,
			RSI:    30.5,
		},
		{
			Date:   parseDate("2023-09-13"),
			Open:   105.3,
			Close:  102.0,
			High:   108.5,
			Low:    101.0,
			RSI:    40.2,
		},
	}

	// 結果が期待通りかをチェック
	for i, marketData := range markets {
		expected := expectedMarkets[i]

		if marketData.Date != expected.Date {
			t.Errorf("期待した日付 %v, しかし得た日付 %v", expected.Date, marketData.Date)
		}
		if marketData.Open != expected.Open {
			t.Errorf("期待したOpen値 %v, しかし得たOpen値 %v", expected.Open, marketData.Open)
		}
		if marketData.Close != expected.Close {
			t.Errorf("期待したClose値 %v, しかし得たClose値 %v", expected.Close, marketData.Close)
		}
		if marketData.High != expected.High {
			t.Errorf("期待したHigh値 %v, しかし得たHigh値 %v", expected.High, marketData.High)
		}
		if marketData.Low != expected.Low {
			t.Errorf("期待したLow値 %v, しかし得たLow値 %v", expected.Low, marketData.Low)
		}
		if marketData.RSI != expected.RSI {
			t.Errorf("期待したRSI値 %v, しかし得たRSI値 %v", expected.RSI, marketData.RSI)
		}
	}
}

// テストのために日付を解析するヘルパー関数
func parseDate(dateStr string) time.Time {
	date, _ := time.Parse("2006-01-02", dateStr)
	return date
}