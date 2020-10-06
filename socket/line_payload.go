package socket

const (
	TYPE_ERROR = "WRITE_ERROR"
	TYPE_OUT = "WRITE_OUT"
)

type OutMessage struct {
	OutType string `json:"type"`
	Message string `json:"message"`
}

func MakeOutError(e string) OutMessage {
	return OutMessage{
		TYPE_ERROR,
		e,
	}
}

func MakeOutNormal(e string) OutMessage {
	return OutMessage{
		TYPE_OUT,
		e,
	}
}