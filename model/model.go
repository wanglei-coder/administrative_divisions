package model

// AdministrativeDivision 统计用区划代码	城乡分类代码	名称
type AdministrativeDivision struct {
	Street *Street
	// StatisticalAreaCode 统计用区划代码
	StatisticalAreaCode string

	// UrbanAndRuralClassificationCode 城乡分类代码
	UrbanAndRuralClassificationCode string

	// Name 名称
	Name string
}

func (ad *AdministrativeDivision) Row() []string {
	//row := []string{b.TeamCode, b.LineName, b.StartDate, b.EndDate, b.GuideName, b.IDNumber, b.NumberOfPeople, b.OperatorName,
	//	b.Company, b.CarNumber, b.DriverName, b.DriverPhone, b.DriverNumber, b.TouristSourceCity,
	//}

	row := []string{
		ad.Street.Area.City.Province.Name,
		ad.Street.Area.City.Province.StatisticalAreaCode,
		ad.Street.Area.City.Name,
		ad.Street.Area.City.StatisticalAreaCode,
		ad.Street.Area.Name,
		ad.Street.Area.StatisticalAreaCode,
		ad.Street.URL,
		ad.Street.Name,
		ad.Street.StatisticalAreaCode,
		ad.StatisticalAreaCode,
		ad.UrbanAndRuralClassificationCode,
		ad.Name,
	} //b.Company, b.CarNumber, b.DriverName, b.DriverPhone, b.DriverNumber, b.TouristSourceCity,

	return row
}
