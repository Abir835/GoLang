package main

import (
	"errors"
	"fmt"
)

type Category struct {
	Id       int         `json:"id"`
	ParentId int         `json:"parent_id"`
	Name     string      `json:"name"`
	Children []*Category `json:"children,omitempty"`
}

func mergeCategoriesAndMeta(
	categoryTree *Category,
	metaCategories []*Category,
) ([]*Category, error) {
	if categoryTree == nil {
		return nil, errors.New("categoryTree cannot be nil")
	}
	rootNodes := []*Category{categoryTree}

	for _, metaCategory := range metaCategories {
		if metaCategory.ParentId == 0 {
			rootNodes = append(rootNodes, metaCategory)
		} else {
			parent := findCategoryById(categoryTree, metaCategory.ParentId)
			if parent == nil {
				return nil, fmt.Errorf("parent with ID %d not found", metaCategory.ParentId)
			}
			parent.Children = append(parent.Children, metaCategory)
		}
	}

	return rootNodes, nil
}

func findCategoryById(category *Category, id int) *Category {
	if category.Id == id {
		return category
	}
	for _, child := range category.Children {
		if result := findCategoryById(child, id); result != nil {
			return result
		}
	}
	return nil
}

func main() {
	categoryTree := &Category{
		Id:       1,
		ParentId: 0,
		Name:     "Root",
		Children: []*Category{
			{Id: 2, ParentId: 1, Name: "Child 1"},
			{Id: 3, ParentId: 1, Name: "Child 2"},
		},
	}

	metaCategories := []*Category{
		{Id: 5, ParentId: 0, Name: "Meta Root"},
		{Id: 6, ParentId: 0, Name: "New Meta Root"},
		{Id: 10, ParentId: 0, Name: "Another Meta Root"},
	}

	mergedTree, err := mergeCategoriesAndMeta(categoryTree, metaCategories)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, root := range mergedTree {
		printCategoryTree(root, 0)
	}
}

func printCategoryTree(category *Category, level int) {
	// Print the current category

	// Print the children recursively if they exist
	for _, child := range category.Children {
		printCategoryTree(child, level+1)
	}
}
