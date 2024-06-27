# Cron Parser

This is a command line application that parses a cron string and expands each field to show the times at which it will run. The cron string follows the standard format with five time fields (minute, hour, day of month, month, and day of week) plus a command.

## Prerequisites

- Go (version 1.16 or later)

## Installation

Clone the folder to your local machine.

"https://github.com/mskKandula/Golang/tree/main/Assignment"

## Run with the below command

go run main.go "*/15 0 1,15 * 1-5 /usr/bin/find"

## Expected Output

minute         0 15 30 45 
hour           0 
day of month   1 15 
month          1 2 3 4 5 6 7 8 9 10 11 12 
day of week    1 2 3 4 5 
command        /usr/bin/find


Unit tests are provided to test the Parse method. To run the tests, use the following command:

go test -v
