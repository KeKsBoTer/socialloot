package inits

import (
	"html/template"
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/KeKsBoTer/socialloot/lib"
	"github.com/beego/i18n"
	"github.com/mattn/go-runewidth"

	"github.com/astaxie/beego"
)

func init() {

	beego.AddFuncMap("i18n", i18n.Tr)

	beego.AddFuncMap("i18nja", func(format string, args ...interface{}) string {
		return i18n.Tr("ja-JP", format, args...)
	})

	beego.AddFuncMap("datenow", func(format string) string {
		return time.Now().Add(time.Duration(9) * time.Hour).Format(format)
	})

	beego.AddFuncMap("dateformat", func(in time.Time) string {
		in = in.Add(time.Duration(9) * time.Hour)
		return in.Format("01/02/2006 15:04")
	})

	beego.AddFuncMap("qescape", func(in string) string {
		return url.QueryEscape(in)
	})

	beego.AddFuncMap("nl2br", func(in string) string {
		return strings.Replace(in, "\n", "<br>", -1)
	})

	beego.AddFuncMap("tostr", func(in interface{}) string {
		return lib.ToStr(reflect.ValueOf(in).Interface())
	})

	beego.AddFuncMap("first", func(in interface{}) interface{} {
		return reflect.ValueOf(in).Index(0).Interface()
	})

	beego.AddFuncMap("last", func(in interface{}) interface{} {
		s := reflect.ValueOf(in)
		return s.Index(s.Len() - 1).Interface()
	})

	beego.AddFuncMap("truncate", func(in string, length int) string {
		return runewidth.Truncate(in, length, "...")
	})

	beego.AddFuncMap("cleanurl", func(in string) string {
		return strings.Trim(strings.Trim(in, " "), "　")
	})

	beego.AddFuncMap("append", func(data map[interface{}]interface{}, key string, value interface{}) template.JS {
		if _, ok := data[key].([]interface{}); !ok {
			data[key] = []interface{}{value}
		} else {
			data[key] = append(data[key].([]interface{}), value)
		}
		return template.JS("")
	})

	beego.AddFuncMap("appendmap", func(data map[interface{}]interface{}, key string, name string, value interface{}) template.JS {
		v := map[string]interface{}{name: value}

		if _, ok := data[key].([]interface{}); !ok {
			data[key] = []interface{}{v}
		} else {
			data[key] = append(data[key].([]interface{}), v)
		}
		return template.JS("")
	})

	beego.AddFuncMap("URL", func(data interface{}) string {
		return lib.URLForItem(data)
	})

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

	beego.AddFuncMap("GetParam", func(urlPath, key string) string {
		urlParsed, err := url.Parse(urlPath)
		if err != nil {
			return ""
		}
		return urlParsed.Query().Get(key)
	})

	beego.AddFuncMap("host", func(urlPath string) string {
		urlParsed, err := url.Parse(urlPath)
		if err != nil {
			return ""
		}
		return urlParsed.Host
	})
	beego.AddFuncMap("favicon", func(urlPath string) string {
		urlParsed, err := url.Parse(urlPath)
		if err != nil {
			return ""
		}
		return "https://" + urlParsed.Host + "/favicon.ico"
	})
	beego.AddFuncMap("cut", func(text string, length int) string {
		if len(text) < length {
			return text
		}
		return text[:length]
	})
	beego.AddFuncMap("choice", func(path string) string {
		urlParsed, err := url.Parse(path)
		if err != nil {
			return ""
		}
		matched, err := regexp.MatchString("*/topic/[a-z0-9]/[a-z]", urlParsed.Path)
		if !matched || err != nil {
			return ""
		}
		splitted := strings.Split(urlParsed.Path, "/")
		return splitted[0]
	})
}
