Ptor.go - Path to Regexp
=====

Path to Regexp converts paths to regexps

##Example
```
package main

import (
	"github.com/nathanfaucett/ptor"
	"fmt"
)

func main() {
	// rails like syntax, but with optional regexp matching defaults to [a-zA-Z0-9-_]
	regex, params := ptor.PathToRegexp("/path/:param1/other/:param2[0-9](.:format)", true /* case sensitive */, true /* read whole path to end of string */)
	fmt.Println(regex, params)
}
```