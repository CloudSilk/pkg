package utils

import (
	"io/ioutil"
	"regexp"
	"strings"
)

func GenIconFont(fileName string) []IconFont {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	reg, err := regexp.Compile(`content: ".*";`)
	if err != nil {
		panic(err)
	}
	str := strings.ReplaceAll(string(data), ":before {", "")
	str = strings.ReplaceAll(str, "}\n", "")
	str = strings.ReplaceAll(str, "}\n", "")

	str = reg.ReplaceAllString(str, "")
	str = strings.ReplaceAll(str, ".", "")
	str = strings.ReplaceAll(str, "  \n", "")
	list := strings.Split(str, "\n")
	var iconFonts []IconFont
	for _, s := range list {
		if s == "" {
			continue
		}
		iconFonts = append(iconFonts, IconFont{
			Key:   s,
			Label: s,
		})
		// fmt.Printf(`{ key: "%s", label: "%s" },`, s, s)
		// fmt.Println()
	}
	return iconFonts
}

type IconFont struct {
	Key   string `json:"key"`
	Label string `json:"label"`
}
