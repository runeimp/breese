APP_NAME := "Breese Morse Code"


@_default:
	just _term-title "{{APP_NAME}}"
	just _term-wipe
	just --list


# Run the breese command line app
run +args='':
	go run cmd/breese/main.go


# Set Terminal Title
_term-title title='':
	@printf "\033]0;%s\007" "{{title}}"

# Wipe Terminal Buffer and Scrollback Buffer
_term-wipe:
	#!/bin/sh
	if [ ${#VISUAL_STUDIO_CODE} -gt 0 ]; then
		clear
	elif [ ${KITTY_WINDOW_ID} -gt 0 ]; then
		printf '\033c'
	elif [ "$(uname)" == 'Darwin' ] && [ ${#TMUX} -eq 0 ]; then
		osascript -e 'tell application "System Events" to keystroke "k" using command down'
	elif [ -f "$(which tput)" ]; then
		tput reset
		if [ ${#TMUX} -gt 0 ]; then
			tput cup 0 0
		fi
	elif [ -f "$(which reset)" ]; then
		reset
	else
		clear
		# alias cls="clear; printf '\33[3J'"
		# echo -ne '\033]50;ClearScrollback\a' # For iTerm2
	fi
	just _term-title "{{APP_NAME}}"

