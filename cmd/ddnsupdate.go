package cmd

import (
	"fmt"
	"net"
	"strings"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [subdomain] [token]",
	Short: "Update the IP address associated with the domain",
	Long: `This command updates the IP address associated with the domain.
For example:

# Updates with public IP address:
duckdnsupdate update mysubdomain 0000-0000-0000-0000

# Updates with a specific IP address:
duckdnsupdate update mysubdomain 0000-0000-0000-0000 --ip-addr X.X.X.X
	`,
	Args: cobra.ExactArgs(2),
	Run:  update,
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().String("ip-addr", "", "The IP address to associate with the domain")
}

func update(cmd *cobra.Command, args []string) {
	subdomain := args[0]
	token := args[1]
	ip, _ := cmd.Flags().GetString("ip-addr")
	if net.ParseIP(ip) == nil && ip != "" {
		fmt.Println("Invalid IP address: " + ip)
		return
	}

	if subdomain == "" {
		fmt.Println("Invalid domain: field cannot be empty")
		return
	}

	if token == "" || len(token) != 36 || strings.Count(token, "-") != 4 {
		fmt.Println("Invalid token: " + token)
		return
	}

	var publicIP string
	if ip != "" {
		publicIP = ip
	} else {
		var err error
		publicIP, err = GetPublicIP()
		if err != nil {
			fmt.Println("Error: " + err.Error())
			return
		}
	}

	// Check if the IP is already associated with the domain
	associatedIP, _ := net.LookupIP(subdomain + ".duckdns.org")
	fmt.Printf("Subdomain: %s\nAssociated IP was: %s\nAssigning with: %s\n", subdomain, associatedIP[0].String(), publicIP)
	if associatedIP[0].String() == publicIP {
		fmt.Println("IP is already associated with the domain")
	} else {
		_, err := MakeAPICall(subdomain, token, publicIP)
		if err != nil {
			fmt.Println("Error: " + err.Error())
			return
		} else {
			fmt.Println("IP address successfully updated")
		}
	}
}
