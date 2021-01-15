package response

import (
	"errors"
	"strings"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

var (
	errHTTPHeaderCleanDidNotSet = "http header did not set"
)

type headerClean struct {
	prefix string
	suffix string
	next   Parser
}

func (t *headerClean) Parse(info *api.HTTPMethod, firstTag string, tags ...string) (err error) {
	if strings.HasPrefix(firstTag, t.prefix) && strings.HasSuffix(firstTag, t.suffix) {
		if len(tags) == 2 {
			if info.ResponseHeadersClean == nil {
				info.ResponseHeadersClean = make(map[string]string)
			}
			info.ResponseHeadersClean[tags[0]] = strings.TrimRight(strings.TrimLeft(strings.TrimSpace(tags[1]), "{"), "}")
			return
		}
		return errors.New(errHTTPHeaderCleanDidNotSet)
	}
	return t.next.Parse(info, firstTag, tags...)
}

// NewHeader ...
func NewHeaderClean(prefix string, suffix string, next Parser) Parser {
	return &headerClean{
		prefix: prefix,
		suffix: suffix,
		next:   next,
	}
}
