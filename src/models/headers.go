package models

type Headers map[string]string

func (hdrs Headers) String() (string) {
	str := ""
	for key, val := range hdrs {
		str += key + ": " + val + "\r\n"
	}
	return str
}
