package gosrisdk

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	SriLayout = "2006-01-02 15:04:05.0"
)

type SriDate struct {
	time.Time
}

func (this *SriDate) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), "\"")
	if s == "null" || s == "" {
		return nil
	}
	parsed, err := time.Parse(SriLayout, s)
	if err != nil {
		return err
	}
	this.Time = parsed
	return nil
}

type SriBool bool

func (this *SriBool) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), "\"")
	if s == "null" || s == "" {
		return nil
	}
	parsed, err := strconv.ParseBool(s)
	if err != nil {
		switch s {
		case "SI":
			*this = true
			return nil
		case "NO":
			*this = false
			return nil
		default:
			return errors.New("Formato de boolean incorrecto")
		}
	}
	fmt.Println("Era bool")
	*this = SriBool(parsed)
	return nil
}
