package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func fetchCourses() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.udemy.com/api-2.0/courses/", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer YOUR_ACCESS_TOKEN")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}

func main() {
	fetchCourses()
}
