package waitresses

import (
	"fmt"
	// "github.com/gorilla/mux"
	"../beers"
	"../buckets"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

func init() {
	fmt.Println("init in mis.go")

	var misResourceUrl = "/mis"

	// 使用路由设置对应的方法
	router.HandleFunc(fmt.Sprint(misResourceUrl, "/{id}"), GetMis).
		Methods("GET")
	router.HandleFunc(misResourceUrl, PostMis).
		Methods("POST")
	router.HandleFunc(fmt.Sprint(misResourceUrl, "/{id}"), DeleteMis).
		Methods("DELETE")
}

func GetMis(w http.ResponseWriter, r *http.Request) {
	/*
		err := beers.AddMis()
		if err != nil {
			fmt.Fprintln(w, err.Error())
			return
		}
	*/

	mis := &beers.Mis{
		"The first Mis.",
		"It's just a test.",
		time.Now(),
	}

	err := buckets.SharedMisBucket.Add(mis)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	fmt.Fprintln(w, "Success!")
	/*
		err := r.ParseForm()

		start := r.FormValue("start")
		fmt.Fprint(w, start)

		if err != nil {
			oops(err, w)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id := mux.Vars(r)["id"]
	*/
}

func PostMis(w http.ResponseWriter, r *http.Request) {
	var err error

	// 转换post请求中的json
	decoder := json.NewDecoder(r.Body)
	mis := beers.Mis{}
	err = decoder.Decode(&mis)
	if err != nil {
		// 请求的数据格式有误
		JsonResponse{"Error": err.Error()}.Output(w, http.StatusBadRequest)
		return
	}

	// 判断字段是否合法
	if strings.TrimSpace(mis.Title) == "" {
		// 请求有误
		JsonResponse{"Error": "Title不能为空"}.Output(w, http.StatusBadRequest)
		return
	}

	// 添加到数据库
	err = buckets.SharedMisBucket.Add(mis)
	if err != nil {
		// 内部错误
		JsonResponse{"Error": err.Error()}.Output(w, http.StatusInternalServerError)
		return
	}

	// 创建成功
	JsonResponse{"Title": mis.Title}.Output(w, http.StatusCreated)
}

func DeleteMis(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Delete a mis!")
}
