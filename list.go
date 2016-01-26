package main

import (
	"fmt"
	"strings"
)

type stringList []string

func (list *stringList) String() string {
	return fmt.Sprint(*list)
}

func (list *stringList) Set(value string) error {
	for _, elem := range strings.Split(value, ",") {
		list.Add(elem)
	}
	return nil
}

func (list *stringList) Add(value string) error {
	*list = append(*list, value)
	return nil
}

func (list *stringList) Len() int {
	return len(*list)
}

func (list *stringList) stringInList(a string) bool {
	for _, b := range *list {
		if b == a {
			return true
		}
	}
	return false
}
