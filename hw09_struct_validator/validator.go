package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type (
	ValidationError struct {
		Field string
		Err   error
	}

	ValidationErrors []ValidationError
)

var (
	ErrStringLen                 = errors.New("string length is invalid")
	ErrStringRegex               = errors.New("string is not compatible with regex")
	ErrStringSet                 = errors.New("string is not compatible with set")
	ErrNumberMin                 = errors.New("number is less than min value")
	ErrNumberMax                 = errors.New("number is greater then max value")
	ErrNumberSet                 = errors.New("number is not compatible with set")
	ErrUnsupportedInterfaceValue = errors.New("received interface is not a struct")
)

var (
	reLen    = regexp.MustCompile(`(?:^|\|)len:(\d+)(?:$|\|.*)`)
	reRegexp = regexp.MustCompile(`(?:^|\|)regexp:(.+)(?:$|\|.*)`)
	reInStr  = regexp.MustCompile(`(?:^|\|)in:(.+)(?:$|\|.*)`)
	reMin    = regexp.MustCompile(`(?:^|\|)min:(\d+)(?:$|\|.*)`)
	reMax    = regexp.MustCompile(`(?:^|\|)max:(\d+)(?:$|\|.*)`)
	reInInt  = regexp.MustCompile(`(?:^|\|)in:([\d,]+)(?:$|\|.*)`)
)

func (v ValidationErrors) Error() string {
	var errorText strings.Builder
	fmt.Fprintf(&errorText, "validation errors count: %v; ", len(v))
	for i, vErr := range v {
		fmt.Fprintf(&errorText, "%v: field '%s' with error '%s';", i, vErr.Field, vErr.Err)
	}
	return errorText.String()
}

func Validate(v interface{}) error {
	var valErr ValidationErrors
	rt := reflect.TypeOf(v)
	rv := reflect.ValueOf(v)

	if rt.Kind() != reflect.Struct {
		return ErrUnsupportedInterfaceValue
	}

	for i := 0; i < rt.NumField(); i++ {
		fieldV := rv.Field(i)
		fieldT := rt.Field(i)
		val, ok := fieldT.Tag.Lookup("validate")
		if !ok {
			// skip fields without validate tag
			continue
		}
		switch fieldV.Kind().String() {
		case "string":
			err := validateString(val, fieldV, fieldT, &valErr)
			if err != nil {
				return err
			}
		case "int":
			err := validateInt(val, fieldV, fieldT, &valErr)
			if err != nil {
				return err
			}
		case "slice":
			sliceType := fieldV.Type().String()
			sliceLen := fieldV.Len()
			for i := 0; i < sliceLen; i++ {
				switch sliceType {
				case "[]string":
					err := validateString(val, fieldV.Index(i), fieldT, &valErr)
					if err != nil {
						return err
					}
				case "[]int":
					err := validateInt(val, fieldV.Index(i), fieldT, &valErr)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	if len(valErr) > 0 {
		return valErr
	}
	return nil
}

func validateString(val string, fieldV reflect.Value, fieldT reflect.StructField, valErr *ValidationErrors) error {
	// check on string length
	lenCond := reLen.FindStringSubmatch(val)
	if len(lenCond) > 0 {
		lenCondInt, err := strconv.Atoi(lenCond[1])
		if err != nil {
			return err
		}
		if fieldV.Len() != lenCondInt {
			*valErr = append(*valErr, ValidationError{Field: fieldT.Name, Err: ErrStringLen})
		}
	}
	// check on string regex
	regexpCond := reRegexp.FindStringSubmatch(val)
	if len(regexpCond) > 0 {
		matched, err := regexp.Match(regexpCond[1], []byte(fieldV.String()))
		if err != nil {
			return err
		}
		if !matched {
			*valErr = append(*valErr, ValidationError{Field: fieldT.Name, Err: ErrStringRegex})
		}
	}
	// check on string compatible with set of values
	inCond := reInStr.FindStringSubmatch(val)
	if len(inCond) > 0 {
		fmt.Printf("[INFO] checking condition for %s\n", inCond[1])
		inCondSplitted := strings.Split(inCond[1], ",")
		valueInList := false
		for _, value := range inCondSplitted {
			if value == fieldV.String() {
				valueInList = true
			}
		}
		if !valueInList {
			*valErr = append(*valErr, ValidationError{Field: fieldT.Name, Err: ErrStringSet})
		}
	}
	return nil
}

func validateInt(val string, fieldV reflect.Value, fieldT reflect.StructField, valErr *ValidationErrors) error {
	// check on min number value
	minCond := reMin.FindStringSubmatch(val)
	if len(minCond) > 0 {
		minCondInt, err := strconv.Atoi(minCond[1])
		if err != nil {
			return err
		}
		if int(fieldV.Int()) < minCondInt {
			*valErr = append(*valErr, ValidationError{Field: fieldT.Name, Err: ErrNumberMin})
		}
	}
	// check on max number value
	maxCond := reMax.FindStringSubmatch(val)
	if len(maxCond) > 0 {
		maxCondInt, err := strconv.Atoi(maxCond[1])
		if err != nil {
			return err
		}
		if int(fieldV.Int()) > maxCondInt {
			*valErr = append(*valErr, ValidationError{Field: fieldT.Name, Err: ErrNumberMax})
		}
	}
	// check on number in list
	inCond := reInInt.FindStringSubmatch(val)
	if len(inCond) > 0 {
		inCondSplitted := strings.Split(inCond[1], ",")
		valueInList := false
		for _, value := range inCondSplitted {
			valueInt, err := strconv.Atoi(value)
			if err != nil {
				return err
			}
			if int64(valueInt) == fieldV.Int() {
				valueInList = true
			}
		}
		if !valueInList {
			*valErr = append(*valErr, ValidationError{Field: fieldT.Name, Err: ErrNumberSet})
		}
	}
	return nil
}
