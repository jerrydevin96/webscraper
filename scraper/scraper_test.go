package scraper

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"runtime"
	"strings"
	"testing"
)

var (
	webPageDetailsTestData = WebPageDetails{
		Title:   "page 1",
		Version: "HTML 5",
		H1:      []string{"heading 1"},
		H2:      []string{"heading 2"},
		H3:      []string{"heading 3"},
		H4:      []string{"heading 4"},
		H5:      []string{"heading 5"},
		H6:      []string{"heading 6"},
		Links:   []string{"/page1", "/page2", "https://medium.com"},
	}
)

func compareStringSliceEqulaity(baseDataSlice []string, compareDataSlice []string) bool {
	equality := false
	if len(baseDataSlice) != len(compareDataSlice) {
		return false
	}
	for _, dataSliceValue := range baseDataSlice {
		valuePresent := false
		for _, compareSliceValue := range compareDataSlice {
			if strings.EqualFold(compareSliceValue, dataSliceValue) {
				valuePresent = true
				break
			}
		}
		if !valuePresent {
			return false
		}
	}
	equality = true
	return equality
}

func TestFetchWebPageDetails(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestURL := r.RequestURI
		returnContent := ""
		if requestURL == "/blog" {
			w.WriteHeader(http.StatusOK)
			if runtime.GOOS == "linux" {
				fileContent, err := ioutil.ReadFile("../test-files/page1.html")
				if err != nil {
					t.Log(err.Error())
					t.Fail()
				}
				returnContent = string(fileContent)
			} else if runtime.GOOS == "windows" {
				fileContent, err := ioutil.ReadFile(`..\test-files\page1.html`)
				if err != nil {
					t.Log(err.Error())
					t.Fail()
				}
				returnContent = string(fileContent)
			} else {
				t.Log(runtime.GOOS)
				t.Log("invalid os")
				t.Fail()
			}
			fmt.Fprint(w, returnContent)
			return
		} else if requestURL == "/page1" {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, returnContent)
			return
		} else if requestURL == "/page2" {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, returnContent)
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, returnContent)
			return
		}
	}))
	defer ts.Close()
	pageDetails := FetchWebPageDetails(ts.URL + "/blog")
	if pageDetails.Title != webPageDetailsTestData.Title {
		t.Log("expected title to be " + webPageDetailsTestData.Title + " but got " + pageDetails.Title)
		t.Fail()
	}
	if pageDetails.Version != webPageDetailsTestData.Version {
		t.Log("expected title to be " + webPageDetailsTestData.Version + " but got " + pageDetails.Version)
		t.Fail()
	}
	var linkChecks bool = false
	linkChecks = compareStringSliceEqulaity(webPageDetailsTestData.H1, pageDetails.H1)
	if !linkChecks {
		t.Log("returned H1 list doesnt match up with base H1 list")
		t.Fail()
	}
	linkChecks = compareStringSliceEqulaity(webPageDetailsTestData.H2, pageDetails.H2)
	if !linkChecks {
		t.Log("returned H2 list doesnt match up with base H2 list")
		t.Fail()
	}
	linkChecks = compareStringSliceEqulaity(webPageDetailsTestData.H3, pageDetails.H3)
	if !linkChecks {
		t.Log("returned H3 list doesnt match up with base H3 list")
		t.Fail()
	}
	linkChecks = compareStringSliceEqulaity(webPageDetailsTestData.H4, pageDetails.H4)
	if !linkChecks {
		t.Log("returned H4 list doesnt match up with base H4 list")
		t.Fail()
	}
	linkChecks = compareStringSliceEqulaity(webPageDetailsTestData.H5, pageDetails.H5)
	if !linkChecks {
		t.Log("returned H5 list doesnt match up with base H5 list")
		t.Fail()
	}
	linkChecks = compareStringSliceEqulaity(webPageDetailsTestData.H6, pageDetails.H6)
	if !linkChecks {
		t.Log("returned H6 list doesnt match up with base H6 list")
		t.Fail()
	}
	linkChecks = compareStringSliceEqulaity(webPageDetailsTestData.Links, pageDetails.Links)
	if !linkChecks {
		t.Log("returned links list doesnt match up with base links list")
		t.Fail()
	}
}
