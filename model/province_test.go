package model

import (
	"administrative_divisions/pkg/request"
	"testing"
)

func TestGetProvinces(t *testing.T) {
	client := request.DefaultCustomHTTPClient
	url := "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2021/index.html"
	ps, err := GetProvinces(url, client)
	if err != nil {
		t.Fatal(err)
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
						t.Log(ad.Name,
							ad.StatisticalAreaCode,
							ad.UrbanAndRuralClassificationCode,
							ad.Street.Name,
							ad.Street.URL,
							ad.Street.Area.Name,
							ad.Street.Area.City.Name,
							ad.Street.Area.City.Province.Name,
						)
					}
				}
			}
		}
	}
}
