package waitresses

import (
	"fmt"
	// "github.com/gorilla/mux"
	"net/http"
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
	/*
		router.HandleFunc(misResourceUrl, PostMis).
			Methods("POST")
		router.HandleFunc(fmt.Sprint(misResourceUrl, "/{id}"), DeleteMis).
			Methods("DELETE")
	*/
}

func GetMis(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Get a mis.")
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
