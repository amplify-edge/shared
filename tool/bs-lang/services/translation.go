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
	Words []string
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
	tlClient, err := tl.NewTranslationClient(ctx)
	if err != nil {
		return nil, err
	}
	defer tlClient.Close()
	words := getTranslateWords(m)
	var wordsTranslated []Translate
	for _, lang := range languages {
		t := Translate{}
		for _, w := range words {
			// skip english
			if lang == "en" {
				continue
			}
			esc := escapeWord(w)
			// hit cache first
			req := &translatepb.TranslateTextRequest{
				Contents:           []string{fmt.Sprintf("%q", esc)},
				MimeType:           "text/plain",
				SourceLanguageCode: "en",
				Parent:             project,
				TargetLanguageCode: lang,
			}
			translated, err := FindTransFromCache(cacheFile, w, lang)
			if err != nil {
				// get translation from google
				gtranslated, err := tlClient.TranslateText(ctx, req)
				if err != nil {
					// retry
					var retryErr []error
					for i := 0; i < tries; i++ {
						time.Sleep(delay)
						gtranslated, err = tlClient.TranslateText(ctx, req)
						if err != nil {
							retryErr = append(retryErr, err)
						}
					}
					if len(retryErr) == tries {
						return nil, err
					}
				}
				if gtranslated != nil {
					for _, v := range gtranslated.Translations {
						translated = v.GetTranslatedText()
						// Add it to cache
						if err = AddToCache(cacheFile, w, lang, translated); err != nil {
							return nil, err
						}
					}
				}
			}
			t.Words = append(t.Words, translated)
		}
		t.Lang = lang
		wordsTranslated = append(wordsTranslated, t)
	}
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
					mapLang.Put(it.Key(), strings.TrimSpace(tr.Words[i]))
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
