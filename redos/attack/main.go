package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"
)

var wg sync.WaitGroup

func main() {
	requestUrl := "http://someip.com/final/index.php"
	concurrentRequester := 20
	for i := 0; i < concurrentRequester; i++ {
		wg.Add(1)
		go func() {
			formValue := url.Values{
				"email":    {"aaaaaaaaaaaa"},
				"password": {"aaaaaaaaaaaa"},
			}
			response, err := http.PostForm(requestUrl, formValue)
			if err != nil {
				log.Println(err.Error())
			}else{
				defer response.Body.Close()
				body, _ := ioutil.ReadAll(response.Body)
				fmt.Println(string(body))
			}
			wg.Done()
		}()
	}
	wg.Done()
	fmt.Println("done")
}
