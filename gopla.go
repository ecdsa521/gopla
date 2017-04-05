package gopla

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/headzoo/surf"
	"github.com/tidwall/gjson"
)

func init() {

}

//VideoStream has actual video file data
type VideoStream struct {
	URL     string
	Bitrate string
	Format  string
	Size    string
	Quality string
}

//VideoData has all information about video page
type VideoData struct {
	Title       string
	Description string
	Thumbnails  []string
	Duration    string
	Hash        string
	Videos      VideoStreams
}

//SearchData has all information from search page
type SearchData struct {
	Title       string
	Description string
	ID          string
	Date        string
	Image       string
	URL         string
	Hash        string
}

//SearchDatas ...
type SearchDatas []SearchData

//VideoDatas ..
type VideoDatas []VideoData

//VideoStreams ..
type VideoStreams []VideoStream

func (slice VideoStreams) Len() int {
	return len(slice)
}

func (slice VideoStreams) Less(i, j int) bool {
	var s1, _ = strconv.Atoi(slice[i].Size)
	var s2, _ = strconv.Atoi(slice[j].Size)
	return s1 < s2
}

func (slice VideoStreams) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func getJSON(url string) gjson.Result {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("User-Agent", "Mozilla")
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	json, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var jsonR = gjson.Parse(string(json))
	return jsonR
}

//GetAllHashes returns all video hashes from given search
func GetAllHashes(id string, name string) []string {
	var link = fmt.Sprintf("https://seeker.redefine.pl/ipla/multi.json?free=1&page=1&platform=www_ipla_tv&portal_id=ipla&query=%s&category_ids%%5B%%5D=%s&size=5000&sort=recent&sound=&top_category=", url.QueryEscape(name), url.QueryEscape(id))
	var jsonR = getJSON(link)
	var data []string
	for _, v := range jsonR.Array() {
		data = append(data, v.Get("media_id").String())
	}
	return data
}

//FindVideo finds the material and returns informations about it
func FindVideo(name string) SearchDatas {
	var link = fmt.Sprintf("https://seeker.redefine.pl/ipla/multi.json?free=1&page=1&platform=www_ipla_tv&portal_id=ipla&query=%s&size=150&sort=recent&sound=&top_category=KATEGORIE", url.QueryEscape(name))
	var jsonR = getJSON(link)

	if len(jsonR.Array()) == 0 {
		return nil
	}
	var data SearchDatas
	for _, v := range jsonR.Array() {
		data = append(data, SearchData{
			ID:          v.Get("id").String(),
			Description: v.Get("description").String(),
			Title:       v.Get("title").String(),
			Date:        v.Get("created_date").String(),
			Image:       v.Get("image").String(),
			URL:         fmt.Sprintf("http://www.ipla.tv/kategoria/%s", v.Get("id").String()),
			Hash:        v.Get("media_id").String(),
		})
	}
	return data
}

//GetHash returns playvod| link from http link
func GetHash(link string) string {
	bow := surf.NewBrowser()
	bow.AddRequestHeader("Accept", "text/html")
	bow.AddRequestHeader("Accept-Charset", "utf8")
	bow.AddRequestHeader("User-Agent", "Mozilla")
	err := bow.Open(link)
	if err != nil {
		panic(err)
	}
	bow.Find("a.start-watch").Each(func(_ int, s *goquery.Selection) {
		link, _ = s.Attr("href")

		re := regexp.MustCompile("^ipla://playvod-1\\|([a-f0-9]{16})")
		link = re.ReplaceAllString(link, "$1")
	})
	return link
}

//GetVideo returns VideoStream object with basic link informations
func GetVideo(hash string) VideoData {

	var jsonR = getJSON("http://getmedia.redefine.pl/vods/get_vod/?cpid=1&ua=mipla/23&media_id=" + hash)
	var data = VideoData{
		Hash:        hash,
		Title:       jsonR.Get("vod.title").String(),
		Description: jsonR.Get("vod.text").String(),
		Duration:    jsonR.Get("vod.duration").String(),
		Videos:      VideoStreams{},
	}

	jsonR.Get("vod.copies").ForEach(func(key gjson.Result, value gjson.Result) bool {

		data.Videos = append(data.Videos, VideoStream{
			URL:     value.Get("url").String(),
			Bitrate: value.Get("bitrate").String(),
			Format:  value.Get("format").String(),
			Quality: value.Get("quality_p").String(),
			Size:    value.Get("size").String(),
		})
		return true
	})

	return data
}
