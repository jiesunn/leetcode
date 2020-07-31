package main;
import (
	"fmt"
    "strconv"
)
func main() {
    var m = make(map[string]int)
    var res = make(map[string]int)
    var x, y, n int
    fmt.Scanln(&x, &y, &n)
    xFlag, yFlag := true, true
    if x < 0 {
        xFlag = false
        x = 0-x
    }
    if y < 0 {
        yFlag = false
        y = 0-y
    }
    
    temp := n
    for temp > 0 {
        temp--
        var tx, ty int
        fmt.Scanln(&tx, &ty)    
        if !xFlag {
        	tx = 0-tx
    	}
    	if !yFlag {
        	ty = 0-ty
    	}
        key := strconv.Itoa(tx) + "_" + strconv.Itoa(ty)
        m[key] = 1
    }
    
    for i := 0; i <= x; i++ {
		for j := 0; j <= y; j++ {
			key := strconv.Itoa(i) + "_" + strconv.Itoa(j)
            _, ok := m[key]
			if ok {
				res[key] = -1
				continue
			}
            if i == 0 {
                res[key] = j
                continue
            }
            if j == 0 {
                res[key] = i
                continue
            }
			key1 := strconv.Itoa(i-1) + "_" + strconv.Itoa(j)
			key2 := strconv.Itoa(i) + "_" + strconv.Itoa(j-1)
            res[key] = res[key1]
            if res[key1] == -1 || res[key1] > res[key2] {
                res[key] = res[key2]
            }
		}
	}
	key := strconv.Itoa(x) + "_" + strconv.Itoa(y)
	fmt.Println(res[key])
}