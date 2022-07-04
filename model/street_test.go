package model

import (
	"reflect"
	"testing"
)

func TestStreet_GetAdministrativeDivisions(t *testing.T) {
	type fields struct {
		Area *Area
		Base Base
	}
	tests := []struct {
		name   string
		fields fields
		want   []*AdministrativeDivision
	}{
		{
			name: "劝业场街道",
			fields: fields{
				Base: Base{Name: "", URL: "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2021/12/01/01/120101001.html"},
			},
			want: []*AdministrativeDivision{
				{
					StatisticalAreaCode:             "120101001001",
					UrbanAndRuralClassificationCode: "111",
					Name:                            "花园路社区居委会",
				},
				{
					StatisticalAreaCode:             "120101001002",
					UrbanAndRuralClassificationCode: "111",
					Name:                            "滨西社区居委会",
				},
				{
					StatisticalAreaCode:             "120101001004",
					UrbanAndRuralClassificationCode: "111",
					Name:                            "兆丰路社区居委会",
				},
				{
					StatisticalAreaCode:             "120101001005",
					UrbanAndRuralClassificationCode: "111",
					Name:                            "林泉社区居委会",
				},
				{
					StatisticalAreaCode:             "120101001006",
					UrbanAndRuralClassificationCode: "111",
					Name:                            "新疆路社区居委会",
				},
				{
					StatisticalAreaCode:             "120101001007",
					UrbanAndRuralClassificationCode: "111",
					Name:                            "南京路社区居委会",
				},
				{
					StatisticalAreaCode:             "120101001008",
					UrbanAndRuralClassificationCode: "111",
					Name:                            "宁夏路社区居委会",
				},
				{
					StatisticalAreaCode:             "120101001009",
					UrbanAndRuralClassificationCode: "111",
					Name:                            "福明社区居委会",
				},
				{
					StatisticalAreaCode:             "120101001010",
					UrbanAndRuralClassificationCode: "111",
					Name:                            "百货大楼社区居委会",
				},
				{
					StatisticalAreaCode:             "120101001012",
					UrbanAndRuralClassificationCode: "111",
					Name:                            "新津社区居委会",
				},
				{
					StatisticalAreaCode:             "120101001013",
					UrbanAndRuralClassificationCode: "111",
					Name:                            "蒙古路社区居委会",
				},
				{
					StatisticalAreaCode:             "120101001014",
					UrbanAndRuralClassificationCode: "111",
					Name:                            "静园社区居委会",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Street{
				Base: tt.fields.Base,
			}
			got, err := v.GetAdministrativeDivisions()
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCities() got = %v, want %v", got, tt.want)
			}
		})
	}
}
