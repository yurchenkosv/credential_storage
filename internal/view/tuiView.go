package view

import (
	"context"
	"github.com/rivo/tview"
	"github.com/yurchenkosv/credential_storage/internal/service"
)

type TuiView struct {
	app                *tview.Application
	credentialsService *service.ClientCredentialsService
	authService        *service.ClientAuthService
	pages              *tview.Pages
}

func NewTuiView(authservice *service.ClientAuthService) *TuiView {
	return &TuiView{
		authService: authservice,
		app:         tview.NewApplication(),
		pages:       tview.NewPages(),
	}
}

func (v *TuiView) GetApp() *tview.Application {
	return v.app
}

func (v *TuiView) DrawLoginForm() *tview.Form {
	var (
		login    string
		password string
		form     *tview.Form
		modal    *tview.Modal
	)

	modal = tview.NewModal().
		AddButtons([]string{"ok"}).
		SetText("wrong credentials").
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			v.pages.SwitchToPage("login form")
		})

	form = tview.NewForm().
		AddInputField("login", "", 20, nil, func(text string) {
			login = text
		}).
		AddPasswordField("password", "", 20, '*', func(text string) {
			password = text
		}).
		AddButton("login", func() {
			_, err := v.authService.Authenticate(context.Background(), login, password)
			if err != nil {
				v.app.SetFocus(modal).SetRoot(modal, true)
			}
		}).
		AddButton("cancel", func() {
			v.app.Stop()
		})
	form.SetTitle("Enter credentials to login").SetTitleAlign(tview.AlignLeft)
	return form
}

func (v *TuiView) drawAllDataPage() tview.Primitive {
	list := tview.NewList()
	list.AddItem("test", "test help", '-', nil)

	return list
}

func (v *TuiView) DrawMainApp() *tview.Pages {
	v.pages.AddPage("login form", v.DrawLoginForm(), true, true)
	v.pages.AddPage("all data", v.drawAllDataPage(), true, true)
	//pages.AddPage("add new")
	//pages.AddPage("new credentials")
	//pages.AddPage("new banking")
	//pages.AddPage("new text")
	//pages.AddPage("new binary")
	//pages.SwitchToPage("add data")
	v.pages.SwitchToPage("login form")
	return v.pages
}
