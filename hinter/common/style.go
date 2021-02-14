package common

import te "github.com/muesli/termenv"

const FocusedTextColor = "205"

var (
	Color         = te.ColorProfile().Color
	FocusedPrompt = te.String("> ").Foreground(Color(FocusedTextColor)).String()
	BlurredPrompt = "> "
)
