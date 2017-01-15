package parser

import (
	"fmt"
	"regexp"
)

var (
	MessageCompiler           *regexp.Regexp
	ServiceCompiler           *regexp.Regexp
	SingleCommentCompiler     *regexp.Regexp
	MultiCommentStartCompiler *regexp.Regexp
	MultiCommentEndCompiler   *regexp.Regexp
)

func init() {
	fmt.Println("init...")
	MessageCompiler, _ = regexp.Compile("^\\s?message ")
	ServiceCompiler, _ = regexp.Compile("^\\s?service ")
	SingleCommentCompiler, _ = regexp.Compile("^\\s*(//.*)|(/\\*.*\\*/)")
	MultiCommentStartCompiler, _ = regexp.Compile("^\\s?/\\*.*")
	MultiCommentEndCompiler, _ = regexp.Compile("^\\s?.*\\*/")
}

func StartWithMessage(line string) bool {
	return MessageCompiler.MatchString(line)
}

func StartWithService(line string) bool {

	return false
}

func IsSingleComment(line string) bool {
	return SingleCommentCompiler.MatchString(line)
}

func StartMultiComment(line string) bool {
	return MultiCommentStartCompiler.MatchString(line)
}

func EndMultiComment(line string) bool {
	return MultiCommentEndCompiler.MatchString(line)
}

func GetMessageName(line string) string {
	c, _ := regexp.Compile("^\\s?message (.*)\\s{")
	matches := c.FindStringSubmatch(line)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func EndWithBrace(line string) bool {
	c, _ := regexp.Compile("^\\s?}$")
	return c.MatchString(line)
}

func ReadComment(lines []string) (string, int) {
	total := len(lines)
	var (
		i       int
		step    int
		comment string
	)
	for {
		if i >= total {
			break
		}
		line := lines[i]
		// Must done otherwise dead foreach
		i++
		if IsSingleComment(line) {
			step += 1
			comment += line + "\n"
		} else if StartMultiComment(line) {
			for {
				step += 1
				comment += line + "\n"
				line = lines[i]

				if EndMultiComment(line) {
					step++
					comment += line + "\n"
					break
				}
				i++
			}
		} else {
			return comment, step
		}
	}
	return comment, step
}
