package controllers

import "regexp"
import "os"
import "fmt"
import "io/ioutil"
import "net/url"

type PathContr interface {
  possible(path string) (bool)
  folder(path string)
  contentType(path string) (string)
}

type PathController struct {
  path string
}

const (
  File = "index.html"
)

func (contr *PathController) setPath(root string, path string) {
  // wd, _ := os.Getwd()
  rPath, _ := url.QueryUnescape(path)
  contr.path = root + rPath
}

func (contr *PathController) possible() (bool) {
  _, err := os.Stat(contr.path)
  if err == nil {
    contr.folder()
    return true
  }
  fmt.Println(contr.path)
  fmt.Println(err)
  return false
}

func (contr *PathController) folder() {
  if res, _ := regexp.MatchString(".*/$", contr.path); res {
    contr.path += File
  }
}

func (contr *PathController) read() (f []byte, err error) {
  f, err = ioutil.ReadFile(contr.path)
  return f, err
}

func (contr *PathController) contentType() (string) {
  re := regexp.MustCompile(".*\\.").Split(contr.path, -1)
  fmt.Print("re: ")
  fmt.Println(re)
	return PossibleTypes[re[1]]
}

var (
  PossibleTypes = map[string]string {
  "png": "image/png",
  "jpg": "image/jpeg",
  "jpeg": "image/jpeg",
  "swf": "application/x-shockwave-flash",
  "html": "text/html",
  "gif": "image/gif",
  "css": "text/css",
  "js": "application/javascript",
  "json": "application/json",
  "txt": "application/text",
}
)
