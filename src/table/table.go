package table

import (
	"fmt"
	"gitlab.com/e.ilmenderov/philo-go/src/stRuct"
	"strconv"
)

func TableMsg(tableStruct *stRuct.TableStruct) {
	fmt.Println("This table have ", tableStruct.PhiloCount, " philosophers")
}

func ParseArgs(args []string) *stRuct.TableStruct {

	table := new(stRuct.TableStruct)
	table.PhiloCount, _ = strconv.Atoi(args[1])
	tmp, _ := strconv.Atoi(args[2])
	table.TimeDie = int64(tmp)
	tmp, _ = strconv.Atoi(args[3])
	table.TimeEat = int64(tmp)
	tmp, _ = strconv.Atoi(args[4])
	table.TimeSleep = int64(tmp)
	if len(args) == 6 {
		table.MustEat, _ = strconv.Atoi(args[5])
	}
	table.Forks = append(table.Forks, make([]bool, table.PhiloCount)...)
	table.Philos = append(table.Philos, make([]stRuct.PhiloStruct, table.PhiloCount)...)
	for i := 0; i < table.PhiloCount; i++ {
		table.Philos[i].Number = i
	}
	return table
}
