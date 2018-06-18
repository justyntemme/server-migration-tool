package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/tealeg/xlsx"
)

func main() {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	domainFile, err := os.Open("./domains.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer domainFile.Close()

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {

		fmt.Printf(err.Error())
	}
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "Domain"
	cell = row.AddCell()
	cell.Value = "IP"
	cell = row.AddCell()
	cell.Value = "MX Record"
	cell = row.AddCell()
	cell.Value = "Name Server"

	scanner := bufio.NewScanner(domainFile)
	for scanner.Scan() {
		row = sheet.AddRow()
		//insert domain
		cell = row.AddCell()
		cell.Value = (scanner.Text())
		//ip lookup
		cell = row.AddCell()
		ips, err := net.LookupIP(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		for _, ip := range ips {
			cell.Value = ip.String()
		}
		//mx lookup
		cell = row.AddCell()
		mxs, err := net.LookupMX(scanner.Text())
		cell.Value = (mxs[0].Host)

		//nameServer lookup
		cell = row.AddCell()
		nss, err := net.LookupNS(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		for _, ns := range nss {
			cell.Value = ns.Host
		}

	}
	err = file.Save("domain-info.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
