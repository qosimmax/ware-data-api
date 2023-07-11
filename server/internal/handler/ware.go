package handler

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
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

		wares := []user.WareData{
			{
				ItemCode:  101,
				Name:      "SHAFTOLI",
				Price:     30.50,
				Count:     0,
				GoodsType: 0,
			},
			{
				ItemCode:  102,
				Name:      "SAMSUNG A2",
				Price:     100.30,
				Count:     0,
				GoodsType: 1,
			},
		}

		for i := 0; i < ldCount; i++ {
			err := dw.AddWareData(ctx, i, wares)
			if err != nil {
				log.Error(err)
				//handleError(w, err, http.StatusInternalServerError, true)
			}
		}

		response, _ := json.Marshal(map[string]string{"status": "ok"})
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(response)

	}
}
