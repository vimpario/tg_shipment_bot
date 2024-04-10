package model

func StartMenu() []Button {
	buttonContents := []Button{
		{
			Name: "Профиль",
			Data: "/profile",
		},
		{
			Name: "Мои посылки",
			Data: "package",
		},
		{
			Name: "Добавить трек-код",
			Data: "add",
		},
		{
			Name: "Информация",
			Data: "info",
		},
	}

	return buttonContents
}

func ProfileMenu() []Button {
	buttonContents := []Button{
		{
			Name: "Назад",
			Data: "back",
		},
	}
	return buttonContents
}
