package pageFetcher

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestFetchHTMLPage(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "page accessed")
	}))
	defer ts.Close()
	response, _ := FetchHTMLPage(ts.URL)
	if response != "page accessed" {
		t.Log("gotunexpected value from API call")
		t.FailNow()
	}
}

func TestCheckPageAccessibility(t *testing.T) {
	c := make(chan PageAccessibilityStatus)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		link := r.RequestURI
		if link == "/accessible" {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
		fmt.Fprint(w, "")
	}))
	defer ts.Close()
	go CheckPageAccessibility(ts.URL+"/accessible", c)
	status := <-c
	if status.StatusCode != 200 && status.StatusCode != 201 {
		t.Log("expected status 200/201 but got " + strconv.Itoa(status.StatusCode))
		t.FailNow()
	}
	go CheckPageAccessibility(ts.URL+"/fail", c)
	status = <-c
	if status.StatusCode == 200 || status.StatusCode == 201 {
		t.Log("expected status code other than 200/201 but got " + strconv.Itoa(status.StatusCode))
		t.FailNow()
	}
}
