package entities


type Payment struct {
    ID        uint32    `gorm:"primaryKey"`
    OrderID   uint32    `gorm:"not null"`
    Status    string    `gorm:"not null"`
}