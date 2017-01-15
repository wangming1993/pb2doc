package pb

type Services struct {
	List    *Service
	Comment string
}

type Service struct {
	Name     string
	Request  *Message
	Response *Message
	Comment  string
}
