package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)

		if err != nil {
			log.Println(err)
			continue
		}

		out, err := dp.ActionInfo()

		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println(out)
	}
}
