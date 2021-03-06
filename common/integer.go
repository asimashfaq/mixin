package common

import (
	"fmt"
	"math"
	"math/big"
	"strconv"

	"github.com/vmihailenco/msgpack"
)

const Precision = 8

func init() {
	msgpack.RegisterExt(0, (*Integer)(nil))
}

type Integer struct {
	i big.Int
}

func NewIntegerFromString(x string) (v Integer) {
	var f big.Float
	p, _, _ := big.ParseFloat(x, 10, 64, big.ToZero)
	d := big.NewFloat(math.Pow(10, Precision))
	f.Mul(p, d).Int(&v.i)
	return
}

func NewInteger(x uint64) (v Integer) {
	p := new(big.Int).SetUint64(x)
	d := big.NewInt(int64(math.Pow(10, Precision)))
	v.i.Mul(p, d)
	return
}

func (x Integer) Add(y Integer) (v Integer) {
	var t Integer
	t.i.Add(&x.i, &y.i)
	if t.Cmp(x) < 0 || t.Cmp(y) < 0 {
		panic(fmt.Sprint(x, y))
	}

	v.i.Add(&x.i, &y.i)
	return
}

func (x Integer) Sub(y Integer) (v Integer) {
	if x.Cmp(y) < 0 {
		panic(fmt.Sprint(x, y))
	}

	v.i.Sub(&x.i, &y.i)
	return
}

func (x Integer) Cmp(y Integer) int {
	return x.i.Cmp(&y.i)
}

func (x Integer) String() string {
	return x.i.String()
}

func (x Integer) MarshalMsgpack() ([]byte, error) {
	return x.i.Bytes(), nil
}

func (x *Integer) UnmarshalMsgpack(data []byte) error {
	x.i.SetBytes(data)
	return nil
}

func (x Integer) MarshalJSON() ([]byte, error) {
	s := x.String()
	p := len(s) - Precision
	s = s[:p] + "." + s[p:]
	return []byte(strconv.Quote(s)), nil
}

func (x *Integer) UnmarshalJSON(b []byte) error {
	unquoted, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}
	i := NewIntegerFromString(unquoted)
	x.i.SetBytes(i.i.Bytes())
	return nil
}
