package main

import (
	"fmt"
	"iterator"
)

func main() {
	var bookshelf = iterator.GetBookShelf()
	bookshelf.Add(iterator.Book{Name: "呼啸山庄"})
	bookshelf.Add(iterator.Book{Name: "百年孤独"})
	bookshelf.Delete(iterator.Book{Name: "百年孤独"})
	bookshelf.Add(iterator.Book{Name: "百年孤独"})
	bookshelf.Add(iterator.Book{Name: "雾都孤儿"})
	bookshelf.Add(iterator.Book{Name: "傲慢与偏见"})
	bookshelf.Add(iterator.Book{Name: "简爱"})
	bookshelf.Delete(iterator.Book{Name: "傲慢与偏见"})
	bookshelf.Add(iterator.Book{Name: "傲慢与偏见"})

	bookshelfiterator := bookshelf.Iterator()
	for bookshelfiterator.HasNext() {
		fmt.Println(bookshelfiterator.Next())
	}
}
