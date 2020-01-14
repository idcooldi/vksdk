package api // import "github.com/idcooldi/vksdk/5.92/api"

import (
	"github.com/idcooldi/vksdk/5.92/object"
)

// WidgetsGetCommentsResponse struct
type WidgetsGetCommentsResponse struct {
	Count int                           `json:"count"`
	Posts []object.WidgetsWidgetComment `json:"posts"`
}

// WidgetsGetComments gets a list of comments for the page added through the Comments widget.
//
// https://vk.com/dev/widgets.getComments
func (vk *VK) WidgetsGetComments(params map[string]string) (response WidgetsGetCommentsResponse, vkErr Error) {
	vk.RequestUnmarshal("widgets.getComments", params, &response, &vkErr)
	return
}

// WidgetsGetPagesResponse struct
type WidgetsGetPagesResponse struct {
	Count int                        `json:"count"`
	Pages []object.WidgetsWidgetPage `json:"pages"`
}

// WidgetsGetPages gets a list of application/site pages where the Comments widget or Like widget is installed.
//
// https://vk.com/dev/widgets.getPages
func (vk *VK) WidgetsGetPages(params map[string]string) (response WidgetsGetPagesResponse, vkErr Error) {
	vk.RequestUnmarshal("widgets.getPages", params, &response, &vkErr)
	return
}
