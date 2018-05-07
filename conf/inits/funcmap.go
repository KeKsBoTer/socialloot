package inits

import (
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/KeKsBoTer/socialloot/lib"

	"github.com/astaxie/beego"
)

// various functions for template rendering
// functions are global and can be used in every template 

func init() {

	// format date to dd.mm.yyyy hh:mm and set timezone to local one
	beego.AddFuncMap("dateformat", func(in time.Time) string {
		return in.Local().Format("02.01.2006 15:04")
	})

	// get url for models like post,topic and user
	beego.AddFuncMap("URL", func(data interface{}) string {
		return lib.URLForItem(data)
	})

	// changes/adds the HTTP GET parameter in a URL
	// e.g. ChangeParam "localhost:8080/?test=2" "test" "new_value"
	// 		=> localhost:8080/?test=new_value
	beego.AddFuncMap("ChangeParam", func(urlPath, key, value string) string {
		urlParsed, err := url.Parse(urlPath)
		if err != nil {
			return urlPath
		}
		query := urlParsed.Query()
		query.Set(key, value)
		urlParsed.RawQuery = query.Encode()
		return urlParsed.String()
	})

	// extracts HTTP GET paramter from URL
	// e.g. GetParam "localhost:8080/?test=1234" "test"
	//		=> test
	beego.AddFuncMap("GetParam", func(urlPath, key string) string {
		urlParsed, err := url.Parse(urlPath)
		if err != nil {
			return ""
		}
		return urlParsed.Query().Get(key)
	})

	// returns host name from URL
	// e.g. host http://google.de/someurl?id=2
	//		=> google.de
	beego.AddFuncMap("host", func(urlPath string) string {
		if strings.HasPrefix(urlPath, "/") {
			return "socialloot"
		}
		urlParsed, err := url.Parse(urlPath)
		if err != nil {
			return ""
		}
		return urlParsed.Host
	})

	// gets URL for a page's favicon
	// if no schema is defined in URL https is used
	// e.g. favicon http://google.de/someurl?id=2
	//		=> http://google.de/favicon.ico
	beego.AddFuncMap("favicon", func(urlPath string) string {
		urlParsed, err := url.Parse(urlPath)
		if err != nil {
			return ""
		}
		scheme := urlParsed.Scheme
		if len(scheme) < 1 {
			scheme = "https" // use https as default
		}
		return scheme + "://" + urlParsed.Host + "/favicon.ico"
	})

	// cuts text after n characters
	beego.AddFuncMap("cut", func(text string, n int) string {
		if len(text) < n {
			return text
		}
		return text[:n]
	})

	// checks if the given value is a nil pointer or the value has length zero
	// for arrays, slices, strings etc.
	// It panics if item's Kind is not Array, Chan, Map, Slice, or String.
	beego.AddFuncMap("isempty", func(item interface{}) bool {
		if item == nil {
			return true
		}
		value := reflect.ValueOf(item)
		// get value from pointer
		for value.Kind() == reflect.Ptr || value.Kind() == reflect.Interface {
			value = value.Elem()
		}
		return value.Len() < 1
	})
}
