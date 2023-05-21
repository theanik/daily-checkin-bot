package main

import (
	// "encoding/json"
    "io/ioutil"
    "log"
	"fmt"
	"sort"

	"github.com/tidwall/gjson"
	"github.com/theanik/goapp/app/utils"
)

type XY struct {
	tokens []string
	tag string
}

func main() {
    content, err := ioutil.ReadFile("./traning_data.json")
    if err != nil {
        log.Fatal("Error when opening file: ", err)
    }

	intents := gjson.Get(string(content), "intents")
	intentsVal := intents.Array()
	intentsLen := len(intentsVal)

	var tags []string
	var allUnsanitizeWords []string
	var allSanitizeWords []string
	var xyBag []XY
	
	var patternStr string
	for i := 0;i < intentsLen;i+=1 {
		tag := gjson.Get(intentsVal[i].String(), "tag")
		patterns := gjson.Get(intentsVal[i].String(), "patterns")

		tags = append(tags, tag.String())

		for _, val := range patterns.Array() {
			patternStr = val.String()
			tokens := utils.Tokenizer(patternStr)
			xy := XY {
				tokens: tokens,
				tag: tag.String(),
			}
			xyBag = append(xyBag, xy)
			allUnsanitizeWords = append(allUnsanitizeWords, tokens...)
		}

	}

	for _, word := range allUnsanitizeWords {
		if isSpecialChars(word) == false {
			stemWord := utils.Stemmer(word)
			allSanitizeWords = append(allSanitizeWords, stemWord)
		}
		
	}

	sort.Strings(allSanitizeWords)
	sort.Strings(tags)

	allSanitizeWords = removeDuplicate(allSanitizeWords)
	tags = removeDuplicate(tags)

	var XTrain []map[int]float32

	var YTrain []int

	for _, value := range xyBag {
		
		bag := utils.BagOfWord(value.tokens, allSanitizeWords)
		XTrain = append(XTrain, bag)

		label := indexOf(value.tag, tags)
		YTrain = append(YTrain, label)
	}

	fmt.Println(XTrain)
	fmt.Println(YTrain)

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

func removeDuplicate[T string | int](sliceList []T) []T {
    allKeys := make(map[T]bool)
    list := []T{}
    for _, item := range sliceList {
        if _, value := allKeys[item]; !value {
            allKeys[item] = true
            list = append(list, item)
        }
    }
    return list
}

func indexOf(element string, data []string) (int) {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
 }