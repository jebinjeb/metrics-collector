package metrics

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptrace"
	"time"
)

// https://golang.org/pkg/net/http/

func Metrics(c *gin.Context) {
	urls := []string{
		"https://cisco.com",
		"https://oracle.com",
		"https://www.amazon.com",
		"https://www.amazon.in",
		"https://www.teeeeeest.com/" }
	c.String(http.StatusOK, loadData(urls))
}

func loadData(urls []string) string {
	var data bytes.Buffer
	message := make(chan string, 2)

	for _, url := range urls {
		fmt.Println("Checking http response for website", url)
		go getMetrics(url, message)
	}
	for _, _ = range urls {
		data.WriteString(<-message + "\n")
	}

	return data.String()
}

func getMetrics(url string, message chan string) {
	var timeElapsed = httpResponseTime(url, "GET")
	message <- url + "," + timeElapsed
}

func httpResponseTime(url string, method string) string {
	// prepare the request
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	var start time.Time
	trace := &httptrace.ClientTrace{
		GotFirstResponseByte: func() {
			fmt.Println("GotFirstResponseByte in time", time.Since(start))
		},
	}

	request = request.WithContext(httptrace.WithClientTrace(request.Context(), trace))
	start = time.Now()
	_, err = http.DefaultTransport.RoundTrip(request)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	return time.Since(start).String()
}