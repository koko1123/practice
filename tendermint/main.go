package tendermint

import "net/http"

func receiveMsg(w http.ResponseWriter, r *http.Request) {
	//
}

func main() {
	// create server in http and call handle on POST on port 8080
	http.HandleFunc("/send", receiveMsg)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
