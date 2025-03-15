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

	// prepare command
	cmd := exec.Command(input)

	// set output device
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// execute and return error
	return cmd.Run()
}

