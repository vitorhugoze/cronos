package cmd

import (
	"cronos/internal"
	"cronos/types"
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cronos",
	Short: "Back scheaduler",
	Long:  "Customizable backup scheaduler that can run just once or on configurable intervals",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("running a command on my first app")
	},
}

func init() {

	var instant, interval, delete string

	addScheadule := &cobra.Command{
		Use:   "add",
		Short: "Add backup scheadule",
		Long:  "Create a backup scheadule to run once or on a interval Usage: add <source_path> <dest_path>",
		Args:  cobra.ExactArgs(2),

		Run: func(cmd *cobra.Command, args []string) {

			var (
				bkpType   types.ScheduleType
				timeOrInt string
				startTime = ""
			)

			if instant != "" {

				bkpType = types.TimeSchedule

				//check instant format
				instant, err := time.Parse("2006-01-02T15:04", instant)
				if err != nil {
					log.Fatal(err)
				}

				timeOrInt = instant.Format("2006-01-02T15:04")
			}

			if interval != "" {
				var d, h, m int

				bkpType = types.IntervalSchedule

				startTime = time.Now().Format("2006-01-02T15:04")

				_, err := fmt.Sscanf(interval, "%d-%d-%d", &d, &h, &m)
				if err != nil {
					log.Fatal(err)
				}

				timeOrInt = interval
			}

			if delete != "" {
				var d, h, m int

				_, err := fmt.Sscanf(delete, "%d-%d-%d", &d, &h, &m)
				if err != nil {
					log.Fatal(err)
				}
			}

			bkpCfg := types.BkpConf{
				BkpSchedules: []types.Schedule{
					{
						Type:           bkpType,
						StartTime:      startTime,
						TimeOrInterval: timeOrInt,
						DeleteInterval: delete,
						SourcePath:     args[0],
						DestPath:       args[1],
					},
				},
			}

			internal.ConfigWriter(bkpCfg)
		},
	}

	addScheadule.Flags().StringVarP(&instant, "time", "t", "", "Set time when backup will run Format: yyyy-MM-ddTHH:mm")
	addScheadule.Flags().StringVarP(&interval, "interval", "i", "", "Set the interval for continuously running backups Format: dd-HH-mm")
	addScheadule.Flags().StringVarP(&delete, "delete", "d", "", "Set the interval for the backup exclusion afther the backup is done Format: dd-HH-mm")

	rootCmd.AddCommand(addScheadule)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
