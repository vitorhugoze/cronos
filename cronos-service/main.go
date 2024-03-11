package main

import (
	"log"
	"main/internals/backup"
	"time"
)

func main() {

	for {
		<-time.NewTicker(time.Second * 10).C

		go func() {
			err := backup.ScheaduleBackups()
			if err != nil {
				log.Fatal(err)
			}
		}()

		go func() {
			err := backup.ScheaduleDeletions()
			if err != nil {
				log.Fatal(err)
			}
		}()
	}
}
