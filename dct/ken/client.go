package ken

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"

	"github.com/RPA-Gang/gograph"
)

const (
	kenSiteId    string = "371338c7-a49f-4a22-94b6-6be107be6a66,52546bf9-b7da-47af-86a2-53234a925c76"
	kenListId    string = "c91934b5-f989-4251-810b-dd458b9b5680"
	kenUatSiteId string = "371338c7-a49f-4a22-94b6-6be107be6a66,66fb87e2-113e-455d-a6a6-48145bceb5d9"
	kenUatListId string = "b705763b-9854-444b-b88c-30bc0bec9498"
)

var DefaultKenClient IKenClient

type IKenClient interface {
	gograph.IGraphClient
	GetListItemResponse(reqUrl string) (*GetListItemResponse, error)
	GetListItems() ([]ListItem, error)
	CreateListItem(fields NewListItem) (ListItem, error)
	DeleteListItem(id int) error
}

func NewKenClient(graphClient gograph.IGraphClient, environment IApiEnvironment) (IKenClient, error) {
	if _, ok := graphClient.(*gograph.GraphClient); !ok {
		return nil, errors.New("graphClient must satisfy gograph.IGraphClient interface")
	} else {
		return &kenClient{graphClient.(*gograph.GraphClient), environment}, nil
	}
}

func SetDefaultKenClient(client IKenClient) {
	DefaultKenClient = client
}

type kenClient struct {
	*gograph.GraphClient
	environment IApiEnvironment
}

func (k *kenClient) GetListItemResponse(reqUrl string) (*GetListItemResponse, error) {
	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to create request")
	}
	resp, err := k.Client().Do(req)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to execute request")
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	var parsedResp GetListItemResponse
	if err = json.NewDecoder(resp.Body).Decode(&parsedResp); err != nil {
		return nil, errors.WithMessage(err, "failed to parse response")
	}
	return &parsedResp, nil
}

func (k *kenClient) GetListItems() ([]ListItem, error) {
	reqUrl := fmt.Sprintf(
		"%s/sites/%s/lists/%s/items?$expand=fields",
		k.BaseUrl(),
		k.environment.SiteId(),
		k.environment.ListId(),
	)
	parsedResp, err := k.GetListItemResponse(reqUrl)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to get first page of list items")
	}
	var listItems []ListItem
	for i := 0; ; i++ {
		if parsedResp.Value != nil {
			listItems = append(listItems, parsedResp.Value...)
		} else {
			break
		}
		if len(parsedResp.OdataNextLink) > 0 {
			parsedResp, err = k.GetListItemResponse(parsedResp.OdataNextLink)
			if err != nil {
				return nil, errors.Errorf("failed to get page %d of list items: %v", i+1, err)
			}
		} else {
			break
		}
	}
	return listItems, nil
}

func (k *kenClient) CreateListItem(fields NewListItem) (ListItem, error) {
	var result ListItem
	jsonFields, err := json.MarshalIndent(fields, "", "\t")
	if err != nil {
		return result, errors.WithMessage(err, "failed to marshal fields")
	}
	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf(
			"%s/sites/%s/lists/%s/items",
			k.BaseUrl(),
			k.environment.SiteId(),
			k.environment.ListId(),
		),
		bytes.NewBuffer(jsonFields),
	)
	//goland:noinspection GoDfaErrorMayBeNotNil
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept-Encoding", "gzip")
	resp, err := k.Client().Do(req)
	if err != nil {
		return result, errors.WithMessage(err, "failed to create list item")
	}
	//goland:noinspection GoUnhandledErrorResult
	defer resp.Body.Close()
	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return result, errors.WithMessage(err, "failed to create gzip reader")
		}
	default:
		reader = resp.Body
	}
	//goland:noinspection GoUnhandledErrorResult
	defer reader.Close()
	var buf bytes.Buffer
	if _, err = io.Copy(&buf, reader); err != nil {
		return result, errors.WithMessage(err, "failed to copy response body")
	}
	if resp.StatusCode != http.StatusCreated {
		return result, errors.Errorf("failed to create list item: %s", buf.String())
	}
	var parsedListItem ListItem
	if err = json.NewDecoder(&buf).Decode(&parsedListItem); err != nil {
		return result, errors.WithMessage(err, "failed to parse response")
	}
	return parsedListItem, nil
}

func (k *kenClient) DeleteListItem(id int) error {
	req, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf(
			"%s/sites/%s/lists/%s/items/%d",
			k.BaseUrl(),
			k.environment.SiteId(),
			k.environment.ListId(),
			id,
		),
		nil,
	)
	resp, err := k.Client().Do(req)
	if err != nil {
		return errors.WithMessage(err, "deletion failed")
	} else if resp.StatusCode != http.StatusNoContent {
		err = errors.Errorf("failed to delete list item: %d - %s", id, resp.Status)
		return errors.WithMessage(err, "deletion failed")
	}
	//goland:noinspection GoUnhandledErrorResult
	defer resp.Body.Close()
	return nil
}
