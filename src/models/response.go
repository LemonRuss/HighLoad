package models

import "time"
import "bytes"
import "strconv"

type Response struct {
	Status  	string
	HTTP	    string
  Headers 	Headers
	Body		  []byte
}

const (
	OK = "200 OK"
	Forbidden = "403 Forbidden"
	NotFound = "404 Not Found"
	NotAllowed = "405 Method Not Allowed"
	HTTP = "HTTP/1.1"
  Serv = "golang server"
  Close = "Close"
)

func (res *Response) Base() {
  res.HTTP = HTTP
  res.Headers = Headers {
   "Date": time.Now().String(),
   "Server": Serv,
   "Connection": Close,
 }
}

func (res *Response) NotAllowed() {
  res.Status = NotAllowed
}

func (res * Response) Forbidden() {
  res.Status = Forbidden
}

func (res *Response) NotFound() {
  res.Status = NotFound
}

func (res *Response) Succses(cont string, file []byte) {
  res.Status = OK
  res.Headers["Content-Length"] = strconv.Itoa(len(file))
  res.Headers["Content-Type"] = cont
}

func (res *Response) Byte() []byte  {
	var result bytes.Buffer

	result.WriteString(res.HTTP + " " + res.Status + "\r\n")
	result.WriteString(res.Headers.String() + "\r\n")
	result.WriteString(string(res.Body))

	return result.Bytes()
}
