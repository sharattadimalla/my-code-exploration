package main

import (
	"log"
	"os"
	
	"github.com/chenjiandongx/go-echarts/charts"
	)

func main() {

	days := []string{"sun", "mon", "tue", "wed", "thu", "fri", "sat"}
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Sample Go Data Viz"})	
	bar.AddXAxis(days).
		AddYAxis("valA", []int{20, 30, 40, 10, 24, 36}).
		AddYAxis("valB", []int{10,45,50,5,20,30})

	f, err := os.Create("bar.html")
	if err != nil {
		log.Println(err)
	}
	bar.Render(f)
}
