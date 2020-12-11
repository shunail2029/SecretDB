package cli

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

func findItemBeginnings(data []byte) []int {
	var idxs []int
	cnt := 0
	for i, d := range data {
		if cnt == 0 {
			idxs = append(idxs, i)
		}
		if d == '{' {
			cnt++
		} else if d == '}' {
			cnt--
		}
	}
	return idxs
}

func printOutput(toPrint interface{}, outputFormat string, indent bool) error {
	var (
		out []byte
		err error
	)

	switch outputFormat {
	case "text":
		out, err = yaml.Marshal(&toPrint)

	case "json":
		if indent {
			out, err = json.MarshalIndent(toPrint, "", "  ")
		} else {
			out, err = json.Marshal(toPrint)
		}
	}

	if err != nil {
		return err
	}

	fmt.Println(string(out))
	return nil
}
