package ken

import (
	"time"
)

type ListItem struct {
	OdataEtag            string    `json:"@odata.etag"`
	CreatedDateTime      time.Time `json:"createdDateTime"`
	ETag                 string    `json:"eTag"`
	Id                   int       `json:"id,string"`
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime"`
	WebUrl               string    `json:"webUrl"`
	CreatedBy            struct {
		User `json:"user"`
	} `json:"createdBy"`
	LastModifiedBy struct {
		User `json:"user"`
	} `json:"lastModifiedBy"`
	ParentReference struct {
		Id     string `json:"id"`
		SiteId string `json:"siteId"`
	} `json:"parentReference"`
	ContentType struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"contentType"`
	FieldsOdataContext string         `json:"fields@odata.context"`
	Fields             ListItemFields `json:"fields"`
}

type ListItemFields struct {
	OdataEtag                         string    `json:"@odata.etag"`
	ActionRequiredCompleted           string    `json:"ActionRequiredCompleted"`
	AutomationOutcome                 string    `json:"AutomationOutcome"`
	Checked                           string    `json:"Checked"`
	EventDescriptionLookupId          int       `json:"EventDescriptionLookupId,string"`
	EventDescriptionEventDescLookupId int       `json:"EventDescription_x003a_EventDescLookupId,string"`
	IProcessNumber                    string    `json:"iProcess_x0020_Number"`
	MemberNumber                      string    `json:"Member_x0020_Number"`
	PlanShortNameLookupId             int       `json:"PlanShortNameLookupId,string"`
	PlanShortNamePlanShortNamLookupId int       `json:"PlanShortName_x003a_PlanShortNamLookupId,string"`
	PlanShortNameQueueLookupId        int       `json:"PlanShortName_x003a_Queue_x0020_LookupId,string"`
	ReceivedAfterCutoff               bool      `json:"Received_x0020_After_x0020_Cutof"`
	SentToQueue                       string    `json:"Sent_x0020_To_x0020_Queue"`
	Comment                           string    `json:"Comment"`
	IProcessNumberHidden              string    `json:"iProcess_x0020_Number_"`
	MemberNumberHidden                string    `json:"Member_x0020_Number_"`
	ReceivedAfterCutoffHidden         string    `json:"Received_x0020_After_x0020_Cutof0"`
	Status                            string    `json:"Status_"`
	CancelRequest                     string    `json:"Cancel_x0020_Request"`
	DateOfDeath                       time.Time `json:"Date_x0020_Of_x0020_Death"`
	MemberDOB                         time.Time `json:"Member_x0020_DOB0"`
	NotificationDate                  time.Time `json:"Notification_x0020_Date"`
	Modified                          time.Time `json:"Modified"`
	Created                           time.Time `json:"Created"`
	Id                                int       `json:"id,string"`
	ContentType                       string    `json:"ContentType"`
	AuthorLookupId                    string    `json:"AuthorLookupId"`
	EditorLookupId                    string    `json:"EditorLookupId"`
	UIVersionString                   string    `json:"_UIVersionString"`
	Attachments                       bool      `json:"Attachments"`
	Edit                              string    `json:"Edit"`
	ItemChildCount                    string    `json:"ItemChildCount"`
	FolderChildCount                  string    `json:"FolderChildCount"`
	ComplianceFlags                   string    `json:"_ComplianceFlags"`
	ComplianceTag                     string    `json:"_ComplianceTag"`
	ComplianceTagWrittenTime          string    `json:"_ComplianceTagWrittenTime"`
	ComplianceTagUserId               string    `json:"_ComplianceTagUserId"`
}

type User struct {
	Email       string `json:"email"`
	Id          string `json:"id"`
	DisplayName string `json:"displayName"`
}

type NewListItem struct {
	Fields NewListItemFields `json:"fields"`
}

type NewListItemFields struct {
	ActionRequiredCompleted  string    `json:"ActionRequiredCompleted"`
	AutomationOutcome        string    `json:"AutomationOutcome,omitempty"`
	Checked                  string    `json:"Checked,omitempty"`
	EventDescriptionLookupId int       `json:"EventDescriptionLookupId,string"`
	IProcessNumber           string    `json:"iProcess_x0020_Number"`
	MemberNumber             string    `json:"Member_x0020_Number"`
	PlanShortNameLookupId    int       `json:"PlanShortNameLookupId,string"`
	ReceivedAfterCutoff      bool      `json:"Received_x0020_After_x0020_Cutof"`
	Comment                  string    `json:"Comment"`
	DateOfDeath              time.Time `json:"Date_x0020_Of_x0020_Death"`
	MemberDOB                string    `json:"Member_x0020_DOB0"`
	NotificationDate         time.Time `json:"Notification_x0020_Date"`
	ContentType              string    `json:"ContentType,omitempty"`
}
