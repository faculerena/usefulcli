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
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"strings"
)

// PingCmd represents the Ping command
var PingCmd = &cobra.Command{
	Use:   "Ping",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		var ip, port string

		fmt.Print("Enter an ip, leave blank for localhost.\n > ")
		fmt.Scanln(&ip)

		fmt.Print("Enter a port.\n > ")
		fmt.Scanln(&port)

		var ipToPing string
		if ip == "" {
			ipToPing = "localhost:" + port
		} else {
			ipToPing = ip + ":" + port
		}

		out, _ := exec.Command("ping", ipToPing, "-c 5", "-i 3", "-w 10").Output()
		if strings.Contains(string(out), "Unreachable") {
			fmt.Println("Dead")
		} else {
			fmt.Println("Alive")
		}
	},
}

func init() {
	rootCmd.AddCommand(PingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// PingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// PingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
