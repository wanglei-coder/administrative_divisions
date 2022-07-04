package main

import (
	"administrative_divisions/model"
	"administrative_divisions/pkg/request"
	"encoding/csv"
	"go.uber.org/ratelimit"
	"go.uber.org/zap"
	"os"
	"time"
)

var Logger *zap.SugaredLogger

func init() {
	getLogger()
}

func getLogger() {
	lg, _ := zap.NewDevelopment()
	Logger = lg.Sugar()
}
func main() {
	f, err := os.Create("StatisticalAreaCode.csv")
	if err != nil {
		Logger.Panic(err)
	}
	// 写入UTF-8 BOM
	f.WriteString("\xEF\xBB\xBF")
	//创建一个新的写入文件流
	writer := csv.NewWriter(f)
	writer.Write(model.RowName)

	httpClient := request.NewCustomHTTPClient(request.WithLimiter(ratelimit.New(60,
		ratelimit.Per(time.Minute))))
	//httpClient = request.DefaultCustomHTTPClient
	url := "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2021/index.html"
	ps, err := model.GetProvinces(url, httpClient)
	if err != nil {
		Logger.Fatal(err)
	}

	for _, p := range ps {
		cities, _ := p.GetCities()
		for _, city := range cities {
			areas, _ := city.GetArea()
			for _, area := range areas {
				streets, _ := area.GetStreets()
				for _, street := range streets {
					ads, _ := street.GetAdministrativeDivisions()
					for _, ad := range ads {
						writer.Write(ad.Row())
						Logger.Info(ad.Row())
					}
				}
			}
		}
	}
}
