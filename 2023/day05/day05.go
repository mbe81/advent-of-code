package day05

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/advent-of-code/2023/util"
)

func Run() {
	input, err := util.ReadLines("./day05/input/input.txt")
	if err != nil {
		fmt.Printf("error reading file: %s", err.Error())
		return
	}
	start := time.Now()
	fmt.Println("Running Day 5, Part 1")
	fmt.Println("Result: ", Part1(input))
	fmt.Println("Running time: ", time.Since(start))

	start = time.Now()
	fmt.Println("Running Day 5, Part 2")
	fmt.Println("Result: ", Part2(input))
	fmt.Println("Running time: ", time.Since(start))
}

type Seed struct {
	Number      int
	Soil        int
	Fertilizer  int
	Water       int
	Light       int
	Temperature int
	Humidity    int
	Location    int
}

type Mapping struct {
	DestinationStart int
	SourceStart      int
	RangeLength      int
}

var seedToSoil,
	soilToFertilizer,
	fertilizerToWater,
	waterToLight,
	lightToTemperature,
	temperatureToHumidity,
	humidityToLocation []Mapping

func Part1(input []string) int {
	seeds := parseSeeds(input)
	parseAlmanac(input)

	minLocation := 0
	for _, seed := range seeds {
		soil := getMappingValue(seedToSoil, seed)
		fertilizer := getMappingValue(soilToFertilizer, soil)
		water := getMappingValue(fertilizerToWater, fertilizer)
		light := getMappingValue(waterToLight, water)
		temperature := getMappingValue(lightToTemperature, light)
		humidity := getMappingValue(temperatureToHumidity, temperature)
		location := getMappingValue(humidityToLocation, humidity)
		if minLocation == 0 || location <= minLocation {
			minLocation = location
		}
	}

	return minLocation
}

func Part2(input []string) int {
	seeds := parseSeedRanges(input)
	parseAlmanac(input)

	minLocation := 0
	for _, seed := range seeds {
		soil := getMappingValue(seedToSoil, seed)
		fertilizer := getMappingValue(soilToFertilizer, soil)
		water := getMappingValue(fertilizerToWater, fertilizer)
		light := getMappingValue(waterToLight, water)
		temperature := getMappingValue(lightToTemperature, light)
		humidity := getMappingValue(temperatureToHumidity, temperature)
		location := getMappingValue(humidityToLocation, humidity)
		//fmt.Println("seed", seed, "location", location)
		if minLocation == 0 || location <= minLocation {
			minLocation = location
		}
	}

	return minLocation
}

func parseAlmanac(input []string) {
	i := 2
	seedToSoil = parseMap(input, &i)
	soilToFertilizer = parseMap(input, &i)
	fertilizerToWater = parseMap(input, &i)
	waterToLight = parseMap(input, &i)
	lightToTemperature = parseMap(input, &i)
	temperatureToHumidity = parseMap(input, &i)
	humidityToLocation = parseMap(input, &i)
}

func parseSeeds(input []string) []int {
	var s []int
	for _, seedString := range strings.Fields(strings.Replace(input[0], "seeds: ", "", -1)) {
		seedNumber, _ := strconv.Atoi(seedString)
		s = append(s, seedNumber)
	}
	return s
}

func parseSeedRanges(input []string) []int {
	var s []int
	seedStrings := strings.Fields(strings.Replace(input[0], "seeds: ", "", -1))
	for i := 0; i < len(seedStrings); i = i + 2 {
		seedStart, _ := strconv.Atoi(seedStrings[i])
		seedRange, _ := strconv.Atoi(seedStrings[i+1])
		for j := seedStart; j <= seedStart+seedRange; j++ {
			s = append(s, j)
		}
	}
	return s
}

func parseMap(input []string, start *int) []Mapping {
	var mappings []Mapping

	i := *start
	for i < len(input) {
		if input[i] == "" {
			break
		}
		numbers := strings.Fields(input[i])
		if len(numbers) == 3 {
			var m Mapping
			m.DestinationStart, _ = strconv.Atoi(numbers[0])
			m.SourceStart, _ = strconv.Atoi(numbers[1])
			m.RangeLength, _ = strconv.Atoi(numbers[2])
			mappings = append(mappings, m)
		}
		i++
	}
	*start = i + 1
	return mappings
}

func getMappingValue(mappings []Mapping, sourceValue int) int {
	for _, mapping := range mappings {
		if sourceValue >= mapping.SourceStart && sourceValue <= mapping.SourceStart+mapping.RangeLength {
			return sourceValue + mapping.DestinationStart - mapping.SourceStart
		}
	}
	return sourceValue
}
