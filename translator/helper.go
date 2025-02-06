package translate

import (
	"strings"
	"unicode"
)

// getMatra returns the matra (sandhangan swara) for a given string
func getMatra(str string, config TranslateConfiguration) string {
	if len(str) < 1 {
		return "꧀"
	}

	i := 0
	for str[i] == 'h' {
		i++
		if i >= len(str) {
			break
		}
	}

	if i < len(str) {
		str = str[i:]
	}

	var matraMap map[string]string
	if config.TypeMode {
		matraMap = matraMap1
	} else {
		matraMap = matraMap2
	}

	if val, ok := matraMap[str]; ok {
		return val
	}

	return ""
}

// getCoreSound returns the core sound of a character
func getCoreSound(str string, config TranslateConfiguration) (string, int) {
	str = strings.ToLower(str)

	var consonantMap map[rune]string
	if config.WithMurda {
		consonantMap = consonantMap2
	} else {
		consonantMap = consonantMap1
	}

	if val, ok := consonantMap[rune(str[0])]; ok {
		return val, 1
	}

	return "", 0
}

// getSound translates a syllable to Javanese script
func getSound(str string, config TranslateConfiguration, vowelPrev bool) string {
	str = strings.TrimSpace(str)

	if str == "" {
		return ""
	}

	if specialSound := getSpecialSound(rune(str[0])); specialSound != "" {
		return specialSound
	}

	if len(str) == 1 {
		return resolveCharacterSound(rune(str[0]), config)
	}

	coreSound, coreLen := getCoreSound(str, config)
	konsonan := ""
	matra := ""

	if coreLen >= 1 {
		matra = getMatra(str[coreLen:], config)
	}

	// Handle special cases for consonant clusters
	switch {
	case strings.HasPrefix(str, "nggr"):
		if vowelPrev {
			konsonan = "ꦁꦒꦿ" // <vowel>nggr-, e.g. panggrahita
		} else {
			konsonan = "ꦔ꧀ꦒꦿ" // <nonvowel>nggr-, i.e. nggronjal
		}
	case strings.HasPrefix(str, "nggl"):
		konsonan = "ꦔ꧀ꦒ꧀ꦭ"
	case strings.HasPrefix(str, "nggw"):
		konsonan = "ꦔ꧀ꦒ꧀ꦮ"
	case strings.HasPrefix(str, "nggy"):
		konsonan = "ꦔ꧀ꦒꦾ"
	case strings.HasPrefix(str, "ngg"):
		if vowelPrev {
			konsonan = "ꦁꦒ" // <vowel>ngg-, e.g. tunggal
		} else {
			konsonan = "ꦔ꧀ꦒ" // <nonvowel>ngg-, i.e. nggambar
		}
	case strings.HasPrefix(str, "ngl"):
		konsonan = "ꦔ꧀ꦭ"
	case strings.HasPrefix(str, "ngw"):
		konsonan = "ꦔ꧀ꦮ"
	case strings.HasPrefix(str, "ncl"):
		konsonan = "ꦚ꧀ꦕ꧀ꦭ"
	case strings.HasPrefix(str, "ncr"):
		konsonan = "ꦚ꧀ꦕꦿ"
	case strings.HasPrefix(str, "njl"):
		konsonan = "ꦚ꧀ꦗ꧀ꦭ"
	case strings.HasPrefix(str, "njr"):
		konsonan = "ꦚ꧀ꦗꦿ"
	default:
		konsonan = coreSound
	}

	return konsonan + matra
}

// resolveCharacterSound resolves the core sound of a character
func resolveCharacterSound(c rune, config TranslateConfiguration) string {
	if unicode.IsDigit(c) {
		return "꧇" + string(c)
	} else if isHR(c) {
		coreSound, _ := getCoreSound(string(c), config)
		return coreSound
	} else if isCJ(c) {
		coreSound, _ := getCoreSound(string(c), config)
		return coreSound + "꧀"
	} else if isConsonant(c) {
		coreSound, _ := getCoreSound(string(c), config)
		return coreSound + "꧀"
	} else {
		coreSound, _ := getCoreSound(string(c), config)
		return coreSound
	}
}

// Helper function to convert Latin numbers to Javanese numerals
func convertNumberToJavanese(number string) string {
	var result strings.Builder
	for _, digit := range number {
		if javaneseNumeral, exists := javaneseNumerals[digit]; exists {
			result.WriteString(javaneseNumeral)
		}
	}
	return result.String()
}

// Helper function to handle number sequences
func handleNumbers(str string, start int) (string, int) {
	var numStr strings.Builder
	i := start

	// Collect all consecutive digits
	for i < len(str) && isDigit(rune(str[i])) {
		numStr.WriteByte(str[i])
		i++
	}

	if numStr.Len() > 0 {
		return convertNumberToJavanese(numStr.String()), i - 1
	}
	return "", start
}

// isDigit checks if a character is a digit
func isDigit(c rune) bool {
	return unicode.IsDigit(c)
}

// isVowel checks if a character is a vowel
func isVowel(c rune) bool {
	vowels := "aeiouāēīōū"
	return strings.ContainsRune(vowels, c)
}

// isConsonant checks if a character is a consonant
func isConsonant(c rune) bool {
	consonants := "bcdfghjklmnpqrstvwxyz"
	return strings.ContainsRune(consonants, c)
}

// isHR checks if a character is a layar or wignyan
func isHR(c rune) bool {
	return c == 'ꦃ' || c == 'ꦂ'
}

// isCJ checks if a character is a cakra or jang
func isCJ(c rune) bool {
	return c == 'ꦿ' || c == 'ꦾ'
}

// getSpecialSound returns the special sound of a character
func getSpecialSound(c rune) string {
	if val, ok := specialSoundMap[c]; ok {
		return val
	}
	return ""
}

// hasAksara checks if a text contains Javanese Script
func HasAksara(text string) bool {
	for _, char := range text {
		if _, ok := javaneseToLatin[char]; ok {
			return true
		}
	}
	return false
}

// Helper function to check if a rune is at the start of a word
func isWordStart(runes []rune, pos int) bool {
	if pos == 0 {
		return true
	}
	return runes[pos-1] == ' '
}
