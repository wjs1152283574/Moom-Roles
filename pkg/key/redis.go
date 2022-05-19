package key

import "strconv"

func TokenKey(userID uint64) string {
	return "access_token:" + strconv.Itoa(int(userID))
}
