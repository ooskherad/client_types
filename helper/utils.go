package helper

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net"
	"reflect"
	"sort"
	"strconv"
	"time"
)

type intOrFloat interface {
	int | float32
}
type StringOrNumber interface {
	int | float32 | uint | int32 | int16 | float64 | int8 | string
}

func ConvertStringToDate(strTime string) time.Time {
	layout := "20060102"
	t, err := time.Parse(layout, strTime)

	if err != nil {
		log.Panicln(err)
	}
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}
func NowTimeString() string {
	t := time.Now()
	return t.Format("20060102")
}

// SetField convert map to struct
func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("no such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("provided value type didn't match obj field type")
	}

	structFieldValue.Set(val)
	return nil
}

func FillStruct(m map[string]interface{}, s interface{}) error {
	for k, v := range m {
		err := SetField(s, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func StructToList(model interface{}) []interface{} {
	r := reflect.ValueOf(model)
	values := make([]interface{}, r.NumField())
	for i := 0; i < r.NumField(); i++ {
		values[i] = r.Field(i).Interface()
	}
	return values
}

func Divide[T intOrFloat](number1 T, number2 T) float32 {
	if number2 != 0 {
		return float32(number1) / float32(number2)
	}
	return 0
}

func StringToInt(input string) int {
	return val(strconv.Atoi(input))
}

func val[T, U any](val T, _ U) T {
	return val
}
func ErrorHandler(err error) {
	if err, ok := err.(net.Error); ok && err.Timeout() {
		log.Println("Tabangoo Custom ErrorHandler: ", err.Error())
		return
	}
	if err != nil {
		log.Fatal("Tabangoo Custom ErrorHandler: ", err.Error())
	}
}
func StringValErrorHandler(val string, err error) string {
	if err != nil {
		log.Fatal("Tabangoo Custom Val ErrorHandler: ", err.Error())
	}
	return val
}
func TimeValErrorHandler(val time.Time, err error) time.Time {
	if err != nil {
		log.Fatal("Tabangoo Custom Time Val ErrorHandler: ", err.Error())
	}
	return val
}
func StringStringMapValErrorHandler(val map[string]string, err error) map[string]string {
	if err != nil {
		log.Fatal("Tabangoo Custom Map Val ErrorHandler: ", err.Error())
	}
	return val
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func GetIranTimeZone() *time.Location {
	loc, _ := time.LoadLocation("Asia/Tehran")
	return loc
}

// SubtractionStringArray Substring a2 from a1
func SubtractionStringArray(a1 []string, a2 []string) []string {
	var result []string
	i := 0
	j := 0
	if len(a2) == 0 {
		return a1
	}
	var unique []string
	for _, v := range a1 {
		skip := false
		for _, u := range unique {
			if v == u {
				skip = true
				break
			}
		}
		if !skip {
			unique = append(unique, v)
		}
	}
	a1 = unique
	sort.Strings(a1)
	sort.Strings(a2)
	var b []string
	for j < len(a1) {
		if a1[j] < a2[i] {
			result = append(result, a1[j])
			j++
			b = append(b, fmt.Sprintf("%d < %d", i, j))
		} else if a1[j] == a2[i] {
			j++
			if i < len(a2)-1 {
				i++
			}
			b = append(b, fmt.Sprintf("%d = %d", i, j))
		} else if a1[j] > a2[i] {
			if i < len(a2)-1 {
				i++
			} else {
				result = append(result, a1[j])
				j++
			}
			b = append(b, fmt.Sprintf("%d > %d", i, j))
		}
	}
	return result
}

func IsIn[T StringOrNumber](list []T, elem T) bool {
	for _, item := range list {
		if elem == item {
			return true
		}
	}
	return false
}

func GetWeekday(weekday int) string {
	weekdayStatus := map[int]string{0: "یکشنبه", 1: "دوشنبه", 2: "سه شنبه", 3: "چهارشنبه", 4: "پنج شنبه", 5: "جمعه", 6: "شنبه"}
	return weekdayStatus[weekday]
}
