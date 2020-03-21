package main

import (
	"log"
	"net/http"
	"net/url"
)

func main() {
	requestUrl := "http://someip.com/final/index.php"
	concurrentRequester := 20
	for i := 0; i < concurrentRequester; i++ {
		go func() {
			formValue := url.Values{
				"email":    {"aaaaaaaaaaaa"},
				"password": {"aaaaaaaaaaaa"},
			}
			_, err := http.PostForm(requestUrl, formValue)
			if err != nil {
				log.Println(err.Error())
			}
		}()
	}
}
