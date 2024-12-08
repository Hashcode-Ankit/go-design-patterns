package builder

import "fmt"

func (c *CarBuilder) Drive() string {
	return "Hey I have started driving"
}

func (c *CarBuilder) Stop() string {
	return fmt.Sprintf("Applying breaks on my %s wheels", c.car.wheels)
}
