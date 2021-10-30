package dev

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type DevAPIError struct {
	msg  string
	code int
}

func (d *DevAPIError) Error() string {
	return fmt.Sprintf("%s: %d", d.msg, d.code)
}

func assertError(err error) bool {
	t := fmt.Sprintf("%T", err)

	return t == "*dev.DevAPIError"
}

func extractDevError(resp *http.Response) error {
	var v map[string]interface{}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}

	return &DevAPIError{
		msg:  v["error"].(string),
		code: int(v["status"].(float64)),
	}
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
