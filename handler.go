package urlshort

import (
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	//return a redirect if there's a url of the PATH being served
	return func(w http.ResponseWriter, req *http.Request) {
		if path, ok := pathsToUrls[req.URL.Path]; ok {
			http.Redirect(w, req, path, http.StatusFound)
		}
		fallback.ServeHTTP(w, req)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {

	// define the object to unmarshal
	var out []struct {
		Path string `yaml:"path"`
		url  string `yaml:"url"`
	}

	if err := yaml.Unmarshal(yml, &out); err != nil {
		log.Fatalf("cannot unmarshat data: %v", err)
	}

	return func(w http.ResponseWriter, req *http.Request) {
		for _, path := range out {
			if path.Path == req.URL.Path {
				http.Redirect(w, req, path.url, http.StatusFound)
				return
			}
		}
		fallback.ServeHTTP(w, req)
	}, nil
}
