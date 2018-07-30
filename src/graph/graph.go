package main

import (
	"fmt"
	"dfs"
	"bfs"
)

func addEdge(node[][] int, a,b int){
	node[a]=append(node[a],b)
	node[b]=append(node[b],a)
	
}
func main(){
	var a,b,v,e int
	fmt.Printf("No of nodes: ");
	fmt.Scanf("%d",&v)
	fmt.Printf("No of edges: ");
	fmt.Scanf("%d",&e)
	node:=make([][]int,v);
	fmt.Println("Enter edges");
	for i:=0;i<e;i++{
		fmt.Scanf("%d %d",&a,&b)
		addEdge(node,a,b)
	}
	fmt.Printf("Depth First Search: ")
	dfs.Search(node,0);
	fmt.Printf("Breadth First Search: ")
	bfs.Search(node,0);
}
