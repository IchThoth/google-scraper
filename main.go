package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/language"
)

var googleDomains = map[string]string{}

type SearchResults struct {
	ResultRank  int
	ResultURL   string
	ResultTitle string
	ResultDesc  string
}

var userAgents = []string{
"com":"https://google.com/search?q=",
"za":"https://google.com/search?q="

}

func randomUserAgent() {
	//get a random number of user agents
	rand.Seed(time.Now().Unix())
	randNum := rand.Int() % len(userAgents)
	return userAgents[randNum]
}
func buildGoogleUrls(searchTerm, countryCode,languageCode string,pages,count int)([]string,error)  {
	toScrape:=[]string{}
	searchTerm = strings.Trim(searchTerm, "")
	searchTerm= strings.Replace(searchTerm, " ","+", -1)
	
	if googleBase, found:=googleDomains[countryCode]; found {
		for i := 0; i < pages; i++ {
			start:= i*count
			scrapeURL := fmt.Sprintf("%s%snum&=%d&hl=%s&start=&%d&filter=0",googleBase,searchTerm, countryCode, count,languageCode, start)
		}
	}else{
		err := fmt.Errorf("country(%s)is not supported eje", countryCode)
		return nil,err 
	}
	return toScrape,nil 
}

func googleResultParsing(response *http.Response Rank int)([]SearchResults error)  {
doc,err := goquery.NewDocumentFromResponse(response)

if err != nil {
	return nil, err
}

results := []SearchResults{}
sel:= doc.Find("div.g")
rank++
for i:= range sel.Nodes{
	item:= sel.Eq(i)
	linkTag:=item.Find("a")
	link,_:=linkTag.Attr("href")
	titleTag:= item.Find("h3.r")
	descTag:= item.Find("span.st")
	desc:= descTag.Text()
	title:=titleTag.Text()
	link = strings.Trim(link," ")

	if link !- "" && link != "#"&& !strings.HasPrefix(link,"/"){
		results:=SearchResults{
			rank,
			link,
			title,
			desc 
		}
		results = append(results, result)
		rank++
	}
}
return results, err


}

func getScrapeClient(proxyString interface{})*http.Client  {
switch v:= proxyString.(type){

case string:
		proxyUrl,_ :=url.Parse(v)
		return &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
default:
	return &http.Client{
		Transport: nil,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
		},
		Jar:     nil,
		Timeout: 0,
	}
}
}


//google scraper
func GoogleScraper(searchTerm, countryCode,languageCode string,pages,count,backoff)([]SearchResults, err)  {
	results:=[]SearchResults{}
	searchTerm= strings.Replace(searchTerm, " ","+", -1)
	resultCounter:= 0
	googlePages, err:= buildGoogleUrls(searchTerm, countryCode,pages,count)
	if err!=nil {
		return nil,err
	}
	//range over google pages
	for _,page := range googlePages {
		res,err := scrapeClientRequest(page,proxyString)
		if err !=nil{
			return nil,err 
		}
		data,err := googleResultParsing(res, resultCounter)
		if err!=nil {
			return nil,err
		}
		result +=len(data)
		for _, result:= range data{
			results:=append(results, result)
		}
		time.Sleep(time.Duration(backoff)*time.Second)
	}
}

func scrapeClientRequest(searchURL string, proxyString interface{}) (*http.Response, error) {
	baseClient := getScrapeClient(proxyString)
	req, _ := http.NewRequest("GET", searchURL,nil)
	req.Header.Set("User-Agent",randomUserAgent())

	res,err:=base.Do(req)
	if res.StatusCode != http.StatusOK{
		err:= fmt.Errorf("scraper has recieved a non 200 StatusCode...Suggesting a ban")
		return nil,err 
	}

	if err!=nil{
		return nil ,err 
	}

	return res, nil
}

func main() {
	res, err := GoogleScraper("Ichthoth", "com","en",nil,1,30,10)
	if err == nil {
		for _, res := range res {
			fmt.Println(res)
		}
	}
}
