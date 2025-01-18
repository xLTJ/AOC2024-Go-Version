package lib

import (
	"slices"
	"utils"
)

const seperator = ""

func GetMiddleSum(fileName string) (int, error) {
	rules, updates, err := utils.ParseInputTwoParts(fileName, seperator, parseRule, parseUpdate)

	ruleMap := map[int][]int{}

	for _, rule := range rules {
		ruleMap[rule[0]] = append(ruleMap[rule[0]], rule[1])
	}

	sum := 0
	for _, update := range updates {
		if !isUpdateValid(update, ruleMap) {
			continue
		}
		sum += update[len(update)/2]
	}

	return sum, err
}

func parseRule(rule string) ([]int, error) {
	return utils.StringToIntSlice(rule, "|")
}

func parseUpdate(update string) ([]int, error) {
	return utils.StringToIntSlice(update, ",")
}

func isUpdateValid(update []int, ruleMap map[int][]int) bool {
	for i, number := range update {
		for _, prevNumber := range update[:i] {
			if slices.Contains(ruleMap[number], prevNumber) {
				return false
			}
		}
	}
	return true
}
