package main

import (
	"bufio"
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
	addName := addCmd.String("name", "", "Student name")
	addAddress := addCmd.String("address", "", "Address")
	addReason := addCmd.String("reason", "", "Reason")

	// flag for "edit biodata" command
	editCmd := flag.NewFlagSet("edit", flag.ExitOnError)
	// input for "edit biodata" command
	editId := editCmd.String("id", "", "Biodata Id")

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
		HandleAdd(addCmd, addName, addAddress, addReason)
	} else if os.Args[1] == "delete" {
		HandleDelete(deleteCmd, deleteId)
	} else if os.Args[1] == "edit" {
		HandleEdit(editCmd, editId)
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

func ValidateBiodata(addCmd *flag.FlagSet, name *string, address *string, reason *string) {
	addCmd.Parse(os.Args[2:])

	if *name == "" || *address == "" || *reason == "" {
		fmt.Println("please input all field that required for adding biodata!")
		addCmd.PrintDefaults()
		os.Exit(1)
	}

}

func HandleAdd(addCmd *flag.FlagSet, name *string, address *string, reason *string) {

	ValidateBiodata(addCmd, name, address, reason)

	biodatas := getBiodatas()

	lastId := biodatas[len(biodatas)-1].Id

	// string to int
	intId, _ := strconv.Atoi(lastId)

	intId += 1

	// int to string
	newId := strconv.Itoa(intId)

	biodata := biodata{
		Id:      newId,
		Name:    *name,
		Address: *address,
		Reason:  *reason,
	}

	biodatas = append(biodatas, biodata)

	writeBiodatas(biodatas)
}

func HandleEdit(editCmd *flag.FlagSet, id *string) {
	editCmd.Parse(os.Args[2:])

	if *id == "" {
		fmt.Println("please input the id of biodata that you to edit!")
		editCmd.PrintDefaults()
		os.Exit(1)
	}

	biodatas := getBiodatas()

	if *id != "" {
		id := *id
		fmt.Println(id)
		for _, biodata := range biodatas {
			if id == biodata.Id {
				scanner := bufio.NewScanner(os.Stdin)

				fmt.Println("Enter the name: ")
				scanner.Scan()
				name := scanner.Text()

				fmt.Println("Enter the address: ")
				scanner.Scan()
				address := scanner.Text()

				fmt.Println("Enter the reason: ")
				scanner.Scan()
				reason := scanner.Text()

				biodata.Name = name
				biodata.Address = address
				biodata.Reason = reason

				biodatas = append(biodatas, biodata)
				id, _ := strconv.Atoi(id)
				biodatas = RemoveBiodataIdx(biodatas, id-1)

				writeBiodatas(biodatas)
			}
		}
	}

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

		var idx int = 0

		for _, biodata := range biodatas {

			if id == biodata.Id {
				biodatas = RemoveBiodataIdx(biodatas, idx)
				writeBiodatas(biodatas)
				return
			}
			idx += 1
		}
		fmt.Println("the id is not found")
	}
}
