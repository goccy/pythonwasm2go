package wasm2go

import (
	base "github.com/goccy/pythonwasm2go/base"
	"fmt"
	"sync"
	"sync/atomic"
	"unsafe"
	_ "github.com/goccy/pythonwasm2go/p2"
	_ "embed"
)

func NewWithWASIReserve(wasi_snapshot_preview1 base.Wasi_snapshot_preview1Imports, env base.EnvImports, reserveBytes int) *base.Module {
	m := &base.Module{Wasi_snapshot_preview1: wasi_snapshot_preview1, Env: env}
	__memcap := reserveBytes
	if __memcap < 11599872 {
		__memcap = 11599872
	}
	m.Memory = make([]byte, 11599872, __memcap)
	m.MemMu = &sync.Mutex{}
	m.MemSize = &atomic.Uint64{}
	m.Threads = &base.ThreadPool{}
	m.MemSize.Store(11599872)
	m.M = unsafe.Pointer(unsafe.SliceData(m.Memory))
	m.MaxMem = 4294967296
	m.T0 = make([]any, 5527)
	m.G0 = int32(8388608)
	InitElemSeg_0_0(m)
	InitElemSeg_1_0(m)
	InitElemSeg_1_1(m)
	InitElemSeg_1_2(m)
	InitElemSeg_1_3(m)
	InitElemSeg_1_4(m)
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
	m.DataEnd = 11539179
	initData_0(m)
	return m
}

// NewWithWASI constructs a *Module with a custom
// wasi_snapshot_preview1 implementation and a default initial
// linear-memory reservation. Use NewWithWASIReserve to pre-size
// the reservation (e.g. to cover an interpreter's whole boot and
// avoid reallocating/copying linear memory on the first grow).
func NewWithWASI(wasi_snapshot_preview1 base.Wasi_snapshot_preview1Imports, env base.EnvImports) *base.Module {
	return NewWithWASIReserve(wasi_snapshot_preview1, env, 14499840)
}

// New constructs a *Module using DefaultWASI() for the
// wasi_snapshot_preview1 import. Use NewWithWASI to plug in a
// custom implementation (sandboxed FS, captured stdout, ...).
func New(env base.EnvImports) *base.Module {
	return NewWithWASI(base.DefaultWASI(), env)
}
func NewWithMemory(wasi_snapshot_preview1 base.Wasi_snapshot_preview1Imports, env base.EnvImports, memory []byte, memSize uint64) *base.Module {
	m := &base.Module{Wasi_snapshot_preview1: wasi_snapshot_preview1, Env: env}
	m.Memory = memory
	m.MemMu = &sync.Mutex{}
	m.MemSize = &atomic.Uint64{}
	m.Threads = &base.ThreadPool{}
	m.MemSize.Store(memSize)
	m.M = unsafe.Pointer(unsafe.SliceData(m.Memory))
	m.MaxMem = uint64(len(memory))
	m.T0 = make([]any, 5527)
	m.G0 = int32(8388608)
	InitElemSeg_0_0(m)
	InitElemSeg_1_0(m)
	InitElemSeg_1_1(m)
	InitElemSeg_1_2(m)
	InitElemSeg_1_3(m)
	InitElemSeg_1_4(m)
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
	m.DataEnd = 11539179
	return m
}
func NewFromSnapshot(wasi_snapshot_preview1 base.Wasi_snapshot_preview1Imports, env base.EnvImports, memory []byte, memSize uint64, globals []uint64) *base.Module {
	m := &base.Module{Wasi_snapshot_preview1: wasi_snapshot_preview1, Env: env}
	m.Memory = memory
	m.MemMu = &sync.Mutex{}
	m.MemSize = &atomic.Uint64{}
	m.Threads = &base.ThreadPool{}
	m.MemSize.Store(memSize)
	m.M = unsafe.Pointer(unsafe.SliceData(m.Memory))
	m.MaxMem = uint64(len(memory))
	m.T0 = make([]any, 5527)
	m.G0 = int32(8388608)
	InitElemSeg_0_0(m)
	InitElemSeg_1_0(m)
	InitElemSeg_1_1(m)
	InitElemSeg_1_2(m)
	InitElemSeg_1_3(m)
	InitElemSeg_1_4(m)
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
	m.DataEnd = 11539179
	base.RestoreGlobals(m, globals)
	return m
}
func initData_0(m *base.Module) {
	copy(m.Memory[8388608:], wasm2goData_data_bin[0:523217])
	copy(m.Memory[8915444:], wasm2goData_data_bin[523217:523550])
	copy(m.Memory[8919155:], wasm2goData_data_bin[523550:565580])
	copy(m.Memory[8964816:], wasm2goData_data_bin[565580:565893])
	copy(m.Memory[8966800:], wasm2goData_data_bin[565893:565902])
	copy(m.Memory[8969680:], wasm2goData_data_bin[565902:565911])
	copy(m.Memory[9002468:], wasm2goData_data_bin[565911:566978])
	copy(m.Memory[9075816:], wasm2goData_data_bin[566978:730936])
	copy(m.Memory[9240864:], wasm2goData_data_bin[730936:759998])
	copy(m.Memory[9270968:], wasm2goData_data_bin[759998:816612])
	copy(m.Memory[9328752:], wasm2goData_data_bin[816612:822326])
	copy(m.Memory[9335768:], wasm2goData_data_bin[822326:838020])
	copy(m.Memory[9352552:], wasm2goData_data_bin[838020:970850])
	copy(m.Memory[9486496:], wasm2goData_data_bin[970850:984096])
	copy(m.Memory[9500912:], wasm2goData_data_bin[984096:1035590])
	copy(m.Memory[9553440:], wasm2goData_data_bin[1035590:1041468])
	copy(m.Memory[9560352:], wasm2goData_data_bin[1041468:1099794])
	copy(m.Memory[9619712:], wasm2goData_data_bin[1099794:1104920])
	copy(m.Memory[9625872:], wasm2goData_data_bin[1104920:1125478])
	copy(m.Memory[9648128:], wasm2goData_data_bin[1125478:1148840])
	copy(m.Memory[9672536:], wasm2goData_data_bin[1148840:1167390])
	copy(m.Memory[9692128:], wasm2goData_data_bin[1167390:1255010])
	copy(m.Memory[9780792:], wasm2goData_data_bin[1255010:1335574])
	copy(m.Memory[9862648:], wasm2goData_data_bin[1335574:1339998])
	copy(m.Memory[9868360:], wasm2goData_data_bin[1339998:1489579])
	copy(m.Memory[10025440:], wasm2goData_data_bin[1489579:2732588])
	copy(m.Memory[11272068:], wasm2goData_data_bin[2732588:2732921])
	copy(m.Memory[11275779:], wasm2goData_data_bin[2732921:2774951])
	copy(m.Memory[11321440:], wasm2goData_data_bin[2774951:2775264])
	copy(m.Memory[11323424:], wasm2goData_data_bin[2775264:2775273])
	copy(m.Memory[11326304:], wasm2goData_data_bin[2775273:2775282])
	copy(m.Memory[11359092:], wasm2goData_data_bin[2775282:2776349])
	copy(m.Memory[11432440:], wasm2goData_data_bin[2776349:2807807])
	copy(m.Memory[11465264:], wasm2goData_data_bin[2807807:2881722])
}
func Initialize(m *base.Module) {
	Fn93(m)
}
func WasmAlloc(m *base.Module, l0 int32) int32 {
	return Fn94(m, l0)
}
func WasmFree(m *base.Module, l0 int32) {
	Fn95(m, l0)
}
func WasmifyGetTypeName(m *base.Module, l0 int32, l1 int32) int64 {
	return Fn115(m, l0, l1)
}
func WasmInit(m *base.Module) int32 {
	return Fn117(m)
}
func WasmShutdown(m *base.Module) {
	Fn118(m)
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
	packed = Fn96(m, l0, l1)
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
	packed = Fn106(m, l0, l1)
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
	packed = Fn107(m, l0, l1)
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
	packed = Fn112(m, l0, l1)
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
	packed = Fn113(m, l0, l1)
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
	packed = Fn114(m, l0, l1)
	return
}
func Memory(m *base.Module) []byte {
	return m.Memory
}

//go:embed data.bin
var wasm2goData_data_bin []byte
