package ken

type GetListItemResponse struct {
	OdataContext  string     `json:"@odata.context"`
	OdataNextLink string     `json:"@odata.nextLink"`
	Value         []ListItem `json:"value"`
}
