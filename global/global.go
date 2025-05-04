package global

import "time"

var (
	G_UrlChan      chan string
	G_ErrorArr     []Errorstruct
	WithChannel    time.Duration
	WithoutChannel time.Duration
)

type Errorstruct struct {
	Indicator string
	Function  string
	Line      string
	Errorstr  string
}
