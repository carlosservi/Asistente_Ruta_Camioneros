package route

func restAreaEqual(ra1, ra2 RestArea) bool {
	// Compara los campos uno por uno
	if ra1.Id != ra2.Id {
		return false
	}

	if ra1.Retraso != ra2.Retraso {
		return false
	}

	// Compara las listas de OpeningHours
	if len(ra1.OpeningHours) != len(ra2.OpeningHours) {
		return false
	}

	for i := range ra1.OpeningHours {
		if !openingHoursEqual(ra1.OpeningHours[i], ra2.OpeningHours[i]) {
			return false
		}
	}

	return true
}

func openingHoursEqual(oh1, oh2 OpeningHours) bool {
	// Compara los campos uno por uno
	if oh1.StartTime != oh2.StartTime {
		return false
	}

	if oh1.CloseTime != oh2.CloseTime {
		return false
	}

	if oh1.Weekday != oh2.Weekday {
		return false
	}

	return true
}
