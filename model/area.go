package model

import (
	"path/filepath"
	"strings"
)

var StreetTag = []string{"tr", "class", "towntr"}

type Area struct {
	City *City
	Base
}

func (a *Area) GetStreets() (streets []*Street, err error) {
	root, err := a.GetSoupDocument()
	if err != nil {
		return nil, err
	}
	streetTag := root.FindAll(StreetTag...)
	for _, areaNode := range streetTag {
		attrs := areaNode.FindAll("td")
		var street Street
		street.Area = a
		street.Client = a.Client
	Loop:
		for idx, attr := range attrs {
			attrA := attr.Find("a")
			switch idx {
			case 0:
				h := attrA.Attrs()["href"]
				if h == "" {
					break Loop
				}
				basePath := filepath.Base(h)
				street.URL = BaseURL + strings.Join([]string{basePath[0:2], basePath[2:4], h}, "/")
				street.StatisticalAreaCode = attrA.Text()
			case 1:
				street.Name = attrA.Text()
			}
		}
		streets = append(streets, &street)
	}

	return streets, nil
}
