package view

import (
	"context"
	"fmt"
	"github.com/rivo/tview"
	log "github.com/sirupsen/logrus"
	"github.com/yurchenkosv/credential_storage/internal/model"
	"github.com/yurchenkosv/credential_storage/internal/service"
)

type TUI struct {
	ctx   context.Context
	svc   *service.ClientCredentialsService
	menu  *tview.List
	form  *tview.Form
	items *tview.List
	pages *tview.Pages
	app   *tview.Application
}

func NewTUI(svc *service.ClientCredentialsService, ctx context.Context) *TUI {
	return &TUI{
		svc:   svc,
		ctx:   ctx,
		menu:  tview.NewList(),
		form:  tview.NewForm(),
		items: tview.NewList(),
		pages: tview.NewPages(),
		app:   tview.NewApplication(),
	}
}

func (t *TUI) RunApp() error {
	t.pages.
		AddPage("data_form", t.form, true, false).
		AddPage("menu", t.menu, true, true).
		AddPage("cred_list", t.items, true, false)
	t.app.SetRoot(t.pages, true)
	t.drawMainMenu()
	return t.app.Run()
}

func (t *TUI) drawMainMenu() {
	t.menu.
		AddItem("view data", "", 'v', func() {
			t.items.Clear()
			creds, err := t.svc.GetData()
			if err != nil {
				log.Error(err)
			}
			drawDataList(t.items, t.pages, creds)
			t.pages.SwitchToPage("cred_list")
		}).
		AddItem("add credentials data", "", 'c', func() {
			t.form.Clear(true)
			t.drawCredentialsForm(t.form, t.pages)
			t.pages.SwitchToPage("data_form")
		}).
		AddItem("add bank data", "", 'b', func() {
			t.form.Clear(true)
			t.drawBankForm(t.form, t.pages)
			t.pages.SwitchToPage("data_form")
		}).
		AddItem("add binary data", "", 'd', func() {
			t.form.Clear(true)
			t.drawBinaryForm(t.form, t.pages)
			t.pages.SwitchToPage("data_form")
		}).
		AddItem("add text data", "", 't', func() {
			t.form.Clear(true)
			t.drawTextForm(t.form, t.pages)
			t.pages.SwitchToPage("data_form")
		}).
		AddItem("quit", "", 'q', func() {
			t.app.Stop()
		})
}

func drawDataList(itemList *tview.List, pages *tview.Pages, credentials []model.Credentials) {
	for _, cred := range credentials {
		if cred.BankingCardData != nil {
			list := tview.NewList()
			bd := cred.BankingCardData
			drawBankInfo(list, bd, pages)
			drawMetadataInfo(list, cred.Metadata)
			itemList.AddItem(cred.Name, "", 0, func() {
				pages.AddPage(cred.Name, list, true, false)
				pages.SwitchToPage(cred.Name)
			})
			list.AddItem("back", "", 'b', func() {
				pages.SwitchToPage("cred_list")
			})
		}
		if cred.CredentialsData != nil {
			cd := cred.CredentialsData
			list := tview.NewList()
			drawCredInfo(list, cd)
			drawMetadataInfo(list, cred.Metadata)
			itemList.AddItem(cred.Name, "", 0, func() {
				pages.AddPage(cred.Name, list, true, false)
				pages.SwitchToPage(cred.Name)
			})
			list.AddItem("back", "", 'b', func() {
				pages.SwitchToPage("cred_list")
			})
		}
		if cred.BinaryData != nil {
			bd := cred.BinaryData
			list := tview.NewList()
			drawBinInfo(list, bd)
			drawMetadataInfo(list, cred.Metadata)
			itemList.AddItem(cred.Name, "", 0, func() {
				pages.AddPage(cred.Name, list, true, false)
				pages.SwitchToPage(cred.Name)
			})
			list.AddItem("back", "", 'b', func() {
				pages.SwitchToPage("cred_list")
			})
		}
		if cred.TextData != nil {
			td := cred.TextData
			list := tview.NewList()
			drawTextInfo(list, td)
			drawMetadataInfo(list, cred.Metadata)
			itemList.AddItem(cred.Name, "", 0, func() {
				pages.AddPage(cred.Name, list, true, false)
				pages.SwitchToPage(cred.Name)
			})
			list.AddItem("back", "", 'b', func() {
				pages.SwitchToPage("cred_list")
			})
		}
	}
	itemList.AddItem("back", "", 'b', func() {
		pages.SwitchToPage("menu")
	})
}

func (t *TUI) drawCredentialsForm(form *tview.Form, pages *tview.Pages) {
	credData := &model.CredentialsData{}
	form.
		AddInputField("Name", "", 20, nil, func(name string) {
			credData.Name = name
		}).
		AddInputField("Login", "", 20, nil, func(login string) {
			credData.Login = login
		}).
		AddPasswordField("Password", "", 20, '*', func(password string) {
			credData.Password = password
		}).
		AddButton("Save", func() {
			err := t.svc.SendCredentials(*credData)
			if err != nil {
				log.Error(err)
			}
			pages.SwitchToPage("menu")
		}).
		AddButton("Cancel", func() {
			pages.SwitchToPage("menu")
		})
}

func (t *TUI) drawBankForm(form *tview.Form, pages *tview.Pages) {
	bankData := &model.BankingCardData{}
	form.
		AddInputField("Name", "", 20, nil, func(name string) {
			bankData.Name = name
		}).
		AddInputField("Card Number", "", 20, nil, func(number string) {
			bankData.Number = number
		}).
		AddInputField("CVV", "", 20, nil, func(cvv string) {
			bankData.CVV = cvv
		}).
		AddInputField("Cardholder", "", 20, nil, func(cardholder string) {
			bankData.CardholderName = cardholder
		}).
		AddInputField("Valid until", "", 20, nil, func(valid string) {
			bankData.ValidUntil = valid
		}).
		AddButton("Save", func() {
			err := t.svc.SendBankCard(*bankData)
			if err != nil {
				log.Error(err)
			}
			pages.SwitchToPage("menu")
		}).
		AddButton("Cancel", func() {
			pages.SwitchToPage("menu")
		})
}

func (t *TUI) drawTextForm(form *tview.Form, pages *tview.Pages) {
	textData := &model.TextData{}
	form.
		AddInputField("Name", "", 20, nil, func(name string) {
			textData.Name = name
		}).
		AddTextArea("Text", "", 20, 40, 800, func(data string) {
			textData.Data = data
		}).
		AddButton("Save", func() {
			err := t.svc.SendText(*textData)
			if err != nil {
				log.Error(err)
			}
			pages.SwitchToPage("menu")
		}).
		AddButton("Cancel", func() {
			pages.SwitchToPage("menu")
		})

}

func (t *TUI) drawBinaryForm(form *tview.Form, pages *tview.Pages) {
	binData := &model.BinaryData{}
	form.
		AddInputField("Name", "", 20, nil, func(name string) {
			binData.Name = name
		}).
		AddInputField("File location", "", 20, nil, func(data string) {
			binData.Data = []byte(data)
		}).
		AddButton("Save", func() {
			err := t.svc.SendBinary(*binData)
			if err != nil {
				log.Error(err)
			}
			pages.SwitchToPage("menu")
		}).
		AddButton("Cancel", func() {
			pages.SwitchToPage("menu")
		})
}

func drawCredInfo(credList *tview.List, data *model.CredentialsData) {
	credList.
		AddItem(fmt.Sprint("Login: ", data.Login), "", 0, nil).
		AddItem(fmt.Sprint("Password: ", data.Password), "", 0, nil)
}

func drawBankInfo(credList *tview.List, data *model.BankingCardData, pages *tview.Pages) {
	credList.
		AddItem(fmt.Sprint("Cardholder: ", data.CardholderName), "", 0, nil).
		AddItem(fmt.Sprint("Number: ", data.Number), "", 0, nil).
		AddItem(fmt.Sprint("CVV: ", data.CVV), "", 0, nil).
		AddItem(fmt.Sprint("Valid Until: ", data.ValidUntil), "", 0, nil)
}

func drawTextInfo(credList *tview.List, data *model.TextData) {
	credList.
		AddItem(fmt.Sprint("Text: ", data.Data), "", 0, nil)
}

func drawBinInfo(credList *tview.List, data *model.BinaryData) {
	credList.
		AddItem(fmt.Sprint("Data: ", data.Data), "", 0, nil)
}

func drawMetadataInfo(credList *tview.List, data []model.Metadata) {
	credList.AddItem("Metadata: ", "", 0, nil)
	for _, meta := range data {
		credList.AddItem(meta.Value, "", 0, nil)
	}
}
