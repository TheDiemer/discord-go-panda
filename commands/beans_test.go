package commands

import (
	"strings"
	"testing"
)

func testBeans(t *testing.T) {
	bean := GimmeBeans("test")
	expected := "hey test, https://"

	if !strings.Contains(bean.String(), expected) {
		t.Errorf("expected '%s' but got '%s'", expected, bean.String())
	}
}
