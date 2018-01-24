package main

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

func Test_makeFormatter_makesJSONFormatter_by_default(t *testing.T) {
	actualFormatter := _makeFormatter("")
	expect(t, actualFormatter, &log.JSONFormatter{})
}

func Test_makeFormatter_makesJSONFormatter_when_json_specified(t *testing.T) {
	actualFormatter := _makeFormatter("json")
	expect(t, actualFormatter, &log.JSONFormatter{})
}

func Test_makeFormatter_makesTextFormatter_when_text_specified(t *testing.T) {
	actualFormatter := _makeFormatter("text")
	expect(t, actualFormatter, &log.TextFormatter{DisableColors: true})
}
