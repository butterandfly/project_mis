package waitresses

import (
	"fmt"
	// "github.com/gorilla/mux"
	"../beers"
	"../buckets"
	"net/http"
	"time"
)

/*
type Mis struct {

}
*/

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
	fmt.Fprintln(w, "Post a mis!")
}

func DeleteMis(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Delete a mis!")
}
