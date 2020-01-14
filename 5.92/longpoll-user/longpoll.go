package longpoll // import "github.com/idcooldi/vksdk/5.92/longpoll-user"

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync/atomic"

	"github.com/idcooldi/vksdk/5.92/api"
	"github.com/idcooldi/vksdk/5.92/object"
)

// Longpoll struct
type Longpoll struct {
	Key     string
	Server  string
	Ts      int
	Mode    int
	Version int
	Wait    int
	VK      *api.VK
	Client  *http.Client

	funcList             FuncList
	funcFullResponseList []func(object.LongpollResponse)
	inShutdown           int32
}

// Init Longpoll
func Init(vk *api.VK, mode int) (lp Longpoll, err error) {
	// NOTE: what about group_id?
	lp.VK = vk
	lp.Mode = mode
	lp.Version = 3
	lp.Wait = 25
	lp.funcList = make(FuncList)
	lp.Client = &http.Client{}
	err = lp.updateServer(true)
	return
}

func (lp *Longpoll) updateServer(updateTs bool) error {
	params := map[string]string{
		"lp_version": strconv.Itoa(lp.Version),
	}
	serverSetting, vkErr := lp.VK.MessagesGetLongPollServer(params)
	if vkErr.Code != 0 {
		return fmt.Errorf(vkErr.Message)
	}

	lp.Key = serverSetting.Key
	lp.Server = serverSetting.Server
	if updateTs {
		lp.Ts = serverSetting.Ts
	}
	return nil
}

func (lp *Longpoll) check() (response object.LongpollResponse, err error) {
	u := fmt.Sprintf(
		"https://%s?act=a_check&key=%s&ts=%d&wait=%d&mode=%d&version=%d",
		lp.Server,
		lp.Key,
		lp.Ts,
		lp.Wait,
		lp.Mode,
		lp.Version,
	)

	resp, err := lp.Client.Get(u)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}

	err = lp.checkResponse(response)
	return
}

func (lp *Longpoll) checkResponse(response object.LongpollResponse) (err error) {
	switch response.Failed {
	case 0:
		lp.Ts = response.Ts
	case 1:
		lp.Ts = response.Ts
	case 2:
		err = lp.updateServer(false)
	case 3:
		err = lp.updateServer(true)
	case 4:
		err = fmt.Errorf("not valid version")
	default:
		err = fmt.Errorf(`"failed":%d`, response.Failed)
	}
	return
}

// Run handler
func (lp *Longpoll) Run() error {
	atomic.StoreInt32(&lp.inShutdown, 0)
	for atomic.LoadInt32(&lp.inShutdown) == 0 {
		resp, err := lp.check()
		if err != nil {
			return err
		}
		for _, event := range resp.Updates {
			lp.funcList.Handler(event)
		}
		for _, f := range lp.funcFullResponseList {
			f(resp)
		}
	}
	return nil
}

// Shutdown gracefully shuts down the longpoll without interrupting any active connections.
func (lp *Longpoll) Shutdown() {
	atomic.StoreInt32(&lp.inShutdown, 1)
}

// EventNew handler
func (lp *Longpoll) EventNew(key int, f EventNewFunc) {
	lp.funcList[key] = append(lp.funcList[key], f)
}

// FullResponse handler
func (lp *Longpoll) FullResponse(f func(object.LongpollResponse)) {
	lp.funcFullResponseList = append(lp.funcFullResponseList, f)
}
