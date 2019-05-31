package pb

import (
	context "context"
	x "database/sql"

	sqrl "github.com/elgris/sqrl"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	chaku_globals "go.appointy.com/chaku/chaku-globals"
	driver "go.appointy.com/chaku/driver"
	sql "go.appointy.com/chaku/driver/sql"
	errors "go.appointy.com/chaku/errors"
)

var objectTableMap = chaku_globals.ObjectTable{
	"accouting_employee_link": {
		"id":          "accouting_employee_link",
		"appointy_id": "accouting_employee_link",
		"external_id": "accouting_employee_link",
		"app_type":    "accouting_employee_link",
		"metadata":    "accouting_employee_link",
	},
}

func (m *AccoutingEmployeeLink) PackageName() string {
	return "appointy_accouting_v1"
}

func (m *AccoutingEmployeeLink) TableOfObject(f, s string) string {
	return objectTableMap[f][s]
}

func (m *AccoutingEmployeeLink) GetDescriptorsOf(f string) (driver.Descriptor, error) {
	switch f {
	default:
		return nil, errors.ErrInvalidField
	}
}

func (m *AccoutingEmployeeLink) ObjectName() string {
	return "accouting_employee_link"
}

func (m *AccoutingEmployeeLink) Fields() []string {
	return []string{
		"id", "appointy_id", "external_id", "app_type", "metadata",
	}
}

func (m *AccoutingEmployeeLink) IsObject(field string) bool {
	switch field {
	default:
		return false
	}
}

func (m *AccoutingEmployeeLink) ValuerSlice(field string) ([]driver.Descriptor, error) {
	if m == nil {
		return nil, nil
	}
	switch field {
	default:
		return []driver.Descriptor{}, errors.ErrInvalidField
	}
}

func (m *AccoutingEmployeeLink) Valuer(field string) (interface{}, error) {
	if m == nil {
		return nil, nil
	}
	switch field {
	case "id":
		return m.Id, nil
	case "appointy_id":
		return m.AppointyId, nil
	case "external_id":
		return m.ExternalId, nil
	case "app_type":
		return m.AppType, nil
	case "metadata":
		return m.Metadata, nil
	default:
		return nil, errors.ErrInvalidField
	}
}

func (m *AccoutingEmployeeLink) Addresser(field string) (interface{}, error) {
	if m == nil {
		return nil, nil
	}
	switch field {
	case "id":
		return &m.Id, nil
	case "appointy_id":
		return &m.AppointyId, nil
	case "external_id":
		return &m.ExternalId, nil
	case "app_type":
		return &m.AppType, nil
	case "metadata":
		return &m.Metadata, nil
	default:
		return nil, errors.ErrInvalidField
	}
}

func (m *AccoutingEmployeeLink) New(field string) error {
	switch field {
	case "id":
		return nil
	case "appointy_id":
		return nil
	case "external_id":
		return nil
	case "app_type":
		return nil
	case "metadata":
		if m.Metadata == nil {
			m.Metadata = map[string]string{}
		}
		return nil
	default:
		return errors.ErrInvalidField
	}
}

func (m *AccoutingEmployeeLink) Type(field string) string {
	switch field {
	case "id":
		return "string"
	case "appointy_id":
		return "string"
	case "external_id":
		return "string"
	case "app_type":
		return "enum"
	case "metadata":
		return "map"
	default:
		return ""
	}
}

func (_ *AccoutingEmployeeLink) GetEmptyObject() (m *AccoutingEmployeeLink) {
	m = &AccoutingEmployeeLink{}
	return
}

func (m *AccoutingEmployeeLink) GetPrefix() string {
	return "ael"
}

func (m *AccoutingEmployeeLink) GetID() string {
	return m.Id
}

func (m *AccoutingEmployeeLink) SetID(id string) {
	m.Id = id
}

func (m *AccoutingEmployeeLink) IsRoot() bool {
	return true
}

func (m *AccoutingEmployeeLink) IsFlatObject(f string) bool {
	return false
}

type AccoutingEmployeeLinkStore struct {
	d driver.Driver
}

func NewAccoutingEmployeeLinkStore(d driver.Driver) AccoutingEmployeeLinkStore {
	return AccoutingEmployeeLinkStore{d: d}
}

func NewPostgresAccoutingEmployeeLinkStore(db *x.DB, usr driver.IUserInfo) AccoutingEmployeeLinkStore {
	return AccoutingEmployeeLinkStore{
		&sql.Sql{DB: db, UserInfo: usr, Placeholder: sqrl.Dollar},
	}
}

func (s AccoutingEmployeeLinkStore) StartTransaction(ctx context.Context) error {
	return s.d.StartTransaction(ctx)
}

func (s AccoutingEmployeeLinkStore) CommitTransaction(ctx context.Context) error {
	return s.d.CommitTransaction(ctx)
}

func (s AccoutingEmployeeLinkStore) RollBackTransaction(ctx context.Context) error {
	return s.d.RollBackTransaction(ctx)
}

func (s AccoutingEmployeeLinkStore) CreateAccoutingEmployeeLinkPGStore(ctx context.Context) error {
	const queries = `
CREATE SCHEMA IF NOT EXISTS appointy_accouting_v1;
CREATE TABLE IF NOT EXISTS appointy_accouting_v1.accouting_employee_link_parent (id text DEFAULT ''::text  , parent text DEFAULT ''::text  , is_deleted boolean DEFAULT false  , deleted_by text DEFAULT ''::text  , deleted_on timestamp without time zone DEFAULT '1970-01-01 00:00:00'::timestamp without time zone  , updated_by text DEFAULT ''::text  , updated_on timestamp without time zone DEFAULT '1970-01-01 00:00:00'::timestamp without time zone  , created_by text DEFAULT ''::text  , created_on timestamp without time zone DEFAULT '1970-01-01 00:00:00'::timestamp without time zone  );
CREATE TABLE IF NOT EXISTS appointy_accouting_v1.accouting_employee_link (id text DEFAULT ''::text  , appointy_id text DEFAULT ''::text  , external_id text DEFAULT ''::text  , app_type integer DEFAULT 0  , metadata jsonb DEFAULT '{}'::jsonb  , parent text DEFAULT ''::text  , is_deleted boolean DEFAULT false  , deleted_by text DEFAULT ''::text  , deleted_on timestamp without time zone DEFAULT '1970-01-01 00:00:00'::timestamp without time zone  , updated_by text DEFAULT ''::text  , updated_on timestamp without time zone DEFAULT '1970-01-01 00:00:00'::timestamp without time zone  , created_by text DEFAULT ''::text  , created_on timestamp without time zone DEFAULT '1970-01-01 00:00:00'::timestamp without time zone  , PRIMARY KEY(id) );
`
	if err := s.d.ExecuteQuery(ctx, queries); err != nil {
		return err
	}
	return nil
}

type OffsetBasedPagination struct {
	Offset int
	Limit  int
}

type CursorBasedPagination struct {
	// Set UpOrDown = true for getting list of data above Offset-ID,
	// limited to 'limit' amount, when ordered by ID in Ascending order.
	// Set UpOrDown = false for getting list of data below Offset-ID,
	// limited to 'limit' amount, when ordered by ID in Ascending order.
	Offset   string
	Limit    int
	UpOrDown bool
}

func (s AccoutingEmployeeLinkStore) CreateAccoutingEmployeeLinks(ctx context.Context, list ...*AccoutingEmployeeLink) ([]string, error) {
	vv := make([]driver.Descriptor, len(list))
	for i := range list {
		vv[i] = list[i]
	}
	return s.d.Insert(ctx, vv, &AccoutingEmployeeLink{}, &AccoutingEmployeeLink{}, "")
}

func (s AccoutingEmployeeLinkStore) DeleteAccoutingEmployeeLink(ctx context.Context, cond accoutingEmployeeLinkCondition) error {
	return s.d.Delete(ctx, cond.accoutingEmployeeLinkCondToDriverCond(s.d), &AccoutingEmployeeLink{}, &AccoutingEmployeeLink{})
}

func (s AccoutingEmployeeLinkStore) UpdateAccoutingEmployeeLink(ctx context.Context, req *AccoutingEmployeeLink, fields []string, cond accoutingEmployeeLinkCondition) error {
	return s.d.Update(ctx, cond.accoutingEmployeeLinkCondToDriverCond(s.d), req, &AccoutingEmployeeLink{}, fields...)
}

func (s AccoutingEmployeeLinkStore) GetAccoutingEmployeeLink(ctx context.Context, fields []string, cond accoutingEmployeeLinkCondition) (*AccoutingEmployeeLink, error) {
	if len(fields) == 0 {
		fields = (&AccoutingEmployeeLink{}).Fields()
	}
	objList, err := s.ListAccoutingEmployeeLinks(ctx, fields, cond)
	if len(objList) == 0 && err == nil {
		err = errors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return objList[0], nil
}

func (s AccoutingEmployeeLinkStore) ListAccoutingEmployeeLinks(ctx context.Context, fields []string, cond accoutingEmployeeLinkCondition, opt ...listAccoutingEmployeeLinksOption) ([]*AccoutingEmployeeLink, error) {
	if len(fields) == 0 {
		fields = (&AccoutingEmployeeLink{}).Fields()
	}
	for _, o := range opt {
		var err error
		cond, err = o.applyOffsetToAccoutingEmployeeLinksList(ctx, s.d, cond)
		if err != nil {
			return nil, err
		}
	}
	res, err := s.d.Get(ctx, cond.accoutingEmployeeLinkCondToDriverCond(s.d), &AccoutingEmployeeLink{}, &AccoutingEmployeeLink{}, fields...)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	list := make([]*AccoutingEmployeeLink, 0, 1000)
	for res.Next(ctx) {
		obj := &AccoutingEmployeeLink{}
		if err := res.Scan(obj); err != nil {
			return nil, err
		}
		list = append(list, obj)
	}

	if err := res.Close(); err != nil {
		return nil, err
	}

	for _, o := range opt {
		list = o.applyLimitToAccoutingEmployeeLinksList(MapperAccoutingEmployeeLink(list))
	}
	return list, nil
}

type listAccoutingEmployeeLinksOption interface {
	listAccoutingEmployeeLinks() // method of no use
	applyOffsetToAccoutingEmployeeLinksList(context.Context, driver.Driver, accoutingEmployeeLinkCondition) (accoutingEmployeeLinkCondition, error)
	applyLimitToAccoutingEmployeeLinksList([]*AccoutingEmployeeLink) []*AccoutingEmployeeLink // temporary
}

func (*OffsetBasedPagination) listAccoutingEmployeeLinks() {
}

func (p *OffsetBasedPagination) applyLimitToAccoutingEmployeeLinksList(ls []*AccoutingEmployeeLink) []*AccoutingEmployeeLink {
	if len(ls) <= p.Limit {
		return ls
	}
	return ls[0:p.Limit]
}

func (*CursorBasedPagination) listAccoutingEmployeeLinks() {
}

func (p *CursorBasedPagination) applyLimitToAccoutingEmployeeLinksList(ls []*AccoutingEmployeeLink) []*AccoutingEmployeeLink {
	if len(ls) <= p.Limit {
		return ls
	}
	if p.UpOrDown {
		return ls[len(ls)-p.Limit:]
	} else {
		return ls[:p.Limit]
	}
}

func (p *OffsetBasedPagination) applyOffsetToAccoutingEmployeeLinksList(ctx context.Context, d driver.Driver, cond accoutingEmployeeLinkCondition) (accoutingEmployeeLinkCondition, error) {
	offID, err := d.OffsetValue(ctx, &AccoutingEmployeeLink{}, &AccoutingEmployeeLink{}, p.Offset)
	if err != nil {
		return cond, err
	}
	return AccoutingEmployeeLinkAnd{cond, AccoutingEmployeeLinkIdGtOrEq{offID}}, nil
}

func (p *CursorBasedPagination) applyOffsetToAccoutingEmployeeLinksList(ctx context.Context, d driver.Driver, cond accoutingEmployeeLinkCondition) (accoutingEmployeeLinkCondition, error) {
	if p.UpOrDown {
		return AccoutingEmployeeLinkAnd{cond, AccoutingEmployeeLinkIdLt{p.Offset}}, nil
	} else {
		return AccoutingEmployeeLinkAnd{cond, AccoutingEmployeeLinkIdGt{p.Offset}}, nil
	}
}

type TrueCondition struct{}

type accoutingEmployeeLinkCondition interface {
	accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner
}

type AccoutingEmployeeLinkAnd []accoutingEmployeeLinkCondition

func (p AccoutingEmployeeLinkAnd) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	dc := make([]driver.Conditioner, 0, len(p))
	for _, c := range p {
		dc = append(dc, c.accoutingEmployeeLinkCondToDriverCond(d))
	}
	return driver.And{Conditioners: dc, Operator: d}
}

type AccoutingEmployeeLinkOr []accoutingEmployeeLinkCondition

func (p AccoutingEmployeeLinkOr) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	dc := make([]driver.Conditioner, 0, len(p))
	for _, c := range p {
		dc = append(dc, c.accoutingEmployeeLinkCondToDriverCond(d))
	}
	return driver.And{Conditioners: dc, Operator: d}
}

type AccoutingEmployeeLinkParentEq struct {
	Parent string
}

func (c AccoutingEmployeeLinkParentEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "parent", Value: c.Parent, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkIdEq struct {
	Id string
}

func (c AccoutingEmployeeLinkIdEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "id", Value: c.Id, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "id", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkAppointyIdEq struct {
	AppointyId string
}

func (c AccoutingEmployeeLinkAppointyIdEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "appointy_id", Value: c.AppointyId, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "appointy_id", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkExternalIdEq struct {
	ExternalId string
}

func (c AccoutingEmployeeLinkExternalIdEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "external_id", Value: c.ExternalId, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "external_id", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkAppTypeEq struct {
	AppType AccountingIntegrationType
}

func (c AccoutingEmployeeLinkAppTypeEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "app_type", Value: c.AppType, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "app_type", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkMetadataEq struct {
	Metadata map[string]string
}

func (c AccoutingEmployeeLinkMetadataEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "metadata", Value: c.Metadata, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "metadata", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkIdNotEq struct {
	Id string
}

func (c AccoutingEmployeeLinkIdNotEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "id", Value: c.Id, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "id", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkAppointyIdNotEq struct {
	AppointyId string
}

func (c AccoutingEmployeeLinkAppointyIdNotEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "appointy_id", Value: c.AppointyId, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "appointy_id", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkExternalIdNotEq struct {
	ExternalId string
}

func (c AccoutingEmployeeLinkExternalIdNotEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "external_id", Value: c.ExternalId, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "external_id", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkAppTypeNotEq struct {
	AppType AccountingIntegrationType
}

func (c AccoutingEmployeeLinkAppTypeNotEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "app_type", Value: c.AppType, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "app_type", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkMetadataNotEq struct {
	Metadata map[string]string
}

func (c AccoutingEmployeeLinkMetadataNotEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "metadata", Value: c.Metadata, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "metadata", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkIdGt struct {
	Id string
}

func (c AccoutingEmployeeLinkIdGt) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "id", Value: c.Id, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "id", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkAppointyIdGt struct {
	AppointyId string
}

func (c AccoutingEmployeeLinkAppointyIdGt) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "appointy_id", Value: c.AppointyId, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "appointy_id", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkExternalIdGt struct {
	ExternalId string
}

func (c AccoutingEmployeeLinkExternalIdGt) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "external_id", Value: c.ExternalId, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "external_id", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkAppTypeGt struct {
	AppType AccountingIntegrationType
}

func (c AccoutingEmployeeLinkAppTypeGt) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "app_type", Value: c.AppType, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "app_type", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkMetadataGt struct {
	Metadata map[string]string
}

func (c AccoutingEmployeeLinkMetadataGt) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "metadata", Value: c.Metadata, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "metadata", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkIdLt struct {
	Id string
}

func (c AccoutingEmployeeLinkIdLt) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "id", Value: c.Id, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "id", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkAppointyIdLt struct {
	AppointyId string
}

func (c AccoutingEmployeeLinkAppointyIdLt) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "appointy_id", Value: c.AppointyId, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "appointy_id", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkExternalIdLt struct {
	ExternalId string
}

func (c AccoutingEmployeeLinkExternalIdLt) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "external_id", Value: c.ExternalId, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "external_id", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkAppTypeLt struct {
	AppType AccountingIntegrationType
}

func (c AccoutingEmployeeLinkAppTypeLt) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "app_type", Value: c.AppType, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "app_type", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkMetadataLt struct {
	Metadata map[string]string
}

func (c AccoutingEmployeeLinkMetadataLt) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "metadata", Value: c.Metadata, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "metadata", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkIdGtOrEq struct {
	Id string
}

func (c AccoutingEmployeeLinkIdGtOrEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "id", Value: c.Id, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "id", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkAppointyIdGtOrEq struct {
	AppointyId string
}

func (c AccoutingEmployeeLinkAppointyIdGtOrEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "appointy_id", Value: c.AppointyId, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "appointy_id", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkExternalIdGtOrEq struct {
	ExternalId string
}

func (c AccoutingEmployeeLinkExternalIdGtOrEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "external_id", Value: c.ExternalId, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "external_id", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkAppTypeGtOrEq struct {
	AppType AccountingIntegrationType
}

func (c AccoutingEmployeeLinkAppTypeGtOrEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "app_type", Value: c.AppType, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "app_type", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkMetadataGtOrEq struct {
	Metadata map[string]string
}

func (c AccoutingEmployeeLinkMetadataGtOrEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "metadata", Value: c.Metadata, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "metadata", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkIdLtOrEq struct {
	Id string
}

func (c AccoutingEmployeeLinkIdLtOrEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "id", Value: c.Id, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "id", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkAppointyIdLtOrEq struct {
	AppointyId string
}

func (c AccoutingEmployeeLinkAppointyIdLtOrEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "appointy_id", Value: c.AppointyId, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "appointy_id", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkExternalIdLtOrEq struct {
	ExternalId string
}

func (c AccoutingEmployeeLinkExternalIdLtOrEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "external_id", Value: c.ExternalId, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "external_id", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkAppTypeLtOrEq struct {
	AppType AccountingIntegrationType
}

func (c AccoutingEmployeeLinkAppTypeLtOrEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "app_type", Value: c.AppType, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "app_type", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkMetadataLtOrEq struct {
	Metadata map[string]string
}

func (c AccoutingEmployeeLinkMetadataLtOrEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "metadata", Value: c.Metadata, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "metadata", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkDeleted struct{}

func (c AccoutingEmployeeLinkDeleted) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "is_deleted", Value: true, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkNotDeleted struct{}

func (c AccoutingEmployeeLinkNotDeleted) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "is_deleted", Value: false, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkCreatedByEq struct {
	By string
}

func (c AccoutingEmployeeLinkCreatedByEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkCreatedOnEq struct {
	On *timestamp.Timestamp
}

func (c AccoutingEmployeeLinkCreatedOnEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkCreatedByNotEq struct {
	By string
}

func (c AccoutingEmployeeLinkCreatedByNotEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkCreatedOnNotEq struct {
	On *timestamp.Timestamp
}

func (c AccoutingEmployeeLinkCreatedOnNotEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkCreatedByGt struct {
	By string
}

func (c AccoutingEmployeeLinkCreatedByGt) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "created_by", Value: c.By, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkCreatedOnGt struct {
	On *timestamp.Timestamp
}

func (c AccoutingEmployeeLinkCreatedOnGt) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "created_on", Value: c.On, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkCreatedByLt struct {
	By string
}

func (c AccoutingEmployeeLinkCreatedByLt) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "created_by", Value: c.By, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkCreatedOnLt struct {
	On *timestamp.Timestamp
}

func (c AccoutingEmployeeLinkCreatedOnLt) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "created_on", Value: c.On, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkCreatedByGtOrEq struct {
	By string
}

func (c AccoutingEmployeeLinkCreatedByGtOrEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkCreatedOnGtOrEq struct {
	On *timestamp.Timestamp
}

func (c AccoutingEmployeeLinkCreatedOnGtOrEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkCreatedByLtOrEq struct {
	By string
}

func (c AccoutingEmployeeLinkCreatedByLtOrEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkCreatedOnLtOrEq struct {
	On *timestamp.Timestamp
}

func (c AccoutingEmployeeLinkCreatedOnLtOrEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkUpdatedByEq struct {
	By string
}

func (c AccoutingEmployeeLinkUpdatedByEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkUpdatedOnEq struct {
	On *timestamp.Timestamp
}

func (c AccoutingEmployeeLinkUpdatedOnEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkUpdatedByNotEq struct {
	By string
}

func (c AccoutingEmployeeLinkUpdatedByNotEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkUpdatedOnNotEq struct {
	On *timestamp.Timestamp
}

func (c AccoutingEmployeeLinkUpdatedOnNotEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkUpdatedByGt struct {
	By string
}

func (c AccoutingEmployeeLinkUpdatedByGt) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkUpdatedOnGt struct {
	On *timestamp.Timestamp
}

func (c AccoutingEmployeeLinkUpdatedOnGt) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkUpdatedByLt struct {
	By string
}

func (c AccoutingEmployeeLinkUpdatedByLt) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkUpdatedOnLt struct {
	On *timestamp.Timestamp
}

func (c AccoutingEmployeeLinkUpdatedOnLt) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkUpdatedByGtOrEq struct {
	By string
}

func (c AccoutingEmployeeLinkUpdatedByGtOrEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkUpdatedOnGtOrEq struct {
	On *timestamp.Timestamp
}

func (c AccoutingEmployeeLinkUpdatedOnGtOrEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkUpdatedByLtOrEq struct {
	By string
}

func (c AccoutingEmployeeLinkUpdatedByLtOrEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkUpdatedOnLtOrEq struct {
	On *timestamp.Timestamp
}

func (c AccoutingEmployeeLinkUpdatedOnLtOrEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkDeletedByEq struct {
	By string
}

func (c AccoutingEmployeeLinkDeletedByEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkDeletedOnEq struct {
	On *timestamp.Timestamp
}

func (c AccoutingEmployeeLinkDeletedOnEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkDeletedByNotEq struct {
	By string
}

func (c AccoutingEmployeeLinkDeletedByNotEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkDeletedOnNotEq struct {
	On *timestamp.Timestamp
}

func (c AccoutingEmployeeLinkDeletedOnNotEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkDeletedByGt struct {
	By string
}

func (c AccoutingEmployeeLinkDeletedByGt) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkDeletedOnGt struct {
	On *timestamp.Timestamp
}

func (c AccoutingEmployeeLinkDeletedOnGt) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkDeletedByLt struct {
	By string
}

func (c AccoutingEmployeeLinkDeletedByLt) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkDeletedOnLt struct {
	On *timestamp.Timestamp
}

func (c AccoutingEmployeeLinkDeletedOnLt) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkDeletedByGtOrEq struct {
	By string
}

func (c AccoutingEmployeeLinkDeletedByGtOrEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkDeletedOnGtOrEq struct {
	On *timestamp.Timestamp
}

func (c AccoutingEmployeeLinkDeletedOnGtOrEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkDeletedByLtOrEq struct {
	By string
}

func (c AccoutingEmployeeLinkDeletedByLtOrEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkDeletedOnLtOrEq struct {
	On *timestamp.Timestamp
}

func (c AccoutingEmployeeLinkDeletedOnLtOrEq) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &AccoutingEmployeeLink{}, RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkIdIn struct {
	Id []string
}

func (c AccoutingEmployeeLinkIdIn) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "id", Value: c.Id, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "id", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkAppointyIdIn struct {
	AppointyId []string
}

func (c AccoutingEmployeeLinkAppointyIdIn) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "appointy_id", Value: c.AppointyId, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "appointy_id", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkExternalIdIn struct {
	ExternalId []string
}

func (c AccoutingEmployeeLinkExternalIdIn) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "external_id", Value: c.ExternalId, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "external_id", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkAppTypeIn struct {
	AppType []AccountingIntegrationType
}

func (c AccoutingEmployeeLinkAppTypeIn) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "app_type", Value: c.AppType, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "app_type", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkIdNotIn struct {
	Id []string
}

func (c AccoutingEmployeeLinkIdNotIn) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "id", Value: c.Id, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "id", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkAppointyIdNotIn struct {
	AppointyId []string
}

func (c AccoutingEmployeeLinkAppointyIdNotIn) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "appointy_id", Value: c.AppointyId, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "appointy_id", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkExternalIdNotIn struct {
	ExternalId []string
}

func (c AccoutingEmployeeLinkExternalIdNotIn) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "external_id", Value: c.ExternalId, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "external_id", RootDescriptor: &AccoutingEmployeeLink{}}
}

type AccoutingEmployeeLinkAppTypeNotIn struct {
	AppType []AccountingIntegrationType
}

func (c AccoutingEmployeeLinkAppTypeNotIn) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "app_type", Value: c.AppType, Operator: d, Descriptor: &AccoutingEmployeeLink{}, FieldMask: "app_type", RootDescriptor: &AccoutingEmployeeLink{}}
}

func (c TrueCondition) accoutingEmployeeLinkCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.TrueCondition{Operator: d}
}

type accoutingEmployeeLinkMapperObject struct {
	id         string
	appointyId string
	externalId string
	appType    AccountingIntegrationType
	metadata   map[string]string
}

func (s *accoutingEmployeeLinkMapperObject) GetUniqueIdentifier() string {
	return s.id
}

func MapperAccoutingEmployeeLink(rows []*AccoutingEmployeeLink) []*AccoutingEmployeeLink {

	ids := make([]string, 0, len(rows))
	for _, r := range rows {
		ids = append(ids, r.Id)
	}

	combinedAccoutingEmployeeLinkMappers := map[string]*accoutingEmployeeLinkMapperObject{}

	for _, rw := range rows {

		tempAccoutingEmployeeLink := &accoutingEmployeeLinkMapperObject{}

		if rw == nil {
			rw = rw.GetEmptyObject()
		}
		tempAccoutingEmployeeLink.id = rw.Id
		tempAccoutingEmployeeLink.appointyId = rw.AppointyId
		tempAccoutingEmployeeLink.externalId = rw.ExternalId
		tempAccoutingEmployeeLink.appType = rw.AppType
		tempAccoutingEmployeeLink.metadata = rw.Metadata

		if combinedAccoutingEmployeeLinkMappers[tempAccoutingEmployeeLink.GetUniqueIdentifier()] == nil {
			combinedAccoutingEmployeeLinkMappers[tempAccoutingEmployeeLink.GetUniqueIdentifier()] = tempAccoutingEmployeeLink
		}
	}

	combinedAccoutingEmployeeLinks := make(map[string]*AccoutingEmployeeLink, 0)

	for _, accoutingEmployeeLink := range combinedAccoutingEmployeeLinkMappers {
		tempAccoutingEmployeeLink := &AccoutingEmployeeLink{}
		tempAccoutingEmployeeLink.Id = accoutingEmployeeLink.id
		tempAccoutingEmployeeLink.AppointyId = accoutingEmployeeLink.appointyId
		tempAccoutingEmployeeLink.ExternalId = accoutingEmployeeLink.externalId
		tempAccoutingEmployeeLink.AppType = accoutingEmployeeLink.appType
		tempAccoutingEmployeeLink.Metadata = accoutingEmployeeLink.metadata

		if tempAccoutingEmployeeLink.Id == "" {
			continue
		}

		combinedAccoutingEmployeeLinks[tempAccoutingEmployeeLink.Id] = tempAccoutingEmployeeLink

	}
	list := make([]*AccoutingEmployeeLink, 0, len(combinedAccoutingEmployeeLinks))
	for _, i := range ids {
		list = append(list, combinedAccoutingEmployeeLinks[i])
	}
	return list
}
