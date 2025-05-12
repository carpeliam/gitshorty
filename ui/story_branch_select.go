package ui

import (
	"fmt"
	"slices"

	"github.com/carpeliam/gitshorty/gitshorty"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type selection struct {
	name     string
	branch   gitshorty.Branch
	children []*selection
}

type Model struct {
	cursor     int
	input      *huh.MultiSelect[*selection]
	form       *huh.Form
	selections []*selection
	acc        huh.Accessor[[]*selection]
	Branches   []gitshorty.Branch
}

func NewModel(stories []gitshorty.Story) Model {
	selections := []*selection{}
	for _, story := range stories {
		if len(story.Branches) == 0 {
			continue
		}
		sel := selection{name: fmt.Sprintf("[sc-%d] %s", story.Id, story.Name), children: make([]*selection, len(story.Branches))}
		selections = append(selections, &sel)
		for i, branch := range story.Branches {
			sel.children[i] = &selection{name: fmt.Sprintf(" %s", branch.Name), branch: branch}
			selections = append(selections, sel.children[i])
		}

	}
	acc := &huh.EmbeddedAccessor[[]*selection]{}
	input := huh.NewMultiSelect[*selection]().Accessor(acc).Options(selectionsToOptions(selections)...).Filterable(false /* 😭 */).Title("abc")
	form := huh.NewForm(huh.NewGroup(input))
	form.SubmitCmd = tea.Quit
	form.CancelCmd = tea.Quit
	return Model{
		cursor:     0,
		input:      input,
		form:       form,
		selections: selections,
		acc:        acc,
		Branches:   []gitshorty.Branch{},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			// 	return m, tea.Quit
			// FIXME: I can't just quit because I might be filtering, but there's no way to know if I'm filtering?
		case tea.KeyUp:
			// TODO: are these values right? what about page up/down or filter, does that affect cursor?
			m.cursor = max(0, m.cursor-1)
		case tea.KeyDown:
			m.cursor = min(m.cursor+1, len(m.selections)-1)
		case tea.KeySpace:
			selectedOption := m.selections[m.cursor]
			selections := m.acc.Get()
			isParentSelected := !slices.Contains(selections, selectedOption)
			if isParentSelected {
				for _, child := range selectedOption.children {
					if !slices.Contains(selections, child) {
						selections = append(selections, child)
					}
				}
			} else {
				selections = slices.DeleteFunc(selections, func(s *selection) bool {
					return slices.Contains(selectedOption.children, s)
				})
			}
			m.acc.Set(selections)
			m.input.Options(selectionsToOptions(m.selections)...)
		case tea.KeyEnter:
			branches := []gitshorty.Branch{}
			for _, selection := range m.acc.Get() {
				if (selection.branch != gitshorty.Branch{}) {
					branches = append(branches, selection.branch)
				}
			}
			m.Branches = branches
		case tea.KeyRunes:
			switch string(msg.Runes) {
			case "q":
				return m, tea.Quit
			}
		}
	}
	return update(m, msg)
}

func update(m Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	_, cmd := m.form.Update(msg)
	// form := model.(*huh.Form)
	return m, cmd
}

func (m Model) View() string {
	return m.form.View()
}

func selectionsToOptions(selections []*selection) []huh.Option[*selection] {
	options := make([]huh.Option[*selection], len(selections))
	for i, sel := range selections {
		options[i] = huh.NewOption(sel.name, sel)
	}
	return options
}
