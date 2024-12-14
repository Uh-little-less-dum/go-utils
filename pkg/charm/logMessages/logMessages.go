package charm_debug

import (
	"fmt"
	"os"
	"reflect"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
)

func LogCharmMessages(filePath string, prefix string, msg tea.Msg) {
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
