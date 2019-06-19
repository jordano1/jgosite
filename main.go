// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Sample helloworld is a basic App Engine flexible app.
package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/search", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	io.WriteString(w, "<p>do my search: "+v+"<p>")
	io.WriteString(w, "this is a little example of using golang to get the value in the URL. to test put at the end of .com/, put ?q=[value] where value is what you type to show on the page. (ie:https://jgosite.appspot.com/?q=dog)")
}
