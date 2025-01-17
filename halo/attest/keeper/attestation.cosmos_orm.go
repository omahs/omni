// Code generated by protoc-gen-go-cosmos-orm. DO NOT EDIT.

package keeper

import (
	context "context"
	ormlist "cosmossdk.io/orm/model/ormlist"
	ormtable "cosmossdk.io/orm/model/ormtable"
	ormerrors "cosmossdk.io/orm/types/ormerrors"
)

type AttestationTable interface {
	Insert(ctx context.Context, attestation *Attestation) error
	InsertReturningId(ctx context.Context, attestation *Attestation) (uint64, error)
	LastInsertedSequence(ctx context.Context) (uint64, error)
	Update(ctx context.Context, attestation *Attestation) error
	Save(ctx context.Context, attestation *Attestation) error
	Delete(ctx context.Context, attestation *Attestation) error
	Has(ctx context.Context, id uint64) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, id uint64) (*Attestation, error)
	HasByChainIdConfLevelOffsetHeightHashAttestationRoot(ctx context.Context, chain_id uint64, conf_level uint32, offset uint64, height uint64, hash []byte, attestation_root []byte) (found bool, err error)
	// GetByChainIdConfLevelOffsetHeightHashAttestationRoot returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByChainIdConfLevelOffsetHeightHashAttestationRoot(ctx context.Context, chain_id uint64, conf_level uint32, offset uint64, height uint64, hash []byte, attestation_root []byte) (*Attestation, error)
	List(ctx context.Context, prefixKey AttestationIndexKey, opts ...ormlist.Option) (AttestationIterator, error)
	ListRange(ctx context.Context, from, to AttestationIndexKey, opts ...ormlist.Option) (AttestationIterator, error)
	DeleteBy(ctx context.Context, prefixKey AttestationIndexKey) error
	DeleteRange(ctx context.Context, from, to AttestationIndexKey) error

	doNotImplement()
}

type AttestationIterator struct {
	ormtable.Iterator
}

func (i AttestationIterator) Value() (*Attestation, error) {
	var attestation Attestation
	err := i.UnmarshalMessage(&attestation)
	return &attestation, err
}

type AttestationIndexKey interface {
	id() uint32
	values() []interface{}
	attestationIndexKey()
}

// primary key starting index..
type AttestationPrimaryKey = AttestationIdIndexKey

type AttestationIdIndexKey struct {
	vs []interface{}
}

func (x AttestationIdIndexKey) id() uint32            { return 0 }
func (x AttestationIdIndexKey) values() []interface{} { return x.vs }
func (x AttestationIdIndexKey) attestationIndexKey()  {}

func (this AttestationIdIndexKey) WithId(id uint64) AttestationIdIndexKey {
	this.vs = []interface{}{id}
	return this
}

type AttestationChainIdConfLevelOffsetHeightHashAttestationRootIndexKey struct {
	vs []interface{}
}

func (x AttestationChainIdConfLevelOffsetHeightHashAttestationRootIndexKey) id() uint32 { return 1 }
func (x AttestationChainIdConfLevelOffsetHeightHashAttestationRootIndexKey) values() []interface{} {
	return x.vs
}
func (x AttestationChainIdConfLevelOffsetHeightHashAttestationRootIndexKey) attestationIndexKey() {}

func (this AttestationChainIdConfLevelOffsetHeightHashAttestationRootIndexKey) WithChainId(chain_id uint64) AttestationChainIdConfLevelOffsetHeightHashAttestationRootIndexKey {
	this.vs = []interface{}{chain_id}
	return this
}

func (this AttestationChainIdConfLevelOffsetHeightHashAttestationRootIndexKey) WithChainIdConfLevel(chain_id uint64, conf_level uint32) AttestationChainIdConfLevelOffsetHeightHashAttestationRootIndexKey {
	this.vs = []interface{}{chain_id, conf_level}
	return this
}

func (this AttestationChainIdConfLevelOffsetHeightHashAttestationRootIndexKey) WithChainIdConfLevelOffset(chain_id uint64, conf_level uint32, offset uint64) AttestationChainIdConfLevelOffsetHeightHashAttestationRootIndexKey {
	this.vs = []interface{}{chain_id, conf_level, offset}
	return this
}

func (this AttestationChainIdConfLevelOffsetHeightHashAttestationRootIndexKey) WithChainIdConfLevelOffsetHeight(chain_id uint64, conf_level uint32, offset uint64, height uint64) AttestationChainIdConfLevelOffsetHeightHashAttestationRootIndexKey {
	this.vs = []interface{}{chain_id, conf_level, offset, height}
	return this
}

func (this AttestationChainIdConfLevelOffsetHeightHashAttestationRootIndexKey) WithChainIdConfLevelOffsetHeightHash(chain_id uint64, conf_level uint32, offset uint64, height uint64, hash []byte) AttestationChainIdConfLevelOffsetHeightHashAttestationRootIndexKey {
	this.vs = []interface{}{chain_id, conf_level, offset, height, hash}
	return this
}

func (this AttestationChainIdConfLevelOffsetHeightHashAttestationRootIndexKey) WithChainIdConfLevelOffsetHeightHashAttestationRoot(chain_id uint64, conf_level uint32, offset uint64, height uint64, hash []byte, attestation_root []byte) AttestationChainIdConfLevelOffsetHeightHashAttestationRootIndexKey {
	this.vs = []interface{}{chain_id, conf_level, offset, height, hash, attestation_root}
	return this
}

type AttestationStatusChainIdConfLevelOffsetIndexKey struct {
	vs []interface{}
}

func (x AttestationStatusChainIdConfLevelOffsetIndexKey) id() uint32            { return 2 }
func (x AttestationStatusChainIdConfLevelOffsetIndexKey) values() []interface{} { return x.vs }
func (x AttestationStatusChainIdConfLevelOffsetIndexKey) attestationIndexKey()  {}

func (this AttestationStatusChainIdConfLevelOffsetIndexKey) WithStatus(status uint32) AttestationStatusChainIdConfLevelOffsetIndexKey {
	this.vs = []interface{}{status}
	return this
}

func (this AttestationStatusChainIdConfLevelOffsetIndexKey) WithStatusChainId(status uint32, chain_id uint64) AttestationStatusChainIdConfLevelOffsetIndexKey {
	this.vs = []interface{}{status, chain_id}
	return this
}

func (this AttestationStatusChainIdConfLevelOffsetIndexKey) WithStatusChainIdConfLevel(status uint32, chain_id uint64, conf_level uint32) AttestationStatusChainIdConfLevelOffsetIndexKey {
	this.vs = []interface{}{status, chain_id, conf_level}
	return this
}

func (this AttestationStatusChainIdConfLevelOffsetIndexKey) WithStatusChainIdConfLevelOffset(status uint32, chain_id uint64, conf_level uint32, offset uint64) AttestationStatusChainIdConfLevelOffsetIndexKey {
	this.vs = []interface{}{status, chain_id, conf_level, offset}
	return this
}

type AttestationCreatedHeightIndexKey struct {
	vs []interface{}
}

func (x AttestationCreatedHeightIndexKey) id() uint32            { return 3 }
func (x AttestationCreatedHeightIndexKey) values() []interface{} { return x.vs }
func (x AttestationCreatedHeightIndexKey) attestationIndexKey()  {}

func (this AttestationCreatedHeightIndexKey) WithCreatedHeight(created_height uint64) AttestationCreatedHeightIndexKey {
	this.vs = []interface{}{created_height}
	return this
}

type attestationTable struct {
	table ormtable.AutoIncrementTable
}

func (this attestationTable) Insert(ctx context.Context, attestation *Attestation) error {
	return this.table.Insert(ctx, attestation)
}

func (this attestationTable) Update(ctx context.Context, attestation *Attestation) error {
	return this.table.Update(ctx, attestation)
}

func (this attestationTable) Save(ctx context.Context, attestation *Attestation) error {
	return this.table.Save(ctx, attestation)
}

func (this attestationTable) Delete(ctx context.Context, attestation *Attestation) error {
	return this.table.Delete(ctx, attestation)
}

func (this attestationTable) InsertReturningId(ctx context.Context, attestation *Attestation) (uint64, error) {
	return this.table.InsertReturningPKey(ctx, attestation)
}

func (this attestationTable) LastInsertedSequence(ctx context.Context) (uint64, error) {
	return this.table.LastInsertedSequence(ctx)
}

func (this attestationTable) Has(ctx context.Context, id uint64) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, id)
}

func (this attestationTable) Get(ctx context.Context, id uint64) (*Attestation, error) {
	var attestation Attestation
	found, err := this.table.PrimaryKey().Get(ctx, &attestation, id)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &attestation, nil
}

func (this attestationTable) HasByChainIdConfLevelOffsetHeightHashAttestationRoot(ctx context.Context, chain_id uint64, conf_level uint32, offset uint64, height uint64, hash []byte, attestation_root []byte) (found bool, err error) {
	return this.table.GetIndexByID(1).(ormtable.UniqueIndex).Has(ctx,
		chain_id,
		conf_level,
		offset,
		height,
		hash,
		attestation_root,
	)
}

func (this attestationTable) GetByChainIdConfLevelOffsetHeightHashAttestationRoot(ctx context.Context, chain_id uint64, conf_level uint32, offset uint64, height uint64, hash []byte, attestation_root []byte) (*Attestation, error) {
	var attestation Attestation
	found, err := this.table.GetIndexByID(1).(ormtable.UniqueIndex).Get(ctx, &attestation,
		chain_id,
		conf_level,
		offset,
		height,
		hash,
		attestation_root,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &attestation, nil
}

func (this attestationTable) List(ctx context.Context, prefixKey AttestationIndexKey, opts ...ormlist.Option) (AttestationIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return AttestationIterator{it}, err
}

func (this attestationTable) ListRange(ctx context.Context, from, to AttestationIndexKey, opts ...ormlist.Option) (AttestationIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return AttestationIterator{it}, err
}

func (this attestationTable) DeleteBy(ctx context.Context, prefixKey AttestationIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this attestationTable) DeleteRange(ctx context.Context, from, to AttestationIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this attestationTable) doNotImplement() {}

var _ AttestationTable = attestationTable{}

func NewAttestationTable(db ormtable.Schema) (AttestationTable, error) {
	table := db.GetTable(&Attestation{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Attestation{}).ProtoReflect().Descriptor().FullName()))
	}
	return attestationTable{table.(ormtable.AutoIncrementTable)}, nil
}

type SignatureTable interface {
	Insert(ctx context.Context, signature *Signature) error
	InsertReturningId(ctx context.Context, signature *Signature) (uint64, error)
	LastInsertedSequence(ctx context.Context) (uint64, error)
	Update(ctx context.Context, signature *Signature) error
	Save(ctx context.Context, signature *Signature) error
	Delete(ctx context.Context, signature *Signature) error
	Has(ctx context.Context, id uint64) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, id uint64) (*Signature, error)
	HasByAttIdValidatorAddress(ctx context.Context, att_id uint64, validator_address []byte) (found bool, err error)
	// GetByAttIdValidatorAddress returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByAttIdValidatorAddress(ctx context.Context, att_id uint64, validator_address []byte) (*Signature, error)
	List(ctx context.Context, prefixKey SignatureIndexKey, opts ...ormlist.Option) (SignatureIterator, error)
	ListRange(ctx context.Context, from, to SignatureIndexKey, opts ...ormlist.Option) (SignatureIterator, error)
	DeleteBy(ctx context.Context, prefixKey SignatureIndexKey) error
	DeleteRange(ctx context.Context, from, to SignatureIndexKey) error

	doNotImplement()
}

type SignatureIterator struct {
	ormtable.Iterator
}

func (i SignatureIterator) Value() (*Signature, error) {
	var signature Signature
	err := i.UnmarshalMessage(&signature)
	return &signature, err
}

type SignatureIndexKey interface {
	id() uint32
	values() []interface{}
	signatureIndexKey()
}

// primary key starting index..
type SignaturePrimaryKey = SignatureIdIndexKey

type SignatureIdIndexKey struct {
	vs []interface{}
}

func (x SignatureIdIndexKey) id() uint32            { return 0 }
func (x SignatureIdIndexKey) values() []interface{} { return x.vs }
func (x SignatureIdIndexKey) signatureIndexKey()    {}

func (this SignatureIdIndexKey) WithId(id uint64) SignatureIdIndexKey {
	this.vs = []interface{}{id}
	return this
}

type SignatureAttIdValidatorAddressIndexKey struct {
	vs []interface{}
}

func (x SignatureAttIdValidatorAddressIndexKey) id() uint32            { return 1 }
func (x SignatureAttIdValidatorAddressIndexKey) values() []interface{} { return x.vs }
func (x SignatureAttIdValidatorAddressIndexKey) signatureIndexKey()    {}

func (this SignatureAttIdValidatorAddressIndexKey) WithAttId(att_id uint64) SignatureAttIdValidatorAddressIndexKey {
	this.vs = []interface{}{att_id}
	return this
}

func (this SignatureAttIdValidatorAddressIndexKey) WithAttIdValidatorAddress(att_id uint64, validator_address []byte) SignatureAttIdValidatorAddressIndexKey {
	this.vs = []interface{}{att_id, validator_address}
	return this
}

type SignatureAttIdIndexKey struct {
	vs []interface{}
}

func (x SignatureAttIdIndexKey) id() uint32            { return 2 }
func (x SignatureAttIdIndexKey) values() []interface{} { return x.vs }
func (x SignatureAttIdIndexKey) signatureIndexKey()    {}

func (this SignatureAttIdIndexKey) WithAttId(att_id uint64) SignatureAttIdIndexKey {
	this.vs = []interface{}{att_id}
	return this
}

type signatureTable struct {
	table ormtable.AutoIncrementTable
}

func (this signatureTable) Insert(ctx context.Context, signature *Signature) error {
	return this.table.Insert(ctx, signature)
}

func (this signatureTable) Update(ctx context.Context, signature *Signature) error {
	return this.table.Update(ctx, signature)
}

func (this signatureTable) Save(ctx context.Context, signature *Signature) error {
	return this.table.Save(ctx, signature)
}

func (this signatureTable) Delete(ctx context.Context, signature *Signature) error {
	return this.table.Delete(ctx, signature)
}

func (this signatureTable) InsertReturningId(ctx context.Context, signature *Signature) (uint64, error) {
	return this.table.InsertReturningPKey(ctx, signature)
}

func (this signatureTable) LastInsertedSequence(ctx context.Context) (uint64, error) {
	return this.table.LastInsertedSequence(ctx)
}

func (this signatureTable) Has(ctx context.Context, id uint64) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, id)
}

func (this signatureTable) Get(ctx context.Context, id uint64) (*Signature, error) {
	var signature Signature
	found, err := this.table.PrimaryKey().Get(ctx, &signature, id)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &signature, nil
}

func (this signatureTable) HasByAttIdValidatorAddress(ctx context.Context, att_id uint64, validator_address []byte) (found bool, err error) {
	return this.table.GetIndexByID(1).(ormtable.UniqueIndex).Has(ctx,
		att_id,
		validator_address,
	)
}

func (this signatureTable) GetByAttIdValidatorAddress(ctx context.Context, att_id uint64, validator_address []byte) (*Signature, error) {
	var signature Signature
	found, err := this.table.GetIndexByID(1).(ormtable.UniqueIndex).Get(ctx, &signature,
		att_id,
		validator_address,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &signature, nil
}

func (this signatureTable) List(ctx context.Context, prefixKey SignatureIndexKey, opts ...ormlist.Option) (SignatureIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return SignatureIterator{it}, err
}

func (this signatureTable) ListRange(ctx context.Context, from, to SignatureIndexKey, opts ...ormlist.Option) (SignatureIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return SignatureIterator{it}, err
}

func (this signatureTable) DeleteBy(ctx context.Context, prefixKey SignatureIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this signatureTable) DeleteRange(ctx context.Context, from, to SignatureIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this signatureTable) doNotImplement() {}

var _ SignatureTable = signatureTable{}

func NewSignatureTable(db ormtable.Schema) (SignatureTable, error) {
	table := db.GetTable(&Signature{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Signature{}).ProtoReflect().Descriptor().FullName()))
	}
	return signatureTable{table.(ormtable.AutoIncrementTable)}, nil
}

type AttestationStore interface {
	AttestationTable() AttestationTable
	SignatureTable() SignatureTable

	doNotImplement()
}

type attestationStore struct {
	attestation AttestationTable
	signature   SignatureTable
}

func (x attestationStore) AttestationTable() AttestationTable {
	return x.attestation
}

func (x attestationStore) SignatureTable() SignatureTable {
	return x.signature
}

func (attestationStore) doNotImplement() {}

var _ AttestationStore = attestationStore{}

func NewAttestationStore(db ormtable.Schema) (AttestationStore, error) {
	attestationTable, err := NewAttestationTable(db)
	if err != nil {
		return nil, err
	}

	signatureTable, err := NewSignatureTable(db)
	if err != nil {
		return nil, err
	}

	return attestationStore{
		attestationTable,
		signatureTable,
	}, nil
}
