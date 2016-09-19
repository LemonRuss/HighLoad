package controllers

import "models"
import "strings"

type ReqContr interface {
    Parse(str string) (*models.Request, error)
}

type ReqController struct {}

func (contr* ReqController) Parse(str string) (*models.Request, error) {
	req := new(models.Request)
	fields := strings.Split(str, "\r\n")
	err := req.Fill(fields)
	if err != nil {
		return nil, err
	}
	return req, nil
}
