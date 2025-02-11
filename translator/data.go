package translate

// TranslateConfiguration holds the configuration for translation
type TranslateConfiguration struct {
	TypeMode  bool
	WithSpace bool
	WithMurda bool
}

// DefaultConfig is the default configuration for translation
var DefaultConfig = TranslateConfiguration{
	TypeMode:  true,
	WithSpace: true,
	WithMurda: false,
}

// Peta transliterasi aksara Jawa ke Latin
var javaneseToLatin = map[rune]string{
	'ꦀ': "",   // archaic
	'ꦁ': "ng", // cecak
	'ꦂ': "r",  // layar
	'ꦃ': "h",  // wignyan
	'ꦄ': "a",  // swara-A
	'ꦅ': "i",  // I-Kawi -- archaic
	'ꦆ': "i",  // I
	'ꦈ': "u",  // U
	'ꦉ': "re", // pa cerek
	'ꦊ': "le", // nga lelet
	'ꦌ': "e",  // E
	'ꦍ': "ai", // Ai
	'ꦎ': "o",  // O

	'ꦏ': "ka",
	'ꦐ': "qa",  // Ka Sasak
	'ꦑ': "kha", // Murda
	'ꦒ': "ga",
	'ꦓ': "gha", // Murda
	'ꦔ': "nga", // ṅa
	'ꦕ': "ca",
	'ꦖ': "cha", // Murda
	'ꦗ': "ja",
	'ꦘ': "nya",  // Ja Sasak, Nya Murda
	'ꦙ': "jha",  // Ja Mahaprana
	'ꦚ': "nya",  // ña
	'ꦛ': "tha",  // ṭa
	'ꦜ': "thha", // Murda
	'ꦝ': "dha",  // ḍa
	'ꦞ': "dhha", // Murda
	'ꦠ': "ta",
	'ꦡ': "ṭa", // Murda
	'ꦢ': "da",
	'ꦣ': "ḍa", // Murda
	'ꦤ': "na",
	'ꦥ': "pa",
	'ꦦ': "pha", // Murda
	'ꦧ': "ba",
	'ꦨ': "bha", // Murda
	'ꦩ': "ma",
	'ꦪ': "ya",
	'ꦫ': "ra",
	'ꦬ': "ra", // Ra Agung
	'ꦭ': "la",
	'ꦮ': "wa",
	'ꦯ': "sha", // Murda
	'ꦰ': "ṣa",  // Sa Mahaprana
	'ꦱ': "sa",
	'ꦲ': "ha", // could also be "a" or any sandhangan swara

	'꦳': "",   // cecak telu -- diganti zero-width joiner (tmp)
	'ꦴ': "a",  // tarung
	'ꦶ': "i",  // suku
	'ꦷ': "ii", // suku panjang
	'ꦸ': "u",  // taling
	'ꦹ': "uu", // taling panjang
	'ꦺ': "e",  // taling
	'ꦻ': "ai", // taling tarung
	'ꦼ': "ê",  // pepet
	'ꦽ': "re", // pa cerek
	'ꦾ': "ya", // pengkal
	'ꦿ': "ra", // cakra

	'꧀': "", // pangkon -- menghapus suara vokal sebelumnya
	'꧁': "—",
	'꧂': "—",
	'꧃': "–",
	'꧄': "–",
	'꧅': "–",
	'꧆': "",
	'꧇': "",
	'꧈': ",",
	'꧉': ".",
	'꧐': "0",
	'꧑': "1",
	'꧒': "2",
	'꧓': "3",
	'꧔': "4",
	'꧕': "5",
	'꧖': "6",
	'꧗': "7",
	'꧘': "8",
	'꧙': "9",
	'​': " ", // zero-width space
}

// Map for Javanese numerals
var javaneseNumerals = map[rune]string{
	'0': "꧐",
	'1': "꧑",
	'2': "꧒",
	'3': "꧓",
	'4': "꧔",
	'5': "꧕",
	'6': "꧖",
	'7': "꧗",
	'8': "꧘",
	'9': "꧙",
}

var specialCombinations = map[string]string{
	// Basic consonant combinations
	"ꦒ꧀ꦒ": "g",  // Handle double consonants (gagga)
	"ꦤ꧀ꦒ": "ng", // Handle ng sound (nga)
	"ꦢ꧀ꦤ": "n",  // Handle dn combination
	"ꦱ꧀ꦠ": "st", // Handle st combination
	"ꦂ꧀ꦱ": "rs", // Handle rs combination
	"ꦏ꧀ꦤ": "kn", // Handle kn combination

	// Taling combinations
	"ꦺꦴ": "o",  // Handle taling tarung
	"ꦺꦤ": "en", // Handle taling + na
	"ꦺꦏ": "ek", // Handle taling + ka
	"ꦺꦭ": "el", // Handle taling + la
	"ꦺꦂ": "er", // Handle taling + layar
	"ꦺꦃ": "eh", // Handle taling + wignyan

	// Common consonant clusters
	"ꦤ꧀ꦝ": "ndh", // Handle ndh sound
	"ꦤ꧀ꦠ": "nt",  // Handle nt combination
	"ꦤ꧀ꦢ": "nd",  // Handle nd combination
	"ꦤ꧀ꦗ": "nj",  // Handle nj combination
	"ꦤ꧀ꦕ": "nc",  // Handle nc combination
	"ꦩ꧀ꦧ": "mb",  // Handle mb combination
	"ꦩ꧀ꦥ": "mp",  // Handle mp combination
	"ꦚ꧀ꦕ": "nc",  // Handle nyc combination
	"ꦚ꧀ꦗ": "nj",  // Handle nyj combination
	"ꦔ꧀ꦒ": "ngg", // Handle ngg combination
	"ꦔ꧀ꦏ": "ngk", // Handle ngk combination

	// Retroflex combinations
	"ꦠ꧀ꦫ": "tr", // Handle tr combination
	"ꦢ꧀ꦫ": "dr", // Handle dr combination
	"ꦱ꧀ꦫ": "sr", // Handle sr combination

	// Special sounds
	"ꦏ꧀ꦱ":   "ksa", // Handle ksa combination
	"ꦥ꧀ꦫ":   "pra", // Handle pra combination
	"ꦧ꧀ꦫ":   "bra", // Handle bra combination
	"ꦩ꧀ꦥ꧀ꦭ": "mpl", // Handle mpl combination
	"ꦩ꧀ꦧ꧀ꦭ": "mbl", // Handle mbl combination
	"ꦤ꧀ꦠ꧀ꦫ": "ntr", // Handle ntr combination
	"ꦤ꧀ꦢ꧀ꦫ": "ndr", // Handle ndr combination

	// Sandhangan combinations
	"ꦲꦶ":  "i", // Handle i sound at start of word
	"ꦲꦸ":  "u", // Handle u sound at start of word
	"ꦲꦺ":  "e", // Handle e sound at start of word
	"ꦲꦺꦴ": "o", // Handle o sound at start of word

	// Common word endings
	"ꦁ": "ng", // Cecak (ng ending)
	"ꦃ": "h",  // Wignyan (h ending)
	"ꦂ": "r",  // Layar (r ending)

	// Additional complex combinations
	"ꦱ꧀ꦠ꧀ꦫ": "stra", // Handle stra combination
	"ꦤ꧀ꦢ꧀ꦮ": "ndw",  // Handle ndw combination
	"ꦱ꧀ꦥ꧀ꦭ": "spl",  // Handle spl combination
	"ꦱ꧀ꦏ꧀ꦫ": "skr",  // Handle skr combination

	// Modern additions for foreign sounds
	"ꦲ꦳": "fa",  // Handle fa sound
	"ꦗ꦳": "za",  // Handle za sound
	"ꦏ꦳": "kha", // Handle kha sound
	"ꦱ꦳": "sya", // Handle sya sound

	// Special vowel combinations
	"ꦺꦵ": "ai", // Handle ai diphthong
	"ꦺꦷ": "au", // Handle au diphthong

	// Double consonant combinations
	"ꦏ꧀ꦏ": "kk", // Handle kk combination
	"ꦥ꧀ꦥ": "pp", // Handle pp combination
	"ꦠ꧀ꦠ": "tt", // Handle tt combination
	"ꦱ꧀ꦱ": "ss", // Handle ss combination

	// Additional modern combinations
	"ꦝ꧀ꦮ": "dhw", // Handle dhw combination
	"ꦠ꧀ꦮ": "tw",  // Handle tw combination
	"ꦱ꧀ꦮ": "sw",  // Handle sw combination
	"ꦏ꧀ꦮ": "kw",  // Handle kw combination
	"ꦢ꧀ꦮ": "dw",  // Handle dw combination

	// Basic consonant combinations
	"ꦭꦺꦴ":   "lo",  // lo with taling tarung
	"ꦤꦺꦴ":   "no",  // no with taling tarung
	"ꦒꦺꦴ":   "go",  // go with taling tarung
	"ꦪꦺꦴ":   "yo",  // yo with taling tarung
	"ꦮ꦳ꀀꦱꦶ": "ver", // ver (for universitas)
	"ꦤꦶ":    "ni",  // ni
}

// Matra maps vowels to their corresponding Javanese script
var matraMap1 = map[string]string{
	"ā":  "ꦴ",
	"â":  "ꦴ",
	"e":  "ꦺ",
	"è":  "ꦺ",
	"é":  "ꦺ",
	"i":  "ꦶ",
	"ī":  "ꦷ",
	"o":  "ꦺꦴ",
	"u":  "ꦸ",
	"ū":  "ꦹ",
	"x":  "ꦼ",
	"ě":  "ꦼ",
	"ĕ":  "ꦼ",
	"ê":  "ꦼ",
	"ō":  "ꦼꦴ",
	"ô":  "",
	"A":  "ꦄ",
	"E":  "ꦌ",
	"È":  "ꦌ",
	"É":  "ꦌ",
	"I":  "ꦆ",
	"U":  "ꦈ",
	"O":  "ꦎ",
	"X":  "ꦄꦼ",
	"Ě":  "ꦄꦼ",
	"Ê":  "ꦄꦼ",
	"ṛ":  "ꦽ",
	"aa": "ꦴ",
	"ai": "ꦻ",
	"au": "ꦻꦴ",
	"ii": "ꦷ",
	"uu": "ꦹ",
}

var matraMap2 = map[string]string{
	"ā":  "ꦴ",
	"â":  "ꦴ",
	"e":  "ꦼ",
	"è":  "ꦺ",
	"é":  "ꦺ",
	"i":  "ꦶ",
	"ī":  "ꦷ",
	"u":  "ꦸ",
	"ū":  "ꦹ",
	"o":  "ꦺꦴ",
	"x":  "ꦼ",
	"ě":  "ꦼ",
	"ĕ":  "ꦼ",
	"ê":  "ꦼ",
	"ô":  "",
	"ō":  "ꦼꦴ",
	"A":  "ꦄ",
	"E":  "ꦄꦼ",
	"È":  "ꦌ",
	"É":  "ꦌ",
	"I":  "ꦆ",
	"U":  "ꦈ",
	"O":  "ꦎ",
	"X":  "ꦄꦼ",
	"Ě":  "ꦄꦼ",
	"Ê":  "ꦄꦼ",
	"ṛ":  "ꦽ",
	"aa": "ꦴ",
	"ai": "ꦻ",
	"au": "ꦻꦴ",
	"ii": "ꦷ",
	"uu": "ꦹ",
}

var consonantMap1 = map[rune]string{
	'A': "ꦄ",   // A
	'B': "ꦧ",   // B
	'C': "ꦕ",   // C
	'D': "ꦢ",   // D
	'E': "ꦌ",   // E
	'F': "ꦥ꦳",  // F
	'G': "ꦒ",   // G
	'H': "ꦲ",   // H
	'I': "ꦆ",   // I
	'J': "ꦗ",   // J
	'K': "ꦏ",   // K
	'L': "ꦭ",   // L
	'M': "ꦩ",   // M
	'N': "ꦤ",   // N
	'O': "ꦎ",   // O
	'P': "ꦥ",   // P
	'Q': "꧀",   // Q
	'R': "ꦂ",   // R
	'S': "ꦱ",   // S
	'T': "ꦠ",   // T
	'U': "ꦈ",   // U
	'V': "ꦮ꦳",  // V
	'W': "ꦮ",   // W
	'X': "ꦼ",   // X
	'Y': "ꦪ",   // Y
	'Z': "ꦗ꦳",  // Z
	'a': "ꦄ",   // a
	'b': "ꦧ",   // b
	'c': "ꦕ",   // c
	'd': "ꦢ",   // d
	'e': "ꦌ",   // e
	'f': "ꦥ꦳",  // f
	'g': "ꦒ",   // g
	'h': "ꦲ",   // h
	'i': "ꦲꦶ",  // i
	'j': "ꦗ",   // j
	'k': "ꦏ",   // k
	'l': "ꦭ",   // l
	'm': "ꦩ",   // m
	'n': "ꦤ",   // n
	'o': "ꦲꦺꦴ", // o
	'p': "ꦥ",   // p
	'q': "꧀",   // q
	'r': "",    // r
	's': "ꦱ",   // s
	't': "ꦠ",   // t
	'u': "ꦈ",   // u
	'v': "ꦮ꦳",  // v
	'w': "ꦮ",   // w
	'x': "ꦼ",   // x
	'y': "ꦪ",   // y
	'z': "ꦗ꦳",  // z
	'È': "ꦌ",   // È
	'É': "ꦌ",   // É
	'Ê': "ꦄꦼ",  // Ê
	'Ě': "ꦄꦼ",  // Ě
	'è': "ꦌ",   // è
	'é': "ꦌ",   // é
	'ê': "ꦄꦼ",  // ê
	'ě': "ꦄꦼ",  // ě
	'ô': "ꦲ",   // ô
	'ñ': "ꦚ",   // ñ
	'ṇ': "nell",
	'ḍ': "ꦝ",
	'ṭ': "ꦛ",
	'ṣ': "ꦰ",
	'ṛ': "ꦽ",
}

var consonantMap2 = map[rune]string{
	'A': "ꦄ",    // A
	'B': "ꦨ",    // B
	'C': "ꦖ",    // C
	'D': "ꦣ",    // D
	'E': "ꦌ",    // E
	'F': "ꦦ꦳",   // F
	'G': "ꦓ",    // G
	'H': "ꦲ꦳",   // H
	'I': "ꦆ",    // I
	'J': "ꦙ",    // J
	'K': "ꦑ",    // K
	'L': "ꦭ",    // L
	'M': "ꦩ",    // M
	'N': "nell", // N
	'O': "ꦎ",    // O
	'P': "ꦦ",    // P
	'Q': "꧀",    // Q
	'R': "ꦬ",    // R
	'ś': "ꦯ",    // ś
	'S': "ꦯ",    // S
	'T': "ꦡ",    // T
	'U': "ꦈ",    // U
	'V': "ꦮ꦳",   // V
	'W': "ꦮ",    // W
	'X': "ꦼ",    // X
	'Y': "ꦪ",    // Y
	'Z': "ꦗ꦳",   // Z
	'a': "ꦄ",    // a
	'b': "ꦧ",    // b
	'c': "ꦕ",    // c
	'd': "ꦢ",    // d
	'e': "ꦌ",    // e
	'f': "ꦥ꦳",   // f
	'g': "ꦒ",    // g
	'h': "ꦃ",    // h
	'i': "ꦲꦶ",   // i
	'j': "ꦗ",    // j
	'k': "ꦏ",    // k
	'l': "ꦭ",    // l
	'm': "ꦩ",    // m
	'n': "nell", // n
	'o': "ꦎ",    // o
	'p': "ꦥ",    // p
	'q': "꧀",    // q
	'r': "ꦂ",    // r
	's': "ꦱ",    // s
	't': "ꦠ",    // t
	'u': "ꦈ",    // u
	'v': "ꦮ꦳",   // v
	'w': "ꦮ",    // w
	'x': "ꦼ",    // x
	'ĕ': "ꦼ",    // ě
	'ě': "ꦼ",    // ě
	'ê': "ꦼ",    // ê
	'ū': "ꦹ",    // ū
	'y': "ꦪ",    // y
	'z': "ꦗ꦳",   // z
	'È': "ꦌ",    // È
	'É': "ꦌ",    // É
	'Ê': "ꦄꦼ",   // Ê
	'Ě': "ꦄꦼ",   // Ě
	'è': "ꦌ",    // è   // é   // ê,   // ě
	'ô': "ꦲ",    // ô
	'ñ': "ꦚ",    // ñ
	'ṇ': "nell",
	'ḍ': "ꦝ",
	'ṭ': "ꦛ",
	'ṣ': "ꦰ",
	'ṛ': "ꦽ",
}

// SpecialSound maps special sounds to their corresponding Javanese script
var specialSoundMap = map[rune]string{
	'f': "ꦥ꦳꧀",
	'v': "ꦮ꦳꧀",
	'z': "ꦗ꦳꧀",
	'ś': "ꦯ",
	'q': "꧀", // pangkon
}
