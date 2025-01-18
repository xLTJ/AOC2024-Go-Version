package lib

import (
	"slices"
	"utils"
)

func SortAndGetMiddleSum(fileName string) (int, error) {
	rules, updates, err := utils.ParseInputTwoParts(fileName, seperator, parseRule, parseUpdate)

	ruleMap := map[int][]int{}

	for _, rule := range rules {
		ruleMap[rule[0]] = append(ruleMap[rule[0]], rule[1])
	}

	sum := 0
	for _, update := range updates {
		if isUpdateValid(update, ruleMap) {
			continue
		}

		slices.SortFunc(update, func(a, b int) int {
			if slices.Contains(ruleMap[a], b) {
				return -1
			}
			return 1
		})

		sum += update[len(update)/2]
	}

	return sum, err
}
