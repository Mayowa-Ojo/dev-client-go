package dev

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func extractDevError(resp *http.Response) (string, error) {
	var v map[string]interface{}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if err := json.Unmarshal(b, &v); err != nil {
		return "", err
	}

	return v["error"].(string), nil
}

func parseUTCDate(t string) (time.Time, error) {
	layout := "2006-01-02T15:04:05Z"

	parsed, err := time.Parse(layout, t)
	if err != nil {
		return time.Time{}, err
	}

	return parsed, nil
}

func parseMarkdownFile(path string) (string, error) {
	if !(strings.HasSuffix(path, ".md") || strings.HasSuffix(path, ".markdown")) {
		return "", errors.New("file must be a markdown file")
	}

	byt, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(byt), nil
}
