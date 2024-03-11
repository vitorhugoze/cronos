package backup

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"
	"time"
)

var scheaduledFiles []string

func ScheaduleDeletions() error {

	var wg sync.WaitGroup

	cfg, err := readBkpConfig()
	if err != nil {
		return err
	}

	for _, s := range cfg.BkpSchedules {

		if s.DeleteInterval == "" {
			continue
		}

		filepath.Walk(s.DestPath, func(path string, info fs.FileInfo, err error) error {

			if info.IsDir() {
				return nil
			}

			if filepath.Ext(path) == ".zip" && !slices.Contains(scheaduledFiles, path) {

				var d, h, m int

				scheaduledFiles = append(scheaduledFiles, path)

				fileTime, err := time.Parse("2006-01-02T15-04-05", strings.TrimSuffix(filepath.Base(path), ".zip"))
				if err != nil {
					log.Println(err)
					return nil
				}

				_, err = fmt.Sscanf(s.DeleteInterval, "%v-%v-%v", &d, &h, &m)
				if err != nil {
					return err
				}

				delTime := fileTime.Add(time.Duration(int(time.Hour) * 24 * d))
				delTime = delTime.Add(time.Duration(int(time.Hour) * h))
				delTime = delTime.Add(time.Duration(int(time.Minute) * m))
				now, err := time.Parse("2006-01-02T15-04-05", time.Now().Format("2006-01-02T15-04-05"))
				if err != nil {
					return err
				}

				if delTime.Before(now) {

					err = os.Remove(path)

					if err != nil {
						log.Println(err)
						return nil
					}
				} else {

					wg.Add(1)

					go func(p string) {

						<-time.NewTimer(delTime.Sub(now)).C

						err := os.Remove(p)
						if err != nil {
							log.Println(err)
						}

						wg.Done()
					}(path)
				}

			}

			return nil
		})

	}

	wg.Wait()

	return nil
}
