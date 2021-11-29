package service

import (
	"context"
	"errors"
	"log"
	"strings"

	gl "github.com/drhodes/golorem"
)

var ErrRequestTypeNotFound = errors.New("Resquest type only valid for word,sentence and paragraph")

type Service interface {
	LoremGenerate(ctx context.Context, requessType string, minx, max int) (string, error)
}

type LoremService struct{}

func (LoremService) LoremGenerate(_ context.Context, requestType string, min, max int) (string, error) {
	log.Println("hi from service")
	var result string = ""
	var err error = nil
	if strings.EqualFold(requestType, "Word") {
		result = gl.Word(min, max)
	} else if strings.EqualFold(requestType, "Sentence") {
		result = gl.Sentence(min, max)
	} else if strings.EqualFold(requestType, "Paragraph") {
		result = gl.Paragraph(min, max)
	} else {
		err = ErrRequestTypeNotFound
	}
	return result, err
}
