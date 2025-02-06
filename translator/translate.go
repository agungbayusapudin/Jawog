package translate

import (
	"strings"
	// "unicode"
)

// Translate translates Latin to Javanese Script
func Translate(latin string, config *TranslateConfiguration) string {
	if HasAksara(latin) {
		return TranslateAksara(latin)
	}

	if config == nil {
		config = &DefaultConfig
	}

	i := 0
	pi := 0 // offset
	vowelFlag := false
	angkaFlag := false
	vowelPrev := false
	str := latin
	var ret strings.Builder

	for i < len(str) {
		if i > 0 && isVowel(rune(str[i])) && isVowel(rune(str[i-1])) {
			// Handle multiple vowels at the start of a word
			if (str[i-1] == 'a' && str[i] == 'a') ||
				(str[i-1] == 'i' && str[i] == 'i') ||
				(str[i-1] == 'u' && str[i] == 'u') ||
				(str[i-1] == 'a' && str[i] == 'i') ||
				(str[i-1] == 'a' && str[i] == 'u') {
				if i > 1 && !isConsonant(rune(str[i-2])) {
					str = str[:i] + "h" + str[i:]
				}
			} else if (str[i-1] == 'e' || str[i-1] == 'è' || str[i-1] == 'é') &&
				(str[i] == 'a' || str[i] == 'o') {
				str = str[:i] + "y" + str[i:]
			} else if str[i-1] == 'i' &&
				(str[i] == 'a' || str[i] == 'e' || str[i] == 'è' || str[i] == 'é' || str[i] == 'o' || str[i] == 'u') {
				str = str[:i] + "y" + str[i:]
			} else if str[i-1] == 'o' &&
				(str[i] == 'a' || str[i] == 'e' || str[i] == 'è' || str[i] == 'é') {
				str = str[:i] + "w" + str[i:]
			} else if str[i-1] == 'u' &&
				(str[i] == 'a' || str[i] == 'e' || str[i] == 'è' || str[i] == 'é' || str[i] == 'i' || str[i] == 'o') {
				str = str[:i] + "w" + str[i:]
			} else {
				str = str[:i] + "h" + str[i:]
			}
		}

		// Handle number sequences
		if isDigit(rune(str[i])) {
			if !angkaFlag {
				ret.WriteString("꧇") // Add number indicator
				angkaFlag = true
			}
			if numStr, newIndex := handleNumbers(str, i); numStr != "" {
				ret.WriteString(numStr)
				i = newIndex
			}
		} else {
			if angkaFlag {
				ret.WriteString("꧇") // Close number indicator
				angkaFlag = false
			}

			if (rune(str[i]) == 'h' && vowelFlag) ||
				(!isVowel(rune(str[i])) && i > 0) ||
				rune(str[i]) == ' ' ||
				isPunct(rune(str[i])) ||
				i-pi > 5 {

				if pi < i {
					ret.WriteString(getSound(str[pi:i], *config, vowelPrev))
				}

				if rune(str[i]) == ' ' && config.WithSpace {
					ret.WriteRune(' ')
				}

				if isPunct(rune(str[i])) {
					switch str[i] {
					case '.':
						ret.WriteString("꧉") // titik
					case ',':
						ret.WriteString("꧈") // koma
					case '|':
						ret.WriteString("꧋")
					case '(':
						ret.WriteString("꧌")
					case ')':
						ret.WriteString("꧍") // with zws
					case '-':
						// tanda hubung (empty)
					case '?', '!', '"', '\'':
						// tanda tanya/seru/petik (empty)
					default:
						ret.WriteRune(rune(str[i]))
					}
					pi = i + 1
				} else {
					pi = i
				}
				vowelFlag = false
			} else if isVowel(rune(str[i])) && rune(str[i]) != 'h' {
				vowelFlag = true
			}
		}

		if pi > 0 && isVowel(rune(str[pi-1])) {
			vowelPrev = true
		} else {
			vowelPrev = false
		}
		i++
	}

	if pi < i {
		ret.WriteString(getSound(str[pi:i], *config, vowelPrev))
	}

	return strings.TrimSpace(ret.String())
}

func TranslateAksara(aksara string) string {
	var ret strings.Builder
	runes := []rune(aksara)

	for i := 0; i < len(runes); i++ {
		// Look ahead for taling (ꦺ) combinations
		if i < len(runes)-1 && runes[i] == 'ꦺ' {
			// Handle taling + next character
			nextChar := runes[i+1]
			if translation, exists := javaneseToLatin[nextChar]; exists {
				ret.WriteString("e")
				if strings.HasSuffix(translation, "a") {
					ret.WriteString(strings.TrimSuffix(translation, "a"))
				} else {
					ret.WriteString(translation)
				}
				i++ // Skip the next character
				continue
			}
		}

		// Check for special combinations
		if i < len(runes)-2 {
			combo := string(runes[i]) + string(runes[i+1]) + string(runes[i+2])
			if translation, exists := specialCombinations[combo]; exists {
				ret.WriteString(translation)
				i += 2
				continue
			}
		}

		char := runes[i]

		// Skip combining marks we've already handled
		if char == 'ꦶ' || char == 'ꦸ' || char == '꧀' || char == 'ꦺ' || char == 'ꦴ' {
			continue
		}

		// Get the next character if available
		var nextChar rune
		if i < len(runes)-1 {
			nextChar = runes[i+1]
		}

		// Special handling for ꦲ (ha) at word start
		if char == 'ꦲ' && isWordStart(runes, i) {
			if nextChar == 'ꦶ' {
				ret.WriteRune('i')
				i++
			} else if nextChar == 'ꦸ' {
				ret.WriteRune('u')
				i++
			} else {
				ret.WriteRune('a')
			}
			continue
		}

		// Handle base character translation
		if translation, exists := javaneseToLatin[char]; exists {
			if nextChar == 'ꦶ' {
				ret.WriteString(strings.ReplaceAll(translation, "a", "i"))
				i++
			} else if nextChar == 'ꦸ' {
				ret.WriteString(strings.ReplaceAll(translation, "a", "u"))
				i++
			} else if nextChar == '꧀' {
				ret.WriteString(strings.TrimSuffix(translation, "a"))
				i++
			} else if nextChar == 'ꦺ' {
				ret.WriteString(strings.ReplaceAll(translation, "a", "e"))
				i++
			} else {
				ret.WriteString(translation)
			}
		}

		// Handle spaces
		if char == ' ' {
			ret.WriteRune(' ')
		}
	}

	// Post-processing to fix common issues
	result := strings.TrimSpace(ret.String())
	result = strings.ReplaceAll(result, "ae", "e") // Fix double vowel issues
	result = strings.ReplaceAll(result, "iy", "i") // Fix common combinations

	return result
}

// isPunct checks if a character is a punctuation mark
func isPunct(c rune) bool {
	punctuations := ".,!?;:'\"()-"
	return strings.ContainsRune(punctuations, c)
}
