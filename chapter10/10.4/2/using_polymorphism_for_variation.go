package upfv

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

func Rating(voyage Voyage, histories []History) string {
	vpf := voyageProfitFactor(voyage, histories)
	vr := voyageRisk(voyage)
	chr := captionHistoryRisk(voyage, histories)
	if vpf*3 > (vr + chr*2) {
		return "A"
	}
	return "B"
}

func voyageRisk(voyage Voyage) int {
	result := 1
	if voyage.Length > 4 {
		result += 2
	}
	if voyage.Length > 8 {
		result += voyage.Length - 8
	}
	if voyage.Zone == "china" || voyage.Zone == "east-indies" {
		result += 4
	}
	return int(math.Max(float64(result), 0))
}

func captionHistoryRisk(voyage Voyage, histories []History) int {
	result := 1
	if len(histories) < 5 {
		result += 4
	}

	for _, v := range histories {
		if v.Profit < 0 {
			result += 1
		}
	}

	// 存在相同的逻辑：都在询问“是否有到中国的航程”以及“船长是否曾去过中国”
	if voyage.Zone == "china" && hasChina(histories) {
		result -= 2
	}
	return int(math.Max(float64(result), 0))
}

func hasChina(histories []History) bool {
	for _, v := range histories {
		if v.Zone == "china" {
			return true
		}
	}
	return false
}

func voyageProfitFactor(voyage Voyage, histories []History) int {
	result := 2
	if voyage.Zone == "china" {
		result += 1
	}
	if voyage.Zone == "east-indies" {
		result += 1
	}
	// 存在相同的逻辑：都在询问“是否有到中国的航程”以及“船长是否曾去过中国”
	if voyage.Zone == "china" && hasChina(histories) {
		result += 3
		if len(histories) > 10 {
			result += 1
		}
		if voyage.Length > 12 {
			result += 1
		}
		if voyage.Length > 18 {
			result -= 1
		}
	} else {
		if len(histories) > 8 {
			result += 1
		}
		if voyage.Length > 14 {
			result -= 1
		}
	}
	return result
}
