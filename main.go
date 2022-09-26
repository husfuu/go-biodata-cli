package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

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

	// flag for "delete biodata" command
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	// input for "delete biodata" command
	deleteId := deleteCmd.String("id", "", "Biodata Id")

	if len(os.Args) < 2 {
		fmt.Print("expected 'get' command")
		os.Exit(1)
	} else if os.Args[1] == "get" {
		HandleGet(getCmd, getAll, getId)
	} else if os.Args[1] == "add" {
		HandleAdd(addCmd, addId, addName, addAddress, addReason)
	} else if os.Args[1] == "delete" {
		HandleDelete(deleteCmd, deleteId)
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
		fmt.Println("please input all field that required for adding biodata!")
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

func RemoveBiodataIdx(s []biodata, index int) []biodata {
	return s[:index+copy(s[index:], s[index+1:])]

}

func HandleDelete(deletCmd *flag.FlagSet, id *string) {
	deletCmd.Parse(os.Args[2:])

	if *id == "" {
		fmt.Println("please input the id of biodata that you want to delete!")
		deletCmd.PrintDefaults()
		os.Exit(1)
	}
	biodatas := getBiodatas()

	if *id != "" {
		id := *id

		for _, biodata := range biodatas {
			if id == biodata.Id {
				id, err := strconv.Atoi(id)
				if err != nil {
					panic(err)
				} else {
					fmt.Print("jalannnn")
					biodatas = RemoveBiodataIdx(biodatas, id-1)
					fmt.Println(biodatas)
					deleteBiodatas(biodatas)
				}

			}
		}

	}

}
