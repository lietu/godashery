package godashery


type Hello struct {
	Type string        `json:"type"`
}

func NewHello() *Hello {
	h := Hello{
		"hello",
	}

	return &h
}


type Notice struct {
	Type    string        `json:"type"`
	Message string        `json:"message"`
}

func NewNotice(message string) *Notice {
	n := Notice{
		"notice",
		message,
	}

	return &n
}
