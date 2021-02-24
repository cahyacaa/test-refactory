package main

import (
	"encoding/json"
	"fmt"
	// "time"
	// "reflect"
)

type UserData struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Profile  struct {
		FullName string   `json:"full_name"`
		Birthday string   `json:"birthday"`
		Phones   []string `json:"phones"`
	} `json:"profile"`
	Articles []struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		PublishedAt string `json:"published_at"`
	} `json:"articles:"`
}

var file = []byte(`[
	{
	  "id": 323,
	  "username": "rinood30",
	  "profile": {
		"full_name": "Shabrina Fauzan",
		"birthday": "1988-10-30",
		"phones": [
		  "08133473821",
		  "082539163912"
		]
	  },
	  "articles:": [
		{
		  "id": 3,
		  "title": "Tips Berbagi Makanan",
		  "published_at": "2019-01-03T16:00:00"
		},
		{
		  "id": 7,
		  "title": "Cara Membakar Ikan",
		  "published_at": "2019-01-07T14:00:00"
		}
	  ]
	},
	{
	  "id": 201,
	  "username": "norisa",
	  "profile": {
		"full_name": "Noor Annisa",
		"birthday": "1986-08-14",
		"phones": []
	  },
	  "articles:": [
		{
		  "id": 82,
		  "title": "Cara Membuat Kue Kering",
		  "published_at": "2019-10-08T11:00:00"
		},
		{
		  "id": 91,
		  "title": "Cara Membuat Brownies",
		  "published_at": "2019-11-11T13:00:00"
		},
		{
		  "id": 31,
		  "title": "Cara Membuat Brownies",
		  "published_at": "2019-11-11T13:00:00"
		}
	  ]
	},
	{
	  "id": 42,
	  "username": "karina",
	  "profile": {
		"full_name": "Karina Triandini",
		"birthday": "1986-04-14",
		"phones": [
		  "06133929341"
		]
	  },
	  "articles:": []
	},
	{
	  "id": 201,
	  "username": "icha",
	  "profile": {
		"full_name": "Annisa Rachmawaty",
		"birthday": "1987-12-30",
		"phones": []
	  },
	  "articles:": [
		{
		  "id": 39,
		  "title": "Tips Berbelanja Bulan Tua",
		  "published_at": "2019-04-06T07:00:00"
		},
		{
		  "id": 43,
		  "title": "Cara Memilih Permainan di Steam",
		  "published_at": "2019-06-11T05:00:00"
		},
		{
		  "id": 58,
		  "title": "Cara Membuat Brownies",
		  "published_at": "2019-09-12T04:00:00"
		}
	  ]
	}
  ]`)

func filterPhone(x []*UserData) {
	var x2 []*UserData
	for _, v := range x {
		if len(v.Profile.Phones) == 0 {
			x2 = append(x2, v)
		}
	}
}

func articlesnotNull(x []*UserData) {
	var x2 []*UserData
	for _, v := range x {
		if len(v.Articles) >0 {
			x2 = append(x2, v)
		}
	}
	data, _ := json.MarshalIndent(x2, "", " ")
	fmt.Println(string(data))
}

func main() {
		var x []*UserData
		err := json.Unmarshal(file, &x)
		if err != nil {
			panic(err)
		}
		// filterPhone(x)
		articlesnotNull(x)

}