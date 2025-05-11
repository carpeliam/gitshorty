package ui_test

import (
	// "fmt"

	tea "github.com/charmbracelet/bubbletea"
	// "github.com/charmbracelet/huh"
	"github.com/charmbracelet/x/ansi"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/carpeliam/gitshorty/gitshorty"
	. "github.com/carpeliam/gitshorty/ui"
)

var _ = Describe("StoryBranchSelect", func() {
	It("shows empty stories initially", func() {
		stories := []gitshorty.Story{
			{Id: 1, Name: "Story 1", Branches: []gitshorty.Branch{
				{Name: "b1-sc-1"},
				{Name: "b2-sc-1"},
			}},
			{Id: 2, Name: "Story 2"},
		}
		model := NewModel(stories)
		model.Update(tea.WindowSize())
		display := ansi.Strip(model.View())

		// fmt.Println(display)
		// fmt.Println(nothingSelected(stories))
		// Expect(display).To(Equal(nothingSelected(stories)))
		Expect(display).To(ContainSubstring("• [sc-1] Story 1"))
		Expect(display).To(ContainSubstring("•  b1-sc-1"))
		Expect(display).To(ContainSubstring("•  b2-sc-1"))
		Expect(display).To(ContainSubstring("• [sc-2] Story 2"))
	})
	It("quits with esc or q", func() {
		_, cmd := NewModel([]gitshorty.Story{}).Update(keys('q'))
		Expect(cmd).NotTo(BeNil())
		Expect(cmd()).To(Equal(tea.Quit()))
		// _, cmd = NewModel([]gitshorty.Story{}).Update(keys(rune(tea.KeyEsc)))
		// Expect(cmd).NotTo(BeNil())
		// Expect(cmd()).To(Equal(tea.Quit()))
	})
	It("selects all branches when selecting a story", func() {
		stories := []gitshorty.Story{
			{Id: 1, Name: "Story 1", Branches: []gitshorty.Branch{
				{Name: "b1-sc-1"},
				{Name: "b2-sc-1"},
			}},
			{Id: 2, Name: "Story 2"},
		}
		model := NewModel(stories)
		m, _ := model.Update(keys(rune(tea.KeySpace)))
		display := ansi.Strip(m.View())

		// Expect(display).To(Equal(storySelected(stories)))
		Expect(display).To(ContainSubstring("✓ [sc-1] Story 1"))
		Expect(display).To(ContainSubstring("✓  b1-sc-1"))
		Expect(display).To(ContainSubstring("✓  b2-sc-1"))
		Expect(display).To(ContainSubstring("• [sc-2] Story 2"))
	})

})

// func nothingSelected(stories []gitshorty.Story) string {
// 	options := []string{}
// 	for _, story := range stories {
// 		options = append(options, fmt.Sprintf("[sc-%d] %s", story.Id, story.Name))
// 		for _, branch := range story.Branches {
// 			options = append(options, fmt.Sprintf(" %s", branch.Name))
// 		}
// 	}
// 	return huh.NewForm(huh.NewGroup(huh.NewMultiSelect[string]().Options(
// 		huh.NewOptions(options...)...
// 	).Title("abc"))).View()
// }

// func storySelected(stories []gitshorty.Story) string {
// 	options := []huh.Option[string]{}
// 	for i, story := range stories {
// 		// options = append(options, fmt.Sprintf("[sc-%d] %s", story.Id, story.Name))
// 		options = append(options, huh.NewOption(fmt.Sprintf("[sc-%d] %s", story.Id, story.Name), "").Selected(i == 0))
// 		for _, branch := range story.Branches {
// 			options = append(options, huh.NewOption(fmt.Sprintf(" %s", branch.Name), "").Selected(i == 0))
// 		}
// 	}
// 	return huh.NewForm(huh.NewGroup(huh.NewMultiSelect[string]().Options(
// 		options...
// 	).Title("abc"))).WithShowHelp(false).View()
// }

func keys(runes ...rune) tea.KeyMsg {
	return tea.KeyMsg{
		Type:  tea.KeyRunes,
		Runes: runes,
	}
}
