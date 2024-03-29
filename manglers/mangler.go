package manglers

import (
	"os"
	"text_mangler/manglers/insurance_companies_to_post_bodies"
	"text_mangler/manglers/parse_location_csv"
)

type Mangler func(*os.File) []byte

var Registry = map[string]Mangler{
	"p":                  parse_location_csv.Mangle,
	"parse_location_csv": parse_location_csv.Mangle,
	"i":                  insurance_companies_to_post_bodies.Mangle,
}
