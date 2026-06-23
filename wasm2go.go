package wasm2go

import (
	base "github.com/goccy/pythonwasm2go/base"
	"fmt"
	"unsafe"
	_ "github.com/goccy/pythonwasm2go/p2"
	_ "embed"
)

func NewWithWASIReserve(wasi_snapshot_preview1 base.Wasi_snapshot_preview1Imports, env base.EnvImports, reserveBytes int) *base.Module {
	m := &base.Module{Wasi_snapshot_preview1: wasi_snapshot_preview1, Env: env}
	__memcap := reserveBytes
	if __memcap < 11665408 {
		__memcap = 11665408
	}
	m.Memory = make([]byte, 11665408, __memcap)
	m.M = unsafe.Pointer(unsafe.SliceData(m.Memory))
	m.MaxMem = 4294967296
	m.T0 = make([]any, 5971)
	m.G0 = int32(8388608)
	InitElemSeg_0_0(m)
	InitElemSeg_1_0(m)
	InitElemSeg_1_1(m)
	InitElemSeg_1_2(m)
	InitElemSeg_1_3(m)
	InitElemSeg_2_0(m)
	InitElemSeg_2_1(m)
	InitElemSeg_2_2(m)
	InitElemSeg_2_3(m)
	InitElemSeg_2_4(m)
	InitElemSeg_2_5(m)
	InitElemSeg_2_6(m)
	InitElemSeg_2_7(m)
	InitElemSeg_2_8(m)
	InitElemSeg_2_9(m)
	InitElemSeg_2_10(m)
	InitElemSeg_2_11(m)
	InitElemSeg_2_12(m)
	InitElemSeg_2_13(m)
	InitElemSeg_2_14(m)
	InitElemSeg_2_15(m)
	InitElemSeg_2_16(m)
	InitElemSeg_2_17(m)
	InitElemSeg_2_18(m)
	InitElemSeg_2_19(m)
	initData_0(m)
	return m
}

// NewWithWASI constructs a *Module with a custom
// wasi_snapshot_preview1 implementation and a default initial
// linear-memory reservation. Use NewWithWASIReserve to pre-size
// the reservation (e.g. to cover an interpreter's whole boot and
// avoid reallocating/copying linear memory on the first grow).
func NewWithWASI(wasi_snapshot_preview1 base.Wasi_snapshot_preview1Imports, env base.EnvImports) *base.Module {
	return NewWithWASIReserve(wasi_snapshot_preview1, env, 14581760)
}

// New constructs a *Module using DefaultWASI() for the
// wasi_snapshot_preview1 import. Use NewWithWASI to plug in a
// custom implementation (sandboxed FS, captured stdout, ...).
func New(env base.EnvImports) *base.Module {
	return NewWithWASI(base.DefaultWASI(), env)
}
func initData_0(m *base.Module) {
	copy(m.Memory[8388608:], wasm2goData_data_bin[0:540097])
	copy(m.Memory[8932324:], wasm2goData_data_bin[540097:540430])
	copy(m.Memory[8936035:], wasm2goData_data_bin[540430:582460])
	copy(m.Memory[8981696:], wasm2goData_data_bin[582460:582773])
	copy(m.Memory[8983680:], wasm2goData_data_bin[582773:582782])
	copy(m.Memory[8986561:], wasm2goData_data_bin[582782:582790])
	copy(m.Memory[9019348:], wasm2goData_data_bin[582790:583857])
	copy(m.Memory[9092696:], wasm2goData_data_bin[583857:747815])
	copy(m.Memory[9257744:], wasm2goData_data_bin[747815:776877])
	copy(m.Memory[9287848:], wasm2goData_data_bin[776877:833491])
	copy(m.Memory[9345632:], wasm2goData_data_bin[833491:839205])
	copy(m.Memory[9352648:], wasm2goData_data_bin[839205:854899])
	copy(m.Memory[9369432:], wasm2goData_data_bin[854899:987729])
	copy(m.Memory[9503376:], wasm2goData_data_bin[987729:1000975])
	copy(m.Memory[9517792:], wasm2goData_data_bin[1000975:1052469])
	copy(m.Memory[9570320:], wasm2goData_data_bin[1052469:1058347])
	copy(m.Memory[9577232:], wasm2goData_data_bin[1058347:1116673])
	copy(m.Memory[9636592:], wasm2goData_data_bin[1116673:1121799])
	copy(m.Memory[9642752:], wasm2goData_data_bin[1121799:1142357])
	copy(m.Memory[9665008:], wasm2goData_data_bin[1142357:1165719])
	copy(m.Memory[9689416:], wasm2goData_data_bin[1165719:1184269])
	copy(m.Memory[9709008:], wasm2goData_data_bin[1184269:1271889])
	copy(m.Memory[9797672:], wasm2goData_data_bin[1271889:1352453])
	copy(m.Memory[9879528:], wasm2goData_data_bin[1352453:1356877])
	copy(m.Memory[9885240:], wasm2goData_data_bin[1356877:1506458])
	copy(m.Memory[10042320:], wasm2goData_data_bin[1506458:2751947])
	copy(m.Memory[11291428:], wasm2goData_data_bin[2751947:2752280])
	copy(m.Memory[11295139:], wasm2goData_data_bin[2752280:2794310])
	copy(m.Memory[11340800:], wasm2goData_data_bin[2794310:2794623])
	copy(m.Memory[11342784:], wasm2goData_data_bin[2794623:2794632])
	copy(m.Memory[11345665:], wasm2goData_data_bin[2794632:2794640])
	copy(m.Memory[11378452:], wasm2goData_data_bin[2794640:2795707])
	copy(m.Memory[11451800:], wasm2goData_data_bin[2795707:2827165])
	copy(m.Memory[11484624:], wasm2goData_data_bin[2827165:2910328])
}
func Initialize(m *base.Module) {
	Fn50(m)
}
func WasmAlloc(m *base.Module, l0 int32) int32 {
	return Fn51(m, l0)
}
func WasmFree(m *base.Module, l0 int32) {
	Fn52(m, l0)
}
func WasmifyGetTypeName(m *base.Module, l0 int32, l1 int32) int64 {
	return Fn72(m, l0, l1)
}
func WasmInit(m *base.Module) int32 {
	return Fn74(m)
}
func WasmShutdown(m *base.Module) {
	Fn75(m)
}
func Inv_0_0(m *base.Module, l0, l1 int32) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			err = fmt.Errorf("wasm trap: %v", r)
		}
	}()
	packed = Fn53(m, l0, l1)
	return
}
func Inv_0_1(m *base.Module, l0, l1 int32) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			err = fmt.Errorf("wasm trap: %v", r)
		}
	}()
	packed = Fn63(m, l0, l1)
	return
}
func Inv_0_2(m *base.Module, l0, l1 int32) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			err = fmt.Errorf("wasm trap: %v", r)
		}
	}()
	packed = Fn64(m, l0, l1)
	return
}
func Inv_0_3(m *base.Module, l0, l1 int32) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			err = fmt.Errorf("wasm trap: %v", r)
		}
	}()
	packed = Fn69(m, l0, l1)
	return
}
func Inv_0_4(m *base.Module, l0, l1 int32) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			err = fmt.Errorf("wasm trap: %v", r)
		}
	}()
	packed = Fn70(m, l0, l1)
	return
}
func Inv_0_5(m *base.Module, l0, l1 int32) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			err = fmt.Errorf("wasm trap: %v", r)
		}
	}()
	packed = Fn71(m, l0, l1)
	return
}
func Memory(m *base.Module) []byte {
	return m.Memory
}

//go:embed data.bin
var wasm2goData_data_bin []byte
