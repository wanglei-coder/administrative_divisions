package model

import (
	"path/filepath"
	"strings"
)

var AreaTag = []string{"tr", "class", "countytr"}

type City struct {
	Province *Province
	Base
}

func (c *City) GetArea() (areas []*Area, err error) {
	root, err := c.GetSoupDocument()
	if err != nil {
		return nil, err
	}
	areaNodes := root.FindAll(AreaTag...)
	for _, areaNode := range areaNodes {
		attrs := areaNode.FindAll("td")
		var area Area
		area.City = c
		area.Client = c.Client
	Loop:
		for idx, attr := range attrs {
			attrA := attr.Find("a")
			if attrA.Pointer == nil {
				break
			}
			switch idx {
			case 0:
				m := attrA.Attrs()
				if m == nil {
					break Loop
				}

				h := m["href"]
				if h == "" {
					break Loop
				}
				basePath := filepath.Base(h)
				area.URL = BaseURL + strings.Join([]string{basePath[0:2], h}, "/")
				area.StatisticalAreaCode = attrA.Text()
			case 1:
				area.Name = attrA.Text()
			}
		}
		areas = append(areas, &area)
	}

	return areas, nil
}
