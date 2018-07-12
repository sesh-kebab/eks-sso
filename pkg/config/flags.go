package config

import (
	"fmt"
	"os"
)

type FlagType int

const (
	FTString FlagType = 0
	FTBool   FlagType = 1
)

type AppFlags struct {
	Name     string
	Value    string
	EnvVar   string
	Required bool
	FlagType FlagType
	Validate func(string) error
}

func GetAppFlags() []AppFlags {
	return []AppFlags{
		{
			Name:     "p, port",
			Value:    DefaultPort,
			EnvVar:   EnvVarPort,
			FlagType: FTString,
		},
		{
			Name:     "d, debug",
			EnvVar:   EnvVarDebug,
			FlagType: FTBool,
		},
		{
			Name:     "aws-region",
			Value:    DefaultAWSRegion,
			EnvVar:   EnvVarAWSRegion,
			FlagType: FTString,
		},
		{
			Name:     FlagAuth0Domain,
			Value:    DefaultAuth0Domain,
			EnvVar:   EnvVarAuth0Domain,
			FlagType: FTString,
			Required: true,
		},
		{
			Name:     "auth0-client-id",
			EnvVar:   EnvVarAuth0ClientID,
			FlagType: FTString,
			Required: true,
		},
		{
			Name:     "auth0-connection",
			EnvVar:   EnvVarAuth0Connection,
			FlagType: FTString,
			Required: true,
		},
		{
			Name:     "cluster-name",
			EnvVar:   EnvVarClusterName,
			FlagType: FTString,
			Required: true,
		},
		{
			Name:     "cluster-region",
			Value:    "us-west-2",
			EnvVar:   EnvVarClusterRegion,
			FlagType: FTString,
			Required: true,
		},
		{
			Name:     FlagInCluster,
			EnvVar:   EnvVarInCluster,
			FlagType: FTBool,
		},
		{
			Name:     FlagKubeConfigPath,
			Value:    "",
			EnvVar:   EnvVarKubeConfigPath,
			FlagType: FTString,
			Validate: func(value string) error {
				if _, err := os.Stat(value); os.IsNotExist(err) {
					return fmt.Errorf("invalid path specified for flag '%s' or env var '%s'",
						FlagKubeConfigPath,
						EnvVarKubeConfigPath,
					)
				}

				return nil
			},
		},
	}
}
