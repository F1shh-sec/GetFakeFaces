package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

//{"generated":"true","src":"\/img\/avatar-09c3a410acbd8ab37fe7128688f803f6.jpg","name":"avatar-09c3a410acbd8ab37fe7128688f803f6.jpg"}
type imageLocation struct {
	Generated bool   `json:"generated"`
	Src       string `json:"src"`
	Name      string `json:"name`
}

var URL = "https://this-person-does-not-exist.com/en?new"

var wg = sync.WaitGroup{}

func main() {
	for i := 1; i < 1000; i++ {
		time.Sleep(time.Duration(1000) * time.Millisecond)
		go makeRequest(i)
	}
	wg.Wait()
}

func makeRequest(i int) {
	wg.Add(1)
	filename := "FILE" + strconv.Itoa(i) + ".json"
	url := URL
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error Getting URL")
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		fmt.Println("PAGE DID NOT 200")
	}
	byteValue, _ := ioutil.ReadAll(response.Body)
	var location imageLocation
	json.Unmarshal(byteValue, &location)
	fmt.Println(location.Name)

	Imageurl := "https://this-person-does-not-exist.com/img/" + location.Name
	imageresponse, err := http.Get(Imageurl)
	if err != nil {
		fmt.Println("Error Getting URL")
	}
	defer imageresponse.Body.Close()
	if imageresponse.StatusCode != 200 {
		fmt.Println("PAGE DID NOT 200")
	}
	imagefilename := "faces/person" + strconv.Itoa(i) + ".jpg"
	imagefile, err := os.Create(imagefilename)
	if err != nil {
		fmt.Println("Error Creating File " + imagefilename)
	}
	_, err = io.Copy(imagefile, imageresponse.Body)
	if err != nil {
		fmt.Println("Error Copying File")
	}
	imagefile.Close()
	os.Remove(filename)
	defer wg.Done()
}
