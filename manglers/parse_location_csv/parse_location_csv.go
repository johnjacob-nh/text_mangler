package parse_location_csv

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"strings"
)

func Mangle(f *os.File) []byte {
	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	//firstRow := data[0]
	var results []*RowJson
	for _, row := range data[1:] {
		rj := &RowJson{

		}
		for i, f := range row {
			//fmt.Printf("%d, %s, %s\n", i, firstRow[i], f)
			switch i {
			case 0:
				rj.ID = "<FILL_ME_OUT>"
			case 1:
				rj.Type = f
			case 2:
				rj.OrgID = f
			case 3:
				rj.Name = f
			case 4:
				rj.LocationCode = f
			case 5:
				rj.CliaID = f
			case 6:
				rj.Address.Street1 = f
			case 7:
				rj.Address.Street2 = f
			case 8:
				rj.Address.City = f
			case 9:
				rj.Address.State = f
			case 10:
				rj.Address.PostalCode = f
			case 11:
				rj.Address.County = f
			case 12:
				rj.Address.Country = f
			case 13:
				rj.Address.Latitude = f
			case 14:
				rj.Address.Longitude = f
			case 15:
				rj.Address.Timezone = f
			case 16:
				rj.InsuranceRequirement = f
			case 17:
				rj.RequiredWithSSN = f
			case 18:
				rj.Npi = f
			case 19:
				rj.DefaultPcrLabID = f
			case 20:
				rj.DefaultAntigenLabID = f
			case 21:
				rj.Registration = f == "true"
			case 22:
				rj.Visible = f
			case 23:
				//procedure_types
				pt := strings.Split(f, "\n")
				for i, s := range pt {
					pt[i] = strings.Trim(s, ",\" \n\r")
				}
				//fmt.Println(f)
				//fmt.Printf("%+v\n", pt)
				rj.ProcedureTypes = pt
			case 24:
			case 25:
				//test_kit_types
				tkt := strings.Split(f, "\n")
				for i, s := range tkt {
					tkt[i] = strings.Trim(s, ",\" \n\r")
				}
				//fmt.Println(f)
				//fmt.Printf("%+v\n", tkt)
				rj.TestKitTypes = tkt
			case 26:
			case 27:
			case 28:
			case 29:
			case 30:
			case 31:
			case 32:
			case 33:
			case 34:
			case 35:
			case 36:

			}
			//fmt.Printf("%+v\n", rowJson)
		}
		results = append(results, rj)
	}

	j, _ := json.MarshalIndent(results, "", "\t")

	return j
}

type RowJson struct {
	ID           string `json:"id"`
	OrgID        string `json:"org_id"`
	Type         string `json:"type"`
	Name         string `json:"name"`
	LocationCode string `json:"location_code"`
	CliaID       string `json:"clia_id"`
	Address      struct {
		Street1    string `json:"street_1"`
		Street2    string `json:"street_2"`
		City       string `json:"city"`
		State      string `json:"state"`
		PostalCode string `json:"postal_code"`
		County    string `json:"county"`
		Country    string `json:"country"`
		Latitude   string `json:"latitude"`
		Longitude  string `json:"longitude"`
		Timezone   string `json:"timezone"`
	} `json:"address"`
	InsuranceRequirement string   `json:"insurance_requirement"`
	RequiredWithSSN string   `json:"required_with_ssn"`
	DefaultPcrLabID      string   `json:"default_pcr_lab_id"`
	DefaultAntigenLabID  string   `json:"default_antigen_lab_id"`
	Registration         bool     `json:"registration"`
	Visible              string   `json:"visible"`
	ProcedureTypes       []string `json:"procedure_types"`
	TestKitTypes         []string `json:"test_kit_types"`
	Npi                  string   `json:"npi"`
}
