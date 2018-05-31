package storageutil

import (
	"testing"
	"time"
)

func TestSelectWithPrefix(t *testing.T) {
	objects := Objects{
		Object{"obj-1", time.Time{}},
		Object{"obj_2", time.Time{}},
		Object{"obj-3", time.Time{}},
	}

	expObjects := Objects{
		Object{"obj-1", time.Time{}},
		Object{"obj-3", time.Time{}},
	}

	resObjects := objects.SelectWithPrefix("obj-")

	if !isObjectSlicesEqual(resObjects, expObjects) {
		t.Error("IsObjectSlicesEqual\nresult:", resObjects, "\nexpected:", expObjects)
	}
}

func TestSelectObjectsWithFromToTime(t *testing.T) {

	objects := Objects{
		Object{"obj-1", time.Date(2017, time.February, 1, 0, 0, 0, 0, time.Local)},
		Object{"obj-2", time.Date(2017, time.April, 1, 0, 0, 0, 0, time.Local)},
		Object{"obj-3", time.Date(2018, time.January, 1, 0, 0, 0, 0, time.Local)},
		Object{"obj-4", time.Date(2018, time.May, 1, 0, 0, 0, 0, time.Local)},
		Object{"obj-5", time.Date(2020, time.April, 1, 0, 0, 0, 0, time.Local)},
	}

	_, err := objects.SelectInTimeRange(
		time.Date(2018, time.April, 1, 0, 0, 0, 0, time.Local),
		time.Date(2017, time.April, 1, 0, 0, 0, 0, time.Local))

	if err == nil {
		t.Error("err is nil, expected not nil")
	}

	resObjects, err := objects.SelectInTimeRange(
		time.Date(2017, time.March, 1, 0, 0, 0, 0, time.Local),
		time.Date(2018, time.April, 1, 0, 0, 0, 0, time.Local),
	)

	expObjects := Objects{
		Object{"obj-2", time.Date(2017, time.April, 1, 0, 0, 0, 0, time.Local)},
		Object{"obj-3", time.Date(2018, time.January, 1, 0, 0, 0, 0, time.Local)},
	}

	if err != nil {
		t.Error("err is not nil, expected nil")
	}

	if !isObjectSlicesEqual(resObjects, expObjects) {
		t.Error("IsObjectSlicesEqual\nresult:", resObjects, "\nexpected:", expObjects)
	}

}
