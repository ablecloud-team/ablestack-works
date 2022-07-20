package main

import "net/http"

func userPasswordChange(userName string, newPassword string) {
	databaseReturnValue := updateUserPassword(userName, newPassword)
	if databaseReturnValue.status == http.StatusOK {
		patchUserPassword(userName, newPassword)
	}
}
