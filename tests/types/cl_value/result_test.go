package cl_value

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/make-software/casper-go-sdk/types/clvalue"
	"github.com/make-software/casper-go-sdk/types/clvalue/cltype"
)

func Test_NewResultFromBuffer_SuccessU32ToString(t *testing.T) {
	source := "010a000000"
	inBytes, err := hex.DecodeString(source)
	require.NoError(t, err)
	res, err := clvalue.NewResultFromBytes(inBytes, cltype.NewResultType(cltype.UInt32))
	require.NoError(t, err)
	assert.Equal(t, "Ok(10)", res.String())
}

func Test_NewResultFromBuffer_ErrorU32ToString(t *testing.T) {
	source := "000a000000"
	inBytes, err := hex.DecodeString(source)
	require.NoError(t, err)
	res, err := clvalue.NewResultFromBytes(inBytes, cltype.NewResultType(cltype.UInt32))
	require.NoError(t, err)
	assert.Equal(t, "Err(10)", res.String())
}

func Test_NewResultFromBuffer_U32ToVal(t *testing.T) {
	source := "010a000000"
	inBytes, err := hex.DecodeString(source)
	require.NoError(t, err)
	res, err := clvalue.NewResultFromBytes(inBytes, cltype.NewResultType(cltype.UInt32))
	require.NoError(t, err)
	assert.Equal(t, uint32(10), res.Value().UI32.Value())
}

func Test_FromBytesByType_ResultU32ToVal(t *testing.T) {
	source := "010a000000"
	inBytes, err := hex.DecodeString(source)
	require.NoError(t, err)
	res, err := clvalue.FromBytesByType(inBytes, cltype.NewResultType(cltype.UInt32))
	require.NoError(t, err)
	assert.Equal(t, uint32(10), res.Result.Value().UI32.Value())
}
