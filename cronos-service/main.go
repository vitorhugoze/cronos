package main

import (
	"fmt"
	"main/internals/backup"
)

func main() {

	/* fmt.Println(time.Now().In(time.UTC))

	t, _ := time.Parse("2006-01-02T15:04", "2024-03-01T00:03")

	timer := time.NewTimer(t.Sub(time.Now().In(time.UTC))).C

	<-timer

	fmt.Println("teste--------------------------------------------------------") */

	err := backup.ScheaduleBackups()
	fmt.Println(err)
}
