package gcpcln

import (
	"log"
	"time"
)

// Fetch periodically fetches data into db from GCP
func (c *Client) Fetch(bktName, prefix string, dt time.Duration) {
	for {
		// pull all csv objects from bucket
		objs, err := c.CsvObjsList(bktName, prefix)
		if err != nil {
			log.Fatalf("in fetch CsvObjList: %v", err)
		}

		// select last report from db with its creation time
		// TODO: use pgcln here to select last report from db

		// filter CsvObjList, save only fresh objects
		// TODO: objs := objs.after(lastReport.Object.Created)

		// make Reports from object content
		reps, err := c.MakeReports(objs)
		_ = reps
		if err != nil {
			log.Fatalf("in fetch MakeReports: %v", err)
		}

		// write Reports into db
		// TODO: use pgcln here to write parsed csv into db

		time.Sleep(dt * time.Hour)
	}
}
