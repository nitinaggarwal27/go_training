package bfs1

import (
	"fmt"
)

func Search(node[][] int, start int){
	v:=len(node)	
	visited:=make([]bool, v)
	for i:=0;i<v;i++ {
		visited[i]=false
	}
	queue:=make([]int,0)
	queue=append(queue,start)
	visited[start]=true
	for len(queue)!=0 {
		x:=queue[0]
		queue=queue[1:]
		fmt.Printf("%d ",x)
		for _,i:=range node[x]{
			if visited[i]==false{
				queue=append(queue,i)
				visited[i]=true
			}
		}
	}
	fmt.Println("")
}
