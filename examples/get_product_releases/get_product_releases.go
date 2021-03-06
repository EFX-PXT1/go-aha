package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/joho/godotenv"

	au "github.com/grokify/go-aha/ahautil"
)

func main() {
	err := godotenv.Load(os.Getenv("ENV_PATH"))
	if err != nil {
		log.Fatal("$ENV_PATH not found")
	}

	apis := au.NewClientAPIs(os.Getenv("AHA_ACCOUNT"), os.Getenv("AHA_API_KEY"))
	ctx := context.Background()

	rs := au.NewReleaseSet()
	rs.ClientAPIs = apis
	products := []string{"PROD"}
	for _, prod := range products {
		err := rs.LoadReleasesForProduct(ctx, prod)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmtutil.PrintJSON(rs.ReleaseMap)

	fmt.Println(len(rs.ReleaseMap))

	dta, err := time.Parse(time.RFC3339, "2017-10-01T00:00:00Z")
	if err != nil {
		log.Fatal(err)
	}
	dtz, err := time.Parse(time.RFC3339, "2018-12-31T23:59:59Z")
	if err != nil {
		log.Fatal(err)
	}

	rels, err := rs.GetReleasesForDates(ctx, dta, dtz)
	if err != nil {
		log.Fatal(err)
	}
	fmtutil.PrintJSON(rels)
	fmt.Println(len(rels))
	fmt.Println("DONE")
}
