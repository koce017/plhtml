package util

import "strings"

func Unescape(str string) string {
    str = strings.Replace(str, "\n", `\n`, -1)
    str = strings.Replace(str, "\t", `\t`, -1)
    str = strings.Replace(str, "\\", `\`, -1)
    return str
}
