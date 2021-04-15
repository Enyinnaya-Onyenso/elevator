package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	//"io"
	"bufio"
	"bytes"
	"os"
	"strconv"
	"sync"
)

func main() {
	var (
		cfloor  int //defining current floor variable
		csignal int //defining current signal
		ESTOP   int64
		buf     bytes.Buffer
		logger  = log.New(&buf, "logger: ", log.Lmicroseconds)
	)
	ESTOP = -999
	fmt.Println("You are currently on Floor", cfloor)
	fmt.Println("Enter Destination Floor Number: ")
	button := bufio.NewScanner(os.Stdin)
	button.Scan()
	dfloor, _ := strconv.ParseInt(button.Text(), 10, 64)
	fmt.Print(dfloor)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() (int64, error) {
		defer wg.Done()
		if dfloor > 0 && dfloor != ESTOP {
			return dfloor, errors.New("Invalid Floor, Please Enter a floor between 0 and 100")
		}
		return dfloor, nil
	}()
	wg.Wait()
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(500 * time.Millisecond)
		logger.Print("You're now on Floor: ", (cfloor + 1))
		fmt.Print(&buf)
	}()
	wg.Wait()
	log.Print(csignal)
}
