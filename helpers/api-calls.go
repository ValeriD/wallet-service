package helpers
import(
	"io/ioutil"
	"net/http"
	"encoding/json"
)
type ApiResponse struct{
}

func get(url string, res *ApiResponse){
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return
	}

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		return
	}
    data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return
	}
	json.Unmarshal(data, &res)
}