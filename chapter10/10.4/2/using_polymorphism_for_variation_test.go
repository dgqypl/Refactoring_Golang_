package upfv

import "testing"

func TestRating(t *testing.T) {
	voyage := Voyage{
		Zone:   "west-indies",
		Length: 10,
	}
	histories := []History{
		{
			Zone:   "east-indies",
			Profit: 5,
		},
		{
			Zone:   "west-indies",
			Profit: 15,
		},
		{
			Zone:   "china",
			Profit: -2,
		},
		{
			Zone:   "west-africa",
			Profit: 7,
		},
	}
	if ans := Rating(voyage, histories); ans != "B" {
		t.Fatalf("expected B, but %s got", ans)
	}
}
