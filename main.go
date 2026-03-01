package main

import (
	"fmt"
	"unsafe"

	"vhost-go/types"
)

func main() {
	fmt.Println(unsafe.Sizeof(types.InflightDesc{}))
}
