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

	fmt.Println(t.V1_SellsSupplierReport_ExcelExport_Success())

}

```
### go.mod
```
module billz-2/topics-test

go 1.19

replace github.com/billz-2/topics => ../topics

require github.com/billz-2/topics v0.0.0
```
