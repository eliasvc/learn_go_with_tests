package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	return url != "fakettp://bunkurl.arg"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://google.com",
		"https://linode.com",
		"https://akamai.com",
		"fakettp://bunkurl.arg",
	}

	want := map[string]bool{
		"https://google.com":    true,
		"https://linode.com":    true,
		"https://akamai.com":    true,
		"fakettp://bunkurl.arg": false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("wanted %v, but got %v", want, got)
	}

}
