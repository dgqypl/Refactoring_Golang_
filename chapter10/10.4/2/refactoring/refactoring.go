package refactoring

import (
	"math"
)

type Voyage struct {
	Zone   string `json:"zone"`
	Length int    `json:"length"`
}

type History struct {
	Zone   string `json:"zone"`
	Profit int    `json:"profit"`
}

type rating struct {
	voyage    Voyage
	histories []History
}

func newRating(voyage Voyage, histories []History) *rating {
	return &rating{
		voyage:    voyage,
		histories: histories,
	}
}

func (r *rating) voyageRisk() int {
	result := 1
	if r.voyage.Length > 4 {
		result += 2
	}
	if r.voyage.Length > 8 {
		result += r.voyage.Length - 8
	}
	if r.voyage.Zone == "china" || r.voyage.Zone == "east-indies" {
		result += 4
	}
	return int(math.Max(float64(result), 0))
}

func (r *rating) captionHistoryRisk() int {
	result := 1
	if len(r.histories) < 5 {
		result += 4
	}

	for _, v := range r.histories {
		if v.Profit < 0 {
			result += 1
		}
	}

	return int(math.Max(float64(result), 0))
}

type lengthFactor interface {
	historyLengthFactor() int
	voyageLengthFactor() int
}

func (r *rating) voyageProfitFactor(p lengthFactor) int {
	result := 2
	if r.voyage.Zone == "china" {
		result += 1
	}
	if r.voyage.Zone == "east-indies" {
		result += 1
	}
	result += p.historyLengthFactor()
	result += p.voyageLengthFactor()

	return result
}

func (r *rating) historyLengthFactor() int {
	if len(r.histories) > 8 {
		return 1
	}
	return 0
}

func (r *rating) voyageLengthFactor() int {
	if r.voyage.Length > 14 {
		return -1
	}
	return 0
}

type experiencedChinaRating struct {
	rating
}

func newExperiencedChinaRating(voyage Voyage, histories []History) *experiencedChinaRating {
	this := &experiencedChinaRating{}
	this.voyage = voyage
	this.histories = histories
	return this
}

func (e *experiencedChinaRating) captionHistoryRisk() int {
	result := e.rating.captionHistoryRisk() - 2
	return int(math.Max(float64(result), 0))
}

func (e *experiencedChinaRating) voyageProfitFactor(p lengthFactor) int {
	return e.rating.voyageProfitFactor(p) + 3
}

func (e *experiencedChinaRating) voyageLengthFactor() int {
	result := 0
	if e.voyage.Length > 12 {
		result += 1
	}
	if e.voyage.Length > 18 {
		result -= 1
	}
	return result
}

func (e *experiencedChinaRating) historyLengthFactor() int {
	if len(e.histories) > 10 {
		return 1
	}
	return 0
}

func hasChina(histories []History) bool {
	for _, v := range histories {
		if v.Zone == "china" {
			return true
		}
	}
	return false
}

func Rating(voyage Voyage, histories []History) string {
	if voyage.Zone == "china" && hasChina(histories) {
		e := newExperiencedChinaRating(voyage, histories)
		return calculateValue(e.voyageProfitFactor, e.voyageRisk, e.captionHistoryRisk, e)
	}
	r := newRating(voyage, histories)
	return calculateValue(r.voyageProfitFactor, r.voyageRisk, r.captionHistoryRisk, r)
}

func calculateValue(voyageProfitFactor func(p lengthFactor) int, voyageRisk func() int, captionHistoryRisk func() int, p lengthFactor) string {
	vpf := voyageProfitFactor(p)
	vr := voyageRisk()
	chr := captionHistoryRisk()
	if vpf*3 > (vr + chr*2) {
		return "A"
	}
	return "B"
}
