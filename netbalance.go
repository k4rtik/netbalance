package main

import (
	"code.google.com/p/gcfg"
	"fmt"
	"github.com/kennygrant/sanitize"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
)

type Config struct {
	Credentials struct {
		Login    string
		Password string
	}
}

func main() {
	regex := regexp.MustCompile(`(\d+\.0)`)

	var cfg Config
	configfile, err := homedir.Expand("~/.netbalance.gcfg")
	err = gcfg.ReadFileInto(&cfg, configfile)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.PostForm("http://itapps.youbroadband.in/default/homeuser/login_sql.jsp",
		url.Values{"login": {cfg.Credentials.Login}, "password": {cfg.Credentials.Password}})
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()
	cookies := resp.Cookies()

	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://itapps.youbroadband.in/default/susage/mybalance.jsp", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.AddCookie(cookies[0])

	resp2, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		log.Fatal(err)
	}
	text := sanitize.HTML(string(body))
	result := regex.FindAllString(text, -1)
	fmt.Printf("MBs: %v, Days: %v\n", result[0], result[1])
}
