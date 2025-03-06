package tmt

import (
	"strings"
	"testing"
)

func Test_NoMissingTech(t *testing.T) {
	for _, v := range migrations {
		for _, w := range v.To {
			_, ok := Tech[strings.ToLower(w)]
			if !ok {
				t.Fatal("missing tech", w)
			}
		}
		for _, w := range v.From {
			_, ok := Tech[strings.ToLower(w)]
			if !ok {
				t.Fatal("missing tech", w)
			}
		}
	}
}
