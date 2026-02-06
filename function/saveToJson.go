package function

import (
	"encoding/json"
	"fmt"
	"os"
	"porigo/model"
)

func saveToJson(finalResult []model.Result) {
	file, err := json.MarshalIndent(finalResult, "", " ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("result.json", file, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("saved result to result.json\n")
}
