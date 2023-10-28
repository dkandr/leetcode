package main

// type MyHashMap struct {
// 	data []int
// }

// func Constructor() MyHashMap {
// 	var h MyHashMap
// 	h.data = make([]int, 10000000, 10000000)

// 	for i := 0; i < 10000000; i++ {
// 		h.data[i] = -1
// 	}

// 	return h
// }

// func (this *MyHashMap) Put(key int, value int) {
// 	if key >= 10000000 {
// 		return
// 	}

// 	this.data[key] = value
// }

// func (this *MyHashMap) Get(key int) int {
// 	if key >= 10000000 {
// 		return -1
// 	}

// 	return this.data[key]
// }

// func (this *MyHashMap) Remove(key int) {
// 	if key >= 10000000 {
// 		return
// 	}

// 	this.data[key] = -1
// }

type MyHashMap struct {
}

func Constructor() MyHashMap {
	var h MyHashMap

	return h
}

func (this *MyHashMap) Put(key int, value int) {

}

func (this *MyHashMap) Get(key int) int {

	return 1
}

func (this *MyHashMap) Remove(key int) {

}
