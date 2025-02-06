package main

import (
	"fmt"
	translate "translate/translator"
)

func main() {
	// Example Latin to Javanese Script translation
	latinText := "saya sangat suka makanan yang berkuah"
	// config := translate.TranslateConfiguration{
	// 	TypeMode:  true,
	// 	WithSpace: true,
	// 	WithMurda: false,
	// }
	config := translate.DefaultConfig

	javaneseText := translate.Translate(latinText, &config)
	fmt.Println("Latin to Javanese:", javaneseText)

	// Example Javanese Script to Latin translation
	javaneseText = "ꦱꦪ ꦱꦤ꧀ꦒꦠ꧀ ꦱꦸꦏ ꦩꦏꦤꦤ꧀ ꦪꦤ꧀ꦒ꧀ ꦧꦺꦂ꧀ꦏꦸꦮꦃ꧀"
	latinText = translate.TranslateAksara(javaneseText)
	fmt.Println("Javanese to Latin:", latinText)
}
