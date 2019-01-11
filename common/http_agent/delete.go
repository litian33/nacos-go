package http_agent

import (
	"log"
	"net/http"
	"time"
)

/**
*
* @description : 
*
* @author : codezhang
*
* @create : 2019-01-08 14:08
**/

func delete(path string, header http.Header, timeoutMs uint64) (response *http.Response, err error) {
	client := http.Client{}
	client.Timeout = time.Millisecond * time.Duration(timeoutMs)
	request, errNew := http.NewRequest(http.MethodDelete, path, nil)
	if errNew != nil {
		log.Println(errNew)
		err = errNew
	}
	if err == nil {
		request.Header = header
		resp, errDo := client.Do(request)
		if errDo != nil {
			log.Println(errDo)
			err = errDo
		} else {
			response = resp
		}
	}
	return
}
