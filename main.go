package main

import (
	"time"

	"cogentcore.org/core/colors"
	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/units"
)

const mongoURI = "mongodb://localhost:27017"

func main() {
	//ctime := time.Now()

	patients, _ := getPatients()

	b := core.NewBody("MedManager")
	u := &user{}
	p := &Patient{}
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

		core.NewButton(buttonFr).SetText("New Patient contact").OnClick(func(e events.Event) {
			pg.Open("sign-in")
		})
		core.NewButton(buttonFr).SetText("New Patient vitals").OnClick(func(e events.Event) {
			pg.Open("sign-in")
		})

		npbt := core.NewButton(buttonFr).SetText("New Patient")
		npbt.OnClick(func(e events.Event) {
			d := core.NewBody("New Patient")
			core.NewForm(d).SetStruct(p)
			d.AddBottomBar(func(bar *core.Frame) {
				d.AddCancel(bar)
				d.AddOK(bar).OnClick(func(e events.Event) {
					core.MessageSnackbar(npbt, "Patient "+p.Name+" added")
					insertPatient(p)
				})
			})
			d.RunDialog(npbt)
		})

		core.NewButton(buttonFr).SetText("Trasnfer Patient").OnClick(func(e events.Event) {
			pg.Open("sign-in")
		})
		listFr := core.NewFrame(bodyFr)
		listFr.Styler(func(s *styles.Style) {
			s.Direction = styles.Column
			s.Border.Width.Set(units.Dp(4))
			s.Border.Color.Set(colors.Scheme.Outline)
		})

		ptRecord := core.NewFrame(bodyFr)
		ptRecord.Styler(func(s *styles.Style) {
			s.Border.Width.Set(units.Dp(4))
			s.Border.Color.Set(colors.Scheme.Outline)
			s.Min.X.Set(700, units.UnitPx)
			s.Min.Y.Set(800, units.UnitPx)
		})

		historyFrames := []*core.Frame{}
		for i := 0; i < len(patients); i++ {
			p = &patients[i]
			core.NewButton(listFr).SetText(patients[i].Name).OnClick(func(e events.Event) {

				for h := 0; h < len(patients[i].Histories); h++ {
					historyFrames[h] := core.NewFrame(ptRecord)
					historyFrames[h].Styler(func(s *styles.Style) {
						s.Direction = styles.Column
						s.Columns = 4
						s.Border.Width.Set(units.Dp(4))
						s.Border.Color.Set(colors.Scheme.Outline)
					})
					core.NewText(historyFrames[h]).SetText("entry date")
					core.NewText(historyFrames[h]).SetText(patients[i].Histories[h].Date.Format(time.RFC850))
					core.NewText(historyFrames[h]).SetText("Recorded by")
					core.NewText(historyFrames[h]).SetText(patients[i].Histories[h].Recorder)
					historylabel := core.NewText(historyFrames[h]).SetText("History entry")
					historylabel.Styler(func(s *styles.Style) {
						s.Grow.Set(1, 1)
					})
					core.NewText(historyFrames[h].SetText(patients[i].Histories[h].Body))
				}
				//ptRecord.SetText(patients[i].Histories[0].Body)
				ptRecord.UpdateChange()
			})
		}

	})
	b.RunMainWindow()
}
