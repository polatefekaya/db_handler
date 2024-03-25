package internals

import "DatabaseHandler/internals/log"

type Api struct {
	Base    string
	Url     string
	Headers *map[string]string
	Key     string
}

func CreateSportsApi(key, query string) *Api {
	sa := Api{
		Base: "https://v3.football.api-sports.io/",
		Url:  *Concat("https://v3.football.api-sports.io/", query),
		Headers: &map[string]string{
			"X-rapidapi-host": "v3.football.api-sports.io",
			"X-rapidapi-key":  key,
		},
		Key: key,
	}
	log.DEBUG(sa.Url, "base", sa.Base)
	return &sa
}
