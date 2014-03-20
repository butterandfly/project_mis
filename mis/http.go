package mis

import (
	// "github.com/gorilla/mux"
	"net/http"
	// "time"
	"./waitresses"
)

func init() {
	/*
		MisResourceUrl := "/mis"

		router := mux.NewRouter()

		// 路由
		router.HandleFunc(fmt.Sprint(articleResourceUrl, "/{id}"), GetArticle).
			Methods("GET")
		router.HandleFunc(articleResourceUrl, PostArticle).
			Methods("POST")
		router.HandleFunc(fmt.Sprint(articleResourceUrl, "/{id}"), DeleteArticle).
			Methods("DELETE")
	*/

	// http.Handle("/", router)

	http.Handle("/", waitresses.GetRouter())
}

/*

func GetArticle(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	err := r.ParseForm()

	start := r.FormValue("start")
	fmt.Fprint(w, start)

	if err != nil {
		oops(err, w)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := mux.Vars(r)["id"]

	key, err := datastore.DecodeKey(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	articleEntity := &model.ArticleEntity{}
	err = datastore.Get(c, key, articleEntity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Article: %v", articleEntity)
}

func PostArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		c := appengine.NewContext(r)

		r.ParseForm()
		article := &model.Article{
			Title:      r.Form.Get("title"),
			Content:    r.Form.Get("content"),
			CreateTime: time.Now(),
			ID:         "",
		}

		ds.AddArticle(article, c)

		w.WriteHeader(http.StatusCreated)

		fmt.Fprintf(w, "Success!")
	}
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	id := mux.Vars(r)["id"]

	if id == "all" {
		// TODO: Delete all articles.
		if err := DeleteAllArticle(c); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusNoContent)
		return
	}

	if err := DeleteArticleById(c, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteArticleById(c appengine.Context, id string) error {
	key, err := datastore.DecodeKey(id)
	if err != nil {
		return err
	}

	err = datastore.Delete(c, key)
	if err != nil {
		return err
	}

	return nil
}

func DeleteAllArticle(c appengine.Context) error {
	articleEntitys := make([]model.ArticleEntity, 0)

	q := datastore.NewQuery("ArticleEntity").KeysOnly()
	keys, err := q.GetAll(c, &articleEntitys)
	if err != nil {
		return err
	}

	if err = datastore.DeleteMulti(c, keys); err != nil {
		return err
	}

	return nil
}
*/
