package parser

import "regexp"

var (
	SingleCommentCompiler     *regexp.Regexp
	MultiCommentStartCompiler *regexp.Regexp
	MultiCommentEndCompiler   *regexp.Regexp
)

func init() {
	SingleCommentCompiler, _ = regexp.Compile("^\\s*(//.*)|(/\\*.*\\*/)")
	MultiCommentStartCompiler, _ = regexp.Compile("^\\s*/\\*.*")
	MultiCommentEndCompiler, _ = regexp.Compile("^.*\\*/")
}

func StartWithMessage(line string) bool {
	return startWith(line, "message")
}

func StartWithService(line string) bool {
	return startWith(line, "service")
}

func StartWithEnum(line string) bool {
	return startWith(line, "enum")
}

func StartWithOneof(line string) bool {
	return startWith(line, "oneof")
}

func startWith(line, prefix string) bool {
	cp, _ := regexp.Compile("^\\s*" + prefix + " ")
	return cp.MatchString(line)
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
	c, _ := regexp.Compile("^\\s*message (.*)\\s{")
	matches := c.FindStringSubmatch(line)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func GetEnumName(line string) string {
	c, _ := regexp.Compile("^\\s*enum\\s+(.*)\\s{")
	matches := c.FindStringSubmatch(line)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func GetOneofName(line string) string {
	c, _ := regexp.Compile("^\\s*oneof\\s+(.*)\\s{")
	matches := c.FindStringSubmatch(line)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func EndWithBrace(line string) bool {
	c, _ := regexp.Compile("^\\s*}$")
	return c.MatchString(line)
}

func IsExtendType(line string) bool {
	return StartWithEnum(line) ||
		StartWithMessage(line) ||
		StartWithOneof(line) ||
		StartWithService(line)
}
