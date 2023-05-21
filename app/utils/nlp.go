package utils

import (
	"log"
	"strings"

    "github.com/jdkato/prose/v2"
	"github.com/surgebase/porter2"
)


func Tokenizer(text string) []string {
	var tokens []string
	doc, err := prose.NewDocument(text)
	if err != nil {
        log.Fatal(err)
    }
	for _, tok := range doc.Tokens() {
        tokens = append(tokens, strings.ToLower(tok.Text))
    }
	return tokens
}

func Stemmer(text string) string {
	return porter2.Stem(text) // orange, organization, 
}

func BagOfWord(sentanceTokens []string, allWords []string) map[int]float32 {
	var stemSentanceTokens []string
	bag := make(map[int]float32)

	for _, word := range sentanceTokens {
		if isSpecialChars(word) == false {
			stemWord := Stemmer(word)
			stemSentanceTokens = append(stemSentanceTokens, stemWord)
		}
	}

	for idx, text := range allWords {
		if inArrayStr(stemSentanceTokens, text) {
			bag[idx] = 1.0
		} else {
			bag[idx] = 0.0
		}
	}

	return bag
}

func isSpecialChars(str string) bool {
	s := []string{".", "?", ",", "!"}
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func inArrayStr(strArr []string, str string) bool {
	for _, v := range strArr {
		if v == str {
			return true
		}
	}

	return false
}