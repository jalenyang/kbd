package util

import "kbd/module"

func Contains(arr []module.Realm, value module.Realm) bool {
	for _, v := range arr {
		if v.Name == value.Name {
			return true
		}
	}
	return false
}

func Index(arr []module.Realm, value module.Realm) int {
	for index, v := range arr {
		if v.Name == value.Name {
			return index
		}
	}
	return -1
}
