package main

import "fmt"

func move(disks int, source string, intermed string, dest string) {
	if disks == 1 {
		fmt.Printf("Move a disk from %s to %s\n", source, dest)
	} else if disks > 1 {
		move(disks-1, source, dest, intermed)
		move(1, source, intermed, dest)
		move(disks-1, intermed, source, dest)
	}
}

func main() {
	disks := 4
	move(disks, "1", "2", "3")
}
