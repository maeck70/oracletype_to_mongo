package oracletype_to_mongo

import (
	"fmt"
)

// Exported function to call to translate Oracle types to Mongo types
// Example oracle_type = "NUMBER[19,5]" returns "decimal"
func Xlat(oracle_type string) (mongo_type string, err error) {
	if ok := xlatMap[oracle_type]; ok.ot != "" {
		return xlatMap[oracle_type].mt, nil
	}
	return "unknown", fmt.Errorf("unknown Oracle type: %s", oracle_type)
}

var xlatMap = make(map[string]xlat_t)

// Populate the xlat map on startup
// This explodes the options from the xlatTable across its precision and scale ranges
func init() {
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

// Explode the precision and scale ranges
// and return a list of strings that represent the options
func explode(r num_t) []string {
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
