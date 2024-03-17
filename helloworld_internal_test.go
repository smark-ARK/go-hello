package main

import (
	"testing"
)

func Examplemain() {
	main()
	//Output:
	// Hello World!
}

func TestGreet(t *testing.T) {

	type senario struct {
		lang             language
		expectedGreeting string
	}

	var tests = map[string]senario{
		"English": {
			lang:             "en",
			expectedGreeting: "Hello World!",
		},
		"Arabic": {
			lang:             "ar",
			expectedGreeting: "مرحبًا بالعالم!",
		},
		"French": {
			lang:             "fr",
			expectedGreeting: "Bonjour, le monde !",
		},
		"Spanish": {
			lang:             "es",
			expectedGreeting: "¡Hola, mundo!",
		},
	}

	for senarioName, tc := range tests {
		t.Run(
			senarioName, func(t *testing.T) {
				greeting := greet(tc.lang)
				if greeting != tc.expectedGreeting {
					t.Errorf("Expected %q, but got %q", tc.expectedGreeting, greeting)
				}
			})
	}

}

/* func TestGreet_English(t *testing.T) {
	var expected_language language = "en"
	var expectedGreeting string = "Hello World!"

	greeting := greet(expected_language)

	if greeting != expectedGreeting {
		t.Errorf("expected %v, got %v", expectedGreeting, greeting)
	}
}

func TestGreet_French(t *testing.T) {
	var expected_language language = "fr"
	var expectedGreeting string = "Bonjour le monde!"

	greeting := greet(expected_language)

	if expectedGreeting != greeting {
		t.Errorf("expected %v, got %v", expectedGreeting, greeting)
	}
}
*/
