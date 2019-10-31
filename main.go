package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"

	"github.com/jung-kurt/gofpdf"
)

var (
	headers = []string{
		"id",
		"Product Number",
		"Company Name",
		"< 30",
		"< 60",
		"< 90",
	}
	w   = []float64{25., 25., 25., 25., 25., 25.}
	pdf *gofpdf.Fpdf
)

// type Tabler interface {
// 	Iterator()
// 	Len() int
// }

// type Aging struct {
// 	ID int
// 	ProductNumber
// }

func main() {
	bytesData, err := ioutil.ReadFile("mock_data.json")
	if err != nil {
		log.Fatalln(err)
	}
	var dataMap []map[string]interface{}
	json.Unmarshal(bytesData, &dataMap)

	pdf = gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")
	pdf.Ln(-1)
	for i, d := range dataMap {
		Row(d)
		if i > 10 {
			break
		}
	}
	err = pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		log.Fatalln(err)
	}
}

//Row set a row of data with multicells with height equal to tallest cell
func Row(data map[string]interface{}) {
	var nb int
	var i int
	lines := make([]int, 0, len(data))
	for _, d := range data {
		lines = append(lines, Lines(w[i], fmt.Sprint(d)))
		i++
	}
	nb = Max(lines...)
	h := float64(5 * nb)
	for i, col := range headers {
		x := pdf.GetX()
		y := pdf.GetY()
		pdf.Rect(x, y, w[i], h, "")
		pdf.MultiCell(w[i], 5., fmt.Sprint(data[col]), "", "LP", false)
		pdf.SetXY(x+w[i], y)
	}
	pdf.Ln(h)
}

//Max finds a heights int in array of ints
func Max(nums ...int) int {
	var max int
	for _, num := range nums {
		if max < num {
			max = num
		}
	}
	return max
}

//Lines finds number of lines for a str within given width
func Lines(w float64, str string) int {

	return int(math.Ceil(pdf.GetStringWidth(str) / w))

}

// Headers ignore.
// func Headers(headers DataIterator) {
// 	for {
// 		if str, ok := header.Next(); !ok {
// 			break
// 		}

// 	}
// }
