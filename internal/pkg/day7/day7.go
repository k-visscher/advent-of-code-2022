package day7

import (
	"advent-of-code-2022/internal/pkg/utils"
	"regexp"
	"sort"
	"strings"
)

/*
--- Day 7: No Space Left On Device ---

You can hear birds chirping and raindrops hitting leaves as the expedition proceeds. Occasionally, you can even hear much louder sounds in the distance; how big do the animals get out here, anyway?

The device the Elves gave you has problems with more than just its communication system. You try to run a system update:

$ system-update --please --pretty-please-with-sugar-on-top
Error: No space left on device

Perhaps you can delete some files to make space for the update?

You browse around the filesystem to assess the situation and save the resulting terminal output (your puzzle input). For example:

$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k

The filesystem consists of a tree of files (plain data) and directories (which can contain other directories or files). The outermost directory is called /. You can navigate around the filesystem, moving into or out of directories and listing the contents of the directory you're currently in.

Within the terminal output, lines that begin with $ are commands you executed, very much like some modern computers:

    cd means change directory. This changes which directory is the current directory, but the specific result depends on the argument:
        cd x moves in one level: it looks in the current directory for the directory named x and makes it the current directory.
        cd .. moves out one level: it finds the directory that contains the current directory, then makes that directory the current directory.
        cd / switches the current directory to the outermost directory, /.
    ls means list. It prints out all of the files and directories immediately contained by the current directory:
        123 abc means that the current directory contains a file named abc with size 123.
        dir xyz means that the current directory contains a directory named xyz.

Given the commands and output in the example above, you can determine that the filesystem looks visually like this:

- / (dir)
  - a (dir)
    - e (dir)
      - i (file, size=584)
    - f (file, size=29116)
    - g (file, size=2557)
    - h.lst (file, size=62596)
  - b.txt (file, size=14848514)
  - c.dat (file, size=8504156)
  - d (dir)
    - j (file, size=4060174)
    - d.log (file, size=8033020)
    - d.ext (file, size=5626152)
    - k (file, size=7214296)

Here, there are four directories: / (the outermost directory), a and d (which are in /), and e (which is in a). These directories also contain files of various sizes.

Since the disk is full, your first step should probably be to find directories that are good candidates for deletion. To do this, you need to determine the total size of each directory. The total size of a directory is the sum of the sizes of the files it contains, directly or indirectly. (Directories themselves do not count as having any intrinsic size.)

The total sizes of the directories above can be found as follows:

    The total size of directory e is 584 because it contains a single file i of size 584 and no other directories.
    The directory a has total size 94853 because it contains files f (size 29116), g (size 2557), and h.lst (size 62596), plus file i indirectly (a contains e which contains i).
    Directory d has total size 24933642.
    As the outermost directory, / contains every file. Its total size is 48381165, the sum of the size of every file.

To begin, find all of the directories with a total size of at most 100000, then calculate the sum of their total sizes. In the example above, these directories are a and e; the sum of their total sizes is 95437 (94853 + 584). (As in this example, this process can count files more than once!)

Find all of the directories with a total size of at most 100000. What is the sum of the total sizes of those directories?
*/

type NodeType int

const (
	File NodeType = iota
	Directory
)

type Node struct {
	Name      string
	NodeType  NodeType
	Parent    *Node   // parent if present is always a Node of nodeType == Directory
	Childeren []*Node // only set if nodeType == Directory
	Size      int     // only set if nodeType == File
}

func (n *Node) TotalSize() int {
	if n.NodeType == File {
		return n.Size
	}

	var result = 0
	for _, node := range n.Childeren {
		result += node.TotalSize()
	}
	return result
}

var cum_sum = 0

func (n *Node) CumSum() {
	if n.NodeType == File {
		return
	}

	for _, node := range n.Childeren {
		node.CumSum()
	}

	totalSize := n.TotalSize()
	if totalSize < 100_000 {
		cum_sum += totalSize
	}
}

var could_be_freed = make([]*Node, 0)

func (n *Node) FindDirsToFree(to_be_freed int) {
	if n.NodeType == File {
		return
	}

	for _, node := range n.Childeren {
		node.FindDirsToFree(to_be_freed)
	}

	total_size := n.TotalSize()
	if total_size >= to_be_freed {
		could_be_freed = append(could_be_freed, n)
	}
}

func StarOne(input_path string) int {
	input, err := utils.ReadFileToString(input_path)
	utils.CheckForErr(err)

	lines := strings.Split(input, "\n")[2:]

	cd_cmd_expr := regexp.MustCompile(`^\$ cd (.*)$`)
	//ls_cmd_expr := regexp.MustCompile(`^\$ ls$`)

	dir_item_expr := regexp.MustCompile(`^dir (.*)$`)
	file_item_expr := regexp.MustCompile(`^(\d+) (.*)$`)

	root := &Node{
		Name:      "/",
		NodeType:  Directory,
		Parent:    nil,
		Childeren: make([]*Node, 0),
		Size:      0,
	}

	current_node := root

	for _, line := range lines {
		if dir_item_expr.MatchString(line) {
			groups := dir_item_expr.FindStringSubmatch(line)[1:]

			node := &Node{
				Name:      groups[0],
				NodeType:  Directory,
				Parent:    current_node,
				Childeren: make([]*Node, 0),
				Size:      0,
			}
			current_node.Childeren = append(current_node.Childeren, node)
		} else if file_item_expr.MatchString(line) {
			groups := file_item_expr.FindStringSubmatch(line)[1:]

			node := &Node{
				Name:      groups[1],
				NodeType:  File,
				Parent:    current_node,
				Childeren: nil,
				Size:      utils.MustParseAsInt(groups[0]),
			}
			current_node.Childeren = append(current_node.Childeren, node)
		} else if cd_cmd_expr.MatchString(line) {
			groups := cd_cmd_expr.FindStringSubmatch(line)[1:]

			if groups[0] == ".." {
				current_node = current_node.Parent
			} else {
				for _, child := range current_node.Childeren {
					if child.Name == groups[0] {
						current_node = child
						break
					}
				}
			}
		}
	}

	for _, child := range root.Childeren {
		child.CumSum()
	}

	return cum_sum
}

/*
--- Part Two ---

Now, you're ready to choose a directory to delete.

The total disk space available to the filesystem is 70000000. To run the update, you need unused space of at least 30000000. You need to find a directory you can delete that will free up enough space to run the update.

In the example above, the total size of the outermost directory (and thus the total amount of used space) is 48381165; this means that the size of the unused space must currently be 21618835, which isn't quite the 30000000 required by the update. Therefore, the update still requires a directory with total size of at least 8381165 to be deleted before it can run.

To achieve this, you have the following options:

	Delete directory e, which would increase unused space by 584.
	Delete directory a, which would increase unused space by 94853.
	Delete directory d, which would increase unused space by 24933642.
	Delete directory /, which would increase unused space by 48381165.

Directories e and a are both too small; deleting them would not free up enough space. However, directories d and / are both big enough! Between these, choose the smallest: d, increasing unused space by 24933642.

Find the smallest directory that, if deleted, would free up enough space on the filesystem to run the update. What is the total size of that directory?
*/
func StarTwo(input_path string) int {
	input, err := utils.ReadFileToString(input_path)
	utils.CheckForErr(err)

	lines := strings.Split(input, "\n")[2:]

	cd_cmd_expr := regexp.MustCompile(`^\$ cd (.*)$`)
	//ls_cmd_expr := regexp.MustCompile(`^\$ ls$`)

	dir_item_expr := regexp.MustCompile(`^dir (.*)$`)
	file_item_expr := regexp.MustCompile(`^(\d+) (.*)$`)

	root := &Node{
		Name:      "/",
		NodeType:  Directory,
		Parent:    nil,
		Childeren: make([]*Node, 0),
		Size:      0,
	}

	current_node := root

	for _, line := range lines {
		if dir_item_expr.MatchString(line) {
			groups := dir_item_expr.FindStringSubmatch(line)[1:]

			node := &Node{
				Name:      groups[0],
				NodeType:  Directory,
				Parent:    current_node,
				Childeren: make([]*Node, 0),
				Size:      0,
			}
			current_node.Childeren = append(current_node.Childeren, node)
		} else if file_item_expr.MatchString(line) {
			groups := file_item_expr.FindStringSubmatch(line)[1:]

			node := &Node{
				Name:      groups[1],
				NodeType:  File,
				Parent:    current_node,
				Childeren: nil,
				Size:      utils.MustParseAsInt(groups[0]),
			}
			current_node.Childeren = append(current_node.Childeren, node)
		} else if cd_cmd_expr.MatchString(line) {
			groups := cd_cmd_expr.FindStringSubmatch(line)[1:]

			if groups[0] == ".." {
				current_node = current_node.Parent
			} else {
				for _, child := range current_node.Childeren {
					if child.Name == groups[0] {
						current_node = child
						break
					}
				}
			}
		}
	}

	total_memory := 70_000_000
	update_size := 30_000_000

	used_memory := root.TotalSize()
	free_memory := total_memory - used_memory

	to_be_freed := update_size - free_memory

	root.FindDirsToFree(to_be_freed)

	sort.Slice(could_be_freed, func(i, j int) bool {
		return could_be_freed[i].TotalSize() < could_be_freed[j].TotalSize()
	})

	return could_be_freed[0].TotalSize()
}
