package main

import (
	"fmt"

	"go.amplifyedge.org/shared-v2/tool/bs-lang-srv/x/bregydoc/gtranslate"
)

func main() {
	text := "Hello World"
	translated, err := gtranslate.TranslateWithParams(
		text,
		gtranslate.TranslationParams{
			From: "en",
			To:   "ja",
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("en: %s | ja: %s \n", text, translated)
	// en: Hello World | ja: こんにちは世界
}
