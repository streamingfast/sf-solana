package md

import (
	"bufio"
	"context"
	"fmt"

	"github.com/dfuse-io/dstore"
)

func readFile(filepath string, f func(line string) error) error {
	reader, _, _, err := dstore.OpenObject(context.Background(), filepath)
	if err != nil {
		return fmt.Errorf("opening file: %w", err)
	}
	defer reader.Close()

	bufReader := bufio.NewReader(reader)
	var line string
	for {
		line, err = bufReader.ReadString('\n')
		if err != nil {
			break
		}

		if err := f(line); err != nil {
			return fmt.Errorf("error processing line: %w", err)
		}
	}
	return nil
}