package xkcd

import (
	"testing"
	"unicode/utf8"
)

func TestFixEncoding(t *testing.T) {
	inputs := []string{
		"\u00f0\u009f\u0098\u00b0",
		"\xf0\x9f\x98\xb0",
	}
	expects := []string{
		"\U0001f630",
		"\U0001f630",
	}

	if len(inputs) != len(expects) {
		panic("len(inputs) != len(expects)")
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		expect := expects[i]

		actual := fixEncoding(input)

		t.Logf("fix(%+q)=%+q, expect=%+q", input, actual, expect)

		if !utf8.ValidString(actual) {
			t.Errorf("%+q isn't valid utf-8", actual)
		}

		if actual != expect {
			t.Fail()
		}
	}
}
