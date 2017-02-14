package model

import (
	// "time"
	"fmt"
	"github.com/tealeg/xlsx"
)

type KubeExcel struct {
	
}

type Excel struct {
	Name 			string
	Version 		string
	Status 			string
	Path 			string
	Method 			string
	StatusContent 	string
	Case	 		string
}

func NewKubeExcel(path string, sheet string) *KubeExcel {

    file := xlsx.NewFile()
    file.AddSheet(sheet)
    err := file.Save(path)
    if err != nil {
        fmt.Printf(err.Error())
    }
	return &KubeExcel{}
}

func (kubeExcel *KubeExcel) Write(path string, sheet string, excel *Excel) {

	file , err1 := xlsx.OpenFile(path)
	if err1 != nil {
		fmt.Println("open error:",err1)
	}
	sheetWrite := file.Sheet[sheet]
	row := sheetWrite.AddRow()
	row.WriteStruct(excel, -1)
	file.Save(path)
}

func (kubeExcel *KubeExcel) SetFill(path, sheet, patternType, fgColor, bgColor string, num int) {
	file , err1 := xlsx.OpenFile(path)
	if err1 != nil {
		fmt.Println("open error:",err1)
	}
	sheetWrite := file.Sheet[sheet]
	row := sheetWrite.Rows[sheetWrite.MaxRow - 1]
	style := xlsx.NewStyle()
	fill := *xlsx.NewFill(patternType, fgColor, bgColor)
	style.Fill = fill
    cell := row.Cells[num]
    cell.SetStyle(style)
    file.Save(path)
}

// func Test() {
// 	excel := Excel {
// 		Status: "test",
// 	}
//     file := xlsx.NewFile()
//     sheet ,_ := file.AddSheet("sheet1")
//     row := sheet.AddRow()
//     style := xlsx.NewStyle()
// 	fill := *xlsx.NewFill("solid", "00FF0000", "FF000000")
// 	style.Fill = fill
//     cell := row.Cells[2]
//     cell.SetStyle(style)
//     file.Save("/home/paas/test.xlsx")
// }