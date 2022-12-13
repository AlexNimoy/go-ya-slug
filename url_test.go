package yaslug

import "testing"

func TestUrl(t *testing.T) {
	tests := []struct {
		name string
		got  string
		want string
	}{
		{"empty", "", ""},
		{"cyr", "урл url", "url-url"},
		{"big leters", "It is Awesome", "it-is-awesome"},
		{"spaces", "  It   is   Awesome ", "it-is-awesome"},
		{"symbols", "  It, is,  *Awesome*#", "it-is-awesome"},
		{"spaces and symbols", "-  It is Awesome -", "it-is-awesome"},
		{"digits", "Golang 1.19", "golang-1-19"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Url(test.got)
			if got != test.want {
				t.Errorf("%s: got %s, want %s", test.name, got, test.want)
			}
		})
	}
}
