package users

import "fmt"

type UserId string

func NewUserId(value string) (UserId, error) {
	if value == "" {
		return "", fmt.Errorf("ValueError: UserIdが空です。")
	} else if len(value) != 28 {
		// Firebase AuthenticationのUIDに従う
		return "", fmt.Errorf("ValueError: UserIdが不正です。")
	}
	id := UserId(value)
	return id, nil
}
