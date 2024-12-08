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
			s.Display = styles.Grid
			s.Columns = 2
		})
		buttonFr := core.NewFrame((outerFr))
		buttonFr.Styler(func(s *styles.Style) {
			s.Grow.Set(1, 1)
			s.Direction = styles.Row
		})
		core.NewButton(outerFr).SetText("Sign out").OnClick(func(e events.Event) {
			*u = user{}
			pg.Open("sign-in")
		})

		bodyFr := core.NewFrame(outerFr)
		bodyFr.Styler(func(s *styles.Style) {
			s.Direction = styles.Row
		})

		core.NewButton(buttonFr).SetText("New Patient").OnClick(func(e events.Event) {
			pg.Open("sign-in")
		})

		listFr := core.NewFrame(bodyFr)
		listFr.Styler(func(s *styles.Style) {
			s.Direction = styles.Column
			s.Border.Width.Set(units.Dp(4))
			s.Border.Color.Set(colors.Scheme.Outline)
		})

		ptRecord := core.NewText(bodyFr)
		ptRecord.Styler(func(s *styles.Style) {
			s.Border.Width.Set(units.Dp(4))
			s.Border.Color.Set(colors.Scheme.Outline)
			s.Min.X.Set(700, units.UnitPx)
			s.Min.Y.Set(800, units.UnitPx)
		})

		for i := 0; i < len(patients); i++ {
			core.NewButton(listFr).SetText(patients[i].Name).OnClick(func(e events.Event) {

				ptRecord.SetText(patients[i].History)
				ptRecord.UpdateChange()
			})
		}

	})
	b.RunMainWindow()
}
