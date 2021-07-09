package main

import (
    "fmt"
    "html/template"
    "net/http"
    "os"
    "path/filepath"
    "strconv"
    "strings"
)

type Horizontal struct {
    horizontal [9]Cell
}

type Vertical struct {
    vertical [9]Cell
}

type Sudoku struct {
    flag int
    sudoku [9]Cell
}

type Cell struct {
    flag int
    x, y int
    v int
    temp set
}

type set map[int]struct{}

func main() {
    http.HandleFunc("/submit", submit)
    http.HandleFunc("/", index)
    http.ListenAndServe(":8080", nil)
    /*cells := &[81]Cell{
        {flag:1,x:0,y:0,v:0},{flag:1,x:1,y:0,v:0},{flag:1,x:2,y:0,v:4},  {flag:2,x:3,y:0,v:0},{flag:2,x:4,y:0,v:0},{flag:2,x:5,y:0,v:8},  {flag:3,x:6,y:0,v:0},{flag:3,x:7,y:0,v:2},{flag:3,x:8,y:0,v:0},
        {flag:1,x:0,y:1,v:9},{flag:1,x:1,y:1,v:0},{flag:1,x:2,y:1,v:0},  {flag:2,x:3,y:1,v:5},{flag:2,x:4,y:1,v:0},{flag:2,x:5,y:1,v:2},  {flag:3,x:6,y:1,v:7},{flag:3,x:7,y:1,v:0},{flag:3,x:8,y:1,v:4},
        {flag:1,x:0,y:2,v:2},{flag:1,x:1,y:2,v:0},{flag:1,x:2,y:2,v:1},  {flag:2,x:3,y:2,v:0},{flag:2,x:4,y:2,v:4},{flag:2,x:5,y:2,v:0},  {flag:3,x:6,y:2,v:0},{flag:3,x:7,y:2,v:0},{flag:3,x:8,y:2,v:5},

        {flag:4,x:0,y:3,v:5},{flag:4,x:1,y:3,v:4},{flag:4,x:2,y:3,v:0},  {flag:5,x:3,y:3,v:0},{flag:5,x:4,y:3,v:2},{flag:5,x:5,y:3,v:0},  {flag:6,x:6,y:3,v:0},{flag:6,x:7,y:3,v:3},{flag:6,x:8,y:3,v:0},
        {flag:4,x:0,y:4,v:8},{flag:4,x:1,y:4,v:0},{flag:4,x:2,y:4,v:0},  {flag:5,x:3,y:4,v:0},{flag:5,x:4,y:4,v:6},{flag:5,x:5,y:4,v:5},  {flag:6,x:6,y:4,v:0},{flag:6,x:7,y:4,v:0},{flag:6,x:8,y:4,v:9},
        {flag:4,x:0,y:5,v:0},{flag:4,x:1,y:5,v:6},{flag:4,x:2,y:5,v:0},  {flag:5,x:3,y:5,v:0},{flag:5,x:4,y:5,v:7},{flag:5,x:5,y:5,v:0},  {flag:6,x:6,y:5,v:0},{flag:6,x:7,y:5,v:8},{flag:6,x:8,y:5,v:1},

        {flag:7,x:0,y:6,v:4},{flag:7,x:1,y:6,v:0},{flag:7,x:2,y:6,v:0},  {flag:8,x:3,y:6,v:0},{flag:8,x:4,y:6,v:9},{flag:8,x:5,y:6,v:0},  {flag:9,x:6,y:6,v:3},{flag:9,x:7,y:6,v:0},{flag:9,x:8,y:6,v:8},
        {flag:7,x:0,y:7,v:1},{flag:7,x:1,y:7,v:0},{flag:7,x:2,y:7,v:8},  {flag:8,x:3,y:7,v:2},{flag:8,x:4,y:7,v:0},{flag:8,x:5,y:7,v:3},  {flag:9,x:6,y:7,v:0},{flag:9,x:7,y:7,v:0},{flag:9,x:8,y:7,v:6},
        {flag:7,x:0,y:8,v:0},{flag:7,x:1,y:8,v:9},{flag:7,x:2,y:8,v:0},  {flag:8,x:3,y:8,v:7},{flag:8,x:4,y:8,v:0},{flag:8,x:5,y:8,v:0},  {flag:9,x:6,y:8,v:1},{flag:9,x:7,y:8,v:0},{flag:9,x:8,y:8,v:0},
    }*/
    /*signPossibility4Cell(cells)
    for {
        single(cells)
        pickSingle(cells)
        continue_ := false
        for _, cell := range cells {
            if len(cell.temp) == 1 {
                continue_ = true
            }
        }
        if !continue_ {
            break
        }
    }
    if !checkResult(*cells) {
        ch := make(chan [81]Cell)
        for _, cell := range cells {
            if len(cell.temp) != 0 {
                for i, _ := range cell.temp {
                    go func(i int) {
                        c := *cells
                          c[cell.x + cell.y * 9].v = i
                          c[cell.x + cell.y * 9].temp = set{}
                          signPossibility4Cell(&c)
                          for {
                              single(&c)
                              pickSingle(&c)
                              continue_ := false
                              for _, ce := range c {
                                  if len(ce.temp) == 1 {
                                      continue_ = true
                                  }
                              }
                              if !continue_ {
                                  break
                              }
                          }
                          if checkResult(c) {
                              ch <- c
                              return
                          }
                    }(i)
                }
            }
        }
        select {
        case c := <-ch:
            printCell(c)
            close(ch)
        }
    } else {
        printCell(*cells)
    }*/
}

func checkResult(cells [81]Cell) bool {
    for _, horizontal := range findHorizontalList(&cells) {
        if !checkIn9(horizontal.horizontal) {
            return false
        }
    }
    for _, vertical := range findVerticalList(&cells) {
        if !checkIn9(vertical.vertical) {
            return false
        }

    }
    for _, sudoku := range findSudokuList(&cells) {
        if !checkIn9(sudoku.sudoku) {
            return false
        }
    }
    return true
}

func checkIn9(cells [9]Cell) bool {
    s := set{}
    for _, c := range cells {
        if c.v == 0 {
            return false
        }
        s[c.v] = struct{}{}
    }
    if len(s) != 9 {
        return false
    }
    return true
}

func getFirstKeyFromSet(set set) int {
    for i := range set {
        return i
    }
    return 0
}

func deleteAndReset(temp int, cell Cell, cells *[81]Cell) {
    cells[cell.x + cell.y * 9].v = temp
    cells[cell.x + cell.y * 9].temp = set{}
    horizontal := findHorizontal(cell, cells)
    for _, c := range horizontal.horizontal {
        if len(c.temp) == 0 {
            continue
        }
        if _, exist := c.temp[temp]; !exist {
            continue
        }
        delete(cells[c.x + c.y * 9].temp, temp)
    }
    vertical := findVertical(cell, cells)
    for _, c := range vertical.vertical {
        if len(c.temp) == 0 {
            continue
        }
        if _, exist := c.temp[temp]; !exist {
            continue
        }
        delete(cells[c.x + c.y * 9].temp, temp)
    }
    sudoku := findSudoku(cell, cells)
    for _, c := range sudoku.sudoku {
        if len(c.temp) == 0 {
            continue
        }
        if _, exist := c.temp[temp]; !exist {
            continue
        }
        delete(cells[c.x + c.y * 9].temp, temp)
    }
}

func pi(cells [9]Cell) map[int]int {
    p := map[int]int{}
    for _, c := range cells {
        if len(c.temp) == 0 {
            continue
        }
        for i, _ := range c.temp {
            if _, exist := p[i]; exist {
                p[i] = p[i] + 1
            } else {
                p[i] = 1
            }
        }
    }
    return p
}

func pick(cells [9]Cell, cellss *[81]Cell) {
    p := pi(cells)
    //fmt.Printf("p: %v\n", p)
    for i, n := range p {
        if n == 1 {
            for _, c := range cells {
                if _, exist := c.temp[i]; exist {
                    deleteAndReset(i, c, cellss)
                }
            }
        }
    }
}

func pickSingle(cells *[81]Cell) {
    for {
        //fmt.Print("start\n")
        for _, horizontal := range findHorizontalList(cells) {
            pick(horizontal.horizontal, cells)
        }
        for _, vertical := range findVerticalList(cells) {
            pick(vertical.vertical, cells)
        }
        for _, sudoku := range findSudokuList(cells) {
            pick(sudoku.sudoku, cells)
        }
        continue_ := false
        for _, horizontal := range findHorizontalList(cells) {
            p := pi(horizontal.horizontal)
            for _, v := range p {
                if v == 1 {
                    continue_ = true
                }
            }
        }
        for _, vertical := range findVerticalList(cells) {
            p := pi(vertical.vertical)
            for _, v := range p {
                if v == 1 {
                    continue_ = true
                }
            }
        }
        for _, sudoku := range findSudokuList(cells) {
            p := pi(sudoku.sudoku)
            for _, v := range p {
                if v == 1 {
                    continue_ = true
                }
            }
        }
        if !continue_ {
            return
        }
    }
}

func single(cells *[81]Cell) {
    //for {
        for i, cell := range cells {
            if len(cell.temp) == 1 {
                cells[i].v = getFirstKeyFromSet(cell.temp)
                //cells[i].temp = set{}
                /*horizontal := findHorizontal(cell, cells)
                for j, c := range horizontal.horizontal {
                    if c.temp != nil && len(c.temp) != 0 {
                        delete(horizontal.horizontal[j].temp, cells[i].v)
                    }
                }
                vertical := findVertical(cell, cells)
                for j, c := range vertical.vertical {
                    if c.temp != nil && len(c.temp) != 0 {
                        delete(vertical.vertical[j].temp, cells[i].v)
                    }
                }
                sudoku := findSudoku(cell, cells)
                for j, c := range sudoku.sudoku {
                    if c.temp != nil && len(c.temp) != 0 {
                        delete(sudoku.sudoku[j].temp, cells[i].v)
                    }
                }*/
                deleteAndReset(cells[i].v, cell, cells)
            }
        }
        /*continue_ := false
        for _, cell := range cells {
            if len(cell.temp) == 1 {
                continue_ = true
            }
        }
        if !continue_ {
            return
        }*/
    //}
}

func signPossibility4Cell(cells *[81]Cell) {
    for i, c := range cells {
        if c.v != 0 {
            continue
        }
        s := set{}
        horizontal := findHorizontal(c, cells)
        for _, cell := range horizontal.horizontal {
            if cell.v == 0 {
                continue
            }
            s[cell.v] = struct{}{}
        }
        vertical := findVertical(c, cells)
        for _, cell := range vertical.vertical {
            if cell.v == 0 {
                continue
            }
            s[cell.v] = struct{}{}
        }
        sudoku := findSudoku(c, cells)
        for _, cell := range sudoku.sudoku {
            if cell.v == 0 {
                continue
            }
            s[cell.v] = struct{}{}
        }
        v := set{}
        for j := 1; j <= 9; j++ {
            if _, exists := s[j]; exists {
                continue
            }
            v[j] = struct{}{}
        }
        cells[i].temp = v
    }
}

func findHorizontal(cell Cell, cells *[81]Cell) Horizontal {
    var horizontal [9]Cell
    for _, c := range cells {
        if c.y == cell.y {
            horizontal[c.x] = c
        }
    }
    return Horizontal{horizontal: horizontal}
}

func findVertical(cell Cell, cells *[81]Cell) Vertical {
    var vertical [9]Cell
    for _, c := range cells {
        if c.x == cell.x {
            vertical[c.y] = c
        }
    }
    return Vertical{vertical: vertical}
}

func findSudoku(cell Cell, cells *[81]Cell) Sudoku {
    var sudoku [9]Cell
    j := 0
    for _, c := range cells {
        if c.flag == cell.flag {
            sudoku[j] = c
            j++
        }
    }
    return Sudoku{
        flag:   cell.flag,
        sudoku: sudoku,
    }
}

func findHorizontalList(cells *[81]Cell) *[9]Horizontal {
    var horizontal [9]Cell
    var horizontalList *[9]Horizontal
    horizontalList = new([9]Horizontal)
    for i := 0; i < 9; i++ {
        for _, c := range cells {
            if c.y == i {
                horizontal[c.x] = c
            }
        }
        horizontalList[i] = Horizontal{horizontal: horizontal}
    }
    return horizontalList
}

func findVerticalList(cells *[81]Cell) *[9]Vertical {
    var vertical [9]Cell
    var verticalList *[9]Vertical
    verticalList = new([9]Vertical)
    for i := 0; i < 9; i++ {
        for _, c := range cells {
            if c.x == i {
                vertical[c.y] = c
            }
        }
        verticalList[i] = Vertical{vertical: vertical}
    }
    return verticalList
}

func findSudokuList(cells *[81]Cell) *[9]Sudoku {
    var sudoku [9]Cell
    var sudokuList *[9]Sudoku
    sudokuList = new([9]Sudoku)
    for i := 1; i <= 9; i++ {
        j := 0
        for _, c := range cells {
            if c.flag == i {
                sudoku[j] = c
                j++
            }
        }
        sudokuList[i-1] = Sudoku{
            flag:   i,
            sudoku: sudoku,
        }
    }

    return sudokuList
}

func printCell(cells [81]Cell) {
    for _, cell := range cells {
        if cell.x < 8 {
            fmt.Print(cell)
        } else {
            fmt.Println(cell)
        }
    }
}

func pushCellValue(cells [81]Cell) string {
    s := ""
    for i, cell := range cells {
        if i != 80 {
            s = s + strconv.Itoa(cell.v) + ","
        } else {
            s = s + strconv.Itoa(cell.v)
        }
    }
    return s
}

func submit(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    shuDu := r.FormValue("shuDuArray")
    shuDuArray := strings.Split(shuDu, ",")
    var cells [81]Cell
    var s string
    for i := 0; i < 81; i++ {
        cells[i] = setCell(i, shuDuArray)
    }
    signPossibility4Cell(&cells)
    for {
        single(&cells)
        pickSingle(&cells)
        continue_ := false
        for _, cell := range cells {
            if len(cell.temp) == 1 {
                continue_ = true
            }
        }
        if !continue_ {
            break
        }
    }
    if !checkResult(cells) {
        //ch := make(chan [81]Cell)
        for _, cell := range cells {
            success := false
            if len(cell.temp) != 0 {
                for i, _ := range cell.temp {
                    fmt.Println(i)
                    c := cells
                    c[cell.x + cell.y * 9].v = i
                    c[cell.x + cell.y * 9].temp = set{}
                    signPossibility4Cell(&c)
                    for {
                        single(&c)
                        pickSingle(&c)
                        continue_ := false
                        for _, ce := range c {
                            if len(ce.temp) == 1 {
                                continue_ = true
                            }
                        }
                        if !continue_ {
                            break
                        }
                    }
                    if checkResult(c) {
                        //ch <- c
                        success = true
                        s = pushCellValue(c)
                        break
                    }
                }
            }
            if success {
                break
            }
        }
        /*select {
        case c := <-ch:
            s = pushCellValue(c)
            close(ch)
        }*/
    } else {
        s = pushCellValue(cells)
    }
    wd, _ := os.Getwd()
    t, _ := template.ParseFiles(filepath.Join(wd, "./shudu.gtpl"))
    t.Execute(w, s)
}

func index(w http.ResponseWriter, r *http.Request) {
    wd, _ := os.Getwd()
    t, _ := template.ParseFiles(filepath.Join(wd, "./shudu.gtpl"))
    t.Execute(w, nil)
}

func setCell(i int, array []string) Cell {
    chu := i / 9
    yu := i % 9
    flag := 0
    x, y, v := -1, -1, 0
    x = yu
    y = chu
    v, _ = strconv.Atoi(array[i])
    if chu >= 0 && chu <= 2 {
        switch {
        case yu >= 0 && yu <= 2:
            flag = 1
        case yu >= 3 && yu <= 5:
            flag = 2
        case yu >= 6 && yu <= 8:
            flag = 3
        }
    }
    if chu >= 3 && chu <= 5 {
        switch {
        case yu >= 0 && yu <= 2:
            flag = 4
        case yu >= 3 && yu <= 5:
            flag = 5
        case yu >= 6 && yu <= 8:
            flag = 6
        }
    }
    if chu >= 6 && chu <= 8 {
        switch {
        case yu >= 0 && yu <= 2:
            flag = 7
        case yu >= 3 && yu <= 5:
            flag = 8
        case yu >= 6 && yu <= 8:
            flag = 9
        }
    }
    return Cell{
        flag: flag,
        x:    x,
        y:    y,
        v:    v,
        temp: nil,
    }
}