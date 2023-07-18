package editor

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

const (
	START_SECTION = "<!--START_SECTION:chatgptme-->"
	END_SECTION   = "<!--END_SECTION:chatgptme-->"
)

func ReadMarkdown(mdPath string) (string, error) {
	content, err := ioutil.ReadFile(mdPath)
	if err != nil {
		log.Fatal(err)
	}
	result := string(content)
	return result, nil
}

func WriteMarkdown(mdPath, newTxt string) error {
	f, err := os.Create(mdPath)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(newTxt)
	if err != nil {
		return err
	}
	return nil
}

func FindAndReplace(text, newText string) string {
	pattern := START_SECTION + "[\\s\\S]+" + END_SECTION
	rgx := regexp.MustCompile(pattern)
	result := rgx.ReplaceAllString(text, newText)
	return result
}
