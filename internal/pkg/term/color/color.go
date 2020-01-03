// Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package color provides functionality to displayed colored text on the terminal.
package color

import (
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2/core"
	"github.com/fatih/color"
)

// Predefined colors.
// Refer to https://en.wikipedia.org/wiki/ANSI_escape_code to validate if colors would
// be visible on white or black screen backgrounds.
var (
	Grey       = color.New(color.FgWhite)
	Red        = color.New(color.FgHiRed)
	Cyan       = color.New(color.FgCyan)
	HiCyan     = color.New(color.FgHiCyan)
	Bold       = color.New(color.Bold)
	BoldItalic = color.New(color.Bold).Add(color.Italic)
)

const colorEnvVar = "COLOR"

var lookupEnv = os.LookupEnv

// DisableColorBasedOnEnvVar determines whether the CLI will produce color
// output based on the environment variable, COLOR.
func DisableColorBasedOnEnvVar() {
	value, exists := lookupEnv(colorEnvVar)
	if !exists {
		// if the COLOR environment variable is not set
		// then follow the settings in the color library
		// since it's dynamically set based on the type of terminal
		// and whether stdout is connected to a terminal or not.
		core.DisableColor = color.NoColor
		return
	}

	if strings.ToLower(value) == "false" {
		core.DisableColor = true
		color.NoColor = true
	} else if strings.ToLower(value) == "true" {
		core.DisableColor = false
		color.NoColor = false
	}
}

// Emphasize colors the string to denote that it as important, and returns it.
func Emphasize(s string) string {
	return BoldItalic.Sprint(s)
}

// HighlightUserInput colors the string to denote it as an input from standard input, and returns it.
func HighlightUserInput(s string) string {
	return Cyan.Sprint(s)
}

// HighlightResource colors the string to denote it as a resource created by the CLI, and returns it.
func HighlightResource(s string) string {
	return HiCyan.Sprint(s)
}

// HighlightCode wraps the string with the ` character, colors it to denote it's a code block, and returns it.
func HighlightCode(s string) string {
	return HiCyan.Sprintf("`%s`", s)
}
