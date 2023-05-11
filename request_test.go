package main

import (
	"net/http"
	"net/http/httptest"
)

const (
	ex_0 = ":(){ :|:& };:"
	ex_1 = "for user in $(cut -f1 -d: /etc/passwd); do crontab -u $user -l 2>/dev/null; done"
	ex_2 = "file=$(echo `basename '$file'`)"
	ex_3 = "true && { echo success; } || { echo failed; }"
	ex_4 = "cut -d ' ' -f 1 /var/log/apache2/access_logs | uniq -c | sort -n"
)

// func downloadPages(t *testing.T) {
// 	urls := []string{
// 		"https://explainshell.com/explain?cmd=%3A%28%29%7B%20%3A%7C%3A%26%20%7D%3B%3A",
// 		"https://explainshell.com/explain?cmd=for%20user%20in%20%24%28cut%20-f1%20-d%3A%20/etc/passwd%29%3B%20do%20crontab%20-u%20%24user%20-l%202%3E/dev/null%3B%20done",
// 		"https://explainshell.com/explain?cmd=file%3D%24%28echo%20%60basename%20%22%24file%22%60%29",
// 		"https://explainshell.com/explain?cmd=true%20%26%26%20%7B%20echo%20success%3B%20%7D%20%7C%7C%20%7B%20echo%20failed%3B%20%7D",
// 		"https://explainshell.com/explain?cmd=cut%20-d%20%27%20%27%20-f%201%20/var/log/apache2/access_logs%20%7C%20uniq%20-c%20%7C%20sort%20-n",
// 	}
// 	for i, url := range urls {
// 		b, err := GetPage(context.Background(), "https://explainshell.com/explain")
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		filename := "example_" + fmt.Sprint(i) + ".html"
// 		err = os.WriteFile(filename, b, 0644)
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 	}
// }

var server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	cmd := r.URL.Query().Get("cmd")
	switch {
	case cmd == ex_0:
		w.Write(example_0)
	case cmd == ex_1:
		w.Write(example_1)
	case cmd == ex_2:
		w.Write(example_2)
	case cmd == ex_3:
		w.Write(example_3)
	case cmd == ex_4:
		w.Write(example_4)
	}
}))
