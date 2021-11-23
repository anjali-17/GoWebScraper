package main

import (
	"encoding/json"
	"github.com/gocolly/colly"
	"github.com/gorilla/mux"
	"net/http"
)

type Movie struct{
	Title 			string  `json:"title"`
	ReleaseYear 	string	`json:"release_year"`
	Actors			[]string	`json:"actors"`
}

func getData(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	movieId := vars["amazon_id"]
	c := colly.NewCollector(
	colly.AllowedDomains("www.amazon.de"),
	)
	var title, detail string
	var movie Movie
	c.OnHTML("._2hu-aV",  func(e *colly.HTMLElement){
		title = e.ChildText("h1")
		movie.Title = title
		detail = e.ChildText("span")
		e.ForEach(".XqYSS8", func(i int, element *colly.HTMLElement) {
			if i == 1 {
				movie.ReleaseYear = element.Text
			}
		})
		var actor []string
		e.ForEach("._266mZB", func(i int, el *colly.HTMLElement){
			el.ForEach("dd", func(i int, element *colly.HTMLElement) {
				if i == 1 {
					element.ForEach("a", func(i int, ele *colly.HTMLElement) {
						actor =  append(actor, ele.Text)
					})
				}
			})
		})
		movie.Actors = actor
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		response.WriteHeader(http.StatusNotFound)
	})

	c.Visit("http://www.amazon.de/gp/product/" + movieId)
	jsonBytes, err  := json.Marshal(movie)
	if err != nil{
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
	} else {
		response.Header().Add("content-type", "application/json")
		response.WriteHeader(http.StatusOK)
		response.Write(jsonBytes)
	}
}

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/movie/amazon/{amazon_id}", getData).Methods("GET")
	err:= http.ListenAndServe(":8080", router)
	if err!= nil {
		panic(err)
	}
}