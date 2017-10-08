package main

//go:generate jedi

//Product is a sku representation.
//jedi:
type Product struct {
	ID         int64       `jedi:"@pk"`
	SKU        string      //todo: see if can be pk (string not int)
	categories []*Category `jedi:"@has_many=Category.products"`
	brand      *Brand      `jedi:"@has_one=Brand.products"`
	BrandID    *int64      // in case of has_one, you declare the exported keys.
	brand2     *Brand      `jedi:"@has_one=Brand.products2"`
	Brand2ID   *int64      // in case of has_one, you declare the exported keys.
	master     *Product    `jedi:"@has_one=Product.master"`
	MasterID   *int64      // in case of has_one, you declare the exported keys.
	variances  []*Product  `jedi:"@has_many=Product.master"`
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
	ID        int64      `jedi:"@pk"`
	products  []*Product `jedi:"@has_many=Product.brand"`
	products2 []*Product `jedi:"@has_many=Product.brand2"`
	Name      string
}
