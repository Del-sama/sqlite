package main

import "fmt"

type InputBuffer struct {
	buffer       *string
	bufferLength int
	inputLength  int
}

func (i *InputBuffer) newInputBuffer() {
	inputBuffer := InputBuffer{}
	inputBuffer.buffer = nil
	inputBuffer.bufferLength = 0
	inputBuffer.inputLength = 0
}

func printPrompt() {
	fmt.Printf("db >")
}

func readInput(stdin io.Reader) (string, error) {
	reader := bufio.NewReader(stdin)
	return reader.ReadString('\n')
}

int main(int argc, char* argv[]) {
	InputBuffer* input_buffer = new_input_buffer();
	while (true) {
	  print_prompt();
	  read_input(input_buffer);

	  if (strcmp(input_buffer->buffer, ".exit") == 0) {
		close_input_buffer(input_buffer);
		exit(EXIT_SUCCESS);
	  } else {
		printf("Unrecognized command '%s'.\n", input_buffer->buffer);
	  }
	}
  }


func repl(argc int, argv *[]string) error {
	inputBuffer := newInputBuffer()
	for {
		printPrompt()
		s, err := readInput()
		if err != nil {
			return err
		}


	}
}
