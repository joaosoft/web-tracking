# Web Tracking
[![Build Status](https://travis-ci.org/joaosoft/web-tracking.svg?branch=master)](https://travis-ci.org/joaosoft/web-tracking) | [![codecov](https://codecov.io/gh/joaosoft/web-tracking/branch/master/graph/badge.svg)](https://codecov.io/gh/joaosoft/web-tracking) | [![Go Report Card](https://goreportcard.com/badge/github.com/joaosoft/web-tracking)](https://goreportcard.com/report/github.com/joaosoft/web-tracking) | [![GoDoc](https://godoc.org/github.com/joaosoft/web-tracking?status.svg)](https://godoc.org/github.com/joaosoft/web-tracking)

A simple web tracking that allows you to send the url parameter "tracking=true" in any route with the tracking parameters on the url parameters or the body and send to the tracking service.

## Support for 
> Tracking [github/joaosoft/tracking](https://github.com/joaosoft/tracking).

###### If i miss something or you have something interesting, please be part of this project. Let me know! My contact is at the end.

## Dependecy Management 
>### Dep

Project dependencies are managed using Dep. Read more about [Dep](https://github.com/golang/dep).
* Install dependencies: `dep ensure`
* Update dependencies: `dep ensure -update`

## Configuration
```
{
  "web-tracking": {
    "host": "localhost:8002",
    "tracking_host": "localhost:8001/api/v1/tracking/event"
  }
}
```

## How it works

> Tracking location by street

Method: ```POST``` 
Route: ```http://localhost:8002/api/v1/dummy?tracking=true&category=bananas```
Body:
```
{
	"tracking": {
		"action": "action",
		"label": "label",
		"value": 1,
		"viewer": "joao",
		"viewed": "jessica",
		"street": "rua particular de monsanto",
		"meta_data": {
			"teste_1": "teste",
			"teste_2": 1,
			"teste_3": 1.1
		}
	}
}
```

> Tracking location by latitude, longitude

Method: ```POST``` 
Route: ```http://localhost:8002/api/v1/dummy?tracking=true&category=bananas```
Body:
```
{
	"tracking": {
		"category": "category",
		"action": "action",
		"label": "label",
		"value": 1,
		"viewer": "joao",
        "viewed": "jessica",
		"latitude": 41.1718238,
		"longitude": -8.6186277,
		"meta_data": {
			"teste_1": "teste",
			"teste_2": 1,
			"teste_3": 1.1
		}
	}
}
```

>### Go
```
go get github.com/joaosoft/web-tracking
```

## Usage 
This examples are available in the project at [web-tracking/examples](https://github.com/joaosoft/web-tracking/tree/master/examples)

```go
func main() {
	m, err := web_tracking.NewWebTracking()
	if err != nil {
		panic(err)
	}

	if err := m.Start(); err != nil {
		panic(err)
	}
}
```

## Known issues

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com
