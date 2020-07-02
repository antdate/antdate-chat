package grpc

import (
	"encoding/json"
	"time"

	"github.com/tinode/chat/server/auth"
	"github.com/tinode/chat/server/store/types"
)

type adapter struct {
}

func (a adapter) Open(config json.RawMessage) error {
	panic("implement me")
}

func (a adapter) Close() error {
	panic("implement me")
}

func (a adapter) IsOpen() bool {
	panic("implement me")
}

func (a adapter) GetDbVersion() (int, error) {
	panic("implement me")
}

func (a adapter) CheckDbVersion() error {
	panic("implement me")
}

func (a adapter) GetName() string {
	panic("implement me")
}

func (a adapter) SetMaxResults(val int) error {
	panic("implement me")
}

func (a adapter) CreateDb(reset bool) error {
	panic("implement me")
}

func (a adapter) UpgradeDb() error {
	panic("implement me")
}

func (a adapter) Version() int {
	panic("implement me")
}

func (a adapter) UserCreate(user *types.User) error {
	panic("implement me")
}

func (a adapter) UserGet(uid types.Uid) (*types.User, error) {
	panic("implement me")
}

func (a adapter) UserGetAll(ids ...types.Uid) ([]types.User, error) {
	panic("implement me")
}

func (a adapter) UserDelete(uid types.Uid, hard bool) error {
	panic("implement me")
}

func (a adapter) UserUpdate(uid types.Uid, update map[string]interface{}) error {
	panic("implement me")
}

func (a adapter) UserUpdateTags(uid types.Uid, add, remove, reset []string) ([]string, error) {
	panic("implement me")
}

func (a adapter) UserGetByCred(method, value string) (types.Uid, error) {
	panic("implement me")
}

func (a adapter) UserUnreadCount(uid types.Uid) (int, error) {
	panic("implement me")
}

func (a adapter) CredUpsert(cred *types.Credential) (bool, error) {
	panic("implement me")
}

func (a adapter) CredGetActive(uid types.Uid, method string) (*types.Credential, error) {
	panic("implement me")
}

func (a adapter) CredGetAll(uid types.Uid, method string, validatedOnly bool) ([]types.Credential, error) {
	panic("implement me")
}

func (a adapter) CredDel(uid types.Uid, method, value string) error {
	panic("implement me")
}

func (a adapter) CredConfirm(uid types.Uid, method string) error {
	panic("implement me")
}

func (a adapter) CredFail(uid types.Uid, method string) error {
	panic("implement me")
}

func (a adapter) AuthGetUniqueRecord(unique string) (types.Uid, auth.Level, []byte, time.Time, error) {
	panic("implement me")
}

func (a adapter) AuthGetRecord(user types.Uid, scheme string) (string, auth.Level, []byte, time.Time, error) {
	panic("implement me")
}

func (a adapter) AuthAddRecord(user types.Uid, scheme, unique string, authLvl auth.Level, secret []byte, expires time.Time) error {
	panic("implement me")
}

func (a adapter) AuthDelScheme(user types.Uid, scheme string) error {
	panic("implement me")
}

func (a adapter) AuthDelAllRecords(uid types.Uid) (int, error) {
	panic("implement me")
}

func (a adapter) AuthUpdRecord(user types.Uid, scheme, unique string, authLvl auth.Level, secret []byte, expires time.Time) error {
	panic("implement me")
}

func (a adapter) TopicCreate(topic *types.Topic) error {
	panic("implement me")
}

func (a adapter) TopicCreateP2P(initiator, invited *types.Subscription) error {
	panic("implement me")
}

func (a adapter) TopicGet(topic string) (*types.Topic, error) {
	panic("implement me")
}

func (a adapter) TopicsForUser(uid types.Uid, keepDeleted bool, opts *types.QueryOpt) ([]types.Subscription, error) {
	panic("implement me")
}

func (a adapter) UsersForTopic(topic string, keepDeleted bool, opts *types.QueryOpt) ([]types.Subscription, error) {
	panic("implement me")
}

func (a adapter) OwnTopics(uid types.Uid) ([]string, error) {
	panic("implement me")
}

func (a adapter) TopicShare(subs []*types.Subscription) error {
	panic("implement me")
}

func (a adapter) TopicDelete(topic string, hard bool) error {
	panic("implement me")
}

func (a adapter) TopicUpdateOnMessage(topic string, msg *types.Message) error {
	panic("implement me")
}

func (a adapter) TopicUpdate(topic string, update map[string]interface{}) error {
	panic("implement me")
}

func (a adapter) TopicOwnerChange(topic string, newOwner types.Uid) error {
	panic("implement me")
}

func (a adapter) SubscriptionGet(topic string, user types.Uid) (*types.Subscription, error) {
	panic("implement me")
}

func (a adapter) SubsForUser(user types.Uid, keepDeleted bool, opts *types.QueryOpt) ([]types.Subscription, error) {
	panic("implement me")
}

func (a adapter) SubsForTopic(topic string, keepDeleted bool, opts *types.QueryOpt) ([]types.Subscription, error) {
	panic("implement me")
}

func (a adapter) SubsUpdate(topic string, user types.Uid, update map[string]interface{}) error {
	panic("implement me")
}

func (a adapter) SubsDelete(topic string, user types.Uid) error {
	panic("implement me")
}

func (a adapter) SubsDelForTopic(topic string, hard bool) error {
	panic("implement me")
}

func (a adapter) SubsDelForUser(user types.Uid, hard bool) error {
	panic("implement me")
}

func (a adapter) FindUsers(user types.Uid, req, opt []string) ([]types.Subscription, error) {
	panic("implement me")
}

func (a adapter) FindTopics(req, opt []string) ([]types.Subscription, error) {
	panic("implement me")
}

func (a adapter) MessageSave(msg *types.Message) error {
	panic("implement me")
}

func (a adapter) MessageGetAll(topic string, forUser types.Uid, opts *types.QueryOpt) ([]types.Message, error) {
	panic("implement me")
}

func (a adapter) MessageDeleteList(topic string, toDel *types.DelMessage) error {
	panic("implement me")
}

func (a adapter) MessageGetDeleted(topic string, forUser types.Uid, opts *types.QueryOpt) ([]types.DelMessage, error) {
	panic("implement me")
}

func (a adapter) MessageAttachments(msgId types.Uid, fids []string) error {
	panic("implement me")
}

func (a adapter) DeviceUpsert(uid types.Uid, dev *types.DeviceDef) error {
	panic("implement me")
}

func (a adapter) DeviceGetAll(uid ...types.Uid) (map[types.Uid][]types.DeviceDef, int, error) {
	panic("implement me")
}

func (a adapter) DeviceDelete(uid types.Uid, deviceID string) error {
	panic("implement me")
}

func (a adapter) FileStartUpload(fd *types.FileDef) error {
	panic("implement me")
}

func (a adapter) FileFinishUpload(fid string, status int, size int64) (*types.FileDef, error) {
	panic("implement me")
}

func (a adapter) FileGet(fid string) (*types.FileDef, error) {
	panic("implement me")
}

func (a adapter) FileDeleteUnused(olderThan time.Time, limit int) ([]string, error) {
	panic("implement me")
}
