package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/emirpasic/gods/maps/linkedhashmap"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
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

	_, res, err := findCacheContent(jsonCachePath, key, nil)

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

// AddToCache adds translation to cache
// we don't check for duplicates, nor we will.
func AddToCache(jsonCachePath, k, language, translated string) error {
	//var encodedKey bytes.Buffer
	//json.HTMLEscape(&encodedKey, []byte(k))
	//key := encodedKey.String()
	key := escapeWord(k)
	data, res, err := findCacheContent(jsonCachePath, key, nil)
	if err != nil {
		return err
	}
	if res == "" {
		data, err = sjson.SetBytes(data, key, []translation{
			{
				LangCode: "en",
				Trans:    k,
			},
		})
		if err != nil {
			return err
		}
	}
	data, err = sjson.SetBytes(data, key+".-1", translation{
		LangCode: language,
		Trans:    translated,
	})
	if err != nil {
		return err
	}
	return ioutil.WriteFile(jsonCachePath, data, 0755)
}

func findCacheContent(cachePath, k string, data []byte) ([]byte, string, error) {
	var err error
	if data == nil {
		data, err = getCacheContent(cachePath)
		if err != nil {
			return nil, "", err
		}
	}
	//var encodedKey bytes.Buffer
	//json.HTMLEscape(&encodedKey, []byte(k))
	//key := encodedKey.String()
	key := escapeWord(k)
	value := gjson.Get(string(data), key)
	return data, value.String(), nil
}

// find out if a given translation key exists
func CacheExists(cachePath, key string) bool {
	_, key, err := findCacheContent(cachePath, key, nil)
	if err != nil {
		return false
	}
	return key != ""
}

func escapeWord(w string) string {
	// replace all occurences of these characters
	re := regexp.MustCompile("([?*~])")
	w = re.ReplaceAllStringFunc(w, func(s string) string {
		return fmt.Sprintf("\\%s", s)
	})
	return w
}

func unescapeWord(w string) string {
	re := regexp.MustCompile(`\\`)
	w = re.ReplaceAllString(w, "")
	return w
}
