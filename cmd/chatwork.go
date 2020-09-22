package cmd

import (
	"github.com/bitnami-labs/kubewatch/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// chatworkConfigCmd represents the hipchat subcommand
var chatworkConfigCmd = &cobra.Command{
	Use:   "chatwork",
	Short: "specific chatwork configuration",
	Long:  `specific chatwork configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := config.New()
		if err != nil {
			logrus.Fatal(err)
		}

		token, err := cmd.Flags().GetString("token")
		if err == nil {
			if len(token) > 0 {
				conf.Handler.Chatwork.Token = token
			}
		} else {
			logrus.Fatal(err)
		}
		room, err := cmd.Flags().GetString("room")
		if err == nil {
			if len(room) > 0 {
				conf.Handler.Chatwork.Room = room
			}
		} else {
			logrus.Fatal(err)
		}

		if err = conf.Write(); err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	chatworkConfigCmd.Flags().StringP("token", "t", "", "Specify chatwork token")
	chatworkConfigCmd.Flags().StringP("room", "r", "", "Specify chatwork room")
	chatworkConfigCmd.Flags().StringP("url", "u", "", "Specify chatwork server url")
}
