package youtube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Kind  string `json:"kind"`
	Items []Item `json:"items"`
}

type Item struct {
	Kind  string `json:"kind"`
	Id    string `json:"id"`
	Stats Stats  `json:"statistics"`
}

type Stats struct {
	Views       string `json:"views"`
	Subscribers string `json:"subscriberCount"`
	Videos      string `json:"videoCount"`
}

func GetSubscribers() (Item, error) {
	var response Response

	req, err := http.NewRequest("GET", "https://www.googleapis.com/youtube/v3/channels", nil)
	if err != nil {
		log.Println(err)
		return Item{}, err
	}

	q := req.URL.Query()
	//log.Println(q)
	q.Add("key", os.Getenv("YOUTUBE_KEY"))
	//log.Println(q)
	q.Add("id", os.Getenv("ID"))
	//log.Println(q)
	q.Add("part", "statistics")
	//log.Println(q)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return Item{}, err
	}
	defer res.Body.Close()

	fmt.Println("Response statis:", res.Status)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return Item{}, err
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println(err)
		return Item{}, err
	}
	return response.Items[0], nil
}
