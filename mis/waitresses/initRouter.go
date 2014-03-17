package waitresses

import (
	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func GetRouter() *mux.Router {
	return router
}

/*
func init() {
	fmt.Println("init in waitresses")
}

func InitRouter(r *mux.Router) {
	// router = r
}

func newRouter() int {
	fmt.Println("newRouter func")
	return 0
}

var router = newRouter()
*/
