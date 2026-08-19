package main

import (
	"fmt"
)

func main() {
	for i, ch := range "Héllo" {
		if i == 0 {
			continue
		}
		fmt.Println(i, string(ch))
	}
}
