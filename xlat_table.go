package oracletype_to_mongo

// This table defines the mapping between Oracle types and their corresponding Mongo types
var xlatTable = []xlat_t{
	{"CHAR", num_t{scale: lowhigh_t{0, 0}, precision: lowhigh_t{0, 0}}, "string"},       // CHAR to string
	{"VARCHAR", num_t{scale: lowhigh_t{0, 0}, precision: lowhigh_t{0, 0}}, "string"},    // VARCHAR to string
	{"VARCHAR2", num_t{scale: lowhigh_t{0, 0}, precision: lowhigh_t{0, 0}}, "string"},   // VARCHAR2 to string
	{"NCHAR", num_t{scale: lowhigh_t{0, 0}, precision: lowhigh_t{0, 0}}, "string"},      // NCHAR to string
	{"NVARCHAR", num_t{scale: lowhigh_t{0, 0}, precision: lowhigh_t{0, 0}}, "string"},   // NVARCHAR to string
	{"NVARCHAR2", num_t{scale: lowhigh_t{0, 0}, precision: lowhigh_t{0, 0}}, "string"},  // NVARCHAR2 to string
	{"CLOB", num_t{scale: lowhigh_t{0, 0}, precision: lowhigh_t{0, 0}}, "string"},       // CLOB to string
	{"NCLOB", num_t{scale: lowhigh_t{0, 0}, precision: lowhigh_t{0, 0}}, "string"},      // NCLOB to string
	{"BLOB", num_t{scale: lowhigh_t{0, 0}, precision: lowhigh_t{0, 0}}, "binary data"},  // BLOB to binary data
	{"BOOL", num_t{scale: lowhigh_t{0, 0}, precision: lowhigh_t{0, 0}}, "bool"},         // BOOL to bool
	{"NUMBER", num_t{scale: lowhigh_t{1, 1}, precision: lowhigh_t{0, 0}}, "bool"},       // NUMBER(1) to bool
	{"NUMBER", num_t{scale: lowhigh_t{2, 9}, precision: lowhigh_t{0, 0}}, "int"},        // NUMBER(2-9,0) to int
	{"NUMBER", num_t{scale: lowhigh_t{10, 38}, precision: lowhigh_t{0, 0}}, "long"},     // NUMBER(10-38,0) to long
	{"NUMBER", num_t{scale: lowhigh_t{0, 38}, precision: lowhigh_t{10, 38}}, "decimal"}, // NUMBER(0-38,10-38) to decimal
	{"NUMBER", num_t{scale: lowhigh_t{0, 38}, precision: lowhigh_t{1, 9}}, "double"},    // NUMBER(0-38,1-9) to double
	{"NUMBER", num_t{scale: lowhigh_t{0, 0}, precision: lowhigh_t{0, 0}}, "int"},        // NUMBER to int
	{"FLOAT", num_t{scale: lowhigh_t{0, 0}, precision: lowhigh_t{0, 0}}, "double"},      // FLOAT to double
	{"DATE", num_t{scale: lowhigh_t{0, 0}, precision: lowhigh_t{0, 0}}, "date"},         // DATE to date
	{"TIMESTAMP(4)", num_t{scale: lowhigh_t{0, 0}, precision: lowhigh_t{0, 0}}, "date"}, // TIMESTAMP(4) to date
	{"INT64", num_t{scale: lowhigh_t{0, 0}, precision: lowhigh_t{0, 0}}, "long"},        // INT64 to long
}
