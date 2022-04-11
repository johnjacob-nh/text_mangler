# text_mangler

## How to Use
A `Mangler` is a golang text transformation function with the signature:
```go
func transform(*File) []byte
```

Create a mangler and add it to the registry:
```go
package manglers

//...

var Registry = map[string]Mangler{
	"p": parse_location_csv.Mangle,
	"parse_location_csv": parse_location_csv.Mangle,
}
```
Pass the key as an argument to the `-mangler=` flag:
```bash
# in a pipe
cat ~/example.csv | ./text_mangler -mangler=p > ~/delete_this.json
```
```bash
# with file flags
./text_mangler -mangler=p -infile=../Desktop/texas.csv -outfile=../Desktop/delete_this.json
```