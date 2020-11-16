package ruokalista

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// FoodURL kertoo mistä ruokalista haetaan
const FoodURL = "https://ravintolapalvelut.iss.fi/espoon-tietokyl%C3%A4"

func getFoodHTML() (io.Reader, error) {
	res, err := http.Get(FoodURL)
	if err != nil {
		fmt.Println("Error fetching data: ", err.Error())
		return nil, err
	}
	return res.Body, nil
}

// GetThisWeeksFood parses the foodhtml and returns an array of days
func GetThisWeeksFood() (Viikko, error) {
	viikko := Viikko{}
	count := 0
	htmlReader, err := getFoodHTML()
	if err != nil {
		return Viikko{}, err
	}
	doc, err := goquery.NewDocumentFromReader(htmlReader)

	if err != nil {
		fmt.Println("Jotain meni erittäin pahasti pieleen eikä se ehkä oo mun vika")
		return Viikko{}, err
	}
	var title string
	var lista *goquery.Selection
	doc.Find(".lunch-menus").Children().EachWithBreak(func(i int, s *goquery.Selection) bool {

		if i%2 == 0 {
			title = s.Text()
			if strings.HasPrefix(title, "Lukiolaisten lounaslista viikko") {
				lista = s.Next()
				return false
			}
		}

		return true
	})
	lista.Find(".lunch-menu__days .lunch-menu__day").EachWithBreak(func(i int, s *goquery.Selection) bool {
		if count >= 5 { // kinda scuff but works
			return false
		}

		viikonpaiva := s.Find("h2").First().Text()
		norm := s.Find("p").First().Text()
		veg := s.Find("p").Last().Text()
		päivä := Päivä{
			Perus:       norm,
			Veg:         veg,
			Viikonpäivä: viikonpaiva,
		}
		viikko = append(viikko, päivä)
		count++
		return true
	})

	return viikko, nil
}
