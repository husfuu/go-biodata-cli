package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/alexeyco/simpletable"
)

func main() {
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	// flag for "get biodata" command
	getAll := getCmd.Bool("all", false, "Get all biodata")
	getId := getCmd.String("id", "", "Biodata Id")

	// flag for "add biodata" command
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	// input for "add biodata" command
	addId := addCmd.String("id", "", "Biodata Id")
	addName := addCmd.String("name", "", "Student name")
	addAddress := addCmd.String("address", "", "Address")
	addReason := addCmd.String("reason", "", "Reason")

	if len(os.Args) < 2 {
		fmt.Print("expected 'get' command")
		os.Exit(1)
	} else if os.Args[1] == "get" {
		HandleGet(getCmd, getAll, getId)
	} else if os.Args[1] == "add" {
		HandleAdd(addCmd, addId, addName, addAddress, addReason)
	}

}

func HandleGet(getCmd *flag.FlagSet, isAll *bool, id *string) {
	getCmd.Parse(os.Args[2:])
	table := simpletable.New()

	if *isAll == false && *id == "" {
		fmt.Print("id is required or specify --all for all videos")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	if *isAll {
		biodatas := getBiodatas()
		table.Header = &simpletable.Header{
			Cells: []*simpletable.Cell{
				{Align: simpletable.AlignCenter, Text: "Id"},
				{Align: simpletable.AlignCenter, Text: "Name"},
				{Align: simpletable.AlignCenter, Text: "Address"},
				{Align: simpletable.AlignCenter, Text: "Reason"},
			},
		}
		for _, biodata := range biodatas {
			r := []*simpletable.Cell{
				{Align: simpletable.AlignLeft, Text: fmt.Sprintf(biodata.Id)},
				{Align: simpletable.AlignLeft, Text: fmt.Sprintf(biodata.Name)},
				{Align: simpletable.AlignLeft, Text: fmt.Sprintf(biodata.Address)},
				{Align: simpletable.AlignLeft, Text: fmt.Sprintf(biodata.Reason)},
			}
			table.Body.Cells = append(table.Body.Cells, r)
		}
		table.SetStyle(simpletable.StyleCompactLite)
		fmt.Println(table.String())
		return
	}
	if *id != "" {
		biodatas := getBiodatas()
		table.Header = &simpletable.Header{
			Cells: []*simpletable.Cell{
				{Align: simpletable.AlignCenter, Text: "Id"},
				{Align: simpletable.AlignCenter, Text: "Name"},
				{Align: simpletable.AlignCenter, Text: "Address"},
				{Align: simpletable.AlignCenter, Text: "Reason"},
			},
		}
		id := *id
		for _, biodata := range biodatas {
			if id == biodata.Id {
				r := []*simpletable.Cell{
					{Align: simpletable.AlignLeft, Text: fmt.Sprintf(biodata.Id)},
					{Align: simpletable.AlignLeft, Text: fmt.Sprintf(biodata.Name)},
					{Align: simpletable.AlignLeft, Text: fmt.Sprintf(biodata.Address)},
					{Align: simpletable.AlignLeft, Text: fmt.Sprintf(biodata.Reason)},
				}
				table.Body.Cells = append(table.Body.Cells, r)
			}
		}
		table.SetStyle(simpletable.StyleCompactLite)
		fmt.Println(table.String())
	}

}

func ValidateBiodata(addCmd *flag.FlagSet, id *string, name *string, address *string, reason *string) {
	addCmd.Parse(os.Args[2:])

	if *id == "" || *name == "" || *address == "" || *reason == "" {
		fmt.Print("please input all field that required for adding biodata!")
		addCmd.PrintDefaults()
		os.Exit(1)
	}

}

func HandleAdd(addCmd *flag.FlagSet, id *string, name *string, address *string, reason *string) {

	ValidateBiodata(addCmd, id, name, address, reason)

	biodata := biodata{
		Id:      *id,
		Name:    *name,
		Address: *address,
		Reason:  *reason,
	}

	biodatas := getBiodatas()
	biodatas = append(biodatas, biodata)

	addBiodatas(biodatas)
}
