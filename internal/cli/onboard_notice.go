package cli

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
)

var boxStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("62")).
	Padding(0, 2).
	MarginTop(1).
	MarginBottom(1)

// OnboardWarningForm provides a warning notice
func OnboardWarningForm(ctx context.Context) error {
	return OnboardConfirmForm(ctx, `
# RISK DISCLOSURE & DISCLAIMER
**PLEASE READ THESE WORDS CAREFULLY:**

Manboster is an AI agent able to chat and control your computers like OpenClaw and IronClaw and currently in MVP stage. By proceeding, you acknowledge:
1. WIP means this project is **Work in Progress**, and **it is expected to encounter bugs, crashes, and breaking changes.**
2. If you run 'manboster start', you open a daemon running in your computer. **The background process has persistent resource access to your computer.**
3. WASM sandboxing plugins is strong, but **3rd-party code still carries risks**.
4. **Hachimi scoring reduces decision fatigue, but cannot fully prevent advanced prompt injections or unsafe LLM behaviors.**
5. **Granting access enables data transmission to LLMs and allows device control. We are not liable for any issues arising from these interactions.**
6. This software is provided "AS IS" under Apache 2.0. **You are strictly prohibited from using this application for any criminal or illegal purposes. We disclaim all liability and responsibility for any unlawful activities conducted using this software.**
`, "Do you understand the risks and wish to proceed?", "I Understand & Continue")
}

func OnboardVersionWarningForm(ctx context.Context) error {

	return OnboardConfirmForm(ctx, `
# UNSTABLE VERSION WARNING
**PLEASE READ THESE WORDS CAREFULLY:**
It seems that you're going to use an unstable version of Manboster. Please note that:
1. It's normal to encounter bugs, crashes, and breaking changes in unstable versions.
2. As this is not a stable version, it's not contain ANY security patches and fixes.
3. This version's configuration may be incompatible with older versions and please aware the configuration changes.
4. If you encounter bugs, we appreciate you to commit to issues and we will fix it as soon as possible.
5. PLEASE DO NOT STORE ANY SENSITIVE AND IMPORTANT DATA IN THIS VERSION! As it's unstable and we are unsure that this application will work as is.
`, "Do you understand the risks and wish to proceed?", "I Understand & Continue")

}

func OnboardConfirmForm(ctx context.Context, tips string, confirmTitle string, confirmContent string) error {
	agree := false

	width, _, _ := term.GetSize(os.Stdout.Fd())
	if width == 0 {
		width = 80
	}
	textWidth := width - 8

	renderer, err := glamour.NewTermRenderer(
		glamour.WithEnvironmentConfig(),
		glamour.WithWordWrap(textWidth),
	)
	if err != nil {
		panic(err)
	}
	renderedMD, _ := renderer.Render(tips)
	defer func(renderer *glamour.TermRenderer) {
		err := renderer.Close()
		if err != nil {
			panic(err)
		}
	}(renderer)

	fmt.Println(boxStyle.Render(renderedMD))
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(confirmTitle).
				Affirmative(confirmContent).
				Negative("Exit Now").
				Value(&agree),
		),
	)
	err = form.Run()
	if !agree {
		os.Exit(0)
	}
	clearScreen()
	return err
}

func ContinueConfirm(ctx context.Context, content string) bool {
	agree := false

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(fmt.Sprintf("%s\nContinue?", content)).
				Affirmative("Continue").
				Negative("Skip").
				Value(&agree),
		),
	)

	err := form.Run()
	if err != nil {
		os.Exit(0)
	}
	return agree
}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return
	}
}
