package errors

import (
	"errors"
	errors2 "github.com/go-kratos/kratos/v2/errors"
	"log"
	"fmt"
)


func Err(){
	fmt.Println("hahah")
	s := &errors2.StatusError{Code: 1}
	st := &errors2.StatusError{Code: 2}

	if errors.Is(s, st) {
		log.Printf("error is  match: %+v -> %+v", s, st)
	}else{
		log.Printf("error is not match: %+v -> %+v", s, st)

	}

	s.Code = 1
	st.Code = 1
	if !errors.Is(s, st) {
		log.Printf("error is not match: %+v -> %+v", s, st)
	}else{
		log.Printf("error is  match: %+v -> %+v", s, st)

	}

	s.Reason = "test_reason"
	st.Reason = "test_reason"

	if !errors.Is(s, st) {
		log.Printf("error is not match: %+v -> %+v", s, st)
	}else{
		log.Printf("error is match: %+v -> %+v", s, st)

	}
	log.Printf("s error reason: %+v",  errors2.Reason(s))

	if errors2.Reason(s) != "test_reason" {
		log.Printf("error is not match: %+v -> %+v", s, st)
	}
}
