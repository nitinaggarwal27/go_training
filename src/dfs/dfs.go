package dfs

import (
	"fmt"
)

func SearchUtil(node[][] int, start int, visited[] bool){
	visited[start] = true
	fmt.Printf("%d ",start)
	for _,i:=range node[start]{
		if visited[i]==false{
			SearchUtil(node,i,visited)
		}
	}
}

func Search(node[][] int, start int){
	v:=len(node)	
	visited:=make([]bool, v)
	for i:=0;i<v;i++ {
		visited[i]=false
	}
	SearchUtil(node,start,visited)
	fmt.Println("")
}
