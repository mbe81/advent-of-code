package day06

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/advent-of-code/2023/util"
)

func Run() {
	input, err := util.ReadLines("./day06/input/input.txt")
	if err != nil {
		fmt.Printf("error reading file: %s", err.Error())
		return
	}

	var start time.Time

	start = time.Now()
	util.PrettyPrint(5, 1, Part1(input), start)

	start = time.Now()
	util.PrettyPrint(5, 2, Part2(input), start)
}

func Part1(input []string) int {

	waysToBeatRecordProduct := 1
	raceRecords := parseRaceRecords(input, 1)

	for _, rr := range raceRecords {
		waysToBeatRecord := 0

		for holdTime := 1; holdTime < rr.Time; holdTime++ {
			speed := holdTime
			distance := speed * (rr.Time - holdTime)
			if distance > rr.DistanceRecord {
				waysToBeatRecord++
			}
		}

		waysToBeatRecordProduct = waysToBeatRecordProduct * waysToBeatRecord
	}

	return waysToBeatRecordProduct
}

type RaceRecord struct {
	Time           int
	DistanceRecord int
}

func parseRaceRecords(input []string, part int) []RaceRecord {
	var raceRecords []RaceRecord

	times := strings.Fields(strings.Replace(input[0], "Time:", "", -1))
	distances := strings.Fields(strings.Replace(input[1], "Distance:", "", -1))

	for i := 0; i < len(times); i++ {
		var rr RaceRecord

		rr.Time, _ = strconv.Atoi(times[i])
		rr.DistanceRecord, _ = strconv.Atoi(distances[i])

		raceRecords = append(raceRecords, rr)
	}

	return raceRecords
}

func Part2(input []string) int {

	waysToBeatRecord := 0

	input[0] = strings.Replace(input[0], "Time:", "", -1)
	input[0] = strings.Replace(input[0], " ", "", -1)
	racingTime, _ := strconv.Atoi(input[0])

	input[1] = strings.Replace(input[1], "Distance:", "", -1)
	input[1] = strings.Replace(input[1], " ", "", -1)
	distanceRecord := new(big.Int)
	distanceRecord, _ = distanceRecord.SetString(input[1], 10)

	for holdTime := 1; holdTime < racingTime; holdTime++ {
		speed := holdTime
		distance := big.NewInt(int64(speed * (racingTime - holdTime)))
		if distance.Cmp(distanceRecord) == 1 {
			waysToBeatRecord++
		} else {
			if waysToBeatRecord > 0 {
				break
			}
		}
	}

	return waysToBeatRecord
}
