package handler

import (
	"encoding/json"
	"net/http"
	"ware-data-api/user"
)

func WaresUpload(
	dw user.DeviceWareAdder,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ldCount, err := dw.GetCountLDDevice(ctx)
		if err != nil {
			handleError(w, err, http.StatusInternalServerError, true)
		}

		for i := 0; i < ldCount; i++ {
			dw.AddWareData(ctx, i, []user.WareData{})
		}

		response, _ := json.Marshal(map[string]string{"status": "ok"})
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(response)

	}
}
