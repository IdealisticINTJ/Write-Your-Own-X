
package main

import "testing"

func TestFormatStats(t *testing.T) {
	cases := []struct {
		Description   string
		InputStats    stats
		InputFilename string
		Want          string
	}{
		{"Empty", stats{bytes: 0, words: 0, lines: 0, chars: 0}, "", "0\t0\t0\t"},
		{"Default", stats{bytes: 11, words: 2, lines: 1, chars: 0}, "filename", "1\t2\t11\tfilename"},
		{"Chars", stats{bytes: 0, words: 0, lines: 0, chars: 100}, "filename", "100\tfilename"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := formatStats(true, true, true, true, test.InputStats, test.InputFilename)

			if got != test.Want {
				t.Errorf("got %v, want %v", got, test.Want)
			}
		})
	}
}
