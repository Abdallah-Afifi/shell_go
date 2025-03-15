func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		// read keyboard input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(os.Stderr, err)
		}
		// handle execution
		if err = execute_input(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execute_input(input string) error {
	//  remove newline char
	input = strings.TrimSuffix(input, "\n")

	// split on spaces to separate command and arguments
	args := strings.Split(input, " ")

	// check for built-in commands
	switch args[0] {
	case "cd":
		// cd to home with no path not implemented yet
		if len(args) < 2 {
			return errors.New("cd: no such file or directory")
		}
		// change directory
		return os.Chdir(args[1])
	case "exit":
		// exit the shell
		os.Exit(0)
	}

	// pass arguments
	cmd := exec.Command(args[0], args[1:]...)

	// prepare command
	cmd := exec.Command(input)

	// set output device
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// execute and return error
	return cmd.Run()
}

