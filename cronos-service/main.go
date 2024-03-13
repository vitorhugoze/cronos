package main

import (
	"log"
	"main/internals/backup"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	sigChan := make(chan os.Signal, 2)

	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-sigChan
		os.Exit(1)
	}()

	startBackupTask()
}

func startBackupTask() {

	for {
		<-time.NewTicker(time.Second * 10).C

		go func() {
			err := backup.ScheaduleBackups()

			if err != nil {
				log.Println(err)
			}
		}()

		go func() {
			err := backup.ScheaduleDeletions()

			if err != nil {
				log.Println(err)
			}
		}()
	}

}
