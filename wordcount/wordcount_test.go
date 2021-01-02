package wordcount

import (
	"fmt"
	"log"
	"reflect"
	"testing"
	"time"
)

func TestSuccessfulWordCountForOneUri(t *testing.T) {
	m, err := CountWords([]string{"https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md"})

	if err != nil {
		t.Fatal(err)
	}

	if m == nil {
		t.Fatal("Nil word count map")
	}

	if len(m) <= 0 {
		t.Fatal("Empty word count map")
	}

	for k, v := range m {
		fmt.Printf("%v -> %v \n", k, v)
	}

	cnt, ok := m["markdown"]
	if ok {
		t.Fatal("Found unexpected value")
	}
	if cnt != 0 {
		t.Fatal("Incorrect count for expected value")
	}
}

func TestSuccessfulWordCountForMultipleUris(t *testing.T) {
	m, err := CountWords([]string{"https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md", "https://raw.githubusercontent.com/dkhanaferov/datasciencecoursera/master/HelloWorld.md"})

	if err != nil {
		t.Fatal(err)
	}

	if m == nil {
		t.Fatal("Nil word count map")
	}

	if len(m) <= 0 {
		t.Fatal("Empty word count map")
	}

	for k, v := range m {
		fmt.Printf("%v -> %v \n", k, v)
	}

	cnt, ok := m["markdown"]
	if !ok {
		t.Fatal("Missing expected value")
	}
	if cnt > 1 || cnt <= 0 {
		t.Fatal("Incorrect count for expected value")
	}
}

func TestSuccessfulWordCountConcurrentForMultipleUris(t *testing.T) {
	m, err := CountWordsConcurrent([]string{"https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md", "https://raw.githubusercontent.com/dkhanaferov/datasciencecoursera/master/HelloWorld.md"})

	if err != nil {
		t.Fatal(err)
	}

	if m == nil {
		t.Fatal("Nil word count map")
	}

	if len(m) <= 0 {
		t.Fatal("Empty word count map")
	}

	for k, v := range m {
		fmt.Printf("%v -> %v \n", k, v)
	}

	cnt, ok := m["markdown"]
	if !ok {
		t.Fatal("Missing expected value")
	}
	if cnt > 1 || cnt <= 0 {
		t.Fatal("Incorrect count for expected value")
	}
}

func TestSuccessfulWordCountConcurrentForOneUri(t *testing.T) {
	m1, err1 := CountWordsConcurrent([]string{"https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md"})

	if err1 != nil {
		t.Fatal(err1)
	}

	if m1 == nil {
		t.Fatal("Nil word count map")
	}

	if len(m1) <= 0 {
		t.Fatal("Empty word count map")
	}

	m2, _ := CountWords([]string{"https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md"})

	for k, v := range m2 {
		fmt.Printf("%v -> %v \n", k, v)
	}

	if len(m1) != len(m2) {
		t.Fatal("Length of m1 not equal to m2")
	}

	if !reflect.DeepEqual(m1, m2) {
		t.Fatal("m1 not equal to m2")
	}
}

func TestSuccessfulWordCountConcurrent2ForOneUri(t *testing.T) {
	m1, err1 := CountWordsConcurrent2([]string{"https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md"})

	if err1 != nil {
		t.Fatal(err1)
	}

	if m1 == nil {
		t.Fatal("Nil word count map")
	}

	if len(m1) <= 0 {
		t.Fatal("Empty word count map")
	}

	m2, _ := CountWords([]string{"https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md"})

	for k, v := range m2 {
		fmt.Printf("%v -> %v \n", k, v)
	}

	if len(m1) != len(m2) {
		t.Fatal("Length of m1 not equal to m2")
	}

	if !reflect.DeepEqual(m1, m2) {
		t.Fatal("m1 not equal to m2")
	}
}

func TestSuccessfulWordCountConcurrent2ForMultipleUris(t *testing.T) {
	m1, err1 := CountWordsConcurrent2([]string{"https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md", "https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md"})

	if err1 != nil {
		t.Fatal(err1)
	}

	if m1 == nil {
		t.Fatal("Nil word count map")
	}

	if len(m1) <= 0 {
		t.Fatal("Empty word count map")
	}

	m2, _ := CountWords([]string{"https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md", "https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md"})

	for k, v := range m2 {
		fmt.Printf("%v -> %v \n", k, v)
	}

	if len(m1) != len(m2) {
		log.Fatalf("Length of m1 not equal to m2, expected %v was %v", len(m2), len(m1))
		t.Fatal("Length of m1 not equal to m2")
	}

	if !reflect.DeepEqual(m1, m2) {
		t.Fatal("m1 not equal to m2")
	}
}

func TestSuccessfulWordCountConcurrent3ForMultipleUris(t *testing.T) {
	m1, err1 := CountWordsConcurrent3([]string{"https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md", "https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md"})

	if err1 != nil {
		t.Fatal(err1)
	}

	if m1 == nil {
		t.Fatal("Nil word count map")
	}

	if len(m1) <= 0 {
		t.Fatal("Empty word count map")
	}

	m2, _ := CountWords([]string{"https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md", "https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md"})

	//for k, v := range m2 {
	//	fmt.Printf("%v -> %v \n", k, v)
	//}

	if len(m1) != len(m2) {
		log.Fatalf("Length of m1 not equal to m2, expected %v was %v", len(m2), len(m1))
		t.Fatal("Length of m1 not equal to m2")
	}

	if !reflect.DeepEqual(m1, m2) {

		for k, v := range m1 {
			if v2, ok := m2[k]; !ok || v != v2 {
				log.Fatalf("m1[%v]=%v, m2[%v]=%v", k, v, k, v2)
			}
		}

		t.Fatal("m1 not equal to m2")
	}
}

func TestPerformanceBetweenVersions(t *testing.T) {

	uriSet := make([]string, 6)
	uriSet[0] = "https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md"
	uriSet[1] = "https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md"
	uriSet[2] = "https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md"
	uriSet[3] = "https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md"
	uriSet[4] = "https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md"
	uriSet[5] = "https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md"

	startSerial := time.Now()
	m1, _ := CountWords(uriSet)
	durationSerial := time.Since(startSerial)

	startConcurrent := time.Now()
	m2, _ := CountWordsConcurrent(uriSet)
	durationConcurrent := time.Since(startConcurrent)

	startConcurrent2 := time.Now()
	m3, _ := CountWordsConcurrent2(uriSet)
	durationConcurrent2 := time.Since(startConcurrent2)

	startConcurrent3 := time.Now()
	m4, _ := CountWordsConcurrent3(uriSet)
	durationConcurrent3 := time.Since(startConcurrent3)

	startConcurrent4 := time.Now()
	m5, _ := CountWordsConcurrent3(uriSet)
	durationConcurrent4 := time.Since(startConcurrent4)

	if !reflect.DeepEqual(m1, m2) {
		t.Fatal("m1 not equal to m2")
	}

	if !reflect.DeepEqual(m1, m3) {
		t.Fatal("m1 not equal to m3")
	}

	if !reflect.DeepEqual(m1, m4) {
		t.Fatal("m1 not equal to m4")
	}

	if !reflect.DeepEqual(m1, m5) {
		t.Fatal("m1 not equal to m5")
	}

	if durationSerial < durationConcurrent {
		t.Fatal("Expected concurrent fast than serial, was serial faster than concurrent")
	}

	fmt.Printf("Serial = %v\n", durationSerial)
	fmt.Printf("Concurrent = %v\n", durationConcurrent)
	fmt.Printf("Concurrent2 = %v\n", durationConcurrent2)
	fmt.Printf("Concurrent3 = %v\n", durationConcurrent3)
	fmt.Printf("Concurrent4 = %v\n", durationConcurrent4)

	fmt.Printf("Concurrent is %v%% faster than serial\n", (1.0-float32(durationConcurrent.Milliseconds())/float32(durationSerial.Milliseconds()))*100.0)
	fmt.Printf("Concurrent2 is %v%% faster than serial\n", (1.0-float32(durationConcurrent2.Milliseconds())/float32(durationSerial.Milliseconds()))*100.0)
	fmt.Printf("Concurrent2 is %v%% faster than concurrent\n", (1.0-float32(durationConcurrent2.Milliseconds())/float32(durationConcurrent.Milliseconds()))*100.0)
	fmt.Printf("Concurrent3 is %v%% faster than concurrent2\n", (1.0-float32(durationConcurrent3.Milliseconds())/float32(durationConcurrent2.Milliseconds()))*100.0)
	fmt.Printf("Concurrent4 is %v%% faster than concurrent2\n", (1.0-float32(durationConcurrent4.Milliseconds())/float32(durationConcurrent2.Milliseconds()))*100.0)
}

func TestPerformanceBetweenConcurrent2AndConcurrent3Versions(t *testing.T) {

	uriSet := make([]string, 6)
	uriSet[0] = "https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md"
	uriSet[1] = "https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md"
	uriSet[2] = "https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md"
	uriSet[3] = "https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md"
	uriSet[4] = "https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md"
	uriSet[5] = "https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md"

	diff := make([]int64, 100)

	for i, _ := range diff {
		startConcurrsent1 := time.Now()
		CountWordsConcurrent2(uriSet)
		durationConcurrent1 := time.Since(startConcurrsent1)

		startConcurrent2 := time.Now()
		CountWordsConcurrent3(uriSet)
		durationConcurrent2 := time.Since(startConcurrent2)

		diff[i] = int64(float64(100) - (float64(durationConcurrent2.Milliseconds())/float64(durationConcurrent1.Milliseconds()))*100)
	}

	var sumFaster int64 = 0
	var sumSlower int64 = 0
	prcnt := 0

	for _, v := range diff {

		if v > 0 {
			prcnt += 1
			sumFaster += v
		} else {
			sumSlower += v
		}
	}

	fmt.Printf("Concurrent 3 is faster than Concurrent 2 %v%% of the time on average by %.2f%% \n", prcnt, float64(sumFaster)/float64(prcnt))
	fmt.Printf("Concurrent 3 is slower than Concurrent 2 %v%% of the time on average by %.2f%% \n", 100-prcnt, float64(sumSlower)/float64(prcnt))
}

func TestPerformanceBetweenConcurren3AndConcurrent4Versions(t *testing.T) {

	uriSet := make([]string, 6)
	uriSet[0] = "https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md"
	uriSet[1] = "https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md"
	uriSet[2] = "https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md"
	uriSet[3] = "https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md"
	uriSet[4] = "https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md"
	uriSet[5] = "https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md"

	diff := make([]int64, 100)

	for i, _ := range diff {
		startConcurrsent1 := time.Now()
		CountWordsConcurrent3(uriSet)
		durationConcurrent1 := time.Since(startConcurrsent1)

		startConcurrent2 := time.Now()
		CountWordsConcurrent4(uriSet)
		durationConcurrent2 := time.Since(startConcurrent2)

		diff[i] = int64(float64(100) - (float64(durationConcurrent2.Milliseconds())/float64(durationConcurrent1.Milliseconds()))*100)
	}

	var sumFaster int64 = 0
	var sumSlower int64 = 0
	prcnt := 0

	for _, v := range diff {

		if v > 0 {
			prcnt += 1
			sumFaster += v
		} else {
			sumSlower += v
		}
	}

	fmt.Printf("Concurrent 4 is faster than Concurrent 3 %v%% of the time on average by %.2f%% \n", prcnt, float64(sumFaster)/float64(prcnt))
	fmt.Printf("Concurrent 4 is slower than Concurrent 3 %v%% of the time on average by %.2f%% \n", 100-prcnt, float64(sumSlower)/float64(prcnt))
}

func TestPerformanceBetweenConcurren2AndConcurrent4Versions(t *testing.T) {

	uriSet := make([]string, 6)
	uriSet[0] = "https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md"
	uriSet[1] = "https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md"
	uriSet[2] = "https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md"
	uriSet[3] = "https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md"
	uriSet[4] = "https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md"
	uriSet[5] = "https://raw.githubusercontent.com/dkhanaferov/datasharing/master/README.md"

	diff := make([]int64, 100)

	for i, _ := range diff {
		startConcurrsent1 := time.Now()
		CountWordsConcurrent2(uriSet)
		durationConcurrent1 := time.Since(startConcurrsent1)

		startConcurrent2 := time.Now()
		CountWordsConcurrent4(uriSet)
		durationConcurrent2 := time.Since(startConcurrent2)

		diff[i] = int64(float64(100) - (float64(durationConcurrent2.Milliseconds())/float64(durationConcurrent1.Milliseconds()))*100)
	}

	var sumFaster int64 = 0
	var sumSlower int64 = 0
	prcnt := 0

	for _, v := range diff {

		if v > 0 {
			prcnt += 1
			sumFaster += v
		} else {
			sumSlower += v
		}
	}

	fmt.Printf("Concurrent 4 is faster than Concurrent 2 %v%% of the time on average by %.2f%% \n", prcnt, float64(sumFaster)/float64(prcnt))
	fmt.Printf("Concurrent 4 is slower than Concurrent 2 %v%% of the time on average by %.2f%% \n", 100-prcnt, float64(sumSlower)/float64(prcnt))
}
