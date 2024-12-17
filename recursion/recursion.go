package main

import (
	"encoding/json"
	"fmt"
)

type Node struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Child []Node `json:"child"`
}

func updateNameByID(nodes []Node, id int, newName string) bool {
	for i := range nodes {
		if nodes[i].ID == id {
			nodes[i].Name = newName
			return true
		}
		if updateNameByID(nodes[i].Child, id, newName) {
			return true
		}
	}
	return false
}

func main() {
	jsonData := `
		[
			{
				"id": 1,
				"name": "Parent Node 1",
				"child": [
					{
						"id": 2,
						"name": "Child Node 1.1",
						"child": [
							{
								"id": 5,
								"name": "Grandchild Node 1.1.1",
								"child": []
							},
							{
								"id": 6,
								"name": "Grandchild Node 1.1.2",
								"child": []
							}
						]
					},
					{
						"id": 3,
						"name": "Child Node 1.2",
						"child": [
							{
								"id": 7,
								"name": "Grandchild Node 1.2.1",
								"child": [
									{
										"id": 10,
										"name": "Great-Grandchild Node 1.2.1.1",
										"child": []
									}
								]
							},
							{
								"id": 8,
								"name": "Grandchild Node 1.2.2",
								"child": []
							}
						]
					}
				]
			},
			{
				"id": 4,
				"name": "Parent Node 2",
				"child": [
					{
						"id": 9,
						"name": "Child Node 2.1",
						"child": []
					},
					{
						"id": 11,
						"name": "Child Node 2.2",
						"child": [
							{
								"id": 12,
								"name": "Grandchild Node 2.2.1",
								"child": [
									{
										"id": 15,
										"name": "Great-Grandchild Node 2.2.1.1",
										"child": []
									}
								]
							}
						]
					}
				]
			}
		]`

	var nodes []Node
	if err := json.Unmarshal([]byte(jsonData), &nodes); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	idToUpdate := 6
	newName := "Md Abir Hasan"
	if updateNameByID(nodes, idToUpdate, newName) {
		fmt.Println("Name updated successfully.")
	} else {
		fmt.Println("ID not found.")
	}

	updatedJSON, err := json.MarshalIndent(nodes, "", "  ")
	if err != nil {
		fmt.Println("Error converting back to JSON:", err)
		return
	}
	fmt.Println("Updated JSON:")
	fmt.Println(string(updatedJSON))
}
