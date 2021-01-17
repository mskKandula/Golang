package main

import (
	"fmt"

	"github.com/tidwall/gjson"
)

func main() {

	objectArray := `[
    {
        "id": "38",
        "article_id": "16",
        "news_event": "625",
        "language": "en",
        "channel_partner_id": "625",
        "title": "Test",
        "show_logo": null,
        "description": "test\n\n",
        "schedule": null,
        "event_date": "2012-03-09 10:08:35",
        "link_text": null,
        "guid": null,
        "timestamp": "2012-03-09 11:19:42",
        "website": null,
        "show_hours": null,
        "page_text": null
    },
    {
        "id": "37",
        "article_id": "15",
        "news_event": "625",
        "language": "en",
        "channel_partner_id": "625",
        "title": "Test",
        "show_logo": null,
        "description": "test\n\n",
        "schedule": null,
        "event_date": "2012-03-09 10:08:35",
        "link_text": null,
        "guid": null,
        "timestamp": "2012-03-09 11:19:39",
        "website": null,
        "show_hours": null,
        "page_text": null
    },
    {
        "id": "36",
        "article_id": "13",
        "news_event": "625",
        "language": "en",
        "channel_partner_id": "625",
        "title": "Test",
        "show_logo": null,
        "description": "test\n\n",
        "schedule": null,
        "event_date": "2012-03-09 10:08:35",
        "link_text": null,
        "guid": null,
        "timestamp": "2012-03-09 11:19:35",
        "website": null,
        "show_hours": null,
        "page_text": null
    },
    {
        "id": "35",
        "article_id": "13",
        "news_event": "625",
        "language": "en",
        "channel_partner_id": "625",
        "title": "Test",
        "show_logo": null,
        "description": "test\n\n",
        "schedule": null,
        "event_date": "2012-03-09 10:08:35",
        "link_text": null,
        "guid": null,
        "timestamp": "2012-03-09 11:19:31",
        "website": null,
        "show_hours": null,
        "page_text": null
    }
]
`

	count := make(map[string]uint8)

	jsonArray := gjson.Parse(objectArray).Array()

	for _, val := range jsonArray {
		article := val.Get("article_id").String()
		count[article]++
	}

	for key, value := range count {
		fmt.Printf("%s num of times is %d \n", key, value)
	}

}
