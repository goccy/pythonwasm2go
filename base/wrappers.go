//go:build (amd64 || arm64) && !purego

package base

func callImport_0(m *Module, l0 int32, l1 int32, l2 int32) int32 {
	return m.Wasi_snapshot_preview1.Sock_getaddrinfo(m, l0, l1, l2)
}

func callImport_1(m *Module, l0 int32, l1 int32) int32 {
	return m.Wasi_snapshot_preview1.Sock_socket(m, l0, l1)
}

func callImport_2(m *Module, l0 int32, l1 int32, l2 int32) int32 {
	return m.Wasi_snapshot_preview1.Sock_connect(m, l0, l1, l2)
}

func callImport_3(m *Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	return m.Wasi_snapshot_preview1.Proc_spawn(m, l0, l1, l2, l3, l4, l5, l6)
}

func callImport_4(m *Module, l0 int32) int32 {
	return m.Wasi_snapshot_preview1.Pipe(m, l0)
}

func callImport_5(m *Module, l0 int32, l1 int32, l2 int32) int32 {
	return m.Wasi_snapshot_preview1.Proc_wait(m, l0, l1, l2)
}

func callImport_6(m *Module) int32 {
	return m.Env.Getpid(m)
}

func callImport_7(m *Module, l0 int32, l1 int32) int32 {
	return m.Env.Dlopen(m, l0, l1)
}

func callImport_8(m *Module) int32 {
	return m.Env.Dlerror(m)
}

func callImport_9(m *Module, l0 int32, l1 int32) int32 {
	return m.Env.Dlsym(m, l0, l1)
}

func callImport_10(m *Module, l0 int32, l1 int32) int32 {
	return m.Wasi_snapshot_preview1.Environ_get(m, l0, l1)
}

func callImport_11(m *Module, l0 int32, l1 int32) int32 {
	return m.Wasi_snapshot_preview1.Environ_sizes_get(m, l0, l1)
}

func callImport_12(m *Module, l0 int32, l1 int32) int32 {
	return m.Wasi_snapshot_preview1.Clock_res_get(m, l0, l1)
}

func callImport_13(m *Module, l0 int32, l1 int64, l2 int32) int32 {
	return m.Wasi_snapshot_preview1.Clock_time_get(m, l0, l1, l2)
}

func callImport_14(m *Module, l0 int32, l1 int64, l2 int64, l3 int32) int32 {
	return m.Wasi_snapshot_preview1.Fd_advise(m, l0, l1, l2, l3)
}

func callImport_15(m *Module, l0 int32) int32 {
	return m.Wasi_snapshot_preview1.Fd_close(m, l0)
}

func callImport_16(m *Module, l0 int32) int32 {
	return m.Wasi_snapshot_preview1.Fd_datasync(m, l0)
}

func callImport_17(m *Module, l0 int32, l1 int32) int32 {
	return m.Wasi_snapshot_preview1.Fd_fdstat_get(m, l0, l1)
}

func callImport_18(m *Module, l0 int32, l1 int32) int32 {
	return m.Wasi_snapshot_preview1.Fd_fdstat_set_flags(m, l0, l1)
}

func callImport_19(m *Module, l0 int32, l1 int32) int32 {
	return m.Wasi_snapshot_preview1.Fd_filestat_get(m, l0, l1)
}

func callImport_20(m *Module, l0 int32, l1 int64) int32 {
	return m.Wasi_snapshot_preview1.Fd_filestat_set_size(m, l0, l1)
}

func callImport_21(m *Module, l0 int32, l1 int64, l2 int64, l3 int32) int32 {
	return m.Wasi_snapshot_preview1.Fd_filestat_set_times(m, l0, l1, l2, l3)
}

func callImport_22(m *Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32 {
	return m.Wasi_snapshot_preview1.Fd_pread(m, l0, l1, l2, l3, l4)
}

func callImport_23(m *Module, l0 int32, l1 int32) int32 {
	return m.Wasi_snapshot_preview1.Fd_prestat_get(m, l0, l1)
}

func callImport_24(m *Module, l0 int32, l1 int32, l2 int32) int32 {
	return m.Wasi_snapshot_preview1.Fd_prestat_dir_name(m, l0, l1, l2)
}

func callImport_25(m *Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32 {
	return m.Wasi_snapshot_preview1.Fd_pwrite(m, l0, l1, l2, l3, l4)
}

func callImport_26(m *Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	return m.Wasi_snapshot_preview1.Fd_read(m, l0, l1, l2, l3)
}

func callImport_27(m *Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32 {
	return m.Wasi_snapshot_preview1.Fd_readdir(m, l0, l1, l2, l3, l4)
}

func callImport_28(m *Module, l0 int32, l1 int64, l2 int32, l3 int32) int32 {
	return m.Wasi_snapshot_preview1.Fd_seek(m, l0, l1, l2, l3)
}

func callImport_29(m *Module, l0 int32) int32 {
	return m.Wasi_snapshot_preview1.Fd_sync(m, l0)
}

func callImport_30(m *Module, l0 int32, l1 int32) int32 {
	return m.Wasi_snapshot_preview1.Fd_tell(m, l0, l1)
}

func callImport_31(m *Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	return m.Wasi_snapshot_preview1.Fd_write(m, l0, l1, l2, l3)
}

func callImport_32(m *Module, l0 int32, l1 int32, l2 int32) int32 {
	return m.Wasi_snapshot_preview1.Path_create_directory(m, l0, l1, l2)
}

func callImport_33(m *Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	return m.Wasi_snapshot_preview1.Path_filestat_get(m, l0, l1, l2, l3, l4)
}

func callImport_34(m *Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int64, l6 int32) int32 {
	return m.Wasi_snapshot_preview1.Path_filestat_set_times(m, l0, l1, l2, l3, l4, l5, l6)
}

func callImport_35(m *Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	return m.Wasi_snapshot_preview1.Path_link(m, l0, l1, l2, l3, l4, l5, l6)
}

func callImport_36(m *Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int64, l6 int64, l7 int32, l8 int32) int32 {
	return m.Wasi_snapshot_preview1.Path_open(m, l0, l1, l2, l3, l4, l5, l6, l7, l8)
}

func callImport_37(m *Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	return m.Wasi_snapshot_preview1.Path_readlink(m, l0, l1, l2, l3, l4, l5)
}

func callImport_38(m *Module, l0 int32, l1 int32, l2 int32) int32 {
	return m.Wasi_snapshot_preview1.Path_remove_directory(m, l0, l1, l2)
}

func callImport_39(m *Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	return m.Wasi_snapshot_preview1.Path_rename(m, l0, l1, l2, l3, l4, l5)
}

func callImport_40(m *Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	return m.Wasi_snapshot_preview1.Path_symlink(m, l0, l1, l2, l3, l4)
}

func callImport_41(m *Module, l0 int32, l1 int32, l2 int32) int32 {
	return m.Wasi_snapshot_preview1.Path_unlink_file(m, l0, l1, l2)
}

func callImport_42(m *Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	return m.Wasi_snapshot_preview1.Poll_oneoff(m, l0, l1, l2, l3)
}

func callImport_43(m *Module, l0 int32) {
	m.Wasi_snapshot_preview1.Proc_exit(m, l0)
}

func callImport_44(m *Module) int32 {
	return m.Wasi_snapshot_preview1.Sched_yield(m)
}

func callImport_45(m *Module, l0 int32, l1 int32) int32 {
	return m.Wasi_snapshot_preview1.Random_get(m, l0, l1)
}

func callImport_46(m *Module, l0 int32, l1 int32, l2 int32) int32 {
	return m.Wasi_snapshot_preview1.Sock_accept(m, l0, l1, l2)
}

func callImport_47(m *Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	return m.Wasi_snapshot_preview1.Sock_recv(m, l0, l1, l2, l3, l4, l5)
}

func callImport_48(m *Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	return m.Wasi_snapshot_preview1.Sock_send(m, l0, l1, l2, l3, l4)
}

func callImport_49(m *Module, l0 int32, l1 int32) int32 {
	return m.Wasi_snapshot_preview1.Sock_shutdown(m, l0, l1)
}

func loadGlobal_0(m *Module) int32     { return m.G0 }
func storeGlobal_0(m *Module, v int32) { m.G0 = v }

func callIndirect_type0(m *Module, idx int32, p0 int32, p1 int32) int32 {
	return m.T0[idx].(func(*Module, int32, int32) int32)(m, p0, p1)
}

func callIndirect_type1(m *Module, idx int32, p0 int32, p1 int32, p2 int32) int32 {
	return m.T0[idx].(func(*Module, int32, int32, int32) int32)(m, p0, p1, p2)
}

func callIndirect_type2(m *Module, idx int32, p0 int32) int32 {
	return m.T0[idx].(func(*Module, int32) int32)(m, p0)
}

func callIndirect_type3(m *Module, idx int32, p0 int32) {
	m.T0[idx].(func(*Module, int32))(m, p0)
}

func callIndirect_type4(m *Module, idx int32, p0 int32, p1 int32, p2 int32, p3 int32) int32 {
	return m.T0[idx].(func(*Module, int32, int32, int32, int32) int32)(m, p0, p1, p2, p3)
}

func callIndirect_type5(m *Module, idx int32, p0 int32, p1 int32, p2 int32, p3 int32, p4 int32) int32 {
	return m.T0[idx].(func(*Module, int32, int32, int32, int32, int32) int32)(m, p0, p1, p2, p3, p4)
}

func callIndirect_type6(m *Module, idx int32, p0 int32, p1 int32) {
	m.T0[idx].(func(*Module, int32, int32))(m, p0, p1)
}

func callIndirect_type7(m *Module, idx int32, p0 int32, p1 int32, p2 int32) {
	m.T0[idx].(func(*Module, int32, int32, int32))(m, p0, p1, p2)
}

func callIndirect_type8(m *Module, idx int32) int32 {
	return m.T0[idx].(func(*Module) int32)(m)
}

func callIndirect_type9(m *Module, idx int32, p0 int32, p1 int32, p2 int32, p3 int32, p4 int32, p5 int32) int32 {
	return m.T0[idx].(func(*Module, int32, int32, int32, int32, int32, int32) int32)(m, p0, p1, p2, p3, p4, p5)
}

func callIndirect_type10(m *Module, idx int32, p0 int32, p1 int32, p2 int32, p3 int32) {
	m.T0[idx].(func(*Module, int32, int32, int32, int32))(m, p0, p1, p2, p3)
}

func callIndirect_type11(m *Module, idx int32, p0 int32, p1 int32, p2 int32, p3 int32, p4 int32) {
	m.T0[idx].(func(*Module, int32, int32, int32, int32, int32))(m, p0, p1, p2, p3, p4)
}

func callIndirect_type12(m *Module, idx int32, p0 int32, p1 int32, p2 int32, p3 int32, p4 int32, p5 int32, p6 int32) int32 {
	return m.T0[idx].(func(*Module, int32, int32, int32, int32, int32, int32, int32) int32)(m, p0, p1, p2, p3, p4, p5, p6)
}

func callIndirect_type13(m *Module, idx int32) {
	m.T0[idx].(func(*Module))(m)
}

func callIndirect_type14(m *Module, idx int32, p0 int32, p1 int32, p2 int32, p3 int32, p4 int32, p5 int32, p6 int32, p7 int32, p8 int32) int32 {
	return m.T0[idx].(func(*Module, int32, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, p0, p1, p2, p3, p4, p5, p6, p7, p8)
}

func callIndirect_type15(m *Module, idx int32, p0 float64) float64 {
	return m.T0[idx].(func(*Module, float64) float64)(m, p0)
}

func callIndirect_type17(m *Module, idx int32, p0 int32, p1 int32, p2 int32, p3 int32, p4 int32, p5 int32) {
	m.T0[idx].(func(*Module, int32, int32, int32, int32, int32, int32))(m, p0, p1, p2, p3, p4, p5)
}

func callIndirect_type19(m *Module, idx int32, p0 int32, p1 int64) int32 {
	return m.T0[idx].(func(*Module, int32, int64) int32)(m, p0, p1)
}

func callIndirect_type20(m *Module, idx int32, p0 int32, p1 int32) int64 {
	return m.T0[idx].(func(*Module, int32, int32) int64)(m, p0, p1)
}

func callIndirect_type21(m *Module, idx int32, p0 int32, p1 int32, p2 int32, p3 int32, p4 int32, p5 int32, p6 int32, p7 int32, p8 int32, p9 int32) int32 {
	return m.T0[idx].(func(*Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
}

func callIndirect_type22(m *Module, idx int32, p0 int32, p1 int64, p2 int32) int64 {
	return m.T0[idx].(func(*Module, int32, int64, int32) int64)(m, p0, p1, p2)
}

func callIndirect_type24(m *Module, idx int32, p0 float64, p1 float64) float64 {
	return m.T0[idx].(func(*Module, float64, float64) float64)(m, p0, p1)
}

func callIndirect_type29(m *Module, idx int32, p0 int32, p1 int32, p2 int32, p3 int32, p4 int32, p5 int32, p6 int32, p7 int32, p8 int32) {
	m.T0[idx].(func(*Module, int32, int32, int32, int32, int32, int32, int32, int32, int32))(m, p0, p1, p2, p3, p4, p5, p6, p7, p8)
}

func callIndirect_type31(m *Module, idx int32, p0 int64, p1 int32) int32 {
	return m.T0[idx].(func(*Module, int64, int32) int32)(m, p0, p1)
}

func callIndirect_type41(m *Module, idx int32, p0 int32, p1 int64, p2 int32) {
	m.T0[idx].(func(*Module, int32, int64, int32))(m, p0, p1, p2)
}

func callIndirect_type45(m *Module, idx int32, p0 int32, p1 float32) int32 {
	return m.T0[idx].(func(*Module, int32, float32) int32)(m, p0, p1)
}
