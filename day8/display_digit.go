package main

type DisplayDigit map[rune]interface{}

func NewDisplayDigit() DisplayDigit {
	return make(map[rune]interface{})
}

func FromSectors(sectors string) DisplayDigit {
	dd := NewDisplayDigit()
	for _, sector := range sectors {
		dd[sector] = nil
	}
	return dd
}

func (dd DisplayDigit) Add(other DisplayDigit) DisplayDigit {
	sum := NewDisplayDigit()
	for sector := range dd {
		sum[sector] = nil
	}
	for sector := range other {
		sum[sector] = nil
	}
	return sum
}

func (dd DisplayDigit) Subtract(other DisplayDigit) DisplayDigit {
	diff := NewDisplayDigit()
	for sector := range dd {
		if other.isOn(sector) {
			continue
		}
		diff[sector] = nil
	}
	return diff
}

func (dd DisplayDigit) Contains(other DisplayDigit) bool {
	for sector := range other {
		if !dd.isOn(sector) {
			return false
		}
	}
	return true
}

func (dd DisplayDigit) isOn(sector rune) bool {
	_, ok := dd[sector]
	return ok
}

func (dd DisplayDigit) Intersection(other DisplayDigit) DisplayDigit {
	common := NewDisplayDigit()
	for sector := range dd {
		if !dd.isOn(sector) || !other.isOn(sector) {
			continue
		}
		common[sector] = nil
	}
	return common
}

func (dd DisplayDigit) Equal(other DisplayDigit) bool {
	if len(dd) != len(other) {
		return false
	}

	for sector := range dd {
		if !other.isOn(sector) {
			return false
		}
	}

	return true
}
