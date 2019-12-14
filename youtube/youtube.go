package youtube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	baseURL = "https://www.googleapis.com/youtube/v3/channels"
)

//SetURLQuery sets url for http request
func SetURLQuery(r *http.Request) {
	q := r.URL.Query()
	q.Add("key", os.Getenv("YOUTUBE_KEY"))
	q.Add("id", os.Getenv("ID"))
	q.Add("part", "statistics")

	r.URL.RawQuery = q.Encode()
}

//GetSubscribers gets info from api
func GetSubscribers() (Items, error) {

	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		log.Println(err)
		return Items{}, err
	}

	SetURLQuery(req)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return Items{}, err
	}
	defer res.Body.Close()

	fmt.Println("Response statis:", res.Status)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return Items{}, err
	}

	var response Response
	if err = json.Unmarshal(body, &response); err != nil {
		log.Println(err)
		return Items{}, err
	}

	return response.Items[0], nil
}
