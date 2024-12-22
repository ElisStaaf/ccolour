package mycolour

import (
	"fmt"
	"testing"
)

func TestBlack(t *testing.T) {
	fmt.Printf("normal: %s\n", Black(Normal, "black"))
  fmt.Printf("bold: %s\n", Black(Bold, "black"))
  fmt.Printf("underline: %s\n", Black(Underline, "black"))
  fmt.Printf("negative: %s\n", Black(Negative, "black"))
}

func TestRed(t *testing.T) {
	fmt.Printf("normal: %s\n", Red(Normal, "red"))
  fmt.Printf("bold: %s\n", Red(Bold, "red"))
  fmt.Printf("underline: %s\n", Red(Underline, "red"))
  fmt.Printf("negative: %s\n", Red(Negative, "red"))

}

func TestGreen(t *testing.T) {
	fmt.Printf("normal: %s\n", Green(Normal, "green"))
  fmt.Printf("bold: %s\n", Green(Bold, "green"))
  fmt.Printf("underline: %s\n", Green(Underline, "green"))
  fmt.Printf("negative: %s\n", Green(Negative, "green"))
}

func TestYellow(t *testing.T) {
	fmt.Printf("normal: %s\n", Yellow(Normal, "yellow"))
  fmt.Printf("bold: %s\n", Yellow(Bold, "yellow"))
  fmt.Printf("underline: %s\n", Yellow(Underline, "yellow"))
  fmt.Printf("negative: %s\n", Yellow(Negative, "yellow"))
}

func TestBlue(t *testing.T) {
	fmt.Printf("normal: %s\n", Blue(Normal, "blue"))
  fmt.Printf("bold: %s\n", Blue(Bold, "blue"))
  fmt.Printf("underline: %s\n", Blue(Underline, "blue"))
  fmt.Printf("negative: %s\n", Blue(Negative, "blue"))
}

func TestMagenta(t *testing.T) {
	fmt.Printf("normal: %s\n", Magenta(Normal, "magenta"))
  fmt.Printf("bold: %s\n", Magenta(Bold, "magenta"))
  fmt.Printf("underline: %s\n", Magenta(Underline, "magenta"))
  fmt.Printf("negative: %s\n", Magenta(Negative, "magenta"))
}

func TestCyan(t *testing.T) {
	fmt.Printf("normal: %s\n", Cyan(Normal, "cyan"))
  fmt.Printf("bold: %s\n", Cyan(Bold, "cyan"))
  fmt.Printf("underline: %s\n", Cyan(Underline, "cyan"))
  fmt.Printf("negative: %s\n", Cyan(Negative, "cyan"))
}

func TestWhite(t *testing.T) {
	fmt.Printf("normal: %s\n", White(Normal, "white"))
  fmt.Printf("bold: %s\n", White(Bold, "white"))
  fmt.Printf("underline: %s\n", White(Underline, "white"))
  fmt.Printf("negative: %s\n", White(Negative, "white"))
}
