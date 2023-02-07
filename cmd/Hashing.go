/*
Copyright © 2023 Facundo Lerena  <contacto@faculerena.com.ar>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/term"
	"log"
	"strings"
	"syscall"
)

// HashingCmd represents the Hashing command
var HashingCmd = &cobra.Command{
	Use:   "Hashing",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		oldState, err := term.GetState(syscall.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		defer term.Restore(syscall.Stdin, oldState)

		algos := []algorithms{
			{Name: "SHA1", Output: "160 bit", Recommendation: ""},
			{Name: "SHA256", Output: "256 bit", Recommendation: ""},
			{Name: "SHA512", Output: "512 bit", Recommendation: ""},
		}

		templates := &promptui.SelectTemplates{
			Label:    "{{ . }}?",
			Active:   "\U00002192   {{ .Name | cyan }} ({{ .Output | white }})",
			Inactive: "  {{ .Name | white }} ",
			Selected: "\U00002192 {{ .Name | cyan }}",
			Details: `
--------- Algorithm ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Output:" | faint }}	{{ .Output }}`,
		}

		searcher := func(input string, index int) bool {
			thisAlgo := algos[index]
			name := strings.Replace(strings.ToLower(thisAlgo.Name), " ", "", -1)
			input = strings.Replace(strings.ToLower(input), " ", "", -1)

			return strings.Contains(name, input)
		}

		prompt := promptui.Select{
			Label:     "Spicy Level",
			Items:     algos,
			Templates: templates,
			Size:      4,
			Searcher:  searcher,
		}

		s, err := getInput()
		if err != nil {
			log.Fatal(err)
		}

		i, _, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		fmt.Printf("Hashing input with %q\n", algos[i].Name)

		switch algos[i].Name {
		case "SHA1":
			s.sha1()
		case "SHA256":
			s.sha256()
		case "SHA512":
			s.sha512()
		}
	},
}

func init() {
	rootCmd.AddCommand(HashingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// HashingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// HashingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type algorithms struct {
	Name           string
	Output         string
	Recommendation string
}

type input []byte

func getInput() (input, error) {
	fmt.Println("Input text to hash")
	fmt.Print("> ")
	s, err := terminal.ReadPassword(0)
	if err != nil {
		return nil, err
	}
	return s, nil
}
func (s *input) sha1() {
	h := sha1.New()
	h.Write(*s)
	fmt.Printf("% x\n", h.Sum(nil))
}
func (s *input) sha256() {
	h := sha256.New()
	h.Write(*s)
	fmt.Printf("% x\n", h.Sum(nil))
}
func (s *input) sha512() {
	h := sha512.New()
	h.Write(*s)
	fmt.Printf("% x\n", h.Sum(nil))
}
