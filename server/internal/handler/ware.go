package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"ware-data-api/user"

	"github.com/xuri/excelize/v2"
)

func WaresUpload(
	dw user.DeviceWareAdder,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		f, _, err := r.FormFile("file")
		if err != nil {
			handleError(w, err, http.StatusInternalServerError, true)
			return
		}

		defer f.Close()

		fe, err := excelize.OpenReader(f)
		if err != nil {
			handleError(w, err, http.StatusInternalServerError, true)
			return

		}

		defer fe.Close()

		sheetName := fe.GetSheetName(0)
		// Get all the rows in the Sheet1.
		rows, err := fe.GetRows(sheetName)
		if err != nil {
			handleError(w, err, http.StatusInternalServerError, true)
			return
		}

		var (
			wares []user.WareData
		)

		for i := 1; i < len(rows); i++ {
			row := rows[i]

			price, _ := strconv.Atoi(row[17])
			if price > 999999 {
				price = 0
			}

			itemCode, _ := strconv.Atoi(row[38])
			if itemCode == 0 {
				continue
			}

			goodsType := 0
			if row[39] == "штучный" {
				goodsType = 1
			}

			wares = append(wares, user.WareData{
				ItemCode:  itemCode,
				Name:      row[21],
				Price:     price,
				Count:     0,
				GoodsType: goodsType,
			})

		}

		ldCount, err := dw.GetCountLDDevice(ctx)
		if err != nil {
			handleError(w, err, http.StatusInternalServerError, true)
		}

		for i := 0; i < ldCount; i++ {
			err := dw.AddWareData(ctx, i, wares)
			if err != nil {
				handleError(w, err, http.StatusInternalServerError, true)
				return
			}
		}

		response, _ := json.Marshal(map[string]string{"status": "ok"})
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(response)

	}
}
