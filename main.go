package main

import (
	"thaiid/thaiid"
)

func main() {
	print(thaiid.IsValid("1234567890121"))
	print(thaiid.IsValid("1-2345-67890-12-1"))
}
