package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	var (
		cfloor  int64 //defining current floor variable
		csignal int64 //defining current signal
		vsignal int64
		buf     bytes.Buffer
		logger  = log.New(&buf, "logger: ", log.Lmicroseconds)
		m       sync.RWMutex
	)
	fmt.Println("You are currently on Floor", cfloor)
	fmt.Println("Enter Destination Floor Number: ")
	button := bufio.NewScanner(os.Stdin)
	button.Scan()
	dfloor, _ := strconv.ParseInt(button.Text(), 10, 64)
	for dfloor < 0 && dfloor != -999 {
		fmt.Println(errors.New("Invalid Floor, Please Enter a floor between 0 and 100"))
		button = bufio.NewScanner(os.Stdin)
		button.Scan()
		dfloor, _ = strconv.ParseInt(button.Text(), 10, 64)
	}
	fmt.Println("Now going to Floor: ", dfloor)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() int64 {
		defer wg.Done()
		for dfloor != cfloor {
			for dfloor > cfloor {
				time.Sleep(500 * time.Millisecond)
				m.Lock()
				var ifloor int64
				ifloor = cfloor + 1
				logger.Print("You're now on Floor: ", (ifloor))
				cfloor = ifloor
				m.Unlock()
			}
			for dfloor < cfloor && dfloor != -999 {
				time.Sleep(500 * time.Millisecond)
				m.Lock()
				logger.Print("You're now on Floor: ", (cfloor - 1))
				cfloor--
				m.Unlock()
			}
			if dfloor == -999 {
				logger.Print("Emergency Stop Initiated. You are currently on Floor: ", cfloor)
				break
			}
			//fmt.Println(&buf)
		}
		fmt.Println(&buf)
		return (cfloor)
	}()
	go func() {
		defer wg.Done()
		for dfloor != cfloor {
			for dfloor > cfloor {
				time.Sleep(100 * time.Millisecond)
				m.Lock()
				csignal = rand.Int63n(100)
				vsignal = csignal * (dfloor - cfloor)
				logger.Print("Current: ", csignal, " Voltage: ", vsignal)
				m.Unlock()
			}
			for dfloor < cfloor && dfloor != -999 {
				time.Sleep(100 * time.Millisecond)
				m.Lock()
				csignal = rand.Int63n(100)
				vsignal = csignal * (dfloor - cfloor)
				logger.Print("Current: ", csignal, " Voltage: ", vsignal)
				m.Unlock()
			}
			if dfloor == -999 {
				logger.Print("Emergency Stop, No Signal")
				break
			}
			//fmt.Println(&buf)
		}
	}()
	wg.Wait()
	fmt.Println(cfloor)
	//log.Print(csignal)
}
