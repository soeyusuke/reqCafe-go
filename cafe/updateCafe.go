package cafe

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type menus struct {
	Week       string `json:"week"`
	Higawari   string `json:"higawari"`
	Donmono    string `json:"donmono"`
	Fish       string `json:"fish"`
	Salada      string `json:"salada"`
	Dessert    string `json:"desert"`
	Onepin     string `json:"onepin"`
	Pasta      string `json:"pasta"`
	Menrui     string `json:"menrui"`
	EveningSet string `json:"eveningset"`
}

//[]string
func GetmenuMon(url string, c1 chan []string) {
	doc, _ := goquery.NewDocument(url)
	var m []string
	doc.Find("div > div > section > table > tbody > tr").Each(func(_ int, s *goquery.Selection) {
		t := s.Find("td").First().Text()
		m = append(m, t)
	})
	c1 <- m
	// return m
}

func GetmenuTue(url string, c2 chan []string) {
	doc, _ := goquery.NewDocument(url)
	var m []string
	doc.Find("div > div > section > table > tbody > tr").Each(func(_ int, s *goquery.Selection) {
		t := s.Find("td").First().Next().Text()
		m = append(m, t)
	})
	c2 <- m
	// return m
}

func GetmenuWen(url string, c3 chan []string) {
	doc, _ := goquery.NewDocument(url)
	var m []string
	doc.Find("div > div > section > table > tbody > tr").Each(func(_ int, s *goquery.Selection) {
		t := s.Find("td").First().Next().Next().Text()
		m = append(m, t)
	})
	c3 <- m
	// return m
}

func GetmenuThu(url string, c4 chan []string) {
	doc, _ := goquery.NewDocument(url)
	var m []string
	doc.Find("div > div > section > table > tbody > tr").Each(func(_ int, s *goquery.Selection) {
		t := s.Find("td").First().Next().Next().Next().Text()
		m = append(m, t)
	})
	c4 <- m
	// return m
}

func GetmenuFri(url string, c5 chan []string) {
	doc, _ := goquery.NewDocument(url)
	var m []string
	doc.Find("div > div > section > table > tbody > tr").Each(func(_ int, s *goquery.Selection) {
		t := s.Find("td").First().Next().Next().Next().Next().Text()
		m = append(m, t)
	})
	c5 <- m
	// return m
}

//並列処理でそれぞれの曜日のmenuをとってくる
func UpdateCafe() {
	var menu []menus //menuのjson型構造体
	url := "http://www.gakushoku.com/univ_mn2.php"
	c1 := make(chan []string)
	c2 := make(chan []string)
	c3 := make(chan []string)
	c4 := make(chan []string)
	c5 := make(chan []string)

	go GetmenuMon(url, c1)
	go GetmenuTue(url, c2)
	go GetmenuWen(url, c3)
	go GetmenuThu(url, c4)
	go GetmenuFri(url, c5)
	mon, tue, wen, thu, fri := <-c1, <-c2, <-c3, <-c4, <-c5

	menu = append(menu, menus{
		Week:       "Monday",
		Higawari:   mon[1],
		Donmono:    mon[2],
		Fish:       mon[3],
		Salad:      mon[4],
		Dessert:    mon[5],
		Onepin:     mon[6],
		Pasta:      mon[7],
		Menrui:     mon[8],
		EveningSet: mon[9],
	})

	menu = append(menu, menus{
		Week:       "Tuesday",
		Higawari:   tue[1],
		Donmono:    tue[2],
		Fish:       tue[3],
		Salad:      tue[4],
		Dessert:    tue[5],
		Onepin:     tue[6],
		Pasta:      tue[7],
		Menrui:     tue[8],
		EveningSet: tue[9],
	})

	menu = append(menu, menus{
		Week:       "Wednesday",
		Higawari:   wen[1],
		Donmono:    wen[2],
		Fish:       wen[3],
		Salad:      wen[4],
		Dessert:    wen[5],
		Onepin:     wen[6],
		Pasta:      wen[7],
		Menrui:     wen[8],
		EveningSet: wen[9],
	})

	menu = append(menu, menus{
		Week:       "Thursday",
		Higawari:   thu[1],
		Donmono:    thu[2],
		Fish:       thu[3],
		Salad:      thu[4],
		Dessert:    thu[5],
		Onepin:     thu[6],
		Pasta:      thu[7],
		Menrui:     thu[8],
		EveningSet: thu[9],
	})

	menu = append(menu, menus{
		Week:       "Friday",
		Higawari:   fri[1],
		Donmono:    fri[2],
		Fish:       fri[3],
		Salad:      fri[4],
		Dessert:    fri[5],
		Onepin:     fri[6],
		Pasta:      fri[7],
		Menrui:     fri[8],
		EveningSet: fri[9],
	})

	// fmt.Println(reflect.TypeOf(menu))

	bytes, _ := json.Marshal(menu)
	ioutil.WriteFile("./json/meshi.json", bytes, os.ModePerm)
	fmt.Println("success")
}
