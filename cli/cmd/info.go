package cmd

import (
	"fmt"

	"github.com/gocomply/oscalkit/pkg/oscal/constants"
	"github.com/gocomply/oscalkit/pkg/oscal_source"
	"github.com/gocomply/oscalkit/types/oscal/catalog"
	"github.com/urfave/cli"
)

// Catalog generates json/xml catalogs
var Info = cli.Command{
	Name:      "info",
	Usage:     "Provides information about particular OSCAL resource",
	ArgsUsage: "[file]",
	Action: func(c *cli.Context) error {
		for _, filePath := range c.Args() {
			os, err := oscal_source.Open(filePath)
			if err != nil {
				return cli.NewExitError(fmt.Sprintf("Could not open oscal file: %v", err), 1)
			}
			defer os.Close()

			o := os.OSCAL()
			switch o.DocumentType() {
			case constants.SSPDocument:
				fmt.Println("OSCAL System Security Plan")
				fmt.Println("ID:\t", o.SystemSecurityPlan.Uuid)
				printMetadata(o.SystemSecurityPlan.Metadata)
				return nil
			case constants.ComponentDocument:
				fmt.Println("OSCAL Component (represents information about particular software asset/component)")
				printMetadata(o.Component.Metadata)
				return nil
			case constants.ProfileDocument:
				fmt.Println("OSCAL Profile (represents subset of controls from OSCAL catalog(s))")
				fmt.Println("ID:\t", o.Profile.Uuid)
				printMetadata(o.Profile.Metadata)
				return nil
			case constants.CatalogDocument:
				fmt.Println("OSCAL Catalog (represents library of control assessment objectives and activities)")
				fmt.Println("ID:\t", o.Catalog.Uuid)
				printMetadata(o.Catalog.Metadata)
				return nil
			}
			return cli.NewExitError("Unrecognized OSCAL resource", 1)
		}
		return cli.NewExitError("No file provided", 1)
	},
}

func printMetadata(m *catalog.Metadata) {
	if m == nil {
		return
	}
	fmt.Println("Metadata:")
	fmt.Println("\tTitle:\t\t\t", m.Title)
	if m.Published != "" {
		fmt.Println("\tPublished:\t\t", m.Published)
	}
	if m.LastModified != "" {
		fmt.Println("\tLast Modified:\t\t", m.LastModified)
	}
	if m.Version != "" {
		fmt.Println("\tDocument Version:\t", m.Version)
	}
	if m.OscalVersion != "" {
		fmt.Println("\tOSCAL Version:\t\t", m.OscalVersion)
	}
}
