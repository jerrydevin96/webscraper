package scraper

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type HTMLVersionTestData struct {
	link    string
	html    string
	version string
}

var (
	htmlVersionDataSet = []HTMLVersionTestData{{
		link: "/page1",
		html: `<!DOCTYPE html>
		<html lang="en-US">
		<head>
		<title>HTML doctype declaration</title>`,
		version: "HTML 5",
	}, {
		link: "/page2",
		html: `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
		<html lang="en-US">
		<head>
		<title>HTML doctype declaration</title>`,
		version: "HTML 4.01",
	}, {
		link: "/page3",
		html: `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.1//EN" "http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">
		<html lang="en-US">
		<head>
		<title>HTML doctype declaration</title>`,
		version: "XHTML 1.1",
	}, {
		link: "/page4",
		html: `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 3.2//EN">
		<html lang="en-US">
		<head>
		<title>HTML doctype declaration</title>`,
		version: "HTML 3.2",
	}}
)

func TestVersionHandler(t *testing.T) {
	getHTMLVersionHandler("https://medium.com/")
}

func TestProcessHTMLForVersion(t *testing.T) {
	for _, testData := range htmlVersionDataSet {
		version, _ := processHTMLForVersion(testData.html)
		if version != testData.version {
			t.Log(testData.html)
			t.Log("expected HTML version " + testData.version + " but function returned " + version)
		}
	}
}

func TestGetHTMLVersionHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		link := r.RequestURI
		responseData := ""
		for _, testData := range htmlVersionDataSet {
			dataLink := testData.link
			if link == dataLink {
				responseData = testData.html
				break
			}
		}
		fmt.Fprint(w, responseData)
	}))
	defer ts.Close()
	for _, testData := range htmlVersionDataSet {
		version, _ := getHTMLVersionHandler(ts.URL + testData.link)
		if version != testData.version {
			t.Log(testData.html)
			t.Log("expected HTML version " + testData.version + " but function returned " + version)
		}
	}
}
