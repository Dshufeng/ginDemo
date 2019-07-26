package utils

import (
	"time"
	"runtime"
	"fmt"
	"path/filepath"
	"strings"
	"encoding/json"
)

type ErrString struct {
	s string
}
type ErrInfo struct {
	Time string `json:"time"`
	Alarm string `json:"alarm"`
	Filename string `json:"filename"`
	Line int `json:"line"`
	Message string `json:"message"`
	Function string `json:"function"`
}

func (err *ErrString) Error()  string{
	return err.s
}

func New(text string) error {
	alarm("INFO",text)
	return &ErrString{text}
}

func WeChat(text string) error  {
	alarm("WX",text)
	return &ErrString{text}
}

func alarm(level string,str string)  {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fileName,line,functionName := "?",0,"?"

	pc,fileName,line,ok := runtime.Caller(2)
	if ok{
		functionName = runtime.FuncForPC(pc).Name()
		functionName = filepath.Ext(functionName)
		functionName = strings.TrimPrefix(functionName,".")
	}

	var info = ErrInfo{
		Time:currentTime,
		Alarm:level,
		Filename:fileName,
		Line:line,
		Message:str,
		Function:functionName,
	}

	jsonInfo,err := json.Marshal(info)
	if err != nil{
		fmt.Println("json marshal err: ",err)
	}
	errJsonInfo := string(jsonInfo)
	fmt.Println(errJsonInfo)

	if level == "EMAIL" {

	}else if level == "SMS" {

	}else if level == "WX"{

	}
}
