package types

const (
	TimeSchedule     ScheduleType = 1
	IntervalSchedule ScheduleType = 2
)

type ScheduleType int

type BkpConf struct {
	BkpSchedules []Schedule `mapstructure:"bkpschedules"`
}

type Schedule struct {
	Type           ScheduleType `mapstructure:"type"`
	StartTime      string       `mapstructure:"starttime"`
	TimeOrInterval string       `mapstructure:"timeorinterval"`
	SourcePath     string       `mapstructure:"sourcepath"`
	DestPath       string       `mapstructure:"destpath"`
}
