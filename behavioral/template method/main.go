package main

import "fmt"

func main() {

	smsOTP := &Sms{}
	o := Otp{
		iOtp: smsOTP,
	}
	err := o.genAndSendOTP(4)
	if err != nil {
		return
	}

	fmt.Println("")
	emailOTP := &Email{}
	o = Otp{
		iOtp: emailOTP,
	}
	err = o.genAndSendOTP(4)
	if err != nil {
		return
	}

}