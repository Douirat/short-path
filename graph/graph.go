package graph

import "fmt"

// A class to represent the graph:
type SocialGraph struct {
	Size   int
	Matrix [][]int
}

// A type to represent the vertice it self:
type User struct {
	UserId   int
	UserName string
}

// Instantiate the graph:
func InstGraph(size int) *SocialGraph {
	graph := new(SocialGraph)
	graph.Size = size
	graph.Matrix = make([][]int, size)
	for i := 0; i < size; i++ {
		graph.Matrix[i] = make([]int, size)
	}
	return graph
}

// Add a new relationship:
func (grqph *SocialGraph) AddFriendship(sender, reciever int) {
	grqph.Matrix[sender-1][reciever-1] = 1
}

// Create a new user:
func NewUser(id int, name string) *User {
	user := new(User)
	user.UserId = id
	user.UserName = name
	return user
}

func (graph SocialGraph) Display() {
		fmt.Printf("  ")
	for i := 0; i < graph.Size; i++ {
		fmt.Printf("%d ", i+1)
	}
	fmt.Println()
	for j:=0; j<graph.Size; j++ {
		fmt.Printf("%d ", j)
		for z:=0; z<graph.Size; z++ {
			fmt.Printf("%d ", graph.Matrix[j][z])
		}
		fmt.Println()
	}
}
