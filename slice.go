package main

import(
  "log"
)

func main() {
  slice := []int{1, 2, 3}
  sliceTwo := []int{55, 66, 73}

  log.Println("start slice: ", slice)
  log.Println("start sliceTwo: ", sliceTwo)

  // add item to a slice
  slice = append(slice, sliceTwo...)
  log.Println("after adding one slice: ", slice)

  // make a copy of the slice
  sliceThree := append([]int(nil), slice...)
  log.Println("copy a slice: ", sliceThree)

  // copy slice to end of itself
  log.Println("before append to self: ", slice)
  slice = append(slice, slice...)

  log.Println("after append to self: ", slice)

  // strings
  slash := "/usr/ken"
  log.Println("first character: ", slash[0])

  // substring
  log.Println("slash[0:4] ", slash[0:4])


}
