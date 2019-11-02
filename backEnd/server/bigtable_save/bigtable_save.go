package bigtable_save

import (
	"../constants"
	"../entity"
	"cloud.google.com/go/bigtable"
	"context"
	"fmt"
	"strconv"
)

func saveToBigTable(p *Post, id string) {
	ctx := context.Background()
	bt_client, err := bigtable.NewClient(ctx, PROJECT_ID, BT_INSTANCE)
	if err != nil {
		panic(err)
		return
	}
	// open table "post"
	tbl := bt_client.Open("post")
	mut := bigtable.NewMutation()
	// timestamp
	t := bigtable.Now()
	mut.Set("post", "user", t, []byte(p.User))
	mut.Set("post", "message", t, []byte(p.Message))
	mut.Set("location", "lat", t, []byte(strconv.FormatFloat(p.Location.Lat, 'f', 1, 64)))
	mut.Set("location", "lon", t, []byte(strconv.FormatFloat(p.Location.Lon, 'f', 1, 64)))
	err = tbl.Apply(ctx, id, mut)
	if err != nil {
		panic(err)
		return
	}
	fmt.Printf("Post is saved to BigTable: %s\n", p.Message)
}
