package api // import "github.com/idcooldi/vksdk/5.92/api"

// AuthCheckResponse struct
type AuthCheckResponse int

// AuthCheckPhone checks a user's phone number for correctness.
//
// https://vk.com/dev/auth.checkPhone
func (vk *VK) AuthCheckPhone(params map[string]string) (response AuthCheckResponse, vkErr Error) {
	vk.RequestUnmarshal("auth.checkPhone", params, &response, &vkErr)
	return
}

// AuthRetoreResponse struct
type AuthRetoreResponse struct {
	Success int    `json:"success"`
	SID     string `json:"sid"`
}

// AuthRetore allows to restore account access using a code received via SMS.
//
// https://vk.com/dev/auth.restore
func (vk *VK) AuthRetore(params map[string]string) (response AuthRetoreResponse, vkErr Error) {
	vk.RequestUnmarshal("auth.restore", params, &response, &vkErr)
	return
}
