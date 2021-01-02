package tool

import "strconv"

func Int64ToInt (i64 int64) (i int) {
	str := strconv.FormatInt(i64, 10)
	i, _ = strconv.Atoi(str)
	return 
}
