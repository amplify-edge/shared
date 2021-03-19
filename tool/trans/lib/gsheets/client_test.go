package gsheets_test

import (
	"github.com/stretchr/testify/require"
	"go.amplifyedge.org/shared-v2/tool/trans/lib"
	"go.amplifyedge.org/shared-v2/tool/trans/lib/gsheets"
	"testing"
)

var (
	cfg *lib.Config
	c   *gsheets.Client
)

func init() {
	cfg, _ = lib.NewConfigFromFile("./testdata/config.json")
}

func TestGsheetsApi(t *testing.T) {
	var err error
	c, err = gsheets.NewClient("./testdata/config.json", cfg)
	require.NoError(t, err)

	locName := c.WorksheetName()
	require.Equal(t, "localizations", locName)

	cells, err := c.Localizations()
	require.NoError(t, err)
	require.NotEqual(t, nil, cells)
	t.Logf("localizations sheets cells: %v", cells)

	lastIdx, err := c.LastIdx()
	require.NoError(t, err)
	require.NotEqual(t, "", lastIdx)
	t.Logf("indexes range: %s", lastIdx)
}
