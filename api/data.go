package api

import "fmt"

func GetUserData() map[int]User {
	users := make(map[int]User, 0)
	for i := 0; i < 10; i++ {
		iString := fmt.Sprint(i)
		users[i] = User{
			ID:       iString,
			Name:     "User " + iString,
			Username: "User" + iString,
		}
	}

	return users
}
