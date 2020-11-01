

func preorder(root *Node) []int {
	var res []int
	var stack = []*Node{root}
	for 0 < len(stack) {
		for root != nil {
			res = append(res, root.Val) //Ç°ÐòÊä³ö
			if 0 == len(root.Children) {
				break
			}
			for i := len(root.Children) - 1; 0 < i; i-- {
				stack = append(stack, root.Children[i]) //ÈëÕ»
			}
			root = root.Children[0]
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return res
}

