package main
import(
	"os/exec"
	"strings"
	"bufio"
	"fmt"
)

func ListInstalledApp(){
	for _, d := range GetInstalledApp(){
		fmt.Println(d)
	}

}


func GetInstalledApp() (csvData []string){
	c := exec.Command("cmd", "/C", "wmic product get name,version,vendor /format:csv")
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