# Oracle Type to Mongo type converter

This library provides the function to convert an Oracle column type (eg. VARCHAR2) to a MongoDB type (eg. string). 

Simply import the library and call 
` mongoType, err := Xlat(<OracleType>)`

## Supported Oracle types

The supported Oracle types are defied in the xlat_table file.