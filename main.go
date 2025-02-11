package main

import (
	"fmt"
	"strings"
	translate "translate/translator"
)

// RemoveInvisibleCharacters menghapus karakter tidak terlihat dari string
func RemoveInvisibleCharacters(input string) string {
	// Menghapus karakter ZWSP (U+200B)
	return strings.ReplaceAll(input, "\u200B", "")
}

func main() {
	// Example Latin to Javanese Script translation
	latinText := "Menerjemahkan teks Latin ke Aksara Jawa bisa menjadi tantangan, terutama ketika ada kata-kata krusial yang sulit untuk diterjemahkan secara langsung. Kata-kata tersebut biasanya adalah kata serapan dari bahasa asing, kata yang mengandung kombinasi huruf khusus, atau kata yang memiliki aturan penulisan khusus dalam Aksara Jawa. Berikut adalah deskripsi langkah-langkah yang dapat dilakukan untuk menangani kata-kata krusial tersebut:"
	// config := translate.TranslateConfiguration{
	// 	TypeMode:  true,
	// 	WithSpace: true,
	// 	WithMurda: false,
	// }
	config := translate.DefaultConfig

	javaneseText := translate.Translate(latinText, &config)
	cleaned := RemoveInvisibleCharacters(javaneseText)
	fmt.Println("Latin to Javanese:", cleaned)

	// Example Javanese Script to Latin translation
	javaneseText = "ꦧꦲꦱ​"
	latinText = translate.Translate(javaneseText, &config)
	fmt.Println("Javanese to Latin:", latinText)
}
