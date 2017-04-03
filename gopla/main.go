package main

import (
	"flag"
	"fmt"
	"regexp"
	"sort"

	"github.com/ecdsa521/gopla"
	"github.com/fatih/color"
)

var verbose = flag.Bool("v", false, "Verbose operation")
var best = flag.Bool("b", false, "Only print best option")

func grabByLink(link string) {
	var hash = gopla.GetHash(link)
	grabByHash(hash)
}

func grabByHash(hash string) {
	var data = gopla.GetVideo(hash)
	if *verbose {
		color.Green("%s (%s)", data.Title, data.Duration)
		color.Cyan(data.Description)
		fmt.Println("Available copies:")
	}
	sort.Sort(sort.Reverse(&data.Videos))

	for _, v := range data.Videos {
		if *verbose {
			color.Green("%s - %s kbps - %s - %s", v.Quality, v.Bitrate, v.Size, v.Format)
		}
		color.Blue("%s", v.URL)
		if *best {
			break
		}
	}
}

func main() {

	flag.Parse()
	for _, link := range flag.Args() {

		if match, _ := regexp.MatchString("[a-f0-9]{16}", link); match == true {
			grabByHash(link)
		}
		if match, _ := regexp.MatchString("www.ipla.tv", link); match == true {
			grabByLink(link)
		}
	}

}
