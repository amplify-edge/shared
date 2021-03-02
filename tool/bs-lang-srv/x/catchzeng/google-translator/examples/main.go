package main

import (
	"log"

	"go.amplifyedge.org/shared-v2/tool/bs-lang-srv/x/catchzeng/google-translator/translator"
	"golang.org/x/text/language"
)

func main() {
	translatedText, err := translator.Translate("Hello World", language.English, language.SimplifiedChinese)
	if err != nil {
		panic(err)
	}
	log.Printf("translated: %s", translatedText)

	err = translator.TranslateFile("./README.md", ".", "README_T.md", true, language.English, language.SimplifiedChinese)
	if err != nil {
		panic(err)
	}
}
