// language-filter filters a text file, including only lines with a percentage of words in the given language.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/BluntSporks/lexicon"
)

func main() {
	// Parse flags.
	lexDir := flag.String("lexdir", lex.DefaultDataPath(), "Location of lexicon data directory")
	langFile := flag.String("lang", "english", "Name of language file to use for filtering")
	percentMin := flag.Int("percent", 75, "Minimum percentage of language to require to be included")
	flag.Parse()

	// Check flags.
	if len(flag.Args()) < 1 {
		log.Fatal("Missing text filename parameter")
	}
	textFile := flag.Arg(0)

	// Load the language.
	path := path.Join(*lexDir, *langFile)
	wordList := lex.LoadLang(path)

	// Open file.
	hdl, err := os.Open(textFile)
	if err != nil {
		log.Fatal(err)
	}
	defer hdl.Close()

	// Match words of at least two in length.
	wordRegExp := regexp.MustCompile(`\pL{2,}`)

	// Scan file line by line.
	scanner := bufio.NewScanner(hdl)
	lastBlank := false
	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		// Process line
		lowLine := strings.TrimRight(line, " \t")
		if len(lowLine) > 0 {
			lowLine = strings.ToLower(line)
			matches := wordRegExp.FindAllString(line, -1)
			total := len(matches)
			if total == 0 {
				continue
			}
			known := 0
			for _, match := range matches {
				if wordList[match] {
					known++
				}
			}
			percent := int(100.0 * float32(known) / float32(total))
			if percent > *percentMin {
				fmt.Println(line)
				lastBlank = false
			}
		} else if !lastBlank {
			fmt.Println()
			lastBlank = true
		}
	}
}
