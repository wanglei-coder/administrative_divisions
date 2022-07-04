package model

import (
	"administrative_divisions/pkg/request"
)

var (
	provinceTableTag = []string{"table", "class", "provincetable"}
	ProvinceTag      = []string{"tr", "class", "provincetr"}
	CityTag          = []string{"tr", "class", "citytr"}
)

type Province struct {
	Base
}

func (p *Province) GetCities() (cities []*City, err error) {
	root, err := p.GetSoupDocument()
	if err != nil {
		return nil, err
	}
	cityNodes := root.FindAll(CityTag...)
	for _, cityNode := range cityNodes {
		attrs := cityNode.FindAll("td")
		var city City
		city.Province = p
		city.Client = p.Client
		for idx, attr := range attrs {
			attrA := attr.Find("a")
			switch idx {
			case 0:
				city.URL = BaseURL + attrA.Attrs()["href"]
				city.StatisticalAreaCode = attrA.Text()
			case 1:
				city.Name = attrA.Text()
			}
		}
		cities = append(cities, &city)
	}

	return cities, nil
}

func GetProvinces(url string, client *request.CustomHTTPClient) (provinces []*Province, err error) {
	soupDoc, err := client.GetSoupDocument(url)
	if err != nil {
		return nil, err
	}
	provinceTable := soupDoc.Find(provinceTableTag...)
	provinceTr := provinceTable.FindAll(ProvinceTag...)

	for _, provinceGroup := range provinceTr {
		provinceTds := provinceGroup.FindAll("td")
		for _, provinceTd := range provinceTds {
			var province Province
			province.Client = client
			attrA := provinceTd.Find("a")
			province.Name = attrA.Text()
			province.URL = BaseURL + attrA.Attrs()["href"]
			provinces = append(provinces, &province)
		}
	}
	return provinces, err
}
