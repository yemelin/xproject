package gcpcln

import "time"

type Object struct {
	Name    string
	Bucket  string
	Created time.Time
}

type Objects []Object

func (objs Objects) after(t time.Time) (res Objects) {
	for _, o := range objs {
		if o.Created.After(t) {
			res = append(res, o)
		}
	}

	return res
}
