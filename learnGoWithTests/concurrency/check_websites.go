package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)
	// channels are a Go data structure that can both receive and send values.

	for _, url := range urls {
		go func() {
			resultChannel <- result{url, wc(url)}
			// send statement: sending to result struct instead of directly writing into map
			//results[url] = wc(url)
		}()
		// goroutine: often anonymous func
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		//receive expression
		results[r.string] = r.bool
	}
	//time.Sleep(2 * time.Second)
	return results
}
