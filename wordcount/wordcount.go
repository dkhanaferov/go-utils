/**
 *
 */
package wordcount

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"sync/atomic"
)

func CountWords(uris []string) (map[string]int, error) {

	if uris == nil || len(uris) == 0 {
		return nil, errors.New("Missiing URIs list")
	}

	aggregator := make(map[string]int)

	for _, uri := range uris {
		wordsForUri := doCountWords(uri)
		mergeMaps(&aggregator, &wordsForUri)
	}

	return aggregator, nil
}

func CountWordsConcurrent(uris []string) (map[string]int, error) {

	if uris == nil || len(uris) == 0 {
		return nil, errors.New("Missiing URIs list")
	}

	aggMap := make(map[string]int)
	aggChannel := make(chan map[string]int)

	for _, uri := range uris {
		go doCountWordsAsync(uri, aggChannel)
	}

	mapsToAggregateCount := len(uris)
	aggregagedSoFar := 0

	for {
		if mapsToAggregateCount == aggregagedSoFar {
			break
		}

		m := <-aggChannel

		mergeMaps(&aggMap, &m)
		aggregagedSoFar += 1
	}

	return aggMap, nil
}

func CountWordsConcurrent2(uris []string) (map[string]int, error) {

	if uris == nil || len(uris) == 0 {
		return nil, errors.New("Missiing URIs list")
	}

	aggMap := make(map[string]int)
	aggChannel := make(chan string)

	for _, uri := range uris {
		go doCountWordsAsync2(uri, aggChannel)
	}

	mapsToAggregateCount := len(uris)
	aggregagedSoFar := 0

	for word := range aggChannel {

		if word == "**end**" {
			aggregagedSoFar += 1
		}

		if aggregagedSoFar >= mapsToAggregateCount {
			break
		}

		if word == "**end**" {
			continue
		}

		cnt, ok := aggMap[word]

		if !ok {
			aggMap[word] = 1
		} else {
			aggMap[word] = cnt + 1
		}
	}

	return aggMap, nil
}

func CountWordsConcurrent3(uris []string) (map[string]int, error) {

	if uris == nil || len(uris) == 0 {
		return nil, errors.New("Missiing URIs list")
	}

	aggMap := new(sync.Map)

	var mu sync.Mutex

	var wg sync.WaitGroup
	wg.Add(len(uris))

	for _, uri := range uris {
		go doCountWordsAsync3(uri, aggMap, &wg, &mu)
	}

	//Wait for all the goroutines to finish
	wg.Wait()

	outputMap := make(map[string]int)

	aggMap.Range(func(key interface{}, value interface{}) bool {

		if word, ok := key.(string); ok {
			if cnt, cntOk := value.(int32); cntOk {
				outputMap[word] = int(cnt)
				return true
			}
		}

		return false
	})

	return outputMap, nil
}

func CountWordsConcurrent4(uris []string) (map[string]int, error) {

	if uris == nil || len(uris) == 0 {
		return nil, errors.New("Missiing URIs list")
	}

	aggMap := make(map[string]int)

	var mu sync.Mutex

	var wg sync.WaitGroup
	wg.Add(len(uris))

	for _, uri := range uris {
		go doCountWordsAsync4(uri, &aggMap, &wg, &mu)
	}

	//Wait for all the goroutines to finish
	wg.Wait()

	return aggMap, nil
}

// Merges M2 into M1
func mergeMaps(m1 *map[string]int, m2 *map[string]int) {

	if *m1 == nil && m2 != nil {
		*m1 = *m2
	}

	if m1 != nil && *m2 == nil {
		return
	}

	for k, v := range *m2 {
		existingWordCount, ok := (*m1)[k]
		if !ok {
			(*m1)[k] = v
		} else {
			(*m1)[k] = existingWordCount + v
		}
	}
}

func doCountWordsAsync(uri string, c chan map[string]int) {
	c <- doCountWords(uri)
}

func doCountWordsAsync2(uri string, c chan string) {

	fmt.Println("Counting " + uri)

	resp, err := http.Get(uri)

	if err == nil && resp != nil && resp.Body != nil {

		defer resp.Body.Close()

		data, e := ioutil.ReadAll(resp.Body)

		if e == nil && len(data) > 0 {

			for _, s := range strings.Fields(string(data)) {

				normalized, nErr := normalizeWord(s)

				if nErr == nil && normalized != "" && len(normalized) >= 1 {
					c <- normalized
				}
			}
		}
	}

	c <- "**end**"
}

func doCountWordsAsync3(uri string, aggMap *sync.Map, wg *sync.WaitGroup, mu *sync.Mutex) {

	fmt.Println("Counting " + uri)

	resp, err := http.Get(uri)

	if err == nil && resp != nil && resp.Body != nil {

		defer resp.Body.Close()

		data, e := ioutil.ReadAll(resp.Body)

		if e == nil && len(data) > 0 {

			for _, s := range strings.Fields(string(data)) {

				normalized, nErr := normalizeWord(s)

				if nErr == nil && normalized != "" && len(normalized) >= 1 {

					var initVal int32 = 1
					var cnt interface{}
					var loaded bool = false

					mu.Lock()
					cnt, loaded = aggMap.LoadOrStore(normalized, initVal)

					if cntInt, isInt32 := cnt.(int32); loaded && isInt32 {
						aggMap.Store(normalized, atomic.AddInt32(&cntInt, 1))
					}

					mu.Unlock()
				}
			}
		}
	}

	wg.Done()
}

func doCountWordsAsync4(uri string, aggMap *map[string]int, wg *sync.WaitGroup, mu *sync.Mutex) {

	fmt.Println("Counting " + uri)

	resp, err := http.Get(uri)

	if err == nil && resp != nil && resp.Body != nil {

		defer resp.Body.Close()

		data, e := ioutil.ReadAll(resp.Body)

		if e == nil && len(data) > 0 {

			for _, s := range strings.Fields(string(data)) {

				normalized, nErr := normalizeWord(s)

				if nErr == nil && normalized != "" && len(normalized) >= 1 {
					mu.Lock()

					cnt, ok := (*aggMap)[normalized]

					if !ok {
						(*aggMap)[normalized] = 1
					} else {
						(*aggMap)[normalized] = cnt
					}

					mu.Unlock()
				}
			}
		}
	}

	wg.Done()
}

func doCountWords(uri string) map[string]int {

	fmt.Println("Counting " + uri)

	resp, err := http.Get(uri)

	if err == nil && resp != nil && resp.Body != nil {

		defer resp.Body.Close()

		data, e := ioutil.ReadAll(resp.Body)

		if e == nil && len(data) > 0 {

			aggregator := make(map[string]int)

			for _, s := range strings.Fields(string(data)) {

				normalized, nErr := normalizeWord(s)

				if nErr == nil && normalized != "" && len(normalized) >= 1 {

					cnt, ok := aggregator[normalized]

					if !ok {
						aggregator[normalized] = 1
					} else {
						aggregator[normalized] = cnt + 1
					}
				}
			}

			return aggregator
		}
	}

	return nil
}

func normalizeWord(raw string) (string, error) {
	pattern := regexp.MustCompile("^[A-Za-z\\'']+$")
	matched := pattern.FindString(raw)

	if len(matched) > 1 || strings.EqualFold("i", matched) || strings.EqualFold("a", matched) {
		return strings.TrimSpace(strings.ToLower(matched)), nil
	}

	return "", errors.New("unsupported string")
}
