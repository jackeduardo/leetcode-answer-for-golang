package main


type Node struct {
	Val int
	Left *Node
	Right *Node
	Next *Node
}

//常数空间dfs
func connect(root *Node) *Node {
	if root==nil||root.Left==nil{
		return root
	}
	root.Left.Next=root.Right
	if root.Next!=nil{
		root.Right.Next=root.Next.Left
	}
	connect(root.Left)
	connect(root.Right)
	return root

}
//bfs层序遍历
func connect2(root *Node) *Node {
	if root==nil{
		return nil
	}
	queue:=make([]*Node,0)
	queue=append(queue,root)
	for len(queue)!=0{
		length:=len(queue)
		for i := 0; i <length; i++ {
			cur:=queue[0]
			queue=queue[1:]
			if i<length-1{
				cur.Next=queue[0]
			}
			if cur.Left!=nil{
				queue=append(queue,cur.Left)
			}
			if cur.Right!=nil{
				queue=append(queue,cur.Right)
			}
		}
	}
	return root
}