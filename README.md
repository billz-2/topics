# Common topics library

Package contains common topics that can be used between services



This library consists of:

- Exported topics
  
    List of topics splitted to 2 versions
  - V1
  - V2

  Feel free to include new topic-functions in any order. Just keep in mind that other developers will also use it :)
  
- internal helper
  
  Just small function to transform the strings. You are more hen welcome to cover it with units


## Example with standalone app and local package

### main.go
```
package main

import (
	"fmt"

	t "github.com/billz-2/topics"
)

func main() {

	fmt.Println(t.V1_ReportService_SellsSupplierReport_ExcelExport_Success())

}

```
### go.mod
```
module billz-2/topics-test

go 1.19

replace github.com/billz-2/topics => ../topics

require github.com/billz-2/topics v0.0.0
```


# How to create new topics?

Well, pretty easy. 

Just create new function with bloody name :)

## No really, what are the rules?
Each topic has unique name that reflects version, service, entity, action and result:

For example the topic **v1.report_service.sells_supplier_report.excel_export.success**
- **v1** is version
- **report_service** is reporting service
- **sells_supplier_report** is entity of Sales BY Suppliers Report
- **excel_export** is an action that request report in Excel format
- **success** is a result of creating excel file with data

To share the topic between services, create a function in related file v1.go or v2.go

Function for example above will have the name **V1_ReportService_SellsSupplierReport_ExcelExport_Success**

When you call it, the func transform own name name to string by following rules

- underscore "**_**" converted to dot "."
- Each section has to be in SnakeCase and will be converted to *SnakeCase*
