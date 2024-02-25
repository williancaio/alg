package main

import (
	"fmt"
	"sort"
)

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][2] < meetings[j][2]
	})
	personKnowAndHour := [][]int{}

	personKnowAndHour = append(personKnowAndHour, []int{firstPerson, 0})
	personKnowAndHour = append(personKnowAndHour, []int{0, 0})

	people := []int{}
	shouldVerifySameHour := true
	lastHourVerify := 0
	people = append(people, 0)
	people = append(people, firstPerson)
	breakPersonKnowLoop := false
	auxY := 0
	for i := 0; i < len(meetings); i++ {
		breakPersonKnowLoop = false
		if !shouldVerifySameHour && lastHourVerify != meetings[i][2] {
			shouldVerifySameHour = true
			auxY = 0
		}
		for y := auxY; y < len(personKnowAndHour); y++ {
			if breakPersonKnowLoop {
				auxY = y
				break
			}

			if meetings[i][0] == personKnowAndHour[y][0] && meetings[i][2] >= personKnowAndHour[y][1] {
				if CanAdd(meetings[i][1], people) {
					personKnowAndHour = append(personKnowAndHour, []int{meetings[i][1], meetings[i][2]})
					people = append(people, meetings[i][1])
				}

				if shouldVerifySameHour {
					lastHourVerify = meetings[i][2]
					shouldVerifySameHour = false
					breakPersonKnowLoop = true

					for i >= 0 {
						i--
						if i == -1 {
							break
						}
						if meetings[i][2] != lastHourVerify {
							i++
							break
						}

					}
				}

			} else if meetings[i][1] == personKnowAndHour[y][0] && meetings[i][2] >= personKnowAndHour[y][1] {
				if CanAdd(meetings[i][0], people) {
					personKnowAndHour = append(personKnowAndHour, []int{meetings[i][0], meetings[i][2]})
					people = append(people, meetings[i][0])
					continue
				}

				if shouldVerifySameHour {
					lastHourVerify = meetings[i][2]
					shouldVerifySameHour = false
					breakPersonKnowLoop = true
					for i >= 0 {
						i--

						if i == -1 {
							break
						}

						if meetings[i][2] != lastHourVerify {
							i++
							break
						}
					}
				}

			}
		}
	}

	return people
}

func CanAdd(person int, personKnow []int) bool {
	for _, v := range personKnow {
		if person == v {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("teste")

	//findAllPeople(12, [][]int{{10, 8, 6}, {9, 5, 11}, {0, 5, 18}, {4, 5, 13}, {11, 6, 17}, {0, 11, 10}, {10, 11, 7}, {5, 8, 3}, {7, 6, 16}, {3, 6, 10}, {3, 11, 1}, {8, 3, 2}, {5, 0, 7}, {3, 8, 20}, {11, 0, 20}, {8, 3, 4}, {1, 9, 4}, {10, 7, 11}, {8, 10, 18}}, 9)
	//result := findAllPeople(12, [][]int{{3, 1, 3}, {1, 2, 2}, {0, 3, 3}}, 3)

	result := findAllPeople(1, [][]int{{1, 4, 3}, {0, 4, 3}}, 3)
	for _, v := range result {
		fmt.Println(v)
	}
}
