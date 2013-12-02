package main

import (
    "net/http"
    "net/url"
    "regexp"
    //"os"
    "io/ioutil"
    "fmt"
    "sync"
)

type TestJar struct {
    m      sync.Mutex
    perURL map[string][]*http.Cookie
}

func (j *TestJar) SetCookies(u *url.URL, cookies []*http.Cookie) {
    j.m.Lock()
    defer j.m.Unlock()
    if j.perURL == nil {
        j.perURL = make(map[string][]*http.Cookie)
    }
    j.perURL[u.Host] = cookies
}

func (j *TestJar) Cookies(u *url.URL) []*http.Cookie {
    j.m.Lock()
    defer j.m.Unlock()
    return j.perURL[u.Host]
}

func httpGet(urlStr string) (string, error){
    client := &http.Client{}

    ReqCookies := []*http.Cookie{
        {Name: "xres", Value: "3"},
    }

    u, err := url.Parse(urlStr)
    if err != nil {
        return "", err
    }
    client.Jar = &TestJar{perURL: make(map[string][]*http.Cookie)}
    client.Jar.SetCookies(u, ReqCookies)
    resp, err := client.Get(urlStr)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    bodyStr := fmt.Sprintf("%s", body)
    return bodyStr, nil
}

func listPage() {
    resp, err := httpGet("http://lofi.e-hentai.org/?f_search=chinese&f_apply=Search")
    if err != nil {
        fmt.Println(err)
        return
    }

    // Sorry for the dirty code...
    re, err := regexp.Compile(
        "<a class=\"b\" href=\"http://lofi.e-hentai.org/g/(.*?)\">(.*?)</a>")
    if err != nil {
        fmt.Println(err)
        return
    }
    allTags := re.FindAllString(resp, -1)

    re, err = regexp.Compile(
        "Category:</td><td>(.*?)<")
    if err != nil {
        fmt.Println(err)
        return
    }
    allType := re.FindAllString(resp, -1)

    for i:=0; i < len(allTags); i++ {
        subIndex := len(allType[i]) - 1
        if allType[i][18:subIndex] == "Manga" {
            re, err = regexp.Compile(">(.*?)<")
            titleAll := re.FindString(allTags[i])
            subIndex = len(titleAll) - 1
            title := titleAll[1:subIndex]
            fmt.Println(title)
        }
    }
}

func main() {
    listPage()
}