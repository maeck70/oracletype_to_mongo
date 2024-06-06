package main

// This table defines the mapping between Oracle types and their corresponding Mongo types
var xlatTable = []xlat_t{
	{"CHAR", num_t{scale: lowhigh_t{0, 0}, precision: lowhigh_t{0, 0}}, "string"},
	{"VARCHAR", num_t{scale: lowhigh_t{0, 0}, precision: lowhigh_t{0, 0}}, "string"},
	{"VARCHAR2", num_t{scale: lowhigh_t{0, 0}, precision: lowhigh_t{0, 0}}, "string"},
	{"NCHAR", num_t{scale: lowhigh_t{0, 0}, precision: lowhigh_t{0, 0}}, "string"},
	{"NVARCHAR", num_t{scale: lowhigh_t{0, 0}, precision: lowhigh_t{0, 0}}, "string"},
	{"NVARCHAR2", num_t{scale: lowhigh_t{0, 0}, precision: lowhigh_t{0, 0}}, "string"},
	{"CLOB", num_t{scale: lowhigh_t{0, 0}, precision: lowhigh_t{0, 0}}, "string"},
	{"NCLOB", num_t{scale: lowhigh_t{0, 0}, precision: lowhigh_t{0, 0}}, "string"},
	{"BLOB", num_t{scale: lowhigh_t{0, 0}, precision: lowhigh_t{0, 0}}, "binary data"},
	{"BOOL", num_t{scale: lowhigh_t{0, 0}, precision: lowhigh_t{0, 0}}, "bool"},
	{"NUMBER", num_t{scale: lowhigh_t{1, 1}, precision: lowhigh_t{0, 0}}, "bool"},
	{"NUMBER", num_t{scale: lowhigh_t{2, 38}, precision: lowhigh_t{0, 0}}, "int"},
	{"NUMBER", num_t{scale: lowhigh_t{1, 38}, precision: lowhigh_t{1, 38}}, "decimal"},
	{"NUMBER", num_t{scale: lowhigh_t{0, 0}, precision: lowhigh_t{0, 0}}, "int"},
	{"FLOAT", num_t{scale: lowhigh_t{0, 0}, precision: lowhigh_t{0, 0}}, "float"},
	{"DATE", num_t{scale: lowhigh_t{0, 0}, precision: lowhigh_t{0, 0}}, "date"},
	{"TIMESTAMP(4)", num_t{scale: lowhigh_t{0, 0}, precision: lowhigh_t{0, 0}}, "timestamp"},
	{"INT64", num_t{scale: lowhigh_t{0, 0}, precision: lowhigh_t{0, 0}}, "long"},
}
