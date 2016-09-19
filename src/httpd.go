package main

import "net"
import "fmt"
import "runtime"
import "models"
import "controllers"
import "flag"

var root string

func main() {
  fmt.Println("Launching models...")

  numCPU := flag.Int("c", runtime.NumCPU(), "")
  doc := flag.String("r", "../httptest", "")
  flag.Parse()
  root = *doc
  fmt.Println(root, " ____")

  ad, _ := net.ResolveTCPAddr("tcp", "localhost:80")
  ln, _ := net.ListenTCP("tcp", ad)

  runtime.GOMAXPROCS(*numCPU)
  for {
    conn, err := ln.Accept()
    if err != nil {
      fmt.Println("Error received:", err)
    }
    if conn != nil {
      fmt.Println("Starting to serve connection:", conn)
      go serve(conn)
    }
  }
}

func serve(conn net.Conn) {
  defer conn.Close()
  req, err := reqFrom(conn)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(req)
  resContr := new(controllers.ResController)
  resContr.SetRoot(root)
  res := new(models.Response)
  res = resContr.Gen(req.Method, req.URI, "/")
  conn.Write(res.Byte())
}

func reqFrom(conn net.Conn) (*models.Request, error) {
  buf := make([]byte, 2048)
  _, rdErr := conn.Read(buf)
  if rdErr != nil {
    fmt.Println(rdErr)
    return nil, rdErr
  }
  contr := new(controllers.ReqController)
  req, psErr := contr.Parse(string(buf))
  return req, psErr
}
