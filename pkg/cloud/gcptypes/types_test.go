package gcptypes

import (
	"testing"
	"time"
)

func Test_Objects_after(t *testing.T) {
	// empty objects list
	var files0 FilesMetadata
	files0 = files0.After(time.Now())
	if len(files0) != 0 {
		t.Errorf("Objs list len Exp: %d, Got: %d", 0, len(files0))
	}

	// 1 element in list
	files1 := FilesMetadata{FileMetadata{
		Created: time.Date(2018, 7, 1, 0, 0, 0, 0, time.UTC)}}

	res1 := files1.After(
		time.Date(2018, 8, 1, 0, 0, 0, 0, time.UTC))
	if len(res1) != 0 {
		t.Errorf("Objs list len Exp: %d, Got: %d", 0, len(res1))
	}

	res1 = files1.After(
		time.Date(2018, 6, 1, 0, 0, 0, 0, time.UTC))
	if len(res1) != 1 {
		t.Errorf("Objs list len Exp: %d, Got: %d", 1, len(res1))
	}
}
