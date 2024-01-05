package main

import "os"

func main() {
	if err := os.WriteFile("/sys/class/power_supply/macsmc-battery/charge_control_start_threshold", []byte("75"), 0644); err != nil {
		os.Exit(1)
	}

	if err := os.WriteFile("/sys/class/power_supply/macsmc-battery/charge_control_end_threshold", []byte("80"), 0644); err != nil {
		os.Exit(1)
	}
}
