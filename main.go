package main

import (
	"fmt"
	"flag"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io/ioutil"
	"strings"

)

func main() {
	var output_file_type string
	flag.StringVar(&output_file_type, "t", "", "target file type")
	flag.Parse()

	if output_file_type == "" {
		fmt.Println("Please provide a target file type using the -t flag.")
		return
	} else {
		fmt.Println("Target file type is:", output_file_type)
	}

	files , _ := ioutil.ReadDir("./Input")

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fmt.Println("Processing file:", file.Name())
		input_file := "./Input/" + file.Name()
		base := strings.Split(file.Name(), ".")[0]
		output_file := "./Output/" + base + "." + output_file_type
		
		err := ffmpeg.
			Input(input_file).
			Output(output_file).
			OverWriteOutput().
			Run()

		if err != nil {
			fmt.Printf("Error converting file %s: %v\n", file.Name(), err)
			return
		}
		fmt.Println("Converted", input_file, "to", output_file)
	}
	fmt.Println("All files converted successfully!")
	
}
