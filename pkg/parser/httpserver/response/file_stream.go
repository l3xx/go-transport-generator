package response

import (
	"errors"
	"strings"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

var (
	errHTTPFileStreamDataSet = "http file stream data did not set"
)

type streamData struct {
	prefix string
	suffix string
	next   Parser
}

func (t *streamData) Parse(info *api.HTTPMethod, firstTag string, tags ...string) (err error) {
	if strings.HasPrefix(firstTag, t.prefix) && strings.HasSuffix(firstTag, t.suffix) {
		if len(tags) == 1 {
			info.ResponseStream = tags[0]
			return
		}
		if len(tags) == 2 {
			info.ResponseStream = tags[0]
			info.ResponseFileName = tags[1]
			return
		}
		return errors.New(errHTTPFileStreamDataSet)
	}
	return t.next.Parse(info, firstTag, tags...)
}

// NewFile ...
func NewFileStream(prefix string, suffix string, next Parser) Parser {
	return &streamData{
		prefix: prefix,
		suffix: suffix,
		next:   next,
	}
}
