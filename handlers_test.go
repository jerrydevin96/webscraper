package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"runtime"
	"testing"
)

var (
	testData = WebPageDetails{
		HTMLVersion:             "HTML 5",
		PageTitle:               "page 1",
		H1Length:                1,
		H2Length:                1,
		H3Length:                1,
		H4Length:                1,
		H5Length:                1,
		H6Length:                1,
		InternalLinksLength:     2,
		ExternalLinksLength:     1,
		InAccessibleLinksLength: 1,
	}
)

func TestPageDetailsHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestURL := r.RequestURI
		returnContent := ""
		if requestURL == "/blog" {
			w.WriteHeader(http.StatusOK)
			if runtime.GOOS == "linux" {
				fileContent, err := ioutil.ReadFile("./test-files/page1.html")
				if err != nil {
					t.Log(err.Error())
					t.Fail()
				}
				returnContent = string(fileContent)
			} else if runtime.GOOS == "windows" {
				fileContent, err := ioutil.ReadFile(`.\test-files\page1.html`)
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
	testURL := ts.URL + "/blog"
	requestData := `{"url": "` + testURL + `"}`
	responseData := pageDetailsHandler(requestData)

	responseParseData := &WebPageDetails{}
	json.Unmarshal([]byte(responseData), &responseParseData)
	if responseParseData.HTMLVersion != testData.HTMLVersion {
		t.Log("expected html version to be " + testData.HTMLVersion + " but got " + responseParseData.HTMLVersion)
		t.Fail()
	}
	if responseParseData.PageTitle != testData.PageTitle {
		t.Log("expected page title to be " + testData.PageTitle + " but got " + responseParseData.PageTitle)
		t.Fail()
	}
	if responseParseData.H1Length != testData.H1Length {
		t.Log("test and actual h1 header length doesnt match")
		t.Fail()
	}
	if responseParseData.H2Length != testData.H2Length {
		t.Log("test and actual h2 header length doesnt match")
		t.Fail()
	}
	if responseParseData.H3Length != testData.H3Length {
		t.Log("test and actual h3 header length doesnt match")
		t.Fail()
	}
	if responseParseData.H4Length != testData.H4Length {
		t.Log("test and actual h4 header length doesnt match")
		t.Fail()
	}
	if responseParseData.H5Length != testData.H5Length {
		t.Log("test and actual h5 header length doesnt match")
		t.Fail()
	}
	if responseParseData.H6Length != testData.H6Length {
		t.Log("test and actual h6 header length doesnt match")
		t.Fail()
	}
	if responseParseData.InternalLinksLength != testData.InternalLinksLength {
		t.Log("test and actual internal links length doesnt match")
		t.Fail()
	}
	if responseParseData.ExternalLinksLength != testData.ExternalLinksLength {
		t.Log("test and actual external links length doesnt match")
		t.Fail()
	}
	if responseParseData.InAccessibleLinksLength != testData.InAccessibleLinksLength {
		t.Log("test and actual inaccessible links length doesnt match")
		t.Fail()
	}

}
