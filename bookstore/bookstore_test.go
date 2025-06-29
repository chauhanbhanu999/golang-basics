package bookstore_test

import (
	"bookstore"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "title 1",
		Author: "author 1",
		Copies: 2,
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "spark joy",
		Author: "marie kondo",
		Copies: 2,
	}
	want := 1
	result, err := bookstore.Buy(b)
	if err != nil {
		t.Fatal(err)
	}
	got := result.Copies
	if want != got {
		t.Errorf("want %d got %d", want, got)
	}
}

func TestNoCopiesLeft(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "spark joy",
		Author: "marie kondo",
		Copies: 0,
	}
	_, err := bookstore.Buy(b)
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	catalog := map[int]bookstore.Book{
		1: {ID: 1, Title: "The Go Programming Language"},
		2: {ID: 2, Title: "Clean Code"},
	}
	want := []bookstore.Book{
		{ID: 1, Title: "The Go Programming Language"},
		{ID: 2, Title: "Clean Code"},
	}
	got := bookstore.GetAllBooks(catalog)
	sort.Slice(got, func(i, j int) bool {
		return got[i].ID < got[j].ID
	})
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()
	catalog := map[int]bookstore.Book{
		1: {ID: 1, Title: "The Go Programming Language"},
		2: {ID: 2, Title: "The Go Programming Language2"},
	}
	want := bookstore.Book{Title: "The Go Programming Language", ID: 1}

	got, err := bookstore.GetBook(catalog, 1)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestPurchaseOfInvalidBookID(t *testing.T) {
	t.Parallel()
	catalog := map[int]bookstore.Book{}
	_, err := bookstore.GetBook(catalog, 9999)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestNetPriceCents(t *testing.T) {
	t.Parallel()
	book := bookstore.Book{
		Title:           "book1",
		Author:          "author1",
		Copies:          1,
		PriceCents:      4000,
		DiscountPercent: 25,
	}
	want := 3000
	got := book.NetPriceCents()
	if want != got {
		t.Errorf("want %d got %d", want, got)
	}
}
