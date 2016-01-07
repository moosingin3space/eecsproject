package main

import (
	"os"
	"fmt"
	"flag"
	"path"
	"templates"
	"text/template"
)

type MakefileData struct {
	Cplusplus bool
	ExecName string
}

var use_cplusplus bool

func init() {
	flag.BoolVar(&use_cplusplus, "cpp", false, "generate a C++ project")
}

func handleError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	flag.Parse()

	project_name := flag.Arg(0)

	if project_name == "" {
		// Unspecified project name
		fmt.Fprintln(os.Stderr, "usage: eecsproject [OPTIONS] project_name")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if use_cplusplus {
		fmt.Printf("Generating a C++ stub project '%s'\n", project_name)
	} else {
		fmt.Printf("Generating a C stub project '%s'\n", project_name)
	}

	// Make the project directory
	fmt.Print("making project directory...")
	cwd, err := os.Getwd()
	handleError(err)
	project_path := path.Join(cwd, project_name)
	info, err := os.Stat(project_path)
	if err == nil && info.IsDir() {
		fmt.Println("already exists.")
		os.Exit(0)
	}
	err = os.Mkdir(project_path, os.FileMode(0755))
	handleError(err)
	fmt.Println("done!")

	// Dump a Vagrantfile in that directory
	fmt.Print("making Vagrantfile...")
	vagrantfile_path := path.Join(project_path, "Vagrantfile")
	file, err := os.Create(vagrantfile_path)
	handleError(err)
	data, err := templates.Asset("assets/Vagrantfile")
	handleError(err)
	_, err = file.Write(data)
	handleError(err)
	fmt.Println("done!")

	// Now, dump a Makefile in that directory
	fmt.Print("making Makefile...")
	makefile_path := path.Join(project_path, "Makefile")
	file, err = os.Create(makefile_path)
	handleError(err)
	data, err = templates.Asset("assets/Makefile.templ")
	handleError(err)
	makefile_templ, err := template.New("makefile").Parse(string(data))
	handleError(err)
	makefile_data := MakefileData{
		Cplusplus: use_cplusplus,
		ExecName: project_name,
	}
	err = makefile_templ.Execute(file, makefile_data)
	handleError(err)
	fmt.Println("done!")

	// Now, save the main executable
	var mainfile_name string
	if use_cplusplus {
		mainfile_name = "main.cpp"
	} else {
		mainfile_name = "main.c"
	}
	mainfile_path := path.Join(project_path, mainfile_name)
	data, err = templates.Asset("assets/" + mainfile_name)
	fmt.Printf("making %s...", mainfile_name)
	handleError(err)
	file, err = os.Create(mainfile_path)
	handleError(err)
	_, err = file.Write(data)
	handleError(err)
	fmt.Println("done!")
	fmt.Println()

	fmt.Printf("Project %s generated!\n", project_name)
}
