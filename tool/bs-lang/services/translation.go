package services

import (
	"context"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	tl "cloud.google.com/go/translate/apiv3"
	"github.com/emirpasic/gods/maps/linkedhashmap"
	translatepb "google.golang.org/genproto/googleapis/cloud/translate/v3"
)

// TranslatedMaps struct
type TranslatedMaps struct {
	Langs []string
	Maps  []*linkedhashmap.Map
}

// Translate struct
type Translate struct {
	Lang string
	// Words string
	Words map[string]string
}

// ArbAttr struct
type ArbAttr struct {
	Description  string            `json:"description"`
	Type         string            `json:"type"`
	Placeholders map[string]string `json:"placeholders"`
}

func getGoogleProjectId() string {
	credPath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	data, err := ioutil.ReadFile(credPath)
	if err != nil {
		log.Fatalf("error reading google credentials path, set it first via GOOGLE_APPLICATION_CREDENTIALS env var")
	}
	res := gjson.GetBytes(data, "project_id")
	return fmt.Sprintf("projects/%s", res.String())
}

// translate a string from languages to language
func getTemplateWords(m *linkedhashmap.Map, delay time.Duration, tries int, languages []string, cacheFile string) ([]Translate, error) {
	project := getGoogleProjectId()
	ctx := context.Background()

	initialWords := getTranslateWords(m)
	var wordsTranslated []Translate

	// google translate client
	tlClient, err := tl.NewTranslationClient(ctx)
	if err != nil {
		return nil, err
	}

	detectLangReq := &translatepb.DetectLanguageRequest{
		Parent:   project,
		Source:   &translatepb.DetectLanguageRequest_Content{Content: initialWords[0]},
		MimeType: "text/plain",
	}

	l, err := tlClient.DetectLanguage(ctx, detectLangReq)
	if err != nil {
		return nil, err
	}

	originLang := l.GetLanguages()[0].LanguageCode

	for _, lang := range languages {
		t := Translate{
			Lang:  lang,
			Words: map[string]string{},
		}
		untranslated, translated := findTransFromCache(cacheFile, originLang, lang, initialWords...)

		for k, v := range translated {
			t.Words[k] = v
		}

		if len(untranslated) > 0 {
			var gtranslated *translatepb.TranslateTextResponse
			req := &translatepb.TranslateTextRequest{
				Contents:           untranslated,
				MimeType:           "text/plain",
				SourceLanguageCode: originLang,
				Parent:             project,
				TargetLanguageCode: lang,
			}

			gtranslated, err = tlClient.TranslateText(ctx, req)
			if err != nil {
				return nil, err
			}

			if gtranslated != nil {
				for i, v := range gtranslated.GetTranslations() {
					t.Words[untranslated[i]] = v.GetTranslatedText()
					if err = AddToCache(cacheFile, untranslated[i], originLang, lang, v.GetTranslatedText()); err != nil {
						return nil, err
					}
				}
			}

		}
		wordsTranslated = append(wordsTranslated, t)
	}
	_ = tlClient.Close()
	return wordsTranslated, nil
}

func getTranslatedMaps(WordsTranslated []Translate, m *linkedhashmap.Map, full bool) (*TranslatedMaps, error) {

	translatedMaps := &TranslatedMaps{}
	for _, tr := range WordsTranslated {
		mapLang := linkedhashmap.New()
		it := m.Iterator()
		i := 0
		for it.Next() {
			if !strings.HasPrefix(it.Key().(string), "@") {
				if len(tr.Words) > i {
					k := it.Value()
					v := tr.Words[k.(string)]
					mapLang.Put(it.Key(), strings.TrimSpace(v))
					i++
				}
			} else if full {
				mapLang.Put(it.Key(), it.Value())
			}
		}

		translatedMaps.Maps = append(translatedMaps.Maps, mapLang)
		translatedMaps.Langs = append(translatedMaps.Langs, tr.Lang)

	}
	return translatedMaps, nil
}

func getTranslateWords(m *linkedhashmap.Map) []string {
	it := m.Iterator()
	var out []string
	for it.Next() {
		if !strings.HasPrefix(it.Key().(string), "@") {
			v, ok := it.Value().(string)
			if ok {
				out = append(out, v)
			}
		}
	}
	return out
}
