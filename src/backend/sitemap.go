package backend

import (
	"encoding/xml"
	"net/http"
	"strconv"
	"time"
	"zdamtosam/src/db"
)

const ZdamtosamHost = "https://zdamtosam.pl"

type URL struct {
	Loc        string  `xml:"loc"`
	LastMod    string  `xml:"lastmod"`
	ChangeFreq string  `xml:"changefreq"`
	Priority   float32 `xml:"priority"`
}

type Sitemap struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	Url     []URL    `xml:"url"`
}

func (s *Sitemap) Add(url URL) *Sitemap {
	url.Loc = ZdamtosamHost + url.Loc
	s.Url = append(s.Url, url)
	return s
}

func (h *Handler) Sitemap(w http.ResponseWriter, r *http.Request) {
	sitemap := &Sitemap{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
	}
	url := URL{Loc: "", Priority: 1, ChangeFreq: "monthly", LastMod: time.Now().String()[0:10]}
	sitemap.Add(url)
	url.Loc = "/privacy-policy"
	url.Priority = 0.1
	sitemap.Add(url)
	url.Loc = "/terms-of-service"
	url.Priority = 0.1
	sitemap.Add(url)
	levels := db.GetLevels(h.db)
	for _, level := range levels {
		levelIdString := strconv.Itoa(level.Id)
		url.Loc = "/level/" + levelIdString
		url.Priority = 0.6
		sitemap.Add(url)
		categories := db.GetCategoriesByLevel(h.db, levelIdString)
		for _, category := range categories {
			categoryIdString := strconv.Itoa(category.Id)
			url.Loc = "/level/" + levelIdString + "/category/" + categoryIdString
			url.Priority = 0.7
			sitemap.Add(url)
			subcategories := db.GetSubcategories(h.db, categoryIdString)
			for _, subcategory := range subcategories {
				subcategoryIdString := strconv.Itoa(subcategory.Id)
				url.Loc = "/level/" + levelIdString + "/category/" + categoryIdString + "/subcategory/" + subcategoryIdString
				url.Priority = 0.8
				sitemap.Add(url)
				exercises := db.GetExercisesBySubcategoryId(h.db, subcategoryIdString)
				for _, exercise := range exercises {
					exerciseIdString := strconv.Itoa(exercise.Id)
					url.Loc = "/level/" + levelIdString + "/category/" + categoryIdString + "/subcategory/" + subcategoryIdString + "/exercise/" + exerciseIdString
					url.Priority = 0.9
					url.LastMod = exercise.Date[0:10]
					sitemap.Add(url)
				}
			}
		}
	}

	w.Write([]byte(xml.Header))
	xml.NewEncoder(w).Encode(sitemap)
}
