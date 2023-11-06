package main

type SeatManager struct {
	seats []int
	min   int
}

func Constructor(n int) SeatManager {
	return SeatManager{
		seats: make([]int, n),
	}
}

func (this *SeatManager) Reserve() int {
	for i := this.min; i < len(this.seats); i++ {
		if this.seats[i] == 0 {
			this.seats[i] = 1
			this.min = i
			return i + 1
		}
	}
	return 0
}

func (this *SeatManager) Unreserve(seatNumber int) {
	this.seats[seatNumber-1] = 0
	if this.min > seatNumber-1 {
		this.min = seatNumber - 1
	}
}

/**
* Your SeatManager object will be instantiated and called as such:
* obj := Constructor(n);
* param_1 := obj.Reserve();
* obj.Unreserve(seatNumber);
 */
