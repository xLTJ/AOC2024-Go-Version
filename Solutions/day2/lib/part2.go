package lib

import "utils"

// CountSafeReports2 counts how many reports are safe from an input file, but now its actually fine if one
// number can be removed to make it safe
func CountSafeReports2(fileName string) (int, error) {
	reports, err := utils.ReadAndParseLines(fileName, ParseLine)
	if err != nil {
		return 0, err
	}

	safeReports := 0
	for _, report := range reports {
		isSafe, err := CheckSafety(report)
		if err != nil {
			return 0, err
		}

		if isSafe {
			safeReports++
			continue
		}

		// if report isnt safe, try removing one number at a time and check again for each of them
		for i := range report {
			modifiedReport := make([]int, 0, len(report)-1)
			modifiedReport = append(modifiedReport, report[:i]...)
			modifiedReport = append(modifiedReport, report[i+1:]...)

			isModifiedSafe, err := CheckSafety(modifiedReport)
			if err != nil {
				return 0, err
			}

			if isModifiedSafe {
				safeReports++
				break
			}
		}
	}
	return safeReports, nil
}
