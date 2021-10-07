package request

type NearByRequest struct {
	Longitude float64 `json:"longitude"`
	Latitude float64 `json:"latitude"`
}

type SearchNameRequest struct {
	Name string `json:"name"`
}

type GetShopServiceRequest struct {
	Id int `json:"id"`
}