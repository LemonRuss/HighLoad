package controllers

import "models"
import "strings"
import "fmt"

type ResContr interface {
  Gen(mthd string, path string, doc string) (res *models.Response)
  SetRoot(root string)
}

type ResController struct {
  doc string
}

const (
  GET = "GET"
  HEAD = "HEAD"
)

func (contr *ResController) SetRoot(root string) {
  contr.doc = root
  fmt.Println("Path:", contr.doc)
}

func (conrt *ResController) Gen(mthd string, path string, doc string) (*models.Response) {
  res := new(models.Response)
  res.Base()
  if mthd != HEAD && mthd != GET {
    res.NotAllowed()
    return res
  }
  if strings.Contains(path, "../") {
    res.Forbidden()
    return res
  }
  contr := new(PathController)

  contr.setPath(conrt.doc, path)
  fmt.Println(contr.path)
  if !contr.possible() {
    res.NotFound()
    return res
  }
  f, err := contr.read()
  if err != nil {
    res.Forbidden()
    return res
  }
  if mthd == "GET" {
    res.Body = f
  }
  res.Succses(contr.contentType(), f)
  return res
}
