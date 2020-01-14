package api

import (
	"os"
	"reflect"
	"testing"

	"github.com/idcooldi/vksdk/5.92/object"
)

func TestVK_GroupsAddCallbackServer(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsAddCallbackServerResponse
		wantVkErr    Error
	}{
		{
			name:      "groups.addCallbackServer error",
			wantVkErr: Error{Code: object.ErrorParam},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkGroup.GroupsAddCallbackServer(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.GroupsAddCallbackServer() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsAddCallbackServer() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_GroupsDeleteAddress(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name      string
		argParams map[string]string
		wantVkErr Error
	}{
		{
			name:      "groups.deleteAddress error",
			wantVkErr: Error{Code: object.ErrorParam},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, gotVkErr := vkGroup.GroupsDeleteAddress(tt.argParams); gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsDeleteAddress() = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_GroupsDeleteCallbackServer(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name      string
		argParams map[string]string
		wantVkErr Error
	}{
		{
			name:      "groups.deleteCallbackServer error",
			wantVkErr: Error{Code: object.ErrorParam},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, gotVkErr := vkGroup.GroupsDeleteCallbackServer(tt.argParams); gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsDeleteCallbackServer() = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_GroupsEditCallbackServer(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name      string
		argParams map[string]string
		wantVkErr Error
	}{
		{
			name:      "groups.editCallbackServer error",
			wantVkErr: Error{Code: object.ErrorParam},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, gotVkErr := vkGroup.GroupsEditCallbackServer(tt.argParams); gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsEditCallbackServer() = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}
func TestVK_GroupsEnableOnline(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name      string
		argParams map[string]string
		wantVkErr Error
	}{
		{
			name: "groups.enableOnline",
			argParams: map[string]string{
				"group_id": os.Getenv("GROUP_ID"),
			},
		},
		{
			name:      "groups.enableOnline error",
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, gotVkErr := vkGroup.GroupsEnableOnline(tt.argParams); gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsEnableOnline() = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_GroupsDisableOnline(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name      string
		argParams map[string]string
		wantVkErr Error
	}{
		{
			name: "groups.disableOnline",
			argParams: map[string]string{
				"group_id": os.Getenv("GROUP_ID"),
			},
		},
		{
			name:      "groups.disableOnline error",
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, gotVkErr := vkGroup.GroupsDisableOnline(tt.argParams); gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsDisableOnline() = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}
func TestVK_GroupsGetBanned(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsGetBannedResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkGroup.GroupsGetBanned(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.GroupsGetBanned() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsGetBanned() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_GroupsGetByID(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	tests := []struct {
		name      string
		argParams map[string]string
		// wantResponse GroupsGetByIDResponse
		wantVkErr Error
	}{
		{
			name: "groups.getById",
			argParams: map[string]string{
				"group_ids": "1",
			},
		},
		{
			name: "groups.getById error",
			argParams: map[string]string{
				"group_ids": "0",
			},
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkService.GroupsGetByID(tt.argParams)
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsGetByID() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
			if gotVkErr.Code == 0 && (gotResponse[0].ID != 1) {
				t.Errorf("VK.GroupsGetByID() gotResponse = %v", gotResponse)
			}
		})
	}
}

func TestVK_GroupsGetCallbackConfirmationCode(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name      string
		argParams map[string]string
		// wantResponse GroupsGetCallbackConfirmationCodeResponse
		wantVkErr Error
	}{
		{
			name: "groups.getCallbackConfirmationCode",
			argParams: map[string]string{
				"group_id": os.Getenv("GROUP_ID"),
			},
		},
		{
			name:      "groups.getCallbackConfirmationCode error",
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkGroup.GroupsGetCallbackConfirmationCode(tt.argParams)
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsGetCallbackConfirmationCode() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
			if gotVkErr.Code == 0 && (gotResponse.Code == "") {
				t.Errorf("VK.GroupsGetCallbackConfirmationCode() gotResponse = %v", gotResponse)
			}
		})
	}
}

func TestVK_GroupsGetCallbackServers(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name      string
		argParams map[string]string
		wantVkErr Error
	}{
		{
			name: "groups.getCallbackSettings",
			argParams: map[string]string{
				"group_id": os.Getenv("GROUP_ID"),
			},
		},
		{
			name:      "groups.getCallbackSettings error",
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkGroup.GroupsGetCallbackServers(tt.argParams)
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsGetCallbackServers() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
			if gotVkErr.Code == 0 && (gotResponse.Count != len(gotResponse.Items)) {
				t.Errorf("VK.GroupsGetCallbackServers() gotResponse = %v", gotResponse)
			}
		})
	}
}

func TestVK_GroupsGetCallbackSettings(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}
	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsGetCallbackSettingsResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkGroup.GroupsGetCallbackSettings(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.GroupsGetCallbackSettings() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsGetCallbackSettings() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_GroupsGetLongPollServer(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name      string
		argParams map[string]string
		wantVkErr Error
	}{
		{
			name: "groups.getLongPollServer enabled",
			argParams: map[string]string{
				"group_id": os.Getenv("GROUP_ID"),
			},
		},
		{
			name:      "groups.getLongPollServer error",
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkGroup.GroupsGetLongPollServer(tt.argParams)
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsGetLongPollServer() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
			if gotVkErr.Code == 0 && (gotResponse.Key == "" || gotResponse.Server == "" || gotResponse.Ts == "") {
				t.Errorf("VK.GroupsGetLongPollServer() gotResponse = %v", gotResponse)
			}
		})
	}
}

func TestVK_GroupsGetLongPollSettings(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name      string
		argParams map[string]string
		wantVkErr Error
	}{
		{
			name: "groups.getLongPollSettings enabled",
			argParams: map[string]string{
				"group_id": os.Getenv("GROUP_ID"),
			},
		},
		{
			name:      "groups.getLongPollSettings error",
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotVkErr := vkGroup.GroupsGetLongPollSettings(tt.argParams)
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsGetLongPollSettings() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_GroupsGetMembers(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}
	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsGetMembersResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkGroup.GroupsGetMembers(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.GroupsGetMembers() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsGetMembers() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_GroupsGetMembersFilterManagers(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsGetMembersFilterManagersResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkGroup.GroupsGetMembersFilterManagers(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.GroupsGetMembersFilterManagers() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsGetMembersFilterManagers() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_GroupsGetOnlineStatus(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsGetOnlineStatusResponse
		wantVkErr    Error
	}{
		{
			name: "groups.getOnlineStatus enabled",
			argParams: map[string]string{
				"group_id": os.Getenv("GROUP_ID"),
			},
			wantResponse: GroupsGetOnlineStatusResponse{
				Status: "none",
			},
		},
		{
			name:      "groups.getOnlineStatus error",
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkGroup.GroupsGetOnlineStatus(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.GroupsGetOnlineStatus() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsGetOnlineStatus() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_GroupsGetTokenPermissions(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	gotResponse, gotVkErr := vkGroup.GroupsGetTokenPermissions()
	if gotResponse.Mask == 0 {
		t.Errorf("VK.GroupsGetTokenPermissions() gotResponse = %v", gotResponse)
	}
	if gotVkErr.Code != 0 {
		t.Errorf("VK.GroupsGetTokenPermissions() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_GroupsSetCallbackSettings(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name      string
		argParams map[string]string
		wantVkErr Error
	}{
		{
			name:      "groups.setCallbackSettings error",
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, gotVkErr := vkGroup.GroupsSetCallbackSettings(tt.argParams); gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsSetCallbackSettings() = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_GroupsSetLongPollSettings(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}
	tests := []struct {
		name      string
		argParams map[string]string
		wantVkErr Error
	}{
		{
			name: "groups.setLongPollSettings enabled",
			argParams: map[string]string{
				"group_id": os.Getenv("GROUP_ID"),
				"enabled":  "1",
			},
		},
		{
			name:      "groups.setLongPollSettings error",
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, gotVkErr := vkGroup.GroupsSetLongPollSettings(tt.argParams); gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsSetLongPollSettings() = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_GroupsGetAddresses(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsGetAddressesResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		{
			name:      "groups.getAddresses error",
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkService.GroupsGetAddresses(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.GroupsGetAddresses() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsGetAddresses() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_GroupsEditAddress(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsEditAddressResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		{
			name:      "groups.editAddresses error",
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkGroup.GroupsEditAddress(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.GroupsEditAddress() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsEditAddress() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_GroupsAddAddress(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsAddAddressResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		{
			name:      "groups.addAddresses error",
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkGroup.GroupsAddAddress(tt.argParams)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.GroupsAddAddress() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsAddAddress() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_GroupsIsMember(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsIsMemberResponse
		wantVkErr    Error
	}{
		{
			name: "groups.isMembers",
			argParams: map[string]string{
				"group_id": "1",
				"user_id":  "117253521",
			},
			wantResponse: 1,
		},
		{
			name:      "groups.isMembers error",
			argParams: map[string]string{},
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkGroup.GroupsIsMember(tt.argParams)
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsIsMember() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.GroupsIsMember() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestVK_GroupsIsMemberExtended(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsIsMemberExtendedResponse
		wantVkErr    Error
	}{
		{
			name: "groups.isMembers",
			argParams: map[string]string{
				"group_id": "1",
				"user_id":  "117253521",
			},
			wantResponse: GroupsIsMemberExtendedResponse{
				Member:    1,
				CanInvite: 0,
			},
		},
		{
			name:      "groups.isMembers error",
			argParams: map[string]string{},
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkGroup.GroupsIsMemberExtended(tt.argParams)
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsIsMemberExtended() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.GroupsIsMemberExtended() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestVK_GroupsIsMemberUserIDsExtended(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsIsMemberUserIDsExtendedResponse
		wantVkErr    Error
	}{
		{
			name: "groups.isMembers",
			argParams: map[string]string{
				"group_id": "1",
				"user_ids": "117253521",
			},
			wantResponse: GroupsIsMemberUserIDsExtendedResponse{
				{
					Member:    1,
					CanInvite: 0,
					UserID:    117253521,
				},
			},
		},
		{
			name:      "groups.isMembers error",
			argParams: map[string]string{},
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkGroup.GroupsIsMemberUserIDsExtended(tt.argParams)
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsIsMemberUserIDsExtended() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.GroupsIsMemberUserIDsExtended() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestVK_GroupsIsMemberUserIDs(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsIsMemberUserIDsResponse
		wantVkErr    Error
	}{
		{
			name: "groups.isMembers",
			argParams: map[string]string{
				"group_id": "1",
				"user_ids": "117253521",
			},
			wantResponse: GroupsIsMemberUserIDsResponse{
				{
					Member: 1,
					UserID: 117253521,
				},
			},
		},
		{
			name:      "groups.isMembers error",
			argParams: map[string]string{},
			wantVkErr: Error{Code: object.ErrorParam},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkGroup.GroupsIsMemberUserIDs(tt.argParams)
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.GroupsIsMemberUserIDs() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.GroupsIsMemberUserIDs() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
