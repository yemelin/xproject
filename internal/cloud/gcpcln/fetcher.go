package gcpcln

import (
	"fmt"
	"log"
	"time"
)

// Fetch periodically fetches data into db from GCP
func (c *Client) Fetch(bktName, prefix string, dt time.Duration) {

	// pull all csv objects from bucket
	objs, err := c.CsvObjsList(bktName, prefix)
	if err != nil {
		log.Fatalf("in fetch c.CsvObjList: %v", err)
	}

	// select last report from db with its creation time
	lstFile, err := c.pgCln.GetLastFile()
	if err != nil {
		log.Printf("in fetch c.pgCln.SelLstGcpCsvObjectl: %v", err)
	}
	fmt.Println(lstFile)

	// filter CsvObjList, save only fresh objects
	if lstFile != nil {
		objs = objs.After(lstFile.Created)
	}

	// make Reports from object content
	reps, err := c.MakeReports(objs)
	if err != nil {
		log.Fatalf("in fetch MakeReports: %v", err)
	}
	fmt.Println(reps)

	// write Reports into db
	// c.pgCln.InsReps(reps)
	// c.pgCln.InsObjs(pgcln.Objects(make(Objects, 10)))

	// time.Sleep(dt * time.Hour)
	// }
}
