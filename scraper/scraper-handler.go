package scraper

import (
	"log"

	"github.com/gocolly/colly"
)

var (
	Title  = ""
	H1List = make([]string, 0)
	H2List = make([]string, 0)
	H3List = make([]string, 0)
	H4List = make([]string, 0)
	H5List = make([]string, 0)
	H6List = make([]string, 0)
)

func getTitleCollyHTMLCallBack(h *colly.HTMLElement) {
	log.Println("Executing get title Colly HTML call back")
	Title = h.Text
	log.Println("Title is " + Title)
	log.Println("Get title Colly HTML call back completed")
}

func getH1CollyHTMLCallBack(h *colly.HTMLElement) {
	log.Println("Executing get H1 Colly HTML call back")
	header := h.Text
	H1List = append(H1List, header)
	log.Println("Get H1 Colly HTML call back completed")
}

func getH2CollyHTMLCallBack(h *colly.HTMLElement) {
	log.Println("Executing get H2 Colly HTML call back")
	header := h.Text
	H2List = append(H2List, header)
	log.Println("Get H2 Colly HTML call back completed")
}

func getH3CollyHTMLCallBack(h *colly.HTMLElement) {
	log.Println("Executing get H3 Colly HTML call back")
	header := h.Text
	H3List = append(H3List, header)
	log.Println("Get H3 Colly HTML call back completed")
}

func getH4CollyHTMLCallBack(h *colly.HTMLElement) {
	log.Println("Executing get H4 Colly HTML call back")
	header := h.Text
	H4List = append(H4List, header)
	log.Println("Get H4 Colly HTML call back completed")
}

func getH5CollyHTMLCallBack(h *colly.HTMLElement) {
	log.Println("Executing get H1 Colly HTML call back")
	header := h.Text
	H5List = append(H5List, header)
	log.Println("Get H5 Colly HTML call back completed")
}

func getH6CollyHTMLCallBack(h *colly.HTMLElement) {
	log.Println("Executing get H1 Colly HTML call back")
	header := h.Text
	H6List = append(H6List, header)
	log.Println("Get H1 Colly HTML call back completed")
}
