package main

import (
	"errors"
	"fmt"
)

var (
	errEmplCount = errors.New("employees  must be != 0 and <= 1000")
	errOffCount  = errors.New("number of offices must be != 0 and <= 1000.")
	errOp        = errors.New("operator must be <= or >=")
)

var (
	opLen = 2
)

func main() {

	var offCount uint
	_, err := fmt.Scan(&offCount)
	if err != nil {
		fmt.Println(err)

		return
	}

	if offCount == 0 || offCount > 1000 {
		fmt.Println("offCount err: ", errOffCount)

		return
	}

	var i uint
	for ; i < offCount; i++ {

		var EmplCount uint
		_, err := fmt.Scan(&EmplCount)
		if err != nil {
			fmt.Println(err)

			return
		}

		if EmplCount == 0 || EmplCount > 1000 {
			fmt.Println("EmplCount err: ", errEmplCount)

			return
		}

		var (
			minTemp uint = 15
			maxTemp uint = 30
		)

		var j uint
		for ; j < EmplCount; j++ {

			var op string
			_, err := fmt.Scan(&op)
			if err != nil {
				fmt.Println(err)

				return
			}

			if (len(op) != opLen) || (op != "<=" && op != ">=") {
				fmt.Println("op err: ", errOp)

				return
			}

			var temp uint
			_, err = fmt.Scan(&temp)
			if err != nil {
				fmt.Println(err)

				return
			}

		}
	}
}
