package main

import (
	elevator "elevator"
)

func main() {
	elevatorSystem := elevator.NewSystem()

	err := elevatorSystem.Init()
	if err != nil {
		elevatorSystem.Logger().Println(err)
		return
	}

	err = elevatorSystem.Open()
	if err != nil {
		elevatorSystem.Logger().Println(err)
		return
	}

	defer elevatorSystem.Close()

	err = elevatorSystem.Run()
	if err != nil {
		elevatorSystem.Logger().Println(err)
		return
	}
}
