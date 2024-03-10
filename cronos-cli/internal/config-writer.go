package internal

import (
	"cronos/types"
	"errors"
	"log"
	"os"

	"github.com/spf13/viper"
)

func ConfigWriter(cfg types.BkpConf) {
	viper.SetConfigName("cronos-conf")
	viper.SetConfigType("yaml")

	configPath := os.Getenv("CRONOS_CONFIG_PATH")

	viper.AddConfigPath(configPath)

	if _, err := os.Stat(configPath + "\\cronos-conf.yaml"); err != nil {

		if errors.Is(err, os.ErrNotExist) {
			viper.Set("bkpschedules", cfg.BkpSchedules)
			if err := viper.SafeWriteConfig(); err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	} else {

		var cfgAux types.BkpConf

		if err := viper.ReadInConfig(); err != nil {
			log.Fatal(err)
		}
		if err := viper.Unmarshal(&cfgAux); err != nil {
			log.Fatal(err)
		}

		cfgAux.BkpSchedules = append(cfgAux.BkpSchedules, cfg.BkpSchedules...)

		viper.Set("bkpschedules", cfgAux.BkpSchedules)

		if err := viper.WriteConfig(); err != nil {
			log.Fatal(err)
		}
	}
}
