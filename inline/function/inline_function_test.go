package function

import (
	"testing"
)

type ratingCase struct {
	Driver   Driver
	Expected int
}

func createRatingTestCase(t *testing.T, r *ratingCase) {
	t.Helper()
	if ans := Rating(r.Driver); ans != r.Expected {
		t.Fatalf("%+v expected %d, but %d got",
			r.Driver, r.Expected, ans)
	}

}

func TestRating(t *testing.T) {
	createRatingTestCase(t, &ratingCase{
		Driver:   Driver{NumberOfLateDeliveries: 6},
		Expected: 2,
	})
	createRatingTestCase(t, &ratingCase{
		Driver:   Driver{NumberOfLateDeliveries: 4},
		Expected: 1,
	})
}
