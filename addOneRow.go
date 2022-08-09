package main



func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	curr := 1
	queue := []*TreeNode{root}
	l := TreeNode{Val: val}
	r := TreeNode{Val: val}
	for len(queue) > 0 {

		size := len(queue)
		for i := 0 ; i < size ;i ++ {
			if curr == depth - 1{
				if queue[i].Left != nil{
					l.Left = queue[i].Left
					queue[i].Left = l
				}
				if queue[i].Right != nil{

				}
			}

			if queue[i].Left != nil {
				queue = append(queue,queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue,queue[i].Right)
			}
			curr ++
		}

		if curr == depth - 1 {

			lft := new TreeNode{val,}

		} else {

		}
	}
}
