package province

import (
	"strconv"

	"github.com/dhpsagala/loan_process_donnysagala/storage"
)

func SearchByNIK(nik int) (IProvinceEntity, error) {
	nikStr := strconv.Itoa(nik)
	if rows, err := storage.Db.Query("SELECT * FROM Province WHERE Code = ? LIMIT 1", nikStr[0:2]); err == nil {
		datas := []IProvinceEntity{}
		defer rows.Close()
		for rows.Next() {
			var ID, Code int
			var Name string
			var AllowLoan bool

			if err := rows.Scan(&ID, &Name, &Code, &AllowLoan); err != nil {
				return nil, err
			}
			var p IProvinceEntity = provinceEntity{ID: ID, Name: Name, Code: Code, AllowLoan: AllowLoan}
			datas = append(datas, p)
		}

		if len(datas) > 0 {
			return datas[0], nil
		}
	} else {
		return nil, err
	}
	return nil, nil
}
