package lib

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/KeKsBoTer/socialloot/models"
	"github.com/astaxie/beego"
)

func URLForItem(data interface{}) string {
	rv := reflect.ValueOf(data)
	// if data is a pointer, get the value to make type switching possible
	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		rv = rv.Elem()
	}
	data = rv.Interface()
	switch data.(type) {
	case models.Post:
		post := data.(models.Post)
		return beego.URLFor("PostController.Get", ":topic", post.Topic.Name, ":post", post.Id)
	case models.Topic:
		topic := data.(models.Topic)
		return beego.URLFor("TopicController.Get", ":topic", topic.Name)
	case models.User:
		user := data.(models.User)
		return beego.URLFor("UserController.Get", ":user", user.Name)
	default:
		beego.Error("Cannot create URL for:", reflect.TypeOf(data))
		return "/"
	}
}

func NumberEncode(number string, alphabet []byte) string {
	token := make([]byte, 0, 12)
	x, ok := new(big.Int).SetString(number, 10)
	if !ok {
		return ""
	}
	y := big.NewInt(int64(len(alphabet)))
	m := new(big.Int)
	for x.Sign() > 0 {
		x, m = x.DivMod(x, y, m)
		token = append(token, alphabet[int(m.Int64())])
	}
	return string(token)
}

func NumberDecode(token string, alphabet []byte) string {
	x := new(big.Int)
	y := big.NewInt(int64(len(alphabet)))
	z := new(big.Int)
	for i := len(token) - 1; i >= 0; i-- {
		v := bytes.IndexByte(alphabet, token[i])
		z.SetInt64(int64(v))
		x.Mul(x, y)
		x.Add(x, z)
	}
	return x.String()
}

// Random generate string
func GetRandomString(n int) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

// convert string to specify type

type StrTo string

func (f *StrTo) Set(v string) {
	if v != "" {
		*f = StrTo(v)
	} else {
		f.Clear()
	}
}

func (f *StrTo) Clear() {
	*f = StrTo(0x1E)
}

func (f StrTo) Exist() bool {
	return string(f) != string(0x1E)
}

func (f StrTo) Bool() (bool, error) {
	if f == "on" {
		return true, nil
	}
	return strconv.ParseBool(f.String())
}

func (f StrTo) Float32() (float32, error) {
	v, err := strconv.ParseFloat(f.String(), 32)
	return float32(v), err
}

func (f StrTo) Float64() (float64, error) {
	return strconv.ParseFloat(f.String(), 64)
}

func (f StrTo) Int() (int, error) {
	v, err := strconv.ParseInt(f.String(), 10, 32)
	return int(v), err
}

func (f StrTo) Int8() (int8, error) {
	v, err := strconv.ParseInt(f.String(), 10, 8)
	return int8(v), err
}

func (f StrTo) Int16() (int16, error) {
	v, err := strconv.ParseInt(f.String(), 10, 16)
	return int16(v), err
}

func (f StrTo) Int32() (int32, error) {
	v, err := strconv.ParseInt(f.String(), 10, 32)
	return int32(v), err
}

func (f StrTo) Int64() (int64, error) {
	v, err := strconv.ParseInt(f.String(), 10, 64)
	return int64(v), err
}

func (f StrTo) Uint() (uint, error) {
	v, err := strconv.ParseUint(f.String(), 10, 32)
	return uint(v), err
}

func (f StrTo) Uint8() (uint8, error) {
	v, err := strconv.ParseUint(f.String(), 10, 8)
	return uint8(v), err
}

func (f StrTo) Uint16() (uint16, error) {
	v, err := strconv.ParseUint(f.String(), 10, 16)
	return uint16(v), err
}

func (f StrTo) Uint32() (uint32, error) {
	v, err := strconv.ParseUint(f.String(), 10, 32)
	return uint32(v), err
}

func (f StrTo) Uint64() (uint64, error) {
	v, err := strconv.ParseUint(f.String(), 10, 64)
	return uint64(v), err
}

func (f StrTo) Md5() string {
	h := md5.New()
	h.Write([]byte(f.String()))
	return hex.EncodeToString(h.Sum(nil))
}

func (f StrTo) String() string {
	if f.Exist() {
		return string(f)
	}
	return ""
}

func (f StrTo) MultiWord() []string {
	str := f.String()

	str = strings.Replace(str, "\u3000", " ", -1)
	str = strings.Replace(str, "、", " ", -1)
	str = strings.Replace(str, ",", " ", -1)
	str = strings.Replace(str, "【", " ", -1)
	str = strings.Replace(str, "】", " ", -1)

	re, _ := regexp.Compile(`\s{2,}`)
	str = re.ReplaceAllString(str, " ")

	str = strings.Trim(str, " ")

	return strings.Split(str, " ")
}

// convert any type to string
func ToStr(value interface{}, args ...int) (s string) {
	switch v := value.(type) {
	case bool:
		s = strconv.FormatBool(v)
	case float32:
		s = strconv.FormatFloat(float64(v), 'f', argInt(args).Get(0, -1), argInt(args).Get(1, 32))
	case float64:
		s = strconv.FormatFloat(v, 'f', argInt(args).Get(0, -1), argInt(args).Get(1, 64))
	case int:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int8:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int16:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int32:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int64:
		s = strconv.FormatInt(v, argInt(args).Get(0, 10))
	case uint:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint8:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint16:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint32:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint64:
		s = strconv.FormatUint(v, argInt(args).Get(0, 10))
	case string:
		s = v
	case []byte:
		s = string(v)
	default:
		s = fmt.Sprintf("%v", v)
	}
	return s
}

type argInt []int

func (a argInt) Get(i int, args ...int) (r int) {
	if i >= 0 && i < len(a) {
		r = a[i]
	}
	if len(args) > 0 {
		r = args[0]
	}
	return
}
