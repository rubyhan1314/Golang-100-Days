package main

import "time"

func main() {

}

type MyDuration = time.Duration

func (m MyDuration) SimpleSet(){ //cannot define new methods on non-local type time.Duration

}
