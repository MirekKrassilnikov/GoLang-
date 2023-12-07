
package main

import (
"fmt"
"math/rand"
)

func main() {

//SLICE CREATE
var columns int
var rows int

fmt.Println("Enter rows number")
fmt.Scan(&rows)

fmt.Println("Enter columns number")
fmt.Scan(&columns)

multiSlice := make([][]int, rows)

for i, _ := range multiSlice {
multiSlice[i] = make([]int, columns)

}
//FILL WITH RANDOM NUMBERS
for i, _ := range multiSlice {
for l, _ := range multiSlice[i] {
multiSlice[i][l] = rand.Intn(199) - 99

}
}

for _, value := range multiSlice {
fmt.Println(value)
}
maxNumber(multiSlice)
minNumber(multiSlice)
allRowsSummarize(multiSlice)
allColumnsSummarize(multiSlice)
sortedRows(multiSlice)
sortedColumns(multiSlice, rows)

}

func maxNumber(slice [][]int) int {
maxNumberRow := slice[0][0]
for k, _ := range slice {
for _, val := range slice[k] {
if val > maxNumberRow {
maxNumberRow = val

}

}
}
fmt.Println("Max number in the slice:", maxNumberRow)
return maxNumberRow
}

func minNumber(slice [][]int) int {
minNumberRow := slice[0][0]
for k, _ := range slice {
for _, val := range slice[k] {
if val < minNumberRow {
minNumberRow = val

}

}
}

fmt.Println("Min number in the slice:", minNumberRow)
return minNumberRow
}

func allRowsSummarize(slice [][]int) int {
sum := 0
for i, value := range slice {
sum = 0
for _, val := range value {
sum = sum + val

}
fmt.Println("Summary of the row num.", fmt.Sprint(i+1), "is", fmt.Sprint(sum))
}
return sum
}

func allColumnsSummarize(slice [][]int) int {
sum := 0
for i, value := range slice {
sum = 0
for n, _ := range value {
sum = sum + slice[n][i]
}
fmt.Println("Summary of the column num.", fmt.Sprint(i+1), "is", fmt.Sprint(sum))
}
return sum
}

func sortedRows(slice [][]int) []int {
sortedSlice := []int{}
for _, value := range slice {
sortedSlice = []int{}
for len(value) > 0 {
lowestNumber := value[0]
lowestIndex := 0

for i, val := range value {
if val < lowestNumber {
lowestNumber = val
lowestIndex = i
}
}

sortedSlice = append(sortedSlice, lowestNumber)

value = append(value[:lowestIndex], value[lowestIndex+1:]...)
}

fmt.Println(sortedSlice)

}
return sortedSlice
}

func sortedColumns(slice [][]int, rows int) {
for k := 0; k < rows; k++ {

for f := 0; f < rows; f++ {
for h := 1; h < rows; h++ {
if slice[h][k] < slice[h-1][k] {
slice[h-1][k], slice[h][k] = slice[h][k], slice[h-1][k]

if h == rows-1 {
break
}
}
}
}

}
fmt.Println("Columns sorted low to high")
for _, value := range slice {
fmt.Println(value)
}

}

/*

   func main() {
   >>>>>>> 5915bff (Разбито на функции)

   	//SLICE CREATE
   	var columns int
   	var rows int

   	fmt.Println("Enter rows number")
   	fmt.Scan(&rows)

   	fmt.Println("Enter columns number")
   	fmt.Scan(&columns)

   	multiSlice := make([][]int, rows)

   	for i, _ := range multiSlice {
   		multiSlice[i] = make([]int, columns)

   	}
   	//FILL WITH RANDOM NUMBERS
   	for i, _ := range multiSlice {
   		for l, _ := range multiSlice[i] {
   			multiSlice[i][l] = rand.Intn(199) - 99

   		}
   	}

   	for _, value := range multiSlice {
   		fmt.Println(value)
   	}
   	//MAX NUMBER
   	maxNumberRow := multiSlice[0][0]
   	for k, _ := range multiSlice {
   		for _, val := range multiSlice[k] {
   			if val > maxNumberRow {
   				maxNumberRow = val

   			}

   		}
   	}
   	fmt.Println("Highest number is ", fmt.Sprint(maxNumberRow))
   	//MIN NUMBER
   	minNumberRow := multiSlice[0][0]
   	for k, _ := range multiSlice {
   		for _, val := range multiSlice[k] {
   			if val < minNumberRow {
   				minNumberRow = val

   			}

   		}
   	}
   	fmt.Println("Lowest number is ", fmt.Sprint(minNumberRow))
   	//ALL ROWS SUMARIZE
   	for i, value := range multiSlice {
   		sum := 0
   		for _, val := range value {
   			sum = sum + val
   		}
   		fmt.Println("Summary of the row num.", fmt.Sprint(i+1), "is", fmt.Sprint(sum))
   	}

   	//ALL COLUMNS SUMMARIZE
   	for i, value := range multiSlice {
   		sum := 0
   		for n, _ := range value {
   			sum = sum + multiSlice[n][i]
   		}
   		fmt.Println("Summary of the column num.", fmt.Sprint(i+1), "is", fmt.Sprint(sum))
   	}
   	//EXACT ROW SUMMARIZE
   	fmt.Println("Enter the row you want to summarize")

   	var rowNum int
   	var sum int
   	fmt.Scan(&rowNum)
   	for _, value := range multiSlice[rowNum-1] {
   		sum = sum + value
   	}
   	fmt.Printf("Summary of the row num. %d is %d", rowNum, sum)

   	//EXACT COLUMN SUMMARIZE
   	fmt.Println("\n Enter the column you want to summarize")

   	var colNum int
   	var colSum int
   	fmt.Scan(&colNum)
   	for i, _ := range multiSlice {
   		colSum = colSum + multiSlice[i][colNum-1]
   	}
   	fmt.Printf("Summary of the col num. %d is %d", rowNum, colSum)

   	//SORT LOW TO HIGH ROWS

   	for _, value := range multiSlice {
   		sortedSlice := []int{}
   		for len(value) > 0 {
   			lowestNumber := value[0]
   			lowestIndex := 0

   			for i, val := range value {
   				if val < lowestNumber {
   					lowestNumber = val
   					lowestIndex = i
   				}
   			}

   			sortedSlice = append(sortedSlice, lowestNumber)

   			value = append(value[:lowestIndex], value[lowestIndex+1:]...)
   		}

   		fmt.Println(sortedSlice)

   	}

   	//SORT COLUMNS LOW TO HIGH
   	for k := 0; k < rows; k++ {

   		for f := 0; f < rows; f++ {
   			for h := 1; h < rows; h++ {
   				if multiSlice[h][k] < multiSlice[h-1][k] {
   					multiSlice[h-1][k], multiSlice[h][k] = multiSlice[h][k], multiSlice[h-1][k]

   					if h == rows-1 {
   						break
   					}
   				}
   			}
   		}

   	}
   	fmt.Println("Columns sorted low to high")
   	for _, value := range multiSlice {
   		fmt.Println(value)
   	}
   }

*/


