package insurance_companies_to_post_bodies

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
)

const namesPerFile = 500
const filenamePrefix = "insurance_companies"

type insuranceCompany struct {
	Name       string `json:"name"`
	ChcPayorId string `json:"chc_payor_id"`
	Type       string `json:"type"`
}

func Mangle(f *os.File) []byte {
	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	insuranceCompanies := make([]*insuranceCompany, 0)
	for _, row := range data[1:] {
		chcPayorId := row[0]
		name := row[4]
		insuranceCompany := &insuranceCompany{
			Name:       name,
			ChcPayorId: chcPayorId,
			Type:       "private_insurance",
		}
		insuranceCompanies = append(insuranceCompanies, insuranceCompany)
	}

	//create enough "files" to hold all of the results
	nFiles := int(math.Ceil(float64(len(insuranceCompanies)) / float64(namesPerFile)))
	namesObjs := make([][]*insuranceCompany, nFiles)
	for i := 0; i < nFiles; i++ {
		namesObjs[i] = make([]*insuranceCompany, 0)
	}
	//for every company...
	//put the company in the corresponding file
	for i, company := range insuranceCompanies {
		resultObjI := i / namesPerFile
		namesObjs[resultObjI] = append(namesObjs[resultObjI], company)
	}

	for i, resultObj := range namesObjs {
		bytes := marshal(resultObj)
		filename := fmt.Sprintf("%s%d.json", filenamePrefix, i)
		err = os.WriteFile(filename, bytes, 0644)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("wrote file %s", filename)
	}
	bytes := marshal(namesObjs)
	return bytes
}

func uniq(collection []string) []string {
	results := make([]string, 0)
	seen := map[string]bool{}
	for _, str := range collection {
		if _, ok := seen[str]; !ok {
			seen[str] = true
			results = append(results, str)
		}
	}
	return results
}

func marshal(data any) []byte {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	_ = enc.Encode(data)
	return buf.Bytes()
}
