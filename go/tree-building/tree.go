package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children Children
}

type Children []*Node

func (children Children) Len() int {
	return len(children)
}

func (children Children) Less(i, j int) bool {
	return children[i].ID < children[j].ID
}

func (children Children) Swap(i, j int) {
	children[i], children[j] = children[j], children[i]
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	nodes := make([]Node, len(records))
	parents := make(map[int]int)
	seen := make(map[int]bool)

	for _, record := range records {
		err := validateRecord(record, &records, &seen)
		if err != nil {
			return nil, err
		}

		seen[record.ID] = true

		if record.ID != record.Parent {
			parents[record.ID] = record.Parent
		}
	}

	for i := 1; i < len(nodes); i++ {
		nodes[i].ID = i
		nodes[parents[i]].Children = append(nodes[parents[i]].Children, &nodes[i])
	}

	for _, node := range nodes {
		sort.Sort(node.Children)
	}

	return &nodes[0], nil
}

func validateRecord(record Record, records *[]Record, seen *map[int]bool) error {
	if record.ID == 0 && record.Parent != 0 {
		return errors.New("Node 0 cannot have parent != 0")
	}

	if record.ID >= len(*records) {
		return errors.New("Node ID cannot be >= number of records")
	}

	if record.Parent > record.ID {
		return errors.New("Node cannot have parent of higher ID")
	}

	if (*seen)[record.ID] {
		return errors.New("A Node cannot appear in tree twice")
	}

	if record.ID != 0 && record.ID == record.Parent {
		return errors.New("Only the root node can be its own parent")
	}

	return nil
}
