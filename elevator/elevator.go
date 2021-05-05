package elevator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Elevator struct {
	currentFloor     int64
	currentSignal    int64
	voltageSignal    int64
	destinationFloor int64
	shutdownChannel  chan struct{} // to be used as shared shutdown channel
	logger           *Logger       // shared logger type from logger.go
	maxFloors        int64
	ESTOP            int64
}

func (e *Elevator) init() error {
	// Initializing  struct variables
	e.currentFloor = 0
	e.destinationFloor = 0
	e.currentSignal = 0
	e.voltageSignal = 0
	e.ESTOP = -999
	e.maxFloors = 100
	return nil
}

func (e *Elevator) open() error {
	// No components to open
	return nil
}

func monitorFunction(runChannel chan struct{}, shutdownChannel chan struct{}, abort func() error, logger *Logger) {
	select {
	case <-runChannel:
		return
	case <-shutdownChannel:
		err := abort()
		if err != nil {
			logger.Println(err)
		}
		return
	}
}

func (e *Elevator) run() error {
	runDoneChannel := make(chan struct{})
	defer close(runDoneChannel)

	go monitorFunction(runDoneChannel, e.shutdownChannel, e.abort, e.logger) // Calling monitor function to check for abort

	fmt.Println("Enter Destination Floor: ")
	button := bufio.NewScanner(os.Stdin)
	button.Scan()
	var err error
	e.destinationFloor, err = strconv.ParseInt(button.Text(), 10, 64)
	if err != nil {
		e.logger.Println("Error: ", err) // Logging errors
		fmt.Println(&e.logger.buf)       // Printing log to console
		e.logger.buf.Reset()
	} else { // Nested if statements to handle logic
		if e.destinationFloor == e.currentFloor { // When user destination is the same as the floor they are on, this branch runs
			e.logger.Println("You are currently on this floor. Floor: ", e.destinationFloor)
			fmt.Println(&e.logger.buf)
			e.logger.buf.Reset()

		} else if e.destinationFloor < 0 && e.destinationFloor != e.ESTOP { // When user enters a negative floor as destination and that floor is not the designated ESTOP value(-999)
			e.logger.Println("Invalid Floor, Please Enter a floor between 0 and 100")
			fmt.Println(&e.logger.buf)
			e.logger.buf.Reset()

		} else {
			if e.destinationFloor > e.maxFloors { // When user destination is higher than the max number of floors 100
				e.logger.Println("This floor does not exist")
				fmt.Println(&e.logger.buf)
				e.logger.buf.Reset()

			} else if e.destinationFloor == e.ESTOP { // When user enters designated ESTOP vaalue (-999)
				e.logger.Println("Estop Initiated")
				fmt.Println(&e.logger.buf)
				e.logger.buf.Reset()

				e.close()
			} else {
				if e.destinationFloor > e.currentFloor { // When a valid destination is entered and the user destination floor is above the current floor
					e.logger.Println("Going up")
					fmt.Println(&e.logger.buf)
					e.logger.buf.Reset()

					err = e.up()
					if err != nil {
						e.logger.Println("Error: ", err)
						fmt.Println(&e.logger.buf)
						e.logger.buf.Reset()
						return err
					}

				} else if e.destinationFloor < e.currentFloor {
					// When a valid destination is entered and the user destination is below current floor (current floor can not be 0 to be valid)
					err = e.down()
					if err != nil {
						e.logger.Println("Error: ", err)
						fmt.Println(&e.logger.buf)
						e.logger.buf.Reset()
						return err
					}

				}
			}
		}
	}
	err = e.run() // Recursive function to  ensure it repeats
	if err != nil {
		e.logger.Println("Error: ", err)
		fmt.Println(&e.logger.buf)
		e.logger.buf.Reset()
		return err
	}

	time.Sleep(2 * time.Second)

	e.logger.Logger.Println()
	return nil
}

func (e *Elevator) close() error {
	// Close the function with a status 0 for a successful run
	os.Exit(0)
	return nil
}

func (e *Elevator) abort() error {
	// Close the function with a status of 1 for all unsuccessful runs regardless of errors
	os.Exit(1)
	return nil
}

func newElevator(currentFloor int64, inpLogger *Logger, shutdownChannel chan struct{}) *Elevator {
	// Creates a new instance of elevator to be used in system.go
	return &Elevator{
		currentFloor:    currentFloor,
		logger:          inpLogger,
		shutdownChannel: shutdownChannel,
	}
}

func (e *Elevator) up() error {
	// This function handles going to higher floors
	for e.destinationFloor > e.currentFloor {
		time.Sleep(500 * time.Millisecond)
		e.currentFloor++
		e.logger.Logger.Println("You are on floor: ", e.currentFloor)
		fmt.Println(&e.logger.buf)
		e.logger.buf.Reset()
	}
	return nil
}

func (e *Elevator) down() error {
	// This function handles going to lower floors
	for e.destinationFloor < e.currentFloor {
		time.Sleep(500 * time.Millisecond)
		e.currentFloor--
		e.logger.Logger.Println("You are on floor: ", e.currentFloor)
		fmt.Println(&e.logger.buf)
		e.logger.buf.Reset()
	}
	return nil
}
