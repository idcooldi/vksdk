package api // import "github.com/idcooldi/vksdk/5.92/api"

import "github.com/idcooldi/vksdk/5.92/object"

// PrettyCardsCreateResponse struct
type PrettyCardsCreateResponse struct {
	OwnerID int    `json:"owner_id"` // Owner ID of created pretty card
	CardID  string `json:"card_id"`  // Card ID of created pretty card
}

// PrettyCardsCreate
//
// https://vk.com/dev/prettyCards.create
func (vk *VK) PrettyCardsCreate(params map[string]string) (response PrettyCardsCreateResponse, vkErr Error) {
	vk.RequestUnmarshal("prettyCards.create", params, &response, &vkErr)
	return
}

// PrettyCardsDeleteResponse struct
type PrettyCardsDeleteResponse struct {
	OwnerID int    `json:"owner_id"` // Owner ID of created pretty card
	CardID  string `json:"card_id"`  // Card ID of created pretty card
	Error   string `json:"error"`    // Error reason if error happened
}

// PrettyCardsDelete
//
// https://vk.com/dev/prettyCards.delete
func (vk *VK) PrettyCardsDelete(params map[string]string) (response PrettyCardsDeleteResponse, vkErr Error) {
	vk.RequestUnmarshal("prettyCards.delete", params, &response, &vkErr)
	return
}

// PrettyCardsEditResponse struct
type PrettyCardsEditResponse struct {
	OwnerID int    `json:"owner_id"` // Owner ID of created pretty card
	CardID  string `json:"card_id"`  // Card ID of created pretty card
}

// PrettyCardsEdit
//
// https://vk.com/dev/prettyCards.edit
func (vk *VK) PrettyCardsEdit(params map[string]string) (response PrettyCardsEditResponse, vkErr Error) {
	vk.RequestUnmarshal("prettyCards.edit", params, &response, &vkErr)
	return
}

// PrettyCardsGetResponse struct
type PrettyCardsGetResponse struct {
	Count int                            `json:"count"` // Total number
	Items []object.PrettyCardsPrettyCard `json:"items"`
}

// PrettyCardsGet
//
// https://vk.com/dev/prettyCards.get
func (vk *VK) PrettyCardsGet(params map[string]string) (response PrettyCardsGetResponse, vkErr Error) {
	vk.RequestUnmarshal("prettyCards.get", params, &response, &vkErr)
	return
}

// PrettyCardsGetByIDResponse struct
type PrettyCardsGetByIDResponse []object.PrettyCardsPrettyCard

// PrettyCardsGetByID
//
// https://vk.com/dev/prettyCards.getById
func (vk *VK) PrettyCardsGetByID(params map[string]string) (response PrettyCardsGetByIDResponse, vkErr Error) {
	vk.RequestUnmarshal("prettyCards.getById", params, &response, &vkErr)
	return
}

// PrettyCardsGetUploadURLResponse struct
type PrettyCardsGetUploadURLResponse string // Upload URL

// PrettyCardsGetUploadURL
//
// https://vk.com/dev/prettyCards.getUploadURL
func (vk *VK) PrettyCardsGetUploadURL() (response PrettyCardsGetUploadURLResponse, vkErr Error) {
	vk.RequestUnmarshal("prettyCards.getUploadURL", map[string]string{}, &response, &vkErr)
	return
}
