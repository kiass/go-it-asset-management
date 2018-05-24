package service
import(
	"os/exec"
	"strings"
	"bufio"
	"encoding/csv"
	"io"
	"log"
)

func CsvByLine(s string) (lines []string){
	reader := csv.NewReader(strings.NewReader(s))
	for {
		lineA, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		lines = append(lines, lineA...)
	}
	return
}

func GetInstalledApp() (csvData []string){
	c := exec.Command("cmd", "/C", "wmic product get name,version,vendor,InstallDate /format:csv")
	bs, _ := c.Output()
	data :=BytesToString(bs)
	csvData,_ = StringToLines(data)
	return
}

func BytesToString(data []byte) string {
	return string(data[:])
}

func StringToLines(s string) (lines []string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}