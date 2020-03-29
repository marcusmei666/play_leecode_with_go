package main

func Dfs (grid [][]byte,x,y,rows,cols int){

	if(x<0 || x>=rows || y<0 || y>=cols){
		return ;
	}
	if(grid[x][y]!='1'){
		return;
	}

	grid[x][y] = '0'

	if x > 0 && grid[x-1][y] == '1' {
		Dfs(grid, x-1, y,rows,cols)
	}
	if x < len(grid)-1 && grid[x+1][y] == '1' {
		Dfs(grid, x, y-1,rows,cols)
	}
	if y > 0 && grid[x][y-1] == '1' {
		Dfs(grid, x+1, y,rows,cols)
	}
	if y < len(grid[0])-1 && grid[x][y+1] == '1' {
		Dfs(grid, x, y+1,rows,cols)
	}

}

func numIslands(grid [][]byte) int {
	cou := 0
	rows:= len(grid)
	cols:= len(grid[0])

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				cou ++
				Dfs(grid,i,j,rows,cols)
			}
		}

	}
	//for k1,v1 := range grid{
	//	cols:= len(grid[k1])
	//	for k2,v2 := range v1{
	//		if v2 == 1{
	//			cou ++
	//			Dfs(grid,k1,k2,rows,cols)
	//		}
	//	}
	//}

	return cou
}