package bastien

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Lilierd/MeGoEnFait/project/bastien/activity"
)

func GetPERTData(path string) [][]string {
	f, err := os.Open(path)

	var tab [][]string

	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		tab = append(tab, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return tab
}

func GetPERTArray(data [][]string) []activity.Activity {
	var tab []activity.Activity

	for _, element := range data {
		var parents []activity.Activity

		if element[2] != "*" {
			parentsString := strings.Split(element[2], ",")
			for _, parent := range parentsString {
				for _, activite := range tab {
					if parent == activite.Name {
						parents = append(parents, activite)
					}
				}
			}
		}

		dur := parseInt(element[1])

		tab = append(tab, activity.New(element[0], dur, parents))
	}

	return tab
}

func ResolvePERT([]activity.Activity) {

}
