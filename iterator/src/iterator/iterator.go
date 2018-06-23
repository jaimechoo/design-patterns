package iterator

import (
	"encoding/json"
	"sync"
)

type iterator interface {
	HasNext() bool
	Next() interface{}
}

type aggregate interface {
	Itetator() iterator
}

type Book struct {
	Name      string
	Price     string
	Publisher string
}

func (b *Book) String() string {
	bytes, _ := json.Marshal(b)
	return string(bytes)
}

func GetBookShelf() *bookShelf {
	var one bookShelf
	one.Init()
	return &one
}

type bookShelf struct {
	Books    []Book
	KeyIndex map[string]int
	sync.RWMutex
}

func (b *bookShelf) Init() {
	b.KeyIndex = make(map[string]int)
	return
}

func (b *bookShelf) Getlength() int {
	return len(b.Books)
}

func (b *bookShelf) Iterator() iterator {
	var one bookShelfIterator
	one.init(b)
	return &one
}

func (b *bookShelf) Add(book Book) {
	b.Lock()
	defer b.Unlock()
	index, ok := b.KeyIndex[book.Name]
	if ok {
		b.Books[index] = book
		return
	} else {
		b.KeyIndex[book.Name] = b.Getlength()
		b.Books = append(b.Books, book)
		return
	}
}

func (b *bookShelf) Delete(book Book) {
	b.Lock()
	defer b.Unlock()
	index, ok := b.KeyIndex[book.Name]
	if ok {
		delete(b.KeyIndex, book.Name)
		b.Books = append(b.Books[:index], b.Books[index+1:]...)
		return
	} else {
		return
	}
}

func (b *bookShelf) GetByKey(key string) *Book {
	index, ok := b.KeyToIndex(key)
	if ok {
		return b.GetByIndex(index)
	} else {
		return nil
	}
}

func (b *bookShelf) KeyToIndex(key string) (int, bool) {
	b.RLock()
	defer b.RUnlock()
	index, ok := b.KeyIndex[key]
	return index, ok
}

func (b *bookShelf) GetByIndex(index int) *Book {
	b.RLock()
	defer b.RUnlock()
	if index < b.Getlength() {
		return &b.Books[index]
	} else {
		return nil
	}
}

type bookShelfIterator struct {
	bookshelf *bookShelf
	index     int
}

func (b *bookShelfIterator) init(i *bookShelf) {
	b.bookshelf = i
}

func (b *bookShelfIterator) HasNext() bool {
	if b.index < b.bookshelf.Getlength() {
		return true
	} else {
		return false
	}
}

func (b *bookShelfIterator) Next() interface{} {
	book := b.bookshelf.GetByIndex(b.index)
	b.index++
	return book
}
