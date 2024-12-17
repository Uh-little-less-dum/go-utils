package charm_debug

import (
	"fmt"
	"os"
	"reflect"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
)

func LogCharmMessages(filePath string, prefix string, msg tea.Msg) {
	if filePath == "" {
		filePath = "/Users/bigsexy/Desktop/Go/projects/ulld/charmLog.log"
	}
	b, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Cannot log charm messages.")
	}
	mt := reflect.TypeOf(msg)
	err = os.WriteFile(filePath, []byte(fmt.Sprintf("%s\n%s: Type=%s Value=%v", string(b), prefix, mt, msg)), 0777)
	if err != nil {
		log.Fatal("Error occurred while logging charm messages.")
	}
}

func LogString(filePath, prefix, value string) {
	if filePath == "" {
		filePath = "/Users/bigsexy/Desktop/Go/projects/ulld/charmLog.log"
	}
	b, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Cannot log charm messages.")
	}
	err = os.WriteFile(filePath, []byte(fmt.Sprintf("%s\n%s: %s", string(b), prefix, value)), 0777)
	if err != nil {
		log.Fatal("Error occurred while logging charm messages.")
	}
}
