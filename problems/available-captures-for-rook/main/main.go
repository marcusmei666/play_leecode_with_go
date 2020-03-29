package main

//
func numRookCaptures(board [][]byte) int {

    for i:= 0; i < 8; i++ {
        for j:= 0; j < 8; j++ {
            //先找到车的位置
            if board[i][j] == 'R' {
                return getNums(i,j,board)
            }
        }
    }
    return 0
}

func getNums(i,j int ,board [][]byte) int {
    res := 0
    for x := i; x>0; x-- {
        if board[x][j] == 'p' {
            res ++
            break
        }else if board[x][j] == 'B' {
            break
        }
    }

    for x := i; x<8; x++ {
        if board[x][j] == 'p' {
            res ++
            break
        }else if board[x][j] == 'B' {
            break
        }
    }

    for x := j; x>0; x-- {
        if board[i][x] == 'p' {
            res ++
            break
        }else if board[i][x] == 'B' {
            break
        }
    }

    for x := j; x<8; x++ {
        if board[i][x] == 'p' {
            res ++
            break
        }else if board[i][x] == 'B' {
            break
        }
    }
    return res
}