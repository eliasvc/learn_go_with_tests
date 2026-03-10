package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "fakettp://bunkurl.arg"
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
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

func BenchmarkWebsiteChecker(b *testing.B) {
	urls := make([]string, 100)

	for i := range 100 {
		urls[i] = "a url"
	}

	for b.Loop() {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
