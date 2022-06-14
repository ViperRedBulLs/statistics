# Statistics

Go Language module for statistical

## Install
```
go get -u github.com/ViperRedBulLs/statistics
```

## Use
```go
package main

import (
    "fmt"
    
    "github.com/ViperRedBulLs/statistics"
)

func main() {
    data := []float64{5, 4, 3, 2, 1}
    fmt.Println(statistics.Sort(data, false))
}
```
Output
```
1 2 3 4 5
```