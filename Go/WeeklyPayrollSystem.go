package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Define days of the week
var daysOfWeek = map[int]string{
	0: "Monday",
	1: "Tuesday",
	2: "Wednesday",
	3: "Thursday",
	4: "Friday",
	5: "Saturday",
	6: "Sunday",
}

// default setting configuration
var workDays = 5
var maxWorkHours = 8
var dailySalary = 500.00
var inTime = "0900"
var outTime = "0900"
var dayTypes = [7]int{0, 0, 0, 0, 0, 0, 0} // 0 = normal, 1 = special non-working, 2 = regular holiday

func main() {
	showSetting()
	menu()
}

func menu() {
	var choice int

	for {
		fmt.Print("\nMain Menu\n[1] Compute Weekly Salary\n[2] Configure Settings\n[3] Exit\n\nEnter your choice: ")

		// Read the whole line
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')

		// Trim spaces and convert to an integer
		input = strings.TrimSpace(input)
		choice, err = strconv.Atoi(input)

		if err != nil || choice < 1 || choice > 3 {
			fmt.Print("\nError: Please enter a valid integer.\n\n")
			continue
		} else if choice == 1 {
			computeWeeklySalary()
		} else if choice == 2 {
			configureSettings()
		} else {
			fmt.Print("\nExiting Program...\n\n")
			// valid choice, break out of the loop
			break
		}
	}
}

func getTimeDifference(inTime, out string) int {
	hours, _ := strconv.ParseFloat(inTime[:2], 64)
	minutes, _ := strconv.ParseFloat(inTime[2:], 64)
	startTime := hours + minutes/60

	outHours, _ := strconv.ParseFloat(out[:2], 64)
	outMinutes, _ := strconv.ParseFloat(out[2:], 64)
	endTime := outHours + outMinutes/60

	diff := 0.0

	if startTime > endTime {
		diff += 24 - startTime + endTime
	} else if startTime < endTime {
		diff = endTime - startTime
	}

	return int(diff)
}

func getNightShift(out string) int {
	nightShiftHours := 0

	// 2200 to 0600
	if out > inTime {
		if out > "2200" {
			nightShiftHours = getTimeDifference("2200", out)
		}
	} else if out < inTime {
		nightShiftHours += 2
		if out >= "0600" {
			nightShiftHours += 6
		} else {
			nightShiftHours += getTimeDifference("0000", out)
		}
	}

	return nightShiftHours
}

func computeWeeklySalary() {
	var outTimes [7]string
	weeklySalary := 0.0

	for i := 0; i < 7; i++ {
		fmt.Println("Out time for", daysOfWeek[i], ":")
		outTimes[i] = scanTime()
	}

	for i := 0; i < 7; i++ {
		salary := 0.0
		fmt.Println("\n-------------", daysOfWeek[i], "-------------")
		fmt.Printf("Daily Rate: %.2f\n", dailySalary)
		fmt.Println("IN Time:", inTime)
		fmt.Println("OUT Time:", outTimes[i])
		fmt.Print("Day Type: ")

		if i >= workDays {
			fmt.Print("Rest Day")
		} else {
			fmt.Print("Work Day")
		}

		switch dayTypes[i] {
		case 0:
			fmt.Println(", Normal Day")
		case 1:
			fmt.Println(", Special Non-Working Holiday")
		case 2:
			fmt.Println((", Regular Holiday"))
		}

		// compute for OT hours, night shift hours, night OT hours if there are
		timeDifference := getTimeDifference(inTime, outTimes[i])

		if timeDifference != 0 {
			if timeDifference < maxWorkHours+1 {
				timeDifference = 9
			}

			OTHours := timeDifference - maxWorkHours - 1
			nightShiftHours := getNightShift(outTimes[i])
			nightOTHours := 0

			if OTHours > 0 && nightShiftHours > 0 {
				if OTHours > nightShiftHours {
					nightOTHours = nightShiftHours
					OTHours -= nightShiftHours
					nightShiftHours = 0
				} else {
					nightOTHours = OTHours
					nightShiftHours -= OTHours
					OTHours = 0
				}
			}

			fmt.Printf("Hours Overtime (Night Shift Overtime): %d (%d)\n", OTHours, nightOTHours)

			// compute for the salary
			workHours := float64(maxWorkHours)

			salary += dailySalary*getDailyRate(i) +
				(float64(OTHours) * dailySalary / workHours * getOTrate(i)) +
				(float64(nightOTHours) * dailySalary / workHours * getOTNightRate(i)) +
				(float64(nightShiftHours) * dailySalary / workHours * 1.1)

			weeklySalary += salary
		}

		fmt.Printf("Salary for the day: %.2f\n", salary)
		fmt.Println("---------------------------------")
	}

	fmt.Printf("\nWeekly Salary: %.2f\n", weeklySalary)

}

func getOTNightRate(i int) float64 {
	multiplier := 1.375

	if i >= workDays {
		multiplier = 1.859

		switch dayTypes[i] {
		case 1:
			multiplier = 2.145
		case 2:
			multiplier = 3.718
		}
	} else {
		switch dayTypes[i] {
		case 1:
			multiplier = 1.859
		case 2:
			multiplier = 2.86
		}
	}

	return multiplier
}

func getDailyRate(i int) float64 {
	multiplier := 1.0

	if i >= workDays {
		multiplier = 1.30

		switch dayTypes[i] {
		case 1:
			multiplier = 1.50
		case 2:
			multiplier = 2.60
		}
	} else {
		switch dayTypes[i] {
		case 1:
			multiplier = 1.30
		case 2:
			multiplier = 2.00
		}
	}

	return multiplier
}

func getOTrate(i int) float64 {
	multiplier := 1.25

	if i >= workDays {
		multiplier = 1.60

		switch dayTypes[i] {
		case 1:
			multiplier = 1.95
		case 2:
			multiplier = 3.38
		}
	} else {
		switch dayTypes[i] {
		case 1:
			multiplier = 1.69
		case 2:
			multiplier = 2.60
		}
	}

	return multiplier
}

func showSetting() {
	fmt.Println("\nCurrent Settings:")
	fmt.Println("--------------------")
	fmt.Println("(1) Daily Salary:", dailySalary)
	fmt.Println("(2) Max Work Hours:", maxWorkHours)
	fmt.Println("(3) Work Days:", workDays)
	fmt.Println("(4) In Time:", inTime)
	fmt.Println("(5) Out Time:", outTime)
	fmt.Println("(6) Day Types for the week:\n  (0 = normal, 1 = special non-working, 2 = regular holiday)\n   =>", dayTypes)
}

func checkInt() int {
	var input int
	for {
		// fmt.Print("Enter an integer: ")
		_, err := fmt.Scan(&input)
		if err == nil && input >= 0 {
			break
		}
		fmt.Println("Invalid input. Try again.")
		fmt.Scanln() // Clear the input buffer
	}
	return input
}

func scanTime() string {
	var time string
	pattern := regexp.MustCompile(`^\d{4}$`)

	for {
		fmt.Print("Enter a time (HHMM): ")
		fmt.Scanln(&time)

		if pattern.MatchString(time) && time >= "0000" && time <= "2359" && time[2:] >= "00" && time[2:] <= "59" {
			break
		} else {
			fmt.Println("Invalid time. Try again.")
		}
	}

	return time
}

func configureSettings() {
	// normal, regular holiday, special non-working

	showSetting()

	// get new daily salary
	fmt.Print("\nNew daily salary: ")
	for {
		_, err := fmt.Scan(&dailySalary)
		if err == nil && dailySalary >= 0 {
			break
		}
		fmt.Println("Invalid input. Try again.")
		fmt.Scanln() // Clear the input buffer
	}

	// get new maximum work hours
	fmt.Print("\nNew maximum regular work hours: ")
	for {
		maxWorkHours = checkInt()
		if maxWorkHours >= 1 && maxWorkHours < 24 {
			break
		}
		fmt.Println("Invalid number of hours (at least 1 hour and less than 24). Try again.")
	}

	// get new number of work days
	fmt.Print("\nNew number of work days per week: ")
	for {
		workDays = checkInt()
		if workDays <= 7 {
			break
		}
		fmt.Println("Invalid number of days. Try again.")
	}

	// get what type of days are in a week
	for ctr := 0; ctr < 7; ctr++ {
		fmt.Printf("Type of day for %s:\n", daysOfWeek[ctr])
		fmt.Println("(0) Normal Day")
		fmt.Println("(1) Special Non-Working Day")
		fmt.Println("(2) Regular Holiday")

		var dayType int
		for {
			dayType = checkInt()
			if dayType >= 0 && dayType <= 2 {
				break
			}
			fmt.Println("Invalid input. Try again.")
		}

		dayTypes[ctr] = dayType
	}

	fmt.Println("New IN time: ")
	inTime = scanTime()
	fmt.Println("New OUT time: ")
	outTime = scanTime()

	showSetting()
}
