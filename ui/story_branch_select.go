package ui

import (
	"fmt"
	"log/slog"
	"slices"

	// "log/slog"

	"github.com/carpeliam/gitshorty/gitshorty"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type selection struct {
	// isSelected bool
	name       string
	// storyId int64
	children []*selection
}

type Model struct {
	// choices  []string           // items on the to-do list
	stories []gitshorty.Story
	cursor  int // which to-do list item our cursor is pointing at
	// selected map[int]struct{}   // which to-do items are selected
	input      *huh.MultiSelect[*selection]
	form       *huh.Form
	selections []*selection
	acc huh.Accessor[[]*selection]
}

func NewModel(stories []gitshorty.Story) Model {
	// options := []string{}
	// for _, story := range stories {
	// 	options = append(options, fmt.Sprintf("[sc-%d] %s", story.Id, story.Name))
	// 	for _, branch := range story.Branches {
	// 		options = append(options, fmt.Sprintf(" %s", branch.Name))
	// 	}
	// }
	// options := []selection{}
	// opts := []huh.Option[*selection]{}
	selections := []*selection{}
	for _, story := range stories {
		sel := selection{name: fmt.Sprintf("[sc-%d] %s", story.Id, story.Name), children: make([]*selection, len(story.Branches))}
		selections = append(selections, &sel)
		// opts = append(opts, huh.NewOption(sel.name, &sel))
		for i, branch := range story.Branches {
			// child :=
			sel.children[i] = &selection{name: fmt.Sprintf(" %s", branch.Name)}
			selections = append(selections, sel.children[i])
			// opts = append(opts, huh.NewOption(sel.children[i].name, sel.children[i]))
		}

	}
	// input := huh.NewMultiSelect[*selection]().Options(opts...).Title("abc")
	acc := &huh.EmbeddedAccessor[[]*selection]{}
	input := huh.NewMultiSelect[*selection]().Accessor(acc).OptionsFunc(func() []huh.Option[*selection] {
		return selectionsToOptions(selections)
	}, &selections).Title("abc")
	form := huh.NewForm(huh.NewGroup(input)).WithShowHelp(false)
	return Model{
		stories:    stories,
		cursor:     0,
		input:      input,
		form:       form,
		selections: selections,
		acc: acc,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func isSelected(sel *selection, acc huh.Accessor[[]*selection]) bool {
	for _, s := range acc.Get() {
		if s == sel {
			return true
		}
	}
	return false
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// fmt.Printf("%v", msg)
	// slog.Info("msg", slog.String("msg", fmt.Sprintf("%#v", msg)))

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc, tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyUp:
			m.cursor = max(0, m.cursor-1)
		case tea.KeyDown:
			m.cursor = min(m.cursor+1, len(m.selections)-1)
		case tea.KeySpace:
			// children = *m.selections[m.cursor]
			// sels := *m.selections
			// sel := sels[m.cursor]
			selectedOption := m.selections[m.cursor]
			isParentSelected := !isSelected(selectedOption, m.acc)
			// isParentSelected := !selectedOption.isSelected
			// selectedOption.name = "abc dif"
			if len(selectedOption.children) > 0 {
				selections := m.acc.Get()
				// selectedOption.children
				slices.Collect(func(yield func(*selection) bool) {
					// give me all the selections, plus children if isParentSelected; or minus if not
					for _, s := range selections {
						if isParentSelected {
							// slices.Contains(selectedOption.children, s)
						}
					}
				})
				// for _, child := range selectedOption.children {
				// 	// child.isSelected = isParentSelected
				// 	x := m.acc.Get()
				// 	if isParentSelected {
				// 		m.acc.Set(append(x, child))
				// 	} else {

				// 	}
				// 	// slog.Info("I would like to see the baby", "child", child.name, "selected", child.isSelected, "eq", child == m.selections[1])
				// }
			}
			// for _, s := range m.selections {
			// 	var checkdot string
			// 	if s.isSelected {
			// 		checkdot = "✓"
			// 	} else {
			// 		checkdot = "•"
			// 	}
				
			// 	slog.Info(fmt.Sprintf("%s %s", checkdot, s.name))
			// }
			// slog.Info("")
			for _, s := range m.acc.Get() {
				checkdot := "✓"
				
				slog.Info(fmt.Sprintf("%s %s", checkdot, s.name))
			}
			// select
			// nextOption := (*m.selections)[1]
			// nextOption.name = "different"
			// i := m.input.Value(&m.selections)
			// m.input = i
			m.input.Options(selectionsToOptions(m.selections)...)
			// m.input.OptionsFunc(func() []huh.Option[*selection] {
			// 	return selectionsToOptions(m.selections)
			// }, nil)

			// selectedOption.isSelected = isParentSelected
			// return m, nil
			// msg = nil
		case tea.KeyRunes:
			switch string(msg.Runes) {
			case "q":
				return m, tea.Quit
			}
		}
	}
	// return m, tea.Quit
	return update(m, msg)
}

func update(m Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	_, cmd := m.form.Update(msg)
	// return m.form.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	// options := []string{}
	// for _, story := range m.stories {
	// 	options = append(options, fmt.Sprintf("[sc-%d] %s", story.Id, story.Name))
	// 	for _, branch := range story.Branches {
	// 		options = append(options, fmt.Sprintf(" %s", branch.Name))
	// 	}
	// }

	// return huh.NewMultiSelect[string]().Options(huh.NewOptions(options...)...).Title("abc").View()
	return m.form.View()
}

func selectionsToOptions(selections []*selection) []huh.Option[*selection] {
	options := make([]huh.Option[*selection], len(selections))
	for i, sel := range selections {
		// slog.Info(fmt.Sprintf("sel %#v", sel))
		options[i] = huh.NewOption(sel.name, sel)
		// slog.Info(fmt.Sprintf("opt %#v", options[i]))
	}
	return options
}
