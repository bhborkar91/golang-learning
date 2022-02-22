package main

import "fmt"

type disk struct {
	size int
}

type tower struct {
	name  string
	disks []disk
}

func newDisk(size int) *disk {
	d := disk{size: size}
	return &d
}

func newTower(name string, initialHeight int) *tower {
	t := tower{name: name}
	t.disks = make([]disk, initialHeight)
	for i := 0; i < initialHeight; i++ {
		t.disks[i] = *newDisk(initialHeight - i)
	}
	return &t
}

func (tower tower) print() {
	fmt.Printf("Tower[%s] : [", tower.name)
	for _, disk := range tower.disks {
		fmt.Printf(" D%d ", disk.size)
	}
	fmt.Println("]")
}

func (tower *tower) takeFrom(source *tower) {
	sourceHeight := len(source.disks)
	diskToMove := source.disks[sourceHeight-1]
	source.disks = source.disks[:sourceHeight-1]

	destHeight := len(tower.disks)
	if destHeight > 0 {
		baseDisk := tower.disks[destHeight-1]
		if baseDisk.size <= diskToMove.size {
			panic("Cannot move this disk")
		}
	}
	tower.disks = append(tower.disks, diskToMove)
}

func move(disks int, source *tower, intermed *tower, dest *tower, allTowers []*tower) {
	if disks == 1 {
		fmt.Printf("Move a disk from %s to %s\n", source.name, dest.name)
		dest.takeFrom(source)
		for _, tower := range allTowers {
			tower.print()
		}
	} else if disks > 1 {
		move(disks-1, source, dest, intermed, allTowers)
		move(1, source, intermed, dest, allTowers)
		move(disks-1, intermed, source, dest, allTowers)
	}
}

func main() {
	height := 5
	tower1 := *newTower("Tower1", height)
	tower2 := *newTower("Tower2", 0)
	tower3 := *newTower("Tower3", 0)
	move(height, &tower1, &tower2, &tower3, []*tower{&tower1, &tower2, &tower3})
}
