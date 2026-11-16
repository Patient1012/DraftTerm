package main

import (
	"charm.land/bubbles/v2/textarea"
)

type Model struct {
	filename   string // name of of the file
	ifModified bool   // state of the file
	width      int    // width of a row or the number of columns set by the terminal
	height     int    // height of text number of rows set by terminal

	editor textarea.Model
}
