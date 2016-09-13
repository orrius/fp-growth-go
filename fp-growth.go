package main

import (
	"math/rand"
	"fmt"
	"sort"
)

type node struct {
	id int
	children []node
	count int
}

type frequentItem struct {
	Id int
	Support int
}

type frequentItemList []frequentItem

// Implement interface for sort
func (p frequentItemList) Len() int { return len(p) }
func (p frequentItemList) Less(i, j int) bool { return p[i].Support > p[j].Support }
func (p frequentItemList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }
func (p frequentItemList) In(id int) bool { 
	for i,_ := range p {
		if id == i {
			return true
		}
	}
	return false
}

func main() {
	orders := generateData()
	fmt.Println(orders)
	fList := collectAndSort(orders, 0)
	fmt.Println(fList)
	fmt.Println(fpTree(orders, fList))
}

func fpTree(orders [][]int, fList frequentItemList) node {
	topNode := node{0, nil, 0}
	for _, order := range orders {
		fmt.Println(order)
		frequentIds := frequentItemList{}
		for _, id := range order {
			if fList.In(id) {
				frequentIds = append(frequentIds, frequentItem{id, fList[id].Support})
			}
		}
		sort.Sort(frequentIds)
		fmt.Println(frequentIds)
	}
	return topNode
}

func collectAndSort(orders [][]int, minimumSupport int) frequentItemList {
	supportList := frequentItemList{}
	itemCount := make(map[int]int)
	for _ , order := range orders {
		for _, item := range order {
			// Default value is zero
			itemCount[item] = itemCount[item] + 1
		}
	}
	for k, v := range itemCount {
		supportList = append(supportList, frequentItem{k, v})
	}
	sort.Sort(supportList)
	finalSupportList := frequentItemList{}
	for k, v := range supportList {
		if v.Support > minimumSupport {
			finalSupportList = append(finalSupportList, supportList[k])
		} else {
			break
		}
	}
	return finalSupportList
}

func generateData() [][]int {
	orders := [][]int{}
	for i := 0;i < 100;i++ {
		numberOfItems := rand.Intn(5)
		basket := []int{}
		for n := 0; n < numberOfItems;n++ {
			basket = append(basket, rand.Intn(100))
		}
		orders = append(orders, basket)
	}
	return orders
}