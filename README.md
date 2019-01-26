# Cloud2Cloud


## GoogleAppEngine

Basic sample to deploy dummy apps into Google Cloud App Engine. See [google Guides](doc/googleGuides.md) as on boarding material

Code Example:


```go
package app

import (
	"fmt"
	"net/http"
)

//  application starts serving.
func init() {
	// Handle all requests
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/test", testInterfaceHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "CesarCorp2Test GCP!")
}

func testInterfaceHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Validation Interface, are you expecting something funny or what?")
}

```

app.yaml sample:

```yaml
runtime: go
api_version: go1

handlers:
- url: /stylesheets
  static_dir: stylesheets

- url: /(.*\.(gif|png|jpg))$
  static_files: static/\1
  upload: static/.*\.(gif|png|jpg)$

- url: /.*
  script: _go_app

```


Deploy:

```shel
cd appGCP
gcloud projects create applauncher
gcloud config set project applauncher
gcloud app deploy
```

Debug:
```shel
gcloud app logs tail -s default
gcloud app browse
```


Running app [https://applauncher.appspot.com](https://applauncher.appspot.com)

Testing Interface [https://applauncher.appspot.com/test](https://applauncher.appspot.com/test)
