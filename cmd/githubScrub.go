/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/kubefirst/kubefirst/internal/flagset"
	"github.com/kubefirst/kubefirst/internal/github"
	"github.com/kubefirst/kubefirst/internal/githubWrapper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// githubRemoveCmd represents the githubRemove command
var githubRealRemoveCmd = &cobra.Command{
	Use:   "real-remove-github",
	Short: "Undo github setup",
	Long:  `TBD`,
	RunE: func(cmd *cobra.Command, args []string) error {
		globalFlags, err := flagset.ProcessGlobalFlags(cmd)
		if err != nil {
			return err
		}

		if globalFlags.DryRun {
			log.Printf("[#99] Dry-run mode, githubRemoveCmd skipped.")
			return nil
		}

		log.Println("Org used:", viper.GetString("github.owner"))
		log.Println("dry-run:", globalFlags.DryRun)

		informUser("Destroying repositories with terraform in GitHub", globalFlags.SilentMode)

		github.DestroyGitHubTerraform(globalFlags.DryRun)

		gitWrapper := githubWrapper.New()
		err = gitWrapper.RemoveRepo(viper.GetString("github.owner"), "gitops")
		if err != nil {
			log.Println(err)
		}
		err = gitWrapper.RemoveRepo(viper.GetString("github.owner"), "metaphor")
		if err != nil {
			log.Println(err)
		}
		err = gitWrapper.RemoveRepo(viper.GetString("github.owner"), "metaphor-go")
		if err != nil {
			log.Println(err)
		}
		err = gitWrapper.RemoveRepo(viper.GetString("github.owner"), "metaphor-frontend")
		if err != nil {
			log.Println(err)
		}
		err = gitWrapper.RemoveSSHKey(viper.GetInt64("github.ssh.keyId"))
		if err != nil {
			log.Println("Trying to remove key failed:", err)
		}

		err = gitWrapper.RemoveTeam(viper.GetString("github.owner"), "admins")
		if err != nil {
			log.Println("Remove team failed:", err)
		}
		err = gitWrapper.RemoveTeam(viper.GetString("github.owner"), "developers")
		if err != nil {
			log.Println("Remove team failed:", err)
		}

		informUser("GitHub terraform destroyed", globalFlags.SilentMode)

		log.Printf("GitHub terraform Executed with Success")
		return nil
	},
}

func init() {
	actionCmd.AddCommand(githubRealRemoveCmd)
	currentCommand := githubRealRemoveCmd
	flagset.DefineGlobalFlags(currentCommand)
}
