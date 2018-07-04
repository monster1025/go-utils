package utils

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"
)

var agents = []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_2) AppleWebKit/601.3.9 (KHTML, like Gecko) Version/9.0.2 Safari/601.3.9",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.111 Safari/537.36",
	"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:15.0) Gecko/20100101 Firefox/15.0.1"}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func randomUserAgent() string {
	return agents[randInt(0, len(agents))]
}

func HttpGet(url string, client *http.Client) string {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", randomUserAgent())
	req.Header.Set("Accept-Language", "en-us")
	// req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Referer", "https://www.papajohns.ru/")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Cannot get url %s, err: %s", url, err.Error())
		return ""
	}

	defer resp.Body.Close()
	contents, _ := ioutil.ReadAll(resp.Body)
	return string(contents)
}

func HttpHead(url string, client *http.Client) int {
	req, _ := http.NewRequest("HEAD", url, nil)
	req.Header.Set("User-Agent", randomUserAgent())
	req.Header.Set("Accept-Language", "en-us")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Cannot get url %s, err: %s", url, err.Error())
		return 0
	}

	return resp.StatusCode
}

func HttpGetDump(url string, client *http.Client, file string) (string, string) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", randomUserAgent())
	req.Header.Set("Accept-Language", "en-us")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Cannot get url %s, err: %s", url, err.Error())
		return "", ""
	}

	defer resp.Body.Close()

	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatal(err)
	}
	dumpStr := string(dump)
	File_put_contents(file, dumpStr)

	contents, _ := ioutil.ReadAll(resp.Body)
	return string(contents), dumpStr
}

func HttpPostJSON(url string, data string, referer_optional string, client *http.Client, isAjax bool) string {
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	req.Header.Set("User-Agent", randomUserAgent())
	if isAjax {
		req.Header.Set("X-Requested-With", "XMLHttpRequest")
	}
	//	req.Header.Set("Origin", "https://www.topcashback.com")
	req.Header.Set("Accept-Language", "en-us")
	req.Header.Set("Content-Type", "application/json")
	if referer_optional != "" {
		req.Header.Set("Referer", referer_optional)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Cannot post url %s, err: %s", url, err.Error())
		return ""
	}

	defer resp.Body.Close()
	contents, _ := ioutil.ReadAll(resp.Body)
	return string(contents)
}

func HttpPostString(url string, data string, referer_optional string, client *http.Client, isAjax bool) string {
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	req.Header.Set("User-Agent", randomUserAgent())
	if isAjax {
		req.Header.Set("X-Requested-With", "XMLHttpRequest")
	}
	req.Header.Set("Accept-Language", "en-us")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	if referer_optional != "" {
		req.Header.Set("Referer", referer_optional)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Cannot post url %s, err: %s", url, err.Error())
		return ""
	}

	defer resp.Body.Close()
	contents, _ := ioutil.ReadAll(resp.Body)
	return string(contents)
}

func HttpPostStringToken(url string, data string, referer_optional string, token string, client *http.Client, isAjax bool) string {
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	req.Header.Set("User-Agent", randomUserAgent())
	if isAjax {
		req.Header.Set("X-Requested-With", "XMLHttpRequest")
	}
	req.Header.Set("Accept-Language", "en-us")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	if referer_optional != "" {
		req.Header.Set("Referer", referer_optional)
	}

	req.Header.Set("X-CSRF-Token", token)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Cannot post url %s, err: %s", url, err.Error())
		return ""
	}

	defer resp.Body.Close()
	contents, _ := ioutil.ReadAll(resp.Body)
	return string(contents)
}

func HttpPost(url string, data url.Values, client *http.Client) string {
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data.Encode())))
	req.Header.Set("User-Agent", randomUserAgent())
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Accept-Language", "en-us")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Referer", "http://user.everbuying.net/m-users-a-sign.htm")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Cannot post url %s, err: %s", url, err.Error())
		return ""
	}

	if resp.StatusCode == 302 && client.CheckRedirect != nil {
		return resp.Header.Get("Location")
	}

	defer resp.Body.Close()
	contents, _ := ioutil.ReadAll(resp.Body)
	return string(contents)
}

func HttpDownloadFile(url string, filepath string, client *http.Client) error {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", randomUserAgent())
	req.Header.Set("Accept-Language", "en-us")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Cannot get url %s, err: %s", url, err.Error())
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
