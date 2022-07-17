package logs

const (
	recommendation int32 = 10071  //"U+2757"
	search         int32 = 128269 //"U+1F50D"
	weather        int32 = 9728   //"U+2600"
)

type application map[int32]string

var (
	applications = application{
		recommendation: "recommendation",
		search:         "search",
		weather:        "weather",
	}
)

//func getCharacter(log string) {
//	for index, char := range log {
//		fmt.Printf("Index: %d\tCharacter: %c\t\tCode Point: %U\t\t int32: %d\n", index, char, char, char)
//	}
//}

func checkForApplication(character int32) (found bool, application string) {
	value, found := applications[character]
	if found {
		return found, value
	}
	return found, ""
}

// Application identifies the application emitting the given log.
func Application(log string) string {
	//getCharacter(log)
	for _, v := range log {
		found, application := checkForApplication(v)
		if found {
			return application
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	out := []rune(log)
	for i, v := range out {
		if v == oldRune {
			out[i] = newRune
		}
	}
	return string(out)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return limit >= len([]rune(log))
}
