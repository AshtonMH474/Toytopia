package models

type ToysInWishlist struct {
	Toy        Toy      `gorm:"foreignKey:ToyId;constraint:OnDelete:CASCADE;"`
	ToyId      int      `json:"toy_id"`
	WishlistId int      `json:"wishlist_id"`
	Wishlist   Wishlist `gorm:"foreignKey:WishlistId;constraint:OnDelete:CASCADE;"`
}
