package speed

// TODO: define the 'Car' type struct
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		100,
		batteryDrain,
		speed,
		0,
	}
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	var newBattery = car.battery
	var newDistance = car.distance
	if car.battery-car.batteryDrain >= 0 {
		newBattery = car.battery - car.batteryDrain
		newDistance = car.distance + car.speed
	}
	return Car{
		newBattery,
		car.batteryDrain,
		car.speed,
		newDistance,
	}
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	batteryneeded := (track.distance / car.speed) * car.batteryDrain
	return batteryneeded <= car.battery
}
