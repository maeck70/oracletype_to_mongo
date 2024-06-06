package main

import (
	"fmt"
)

// Exported function to call to translate Oracle types to Mongo types
func Oracle2MongoXlat(oracle_type string) (string, error) {
	if ok := xlatMap[oracle_type]; ok.ot != "" {
		return xlatMap[oracle_type].mt, nil
	}
	return "unknown", fmt.Errorf("unknown Oracle type: %s", oracle_type)
}

var xlatMap = make(map[string]xlat_t)

func init() {
	// Populate the xlat map on startup
	// This explodes the options from the xlatTable across its precision and scale ranges
	for _, i := range xlatTable {
		switch i.ot {
		case "NUMBER":
			for _, j := range explode(i.otnp) {
				xlatMap[fmt.Sprintf("%s%s", i.ot, j)] = i
			}
		default:
			xlatMap[i.ot] = i
		}
	}
}

func explode(r num_t) []string {
	// Explode the precision and scale ranges
	// and return a list of strings that represent the options
	var a []string
	for i := r.precision.low; i <= r.precision.high; i++ {
		for j := r.scale.low; j <= r.scale.high; j++ {
			if i <= j { // scale can never be bigger than precision
				a = append(a, fmt.Sprintf("[%d,%d]", j, i))
			}
		}
	}
	return a
}

func main() {
	fmt.Println("There is no main to run, just test with go test .")
}
