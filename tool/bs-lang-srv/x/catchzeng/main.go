package main

import (
	"log"

	"github.com/CatchZeng/google-translator/translator"
	"golang.org/x/text/language"
)

func main() {
	translatedText, err := translator.Translate("Hello World", language.English, language.SimplifiedChinese)
	if err != nil {
		panic(err)
	}
	log.Printf("translated: %s", translatedText)
	//translated: 你好，世界

	err = translator.TranslateFile("./README.md", ".", "README_T.md", true, language.English, language.SimplifiedChinese)
	if err != nil {
		panic(err)
	}

}
