package artifact

import (
	"fmt"
	"log"
	"net/http"
)

var urlBase string = "https://search.maven.org/solrsearch/select?rows=20&wt=json&q="

func Sequential(artifacts []string) {
	for _, artifact := range artifacts {
		res, err := http.Get(urlBase + artifact)
		if err != nil {
			log.Fatal(err)
		}

		// body, _ := ioutil.ReadAll(res.Body)
		defer res.Body.Close()

		// fmt.Println(string(body))
		fmt.Println(res.Status)
	}
}

func Parallel(artifacts []string) {
	statusChan := make(chan string)

	for _, artifact := range artifacts {
		go func(artifact string) {
			res, err := http.Get(urlBase + artifact)
			if err != nil {
				log.Fatal(err)
			}

			// body, _ := ioutil.ReadAll(res.Body)
			defer res.Body.Close()

			// fmt.Println(string(body))
			// fmt.Println(url, res.Status)
			statusChan <- res.Status
		}(artifact)
	}

	for i := 0; i < len(artifacts); i++ {
		fmt.Println(<-statusChan)
	}
}
