package main

import (
	"fmt"
	"os/exec"
	"time"
	"os"
	"strconv"
	"regexp"
	"runtime"
	"log"
)


func run_acpi(channel chan float32) {
	reg, _ := regexp.Compile("\\d+\\.\\d+")
	out, err := exec.Command("acpi", "-t").Output()
	if err != nil{
		fmt.Println(err)
	}
	output := string(out)
	t := time.Now()
	core_temp := reg.FindString(output)
	fmt.Printf("%s : %s C\n", t, core_temp)
	float_temp, err := strconv.ParseFloat(core_temp, 32)
	if err != nil{
		fmt.Println(err)
	}
	channel <- float32(float_temp)
}

func main(){
	seconds, err := strconv.Atoi(os.Args[1])
	if err != nil{
		fmt.Println(err)
	}
	if runtime.GOOS != "linux" {
		log.Fatal("You need to run this program on the linux operating system :p")
	}
	_, test_err := exec.Command("acpi").Output()
	if test_err != nil {
		fmt.Println(test_err)
		fmt.Printf("You probably don't have acpi package installed, you can install it by executing the following command: \n sudo apt-get install acpi\n")
		os.Exit(0)
	}
	
	ticker := time.NewTicker(1000 * time.Millisecond)
	done := make(chan bool)
	temp_chan := make(chan float32, seconds)
	var temp_sum, min_temp, max_temp float32
	temp_sum, min_temp, max_temp = 0, 1000.0, 0
	
	go func(){
		for{
			select{
				case <- ticker.C:
					run_acpi(temp_chan)
					current_temp := <- temp_chan
					temp_sum += current_temp
					if current_temp > max_temp{
						max_temp = current_temp
					}
					if current_temp < min_temp{
						min_temp = current_temp
					}
				case <- done:
					return
			}
		}
	}()
	
	time.Sleep(time.Duration(seconds) * time.Second)
	ticker.Stop()
	
	done <- true
	close(temp_chan)
	fmt.Printf("Avg temperature: %.2f\n",temp_sum/float32(seconds))
	fmt.Printf("Min temperature: %.2f\n", min_temp)
	fmt.Printf("Max temperature: %.2f\n", max_temp)
	
	
}
