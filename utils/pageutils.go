package utils

func GetPageParam(pageNumInt, limitInt int64) (page int64, start int64, offset int64, err error) {

	if limitInt <= 0 || limitInt > 100 {
		limitInt = 20
	}
	if pageNumInt <= 0 {
		pageNumInt = 1
	}
	if pageNumInt != 0 {
		pageNumInt--
	}
	pageCount := limitInt * pageNumInt
	return pageNumInt + 1, pageCount, limitInt, nil
}
