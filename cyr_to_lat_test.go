package yaslug

import "testing"

func TestCyrToLat(t *testing.T) {
	tests := []struct {
		name string
		got  string
		want string
	}{
		{"empty", "", ""},
		{"only cyr", "ягода и юндола", "yagoda i yundola"},
		{"only lat", "doom rune", "doom rune"},
		{"symbols", "doom!.#", "doom!.#"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := cyrToLat(test.got)
			if got != test.want {
				t.Errorf("%s: got %s, want %s", test.name, got, test.want)
			}
		})
	}
}
