package models

import "strings"
import "errors"

type Fields []string

type Request struct {
	Method   string
	URI      string
	HTTP     string
  Headers  Headers
	Body     string
}

func (req *Request) Fill(f Fields) (error) {
	parts := strings.Split(f[0], " ")
	if (len(parts) >= 3) {
		req.Method = parts[0]
		req.URI = strings.Split(parts[1], "?")[0]
		req.HTTP = parts[2]
		return nil
	} else {
		return errors.New("Bad request")
	}
  for _, field := range f {
    h := strings.Split(field, ": ")
		if len(h) == 2 {
			req.Headers[h[0]] = h[1]
		}
  }
	return nil
}
