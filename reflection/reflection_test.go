package main

import (
	"reflect"
	"slices"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         any
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with non-string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"nested struct",
			Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"struct pointer",
			&struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"slices",
			[]Profile{
				{33, "London"},
				{22, "Denmark"},
			},
			[]string{"London", "Denmark"},
		},
		{
			"array",
			[1]Profile{
				{33, "London"},
			},
			[]string{"London"},
		},
		{
			"functions",
			func() (Profile, Profile) {
				return Profile{33, "London"}, Profile{22, "Denmark"}
			},
			[]string{"London", "Denmark"},
		},
	}
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		m := map[string]string{
			"Name":     "Chris",
			"LastName": "Foo",
			"City":     "London",
			"Age":      "33",
			"Country":  "UK",
		}

		var got []string

		walk(m, func(input string) {
			got = append(got, input)
		})

		for _, v := range m {
			if !slices.Contains(got, v) {
				t.Errorf("expected %v to contain %q, but it didn't", got, v)
			}
		}
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "London"}
			aChannel <- Profile{23, "Denmark"}
			close(aChannel)
		}()

		want := []string{"London", "Denmark"}
		var got []string
		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
