package services

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"minecraft/cmd/internal/config"
	"minecraft/cmd/internal/models"
	"net/http"
	"strconv"
	"strings"
)

func GetPage(page string) []models.Mod {
	url := config.ModeUrl + "page/" + page
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var modeList []models.Mod

	doc.Find(".posts-grid__item").Each(func(i int, s *goquery.Selection) {
		description := getDescription(s)
		link := getLink(s)
		title := getTitle(link)
		version := getVersion(link)
		imgLink := getImage(s)
		downloadLink := getDownload(link)

		mode := models.NewMode(title, description, link, version, imgLink, downloadLink)
		modeList = append(modeList, *mode)

		//fmt.Printf("Номер :%d\nНазвание: %s\nВерсия: %s\nLink: %s\nОписание:%s\nImg: %s\nDownload:%s\n\n\n", i, mode.Title,
		//	mode.Version, mode.Link, mode.Description, mode.ImgLink, mode.DownloadLink)
	})
	log.Printf("Моды скачаны: %v", modeList)
	return modeList
}
func GetAllPages() []models.Mod {
	lastPageInt, err := strconv.ParseInt(config.LastPage, 10, 64)
	if err != nil {
		log.Printf("Cannot parse LastPage to int")
		log.Fatal(err)
	}

	var arrPages []string
	for i := int64(1); i <= lastPageInt; i++ {
		arrPages = append(arrPages, strconv.FormatInt(i, 10))
	}

	var modeList []models.Mod
	for _, page := range arrPages {
		log.Printf("page %s скачан", page)
		modeList = append(modeList, GetPage(page)...)
	}

	return modeList
}
func GetPages(pages string) []models.Mod {
	arrPages := strings.Split(pages, ",")
	for i := 0; i < len(arrPages); i++ {
		arrPages[i] += arrPages[i]
	}

	var modeList []models.Mod
	for _, page := range arrPages {
		log.Printf("page %s скачан", page)
		modeList = append(modeList, GetPage(page)...)
	}

	return modeList
}

func GetLastPageNumber() {
	maxPage := "1000"
	url := config.ModeUrl + "page/" + maxPage

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	paginationClass := doc.Find(".pagination")
	lastPage := paginationClass.Find(".active").Text()
	config.LastPage = lastPage
}
