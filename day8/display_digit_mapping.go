package main

func (ddm DisplayDigitMapping) MapPattern(pattern string) int {
	soughtDd := FromSectors(pattern)
	for i, dd := range ddm {
		if dd.Equal(soughtDd) {
			return i
		}
	}
	return -1
}

func deductMapping(p Patterns) DisplayDigitMapping {
	sm := make([]DisplayDigit, 10)

	// map simple ones first
	for _, pattern := range p {
		dd := FromSectors(pattern)
		switch len(pattern) {
		case 2: // 1
			sm[1] = dd
		case 3: // 7
			sm[7] = dd
		case 4: // 4
			sm[4] = dd
		case 7: // 8
			sm[8] = dd
		}
	}

	bottomLeftCorner := sm[8].Subtract(sm[4]).Subtract(sm[7])
	middleLeftCorner := sm[4].Subtract(sm[1])

	// guess the rest
	for _, pattern := range p {
		switch len(pattern) {
		case 2, 3, 4, 7:
			continue
		}
		dd := FromSectors(pattern)
		switch {
		case dd.Contains(sm[7]) && dd.Contains(bottomLeftCorner):
			sm[0] = dd
		case dd.Contains(bottomLeftCorner) && dd.Contains(middleLeftCorner):
			sm[6] = dd
		case dd.Contains(bottomLeftCorner) && !dd.Contains(middleLeftCorner):
			sm[2] = dd
		case !dd.Contains(bottomLeftCorner) && !dd.Contains(middleLeftCorner):
			sm[3] = dd
		case !dd.Contains(bottomLeftCorner) && dd.Contains(middleLeftCorner) && !dd.Contains(sm[4]):
			sm[5] = dd
		case !dd.Contains(bottomLeftCorner) && dd.Contains(middleLeftCorner) && dd.Contains(sm[4]):
			sm[9] = dd
		}
	}

	return sm
}
