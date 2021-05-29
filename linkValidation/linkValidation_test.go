package linkValidation

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type linkAccessibilityData struct {
	inputEndpoints        []string
	accessibleEndpoints   []string
	inAccessibleEndpoints []string
}

type linkClassificationData struct {
	inputLinks    []string
	internalLinks []string
	externalLinks []string
	url           string
}

var (
	linkAccessibilityTestData = linkAccessibilityData{
		inputEndpoints:        []string{"/page1", "/page2", "/page3", "/page4", "/page5", "/page6"},
		accessibleEndpoints:   []string{"/page2", "/page3", "/page4", "/page6"},
		inAccessibleEndpoints: []string{"/page1", "/page5"},
	}
	linkClassificationTestData = linkClassificationData{
		inputLinks: []string{"https://internal.com/p1", "https://internal.com/p2", "https://internal.com/p3",
			"https://external.com/p1", "https://external.com/p2"},
		internalLinks: []string{"https://internal.com/p1", "https://internal.com/p2", "https://internal.com/p3"},
		externalLinks: []string{"https://external.com/p1", "https://external.com/p2"},
		url:           "https://internal.com/blog",
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

func TestCheckLinkAccessibility(t *testing.T) {
	inputEndpoints := linkAccessibilityTestData.inputEndpoints
	accessibleEndpoints := linkAccessibilityTestData.accessibleEndpoints
	inAccessibleEndpoints := linkAccessibilityTestData.inAccessibleEndpoints
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		linkCheckStatus := false
		uri := r.RequestURI
		for _, link := range accessibleEndpoints {
			if strings.EqualFold(uri, link) {
				linkCheckStatus = true
				w.WriteHeader(http.StatusOK)
				break
			}
		}
		if !linkCheckStatus {
			for _, link := range inAccessibleEndpoints {
				if strings.EqualFold(uri, link) {
					w.WriteHeader(http.StatusNotFound)
					break
				}
			}
		}
		fmt.Fprintln(w, "")
	}))
	defer ts.Close()
	testServerHost := ts.URL
	inputLinks := make([]string, 0)
	for _, link := range inputEndpoints {
		inputLinks = append(inputLinks, (testServerHost + link))
	}
	inaccessibleLinks := checkLinkAccessibility(inputLinks)
	if len(inaccessibleLinks) == len(inAccessibleEndpoints) {
		for _, link := range inaccessibleLinks {
			linkValid := false
			for _, endPoint := range inAccessibleEndpoints {
				if strings.Contains(link, endPoint) {
					linkValid = true
					break
				}
			}
			if !linkValid {
				t.Log("link returned by function does not match with test data set")
				t.FailNow()
			}
		}
	} else {
		t.Log("function returned more/less links than expected number of links")
		t.FailNow()
	}
}

func TestClassifyLinks(t *testing.T) {
	inputLinks := linkClassificationTestData.inputLinks
	inputURL := linkClassificationTestData.url
	validationInternalLinks := linkClassificationTestData.internalLinks
	validationExternalLinks := linkClassificationTestData.externalLinks

	internalLinks, externalLinks := classifyLinks(inputLinks, inputURL)

	comparisionStatus := compareStringSliceEqulaity(validationInternalLinks, internalLinks)
	if !comparisionStatus {
		t.Log("Test data slice and slice returned by funtion donot match")
		t.FailNow()
	}
	comparisionStatus = compareStringSliceEqulaity(validationExternalLinks, externalLinks)
	if !comparisionStatus {
		t.Log("Test data slice and slice returned by funtion donot match")
		t.FailNow()
	}
}
