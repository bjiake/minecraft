package services

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"minecraft/cmd/internal/config"
	"minecraft/cmd/internal/models"
	"net/http"
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

	return modeList
}
func GetPages(pages string) []models.Mod {
	arrPages := strings.Split(pages, ",")
	for i := 0; i < len(arrPages); i++ {
		arrPages[i] += arrPages[i]
	}

	var modeList []models.Mod
	for _, page := range arrPages {
		modeList = append(modeList, GetPage(page)...)
	}

	return modeList
}
func getDownload(link string) string {
	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal(err)
	}
	var downLoadLink string

	doc.Find(".dl__link").Each(func(i int, s *goquery.Selection) {
		downLoadLink = s.Text()
	})

	return downLoadLink
}
func getLink(s *goquery.Selection) string {
	linkTag := s.Find("a")
	link, _ := linkTag.Attr("href")
	return config.ModeUrl + link[6:]
}
func getTitle(link string) string {

	doc, err := goquery.NewDocument(link)
	if err != nil {
		log.Fatal(err)
	}
	title := strings.TrimSpace(doc.Find(".box__title").First().Text())
	ver := strings.TrimSpace(doc.Find(".post__versions").First().Text())

	// Исключаем текст порт-версии
	title = strings.ReplaceAll(title, ver, "")
	title = strings.ReplaceAll(title, "\n", "")
	title = strings.ReplaceAll(title, " ", "")

	return title
}
func getVersion(link string) string {
	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal(err)
	}
	var version string

	doc.Find(".post__versions").Each(func(i int, s *goquery.Selection) {
		version = s.Text()
	})

	return version
	//name := s.Find(".box__title")
	//v := name.Find(".post__version").Text()
	//parts := strings.Split(strings.TrimSpace(name.Text()), "\n")
	//
	//version := strings.ReplaceAll(parts[1], " ", "")
}
func getImage(s *goquery.Selection) string {
	img := s.Find("img[src]")
	src, exists := img.Attr("src")

	if exists {
		if strings.HasSuffix(src, ".jpg") || strings.HasSuffix(src, ".jpeg") || strings.HasSuffix(src, ".png") || strings.HasSuffix(src, ".webp") {

			return config.BaseUrl + src
		}
	}
	return ""
}
func getDescription(s *goquery.Selection) string {
	description := s.Find(".post__text").Text()
	return description
}
