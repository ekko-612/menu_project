package main

import "fmt"

type Node struct 
{
	value  int
	next *Node
}

type LinkTable struct 
{
	size int
	Head *Node
	Tail *Node
	
}

func CreateLinkTable() *LinkTable 
{
	linktable := &LinkTable
	{
		Head: nil,
		Tail: nil,
		size: 0,
	}
	return linktable
}

func (linktable *LinkTable) DeleteLinkTable() 
{
	linktable.Head = nil
	linktable.Tail = nil
	linktable.size = 0
}

func (linktable *LinkTable) AddNode(val int) 
{
	node := Node{val, nil}

	if linktable.size == 0 
	{
		linktable.Head = &node
		linktable.Tail = &node
	} 
	else 
	{
		linktable.Tail.next = &node
		linktable.Tail = &node
	}
	l.size++
	fmt.Printf("add value %d successfully\n", v)
	return
}

func (linktable *LinkTable) DeleteNode(val int) 
{
	i := linktable.Head

	if linktable.size == 0 
	{
		return
	}

	if i.val == val
	 {
		if linktable.Head.next != nil 
		{
			linktable.Head = linktable.Head.next
		} 
		else 
		{
			linktable.Head = nil
		}
		linktable.size--
		return
	}

	for i.next != nil 
	{
		if i.next.val == v 
		{
			fmt.Printf("Delete value %d successfully\n", v)
			if i.next.next != nil
		    {
				i.next = i.next.next
			} 
			else 
			{
				i.next = nil
			}
			linktable.size--
		}
		i = i.next
	}
}

func (linktable *LinkTable) Find(val int) *Node 
{
	if linktable.size == 0 
	{
		return nil
	} 
	else
	 {
		i := l.Head

		for i.next != nil
		 {
			if i.val == val 
			{
				return i
			}
			i= i.next
		}
		return nil
	}
}

func (linktable *LinkTable) UpdateNode(val int, newval int) *Node 
{
	if linktable.size == 0 
	{
		return nil
	} 
	else
	 {
		i := l.Head

		for i.next != nil
		 {
			if i.val == val 
			{
				i.val = newval
				return i
			}
			i= i.next
		}
		return nil
	}
}

func (linktable *LinkTable) PrintLinkTable() 
{
	i := linktable.Head

	for i.next != nil 
	{
		fmt.Printf("%d ", i.val)
		i = i.next
	}
	fmt.Println(i.val)
}

