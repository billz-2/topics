package topics

import (
	"runtime"
	"strings"
	"unicode"
)

func t() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	str := strings.Split(frame.Function, ".")
	return toSausagesCase(str[len(str)-1])

}
func toSnakeCase(in string) string {
	var (
		runes  = []rune(in)
		length = len(runes)
		out    []rune
	)
	for idx := 0; idx < length; idx++ {
		if idx > 0 && unicode.IsUpper(runes[idx]) &&
			((idx+1 < length && unicode.IsLower(runes[idx+1])) || unicode.IsLower(runes[idx-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[idx]))
	}

	return string(out)
}

func toSausagesCase(in string) string {
	s := toSnakeCase(in)
	return strings.ReplaceAll(s, "__", ".")

}
