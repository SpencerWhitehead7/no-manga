package playground

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ripped shamelessly from https://github.com/99designs/gqlgen/blob/master/graphql/playground/playground.go (and slightly adapted)

// https://github.com/graphql/graphql-playground and SRI version updates have to be done by hand
// <unpkgFilePathURL>?meta gives you the integrity
// and eventually migrated to graphiQL when they merge https://github.com/graphql/graphql-playground/issues/1143
var page = template.Must(template.New("graphiql").Parse(`<!DOCTYPE html>
<html>
<head>
	<meta charset=utf-8/>
	<meta name="viewport" content="user-scalable=no, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, minimal-ui">
	<link
		rel="shortcut icon"
		href="https://unpkg.com/graphql-playground-react@{{.version}}/build/static/media/logo.8c98d067.png"
		integrity="{{ .faviconSRI }}"
		crossorigin="anonymous"
	/>
	<link
		rel="stylesheet"
		href="https://unpkg.com/graphql-playground-react@{{.version}}/build/static/css/index.css"
		integrity="{{ .cssSRI }}"
		crossorigin="anonymous"
	/>
	<script
		type="text/javascript"
		src="https://unpkg.com/graphql-playground-react@{{.version}}/build/static/js/middleware.js"
		integrity="{{ .jsSRI }}"
		crossorigin="anonymous"
	></script>
	<title>{{.title}}</title>
</head>
<body>
<style type="text/css">
	html { font-family: "Open Sans", sans-serif; overflow: hidden; }
	body { margin: 0; background: #172a3a; }
</style>
<div id="root"/>
<script type="text/javascript">
	window.addEventListener('load', function (event) {
		const root = document.getElementById('root');
		root.classList.add('playgroundIn');
		const wsProto = location.protocol == 'https:' ? 'wss:' : 'ws:'
		GraphQLPlayground.init(root, {
			endpoint: location.protocol + '//' + location.host + '{{.endpoint}}',
			subscriptionsEndpoint: wsProto + '//' + location.host + '{{.endpoint }}',
      shareEnabled: true,
			settings: {
				'request.credentials': 'same-origin'
			}
		})
	})
</script>
</body>
</html>
`))

func handler(title string, endpoint string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		err := page.Execute(w, map[string]string{
			"title":      title,
			"endpoint":   endpoint,
			"version":    "1.7.28",
			"faviconSRI": "sha384-EZ+grpIFhvKgOSGq4SO5SzSZBUVNDTcOQdBYKeCDwXsrgpJzoM2Ml2kl5pbFLa5z",
			"cssSRI":     "sha384-xb+UHILNN4fV3NgQMTjXk0x9A80U0hmkraTFvucUYTILJymGT8E1Aq2278NSi5+3",
			"jsSRI":      "sha384-ardaO17esJ2ZxvY24V1OE6X4j+Z3WKgGMptrlDLmD+2w/JC3nbQ5ZfKGY2zfOPEE",
		})
		if err != nil {
			panic(err)
		}
	}
}

func NewHandler(title string, endpoint string) gin.HandlerFunc {
	h := handler(title, endpoint)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
