package main

import (
	"strings"
	"testing"
)

func testBeans(t *testing.T) {
	bean := gimmeBeans("test")
	expected := "hey test, https://"

	if !strings.Contains(bean.String(), expected) {
		t.Errorf("expected '%s' but got '%s'", expected, bean.String())
	}
}
