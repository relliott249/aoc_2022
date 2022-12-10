package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Readln returns a single line (without the ending \n)
// from the input buffered reader.
// An error is returned iff there is an error with the
// buffered reader.
func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix       = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

type Node struct {
	Name     string
	Type     string // file or directory
	Parent   *Node  // pointer to parent directory
	Children []Node // list of pointers to subdirectories or files
	Size     int    // Size of the file directories will have size 0
}

// print the node structure
func (node *Node) printNode(level int) {
	spaces := strings.Repeat(" ", level)
	fmt.Printf("%s- %s (%s)\n", spaces, node.Name, node.Type)
	if node.Children != nil {
		level += 1
		for _, c := range node.Children {
			c.printNode(level)
		}
	}
}

// get the size of a current node
func (node *Node) getNodeSize() (size int) {
	size += node.Size
	if node.Children != nil {
		for _, i := range node.Children {
			size += i.getNodeSize()
		}
	}
	return size
}

func newFile(name string, parent *Node, size int) (*Node, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("No name specified for current Fikle")
	}
	if parent == nil {
		return nil, fmt.Errorf("No Parent node defined for file")
	}
	if size <= 0 {
		return nil, fmt.Errorf("Size must be >= 0")
	}

	return &Node{
		Name:     name,
		Type:     "file",
		Children: nil,
		Parent:   parent,
		Size:     size,
	}, nil

}

func newDir(name string, parent *Node) (*Node, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("No name specified for current Directory")
	}

	return &Node{
		Name:     name,
		Type:     "dir",
		Children: nil,
		Parent:   parent,
		Size:     0,
	}, nil
}

func (curr *Node) addChild(child Node) {
	curr.Children = append(curr.Children, child)
}

func (node *Node) getSizeUnder(limit int) (size int) {
	for _, i := range node.Children {
		if i.Type == "dir" {
			currSize := i.getNodeSize()
			if currSize <= limit {
				fmt.Println("child: ", i.Name, " size: ", currSize)
				size += currSize
			}
			size += i.getSizeUnder(limit)
		}
	}
	return size
}

func (node *Node) getSizeOver(limit int) (size int) {
	for _, i := range node.Children {
		if i.Type == "dir" {
			currSize := i.getNodeSize()
			if currSize >= limit {
				fmt.Println("child: ", i.Name, " size: ", currSize)
				size += currSize
			}
			size += i.getSizeUnder(limit)
		}
	}
	return size
}
func (node *Node) getNodesOverSize(size int) (nodes []Node) {
	if node.Type != "dir" {
		return nil
	}

	currSize := node.getNodeSize()
	if currSize >= size {
		nodes = append(nodes, *node)
	}

	for _, i := range node.Children {
		currSize = i.getNodeSize()
		if currSize >= size {
			nodes = append(nodes, i.getNodesOverSize(size)...)
		}
	}
	return nodes
}

func partTwo(node *Node) (size int) {
	diskSize := 70000000
	spaceNeeded := 30000000
	currSize := node.getNodeSize()
	unUsed := diskSize - currSize
	toDelete := spaceNeeded - unUsed

	// find the smallest node to delete
	potNodes := node.getNodesOverSize(toDelete)
	smallestNode := potNodes[0].getNodeSize()
	for _, i := range potNodes[1:] {
		if i.getNodeSize() < smallestNode {
			smallestNode = i.getNodeSize()
		}
	}
	return smallestNode

}

func main() {

	f, err := os.Open("/home/relliott/aoc_2022/day07/input_day07.txt")
	check(err)
	r := bufio.NewReader(f)

	//var dirList [10]Directory
	var aocInput []string
	var s, e = Readln(r)
	aocInput = append(aocInput, s)
	for e == nil {
		s, e = Readln(r)
		aocInput = append(aocInput, s)
	}
	f.Close()

	var currLine []string
	// first do root directory
	currLine = strings.Fields(aocInput[0])
	rootDir, _ := newDir(currLine[2], nil)
	currDir := rootDir
	for i := 1; i < len(aocInput)-1; i++ {
		currLine = strings.Fields(aocInput[i])
		// Next line is either cd or ls
		if currLine[0] == "$" {
			if currLine[1] == "cd" {
				// go up one directory
				if currLine[2] == ".." {
					currDir = currDir.Parent
				} else { // go to specified subdir
					dirName := currLine[2]
					for j, child := range currDir.Children {
						if child.Name == dirName {
							currDir = &currDir.Children[j]
						}
					}
				}
			} else { // ls
				continue
			}
		} else if currLine[0] == "dir" {
			dirName := currLine[1]
			node, err := newDir(dirName, currDir)
			check(err)
			currDir.addChild(*node)
		} else { // this is a file
			size, _ := strconv.Atoi(currLine[0])
			name := currLine[1]
			child, err := newFile(name, currDir, size)
			check(err)
			currDir.addChild(*child)
		}
	}

	result := rootDir.getSizeUnder(100000)
	result2 := partTwo(rootDir)

	fmt.Println("Part 1 result: ", result)
	fmt.Println("Part 2 result: ", result2)

}
