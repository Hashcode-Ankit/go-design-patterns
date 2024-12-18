package factory

import "fmt"

type Cloud struct {
	name      string
	cloudType string
}

type CloudInterface interface {
	SpinUpMachine() string
	DowngradeMachine() string
}

func (c *Cloud) SpinUpMachine() string {
	if c.cloudType == "AZR" {
		return func() string { return "" }()
	}
	return fmt.Sprintf("hey bro! spinned up machine on cloud")
}

func (c *Cloud) DowngradeMachine() string {
	if c.cloudType == "AZR" {
		return func() string { return "" }()
	}
	return fmt.Sprintf("hey bro! downgraded machine on cloud")
}
