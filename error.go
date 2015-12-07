package simpletwitter

// Error is customized error for tweet get
type Error struct {
	Op            string // the failing Operation (Request, Redirected, Parse)
	GetURL        string // the getting url
	RedirectedURL string // the definitive url
	Err           error  // the reason the get failed
}

func (e *Error) Error() string {
	return "please input error log"
}
