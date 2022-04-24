# TemperatureChecker

Simple tool written in Go to check CPU temperature in the specified time measured in seconds.
To run this program you first need to use Linux operating system, and you need to install acpi by executing command:
``` sudo apt-get install acpi ``` .
Way to execute command: ``` go run temperature.go <no_of_seconds> ```. 
On the output you'll get CPU temperature every second of the program execution
and also average, minimum, and maximum temperature during measurement.
