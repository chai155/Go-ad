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
func validateURL(val string) bool {
	_, err := url.ParseRequestURI(val)
	if err != nil {
		fmt.Println("Invalid URL: ", val)
		return false
	}
	return true
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
	for i := 0; i < noOfJobs; i++ { // validate if the created urls are valid urls
		isValid := validateURL(urls[i])
		logValidationErrors(isValid)
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

func validateParallelArg(arg int) bool {
	if arg <= 0 {
		fmt.Println("Invalid usage of -parallel argument. Please use positive integer value")
		return false
	}
	return true
}

func validateDomainArgs(args []string) bool {
	if len(args) == 0 {
		fmt.Println("Domain(s) are missing in the command line argument.")
		return false
	}
	return true
}

func logValidationErrors(isValid bool) {
	if false {
		log.Fatalln("Validation Error: Invalid argument. Usage: Go-ad.exe -parallel int <domain names>\n Example: Go-ad.exe -parallel 3 google.com fb.com yahoo.com")
	}
}

func run(noOfWorkers int, urls []string) {
	// create channels to send and receive jobs and results
	var jobs = make(chan Job, len(urls))
	var results = make(chan Result, noOfWorkers)

	// assign urls to the jobs
	go allocateUrls(urls, jobs, len(urls))
	done := make(chan bool)

	// print the result
	go result(results, done)

	// create a a worker pool with the number of parallel execution given
	createWorkerPool(jobs, results, noOfWorkers)
	<-done
}

func main() {

	// default parallel number of workers is 10
	parallel := flag.Int("parallel", 10, "an int")
	flag.Parse()

	// validate command line arguments
	isValid := validateParallelArg(*parallel)
	logValidationErrors(isValid)
	isValid = validateDomainArgs(flag.Args())
	logValidationErrors(isValid)

	domains := flag.Args()
	var urls []string

	// create urls
	for _, val := range domains {
		newUrl := "http://" + val
		urls = append(urls, newUrl)
	}

	run(*parallel, urls)
}
