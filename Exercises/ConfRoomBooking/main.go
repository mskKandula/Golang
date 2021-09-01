package main

import (
	"fmt"
	"strings"
)

type room struct {
	startTime int
	endTime   int
	id        string
	name      string
}

var (
	buildings []string
	floors    = make(map[string][]string)
	confRooms = make(map[string][]*room)
)

func main() {
	addBuilding("Building1")
	addFloor("Building1", "floor1")
	addConfRoom("Building1", "floor1", "conferenceRoom1")
	bookConfRoom("Building1", "floor1", "conferenceRoom1", 3, 5)
	bookConfRoom("Building1", "floor1", "conferenceRoom1", 1, 4)
	// cancelConfRoom("Building1", "floor1", "conferenceRoom1", 3, 5)
	bookConfRoom("Building1", "floor1", "conferenceRoom1", 2, 5)
}

func convertToLowerCase(name string) string {
	return strings.ToLower(name)
}

// To check valid time between 0-24hrs
func checkTime(start, end int) bool {
	if (start < 0 || end < 0) || (start > 23 || end > 23) {
		return false
	}
	return true
}

func checkBuildingExists(buildingName string) bool {
	for _, name := range buildings {
		if name == buildingName {
			return true
		}
	}
	return false
}

func checkFloorExists(buildingName, floorName string) bool {

	floorsList, ok := floors[buildingName]

	if !ok {
		return false
	}

	for _, name := range floorsList {
		if name == floorName {
			return true
		}
	}
	return false
}

func checkConfRoomExists(buildingName, floorName, confRoomName string) (bool, *room) {

	conferenceRooms, ok := confRooms[buildingName+floorName]

	if !ok {
		return false, &room{}
	}

	for _, room := range conferenceRooms {
		if room.name == confRoomName {
			return true, room
		}
	}

	return false, &room{}

}

func addBuilding(buildingName string) {

	nameOfBuilding := convertToLowerCase(buildingName)

	buildingExists := checkBuildingExists(nameOfBuilding)

	if buildingExists {
		fmt.Printf("%v Building already exists\n", buildingName)
		return
	}

	buildings = append(buildings, nameOfBuilding)

	fmt.Printf("%v Building added successfully\n", buildingName)
}

func addFloor(buildingName, floorName string) {

	nameOfBuilding := convertToLowerCase(buildingName)

	buildingExists := checkBuildingExists(nameOfBuilding)

	if !buildingExists {
		fmt.Printf("%v Building doesn't exist to add a floor\n", buildingName)
		return
	}

	nameOfFloor := convertToLowerCase(floorName)

	floorExists := checkFloorExists(nameOfBuilding, nameOfFloor)

	if floorExists {
		fmt.Printf("%v Floor already exists,Please try with another name\n", floorName)
		return
	}

	floors[nameOfBuilding] = append(floors[nameOfBuilding], nameOfFloor)

	fmt.Printf("%v Floor added successfully\n", floorName)

}

func addConfRoom(buildingName, floorName, confRoomName string) {
	bNameLower := convertToLowerCase(buildingName)
	fNameLower := convertToLowerCase(floorName)
	cNameLower := convertToLowerCase(confRoomName)

	buildingExists := checkBuildingExists(bNameLower)

	if !buildingExists {
		fmt.Printf("%v Building doesn't exist to add a conference room\n", buildingName)
		return
	}

	floorExists := checkFloorExists(bNameLower, fNameLower)

	if !floorExists {
		fmt.Printf("%v Floor doesn't exist to add a conference room", floorName)
		return
	}

	confRoomExists, _ := checkConfRoomExists(bNameLower, fNameLower, cNameLower)

	if confRoomExists {

		fmt.Printf("%v Conference room already exists\n", confRoomName)

		return

	} else {

		confRoom := &room{
			id:        bNameLower + fNameLower + cNameLower,
			name:      cNameLower,
			startTime: 0,
			endTime:   0,
		}

		confRooms[bNameLower+fNameLower] = append(confRooms[bNameLower+fNameLower], confRoom)

		fmt.Printf("%v Conference room added successfully\n", confRoomName)
	}
}

func bookConfRoom(buildingName, floorName, confRoomName string, startTime, endTime int) {

	bNameLower := convertToLowerCase(buildingName)
	fNameLower := convertToLowerCase(floorName)
	cNameLower := convertToLowerCase(confRoomName)

	buildingExists := checkBuildingExists(bNameLower)

	if !buildingExists {
		fmt.Printf("%v Building doesn't exist to book a conference room\n", buildingName)
		return
	}

	floorExists := checkFloorExists(bNameLower, fNameLower)

	if !floorExists {
		fmt.Printf("%v Floor doesn't exist to book a conference room\n", floorName)
		return
	}

	confRoomExists, confRoom := checkConfRoomExists(bNameLower, fNameLower, cNameLower)

	if !confRoomExists {

		fmt.Printf("%v conference room doesn't exists\n", confRoomName)

		return

	}
	if !checkTime(startTime, endTime) {
		fmt.Println("Please add proper start & end times")
		return
	}
	if (confRoom.startTime == 0 || (confRoom.startTime < startTime && startTime >= confRoom.endTime)) ||
		(confRoom.endTime == 0 || (confRoom.startTime <= endTime && endTime > confRoom.endTime)) {

		confRoom.startTime = startTime
		confRoom.endTime = endTime
		fmt.Printf("%v Conference room booked successfully\n", confRoomName)

	} else {
		fmt.Printf("Conference room not available for this time: %v - %v\n", startTime, endTime)
		return
	}
}

func cancelConfRoom(buildingName, floorName, confRoomName string, startTime, endTime int) {
	bNameLower := convertToLowerCase(buildingName)
	fNameLower := convertToLowerCase(floorName)
	cNameLower := convertToLowerCase(confRoomName)

	buildingExists := checkBuildingExists(bNameLower)

	if !buildingExists {
		fmt.Printf("%v Building doesn't exist to cancel a booked conference room\n", buildingName)
		return
	}

	floorExists := checkFloorExists(bNameLower, fNameLower)

	if !floorExists {
		fmt.Printf("%v Floor doesn't exist to cancel a booked conference room\n", floorName)
		return
	}

	confRoomExists, confRoom := checkConfRoomExists(bNameLower, fNameLower, cNameLower)

	if !confRoomExists {

		fmt.Printf("%v Conference room doesn't exists\n", confRoomName)

		return
	}

	if !checkTime(startTime, endTime) {
		fmt.Println("Please add proper start & end times")
		return
	}

	if confRoom.startTime == startTime && confRoom.endTime == endTime {
		confRoom.startTime = 0
		confRoom.endTime = 0
		fmt.Printf("%v conference room from time %v - %v cancelled successfully\n", confRoomName, startTime, endTime)
		return
	} else {
		fmt.Println("Please check start & end times for this conference room:", confRoomName)
		return
	}
}
