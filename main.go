package main

import "os"
import "fmt"

import "github.com/nsf/termbox-go"
import "github.com/mattn/go-runewidth"

func print_message(col, row int, fg, bg termbox.Attribute, message string) {
  for _, ch := range message{
    termbox.SetCell(col, row, ch, fg, bg)
    col += runewidth.RuneWidth(ch)
  }
}

func run_editor() {
  err := termbox.Init()
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  print_message(25, 11, termbox.ColorDefault, termbox.ColorDefault, "LAMO - A terminal text editor")
  termbox.Flush()
  termbox.PollEvent()
  defer termbox.Close()
}

func main() {
  run_editor();
}
