package utils

import (
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func Get(url string, response interface{}, tlsConfig *tls.Config, setHeader func(req *http.Request) error) error {
	return do("GET", url, nil, response, tlsConfig, setHeader, json.Unmarshal)
}

func Post(url string, body io.Reader, response interface{}, tlsConfig *tls.Config, setHeader func(req *http.Request) error) error {
	return do("POST", url, body, response, tlsConfig, setHeader, json.Unmarshal)
}

func GetByWebService(url string, response interface{}, tlsConfig *tls.Config, setHeader func(req *http.Request) error) error {
	return do("GET", url, nil, response, tlsConfig, func(req *http.Request) error {
		req.Header.Set("Content-Type", "text/xml; charset=utf-8")
		if setHeader != nil {
			return setHeader(req)
		}
		return nil
	}, xml.Unmarshal)
}

func PostByWebService(url string, body io.Reader, response interface{}, tlsConfig *tls.Config, setHeader func(req *http.Request) error) error {
	return do("POST", url, body, response, tlsConfig, func(req *http.Request) error {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		if setHeader != nil {
			return setHeader(req)
		}
		return nil
	}, xml.Unmarshal)
}

var StatusUnauthorized = errors.New("Unauthorized")

func do(method, url string, body io.Reader, response interface{}, tlsConfig *tls.Config,
	setHeader func(req *http.Request) error, unmarshal func(data []byte, v interface{}) error) error {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}

	if setHeader != nil {
		err = setHeader(req)
		if err != nil {
			return err
		}
	}

	client := http.DefaultClient
	if tlsConfig != nil {
		client.Transport = &http.Transport{TLSClientConfig: tlsConfig}
	}

	resp, err := client.Do(req)
	if err != nil {
		write([]byte(err.Error()))
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return StatusUnauthorized
	}

	if resp.StatusCode != http.StatusOK {
		write([]byte(resp.Status))
		return errors.New(url + ":" + resp.Status)
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	write([]byte(fmt.Sprintf("url:%s,resp:%s", url, buf)))
	if unmarshal != nil {
		err = unmarshal(buf, response)
	} else {
		err = json.Unmarshal(buf, response)
	}

	if err != nil {
		return err
	}
	return nil
}

type WriteLog func(data []byte) error

var (
	f        *os.File
	writeLog WriteLog
)

func init() {
	writeLog = writeConsole
}

func SetWriteLogFunc(w WriteLog) {
	writeLog = w
}

func SetLogFileName(name string) {
	if f != nil {
		f.Close()
	}
	var err error
	f, err = os.OpenFile(name, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	writeLog = writeFile
}

func write(data []byte) error {
	if writeLog != nil {
		return writeLog(data)
	}
	return nil
}

func writeFile(data []byte) error {

	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err != nil {
		return err
	}
	n, err = f.Write([]byte("\r\n"))
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	return err
}

func writeConsole(data []byte) error {
	fmt.Println(string(data))
	return nil
}
