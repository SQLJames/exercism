package parsinglogfiles

import "regexp"

var (
	validLine    = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	badSplit     = regexp.MustCompile(`<[-~*=]*>`)
	password     = regexp.MustCompile(`(?i)"[^"]*password[^"]*"`)
	endOfLine    = regexp.MustCompile(`end-of-line(\d+)`)
	containsUser = regexp.MustCompile(`User\s+(\S+)`)
)

func IsValidLine(text string) bool {
	return validLine.MatchString(text)
}

func SplitLogLine(text string) []string {
	matches := badSplit.Split(text, -1)
	if len(matches) == 0 {
		return []string{text}
	}
	return matches
}

func CountQuotedPasswords(lines []string) int {
	var count int
	for _, line := range lines {
		count += len(password.FindAllStringIndex(line, -1))
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	return endOfLine.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {

	var result []string
	for _, line := range lines {
		if match := containsUser.FindStringSubmatch(line); match != nil {
			username := match[1] // Extract the username
			result = append(result, "[USR] "+username+" "+line)
		} else {
			result = append(result, line)
		}
	}
	return result
}
