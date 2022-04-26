package storage

func allPage(allr int, limit int) int {
	var (
		allp int
	)
	ost := allr % limit
	if ost != 0 {
		allp = allr/limit + 1
	} else {
		allp = allr / limit
	}
	return allp
}
