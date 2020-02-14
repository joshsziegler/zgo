package main

// Paginator provides a simple 5-element paginator (first, previous, current,
// next, and last page). Use getPaginator() to generate this struct.
//
// Pages start on 1 and end on NumPages (inclusive).  PrevPage/NextPage will be
// 0 if the CurrPage is already at the first/last page respectively.
type Paginator struct {
	CurrPage int64
	PrevPage int64 // 0 if CurrPage == first page
	NextPage int64 // 0 if CurrPage == last page
	// NumItems is the total number of items that we are paginating over.
	NumItems        int64
	NumItemsPerPage int64
	NumPages        int64
}

// getPaginator returns a Paginator object given the current page, total number
// of items and the number of items per page.
func getPaginator(currPage, numItems, itemsPerPage int64) Paginator {
	p := Paginator{
		CurrPage:        currPage,
		NumItems:        numItems,
		NumItemsPerPage: itemsPerPage,
		NumPages:        numItems / itemsPerPage,
	}
	// Handle remainders by adding an extra partial page
	if numItems%itemsPerPage != 0 {
		p.NumPages += 1
	}
	// Handle cases where request current page is beyond the bounds
	if p.CurrPage < 1 {
		p.CurrPage = 1
	} else if p.CurrPage > p.NumPages {
		p.CurrPage = p.NumPages
	}
	// Calculate previous Page (leave as 0 if currPage is first page)
	if p.CurrPage > 1 {
		p.PrevPage = p.CurrPage - 1
	}
	// Caluclate next Page (leave as 0 if currPage is last page)
	if p.CurrPage < p.NumPages {
		p.NextPage = p.CurrPage + 1
	}
	return p
}
