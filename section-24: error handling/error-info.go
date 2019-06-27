package main

import (
	"fmt"
	"log"
)

type mathError struct {
	lat  string
	long string
	err  error
}

func (me mathError) Error() string {
	return fmt.Sprintf("An error occurred: %v, %v, %v", me.lat, me.long, me.err)
}

func error_info() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("math redux: square root of negative number %v", f)
		return 0, mathError{
			lat:  "50.2289 N",
			long: "99.4656 W",
			err:  nme,
		}
	}

	return 42, nil
}
