package main

import (
	"fmt"
	"tools/uuid"
)

func main() {
	uuid := uuid.GetUUID()
	fmt.Println(uuid)
}
