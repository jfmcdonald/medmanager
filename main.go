package main

import (
	"cogentcore.org/core/colors"
	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/units"
)

func main() {

	type user struct {
		Username string
		Password string
	}

	type Patient struct {
		Name    string
		History string
	}
	patients := []Patient{
		{
			Name:    "Bob",
			History: "This is Bobs patient history",
		},
		{
			Name:    "Jane",
			History: "This is Jane's patient history",
		},
		{
			Name:    "Joe",
			History: "This is Joe's patient history",
		},
	}
	b := core.NewBody("MedManager")
	u := &user{}
	pg := core.NewPages(b)
	pg.AddPage("sign-in", func(pg *core.Pages) {
		logininfo := core.NewForm(pg).SetStruct(u)
		core.NewButton(pg).SetText("Sign in").OnClick(func(e events.Event) {
			if u.Username == "testing" && u.Password == "password" {
				pg.Open("home")
			} else {
				core.MessageSnackbar(b, "login incorrect")
				u.Username = ""
				u.Password = ""
				logininfo.Update()
			}
		})
	})
	pg.AddPage("home", func(pg *core.Pages) {
		core.NewText(pg).SetText("Welcome, " + u.Username + "!").SetType(core.TextHeadlineSmall)

		outerFr := core.NewFrame(pg)
		outerFr.Styler(func(s *styles.Style) {
			s.Direction = styles.Row

		})

		listFr := core.NewFrame(outerFr)
		listFr.Styler(func(s *styles.Style) {
			s.Direction = styles.Column
			s.Border.Width.Set(units.Dp(4))
			s.Border.Color.Set(colors.Scheme.Outline)
		})

		ptRecord := core.NewText(outerFr)
		ptRecord.Styler(func(s *styles.Style) {
			s.Border.Width.Set(units.Dp(4))
			s.Border.Color.Set(colors.Scheme.Outline)
			s.Min.X.Set(70, units.UnitPw)
			s.Min.Y.Set(80, units.UnitPh)
		})

		for i := 0; i < len(patients); i++ {
			core.NewButton(listFr).SetText(patients[i].Name).OnClick(func(e events.Event) {

				ptRecord.SetText(patients[i].History)
			})
		}

		core.NewButton(outerFr).SetText("Sign out").OnClick(func(e events.Event) {
			*u = user{}
			pg.Open("sign-in")
		})
	})
	b.RunMainWindow()
}
