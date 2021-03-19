package gsheets

import (
	"context"
	"go.amplifyedge.org/shared-v2/tool/trans/lib"
	"go.amplifyedge.org/shared-v2/tool/trans/lib/errutil"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"sync"
	"time"
)

type Client struct {
	cfg *lib.Config
	svc *sheets.SpreadsheetsService
}

// create new google sheets client
func NewClient(googleCredsPath string, cfg *lib.Config) (*Client, error) {
	ctx := context.Background()

	s, err := sheets.NewService(ctx, option.WithCredentialsFile(googleCredsPath), option.WithScopes(sheets.SpreadsheetsScope))
	if err != nil {
		return nil, errutil.New(errutil.ErrRetrieveSheetsClient, err)
	}

	return &Client{
		cfg: cfg,
		svc: s.Spreadsheets,
	}, nil

}

func (s *Client) WorksheetName() string {
	return s.cfg.WorksheetName
}

func (s *Client) Localizations() ([][]string, error) {
	return s.rawVals(s.WorksheetName())
}

func (s *Client) LastIdx() (*sheetRange, error) {
	var vr sheets.ValueRange

	myval := []interface{}{"Last Row Plus One"}
	vr.Values = append(vr.Values, myval)

	// Add a value to last row + 1 in order to know the range of cell
	res, err := s.svc.Values.Append(s.cfg.GSheetId, "A1:Z", &vr).ValueInputOption("USER_ENTERED").InsertDataOption("INSERT_ROWS").Do()
	if err != nil {
		return nil, err
	}

	// i am talking about the TableRange
	sr, err := parseGsheetRange(res.TableRange)
	if err != nil {
		return nil, err
	}

	// then we clear the value that we appended
	cvr := &sheets.ClearValuesRequest{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = s.svc.Values.Clear(s.cfg.GSheetId, sr.PostEnd(), cvr).Context(ctx).Do()
	if err != nil {
		return nil, err
	}
	// all in all this will get us the range of the worksheet (from start to finish)
	return sr, nil
}

func (s *Client) rawVals(sheetName string) ([][]string, error) {
	res, err := s.svc.Values.Get(s.cfg.GSheetId, sheetName).Do()
	if err != nil {
		return nil, err
	}
	values := res.Values
	lastOffset := len(values)
	result := make([][]string, lastOffset)
	var wg sync.WaitGroup

	for i, row := range values {
		wg.Add(1)
		i := i
		row := row
		go func() {
			result[i] = make([]string, len(row))
			for j, col := range row {
				result[i][j] = col.(string)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	return result, nil
}
