package main

import(
	"os/exec"
	"fmt"
	"bufio"
	"strings"
	"encoding/csv"
	"io"
	"log"
)
func main(){
	//systeminfo /S xxx /U xxx /P yyy /FO CSV
	c := exec.Command("cmd", "/C", "systeminfo /FO CSV")
	bs, _ := c.Output()
	lines,_ := b(a(bs))
	headerLine := lines[0]
	valueLine := lines[1]
	header := CsvByLine(headerLine)
	line := CsvByLine(valueLine)
	for index, h := range header{
		fmt.Println(h)
		fmt.Println(line[index])
	}
	//
	//fmt.Println("============================================")
	//for _, h := range line{
	//	fmt.Println(h)
	//}
	//for _,line :=range lines{
	//	reader := csv.NewReader(strings.NewReader(line))
	//	for {
	//		lineA, error := reader.Read()
	//		if error == io.EOF {
	//			break
	//		} else if error != nil {
	//			log.Fatal(error)
	//		}
	//		fmt.Print(lineA)
	//		//people = append(people, Person{
	//		//	Firstname: line[0],
	//		//	Lastname:  line[1],
	//		//	Address: &Address{
	//		//		City:  line[2],
	//		//		State: line[3],
	//		//	},
	//		//})
	//	}
	//	fmt.Println("===")

	//}


	//
	//fmt.Println(a(bs))
	//datas := make(map[string]string)
	//header :=lines[0]
	//value :=lines[1]

	//for _,line :=range lines{
	//	fmt.Println(line)
	//	//stringSlice := strings.SplitN(line, ":",2)
	//	//
	//	//if len(stringSlice)>1 {
	//	//	fmt.Println(strings.Trim(stringSlice[0]," ")+"."+strings.Trim(stringSlice[1]," "))
	//	//}
	//
	//	//fmt.Println("----")
	//}


	//data :=BytesToString(string(bs[:]))
	//csvData,_ = StringToLines(data)
	//fmt.Println(data)
}

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

func a(data []byte) string {
	return string(data[:])
}


func b(s string) (lines []string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}

//systeminfo /?

//wmic bios get serialnumber -> serial number
//arp -a 10.1.30.66 -> Mac address
