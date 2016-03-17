package phone

const (
	_ = iota
	NEXUS_UID = iota
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
	Imei int
}



type Nexus struct {
	Hardware
}


func (ip Nexus) Call(voice PhoneNumber) {

}

func (ip Nexus)Text(sms Sms)  {

}


type Iphone struct {
	Hardware
}

func (ip Iphone) Call(voice PhoneNumber) {

}

func (ip Iphone)Text(sms Sms)  {

}



