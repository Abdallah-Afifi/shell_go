func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		// read keyboard input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(os.Stderr, err)
		}
	}
}