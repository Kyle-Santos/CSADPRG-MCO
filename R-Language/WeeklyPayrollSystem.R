normalWorkDays <- 5
maxHours <- 8
dailySalary <- 500
inTime <- "0900"
outTime <- "1800"

# Define a named vector
days_of_the_week <- c(MONDAY = 1, TUESDAY = 2, WEDNESDAY = 3, THURSDAY = 4, FRIDAY = 5, SATURDAY = 6, SUNDAY = 7)


main <- function() {
  c <- menu()
  
  switch(choice,
        "1" = computeSalary(),
        "2" = configureSettings(),
        "3" = cat("Exiting...\n"),
        cat("Invalid choice. Please enter 1, 2, or 3.\n")
  )
}

# Use the names as the enum values
# day <- days_of_the_week[1]

# Print the value
# print(day)

# 9 - 6 normal work hours
# 10pm - 6am + 10% hourly rate
# hourly rate = daily salary / normal_work_hours

# rest = 130% rate
# special non-working 130%
# non-working & rest 150%
# regular holday 200%
# regular holiday & rest day 260%

# exceed hour (non night, night)
# normal  125%   137.5%
# rest    169%   185.9%
# non-working    169%    185.9%
# non-working & rest  195%    214.5%
# holiday 260%    286%
# holiday & rest  338%    371.8%

menu <- function() {
  choice <- ""
  
  while (!(choice %in% c("1", "2", "3"))) {
    cat('Main Menu\n', '[1] Compute Weekly Salary\n', '[2] Configure Settings\n', '[3] Exit\n\n', 'Enter your choice: ')
    
    # Read user input as a string
    choice <- readLines('stdin', n=1)
    
    cat("\n")
  }
  
  return(choice)
  
  # Convert the user input to a numeric value (if necessary)
  # user_choice <- as.numeric(user_input)
}


main()

