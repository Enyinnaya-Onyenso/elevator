# elevator
A  simulation of an elevator project.
The question given to me is written below:
Please write a program that similates the use of an elevator. The Program must have the following
- CLI input to select the floor (assume 100 floors)
    = should accept only ints and on special occasion an ESTOP command (assume -999)
    = Must display error command if command shows < 0
- Logger that logs every command inputted from the CLI
- Assume that the elevator moves at 2 floors a second
     = Logger must log every single time it reaches a floor. 
    = Example (if you're on floor 3, and press 10, there should be a logged timestamp every 500 ms to show that you hit floor 4, 5, 6 and so on...)
- If an estop is pressed, logger logs which floor you stopped on and output message to user on CLI
- Have a data stream that has random voltage and current readings every time elevator moves. You must log this as well
    = Example. If you're on floor 3, and moving to floor 10, you should be outputting a constant positive current and assume voltage output is the difference between each floor.
    = If you're moving from floor 10 to floor 9, you should be outputting a negative current and assume voltage output is the difference between each floor.
    = Readings from current and voltage need to be logged at a rate of 100 ms.
