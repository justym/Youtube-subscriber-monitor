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
	Kind  string  `json:"kind"`
	Items []Items `json:"items"`
}

type Items struct {
	Kind  string `json:"kind"`
	Id    string `json:"id"`
	Stats Stats  `json:"statistics"`
}

type Stats struct {
	Views       string `json:"views"`
	Subscribers string `json:"subscriberCount"`
	Videos      string `json:"videoCount"`
}

func GetSubscribers() (Items, error) {
	var response Response

	req, err := http.NewRequest("GET", "https://www.googleapis.com/youtube/v3/channels", nil)
	if err != nil {
		log.Println("34")
		log.Println(err)
		return Items{}, err
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
		log.Println("51")
		log.Println(err)
		return Items{}, err
	}
	defer res.Body.Close()

	fmt.Println("Response statis:", res.Status)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("L60")
		log.Println(err)
		return Items{}, err
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println(err)
		return Items{}, err
	}
	return response.Items[0], nil
}
