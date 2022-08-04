package insurance_company_names_to_post_body

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
)

//the Name is at index 3 in the csv, the 4th field
const nameIndex = 3
const namesPerFile = 500
const filenamePrefix = "payer_names"

type Names struct {
	Names []string `json:"names"`
}

func Mangle(f *os.File) []byte {
	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	names := make([]string, 0)
	for _, row := range data[1:] {
		name := row[nameIndex]
		names = append(names, name)
	}

	names = uniq(names)

	//create enough "files" to hold all of the results
	nNamesObjs := int(math.Ceil(float64(len(names)) / float64(namesPerFile)))
	namesObjs := make([]*Names, nNamesObjs)
	for i := 0; i < nNamesObjs; i++ {
		namesObjs[i] = new(Names)
	}
	//for every Name...
	//put the Name in the corresponding file
	for i, name := range names {
		resultObjI := i / namesPerFile
		namesObjs[resultObjI].Names = append(namesObjs[resultObjI].Names, name)
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
