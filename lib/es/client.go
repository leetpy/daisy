package es

import (
	"gopkg.in/olivere/elastic.v5"
)

func GetClient() *elastic.Client {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		panic(err)
	}
	return client
}
