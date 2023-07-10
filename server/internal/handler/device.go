package handler

import (
	"encoding/json"
	"net/http"
	"ware-data-api/user"
)

func GetDevices(
	dr user.DeviceFinder,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		devices, err := dr.FindActiveDevices(ctx, "192.168.0.1", "192.168.0.255")
		if err != nil {
			handleError(w, err, http.StatusInternalServerError, true)
		}

		response, _ := json.Marshal(devices)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(response)

	}
}
