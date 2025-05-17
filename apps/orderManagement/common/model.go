package ordercommon

type CsvData struct {
	// Product Data
	ProductID   string `csv:"Product ID"`
	ProductName string `csv:"Product Name"`
	Category    string `csv:"Category"`

	// Orders Data
	OrderID       string  `csv:"Order ID"`
	Region        string  `csv:"Region"`
	DateOfSale    string  `csv:"DateOfSale"`
	ShippingCost  float64 `csv:"ShippingCost"`
	PaymentMethod string  `csv:"PaymentMethod"`

	// Orders Items Data
	Quantity  int     `csv:"Quantity Sold"`
	UnitPrice float64 `csv:"Unit Price"`
	Discount  float64 `csv:"Discount"`

	// Customer Data
	CustomerID      string `csv:"Customer ID"`
	CustomerName    string `csv:"Customer Name"`
	CustomerEmail   string `csv:"Customer Email"`
	CustomerAddress string `csv:"Customer Address"`
}

type Customer struct {
	CustomerID      string `csv:"CustomerID"`
	CustomerName    string `csv:"CustomerName"`
	CustomerEmail   string `csv:"CustomerEmail"`
	CustomerAddress string `csv:"CustomerAddress"`
}

type Product struct {
	ProductID   string `csv:"ProductID"`
	ProductName string `csv:"ProductName"`
	Category    string `csv:"Category"`
}

type Order struct {
	OrderID       string  `csv:"OrderID"`
	CustomerID    string  `csv:"CustomerID"`
	Region        string  `csv:"Region"`
	DateOfSale    string  `csv:"DateOfSale"`
	ShippingCost  float64 `csv:"ShippingCost"`
	PaymentMethod string  `csv:"PaymentMethod"`
}

type OrderItem struct {
	OrderID   string  `csv:"OrderID"`
	ProductID string  `csv:"ProductID"`
	Quantity  int     `csv:"Quantity"`
	UnitPrice float64 `csv:"UnitPrice"`
	Discount  float64 `csv:"Discount"`
}

type RevenueStruct struct {
	RevenueWithDiscount    string `json:"totalRevenueWithDis,omitempty" `
	RevenueWithoutDiscount string `json:"totalRevenueWithoutDis,omitempty" `
}
type RevenueResp struct {
	ProductName   string `json:"product_name,omitempty" `
	CatagoryName  string `json:"catagiryName,omitempty" `
	RegionName    string `json:"regionName,omitempty" `
	RevenueStruct `json:"Revenue" `
}

type RevenueRange struct {
	Month         string `json:"month,omitempty" `
	Year          string `json:"year,omitempty" `
	Quater        string `json:"quater,omitempty" `
	RevenueStruct `json:"Revenue" `
}

type RequestStruct struct {
	FromDate  string `json:"fromDate"`
	ToDate    string `json:"toDate"`
	RangeType string `json:"rangeType"`
}

const (
	GetTotalRevenue    = "GetTotalRevenue"
	GetCategoryRevenue = "GetCategoryRevenue"
	GetProductRevenue  = "GetProductRevenue"
	GetRegionRevenue   = "GetRegionRevenue"
)
