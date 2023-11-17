normal_work_days <- 5
normal_work_hours <- 8
daily_salary <- 500

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



cat('Main Menu\n', '[1] Compute\n', '[2] Idk\n', '[0] Exit\n')

# Read user input as a string
user_input <- readline('Enter your choice: ')

# Convert the user input to a numeric value (if necessary)
user_choice <- as.numeric(user_input)

cat('hi', user_choice)

