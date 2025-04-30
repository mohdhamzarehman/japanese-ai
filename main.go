package main
import {
	"os"
	"github.com/mohdhamzarehman/google-tts/cmd"
}
func main() {
	cli := &cmd.CLI{ErrStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}