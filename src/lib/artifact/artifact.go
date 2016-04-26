package artifact

import (
	"fmt"
	"log"
	"net/http"
)

func Sequential(urls []string) {
	for _, url := range urls {
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		// body, _ := ioutil.ReadAll(res.Body)
		defer res.Body.Close()

		// fmt.Println(string(body))
		fmt.Println(res.Status)
	}
}

func Parallel(urls []string) {
	statusChan := make(chan string)

	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}

			// body, _ := ioutil.ReadAll(res.Body)
			defer res.Body.Close()

			// fmt.Println(string(body))
			// fmt.Println(url, res.Status)
			statusChan <- res.Status
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-statusChan)
	}
}
