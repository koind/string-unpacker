# String-unpacker

A small library that performs a primitive unpacking of a string containing duplicate characters.


## Installation

Run the following command from you terminal:


 ```bash
 go get github.com/koind/string-unpacker
 ```

## Usage

Package usage example.

```go
package main

import (
	"github.com/koind/string-unpacker"
)

func main() {
	unpacker := string_unpacker.NewUnpacker()
	
	resultStr := unpacker.Unpack("a4bc2d5e")
	println(resultStr) // aaaabccddddde
	
	resultStr = unpacker.Unpack("°5€1")
	println(resultStr) // °°°°°€
}
```

## Available Methods

The following methods are available:

##### koind/string-unpacker

```go
Unpack(text string) string
```

## Tests

Run the following command from you terminal:

```
go test -v .
```