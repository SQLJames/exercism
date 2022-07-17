package elon

import "fmt"

// TODO: define the 'Drive()' method
func (c *Car) Drive() *Car {
	if c.batteryDrain > c.battery {
		return c
	}
	c.distance = c.distance + c.speed
	c.battery = c.battery - c.batteryDrain
	return c
}

// TODO: define the 'DisplayDistance() string' method
func (c *Car) DisplayDistance() (distance string) {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

// TODO: define the 'DisplayBattery() string' method
func (c *Car) DisplayBattery() (Battery string) {
	return fmt.Sprintf(`Battery at %d%%`, c.battery)
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (c *Car) CanFinish(trackDistance int) bool {
	return (c.battery / c.batteryDrain) * c.speed >= trackDistance 
}
