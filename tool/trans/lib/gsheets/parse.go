package gsheets

import (
	"fmt"
	"go.amplifyedge.org/shared-v2/tool/trans/lib/errutil"
	"strconv"
	"strings"
)

type sheetRange struct {
	start   string
	end     string
	postEnd string
}

func (s *sheetRange) Start() string {
	return s.start
}

func (s *sheetRange) End() string {
	return s.end
}

func (s *sheetRange) PostEnd() string {
	return s.postEnd
}

func (s *sheetRange) String() string {
	return fmt.Sprintf("%s:%s", s.start, s.end)
}

func parseGsheetRange(rangeString string) (*sheetRange, error) {
	ws := strings.Split(rangeString, "!")
	if len(ws) < 2 {
		return nil, errutil.New(errutil.ErrParsingRangeResponse, fmt.Errorf("invalid range string input"))
	}
	indexes := strings.Split(ws[1], ":")
	n := string(indexes[1][1:])
	endIdx, err := strconv.Atoi(n)
	if err != nil {
		return nil, errutil.New(errutil.ErrParsingRangeResponse, fmt.Errorf("invalid index"))
	}
	return &sheetRange{
		start:   indexes[0],
		end:     indexes[1],
		postEnd: fmt.Sprintf("%s%d", string(indexes[0][:1]), endIdx+1),
	}, nil
}
