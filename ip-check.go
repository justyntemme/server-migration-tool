package main

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/tealeg/xlsx"
)

func main() {
	MTI := []string{"70.32.113.40", "70.32.113.42", "70.32.113.101", "70.32.113.40", "205.186.175.166"}
	excelFileName := "domain-info.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.Fatal(err)
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				lookup, err := net.LookupNS(text)
				if strings.Contains(text, "gotodja") {
					fmt.Println(lookup)
				}
				if err != nil {
					//fmt.Println(text + " " + err.Error())
					continue
				}
				if err == nil {
					for _, host := range lookup {

						if strings.Contains(host.Host, "mediatemple") {
							fmt.Println("Alert! " + lookup[0].Host)
							style := xlsx.NewStyle()
							fill := xlsx.NewFill("solid", "8B2323", "8B2323")
							style.Fill = *fill
							cell.SetStyle(style)

						}
					}
				}
				lookupIP, err := net.LookupIP(text)
				if err != nil {
					continue
				}
				if err == nil {
					for _, host := range lookupIP {
						for _, mt := range MTI {
							if strings.Contains(host.String(), mt) {
								fmt.Println("Alert! " + lookupIP[0].String())
								style := xlsx.NewStyle()
								fill := xlsx.NewFill("solid", "8B2323", "8B2323")
								style.Fill = *fill
								cell.SetStyle(style)
							}
						}
					}
				}

			}
		}
	}

	err = xlFile.Save("domain-info.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
