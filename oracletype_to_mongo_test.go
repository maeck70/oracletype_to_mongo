package oracletype_to_mongo

import (
	"testing"
)

func TestXlat(t *testing.T) {
	// Define test cases
	testCases := []struct {
		oracleType    string
		MongoExpected string
	}{
		// Note, scale can never be smaller than precision
		// Eg.
		//   NUMBER[19,5] is a long (123456.12345),
		//   NUMBER[2,2] is a double (0.12),
		//   NUMBER[0,5] is invalid because scale is smaller than precision
		{"NUMBER[19,5]", "double"},
		{"VARCHAR2", "string"},
		{"NUMBER[1,0]", "bool"},
		{"NUMBER[2,0]", "int"},
		{"NUMBER[9,0]", "int"},
		{"NUMBER[19,0]", "long"},
		{"NUMBER[2,2]", "double"},
		{"NUMBER[19,18]", "decimal"},
		{"NUMBER[38,19]", "decimal"},
		{"NUMBER[0,0]", "int"},
		{"FLOAT", "double"},
		{"NUMBER[5,10]", "unknown"},
		{"DATE", "date"},
		{"TIMESTAMP(4)", "date"},
		{"INT64", "long"},
		{"NUMBER[2,2]", "double"},
		{"NUMBER[24,20]", "decimal"},
	}

	type result_t struct {
		mongoType string
		err       error
	}

	// Run test cases
	for _, tc := range testCases {
		var r result_t
		r.mongoType, r.err = Xlat(tc.oracleType)

		// log.Print(r)

		if r.mongoType != tc.MongoExpected {
			t.Errorf("Oracle2MongoXlat(%s) = (%s, %v), expected %s", tc.oracleType, r.mongoType, r.err, tc.MongoExpected)
		}

	}
}
