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
    d := statistics.Array{
		Data: []float64{5, 4, 3, 2, 1},
    }
	d.Sort(false) // Sorting array
	fmt.Println(d.Data)
	
	// Find median
	fmt.Println(d.Median())
}
```
Output
```
1 2 3 4 5
3
```

# Run example

I have a make example file go in folder running.

To run the program
```bash
go run running/main.go
```