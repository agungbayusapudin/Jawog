# Jawog - Javanese Script Translator

<a href="https://ibb.co.com/Z6V3D131"><img src="https://i.ibb.co.com/Z6V3D131/jawog-logo.png" alt="jawog-logo" border="0"></a>

A Go package for translating between Latin text and Javanese script (Aksara Jawa).

## Features

- Bidirectional translation between Latin and Javanese script
- Configurable translation modes
- Support for special characters and combinations
- Clean output with invisible character removal

## Installation

```bash
go mod init your-project
go get github.com/yourusername/jawog
```

## Usage

```go
package main

import (
    "fmt"
    translate "translate/translator"
)

func main() {
    // Latin to Javanese
    latinText := "Menerjemahkan teks Latin ke Aksara Jawa"
    config := translate.DefaultConfig
    javaneseText := translate.Translate(latinText, &config)
    fmt.Println("Javanese:", javaneseText)
    
    // Javanese to Latin
    javaneseText = "ꦧꦲꦱ"
    latinText = translate.Translate(javaneseText, &config)
    fmt.Println("Latin:", latinText)
}
```

## Configuration

The translator supports various configuration options through `TranslateConfiguration`:

- `TypeMode`: Enable/disable type mode
- `WithSpace`: Include spaces in translation
- `WithMurda`: Use Murda characters

## License

MIT License - see [LICENSE](LICENSE) file for details.
