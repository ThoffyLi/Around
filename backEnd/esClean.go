package main

import (
	elastic "gopkg.in/olivere/elastic.v3"
)

const (
	ES_URL = "http://35.223.74.152:9200"
	INDEX  = "around"
)

/* Delete index "around"  */
func main() {
	// Create a client
	client, err := elastic.NewClient(elastic.SetURL(ES_URL), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	deleteIndex, err := client.DeleteIndex(INDEX).Do()
	if err != nil {
		// Handle error
		panic(err)
	}
	if !deleteIndex.Acknowledged {
		// Not acknowledged
	}

}
