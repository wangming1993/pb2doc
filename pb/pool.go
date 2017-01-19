package pb

var ParsedMessages map[string]*Message
var ParsingProto map[string]bool

func init() {
	ParsedMessages = make(map[string]*Message)
	ParsingProto = make(map[string]bool)
}
