package main

import (
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
)

func selectList(title string, items []string) (int, error) {
	prompt := promptui.Select{
		Label:        title,
		Items:        items,
		Size:         15,
		CursorPos:    0,
		IsVimMode:    false,
		HideHelp:     false,
		HideSelected: false,
		Templates:    nil,
		Keys:         nil,
		Searcher: func(input string, index int) bool {
			return inText(strings.ToLower(input), strings.ToLower(items[index]))
		},
		StartInSearchMode: true,
		Pointer:           nil,
		Stdin:             nil,
		Stdout:            nil,
	}
	idx, result, err := prompt.Run()
	if err != nil {
		return 0, err
	}
	logrus.Infof("select %d, %s", idx, result)

	return idx, nil
}

// inText a in `aaa
// abc in `aaa`bbb`ccc
// 你好世界 in `你`好啊 ，哈哈，这个`世`界
func inText(key, text string) bool {
	// textMin := 0
	// textMax := len(text) - 1
	matchCount := 0
	keyIndex := 0
	textIndex := 0
	for keyIndex < len(key) && keyIndex < len(text) && textIndex < len(text) {
		// if keyIndex > textMax {
		// 	break
		// }

		keyRune := key[keyIndex]
		textRune := text[textIndex]
		// fmt.Println(keyIndex,textIndex,)
		if keyRune == textRune {
			matchCount++
			keyIndex++
		}
		textIndex++
	}

	return len(key) == matchCount
}
