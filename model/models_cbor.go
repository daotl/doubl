// Code generated by github.com/daotl/cbor-gen. DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"math"
	"sort"

	cbg "github.com/daotl/cbor-gen"
	cid "github.com/ipfs/go-cid"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

func (t *Transaction) InitNilEmbeddedStruct() {
	if t != nil {
	}
}

var lengthBufTransaction = []byte{135}

func (t *Transaction) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	t.InitNilEmbeddedStruct()
	if _, err := w.Write(lengthBufTransaction); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Type (model.TransactionType) (uint8)
	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Type)); err != nil {
		return err
	}

	// t.From (crpt.Address) (slice)
	if len(t.From) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.From was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.From))); err != nil {
		return err
	}

	if _, err := w.Write(t.From[:]); err != nil {
		return err
	}

	// t.Nonce (uint64) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Nonce)); err != nil {
		return err
	}

	// t.To (crpt.Address) (slice)
	if len(t.To) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.To was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.To))); err != nil {
		return err
	}

	if _, err := w.Write(t.To[:]); err != nil {
		return err
	}

	// t.Data ([]uint8) (slice)
	if len(t.Data) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Data was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Data))); err != nil {
		return err
	}

	if _, err := w.Write(t.Data[:]); err != nil {
		return err
	}

	// t.Extra ([]uint8) (slice)
	if len(t.Extra) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Extra was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Extra))); err != nil {
		return err
	}

	if _, err := w.Write(t.Extra[:]); err != nil {
		return err
	}

	// t.Sig (crpt.Signature) (slice)
	if len(t.Sig) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Sig was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Sig))); err != nil {
		return err
	}

	if _, err := w.Write(t.Sig[:]); err != nil {
		return err
	}
	return nil
}

func (t *Transaction) UnmarshalCBOR(r io.Reader) error {
	*t = Transaction{}
	t.InitNilEmbeddedStruct()

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 7 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Type (model.TransactionType) (uint8)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint8 field")
	}
	if extra > math.MaxUint8 {
		return fmt.Errorf("integer in input was too large for uint8 field")
	}
	t.Type = TransactionType(extra)
	// t.From (crpt.Address) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.From: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.From = make([]uint8, extra)
	}

	if _, err := io.ReadFull(br, t.From[:]); err != nil {
		return err
	}
	// t.Nonce (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Nonce = uint64(extra)

	}
	// t.To (crpt.Address) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.To: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.To = make([]uint8, extra)
	}

	if _, err := io.ReadFull(br, t.To[:]); err != nil {
		return err
	}
	// t.Data ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Data: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Data = make([]uint8, extra)
	}

	if _, err := io.ReadFull(br, t.Data[:]); err != nil {
		return err
	}
	// t.Extra ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Extra: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Extra = make([]uint8, extra)
	}

	if _, err := io.ReadFull(br, t.Extra[:]); err != nil {
		return err
	}
	// t.Sig (crpt.Signature) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Sig: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Sig = make([]uint8, extra)
	}

	if _, err := io.ReadFull(br, t.Sig[:]); err != nil {
		return err
	}
	return nil
}

func (t *BlockHeader) InitNilEmbeddedStruct() {
	if t != nil {
	}
}

var lengthBufBlockHeader = []byte{136}

func (t *BlockHeader) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	t.InitNilEmbeddedStruct()
	if _, err := w.Write(lengthBufBlockHeader); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Creator (crpt.Address) (slice)
	if len(t.Creator) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Creator was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Creator))); err != nil {
		return err
	}

	if _, err := w.Write(t.Creator[:]); err != nil {
		return err
	}

	// t.Time (model.Timestamp) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Time)); err != nil {
		return err
	}

	// t.PrevHashes ([][]uint8) (slice)
	if len(t.PrevHashes) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.PrevHashes was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajArray, uint64(len(t.PrevHashes))); err != nil {
		return err
	}
	for _, v := range t.PrevHashes {
		if len(v) > cbg.ByteArrayMaxLen {
			return xerrors.Errorf("Byte array in field v was too long")
		}

		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(v))); err != nil {
			return err
		}

		if _, err := w.Write(v[:]); err != nil {
			return err
		}
	}

	// t.Height (model.BlockHeight) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Height)); err != nil {
		return err
	}

	// t.TxRoot ([]uint8) (slice)
	if len(t.TxRoot) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.TxRoot was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.TxRoot))); err != nil {
		return err
	}

	if _, err := w.Write(t.TxRoot[:]); err != nil {
		return err
	}

	// t.TxCount (uint64) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.TxCount)); err != nil {
		return err
	}

	// t.Extra ([]uint8) (slice)
	if len(t.Extra) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Extra was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Extra))); err != nil {
		return err
	}

	if _, err := w.Write(t.Extra[:]); err != nil {
		return err
	}

	// t.Sig (crpt.Signature) (slice)
	if len(t.Sig) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Sig was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Sig))); err != nil {
		return err
	}

	if _, err := w.Write(t.Sig[:]); err != nil {
		return err
	}
	return nil
}

func (t *BlockHeader) UnmarshalCBOR(r io.Reader) error {
	*t = BlockHeader{}
	t.InitNilEmbeddedStruct()

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 8 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Creator (crpt.Address) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Creator: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Creator = make([]uint8, extra)
	}

	if _, err := io.ReadFull(br, t.Creator[:]); err != nil {
		return err
	}
	// t.Time (model.Timestamp) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Time = Timestamp(extra)

	}
	// t.PrevHashes ([][]uint8) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.PrevHashes: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.PrevHashes = make([][]uint8, extra)
	}

	for i := 0; i < int(extra); i++ {
		{
			var maj byte
			var extra uint64
			var err error

			maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
			if err != nil {
				return err
			}

			if extra > cbg.ByteArrayMaxLen {
				return fmt.Errorf("t.PrevHashes[i]: byte array too large (%d)", extra)
			}
			if maj != cbg.MajByteString {
				return fmt.Errorf("expected byte array")
			}

			if extra > 0 {
				t.PrevHashes[i] = make([]uint8, extra)
			}

			if _, err := io.ReadFull(br, t.PrevHashes[i][:]); err != nil {
				return err
			}
		}
	}

	// t.Height (model.BlockHeight) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Height = BlockHeight(extra)

	}
	// t.TxRoot ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.TxRoot: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.TxRoot = make([]uint8, extra)
	}

	if _, err := io.ReadFull(br, t.TxRoot[:]); err != nil {
		return err
	}
	// t.TxCount (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.TxCount = uint64(extra)

	}
	// t.Extra ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Extra: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Extra = make([]uint8, extra)
	}

	if _, err := io.ReadFull(br, t.Extra[:]); err != nil {
		return err
	}
	// t.Sig (crpt.Signature) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Sig: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Sig = make([]uint8, extra)
	}

	if _, err := io.ReadFull(br, t.Sig[:]); err != nil {
		return err
	}
	return nil
}

func (t *Block) InitNilEmbeddedStruct() {
	if t != nil {
	}
}

var lengthBufBlock = []byte{130}

func (t *Block) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	t.InitNilEmbeddedStruct()
	if _, err := w.Write(lengthBufBlock); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Header (model.BlockHeader) (struct)
	if err := t.Header.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Txs (model.Transactions) (slice)
	if len(t.Txs) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Txs was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajArray, uint64(len(t.Txs))); err != nil {
		return err
	}
	for _, v := range t.Txs {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}
	return nil
}

func (t *Block) UnmarshalCBOR(r io.Reader) error {
	*t = Block{}
	t.InitNilEmbeddedStruct()

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Header (model.BlockHeader) (struct)

	{

		if err := t.Header.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Header: %w", err)
		}

	}
	// t.Txs (model.Transactions) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Txs: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Txs = make([]Transaction, extra)
	}

	for i := 0; i < int(extra); i++ {

		var v Transaction
		if err := v.UnmarshalCBOR(br); err != nil {
			return err
		}

		t.Txs[i] = v
	}

	return nil
}