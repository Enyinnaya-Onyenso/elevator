package elevator

import (
	"sync"
)

type System struct { // Adding all components of the system
	elevator        *Elevator
	shutdownChannel chan struct{}
	logger          *Logger
}

func (s *System) Init() error {
	// Initializing all components in the System
	err := s.logger.init()
	if err != nil {
		return err
	}

	err = s.elevator.init()
	if err != nil {
		return err
	}

	return nil
}

func (s *System) Open() error {

	err := s.elevator.open()
	if err != nil {
		return err
	}

	return nil
}

func (s *System) Run() error {

	s.shutdownChannel = make(chan struct{})

	s.elevator.shutdownChannel = s.shutdownChannel // Assigning system channel to elevator channel to be used in abort
	done := make(chan struct{})
	defer close(done) // Ensuring done closes after successful run
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		select {
		case <-done:
			return
		case <-s.shutdownChannel:
			s.Abort() // Calling abort on unsuccessful run
		}
	}()

	err := s.elevator.run()
	if err != nil {
		return err
	}
	wg.Wait()
	return nil
}

func (s *System) Close() error {

	err := s.elevator.close()
	if err != nil {
		s.logger.Println(err)
	}

	return nil
}

func (s *System) Abort() error {

	select {
	case <-s.shutdownChannel:

	default:
		close(s.shutdownChannel)
	}

	return nil
}

func NewSystem() *System {
	// Function to create a new instance of system in main.go
	systemShutDownChannel := make(chan struct{})

	systemLogger := newLogger()

	systemElevator := newElevator(0, systemLogger, systemShutDownChannel)
	return &System{
		elevator:        systemElevator,
		logger:          systemLogger,
		shutdownChannel: systemShutDownChannel,
	}
}

func (s *System) Logger() *Logger {
	// Exposing logger to log in main.go
	return s.logger
}
