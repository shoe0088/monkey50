package printer

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
	"monkey50/market"
)

// ImportCSV はCSVファイルを読み込み、Market型のリストを返却します
func ImportCSV(filename string) ([]market.Market, error) {
	// CSVファイルを開く
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("ファイルを開く際にエラーが発生しました: %v", err)
	}
	defer file.Close()

	// CSVリーダーを作成
	reader := csv.NewReader(file)

	// ヘッダーをスキップ
	_, err = reader.Read() // 1行目がヘッダーの場合はこれで読み飛ばす
	if err != nil {
		return nil, fmt.Errorf("CSVの読み込みに失敗しました: %v", err)
	}

	// Market型のリストを格納するスライス
	var markets []market.Market

	// CSVデータを1行ずつ読み込む
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, fmt.Errorf("CSVファイルの読み込みに失敗しました: %v", err)
		}

		// 日付を time.Time に変換
		date, err := time.Parse("2006-01-02", record[0]) // フォーマットに応じて変更
		if err != nil {
			return nil, fmt.Errorf("日付の変換に失敗しました: %v", err)
		}

		// 各フィールドをMarket型に変換
		open, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return nil, fmt.Errorf("Open値の変換に失敗しました: %v", err)
		}
		closePrice, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return nil, fmt.Errorf("Close値の変換に失敗しました: %v", err)
		}
		high, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			return nil, fmt.Errorf("High値の変換に失敗しました: %v", err)
		}
		low, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			return nil, fmt.Errorf("Low値の変換に失敗しました: %v", err)
		}
		rsi, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			return nil, fmt.Errorf("RSI値の変換に失敗しました: %v", err)
		}

		// Market構造体にデータをマッピング
		marketData := market.Market{
			Date:   date,  // time.Time型に変換された日付
			Open:   open,
			Close:  closePrice,
			High:   high,
			Low:    low,
			RSI:    rsi,
		}

		// Marketリストに追加
		markets = append(markets, marketData)
	}

	// Marketのリストを返却
	return markets, nil
}