# Oracle Type to Mongo type converter

This Go package provides a convenient way to translate Oracle data types to their corresponding MongoDB data types. It simplifies the process of mapping data between these two popular database systems.

### Installation

To use this package in your Go project, run the following command:

```bash
go get -u github.com/maeck70/oracletype_to_mongo
```

### Usage

Import the `oracletype_to_mongo` package and utilize the `Xlat` function:

```go
package main

import (
    "fmt"

    "github.com/maeck70/oracletype_to_mongo"
)

func main() {
    oracleType := "NUMBER[19,5]"
    mongoType, err := oracletype_to_mongo.Xlat(oracleType)

    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Printf("Oracle type: %s -> MongoDB type: %s\n", oracleType, mongoType)
}
```

This code snippet demonstrates how to:

1. Import the package.
2. Define an Oracle data type string.
3. Call the `Xlat` function to translate the type.
4. Handle potential errors (if the Oracle type is unknown).
5. Print the translated MongoDB type.

### Functions

* **`Xlat(oracle_type string) (mongo_type string, err error)`**
   - This is the primary function for translating Oracle data types to MongoDB types.
   - It takes an `oracle_type` string as input and returns a `mongo_type` string and an optional error.
   - The function searches an internal mapping (`xlatMap`) for the provided Oracle type.
   - If a match is found, the corresponding MongoDB type is returned.
   - If the Oracle type is not found, an error is returned with a message indicating the unknown type.

### Internal Mapping

This package utilizes pre-defined mappings between Oracle data types and their corresponding MongoDB data types. These mappings are stored within the `xlatTable` and provide a flexible way to handle various Oracle data type formats. Let's break down the table structure:

- **`xlat_t`**: This is a custom struct that defines two properties:
    - `ot`: The Oracle data type string (e.g., "CHAR", "NUMBER").
    - `mt`: The corresponding MongoDB data type string (e.g., "string", "int").
    - `num_t` (optional): For numeric types (like "NUMBER"), this struct defines the scale and precision ranges using the `lowhigh_t` struct.
- **`lowhigh_t`**: This struct represents a range with low and high values. It's used to capture the potential scale and precision variations within an Oracle numeric data type.

The `xlatTable` contains entries for various Oracle data types, mapping them to their appropriate MongoDB representations. Here are some key points to note:

- **String Types**: Character-based Oracle types like "CHAR", "VARCHAR", and "NVARCHAR" are all mapped to the "string" type in MongoDB.
- **Number Types**: The mapping for "NUMBER" depends on the scale and precision ranges. Numbers with a scale of 0 and precision of 1 are mapped to "bool" (assuming they represent boolean values). Integers with varying precisions are mapped to "int" or "long" depending on the range. Numbers with a non-zero scale are mapped to "decimal" for high precision or "double" for lower precision.
- **Other Types**: BLOBs are mapped to "binary data", and date/timestamp types are mapped to "date" in MongoDB.

**Customization:**

This table serves as a foundation for mapping common Oracle data types. You can extend or modify this table to accommodate additional data types or adjust the mappings based on your specific requirements.

**Note:** The details of this internal mapping may change in future versions of the package.

### Contributing

Feel free to submit pull requests to improve this package. We welcome contributions for bug fixes, enhancements, and additional data type mappings.

### License

This package is licensed under the [MIT License](https://opensource.org/licenses/MIT).
