package utils

import (
	"bufio"
	"net/url"
	"os"
)

func ReadTxtUrl(filename string) ([]string, error) {
	filePath, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(filePath)
	fileScanner.Split(bufio.ScanLines)

	var Urls []string
	for fileScanner.Scan() {
		Urls = append(Urls, fileScanner.Text())
	}
	filePath.Close()
	return Urls, nil
}

func CreateHtmlFile(u string) (string, error) {
	ext := ".html"

	parsed, err := url.Parse(u)
	if err != nil {
		return "", err
	}
	name := parsed.Host
	file := name + ext
	return file, nil

}
