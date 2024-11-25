package main

import "strconv"

func checkMembersNb(artist artistData, membersNb string) bool {
	nb, _ := strconv.Atoi(membersNb)
	return len(artist.Members) == nb
}

func filter(arr []artistData, criteria string, check func (artistData, string) bool) []artistData {
	newArr := []artistData{}
	for _, artist := range arr {
		if check(artist, criteria) {
			newArr = append(newArr, artist)
		}
	}
	return newArr
}