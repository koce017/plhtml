package util

import "strconv"

func FloatToString(number float64) string {
    return strconv.FormatFloat(number, 'f', 6, 64)
}

func StrToInt64(str string) (int64, error) {
    return strconv.ParseInt(str, 10, 64)
}

func StrToFloat64(str string) (float64, error) {
    return strconv.ParseFloat(str, 64)
}

func StrToBool(str string) (bool, error) {
    return strconv.ParseBool(str)
}
