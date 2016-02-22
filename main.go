package main

import (
	"fmt"
	"strings"
	"flag"
	"os"
	"github.com/faryon93/eagle-go"
)

func main() {
	// parse the command line arguments
	var outputFile string
	flag.StringVar(&outputFile, "out", "", "output file, if omitted result is written to stdout")
	flag.Parse()

	// check if the obligatory arguments are present
	inputFile := flag.Arg(0)
	partName := flag.Arg(1)
	if inputFile == "" || partName == "" {
		usage()
		os.Exit(-1)
	}

	// load the schmeatic and get the part which is supplied by the user
	schematic, err := eagle.Load(inputFile)	
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	// get the part an its deviceset
	fpga := schematic.GetPart(partName)
	if fpga == nil {
		fmt.Println("could not find part", partName, "in schematic!")
		os.Exit(-1)
	}
	device := fpga.GetDeviceset()
	if device == nil {
		fmt.Println("could not find deviceset", fpga.Deviceset, "in library", fpga.Library)
		os.Exit(-1)
	}

	// generate the ucf file content
	ucf := ""
	for _, net := range schematic.Nets {
		pin := net.GetPartPin(fpga.Name)
		if (pin != "" && isIoPin(pin)) {
			ucf += fmt.Sprintf("NET %-15s LOC = %-5s | IOSTANDARD = LVCMOS33;\n", "\"" + net.Name + "\"", device.GetPad(pin))
		}
	}

	// decide where the ucf file should go
	if (outputFile == "") {
		fmt.Println(ucf)	

	// we want to write it to a file
	} else {
		f, err := os.OpenFile(outputFile, os.O_WRONLY | os.O_CREATE, 0622)
		if (err != nil) {
			fmt.Println(err)
			os.Exit(-1)
		}
		defer f.Close()

		if _, err = f.WriteString(ucf); err != nil {
			panic(err)
		}
	}
	
}


// ----------------------------------------------------------------------------------
//  helper functions
// ----------------------------------------------------------------------------------

func usage() {
	fmt.Println("usage: sch2ucf [--out <file>] input-file part-name")
}

func isIoPin(pinName string) bool {
	// all pins which may be used in the ucf file
	// start with IP or IO
	return strings.HasPrefix(pinName, "IO") ||
		   strings.HasPrefix(pinName, "IP")
}