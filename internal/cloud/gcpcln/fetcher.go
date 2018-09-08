package gcpcln

import (
	"log"
	"time"
)

// Fetch periodically fetches data into db from GCP
func (c *Client) Fetch(bktName, prefix string, accountID int, dt time.Duration) {

	// pull all csv objects from bucket
	objs, err := c.CsvObjsList(bktName, prefix)
	if err != nil {
		log.Fatalf("in fetch c.CsvObjList: %v", err)
	}
	log.Println(objs)

	// select last report from db with its creation time
	lstFile, err := c.pgCln.GetLastFile()
	if err != nil {
		log.Printf("in fetch c.pgCln.SelLstGcpCsvObjectl: %v", err)
	}
	log.Println(lstFile)

	// filter CsvObjList, save only fresh objects
	if lstFile != nil {
		objs = objs.After(lstFile.Created)
	}

	// // make Reports from object content
	reps, err := c.MakeReports(objs)
	if err != nil {
		log.Fatalf("in fetch MakeReports: %v", err)
	}
	log.Println(reps)

	// write Reports into db
	err = c.pgCln.AddReportsToAccount(reps, accountID)
	if err != nil {
		log.Fatalf("in fetch AddReportsToAccount: %v", err)
	}
}
