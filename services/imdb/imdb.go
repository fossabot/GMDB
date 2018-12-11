// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package imdb

import (
	"log"

	"gmdb/models"
	"gmdb/services/common"

	"github.com/puerkitobio/goquery"
)

type IMDB struct {
	Name    string
	Request models.SearchRequest
}

func New(name string, request models.SearchRequest) *IMDB {
	return &IMDB{
		Name:    name,
		Request: request,
	}
}

func (s *IMDB) SearchMovie(request *models.SearchRequest) *models.SearchResponse {

	url := "https://www.imdb.com/find?q=" + request.Title + "&s=tt"

	rq, err := GetSearchMovies(services.GetDocumentFromURL(url))
	//year, id\ exactsearch

	if err != nil {
		log.Fatalln("nil")
	}

	return rq
}

func (s *IMDB) GetMovie() (*models.Movie, error) {

	urlTL := s.Request.URL + "/taglines"
	urlPS := s.Request.URL + "/plotsummary"
	urlPK := s.Request.URL + "/keywords"
	urlPG := s.Request.URL + "/parentalguide"

	movie := new(models.Movie)

	//TODO: Optimization for spesific arguments

	mi, err1 := GetMovieInfo(services.GetDocumentFromURL(s.Request.URL))

	tl, err2 := GetTagline(services.GetDocumentFromURL(urlTL))
	ps, err3 := GetPlotSummary(services.GetDocumentFromURL(urlPS))
	pk, err4 := GetPlotKeywords(services.GetDocumentFromURL(urlPK))
	pg, err5 := GetParentsGuide(services.GetDocumentFromURL(urlPG))

	if err1 != nil {
		log.Fatalln("nil")
	}
	if err2 != nil {
		log.Fatalln("nil")
	}
	if err3 != nil {
		log.Fatalln("nil")
	}
	if err4 != nil {
		log.Fatalln("nil")
	}
	if err5 != nil {
		log.Fatalln("nil")
	}

	movie.Info = *mi

	movie.TL = *tl
	movie.PS = *ps
	movie.PK = *pk
	movie.PG = *pg

	return movie, nil
}

func GetSearchMovies(doc *goquery.Document) (*models.SearchResponse, error) {
	searches := ParseSearchMovies(doc)
	return searches, nil
}

func GetMovieInfo(doc *goquery.Document) (*models.MovieInfo, error) {
	info := ParseMovieInfo(doc)
	return info, nil
}

func GetTagline(doc *goquery.Document) (*models.Tagline, error) {
	tags := ParseTagline(doc)
	return tags, nil
}

func GetPlotKeywords(doc *goquery.Document) (*models.PlotKeywords, error) {
	keywords := ParsePlotKeywords(doc)
	return keywords, nil
}

func GetPlotSummary(doc *goquery.Document) (*models.PlotSummary, error) {
	summary := ParsePlotSummary(doc)
	return summary, nil
}

func GetParentsGuide(doc *goquery.Document) (*models.ParentsGuide, error) {
	rates := ParseParentsGuide(doc)
	return rates, nil
}
