package managers

import (
	"encoding/json"
	"github/Chidi-creator/go-medic-server/internal/utils"
	"net/http"
)


//sending the response to the front end
func JSONresponse(w http.ResponseWriter, status int, resp utils.ApiResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(resp)

}
