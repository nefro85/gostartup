package phone

import (
	"fmt"
)

const (
	_          = iota
	NEXUS_UID  = iota
	IPHONE_UID = iota
)

type Sms string
type PhoneNumber []int
type Os string

type Phone interface {
	Call(PhoneNumber)
	Text(Sms)
}

type Hardware struct {
	System Os
	Imei   int
}

type Nexus struct {
	Hw Hardware
}

func (nex *Nexus) Call(number PhoneNumber) {
	fmt.Printf("%#v Call : %v \n" ,nex, number)

}

func (nex *Nexus) Text(sms Sms) {

}

type Iphone struct {
	Hw Hardware
}

func (ip *Iphone) Call(number PhoneNumber) {
	fmt.Printf("%#v Call : %v \n" ,ip, number)
}

func (ip *Iphone) Text(sms Sms) {

}
