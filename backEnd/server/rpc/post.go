package rpc

import (
	"../constants"
	"../entity"
	"../es_save"
	"../gcs_save"
	"context"
	"fmt"
	"github.com/pborman/uuid"
	elastic "gopkg.in/olivere/elastic.v3"
	"net/http"

	"strconv"
	"strings"
	// auth
	"github.com/dgrijalva/jwt-go"
)

var (
	mediaTypes = map[string]string{
		"jpeg": "image",
		"jpg":  "image",
		"gif":  "image",
		"png":  "image",
		"mov":  "video",
		"mp4":  "video",
		"avi":  "video",
		"flv":  "video",
		"wmv":  "video",
		"mkv":  "video",
	}
)

func handlerPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
	if r.Method == "OPTIONS" {
		return
	}
	user := r.Context().Value("user")
	claims := user.(*jwt.Token).Claims
	username := claims.(jwt.MapClaims)["username"]

	// Parse from multi form data
	r.ParseMultipartForm(32 << 20)
	fmt.Printf("Received one post request %s\n", r.FormValue("message"))
	lon, _ := strconv.ParseFloat(r.FormValue("lon"), 64)
	lat, _ := strconv.ParseFloat(r.FormValue("lat"), 64)

	id := uuid.New()
	// file type mapping: image or video
	file, header, err := r.FormFile("image")
	filename := header.Filename

	p := &Post{
		User:    username.(string),
		Message: r.FormValue("message"),
		Location: Location{
			Lon: lon,
			Lat: lat,
		},
		Type: mediaTypes[strings.Split(filename, `.`)[1]],
	}

	if err != nil {
		http.Error(w, "GCS is not setup", http.StatusInternalServerError)
		fmt.Printf("GCS is not setup %v\n", err)
		panic(err)
	}
	defer file.Close()
	ctx := context.Background()
	_, attrs, err := saveToGCS(ctx, file, BUCKET_NAME, id)
	if err != nil {
		http.Error(w, "GCS is not setup", http.StatusInternalServerError)
		fmt.Printf("GCS is not setup %v\n", err)
		panic(err)
	}
	p.Url = attrs.MediaLink
	//save to ES
	saveToES(p, id)

	/* save to BigTable*/
	//saveToBigTable(p, id)

}
