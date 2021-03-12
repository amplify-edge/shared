package lib

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	// Id for the Google Sheet to process
	// if you have google sheets open in your browser, take a look at your URL bar
	// it should look like: https://docs.google.com/spreadsheets/d/1pmPPYLrHfSGLM-1MPYEGtbb9Z5iHFUL-xqXNFS0DyaM/edit#gid=1594323569
	// the GSheetId is this part: 1pmPPYLrHfSGLM-1MPYEGtbb9Z5iHFUL-xqXNFS0DyaM
	GSheetId string `json:"gsheets_id"`
	// Name of the spreadsheet tab that contains the localizations
	WorksheetName string `json:"worksheet_name"`
}

func NewConfigFromFile(filepath string) (*Config, error) {
	var c Config
	f, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(f, &c); err != nil {
		return nil, err
	}
	return &c, nil
}
