package model

var AdministrativeDivisionTag = []string{"tr", "class", "villagetr"}

type Street struct {
	Area *Area
	Base
}

func (s *Street) GetAdministrativeDivisions() ([]*AdministrativeDivision, error) {
	root, err := s.GetSoupDocument()
	if err != nil {
		return nil, err
	}
	administrativeDivisionNodes := root.FindAll(AdministrativeDivisionTag...)
	administrativeDivisions := make([]*AdministrativeDivision, 0, len(administrativeDivisionNodes))
	for _, ad := range administrativeDivisionNodes {
		attrs := ad.FindAll("td")
		if len(attrs) != 3 {
			continue
		}
		var administrativeDivision AdministrativeDivision
		administrativeDivision.Street = s
		for idx, attr := range attrs {
			switch idx {
			case 0:
				administrativeDivision.StatisticalAreaCode = attr.Text()
			case 1:
				administrativeDivision.UrbanAndRuralClassificationCode = attr.Text()
			case 2:
				administrativeDivision.Name = attr.Text()
			}
		}
		administrativeDivisions = append(administrativeDivisions, &administrativeDivision)
	}

	return administrativeDivisions, nil
}
