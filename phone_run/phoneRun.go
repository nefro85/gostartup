package main

import "gostartup/phone"

func main() {

	devices := []phone.Phone{
		&phone.Nexus{
			Hw: phone.Hardware{
				System: "Android",
				Imei:   phone.NEXUS_UID,
			},
		},
		&phone.Iphone{
			Hw: phone.Hardware{
				System: "iOS",
				Imei:   phone.IPHONE_UID,
			},
		},
	}

	for _, dev := range devices {

		dev.Call([]int{2, 4, 6, 3})
	}

}
