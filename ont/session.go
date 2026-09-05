package ont

import (
	"io"
	"net/http"
)

type Session struct {
	*http.Client
	Endpoint string
}

func (s *Session) getAndClose(url string) {
	resp, err := s.Get(url)
	if err != nil {
		return
	}
	closeBody(resp.Body)
}

func closeBody(body io.ReadCloser) {
    if body == nil {
        return
    }

    _, _ = io.Copy(io.Discard, body)
    _ = body.Close()
}
