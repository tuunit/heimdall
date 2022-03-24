package extractors

import (
	"errors"
	"strings"

	"github.com/dadrus/heimdall/pipeline"
)

type HeaderValueExtractStrategy struct {
	Name   string
	Prefix string
}

func (es HeaderValueExtractStrategy) GetAuthData(s pipeline.AuthDataSource) (string, error) {
	if val := s.Header(es.Name); len(val) != 0 {
		return strings.TrimSpace(strings.TrimPrefix(val, es.Prefix)), nil
	} else {
		return "", errors.New("no authentication data present")
	}
}
