package data

import "time"

type RollingType uint

const (
	RollingTimeHour = 1
	RollingTimeDay  = 2
	RollingTimeMin  = 3
)

//type Rolling struct {
//	Type RollingType
//}

type Task struct {
	Name        string      `json:"name"`
	RollingType RollingType `json:"rolling_type"`
	Compress    bool        `json:"compress"`
}

type Data struct {
	DataType string
	Time     time.Time
	Data     []byte
}
