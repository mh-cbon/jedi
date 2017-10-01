package main

//go:generate jedi

//Product is a sku representation.
//jedi:
type Product struct {
	ID         int64       `jedi:"@pk"`
	SKU        string      //todo: see if can be pk (string not int)
	categories []*Category `jedi:"@has_many=Category.products"`
	brand      *Brand      `jedi:"@has_one=Brand"`
	BrandID    *int64      // in case of has_one, you declare the exported keys.
}

//Category is a product category representation.
//jedi:
type Category struct {
	ID       int64      `jedi:"@pk"`
	products []*Product `jedi:"@has_many=Product.categories"`
	Name     string
}

//Brand is a product brand representation.
//jedi:
type Brand struct {
	ID       int64      `jedi:"@pk"`
	products []*Product `jedi:"@has_many=Product"`
	Name     string
}
