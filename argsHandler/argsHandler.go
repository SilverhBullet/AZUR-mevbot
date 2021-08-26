package argshandler

import (
	"strconv"
	"strings"
	"time"
)

//rewrite some redis instructions

func RedisCmdRewrite(args [][]string) {
	nowUnixTime := time.Now()

	for i, arg := range args {
		switch strings.ToLower(arg[0]) {
		case "setex":
			if len(arg) == 4 && nowUnixTime != 0 {
				exDuration, err := strconv.ParseInt(arg[2], 10, 64)
				if err == nil {
					args[i] = []string{"setexat", arg[1], strconv.FormatInt(nowUnixTime+exDuration, 10), arg[3]}
				}
			}
		case "expire":
			if len(arg) == 3 && nowUnixTime != 0 {
				exDuration, err := strconv.ParseInt(arg[2], 10, 64)
				if err == nil {
					args[i] = []string{"expireat", arg[1], strconv.FormatInt(nowUnixTime+exDuration, 10)}
				}
			}
		case "lexpire":
			if len(arg) == 3 && nowUnixTime != 0 {
				exDuration, err := strconv.ParseInt(arg[2], 10, 64)
				if err == nil {
					args[i] = []string{"lexpireat", arg[1], strconv.FormatInt(nowUnixTime+exDuration, 10)}
				}
			}
		case "hexpire":
			if len(arg) == 3 && nowUnixTime != 0 {
				exDuration, err := strconv.ParseInt(arg[2], 10, 64)
				if err == nil {
					args[i] = []string{"lexpireat", arg[1], strconv.FormatInt(nowUnixTime+exDuration, 10)}
				}
			}
		default:
		}
	}
}
