package account


func createAccount(c echo.Context) error {
	// id・作成日は自動
	username = c.FormValue("username")
	password = c.FormValue("password")

	if err != nil {
		return error
	}

	return c.HTML(`<p>Hello!`, username, `</p>`)
}

