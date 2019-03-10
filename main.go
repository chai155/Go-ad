package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"net/url"
	"sync"
)

// structures to store the jobs and results
type Result struct {
	job     Job
	md5Hash string
}
type Job struct {
	id     int
	urlVal string
}

// validate if the string is valid url or nots
func validateURL(urls []string) {
	for _, val := range urls {
		_, err := url.ParseRequestURI(val)
		if err != nil {
			log.Fatalln("Validation Error: Invalid url: " + val)
		}
	}
}

// create and return md5 hash for the given string
func getMD5Hash(text string) string {
	hash := md5.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

// convert the url to md5 hash
func worker(jobs chan Job, results chan Result, wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, getMD5Hash(job.urlVal)}
		results <- output
	}
	wg.Done()
}

// allocate the urls to the jobs
func allocateUrls(urls []string, jobs chan Job, noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		job := Job{i, urls[i]}
		jobs <- job
	}
	close(jobs)
}

// worker pool to execute the job parallely
func createWorkerPool(jobs chan Job, results chan Result, noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}
	wg.Wait()
	close(results)
}

// print the result
func result(results chan Result, done chan bool) {
	for result := range results {
		fmt.Println(result.job.urlVal, result.md5Hash)
	}
	done <- true
}

func validateParallelArg(arg int) {
	if arg <= 0 {
		log.Fatalln("Validation Error: parallel argument value has be a positive integer value")
	}
}

func validateDomainArgs(args []string) {
	if len(args) == 0 {
		log.Fatalln("Validation Error: Could not find domain string argument")
	}
}

func main() {

	// default parallel number of workers is 10
	parallel := flag.Int("parallel", 10, "an int")
	flag.Parse()

	// validate command line arguments
	validateParallelArg(*parallel)
	validateDomainArgs(flag.Args())

	noOfWorkers := *parallel
	domains := flag.Args()
	var urls []string

	// create urls
	for _, val := range domains {
		newUrl := "http://" + val
		urls = append(urls, newUrl)
	}

	// validate if the created urls are valid urls
	validateURL(urls)

	// create channels to send and receive jobs and results
	var jobs = make(chan Job, len(urls))
	var results = make(chan Result, *parallel)

	// assign urls to the jobs
	go allocateUrls(urls, jobs, len(urls))
	done := make(chan bool)

	// print the result
	go result(results, done)

	// create a a worker pool with the number of parallele execution given
	createWorkerPool(jobs, results, noOfWorkers)
	<-done
}
