// contains utility functions for a candidates competence matrix
package candidateutil

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"log"
	"strconv"
)

type Candidate struct {
	Id,Age,Comp_java,Comp_aws,Comp_gcp,Comp_go,Soft_eng,Soft_lead,Soft_innov    int
	Name,Email   string
}

func getCellIntValue (val string,err error) int {
	x,err := strconv.Atoi(val)
	if err != nil {
		log.Fatal(err)
	}
return	x
}

func Get(id int) Candidate {

	var (
		Id,Age,Comp_java,Comp_aws,Comp_gcp,Comp_go,Soft_eng,Soft_lead,Soft_innov       int
		Name,Email  string
	)

	f, err := excelize.OpenFile("candidateutil/candidate.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	Name,err = f.GetCellValue("PROFILE", "B1")
	Email,err = f.GetCellValue("PROFILE", "B2")

	Age  = getCellIntValue(f.GetCellValue("PROFILE", "B3"))
	Comp_java = getCellIntValue(f.GetCellValue("COMPETENCY", "B1"))
	Comp_aws = getCellIntValue(f.GetCellValue("COMPETENCY", "B2"))
	Comp_gcp = getCellIntValue(f.GetCellValue("COMPETENCY", "B3"))
	Comp_go = getCellIntValue(f.GetCellValue("COMPETENCY", "B4"))

	Soft_eng  = getCellIntValue(f.GetCellValue("SOFTSKILLS", "B1"))
	Soft_lead = getCellIntValue(f.GetCellValue("SOFTSKILLS", "B2"))
	Soft_innov = getCellIntValue(f.GetCellValue("SOFTSKILLS", "B3"))


	if err != nil {
		log.Fatal(err)
	}

	details := Candidate{
		Id:   (Id),
		Name: (Name),
		Email: (Email),
		Age:(Age),
		Comp_java:(Comp_java),
		Comp_aws:(Comp_aws),
		Comp_gcp:(Comp_gcp),
		Comp_go:(Comp_go),
		Soft_eng:(Soft_eng),
		Soft_lead:(Soft_lead),
		Soft_innov:(Soft_innov),

	}

	return details
}