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
var search = flag.String("s", "", "Search for this show")
var grabHashes = flag.Bool("a", false, "Grab all hashes for search query")
var grabLinks = flag.Bool("l", false, "Grab all links for search query")
var wgetLinks = flag.Bool("w", false, "Generate batch output for wget downloads")

func wgetAllFiles(hash []string) {
	for _, h := range hash {

		var data = gopla.GetVideo(h)
		sort.Sort(sort.Reverse(&data.Videos))
		for _, d := range data.Videos {
			if *best {
				fmt.Printf("wget -O '%s.mp4' '%s'\n", data.Title, d.URL)
				break
			} else {
				fmt.Printf("wget -O '%s (%s).mp4' '%s'\n", data.Title, d.Quality, d.URL)
			}

		}
	}

}

func grabByName(name string) {
	var res = gopla.FindVideo(name, "KATEGORIE")
	if res == nil || len(res) == 0 {
		res = gopla.FindVideo(name, "")

	}
	//for _, v := range gopla.GetAllHashes(data.ID, data.Title) {
	//	var data = gopla.GetVideo(v)

	for _, v := range res {
		//f/mt.Printf("%v\n", v)
		if *verbose {
			color.Green(v.Title)
			color.Cyan(v.Description)
			color.Blue("%s | %s", v.URL, v.Hash)

		}
		if v.Type == "vod" {
			var hashes []string
			hashes = append(hashes, v.Hash)
			if *grabLinks {
				grabByHash(v.Hash)
			}
			if *grabHashes {
				color.Blue(v.Hash)
			}
			if *wgetLinks {
				wgetAllFiles(hashes)
			}
		} else if v.Type == "category" {
			var hashes = gopla.GetAllHashes(v.ID, v.Title)
			if *wgetLinks {
				wgetAllFiles(hashes)
			}
			for _, d := range hashes {

				if *grabLinks {
					grabByHash(d)
				}
				if *grabHashes {
					color.Blue(d)
				}

			}
		}
	}
}

func grabByLink(link string) {
	var hash = gopla.GetHash(link)
	if *verbose {
		color.Cyan("playvod:%s", hash)
	}
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
	if len(*search) > 0 {
		grabByName(*search)
	} else {
		for _, link := range flag.Args() {

			if match, _ := regexp.MatchString("[a-f0-9]+", link); match == true {
				grabByHash(link)
			}
			if match, _ := regexp.MatchString("www.ipla.tv", link); match == true {
				grabByLink(link)
			}
		}
	}
}
