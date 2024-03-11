package backup

import (
	"fmt"
	ziputils "main/pkg/zip"
	"main/types"
	"os"
	"slices"
	"sync"
	"time"

	"github.com/spf13/viper"
)

var scheaduledBkps []types.Schedule

func ScheaduleBackups() error {

	var wg sync.WaitGroup

	cfg, err := readBkpConfig()
	if err != nil {
		return err
	}

	for _, s := range cfg.BkpSchedules {

		if !slices.Contains(scheaduledBkps, s) {

			wg.Add(1)
			go func(s types.Schedule) {
				backupTask(s)
				wg.Done()
			}(s)

			scheaduledBkps = append(scheaduledBkps, s)
		}
	}

	wg.Wait()
	return nil
}

func readBkpConfig() (*types.BkpConf, error) {

	var config types.BkpConf

	configPath := os.Getenv("CRONOS_CONFIG_PATH")

	viper.SetConfigType("yaml")
	viper.SetConfigName("cronos-conf.yaml")

	viper.AddConfigPath(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func backupTask(s types.Schedule) error {

	now, err := time.Parse("2006-01-02T15:04", time.Now().Format("2006-01-02T15:04"))
	if err != nil {
		return err
	}

	if s.Type == types.TimeSchedule {

		bkpTime, err := time.Parse("2006-01-02T15:04", s.TimeOrInterval)
		if err != nil {
			return err
		}

		dur := bkpTime.Sub(now)

		if dur.Milliseconds() > 0 {
			<-time.NewTimer(dur).C

			err = ziputils.CreateZipFile(s.SourcePath, s.DestPath)
			return err
		}

	} else if s.Type == types.IntervalSchedule {

		var d, h, m int

		startTime, err := time.Parse("2006-01-02T15:04", s.StartTime)
		if err != nil {
			return err
		}

		_, err = fmt.Sscanf(s.TimeOrInterval, "%d-%d-%d", &d, &h, &m)
		if err != nil {
			return err
		}

		for startTime.Before(now.Add(time.Minute)) {
			startTime = startTime.Add(time.Duration(d * int(time.Hour) * 24))
			startTime = startTime.Add(time.Duration(h * int(time.Hour)))
			startTime = startTime.Add(time.Duration(m * int(time.Minute)))
		}

		<-time.NewTimer(startTime.Sub(now)).C
		if err = ziputils.CreateZipFile(s.SourcePath, s.DestPath); err != nil {
			return err
		}
		backupTask(s)
	}

	return nil
}
