package services

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/emirpasic/gods/maps/linkedhashmap"
	"github.com/tidwall/gjson"
)

type translation struct {
	LangCode string `json:"lang_code"`
	Trans    string `json:"trans"`
}

// JSONMap convert json to arb
func JSONMap(data []byte) (*linkedhashmap.Map, error) {
	m := linkedhashmap.New()

	err := m.FromJSON(data)

	if err != nil {
		return nil, err
	}

	it := m.Iterator()
	out := linkedhashmap.New()

	for it.Next() {
		key := it.Key().(string)
		keyExists := "@" + key

		_, ok := m.Get(keyExists)
		if !ok {
			item := ArbAttr{
				Description:  "",
				Type:         "text",
				Placeholders: make(map[string]string),
			}
			out.Put(it.Key(), it.Value())
			out.Put(keyExists, item)
		} else {
			out.Put(it.Key(), it.Value())
		}
	}
	return out, nil
}

func getCacheContent(jsonCachePath string) ([]byte, error) {
	_, ext := getFileNameAndExtension(jsonCachePath)
	data, err := ioutil.ReadFile(jsonCachePath)
	if err != nil {
		return nil, err
	}
	if ext != "json" {
		return nil, errors.New("invalid extension, supports only json")
	}
	return data, nil
}

// FindTrans finds a translation from cache file
// it takes raw byte slice or filepath to json / arb cache
func FindTransFromCache(jsonCachePath, key, language string) (string, error) {
	res, err := findCacheContent(jsonCachePath, key)

	var tl []translation

	if err = json.Unmarshal([]byte(res), &tl); err != nil {
		return "", errors.New("error unmarshaling available translations")
	}

	for _, t := range tl {
		if t.LangCode == language {
			return t.Trans, nil
		}
	}
	return "", errors.New("no translations found for " + key + " and " + language)
}

func findCacheContent(cachePath, key string) (string, error) {
	data, err := getCacheContent(cachePath)
	if err != nil {
		return "", err
	}
	value := gjson.Get(string(data), key)
	return value.String(), nil
}

// find out if a given translation key exists
func CacheExists(cachePath, key string) bool {
	key, err := findCacheContent(cachePath, key)
	if err != nil {
		return false
	}
	return key != ""
}
