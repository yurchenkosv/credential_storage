package view

import (
	"bytes"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/rivo/tview"
	"github.com/yurchenkosv/credential_storage/internal/model"
	"github.com/yurchenkosv/credential_storage/internal/service"
	"os"
	"strings"
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
				drawModalError(err.Error(), t.pages)
			}
			t.drawDataList(creds)
			t.pages.SwitchToPage("cred_list")
		}).
		AddItem("add credentials data", "", 'c', func() {
			t.form.Clear(true)
			t.drawCredentialsForm()
			t.pages.SwitchToPage("data_form")
		}).
		AddItem("add bank data", "", 'b', func() {
			t.form.Clear(true)
			t.drawBankForm()
			t.pages.SwitchToPage("data_form")
		}).
		AddItem("add binary data", "", 'd', func() {
			t.form.Clear(true)
			t.drawBinaryForm()
			t.pages.SwitchToPage("data_form")
		}).
		AddItem("add text data", "", 't', func() {
			t.form.Clear(true)
			t.drawTextForm()
			t.pages.SwitchToPage("data_form")
		}).
		AddItem("quit", "", 'q', func() {
			t.app.Stop()
		})
}

func (t *TUI) drawDataList(credentials []model.Credentials) {
	for _, cred := range credentials {
		if cred.BankingCardData != nil {
			list := tview.NewList()
			bd := cred.BankingCardData
			t.drawBankInfo(list, bd)
			t.drawMetadataInfo(list, cred.Metadata)
			t.items.AddItem(cred.Name, "", 0, func() {
				t.pages.AddPage(cred.Name, list, true, false)
				t.pages.SwitchToPage(cred.Name)
			})
			list.AddItem("back", "", 'b', func() {
				t.pages.SwitchToPage("cred_list")
			})
		}
		if cred.CredentialsData != nil {
			cd := cred.CredentialsData
			list := tview.NewList()
			drawCredInfo(list, cd)
			t.drawMetadataInfo(list, cred.Metadata)
			t.items.AddItem(cred.Name, "", 0, func() {
				t.pages.AddPage(cred.Name, list, true, false)
				t.pages.SwitchToPage(cred.Name)
			})
			list.AddItem("back", "", 'b', func() {
				t.pages.SwitchToPage("cred_list")
			})
		}
		if cred.BinaryData != nil {
			bd := cred.BinaryData
			list := tview.NewList()
			t.drawBinInfo(list, bd)
			t.drawMetadataInfo(list, cred.Metadata)
			t.items.AddItem(cred.Name, "", 0, func() {
				t.pages.AddPage(cred.Name, list, true, false)
				t.pages.SwitchToPage(cred.Name)
			})
			list.AddItem("back", "", 'b', func() {
				t.pages.SwitchToPage("cred_list")
			})
		}
		if cred.TextData != nil {
			td := cred.TextData
			list := tview.NewList()
			t.drawTextInfo(list, td)
			t.drawMetadataInfo(list, cred.Metadata)
			t.items.AddItem(cred.Name, "", 0, func() {
				t.pages.AddPage(cred.Name, list, true, false)
				t.pages.SwitchToPage(cred.Name)
			})
			list.AddItem("back", "", 'b', func() {
				t.pages.SwitchToPage("cred_list")
			})
		}
	}
	t.items.AddItem("back", "", 'b', func() {
		t.pages.SwitchToPage("menu")
	})
}

func (t *TUI) drawCredentialsForm() {
	var meta string
	credData := &model.CredentialsData{}
	t.form.
		AddInputField("Name", "", 20, nil, func(name string) {
			credData.Name = name
		}).
		AddInputField("Login", "", 20, nil, func(login string) {
			credData.Login = login
		}).
		AddPasswordField("Password", "", 20, '*', func(password string) {
			credData.Password = password
		}).
		AddTextArea("Metadata", "", 20, 5, 400, func(text string) {
			meta = text
		}).
		AddButton("Save", func() {
			for _, item := range strings.Split(meta, "\n") {
				credData.Metadata = append(credData.Metadata, model.Metadata{Value: item})
			}
			err := t.svc.SendCredentials(*credData)
			t.pages.SwitchToPage("menu")
			if err != nil {
				drawModalError(err.Error(), t.pages)
			}
		}).
		AddButton("Cancel", func() {
			t.pages.SwitchToPage("menu")
		})
}

func (t *TUI) drawBankForm() {
	var meta string
	bankData := &model.BankingCardData{}
	t.form.
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
		AddTextArea("Metadata", "", 20, 5, 400, func(text string) {
			meta = text
		}).
		AddButton("Save", func() {
			for _, item := range strings.Split(meta, "\n") {
				bankData.Metadata = append(bankData.Metadata, model.Metadata{Value: item})
			}
			err := t.svc.SendBankCard(*bankData)
			t.pages.SwitchToPage("menu")
			if err != nil {
				drawModalError(err.Error(), t.pages)
			}
		}).
		AddButton("Cancel", func() {
			t.pages.SwitchToPage("menu")
		})
}

func (t *TUI) drawTextForm() {
	var meta string
	textData := &model.TextData{}
	t.form.
		AddInputField("Name", "", 20, nil, func(name string) {
			textData.Name = name
		}).
		AddTextArea("Text", "", 20, 40, 800, func(data string) {
			textData.Data = data
		}).
		AddButton("Save", func() {
			for _, item := range strings.Split(meta, "\n") {
				textData.Metadata = append(textData.Metadata, model.Metadata{Value: item})
			}
			err := t.svc.SendText(*textData)
			t.pages.SwitchToPage("menu")
			if err != nil {
				drawModalError(err.Error(), t.pages)
			}
		}).
		AddButton("Cancel", func() {
			t.pages.SwitchToPage("menu")
		})

}

func (t *TUI) drawBinaryForm() {
	var meta, fileLocation string
	binData := &model.BinaryData{}
	t.form.
		AddInputField("Name", "", 20, nil, func(name string) {
			binData.Name = name
		}).
		AddInputField("File location", "", 20, nil, func(dataLocation string) {
			fileLocation = dataLocation
		}).
		AddTextArea("Metadata", "", 20, 5, 400, func(text string) {
			meta = text
		}).
		AddButton("Save", func() {
			data, err := os.ReadFile(fileLocation)
			if err != nil {
				drawModalError(err.Error(), t.pages)
			}
			binData.Data = data
			for _, item := range strings.Split(meta, "\n") {
				binData.Metadata = append(binData.Metadata, model.Metadata{Value: item})
			}
			err = t.svc.SendBinary(*binData)
			t.pages.SwitchToPage("menu")
			if err != nil {
				drawModalError(err.Error(), t.pages)
			}
		}).
		AddButton("Cancel", func() {
			t.pages.SwitchToPage("menu")
		})
}

func drawCredInfo(credList *tview.List, data *model.CredentialsData) {
	credList.
		AddItem(fmt.Sprint("Login: ", data.Login), "", 0, nil).
		AddItem(fmt.Sprint("Password: ", data.Password), "", 0, nil)
}

func (t *TUI) drawBankInfo(credList *tview.List, data *model.BankingCardData) {
	credList.
		AddItem(fmt.Sprint("Cardholder: ", data.CardholderName), "", 0, nil).
		AddItem(fmt.Sprint("Number: ", data.Number), "", 0, nil).
		AddItem(fmt.Sprint("CVV: ", data.CVV), "", 0, nil).
		AddItem(fmt.Sprint("Valid Until: ", data.ValidUntil), "", 0, nil)
}

func (t *TUI) drawTextInfo(credList *tview.List, data *model.TextData) {
	credList.
		AddItem(fmt.Sprint("Text: ", data.Data), "", 0, nil)
}

func (t *TUI) drawBinInfo(credList *tview.List, data *model.BinaryData) {
	credList.
		AddItem("Download data", "", 0, func() {
			t.pages.SwitchToPage("menu")
			if data == nil {
				drawModalError("no data found", t.pages)
				return
			}
			reader := bytes.NewReader(data.Data)
			err := t.svc.SaveBinary(reader, uuid.New().String())
			t.pages.SwitchToPage("menu")
			if err != nil {
				drawModalError(err.Error(), t.pages)
				return
			}
		})
}

func (t *TUI) drawMetadataInfo(credList *tview.List, data []model.Metadata) {
	credList.AddItem("Metadata: ", "", 0, nil)
	for _, meta := range data {
		credList.AddItem(meta.Value, "", 0, nil)
	}
}

func drawModalError(errText string, pages *tview.Pages) {
	modal := func(p tview.Primitive, width, height int) tview.Primitive {
		return tview.NewFlex().
			AddItem(nil, 0, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(nil, 0, 1, false).
				AddItem(p, height, 1, true).
				AddItem(nil, 0, 1, false), width, 1, true).
			AddItem(nil, 0, 1, false)
	}
	form := tview.NewForm().
		AddTextView("", errText, 20, 5, false, false).
		AddButton("Ok", func() {
			pages.SwitchToPage("menu")
		})
	pages.AddPage("modal_err", modal(form, 40, 10), true, true)
}
