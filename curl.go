package curl

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type Header map[string]string

func (h Header) Get(key string) string {
	return h[key]
}
func (h Header) Set(key, value string) {
	h[key] = value
}
func (h Header) Del(key string) {
	delete(h, key)
}

type Query map[string]string

func (q Query) Get(key string) string {
	return q[key]
}
func (q Query) Set(key, value string) {
	q[key] = value
}
func (q Query) Del(key string) {
	delete(q, key)
}

type Form map[string]string

func (f Form) Get(key string) string {
	return f[key]
}
func (f Form) Set(key, value string) {
	f[key] = value
}
func (f Form) Del(key string) {
	delete(f, key)
}

type Body string

func (b Body) UnmarshalParse(obj interface{}) error {
	return json.Unmarshal([]byte(b), obj)
}

type Curl struct {
	Url    string
	Method string
	Header Header
	Query  Query
	From   Form
	Body   Body
}

var (
	ErrParseFailed  = errors.New("parse url failed")
	ErrCurlSyntax   = errors.New("curl syntax error")
	ErrHeaderSyntax = errors.New("header syntax error")
	ErrQuerySyntax  = errors.New("query syntax error")
	ErrFromSyntax   = errors.New("from syntax error")
)

type Option string

const (
	// Header
	OptionShortHeader Option = "-H"
	OptionHeader      Option = "--header"

	// Method
	OptionShortMethod Option = "-X"
	OptionMethod      Option = "--request"

	// Body
	OptionShortBody Option = "-d"
	OptionBody      Option = "--data"
	OptionDataRow   Option = "--data-raw"

	// Form-data
	OptionShortForm Option = "-F"
	OptionForm      Option = "--form"

	// Location
	OptionShortLocation Option = "-L"
	OptionLocation      Option = "--location"
)

const (
	ContentTypeKey      = "Content-Type"
	ContentTypeText     = "text/plain"
	ContentTypeFormData = "multipart/form-data"
	ContentTypeJson     = "application/json"
)

func isUrl(s string) bool {
	return strings.HasPrefix(s, "http")
}

func isOption(s string) bool {
	return strings.HasPrefix(s, "-")
}

func parseHeader(header string) (hKey string, hVal string, err error) {
	pos := strings.IndexByte(header, ':')
	if pos == -1 {
		return "", "", ErrHeaderSyntax
	}

	return header[:pos], header[pos+1:], nil
}

func parseUrl(url string) (rawUrl string, query map[string]string, err error) {
	pos := strings.IndexByte(url, '?')
	if pos == -1 {
		return url, nil, nil
	}

	rawUrl = url[:pos]
	query = make(map[string]string)

	qKvs := strings.Split(url[pos+1:], "&")
	for _, kv := range qKvs {
		qPos := strings.IndexByte(kv, '=')
		if qPos == -1 {
			return "", nil, ErrQuerySyntax
		}

		query[kv[:qPos]] = kv[qPos+1:]
	}

	return rawUrl, query, nil
}

func ParseFormData(form string) (fKey string, fVal string, err error) {
	pos := strings.IndexByte(form, '=')
	if pos == -1 {
		return "", "", ErrFromSyntax
	}

	return form[:pos], form[pos+1:], nil
}

func Parse(url string) (*Curl, error) {
	curlCmd, err := GetArgsToken(url)
	if err != nil {
		return nil, err
	}

	if len(curlCmd) == 0 {
		return nil, ErrParseFailed
	}

	if strings.ToLower(curlCmd[0]) != "curl" {
		return nil, ErrCurlSyntax
	}

	curl := &Curl{
		Header: make(map[string]string),
		Query:  make(map[string]string),
		From:   make(map[string]string),
	}

	for idx := 1; idx < len(curlCmd); idx++ {
		if cmd := curlCmd[idx]; isOption(cmd) {
			switch Option(cmd) {
			case OptionShortHeader, OptionHeader:
				if idx+1 >= len(curlCmd) {
					return nil, ErrCurlSyntax
				}

				k, v, err := parseHeader(curlCmd[idx+1])
				if err != nil {
					return nil, err
				}

				curl.Header[k] = v

				idx += 1
			case OptionShortMethod, OptionMethod:
				if idx+1 >= len(curlCmd) {
					return nil, ErrCurlSyntax
				}
				curl.Method = curlCmd[idx+1]
				idx += 1
			case OptionShortBody, OptionBody, OptionDataRow:
				if idx+1 >= len(curlCmd) {
					return nil, ErrCurlSyntax
				}
				curl.Body = Body(curlCmd[idx+1])
				idx += 1
			case OptionShortForm, OptionForm:
				if idx+1 >= len(curlCmd) {
					return nil, ErrCurlSyntax
				}
				k, v, err := ParseFormData(curlCmd[idx+1])
				if err != nil {
					return nil, err
				}
				curl.From[k] = v
				idx += 1
			case OptionShortLocation, OptionLocation:
				continue
			default:
				return nil, ErrCurlSyntax
			}
		} else if isUrl(curlCmd[idx]) {
			curl.Url, curl.Query, err = parseUrl(curlCmd[idx])
			if err != nil {
				return nil, err
			}
		} else {
			fmt.Printf("unsupport curl cmd: %s, ignore it", cmd)
		}
	}

	// set default method
	if curl.Method == "" {
		if curl.Body != "" {
			curl.Method = http.MethodPost
		} else {
			curl.Method = http.MethodGet
		}
	}

	// set default content-type
	if _, ok := curl.Header[ContentTypeKey]; !ok {
		if curl.Body != "" {
			curl.Header[ContentTypeKey] = ContentTypeJson
		} else if curl.From != nil {
			curl.Header[ContentTypeKey] = ContentTypeFormData
		}
	}

	return curl, nil
}
