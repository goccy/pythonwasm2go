//go:build !arm64 && (!amd64 || !amd64.v2)

package p2

import (
	base "github.com/goccy/pythonwasm2go/base"
	_ "unsafe"
)

//go:linkname Fn121 github.com/goccy/pythonwasm2go/p1.Fn121
func Fn121(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn122 github.com/goccy/pythonwasm2go/p1.Fn122
func Fn122(m *base.Module, l0 int32)

//go:linkname Fn145 github.com/goccy/pythonwasm2go/p1.Fn145
func Fn145(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn184 github.com/goccy/pythonwasm2go/p0.Fn184
func Fn184(m *base.Module)

//go:linkname Fn185 github.com/goccy/pythonwasm2go/p0.Fn185
func Fn185(m *base.Module, l0 int32) int32

//go:linkname Fn186 github.com/goccy/pythonwasm2go/p0.Fn186
func Fn186(m *base.Module, l0 int32) int32

//go:linkname Fn187 github.com/goccy/pythonwasm2go/p0.Fn187
func Fn187(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn189 github.com/goccy/pythonwasm2go/p0.Fn189
func Fn189(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn190 github.com/goccy/pythonwasm2go/p0.Fn190
func Fn190(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn191 github.com/goccy/pythonwasm2go/p0.Fn191
func Fn191(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn192 github.com/goccy/pythonwasm2go/p0.Fn192
func Fn192(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn193 github.com/goccy/pythonwasm2go/p0.Fn193
func Fn193(m *base.Module, l0 int32) int32

//go:linkname Fn194 github.com/goccy/pythonwasm2go/p0.Fn194
func Fn194(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn195 github.com/goccy/pythonwasm2go/p0.Fn195
func Fn195(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn196 github.com/goccy/pythonwasm2go/p0.Fn196
func Fn196(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn201 github.com/goccy/pythonwasm2go/p0.Fn201
func Fn201(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn208 github.com/goccy/pythonwasm2go/p0.Fn208
func Fn208(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn214 github.com/goccy/pythonwasm2go/p1.Fn214
func Fn214(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn230 github.com/goccy/pythonwasm2go/p1.Fn230
func Fn230(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn254 github.com/goccy/pythonwasm2go/p1.Fn254
func Fn254(m *base.Module, l0 int32) int32

//go:linkname Fn255 github.com/goccy/pythonwasm2go/p1.Fn255
func Fn255(m *base.Module, l0 int32) int32

//go:linkname Fn258 github.com/goccy/pythonwasm2go/p0.Fn258
func Fn258(m *base.Module, l0 int32) int32

//go:linkname Fn262 github.com/goccy/pythonwasm2go/p0.Fn262
func Fn262(m *base.Module, l0 int32) int32

//go:linkname Fn263 github.com/goccy/pythonwasm2go/p0.Fn263
func Fn263(m *base.Module, l0 int32) int32

//go:linkname Fn264 github.com/goccy/pythonwasm2go/p0.Fn264
func Fn264(m *base.Module, l0 int32) int32

//go:linkname Fn265 github.com/goccy/pythonwasm2go/p0.Fn265
func Fn265(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn266 github.com/goccy/pythonwasm2go/p0.Fn266
func Fn266(m *base.Module, l0 int32) int32

//go:linkname Fn267 github.com/goccy/pythonwasm2go/p0.Fn267
func Fn267(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn268 github.com/goccy/pythonwasm2go/p0.Fn268
func Fn268(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn270 github.com/goccy/pythonwasm2go/p0.Fn270
func Fn270(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn276 github.com/goccy/pythonwasm2go/p0.Fn276
func Fn276(m *base.Module, l0 int32) int32

//go:linkname Fn277 github.com/goccy/pythonwasm2go/p0.Fn277
func Fn277(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn281 github.com/goccy/pythonwasm2go/p0.Fn281
func Fn281(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn282 github.com/goccy/pythonwasm2go/p0.Fn282
func Fn282(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn284 github.com/goccy/pythonwasm2go/p0.Fn284
func Fn284(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn285 github.com/goccy/pythonwasm2go/p0.Fn285
func Fn285(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn286 github.com/goccy/pythonwasm2go/p0.Fn286
func Fn286(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn287 github.com/goccy/pythonwasm2go/p0.Fn287
func Fn287(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn288 github.com/goccy/pythonwasm2go/p0.Fn288
func Fn288(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn315 github.com/goccy/pythonwasm2go/p1.Fn315
func Fn315(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn320 github.com/goccy/pythonwasm2go/p1.Fn320
func Fn320(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn337 github.com/goccy/pythonwasm2go/p0.Fn337
func Fn337(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn340 github.com/goccy/pythonwasm2go/p0.Fn340
func Fn340(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn341 github.com/goccy/pythonwasm2go/p0.Fn341
func Fn341(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn342 github.com/goccy/pythonwasm2go/p0.Fn342
func Fn342(m *base.Module, l0 int32) int32

//go:linkname Fn345 github.com/goccy/pythonwasm2go/p1.Fn345
func Fn345(m *base.Module, l0 int32) int32

//go:linkname Fn363 github.com/goccy/pythonwasm2go/p1.Fn363
func Fn363(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn366 github.com/goccy/pythonwasm2go/p1.Fn366
func Fn366(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn406 github.com/goccy/pythonwasm2go/p1.Fn406
func Fn406(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn435 github.com/goccy/pythonwasm2go/p0.Fn435
func Fn435(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn436 github.com/goccy/pythonwasm2go/p0.Fn436
func Fn436(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn438 github.com/goccy/pythonwasm2go/p0.Fn438
func Fn438(m *base.Module, l0 int32) int32

//go:linkname Fn439 github.com/goccy/pythonwasm2go/p0.Fn439
func Fn439(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn440 github.com/goccy/pythonwasm2go/p0.Fn440
func Fn440(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn444 github.com/goccy/pythonwasm2go/p0.Fn444
func Fn444(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn446 github.com/goccy/pythonwasm2go/p0.Fn446
func Fn446(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn447 github.com/goccy/pythonwasm2go/p0.Fn447
func Fn447(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn453 github.com/goccy/pythonwasm2go/p0.Fn453
func Fn453(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn454 github.com/goccy/pythonwasm2go/p1.Fn454
func Fn454(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn461 github.com/goccy/pythonwasm2go/p1.Fn461
func Fn461(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn462 github.com/goccy/pythonwasm2go/p1.Fn462
func Fn462(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn464 github.com/goccy/pythonwasm2go/p1.Fn464
func Fn464(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn465 github.com/goccy/pythonwasm2go/p1.Fn465
func Fn465(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn466 github.com/goccy/pythonwasm2go/p1.Fn466
func Fn466(m *base.Module, l0 int32) int32

//go:linkname Fn536 github.com/goccy/pythonwasm2go/p0.Fn536
func Fn536(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn538 github.com/goccy/pythonwasm2go/p0.Fn538
func Fn538(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn548 github.com/goccy/pythonwasm2go/p0.Fn548
func Fn548(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn549 github.com/goccy/pythonwasm2go/p0.Fn549
func Fn549(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn550 github.com/goccy/pythonwasm2go/p0.Fn550
func Fn550(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn551 github.com/goccy/pythonwasm2go/p0.Fn551
func Fn551(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn552 github.com/goccy/pythonwasm2go/p0.Fn552
func Fn552(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn555 github.com/goccy/pythonwasm2go/p0.Fn555
func Fn555(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn557 github.com/goccy/pythonwasm2go/p0.Fn557
func Fn557(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn558 github.com/goccy/pythonwasm2go/p0.Fn558
func Fn558(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn559 github.com/goccy/pythonwasm2go/p0.Fn559
func Fn559(m *base.Module, l0 int32)

//go:linkname Fn561 github.com/goccy/pythonwasm2go/p0.Fn561
func Fn561(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn562 github.com/goccy/pythonwasm2go/p0.Fn562
func Fn562(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn563 github.com/goccy/pythonwasm2go/p0.Fn563
func Fn563(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn564 github.com/goccy/pythonwasm2go/p0.Fn564
func Fn564(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn566 github.com/goccy/pythonwasm2go/p0.Fn566
func Fn566(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn567 github.com/goccy/pythonwasm2go/p0.Fn567
func Fn567(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn568 github.com/goccy/pythonwasm2go/p0.Fn568
func Fn568(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn611 github.com/goccy/pythonwasm2go/p0.Fn611
func Fn611(m *base.Module, l0 int32) int32

//go:linkname Fn614 github.com/goccy/pythonwasm2go/p0.Fn614
func Fn614(m *base.Module, l0 int32) int32

//go:linkname Fn616 github.com/goccy/pythonwasm2go/p0.Fn616
func Fn616(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn710 github.com/goccy/pythonwasm2go/p0.Fn710
func Fn710(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn724 github.com/goccy/pythonwasm2go/p0.Fn724
func Fn724(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn725 github.com/goccy/pythonwasm2go/p0.Fn725
func Fn725(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn732 github.com/goccy/pythonwasm2go/p1.Fn732
func Fn732(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn734 github.com/goccy/pythonwasm2go/p1.Fn734
func Fn734(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn736 github.com/goccy/pythonwasm2go/p1.Fn736
func Fn736(m *base.Module, l0 int32) int32

//go:linkname Fn743 github.com/goccy/pythonwasm2go/p0.Fn743
func Fn743(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn744 github.com/goccy/pythonwasm2go/p0.Fn744
func Fn744(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn745 github.com/goccy/pythonwasm2go/p0.Fn745
func Fn745(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn759 github.com/goccy/pythonwasm2go/p0.Fn759
func Fn759(m *base.Module, l0 int32) int32

//go:linkname Fn761 github.com/goccy/pythonwasm2go/p0.Fn761
func Fn761(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn787 github.com/goccy/pythonwasm2go/p0.Fn787
func Fn787(m *base.Module, l0 int32) int32

//go:linkname Fn832 github.com/goccy/pythonwasm2go/p1.Fn832
func Fn832(m *base.Module, l0 int32) int32

//go:linkname Fn835 github.com/goccy/pythonwasm2go/p0.Fn835
func Fn835(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn852 github.com/goccy/pythonwasm2go/p0.Fn852
func Fn852(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn864 github.com/goccy/pythonwasm2go/p1.Fn864
func Fn864(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn867 github.com/goccy/pythonwasm2go/p1.Fn867
func Fn867(m *base.Module, l0 int32) int32

//go:linkname Fn875 github.com/goccy/pythonwasm2go/p1.Fn875
func Fn875(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn878 github.com/goccy/pythonwasm2go/p1.Fn878
func Fn878(m *base.Module, l0 int32) int32

//go:linkname Fn895 github.com/goccy/pythonwasm2go/p1.Fn895
func Fn895(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn900 github.com/goccy/pythonwasm2go/p1.Fn900
func Fn900(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn934 github.com/goccy/pythonwasm2go/p0.Fn934
func Fn934(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn935 github.com/goccy/pythonwasm2go/p0.Fn935
func Fn935(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn937 github.com/goccy/pythonwasm2go/p0.Fn937
func Fn937(m *base.Module, l0 int32) int32

//go:linkname Fn941 github.com/goccy/pythonwasm2go/p0.Fn941
func Fn941(m *base.Module, l0 int32) int32

//go:linkname Fn948 github.com/goccy/pythonwasm2go/p1.Fn948
func Fn948(m *base.Module) int32

//go:linkname Fn949 github.com/goccy/pythonwasm2go/p0.Fn949
func Fn949(m *base.Module, l0 float64) int32

//go:linkname Fn953 github.com/goccy/pythonwasm2go/p1.Fn953
func Fn953(m *base.Module, l0 int32) int32

//go:linkname Fn955 github.com/goccy/pythonwasm2go/p0.Fn955
func Fn955(m *base.Module, l0 int32) float64

//go:linkname Fn964 github.com/goccy/pythonwasm2go/p1.Fn964
func Fn964(m *base.Module, l0 float64, l1 int32, l2 int32) int32

//go:linkname Fn965 github.com/goccy/pythonwasm2go/p1.Fn965
func Fn965(m *base.Module, l0 float64, l1 int32, l2 int32) int32

//go:linkname Fn966 github.com/goccy/pythonwasm2go/p1.Fn966
func Fn966(m *base.Module, l0 float64, l1 int32, l2 int32) int32

//go:linkname Fn968 github.com/goccy/pythonwasm2go/p1.Fn968
func Fn968(m *base.Module, l0 int32, l1 int32) float64

//go:linkname Fn969 github.com/goccy/pythonwasm2go/p1.Fn969
func Fn969(m *base.Module, l0 int32, l1 int32) float64

//go:linkname Fn1013 github.com/goccy/pythonwasm2go/p1.Fn1013
func Fn1013(m *base.Module, l0 int32) int32

//go:linkname Fn1041 github.com/goccy/pythonwasm2go/p1.Fn1041
func Fn1041(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn1080 github.com/goccy/pythonwasm2go/p0.Fn1080
func Fn1080(m *base.Module, l0 int32) int32

//go:linkname Fn1081 github.com/goccy/pythonwasm2go/p1.Fn1081
func Fn1081(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1086 github.com/goccy/pythonwasm2go/p0.Fn1086
func Fn1086(m *base.Module, l0 int32) int32

//go:linkname Fn1087 github.com/goccy/pythonwasm2go/p1.Fn1087
func Fn1087(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1097 github.com/goccy/pythonwasm2go/p1.Fn1097
func Fn1097(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1098 github.com/goccy/pythonwasm2go/p1.Fn1098
func Fn1098(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1126 github.com/goccy/pythonwasm2go/p1.Fn1126
func Fn1126(m *base.Module, l0 int32) int32

//go:linkname Fn1139 github.com/goccy/pythonwasm2go/p1.Fn1139
func Fn1139(m *base.Module, l0 int32) int32

//go:linkname Fn1144 github.com/goccy/pythonwasm2go/p1.Fn1144
func Fn1144(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1195 github.com/goccy/pythonwasm2go/p0.Fn1195
func Fn1195(m *base.Module, l0 int32) int32

//go:linkname Fn1203 github.com/goccy/pythonwasm2go/p0.Fn1203
func Fn1203(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1204 github.com/goccy/pythonwasm2go/p0.Fn1204
func Fn1204(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1205 github.com/goccy/pythonwasm2go/p0.Fn1205
func Fn1205(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1206 github.com/goccy/pythonwasm2go/p0.Fn1206
func Fn1206(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1207 github.com/goccy/pythonwasm2go/p0.Fn1207
func Fn1207(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1208 github.com/goccy/pythonwasm2go/p0.Fn1208
func Fn1208(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1209 github.com/goccy/pythonwasm2go/p0.Fn1209
func Fn1209(m *base.Module, l0 int32) int32

//go:linkname Fn1210 github.com/goccy/pythonwasm2go/p0.Fn1210
func Fn1210(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1211 github.com/goccy/pythonwasm2go/p0.Fn1211
func Fn1211(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1212 github.com/goccy/pythonwasm2go/p0.Fn1212
func Fn1212(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1214 github.com/goccy/pythonwasm2go/p0.Fn1214
func Fn1214(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1219 github.com/goccy/pythonwasm2go/p0.Fn1219
func Fn1219(m *base.Module, l0 int32) int32

//go:linkname Fn1222 github.com/goccy/pythonwasm2go/p0.Fn1222
func Fn1222(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1234 github.com/goccy/pythonwasm2go/p0.Fn1234
func Fn1234(m *base.Module, l0 int32) int32

//go:linkname Fn1284 github.com/goccy/pythonwasm2go/p0.Fn1284
func Fn1284(m *base.Module, l0 int32) int32

//go:linkname Fn1287 github.com/goccy/pythonwasm2go/p0.Fn1287
func Fn1287(m *base.Module, l0 int32) int32

//go:linkname Fn1288 github.com/goccy/pythonwasm2go/p0.Fn1288
func Fn1288(m *base.Module, l0 int32) int32

//go:linkname Fn1289 github.com/goccy/pythonwasm2go/p0.Fn1289
func Fn1289(m *base.Module, l0 int32) int32

//go:linkname Fn1290 github.com/goccy/pythonwasm2go/p0.Fn1290
func Fn1290(m *base.Module, l0 int64) int32

//go:linkname Fn1292 github.com/goccy/pythonwasm2go/p0.Fn1292
func Fn1292(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1293 github.com/goccy/pythonwasm2go/p0.Fn1293
func Fn1293(m *base.Module, l0 int32) int32

//go:linkname Fn1294 github.com/goccy/pythonwasm2go/p0.Fn1294
func Fn1294(m *base.Module, l0 int32) int32

//go:linkname Fn1295 github.com/goccy/pythonwasm2go/p0.Fn1295
func Fn1295(m *base.Module, l0 int32) int32

//go:linkname Fn1299 github.com/goccy/pythonwasm2go/p0.Fn1299
func Fn1299(m *base.Module, l0 int32) int32

//go:linkname Fn1303 github.com/goccy/pythonwasm2go/p1.Fn1303
func Fn1303(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1307 github.com/goccy/pythonwasm2go/p0.Fn1307
func Fn1307(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn1308 github.com/goccy/pythonwasm2go/p1.Fn1308
func Fn1308(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1311 github.com/goccy/pythonwasm2go/p0.Fn1311
func Fn1311(m *base.Module, l0 int64) int32

//go:linkname Fn1312 github.com/goccy/pythonwasm2go/p0.Fn1312
func Fn1312(m *base.Module, l0 int32) int64

//go:linkname Fn1315 github.com/goccy/pythonwasm2go/p0.Fn1315
func Fn1315(m *base.Module, l0 int32) int64

//go:linkname Fn1322 github.com/goccy/pythonwasm2go/p0.Fn1322
func Fn1322(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1323 github.com/goccy/pythonwasm2go/p1.Fn1323
func Fn1323(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn1325 github.com/goccy/pythonwasm2go/p0.Fn1325
func Fn1325(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1330 github.com/goccy/pythonwasm2go/p1.Fn1330
func Fn1330(m *base.Module, l0 int32, l1 int32) float64

//go:linkname Fn1333 github.com/goccy/pythonwasm2go/p0.Fn1333
func Fn1333(m *base.Module, l0 int32) float64

//go:linkname Fn1338 github.com/goccy/pythonwasm2go/p1.Fn1338
func Fn1338(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1341 github.com/goccy/pythonwasm2go/p1.Fn1341
func Fn1341(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1352 github.com/goccy/pythonwasm2go/p1.Fn1352
func Fn1352(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1353 github.com/goccy/pythonwasm2go/p1.Fn1353
func Fn1353(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1361 github.com/goccy/pythonwasm2go/p1.Fn1361
func Fn1361(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1377 github.com/goccy/pythonwasm2go/p1.Fn1377
func Fn1377(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1386 github.com/goccy/pythonwasm2go/p1.Fn1386
func Fn1386(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1408 github.com/goccy/pythonwasm2go/p0.Fn1408
func Fn1408(m *base.Module) int32

//go:linkname Fn1409 github.com/goccy/pythonwasm2go/p0.Fn1409
func Fn1409(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1417 github.com/goccy/pythonwasm2go/p0.Fn1417
func Fn1417(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1421 github.com/goccy/pythonwasm2go/p0.Fn1421
func Fn1421(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1424 github.com/goccy/pythonwasm2go/p1.Fn1424
func Fn1424(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1425 github.com/goccy/pythonwasm2go/p0.Fn1425
func Fn1425(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1426 github.com/goccy/pythonwasm2go/p0.Fn1426
func Fn1426(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1427 github.com/goccy/pythonwasm2go/p0.Fn1427
func Fn1427(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1428 github.com/goccy/pythonwasm2go/p0.Fn1428
func Fn1428(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1430 github.com/goccy/pythonwasm2go/p0.Fn1430
func Fn1430(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1431 github.com/goccy/pythonwasm2go/p0.Fn1431
func Fn1431(m *base.Module, l0 int32) int32

//go:linkname Fn1433 github.com/goccy/pythonwasm2go/p0.Fn1433
func Fn1433(m *base.Module, l0 int32)

//go:linkname Fn1435 github.com/goccy/pythonwasm2go/p0.Fn1435
func Fn1435(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1436 github.com/goccy/pythonwasm2go/p0.Fn1436
func Fn1436(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1437 github.com/goccy/pythonwasm2go/p0.Fn1437
func Fn1437(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1440 github.com/goccy/pythonwasm2go/p0.Fn1440
func Fn1440(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1441 github.com/goccy/pythonwasm2go/p0.Fn1441
func Fn1441(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1442 github.com/goccy/pythonwasm2go/p0.Fn1442
func Fn1442(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1444 github.com/goccy/pythonwasm2go/p0.Fn1444
func Fn1444(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1450 github.com/goccy/pythonwasm2go/p0.Fn1450
func Fn1450(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1451 github.com/goccy/pythonwasm2go/p0.Fn1451
func Fn1451(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1453 github.com/goccy/pythonwasm2go/p0.Fn1453
func Fn1453(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn1456 github.com/goccy/pythonwasm2go/p0.Fn1456
func Fn1456(m *base.Module, l0 int32)

//go:linkname Fn1457 github.com/goccy/pythonwasm2go/p0.Fn1457
func Fn1457(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1458 github.com/goccy/pythonwasm2go/p0.Fn1458
func Fn1458(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1460 github.com/goccy/pythonwasm2go/p1.Fn1460
func Fn1460(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1462 github.com/goccy/pythonwasm2go/p0.Fn1462
func Fn1462(m *base.Module, l0 int32) int32

//go:linkname Fn1464 github.com/goccy/pythonwasm2go/p1.Fn1464
func Fn1464(m *base.Module, l0 int32) int32

//go:linkname Fn1465 github.com/goccy/pythonwasm2go/p1.Fn1465
func Fn1465(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1466 github.com/goccy/pythonwasm2go/p0.Fn1466
func Fn1466(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1467 github.com/goccy/pythonwasm2go/p0.Fn1467
func Fn1467(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1468 github.com/goccy/pythonwasm2go/p0.Fn1468
func Fn1468(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1471 github.com/goccy/pythonwasm2go/p0.Fn1471
func Fn1471(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1472 github.com/goccy/pythonwasm2go/p0.Fn1472
func Fn1472(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1474 github.com/goccy/pythonwasm2go/p0.Fn1474
func Fn1474(m *base.Module, l0 int32) int32

//go:linkname Fn1475 github.com/goccy/pythonwasm2go/p0.Fn1475
func Fn1475(m *base.Module, l0 int32) int32

//go:linkname Fn1477 github.com/goccy/pythonwasm2go/p0.Fn1477
func Fn1477(m *base.Module, l0 int32) int32

//go:linkname Fn1478 github.com/goccy/pythonwasm2go/p0.Fn1478
func Fn1478(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1494 github.com/goccy/pythonwasm2go/p0.Fn1494
func Fn1494(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1495 github.com/goccy/pythonwasm2go/p0.Fn1495
func Fn1495(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1518 github.com/goccy/pythonwasm2go/p0.Fn1518
func Fn1518(m *base.Module, l0 int32) int32

//go:linkname Fn1520 github.com/goccy/pythonwasm2go/p0.Fn1520
func Fn1520(m *base.Module, l0 int32) int32

//go:linkname Fn1524 github.com/goccy/pythonwasm2go/p1.Fn1524
func Fn1524(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1536 github.com/goccy/pythonwasm2go/p0.Fn1536
func Fn1536(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1564 github.com/goccy/pythonwasm2go/p0.Fn1564
func Fn1564(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1577 github.com/goccy/pythonwasm2go/p0.Fn1577
func Fn1577(m *base.Module) int32

//go:linkname Fn1578 github.com/goccy/pythonwasm2go/p0.Fn1578
func Fn1578(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1579 github.com/goccy/pythonwasm2go/p0.Fn1579
func Fn1579(m *base.Module, l0 int32) int32

//go:linkname Fn1584 github.com/goccy/pythonwasm2go/p0.Fn1584
func Fn1584(m *base.Module, l0 int32) int32

//go:linkname Fn1595 github.com/goccy/pythonwasm2go/p1.Fn1595
func Fn1595(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1606 github.com/goccy/pythonwasm2go/p1.Fn1606
func Fn1606(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1655 github.com/goccy/pythonwasm2go/p0.Fn1655
func Fn1655(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1683 github.com/goccy/pythonwasm2go/p1.Fn1683
func Fn1683(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1687 github.com/goccy/pythonwasm2go/p0.Fn1687
func Fn1687(m *base.Module, l0 int32) int32

//go:linkname Fn1700 github.com/goccy/pythonwasm2go/p0.Fn1700
func Fn1700(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1719 github.com/goccy/pythonwasm2go/p1.Fn1719
func Fn1719(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1725 github.com/goccy/pythonwasm2go/p1.Fn1725
func Fn1725(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1743 github.com/goccy/pythonwasm2go/p1.Fn1743
func Fn1743(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1755 github.com/goccy/pythonwasm2go/p1.Fn1755
func Fn1755(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1757 github.com/goccy/pythonwasm2go/p1.Fn1757
func Fn1757(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1795 github.com/goccy/pythonwasm2go/p0.Fn1795
func Fn1795(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn1797 github.com/goccy/pythonwasm2go/p0.Fn1797
func Fn1797(m *base.Module, l0 int32)

//go:linkname Fn1799 github.com/goccy/pythonwasm2go/p1.Fn1799
func Fn1799(m *base.Module, l0 int32)

//go:linkname Fn1804 github.com/goccy/pythonwasm2go/p0.Fn1804
func Fn1804(m *base.Module, l0 int32) int32

//go:linkname Fn1809 github.com/goccy/pythonwasm2go/p0.Fn1809
func Fn1809(m *base.Module, l0 int32) int32

//go:linkname Fn1810 github.com/goccy/pythonwasm2go/p0.Fn1810
func Fn1810(m *base.Module, l0 int32) int32

//go:linkname Fn1811 github.com/goccy/pythonwasm2go/p0.Fn1811
func Fn1811(m *base.Module, l0 int32) int32

//go:linkname Fn1817 github.com/goccy/pythonwasm2go/p0.Fn1817
func Fn1817(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1818 github.com/goccy/pythonwasm2go/p0.Fn1818
func Fn1818(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1820 github.com/goccy/pythonwasm2go/p0.Fn1820
func Fn1820(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1821 github.com/goccy/pythonwasm2go/p0.Fn1821
func Fn1821(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1824 github.com/goccy/pythonwasm2go/p0.Fn1824
func Fn1824(m *base.Module, l0 int32) int32

//go:linkname Fn1825 github.com/goccy/pythonwasm2go/p0.Fn1825
func Fn1825(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1826 github.com/goccy/pythonwasm2go/p0.Fn1826
func Fn1826(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1831 github.com/goccy/pythonwasm2go/p0.Fn1831
func Fn1831(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1835 github.com/goccy/pythonwasm2go/p0.Fn1835
func Fn1835(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1839 github.com/goccy/pythonwasm2go/p0.Fn1839
func Fn1839(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1840 github.com/goccy/pythonwasm2go/p1.Fn1840
func Fn1840(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1845 github.com/goccy/pythonwasm2go/p1.Fn1845
func Fn1845(m *base.Module, l0 int32) int32

//go:linkname Fn1868 github.com/goccy/pythonwasm2go/p1.Fn1868
func Fn1868(m *base.Module) int32

//go:linkname Fn1871 github.com/goccy/pythonwasm2go/p1.Fn1871
func Fn1871(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1872 github.com/goccy/pythonwasm2go/p1.Fn1872
func Fn1872(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1880 github.com/goccy/pythonwasm2go/p1.Fn1880
func Fn1880(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1891 github.com/goccy/pythonwasm2go/p1.Fn1891
func Fn1891(m *base.Module, l0 int32) int32

//go:linkname Fn1894 github.com/goccy/pythonwasm2go/p1.Fn1894
func Fn1894(m *base.Module, l0 int32) int32

//go:linkname Fn1907 github.com/goccy/pythonwasm2go/p1.Fn1907
func Fn1907(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1915 github.com/goccy/pythonwasm2go/p1.Fn1915
func Fn1915(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1925 github.com/goccy/pythonwasm2go/p1.Fn1925
func Fn1925(m *base.Module)

//go:linkname Fn1930 github.com/goccy/pythonwasm2go/p1.Fn1930
func Fn1930(m *base.Module, l0 int32)

//go:linkname Fn1964 github.com/goccy/pythonwasm2go/p1.Fn1964
func Fn1964(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1965 github.com/goccy/pythonwasm2go/p1.Fn1965
func Fn1965(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1973 github.com/goccy/pythonwasm2go/p1.Fn1973
func Fn1973(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1986 github.com/goccy/pythonwasm2go/p1.Fn1986
func Fn1986(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn1994 github.com/goccy/pythonwasm2go/p1.Fn1994
func Fn1994(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn2008 github.com/goccy/pythonwasm2go/p0.Fn2008
func Fn2008(m *base.Module, l0 int32)

//go:linkname Fn2010 github.com/goccy/pythonwasm2go/p0.Fn2010
func Fn2010(m *base.Module, l0 int32)

//go:linkname Fn2040 github.com/goccy/pythonwasm2go/p1.Fn2040
func Fn2040(m *base.Module, l0 int32)

//go:linkname Fn2041 github.com/goccy/pythonwasm2go/p1.Fn2041
func Fn2041(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2074 github.com/goccy/pythonwasm2go/p1.Fn2074
func Fn2074(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2088 github.com/goccy/pythonwasm2go/p1.Fn2088
func Fn2088(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2100 github.com/goccy/pythonwasm2go/p0.Fn2100
func Fn2100(m *base.Module, l0 int32)

//go:linkname Fn2101 github.com/goccy/pythonwasm2go/p0.Fn2101
func Fn2101(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2103 github.com/goccy/pythonwasm2go/p0.Fn2103
func Fn2103(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2107 github.com/goccy/pythonwasm2go/p0.Fn2107
func Fn2107(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2109 github.com/goccy/pythonwasm2go/p0.Fn2109
func Fn2109(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2123 github.com/goccy/pythonwasm2go/p0.Fn2123
func Fn2123(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2129 github.com/goccy/pythonwasm2go/p0.Fn2129
func Fn2129(m *base.Module, l0 int32) int32

//go:linkname Fn2135 github.com/goccy/pythonwasm2go/p0.Fn2135
func Fn2135(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2139 github.com/goccy/pythonwasm2go/p0.Fn2139
func Fn2139(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2140 github.com/goccy/pythonwasm2go/p0.Fn2140
func Fn2140(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2141 github.com/goccy/pythonwasm2go/p0.Fn2141
func Fn2141(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2147 github.com/goccy/pythonwasm2go/p0.Fn2147
func Fn2147(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2148 github.com/goccy/pythonwasm2go/p0.Fn2148
func Fn2148(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2149 github.com/goccy/pythonwasm2go/p0.Fn2149
func Fn2149(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2153 github.com/goccy/pythonwasm2go/p0.Fn2153
func Fn2153(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2155 github.com/goccy/pythonwasm2go/p1.Fn2155
func Fn2155(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2157 github.com/goccy/pythonwasm2go/p1.Fn2157
func Fn2157(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2163 github.com/goccy/pythonwasm2go/p1.Fn2163
func Fn2163(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2192 github.com/goccy/pythonwasm2go/p1.Fn2192
func Fn2192(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2197 github.com/goccy/pythonwasm2go/p1.Fn2197
func Fn2197(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2207 github.com/goccy/pythonwasm2go/p0.Fn2207
func Fn2207(m *base.Module, l0 int32) int32

//go:linkname Fn2208 github.com/goccy/pythonwasm2go/p0.Fn2208
func Fn2208(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2214 github.com/goccy/pythonwasm2go/p1.Fn2214
func Fn2214(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2231 github.com/goccy/pythonwasm2go/p0.Fn2231
func Fn2231(m *base.Module, l0 int32) int32

//go:linkname Fn2232 github.com/goccy/pythonwasm2go/p0.Fn2232
func Fn2232(m *base.Module, l0 int32) int32

//go:linkname Fn2233 github.com/goccy/pythonwasm2go/p0.Fn2233
func Fn2233(m *base.Module, l0 int32) int32

//go:linkname Fn2236 github.com/goccy/pythonwasm2go/p0.Fn2236
func Fn2236(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2237 github.com/goccy/pythonwasm2go/p0.Fn2237
func Fn2237(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2239 github.com/goccy/pythonwasm2go/p0.Fn2239
func Fn2239(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2251 github.com/goccy/pythonwasm2go/p1.Fn2251
func Fn2251(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2282 github.com/goccy/pythonwasm2go/p0.Fn2282
func Fn2282(m *base.Module, l0 int32)

//go:linkname Fn2284 github.com/goccy/pythonwasm2go/p0.Fn2284
func Fn2284(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2285 github.com/goccy/pythonwasm2go/p0.Fn2285
func Fn2285(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2288 github.com/goccy/pythonwasm2go/p0.Fn2288
func Fn2288(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2289 github.com/goccy/pythonwasm2go/p0.Fn2289
func Fn2289(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2294 github.com/goccy/pythonwasm2go/p0.Fn2294
func Fn2294(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2296 github.com/goccy/pythonwasm2go/p0.Fn2296
func Fn2296(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2297 github.com/goccy/pythonwasm2go/p0.Fn2297
func Fn2297(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2302 github.com/goccy/pythonwasm2go/p0.Fn2302
func Fn2302(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2309 github.com/goccy/pythonwasm2go/p0.Fn2309
func Fn2309(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2315 github.com/goccy/pythonwasm2go/p0.Fn2315
func Fn2315(m *base.Module, l0 int32) int32

//go:linkname Fn2319 github.com/goccy/pythonwasm2go/p0.Fn2319
func Fn2319(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2335 github.com/goccy/pythonwasm2go/p0.Fn2335
func Fn2335(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2341 github.com/goccy/pythonwasm2go/p0.Fn2341
func Fn2341(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2357 github.com/goccy/pythonwasm2go/p1.Fn2357
func Fn2357(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2365 github.com/goccy/pythonwasm2go/p1.Fn2365
func Fn2365(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2366 github.com/goccy/pythonwasm2go/p0.Fn2366
func Fn2366(m *base.Module, l0 int32) int32

//go:linkname Fn2382 github.com/goccy/pythonwasm2go/p0.Fn2382
func Fn2382(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2388 github.com/goccy/pythonwasm2go/p1.Fn2388
func Fn2388(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2395 github.com/goccy/pythonwasm2go/p1.Fn2395
func Fn2395(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2397 github.com/goccy/pythonwasm2go/p0.Fn2397
func Fn2397(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2398 github.com/goccy/pythonwasm2go/p0.Fn2398
func Fn2398(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2399 github.com/goccy/pythonwasm2go/p0.Fn2399
func Fn2399(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2400 github.com/goccy/pythonwasm2go/p1.Fn2400
func Fn2400(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2406 github.com/goccy/pythonwasm2go/p0.Fn2406
func Fn2406(m *base.Module, l0 int32) int32

//go:linkname Fn2412 github.com/goccy/pythonwasm2go/p1.Fn2412
func Fn2412(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2446 github.com/goccy/pythonwasm2go/p0.Fn2446
func Fn2446(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2564 github.com/goccy/pythonwasm2go/p1.Fn2564
func Fn2564(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32

//go:linkname Fn2639 github.com/goccy/pythonwasm2go/p0.Fn2639
func Fn2639(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2641 github.com/goccy/pythonwasm2go/p1.Fn2641
func Fn2641(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2642 github.com/goccy/pythonwasm2go/p1.Fn2642
func Fn2642(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2646 github.com/goccy/pythonwasm2go/p0.Fn2646
func Fn2646(m *base.Module, l0 int32) int32

//go:linkname Fn2651 github.com/goccy/pythonwasm2go/p0.Fn2651
func Fn2651(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2652 github.com/goccy/pythonwasm2go/p0.Fn2652
func Fn2652(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2654 github.com/goccy/pythonwasm2go/p0.Fn2654
func Fn2654(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2655 github.com/goccy/pythonwasm2go/p0.Fn2655
func Fn2655(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2656 github.com/goccy/pythonwasm2go/p1.Fn2656
func Fn2656(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2659 github.com/goccy/pythonwasm2go/p0.Fn2659
func Fn2659(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2661 github.com/goccy/pythonwasm2go/p0.Fn2661
func Fn2661(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2662 github.com/goccy/pythonwasm2go/p0.Fn2662
func Fn2662(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2663 github.com/goccy/pythonwasm2go/p0.Fn2663
func Fn2663(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2664 github.com/goccy/pythonwasm2go/p0.Fn2664
func Fn2664(m *base.Module, l0 int32) int32

//go:linkname Fn2665 github.com/goccy/pythonwasm2go/p0.Fn2665
func Fn2665(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2666 github.com/goccy/pythonwasm2go/p0.Fn2666
func Fn2666(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2667 github.com/goccy/pythonwasm2go/p0.Fn2667
func Fn2667(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2669 github.com/goccy/pythonwasm2go/p0.Fn2669
func Fn2669(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2670 github.com/goccy/pythonwasm2go/p0.Fn2670
func Fn2670(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2671 github.com/goccy/pythonwasm2go/p0.Fn2671
func Fn2671(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2672 github.com/goccy/pythonwasm2go/p0.Fn2672
func Fn2672(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2674 github.com/goccy/pythonwasm2go/p0.Fn2674
func Fn2674(m *base.Module, l0 int32) int32

//go:linkname Fn2676 github.com/goccy/pythonwasm2go/p0.Fn2676
func Fn2676(m *base.Module, l0 int32) int32

//go:linkname Fn2678 github.com/goccy/pythonwasm2go/p1.Fn2678
func Fn2678(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2680 github.com/goccy/pythonwasm2go/p0.Fn2680
func Fn2680(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2681 github.com/goccy/pythonwasm2go/p0.Fn2681
func Fn2681(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2682 github.com/goccy/pythonwasm2go/p0.Fn2682
func Fn2682(m *base.Module, l0 int32) int32

//go:linkname Fn2684 github.com/goccy/pythonwasm2go/p0.Fn2684
func Fn2684(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2686 github.com/goccy/pythonwasm2go/p0.Fn2686
func Fn2686(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2690 github.com/goccy/pythonwasm2go/p0.Fn2690
func Fn2690(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2694 github.com/goccy/pythonwasm2go/p0.Fn2694
func Fn2694(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2695 github.com/goccy/pythonwasm2go/p0.Fn2695
func Fn2695(m *base.Module, l0 int32) int32

//go:linkname Fn2698 github.com/goccy/pythonwasm2go/p0.Fn2698
func Fn2698(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2699 github.com/goccy/pythonwasm2go/p0.Fn2699
func Fn2699(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2702 github.com/goccy/pythonwasm2go/p0.Fn2702
func Fn2702(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2703 github.com/goccy/pythonwasm2go/p0.Fn2703
func Fn2703(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2704 github.com/goccy/pythonwasm2go/p0.Fn2704
func Fn2704(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2705 github.com/goccy/pythonwasm2go/p0.Fn2705
func Fn2705(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2711 github.com/goccy/pythonwasm2go/p0.Fn2711
func Fn2711(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2712 github.com/goccy/pythonwasm2go/p0.Fn2712
func Fn2712(m *base.Module, l0 int32) int32

//go:linkname Fn2713 github.com/goccy/pythonwasm2go/p0.Fn2713
func Fn2713(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2714 github.com/goccy/pythonwasm2go/p0.Fn2714
func Fn2714(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2716 github.com/goccy/pythonwasm2go/p0.Fn2716
func Fn2716(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2717 github.com/goccy/pythonwasm2go/p0.Fn2717
func Fn2717(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2718 github.com/goccy/pythonwasm2go/p0.Fn2718
func Fn2718(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2719 github.com/goccy/pythonwasm2go/p0.Fn2719
func Fn2719(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2720 github.com/goccy/pythonwasm2go/p0.Fn2720
func Fn2720(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2721 github.com/goccy/pythonwasm2go/p0.Fn2721
func Fn2721(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2722 github.com/goccy/pythonwasm2go/p0.Fn2722
func Fn2722(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2723 github.com/goccy/pythonwasm2go/p0.Fn2723
func Fn2723(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2728 github.com/goccy/pythonwasm2go/p0.Fn2728
func Fn2728(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2729 github.com/goccy/pythonwasm2go/p0.Fn2729
func Fn2729(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2730 github.com/goccy/pythonwasm2go/p0.Fn2730
func Fn2730(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2731 github.com/goccy/pythonwasm2go/p0.Fn2731
func Fn2731(m *base.Module, l0 int32) int32

//go:linkname Fn2732 github.com/goccy/pythonwasm2go/p0.Fn2732
func Fn2732(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2734 github.com/goccy/pythonwasm2go/p0.Fn2734
func Fn2734(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn2744 github.com/goccy/pythonwasm2go/p0.Fn2744
func Fn2744(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2745 github.com/goccy/pythonwasm2go/p0.Fn2745
func Fn2745(m *base.Module, l0 int32) int32

//go:linkname Fn2750 github.com/goccy/pythonwasm2go/p1.Fn2750
func Fn2750(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2751 github.com/goccy/pythonwasm2go/p0.Fn2751
func Fn2751(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2752 github.com/goccy/pythonwasm2go/p0.Fn2752
func Fn2752(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2753 github.com/goccy/pythonwasm2go/p1.Fn2753
func Fn2753(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2755 github.com/goccy/pythonwasm2go/p1.Fn2755
func Fn2755(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2756 github.com/goccy/pythonwasm2go/p0.Fn2756
func Fn2756(m *base.Module, l0 int32) int32

//go:linkname Fn2759 github.com/goccy/pythonwasm2go/p1.Fn2759
func Fn2759(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2761 github.com/goccy/pythonwasm2go/p1.Fn2761
func Fn2761(m *base.Module, l0 int32) int32

//go:linkname Fn2762 github.com/goccy/pythonwasm2go/p1.Fn2762
func Fn2762(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2764 github.com/goccy/pythonwasm2go/p1.Fn2764
func Fn2764(m *base.Module, l0 int32) int32

//go:linkname Fn2769 github.com/goccy/pythonwasm2go/p1.Fn2769
func Fn2769(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2770 github.com/goccy/pythonwasm2go/p1.Fn2770
func Fn2770(m *base.Module, l0 int32) int32

//go:linkname Fn2771 github.com/goccy/pythonwasm2go/p1.Fn2771
func Fn2771(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2776 github.com/goccy/pythonwasm2go/p0.Fn2776
func Fn2776(m *base.Module, l0 int32) int32

//go:linkname Fn2777 github.com/goccy/pythonwasm2go/p1.Fn2777
func Fn2777(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2785 github.com/goccy/pythonwasm2go/p0.Fn2785
func Fn2785(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2790 github.com/goccy/pythonwasm2go/p1.Fn2790
func Fn2790(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2794 github.com/goccy/pythonwasm2go/p1.Fn2794
func Fn2794(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2795 github.com/goccy/pythonwasm2go/p1.Fn2795
func Fn2795(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2796 github.com/goccy/pythonwasm2go/p1.Fn2796
func Fn2796(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2797 github.com/goccy/pythonwasm2go/p1.Fn2797
func Fn2797(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2803 github.com/goccy/pythonwasm2go/p1.Fn2803
func Fn2803(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2805 github.com/goccy/pythonwasm2go/p1.Fn2805
func Fn2805(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2807 github.com/goccy/pythonwasm2go/p0.Fn2807
func Fn2807(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2808 github.com/goccy/pythonwasm2go/p1.Fn2808
func Fn2808(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2811 github.com/goccy/pythonwasm2go/p1.Fn2811
func Fn2811(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2816 github.com/goccy/pythonwasm2go/p1.Fn2816
func Fn2816(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2821 github.com/goccy/pythonwasm2go/p0.Fn2821
func Fn2821(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2822 github.com/goccy/pythonwasm2go/p0.Fn2822
func Fn2822(m *base.Module, l0 int32) int32

//go:linkname Fn2824 github.com/goccy/pythonwasm2go/p1.Fn2824
func Fn2824(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2827 github.com/goccy/pythonwasm2go/p0.Fn2827
func Fn2827(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2828 github.com/goccy/pythonwasm2go/p1.Fn2828
func Fn2828(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2829 github.com/goccy/pythonwasm2go/p1.Fn2829
func Fn2829(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2830 github.com/goccy/pythonwasm2go/p0.Fn2830
func Fn2830(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2842 github.com/goccy/pythonwasm2go/p0.Fn2842
func Fn2842(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2845 github.com/goccy/pythonwasm2go/p0.Fn2845
func Fn2845(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2862 github.com/goccy/pythonwasm2go/p0.Fn2862
func Fn2862(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2863 github.com/goccy/pythonwasm2go/p0.Fn2863
func Fn2863(m *base.Module, l0 int32) int32

//go:linkname Fn2897 github.com/goccy/pythonwasm2go/p1.Fn2897
func Fn2897(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2946 github.com/goccy/pythonwasm2go/p1.Fn2946
func Fn2946(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2954 github.com/goccy/pythonwasm2go/p1.Fn2954
func Fn2954(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn2989 github.com/goccy/pythonwasm2go/p1.Fn2989
func Fn2989(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2991 github.com/goccy/pythonwasm2go/p1.Fn2991
func Fn2991(m *base.Module, l0 int32) int32

//go:linkname Fn2992 github.com/goccy/pythonwasm2go/p1.Fn2992
func Fn2992(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3022 github.com/goccy/pythonwasm2go/p0.Fn3022
func Fn3022(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3033 github.com/goccy/pythonwasm2go/p0.Fn3033
func Fn3033(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3039 github.com/goccy/pythonwasm2go/p1.Fn3039
func Fn3039(m *base.Module, l0 int32)

//go:linkname Fn3087 github.com/goccy/pythonwasm2go/p0.Fn3087
func Fn3087(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3088 github.com/goccy/pythonwasm2go/p0.Fn3088
func Fn3088(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3089 github.com/goccy/pythonwasm2go/p0.Fn3089
func Fn3089(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3091 github.com/goccy/pythonwasm2go/p0.Fn3091
func Fn3091(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3093 github.com/goccy/pythonwasm2go/p0.Fn3093
func Fn3093(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3094 github.com/goccy/pythonwasm2go/p0.Fn3094
func Fn3094(m *base.Module) int32

//go:linkname Fn3095 github.com/goccy/pythonwasm2go/p0.Fn3095
func Fn3095(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32

//go:linkname Fn3096 github.com/goccy/pythonwasm2go/p0.Fn3096
func Fn3096(m *base.Module, l0 int32) int32

//go:linkname Fn3099 github.com/goccy/pythonwasm2go/p0.Fn3099
func Fn3099(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3194 github.com/goccy/pythonwasm2go/p1.Fn3194
func Fn3194(m *base.Module, l0 int32) int32

//go:linkname Fn3199 github.com/goccy/pythonwasm2go/p0.Fn3199
func Fn3199(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3212 github.com/goccy/pythonwasm2go/p1.Fn3212
func Fn3212(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3246 github.com/goccy/pythonwasm2go/p0.Fn3246
func Fn3246(m *base.Module, l0 int32) int32

//go:linkname Fn3247 github.com/goccy/pythonwasm2go/p0.Fn3247
func Fn3247(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3265 github.com/goccy/pythonwasm2go/p1.Fn3265
func Fn3265(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3316 github.com/goccy/pythonwasm2go/p1.Fn3316
func Fn3316(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3333 github.com/goccy/pythonwasm2go/p1.Fn3333
func Fn3333(m *base.Module, l0 int32)

//go:linkname Fn3335 github.com/goccy/pythonwasm2go/p0.Fn3335
func Fn3335(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3338 github.com/goccy/pythonwasm2go/p1.Fn3338
func Fn3338(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3340 github.com/goccy/pythonwasm2go/p0.Fn3340
func Fn3340(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3374 github.com/goccy/pythonwasm2go/p1.Fn3374
func Fn3374(m *base.Module) int32

//go:linkname Fn3385 github.com/goccy/pythonwasm2go/p0.Fn3385
func Fn3385(m *base.Module, l0 int32) int32

//go:linkname Fn3394 github.com/goccy/pythonwasm2go/p0.Fn3394
func Fn3394(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3395 github.com/goccy/pythonwasm2go/p0.Fn3395
func Fn3395(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3397 github.com/goccy/pythonwasm2go/p0.Fn3397
func Fn3397(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3401 github.com/goccy/pythonwasm2go/p0.Fn3401
func Fn3401(m *base.Module, l0 int32) int32

//go:linkname Fn3402 github.com/goccy/pythonwasm2go/p0.Fn3402
func Fn3402(m *base.Module, l0 int32) int32

//go:linkname Fn3410 github.com/goccy/pythonwasm2go/p1.Fn3410
func Fn3410(m *base.Module, l0 int32) int32

//go:linkname Fn3412 github.com/goccy/pythonwasm2go/p1.Fn3412
func Fn3412(m *base.Module, l0 int32) int32

//go:linkname Fn3422 github.com/goccy/pythonwasm2go/p0.Fn3422
func Fn3422(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3423 github.com/goccy/pythonwasm2go/p0.Fn3423
func Fn3423(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3439 github.com/goccy/pythonwasm2go/p1.Fn3439
func Fn3439(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3440 github.com/goccy/pythonwasm2go/p1.Fn3440
func Fn3440(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3442 github.com/goccy/pythonwasm2go/p0.Fn3442
func Fn3442(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3445 github.com/goccy/pythonwasm2go/p1.Fn3445
func Fn3445(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3471 github.com/goccy/pythonwasm2go/p0.Fn3471
func Fn3471(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3509 github.com/goccy/pythonwasm2go/p0.Fn3509
func Fn3509(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn3520 github.com/goccy/pythonwasm2go/p0.Fn3520
func Fn3520(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3521 github.com/goccy/pythonwasm2go/p0.Fn3521
func Fn3521(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3522 github.com/goccy/pythonwasm2go/p0.Fn3522
func Fn3522(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3538 github.com/goccy/pythonwasm2go/p1.Fn3538
func Fn3538(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3540 github.com/goccy/pythonwasm2go/p1.Fn3540
func Fn3540(m *base.Module, l0 int32)

//go:linkname Fn3557 github.com/goccy/pythonwasm2go/p1.Fn3557
func Fn3557(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3564 github.com/goccy/pythonwasm2go/p1.Fn3564
func Fn3564(m *base.Module, l0 int32) int32

//go:linkname Fn3579 github.com/goccy/pythonwasm2go/p0.Fn3579
func Fn3579(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3580 github.com/goccy/pythonwasm2go/p1.Fn3580
func Fn3580(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3634 github.com/goccy/pythonwasm2go/p1.Fn3634
func Fn3634(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3636 github.com/goccy/pythonwasm2go/p1.Fn3636
func Fn3636(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3661 github.com/goccy/pythonwasm2go/p0.Fn3661
func Fn3661(m *base.Module) int32

//go:linkname Fn3672 github.com/goccy/pythonwasm2go/p0.Fn3672
func Fn3672(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3683 github.com/goccy/pythonwasm2go/p0.Fn3683
func Fn3683(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3690 github.com/goccy/pythonwasm2go/p0.Fn3690
func Fn3690(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3697 github.com/goccy/pythonwasm2go/p0.Fn3697
func Fn3697(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3700 github.com/goccy/pythonwasm2go/p0.Fn3700
func Fn3700(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3701 github.com/goccy/pythonwasm2go/p0.Fn3701
func Fn3701(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3702 github.com/goccy/pythonwasm2go/p0.Fn3702
func Fn3702(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3706 github.com/goccy/pythonwasm2go/p0.Fn3706
func Fn3706(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3707 github.com/goccy/pythonwasm2go/p0.Fn3707
func Fn3707(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3709 github.com/goccy/pythonwasm2go/p0.Fn3709
func Fn3709(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3711 github.com/goccy/pythonwasm2go/p0.Fn3711
func Fn3711(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3712 github.com/goccy/pythonwasm2go/p0.Fn3712
func Fn3712(m *base.Module, l0 int32)

//go:linkname Fn3715 github.com/goccy/pythonwasm2go/p0.Fn3715
func Fn3715(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3716 github.com/goccy/pythonwasm2go/p0.Fn3716
func Fn3716(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3718 github.com/goccy/pythonwasm2go/p0.Fn3718
func Fn3718(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3719 github.com/goccy/pythonwasm2go/p0.Fn3719
func Fn3719(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3720 github.com/goccy/pythonwasm2go/p0.Fn3720
func Fn3720(m *base.Module, l0 int32) int32

//go:linkname Fn3721 github.com/goccy/pythonwasm2go/p0.Fn3721
func Fn3721(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3732 github.com/goccy/pythonwasm2go/p0.Fn3732
func Fn3732(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3733 github.com/goccy/pythonwasm2go/p0.Fn3733
func Fn3733(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3735 github.com/goccy/pythonwasm2go/p0.Fn3735
func Fn3735(m *base.Module)

//go:linkname Fn3736 github.com/goccy/pythonwasm2go/p0.Fn3736
func Fn3736(m *base.Module) int32

//go:linkname Fn3738 github.com/goccy/pythonwasm2go/p0.Fn3738
func Fn3738(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3740 github.com/goccy/pythonwasm2go/p0.Fn3740
func Fn3740(m *base.Module, l0 int32) int32

//go:linkname Fn3744 github.com/goccy/pythonwasm2go/p0.Fn3744
func Fn3744(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3745 github.com/goccy/pythonwasm2go/p0.Fn3745
func Fn3745(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3746 github.com/goccy/pythonwasm2go/p1.Fn3746
func Fn3746(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3748 github.com/goccy/pythonwasm2go/p0.Fn3748
func Fn3748(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3749 github.com/goccy/pythonwasm2go/p0.Fn3749
func Fn3749(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3750 github.com/goccy/pythonwasm2go/p0.Fn3750
func Fn3750(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3751 github.com/goccy/pythonwasm2go/p1.Fn3751
func Fn3751(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn3762 github.com/goccy/pythonwasm2go/p0.Fn3762
func Fn3762(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3774 github.com/goccy/pythonwasm2go/p1.Fn3774
func Fn3774(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3782 github.com/goccy/pythonwasm2go/p1.Fn3782
func Fn3782(m *base.Module, l0 int32) int32

//go:linkname Fn3792 github.com/goccy/pythonwasm2go/p1.Fn3792
func Fn3792(m *base.Module, l0 int32) int32

//go:linkname Fn3794 github.com/goccy/pythonwasm2go/p1.Fn3794
func Fn3794(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3805 github.com/goccy/pythonwasm2go/p0.Fn3805
func Fn3805(m *base.Module, l0 int32) int32

//go:linkname Fn3806 github.com/goccy/pythonwasm2go/p1.Fn3806
func Fn3806(m *base.Module, l0 int32)

//go:linkname Fn3809 github.com/goccy/pythonwasm2go/p1.Fn3809
func Fn3809(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3819 github.com/goccy/pythonwasm2go/p0.Fn3819
func Fn3819(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3828 github.com/goccy/pythonwasm2go/p0.Fn3828
func Fn3828(m *base.Module, l0 int32)

//go:linkname Fn3831 github.com/goccy/pythonwasm2go/p0.Fn3831
func Fn3831(m *base.Module, l0 int32)

//go:linkname Fn3832 github.com/goccy/pythonwasm2go/p0.Fn3832
func Fn3832(m *base.Module, l0 int32) int32

//go:linkname Fn3834 github.com/goccy/pythonwasm2go/p0.Fn3834
func Fn3834(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3835 github.com/goccy/pythonwasm2go/p0.Fn3835
func Fn3835(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3842 github.com/goccy/pythonwasm2go/p0.Fn3842
func Fn3842(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3843 github.com/goccy/pythonwasm2go/p0.Fn3843
func Fn3843(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3848 github.com/goccy/pythonwasm2go/p0.Fn3848
func Fn3848(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3849 github.com/goccy/pythonwasm2go/p0.Fn3849
func Fn3849(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3853 github.com/goccy/pythonwasm2go/p1.Fn3853
func Fn3853(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3859 github.com/goccy/pythonwasm2go/p1.Fn3859
func Fn3859(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32

//go:linkname Fn3861 github.com/goccy/pythonwasm2go/p1.Fn3861
func Fn3861(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3874 github.com/goccy/pythonwasm2go/p0.Fn3874
func Fn3874(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3878 github.com/goccy/pythonwasm2go/p0.Fn3878
func Fn3878(m *base.Module, l0 int32)

//go:linkname Fn3879 github.com/goccy/pythonwasm2go/p0.Fn3879
func Fn3879(m *base.Module) int32

//go:linkname Fn3880 github.com/goccy/pythonwasm2go/p0.Fn3880
func Fn3880(m *base.Module, l0 int32)

//go:linkname Fn3883 github.com/goccy/pythonwasm2go/p0.Fn3883
func Fn3883(m *base.Module, l0 int32)

//go:linkname Fn3884 github.com/goccy/pythonwasm2go/p0.Fn3884
func Fn3884(m *base.Module, l0 int32) int32

//go:linkname Fn3886 github.com/goccy/pythonwasm2go/p0.Fn3886
func Fn3886(m *base.Module, l0 int32) int32

//go:linkname Fn3887 github.com/goccy/pythonwasm2go/p0.Fn3887
func Fn3887(m *base.Module) int32

//go:linkname Fn3890 github.com/goccy/pythonwasm2go/p0.Fn3890
func Fn3890(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3895 github.com/goccy/pythonwasm2go/p0.Fn3895
func Fn3895(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3899 github.com/goccy/pythonwasm2go/p1.Fn3899
func Fn3899(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3902 github.com/goccy/pythonwasm2go/p0.Fn3902
func Fn3902(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3903 github.com/goccy/pythonwasm2go/p0.Fn3903
func Fn3903(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3904 github.com/goccy/pythonwasm2go/p1.Fn3904
func Fn3904(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3949 github.com/goccy/pythonwasm2go/p0.Fn3949
func Fn3949(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3952 github.com/goccy/pythonwasm2go/p0.Fn3952
func Fn3952(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3953 github.com/goccy/pythonwasm2go/p0.Fn3953
func Fn3953(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3954 github.com/goccy/pythonwasm2go/p0.Fn3954
func Fn3954(m *base.Module, l0 int32)

//go:linkname Fn3967 github.com/goccy/pythonwasm2go/p1.Fn3967
func Fn3967(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3976 github.com/goccy/pythonwasm2go/p1.Fn3976
func Fn3976(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3980 github.com/goccy/pythonwasm2go/p1.Fn3980
func Fn3980(m *base.Module, l0 int32) int32

//go:linkname Fn3986 github.com/goccy/pythonwasm2go/p0.Fn3986
func Fn3986(m *base.Module, l0 int32) int32

//go:linkname Fn3987 github.com/goccy/pythonwasm2go/p0.Fn3987
func Fn3987(m *base.Module, l0 int32) int32

//go:linkname Fn3988 github.com/goccy/pythonwasm2go/p0.Fn3988
func Fn3988(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3991 github.com/goccy/pythonwasm2go/p1.Fn3991
func Fn3991(m *base.Module, l0 int32) int32

//go:linkname Fn3993 github.com/goccy/pythonwasm2go/p0.Fn3993
func Fn3993(m *base.Module, l0 int32)

//go:linkname Fn3994 github.com/goccy/pythonwasm2go/p0.Fn3994
func Fn3994(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3995 github.com/goccy/pythonwasm2go/p0.Fn3995
func Fn3995(m *base.Module, l0 int32)

//go:linkname Fn3996 github.com/goccy/pythonwasm2go/p0.Fn3996
func Fn3996(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3999 github.com/goccy/pythonwasm2go/p1.Fn3999
func Fn3999(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4000 github.com/goccy/pythonwasm2go/p1.Fn4000
func Fn4000(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4038 github.com/goccy/pythonwasm2go/p1.Fn4038
func Fn4038(m *base.Module, l0 int32)

//go:linkname Fn4045 github.com/goccy/pythonwasm2go/p1.Fn4045
func Fn4045(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4046 github.com/goccy/pythonwasm2go/p1.Fn4046
func Fn4046(m *base.Module, l0 int32) int32

//go:linkname Fn4074 github.com/goccy/pythonwasm2go/p1.Fn4074
func Fn4074(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn4088 github.com/goccy/pythonwasm2go/p0.Fn4088
func Fn4088(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4101 github.com/goccy/pythonwasm2go/p1.Fn4101
func Fn4101(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4103 github.com/goccy/pythonwasm2go/p1.Fn4103
func Fn4103(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn4179 github.com/goccy/pythonwasm2go/p0.Fn4179
func Fn4179(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn4180 github.com/goccy/pythonwasm2go/p0.Fn4180
func Fn4180(m *base.Module, l0 int32) int32

//go:linkname Fn4188 github.com/goccy/pythonwasm2go/p0.Fn4188
func Fn4188(m *base.Module, l0 int32)

//go:linkname Fn4189 github.com/goccy/pythonwasm2go/p0.Fn4189
func Fn4189(m *base.Module, l0 int32)

//go:linkname Fn4191 github.com/goccy/pythonwasm2go/p0.Fn4191
func Fn4191(m *base.Module, l0 int32) int32

//go:linkname Fn4192 github.com/goccy/pythonwasm2go/p0.Fn4192
func Fn4192(m *base.Module, l0 int32)

//go:linkname Fn4202 github.com/goccy/pythonwasm2go/p0.Fn4202
func Fn4202(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4209 github.com/goccy/pythonwasm2go/p0.Fn4209
func Fn4209(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4211 github.com/goccy/pythonwasm2go/p1.Fn4211
func Fn4211(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4218 github.com/goccy/pythonwasm2go/p0.Fn4218
func Fn4218(m *base.Module, l0 int32) int32

//go:linkname Fn4233 github.com/goccy/pythonwasm2go/p0.Fn4233
func Fn4233(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4234 github.com/goccy/pythonwasm2go/p0.Fn4234
func Fn4234(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4238 github.com/goccy/pythonwasm2go/p0.Fn4238
func Fn4238(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4250 github.com/goccy/pythonwasm2go/p0.Fn4250
func Fn4250(m *base.Module, l0 int32)

//go:linkname Fn4252 github.com/goccy/pythonwasm2go/p0.Fn4252
func Fn4252(m *base.Module, l0 int32, l1 int64) int32

//go:linkname Fn4254 github.com/goccy/pythonwasm2go/p0.Fn4254
func Fn4254(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn4255 github.com/goccy/pythonwasm2go/p0.Fn4255
func Fn4255(m *base.Module, l0 int32)

//go:linkname Fn4256 github.com/goccy/pythonwasm2go/p0.Fn4256
func Fn4256(m *base.Module, l0 int32)

//go:linkname Fn4257 github.com/goccy/pythonwasm2go/p0.Fn4257
func Fn4257(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4258 github.com/goccy/pythonwasm2go/p0.Fn4258
func Fn4258(m *base.Module, l0 int32)

//go:linkname Fn4289 github.com/goccy/pythonwasm2go/p1.Fn4289
func Fn4289(m *base.Module, l0 int32)

//go:linkname Fn4294 github.com/goccy/pythonwasm2go/p0.Fn4294
func Fn4294(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4302 github.com/goccy/pythonwasm2go/p0.Fn4302
func Fn4302(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn4304 github.com/goccy/pythonwasm2go/p0.Fn4304
func Fn4304(m *base.Module, l0 int32)

//go:linkname Fn4305 github.com/goccy/pythonwasm2go/p0.Fn4305
func Fn4305(m *base.Module, l0 int32)

//go:linkname Fn4306 github.com/goccy/pythonwasm2go/p0.Fn4306
func Fn4306(m *base.Module, l0 int32)

//go:linkname Fn4308 github.com/goccy/pythonwasm2go/p0.Fn4308
func Fn4308(m *base.Module, l0 int32)

//go:linkname Fn4309 github.com/goccy/pythonwasm2go/p0.Fn4309
func Fn4309(m *base.Module, l0 int32)

//go:linkname Fn4310 github.com/goccy/pythonwasm2go/p0.Fn4310
func Fn4310(m *base.Module, l0 int32)

//go:linkname Fn4311 github.com/goccy/pythonwasm2go/p0.Fn4311
func Fn4311(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4313 github.com/goccy/pythonwasm2go/p1.Fn4313
func Fn4313(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4318 github.com/goccy/pythonwasm2go/p0.Fn4318
func Fn4318(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4326 github.com/goccy/pythonwasm2go/p0.Fn4326
func Fn4326(m *base.Module, l0 int32)

//go:linkname Fn4327 github.com/goccy/pythonwasm2go/p0.Fn4327
func Fn4327(m *base.Module, l0 int32)

//go:linkname Fn4328 github.com/goccy/pythonwasm2go/p0.Fn4328
func Fn4328(m *base.Module, l0 int32)

//go:linkname Fn4329 github.com/goccy/pythonwasm2go/p0.Fn4329
func Fn4329(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4339 github.com/goccy/pythonwasm2go/p0.Fn4339
func Fn4339(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4340 github.com/goccy/pythonwasm2go/p0.Fn4340
func Fn4340(m *base.Module, l0 int32) int32

//go:linkname Fn4341 github.com/goccy/pythonwasm2go/p0.Fn4341
func Fn4341(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4343 github.com/goccy/pythonwasm2go/p0.Fn4343
func Fn4343(m *base.Module, l0 int32)

//go:linkname Fn4344 github.com/goccy/pythonwasm2go/p0.Fn4344
func Fn4344(m *base.Module, l0 int32)

//go:linkname Fn4353 github.com/goccy/pythonwasm2go/p0.Fn4353
func Fn4353(m *base.Module, l0 int32)

//go:linkname Fn4354 github.com/goccy/pythonwasm2go/p0.Fn4354
func Fn4354(m *base.Module, l0 int32)

//go:linkname Fn4355 github.com/goccy/pythonwasm2go/p0.Fn4355
func Fn4355(m *base.Module, l0 int32) int32

//go:linkname Fn4356 github.com/goccy/pythonwasm2go/p0.Fn4356
func Fn4356(m *base.Module, l0 int32)

//go:linkname Fn4365 github.com/goccy/pythonwasm2go/p0.Fn4365
func Fn4365(m *base.Module) int32

//go:linkname Fn4366 github.com/goccy/pythonwasm2go/p0.Fn4366
func Fn4366(m *base.Module, l0 int32)

//go:linkname Fn4368 github.com/goccy/pythonwasm2go/p0.Fn4368
func Fn4368(m *base.Module) int32

//go:linkname Fn4374 github.com/goccy/pythonwasm2go/p1.Fn4374
func Fn4374(m *base.Module, l0 int32)

//go:linkname Fn4378 github.com/goccy/pythonwasm2go/p1.Fn4378
func Fn4378(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)

//go:linkname Fn4382 github.com/goccy/pythonwasm2go/p1.Fn4382
func Fn4382(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4394 github.com/goccy/pythonwasm2go/p1.Fn4394
func Fn4394(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4395 github.com/goccy/pythonwasm2go/p1.Fn4395
func Fn4395(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4413 github.com/goccy/pythonwasm2go/p1.Fn4413
func Fn4413(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4423 github.com/goccy/pythonwasm2go/p1.Fn4423
func Fn4423(m *base.Module, l0 int32) int32

//go:linkname Fn4426 github.com/goccy/pythonwasm2go/p0.Fn4426
func Fn4426(m *base.Module, l0 int32) int32

//go:linkname Fn4431 github.com/goccy/pythonwasm2go/p0.Fn4431
func Fn4431(m *base.Module, l0 int32) int32

//go:linkname Fn4476 github.com/goccy/pythonwasm2go/p0.Fn4476
func Fn4476(m *base.Module, l0 int32) int32

//go:linkname Fn4486 github.com/goccy/pythonwasm2go/p0.Fn4486
func Fn4486(m *base.Module, l0 int32) int32

//go:linkname Fn4489 github.com/goccy/pythonwasm2go/p0.Fn4489
func Fn4489(m *base.Module, l0 int32) int32

//go:linkname Fn4490 github.com/goccy/pythonwasm2go/p0.Fn4490
func Fn4490(m *base.Module, l0 int32) int32

//go:linkname Fn4491 github.com/goccy/pythonwasm2go/p0.Fn4491
func Fn4491(m *base.Module, l0 int32)

//go:linkname Fn4505 github.com/goccy/pythonwasm2go/p0.Fn4505
func Fn4505(m *base.Module, l0 int32) int32

//go:linkname Fn4510 github.com/goccy/pythonwasm2go/p0.Fn4510
func Fn4510(m *base.Module, l0 int32) int32

//go:linkname Fn4585 github.com/goccy/pythonwasm2go/p1.Fn4585
func Fn4585(m *base.Module, l0 int32) int32

//go:linkname Fn4587 github.com/goccy/pythonwasm2go/p1.Fn4587
func Fn4587(m *base.Module, l0 int32) int32

//go:linkname Fn4588 github.com/goccy/pythonwasm2go/p1.Fn4588
func Fn4588(m *base.Module, l0 int32) int32

//go:linkname Fn4589 github.com/goccy/pythonwasm2go/p1.Fn4589
func Fn4589(m *base.Module, l0 int32) int32

//go:linkname Fn4595 github.com/goccy/pythonwasm2go/p1.Fn4595
func Fn4595(m *base.Module, l0 int32) int32

//go:linkname Fn4639 github.com/goccy/pythonwasm2go/p1.Fn4639
func Fn4639(m *base.Module, l0 int32) int32

//go:linkname Fn4649 github.com/goccy/pythonwasm2go/p1.Fn4649
func Fn4649(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4651 github.com/goccy/pythonwasm2go/p1.Fn4651
func Fn4651(m *base.Module, l0 int32) int32

//go:linkname Fn4653 github.com/goccy/pythonwasm2go/p1.Fn4653
func Fn4653(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn4655 github.com/goccy/pythonwasm2go/p1.Fn4655
func Fn4655(m *base.Module, l0 int32) int32

//go:linkname Fn4658 github.com/goccy/pythonwasm2go/p1.Fn4658
func Fn4658(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4664 github.com/goccy/pythonwasm2go/p1.Fn4664
func Fn4664(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn4669 github.com/goccy/pythonwasm2go/p0.Fn4669
func Fn4669(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4672 github.com/goccy/pythonwasm2go/p0.Fn4672
func Fn4672(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4676 github.com/goccy/pythonwasm2go/p0.Fn4676
func Fn4676(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4690 github.com/goccy/pythonwasm2go/p1.Fn4690
func Fn4690(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4694 github.com/goccy/pythonwasm2go/p0.Fn4694
func Fn4694(m *base.Module)

//go:linkname Fn4695 github.com/goccy/pythonwasm2go/p0.Fn4695
func Fn4695(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4697 github.com/goccy/pythonwasm2go/p1.Fn4697
func Fn4697(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4709 github.com/goccy/pythonwasm2go/p0.Fn4709
func Fn4709(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4712 github.com/goccy/pythonwasm2go/p0.Fn4712
func Fn4712(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4713 github.com/goccy/pythonwasm2go/p0.Fn4713
func Fn4713(m *base.Module, l0 int32) int32

//go:linkname Fn4718 github.com/goccy/pythonwasm2go/p0.Fn4718
func Fn4718(m *base.Module, l0 int64) int64

//go:linkname Fn4720 github.com/goccy/pythonwasm2go/p1.Fn4720
func Fn4720(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4748 github.com/goccy/pythonwasm2go/p1.Fn4748
func Fn4748(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4749 github.com/goccy/pythonwasm2go/p1.Fn4749
func Fn4749(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4752 github.com/goccy/pythonwasm2go/p1.Fn4752
func Fn4752(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4755 github.com/goccy/pythonwasm2go/p0.Fn4755
func Fn4755(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4759 github.com/goccy/pythonwasm2go/p1.Fn4759
func Fn4759(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4762 github.com/goccy/pythonwasm2go/p1.Fn4762
func Fn4762(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4788 github.com/goccy/pythonwasm2go/p1.Fn4788
func Fn4788(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4793 github.com/goccy/pythonwasm2go/p0.Fn4793
func Fn4793(m *base.Module, l0 int32) int32

//go:linkname Fn4795 github.com/goccy/pythonwasm2go/p0.Fn4795
func Fn4795(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4799 github.com/goccy/pythonwasm2go/p0.Fn4799
func Fn4799(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4803 github.com/goccy/pythonwasm2go/p0.Fn4803
func Fn4803(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4804 github.com/goccy/pythonwasm2go/p0.Fn4804
func Fn4804(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4806 github.com/goccy/pythonwasm2go/p0.Fn4806
func Fn4806(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4814 github.com/goccy/pythonwasm2go/p0.Fn4814
func Fn4814(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4815 github.com/goccy/pythonwasm2go/p0.Fn4815
func Fn4815(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4880 github.com/goccy/pythonwasm2go/p0.Fn4880
func Fn4880(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn4888 github.com/goccy/pythonwasm2go/p0.Fn4888
func Fn4888(m *base.Module, l0 int32, l1 int64) int32

//go:linkname Fn4890 github.com/goccy/pythonwasm2go/p0.Fn4890
func Fn4890(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4896 github.com/goccy/pythonwasm2go/p0.Fn4896
func Fn4896(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4901 github.com/goccy/pythonwasm2go/p0.Fn4901
func Fn4901(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4905 github.com/goccy/pythonwasm2go/p0.Fn4905
func Fn4905(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4906 github.com/goccy/pythonwasm2go/p0.Fn4906
func Fn4906(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4907 github.com/goccy/pythonwasm2go/p0.Fn4907
func Fn4907(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4909 github.com/goccy/pythonwasm2go/p0.Fn4909
func Fn4909(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4912 github.com/goccy/pythonwasm2go/p0.Fn4912
func Fn4912(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4928 github.com/goccy/pythonwasm2go/p1.Fn4928
func Fn4928(m *base.Module, l0 int32) int32

//go:linkname Fn4942 github.com/goccy/pythonwasm2go/p1.Fn4942
func Fn4942(m *base.Module) int32

//go:linkname Fn4943 github.com/goccy/pythonwasm2go/p0.Fn4943
func Fn4943(m *base.Module)

//go:linkname Fn4946 github.com/goccy/pythonwasm2go/p0.Fn4946
func Fn4946(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4953 github.com/goccy/pythonwasm2go/p1.Fn4953
func Fn4953(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4968 github.com/goccy/pythonwasm2go/p1.Fn4968
func Fn4968(m *base.Module, l0 float64, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4970 github.com/goccy/pythonwasm2go/p1.Fn4970
func Fn4970(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4972 github.com/goccy/pythonwasm2go/p0.Fn4972
func Fn4972(m *base.Module, l0 int32, l1 int32) float64

//go:linkname Fn4992 github.com/goccy/pythonwasm2go/p1.Fn4992
func Fn4992(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4994 github.com/goccy/pythonwasm2go/p1.Fn4994
func Fn4994(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn5001 github.com/goccy/pythonwasm2go/p1.Fn5001
func Fn5001(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5007 github.com/goccy/pythonwasm2go/p1.Fn5007
func Fn5007(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5010 github.com/goccy/pythonwasm2go/p1.Fn5010
func Fn5010(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5015 github.com/goccy/pythonwasm2go/p1.Fn5015
func Fn5015(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn5016 github.com/goccy/pythonwasm2go/p1.Fn5016
func Fn5016(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5018 github.com/goccy/pythonwasm2go/p1.Fn5018
func Fn5018(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn5031 github.com/goccy/pythonwasm2go/p0.Fn5031
func Fn5031(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5032 github.com/goccy/pythonwasm2go/p0.Fn5032
func Fn5032(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5037 github.com/goccy/pythonwasm2go/p1.Fn5037
func Fn5037(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5044 github.com/goccy/pythonwasm2go/p1.Fn5044
func Fn5044(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn5086 github.com/goccy/pythonwasm2go/p1.Fn5086
func Fn5086(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5176 github.com/goccy/pythonwasm2go/p1.Fn5176
func Fn5176(m *base.Module, l0 int32) int32

//go:linkname Fn5211 github.com/goccy/pythonwasm2go/p0.Fn5211
func Fn5211(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5215 github.com/goccy/pythonwasm2go/p1.Fn5215
func Fn5215(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5219 github.com/goccy/pythonwasm2go/p1.Fn5219
func Fn5219(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5233 github.com/goccy/pythonwasm2go/p1.Fn5233
func Fn5233(m *base.Module, l0 int32) int32

//go:linkname Fn5238 github.com/goccy/pythonwasm2go/p1.Fn5238
func Fn5238(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5322 github.com/goccy/pythonwasm2go/p1.Fn5322
func Fn5322(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5326 github.com/goccy/pythonwasm2go/p1.Fn5326
func Fn5326(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn5359 github.com/goccy/pythonwasm2go/p1.Fn5359
func Fn5359(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5366 github.com/goccy/pythonwasm2go/p0.Fn5366
func Fn5366(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5400 github.com/goccy/pythonwasm2go/p1.Fn5400
func Fn5400(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5407 github.com/goccy/pythonwasm2go/p1.Fn5407
func Fn5407(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5408 github.com/goccy/pythonwasm2go/p1.Fn5408
func Fn5408(m *base.Module, l0 int32) int64

//go:linkname Fn5418 github.com/goccy/pythonwasm2go/p1.Fn5418
func Fn5418(m *base.Module, l0 int32)

//go:linkname Fn5434 github.com/goccy/pythonwasm2go/p0.Fn5434
func Fn5434(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5438 github.com/goccy/pythonwasm2go/p1.Fn5438
func Fn5438(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5439 github.com/goccy/pythonwasm2go/p0.Fn5439
func Fn5439(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5462 github.com/goccy/pythonwasm2go/p1.Fn5462
func Fn5462(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5473 github.com/goccy/pythonwasm2go/p1.Fn5473
func Fn5473(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5492 github.com/goccy/pythonwasm2go/p1.Fn5492
func Fn5492(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5510 github.com/goccy/pythonwasm2go/p1.Fn5510
func Fn5510(m *base.Module, l0 int32) int32

//go:linkname Fn5523 github.com/goccy/pythonwasm2go/p1.Fn5523
func Fn5523(m *base.Module, l0 int32) int32

//go:linkname Fn5559 github.com/goccy/pythonwasm2go/p1.Fn5559
func Fn5559(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5579 github.com/goccy/pythonwasm2go/p1.Fn5579
func Fn5579(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5584 github.com/goccy/pythonwasm2go/p1.Fn5584
func Fn5584(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5689 github.com/goccy/pythonwasm2go/p1.Fn5689
func Fn5689(m *base.Module, l0 int32) int32

//go:linkname Fn5697 github.com/goccy/pythonwasm2go/p1.Fn5697
func Fn5697(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5705 github.com/goccy/pythonwasm2go/p1.Fn5705
func Fn5705(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5707 github.com/goccy/pythonwasm2go/p0.Fn5707
func Fn5707(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5709 github.com/goccy/pythonwasm2go/p1.Fn5709
func Fn5709(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5743 github.com/goccy/pythonwasm2go/p1.Fn5743
func Fn5743(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32) float64

//go:linkname Fn5808 github.com/goccy/pythonwasm2go/p1.Fn5808
func Fn5808(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5814 github.com/goccy/pythonwasm2go/p1.Fn5814
func Fn5814(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5817 github.com/goccy/pythonwasm2go/p1.Fn5817
func Fn5817(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5832 github.com/goccy/pythonwasm2go/p1.Fn5832
func Fn5832(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5836 github.com/goccy/pythonwasm2go/p1.Fn5836
func Fn5836(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5850 github.com/goccy/pythonwasm2go/p1.Fn5850
func Fn5850(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5851 github.com/goccy/pythonwasm2go/p1.Fn5851
func Fn5851(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5858 github.com/goccy/pythonwasm2go/p1.Fn5858
func Fn5858(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5904 github.com/goccy/pythonwasm2go/p1.Fn5904
func Fn5904(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5905 github.com/goccy/pythonwasm2go/p1.Fn5905
func Fn5905(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5912 github.com/goccy/pythonwasm2go/p1.Fn5912
func Fn5912(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5920 github.com/goccy/pythonwasm2go/p1.Fn5920
func Fn5920(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5934 github.com/goccy/pythonwasm2go/p1.Fn5934
func Fn5934(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5943 github.com/goccy/pythonwasm2go/p1.Fn5943
func Fn5943(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn5954 github.com/goccy/pythonwasm2go/p1.Fn5954
func Fn5954(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn5959 github.com/goccy/pythonwasm2go/p1.Fn5959
func Fn5959(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5965 github.com/goccy/pythonwasm2go/p1.Fn5965
func Fn5965(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5971 github.com/goccy/pythonwasm2go/p1.Fn5971
func Fn5971(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5974 github.com/goccy/pythonwasm2go/p1.Fn5974
func Fn5974(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5975 github.com/goccy/pythonwasm2go/p1.Fn5975
func Fn5975(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5978 github.com/goccy/pythonwasm2go/p1.Fn5978
func Fn5978(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6001 github.com/goccy/pythonwasm2go/p1.Fn6001
func Fn6001(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6007 github.com/goccy/pythonwasm2go/p1.Fn6007
func Fn6007(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn6023 github.com/goccy/pythonwasm2go/p1.Fn6023
func Fn6023(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6027 github.com/goccy/pythonwasm2go/p1.Fn6027
func Fn6027(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6041 github.com/goccy/pythonwasm2go/p1.Fn6041
func Fn6041(m *base.Module, l0 int32) int32

//go:linkname Fn6076 github.com/goccy/pythonwasm2go/p1.Fn6076
func Fn6076(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6077 github.com/goccy/pythonwasm2go/p1.Fn6077
func Fn6077(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6097 github.com/goccy/pythonwasm2go/p0.Fn6097
func Fn6097(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6101 github.com/goccy/pythonwasm2go/p0.Fn6101
func Fn6101(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn6288 github.com/goccy/pythonwasm2go/p0.Fn6288
func Fn6288(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6307 github.com/goccy/pythonwasm2go/p1.Fn6307
func Fn6307(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6324 github.com/goccy/pythonwasm2go/p0.Fn6324
func Fn6324(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6333 github.com/goccy/pythonwasm2go/p0.Fn6333
func Fn6333(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6334 github.com/goccy/pythonwasm2go/p1.Fn6334
func Fn6334(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6335 github.com/goccy/pythonwasm2go/p1.Fn6335
func Fn6335(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6378 github.com/goccy/pythonwasm2go/p1.Fn6378
func Fn6378(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6381 github.com/goccy/pythonwasm2go/p1.Fn6381
func Fn6381(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6413 github.com/goccy/pythonwasm2go/p0.Fn6413
func Fn6413(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn6421 github.com/goccy/pythonwasm2go/p0.Fn6421
func Fn6421(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32)

//go:linkname Fn6436 github.com/goccy/pythonwasm2go/p1.Fn6436
func Fn6436(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6444 github.com/goccy/pythonwasm2go/p1.Fn6444
func Fn6444(m *base.Module, l0 int32) int32

//go:linkname Fn6447 github.com/goccy/pythonwasm2go/p1.Fn6447
func Fn6447(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32)

//go:linkname Fn6448 github.com/goccy/pythonwasm2go/p1.Fn6448
func Fn6448(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6449 github.com/goccy/pythonwasm2go/p1.Fn6449
func Fn6449(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32)

//go:linkname Fn6450 github.com/goccy/pythonwasm2go/p1.Fn6450
func Fn6450(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6452 github.com/goccy/pythonwasm2go/p1.Fn6452
func Fn6452(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn6453 github.com/goccy/pythonwasm2go/p1.Fn6453
func Fn6453(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn6454 github.com/goccy/pythonwasm2go/p1.Fn6454
func Fn6454(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn6455 github.com/goccy/pythonwasm2go/p1.Fn6455
func Fn6455(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn6456 github.com/goccy/pythonwasm2go/p1.Fn6456
func Fn6456(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn6457 github.com/goccy/pythonwasm2go/p1.Fn6457
func Fn6457(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn6458 github.com/goccy/pythonwasm2go/p1.Fn6458
func Fn6458(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn6459 github.com/goccy/pythonwasm2go/p1.Fn6459
func Fn6459(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn6460 github.com/goccy/pythonwasm2go/p1.Fn6460
func Fn6460(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn6461 github.com/goccy/pythonwasm2go/p1.Fn6461
func Fn6461(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn6462 github.com/goccy/pythonwasm2go/p1.Fn6462
func Fn6462(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn6463 github.com/goccy/pythonwasm2go/p1.Fn6463
func Fn6463(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn6512 github.com/goccy/pythonwasm2go/p1.Fn6512
func Fn6512(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6525 github.com/goccy/pythonwasm2go/p1.Fn6525
func Fn6525(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn6529 github.com/goccy/pythonwasm2go/p1.Fn6529
func Fn6529(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32

//go:linkname Fn6542 github.com/goccy/pythonwasm2go/p1.Fn6542
func Fn6542(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6544 github.com/goccy/pythonwasm2go/p1.Fn6544
func Fn6544(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6665 github.com/goccy/pythonwasm2go/p1.Fn6665
func Fn6665(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn6672 github.com/goccy/pythonwasm2go/p1.Fn6672
func Fn6672(m *base.Module, l0 int32)

//go:linkname Fn6680 github.com/goccy/pythonwasm2go/p1.Fn6680
func Fn6680(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6681 github.com/goccy/pythonwasm2go/p1.Fn6681
func Fn6681(m *base.Module, l0 int32) int32

//go:linkname Fn6683 github.com/goccy/pythonwasm2go/p1.Fn6683
func Fn6683(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6691 github.com/goccy/pythonwasm2go/p1.Fn6691
func Fn6691(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6724 github.com/goccy/pythonwasm2go/p1.Fn6724
func Fn6724(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6725 github.com/goccy/pythonwasm2go/p1.Fn6725
func Fn6725(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6739 github.com/goccy/pythonwasm2go/p1.Fn6739
func Fn6739(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6740 github.com/goccy/pythonwasm2go/p0.Fn6740
func Fn6740(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32

//go:linkname Fn6743 github.com/goccy/pythonwasm2go/p1.Fn6743
func Fn6743(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6752 github.com/goccy/pythonwasm2go/p1.Fn6752
func Fn6752(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn6754 github.com/goccy/pythonwasm2go/p1.Fn6754
func Fn6754(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn6765 github.com/goccy/pythonwasm2go/p0.Fn6765
func Fn6765(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn6780 github.com/goccy/pythonwasm2go/p1.Fn6780
func Fn6780(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn6795 github.com/goccy/pythonwasm2go/p1.Fn6795
func Fn6795(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn6865 github.com/goccy/pythonwasm2go/p1.Fn6865
func Fn6865(m *base.Module, l0 int32) int32

//go:linkname Fn6883 github.com/goccy/pythonwasm2go/p1.Fn6883
func Fn6883(m *base.Module, l0 int32) int32

//go:linkname Fn6891 github.com/goccy/pythonwasm2go/p1.Fn6891
func Fn6891(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6903 github.com/goccy/pythonwasm2go/p1.Fn6903
func Fn6903(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn6941 github.com/goccy/pythonwasm2go/p1.Fn6941
func Fn6941(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6954 github.com/goccy/pythonwasm2go/p1.Fn6954
func Fn6954(m *base.Module, l0 int32) int32

//go:linkname Fn6975 github.com/goccy/pythonwasm2go/p1.Fn6975
func Fn6975(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn6977 github.com/goccy/pythonwasm2go/p1.Fn6977
func Fn6977(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7029 github.com/goccy/pythonwasm2go/p1.Fn7029
func Fn7029(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7080 github.com/goccy/pythonwasm2go/p1.Fn7080
func Fn7080(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn7083 github.com/goccy/pythonwasm2go/p1.Fn7083
func Fn7083(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7093 github.com/goccy/pythonwasm2go/p1.Fn7093
func Fn7093(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7111 github.com/goccy/pythonwasm2go/p1.Fn7111
func Fn7111(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7133 github.com/goccy/pythonwasm2go/p1.Fn7133
func Fn7133(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7135 github.com/goccy/pythonwasm2go/p1.Fn7135
func Fn7135(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7174 github.com/goccy/pythonwasm2go/p1.Fn7174
func Fn7174(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7210 github.com/goccy/pythonwasm2go/p1.Fn7210
func Fn7210(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7220 github.com/goccy/pythonwasm2go/p1.Fn7220
func Fn7220(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int64) int32

//go:linkname Fn7284 github.com/goccy/pythonwasm2go/p1.Fn7284
func Fn7284(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7285 github.com/goccy/pythonwasm2go/p1.Fn7285
func Fn7285(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32)

//go:linkname Fn7434 github.com/goccy/pythonwasm2go/p1.Fn7434
func Fn7434(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7763 github.com/goccy/pythonwasm2go/p1.Fn7763
func Fn7763(m *base.Module, l0 int32) int32

//go:linkname Fn8556 github.com/goccy/pythonwasm2go/p0.Fn8556
func Fn8556(m *base.Module, l0 int32)

//go:linkname Fn8557 github.com/goccy/pythonwasm2go/p0.Fn8557
func Fn8557(m *base.Module, l0 int32)

//go:linkname Fn8572 github.com/goccy/pythonwasm2go/p1.Fn8572
func Fn8572(m *base.Module, l0 int32) int32

//go:linkname Fn8595 github.com/goccy/pythonwasm2go/p1.Fn8595
func Fn8595(m *base.Module, l0 int32) int32

//go:linkname Fn8598 github.com/goccy/pythonwasm2go/p1.Fn8598
func Fn8598(m *base.Module, l0 int32) int32

//go:linkname Fn8601 github.com/goccy/pythonwasm2go/p1.Fn8601
func Fn8601(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8653 github.com/goccy/pythonwasm2go/p1.Fn8653
func Fn8653(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8709 github.com/goccy/pythonwasm2go/p1.Fn8709
func Fn8709(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8712 github.com/goccy/pythonwasm2go/p1.Fn8712
func Fn8712(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8714 github.com/goccy/pythonwasm2go/p1.Fn8714
func Fn8714(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8721 github.com/goccy/pythonwasm2go/p0.Fn8721
func Fn8721(m *base.Module) int32

//go:linkname Fn8723 github.com/goccy/pythonwasm2go/p0.Fn8723
func Fn8723(m *base.Module, l0 int32) int32

//go:linkname Fn8773 github.com/goccy/pythonwasm2go/p1.Fn8773
func Fn8773(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8790 github.com/goccy/pythonwasm2go/p1.Fn8790
func Fn8790(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8796 github.com/goccy/pythonwasm2go/p1.Fn8796
func Fn8796(m *base.Module, l0 int32) int32

//go:linkname Fn8832 github.com/goccy/pythonwasm2go/p1.Fn8832
func Fn8832(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8850 github.com/goccy/pythonwasm2go/p1.Fn8850
func Fn8850(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8942 github.com/goccy/pythonwasm2go/p1.Fn8942
func Fn8942(m *base.Module, l0 int32)

//go:linkname Fn9018 github.com/goccy/pythonwasm2go/p1.Fn9018
func Fn9018(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9042 github.com/goccy/pythonwasm2go/p1.Fn9042
func Fn9042(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9083 github.com/goccy/pythonwasm2go/p1.Fn9083
func Fn9083(m *base.Module, l0 int32) int32

//go:linkname Fn9134 github.com/goccy/pythonwasm2go/p1.Fn9134
func Fn9134(m *base.Module, l0 int32) int32

//go:linkname Fn9172 github.com/goccy/pythonwasm2go/p1.Fn9172
func Fn9172(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9203 github.com/goccy/pythonwasm2go/p1.Fn9203
func Fn9203(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9211 github.com/goccy/pythonwasm2go/p1.Fn9211
func Fn9211(m *base.Module, l0 int32, l1 int64, l2 int32) int64

//go:linkname Fn9213 github.com/goccy/pythonwasm2go/p1.Fn9213
func Fn9213(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9260 github.com/goccy/pythonwasm2go/p1.Fn9260
func Fn9260(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9262 github.com/goccy/pythonwasm2go/p1.Fn9262
func Fn9262(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn9265 github.com/goccy/pythonwasm2go/p1.Fn9265
func Fn9265(m *base.Module, l0 int32) int32

//go:linkname Fn9269 github.com/goccy/pythonwasm2go/p1.Fn9269
func Fn9269(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9287 github.com/goccy/pythonwasm2go/p1.Fn9287
func Fn9287(m *base.Module, l0 int32) int32

//go:linkname Fn9353 github.com/goccy/pythonwasm2go/p1.Fn9353
func Fn9353(m *base.Module, l0 int32) int32

//go:linkname Fn9458 github.com/goccy/pythonwasm2go/p1.Fn9458
func Fn9458(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9462 github.com/goccy/pythonwasm2go/p0.Fn9462
func Fn9462(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9464 github.com/goccy/pythonwasm2go/p1.Fn9464
func Fn9464(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn9475 github.com/goccy/pythonwasm2go/p0.Fn9475
func Fn9475(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9476 github.com/goccy/pythonwasm2go/p0.Fn9476
func Fn9476(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9477 github.com/goccy/pythonwasm2go/p0.Fn9477
func Fn9477(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9491 github.com/goccy/pythonwasm2go/p1.Fn9491
func Fn9491(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9581 github.com/goccy/pythonwasm2go/p1.Fn9581
func Fn9581(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9600 github.com/goccy/pythonwasm2go/p1.Fn9600
func Fn9600(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9833 github.com/goccy/pythonwasm2go/p1.Fn9833
func Fn9833(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9856 github.com/goccy/pythonwasm2go/p1.Fn9856
func Fn9856(m *base.Module, l0 int32, l1 int64, l2 int32)

//go:linkname Fn9858 github.com/goccy/pythonwasm2go/p1.Fn9858
func Fn9858(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9886 github.com/goccy/pythonwasm2go/p1.Fn9886
func Fn9886(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9887 github.com/goccy/pythonwasm2go/p1.Fn9887
func Fn9887(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9893 github.com/goccy/pythonwasm2go/p1.Fn9893
func Fn9893(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn9913 github.com/goccy/pythonwasm2go/p1.Fn9913
func Fn9913(m *base.Module)

//go:linkname Fn9924 github.com/goccy/pythonwasm2go/p1.Fn9924
func Fn9924(m *base.Module, l0 int32)

//go:linkname Fn9934 github.com/goccy/pythonwasm2go/p1.Fn9934
func Fn9934(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn9948 github.com/goccy/pythonwasm2go/p1.Fn9948
func Fn9948(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9950 github.com/goccy/pythonwasm2go/p1.Fn9950
func Fn9950(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9956 github.com/goccy/pythonwasm2go/p1.Fn9956
func Fn9956(m *base.Module, l0 float64) float64

//go:linkname Fn9957 github.com/goccy/pythonwasm2go/p1.Fn9957
func Fn9957(m *base.Module, l0 float64, l1 float64) float64

//go:linkname Fn9961 github.com/goccy/pythonwasm2go/p0.Fn9961
func Fn9961(m *base.Module, l0 float64, l1 int32) int32

//go:linkname Fn9972 github.com/goccy/pythonwasm2go/p1.Fn9972
func Fn9972(m *base.Module, l0 float64) float64

//go:linkname Fn9975 github.com/goccy/pythonwasm2go/p1.Fn9975
func Fn9975(m *base.Module, l0 float64, l1 float64, l2 float64) float64

//go:linkname Fn9976 github.com/goccy/pythonwasm2go/p1.Fn9976
func Fn9976(m *base.Module, l0 float64, l1 float64) float64

//go:linkname Fn9978 github.com/goccy/pythonwasm2go/p1.Fn9978
func Fn9978(m *base.Module, l0 float64, l1 float64) float64

//go:linkname Fn9981 github.com/goccy/pythonwasm2go/p1.Fn9981
func Fn9981(m *base.Module, l0 float64) float64

//go:linkname Fn9986 github.com/goccy/pythonwasm2go/p1.Fn9986
func Fn9986(m *base.Module, l0 float64, l1 float64) float64

//go:linkname Fn9995 github.com/goccy/pythonwasm2go/p1.Fn9995
func Fn9995(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10038 github.com/goccy/pythonwasm2go/p1.Fn10038
func Fn10038(m *base.Module, l0 int32, l1 int32, l2 int64) int64

//go:linkname Fn10043 github.com/goccy/pythonwasm2go/p1.Fn10043
func Fn10043(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10047 github.com/goccy/pythonwasm2go/p1.Fn10047
func Fn10047(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10049 github.com/goccy/pythonwasm2go/p1.Fn10049
func Fn10049(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10056 github.com/goccy/pythonwasm2go/p1.Fn10056
func Fn10056(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10073 github.com/goccy/pythonwasm2go/p1.Fn10073
func Fn10073(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10077 github.com/goccy/pythonwasm2go/p0.Fn10077
func Fn10077(m *base.Module, l0 int32) int32

//go:linkname Fn10078 github.com/goccy/pythonwasm2go/p1.Fn10078
func Fn10078(m *base.Module, l0 int32)

//go:linkname Fn10080 github.com/goccy/pythonwasm2go/p1.Fn10080
func Fn10080(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10095 github.com/goccy/pythonwasm2go/p1.Fn10095
func Fn10095(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn10121 github.com/goccy/pythonwasm2go/p1.Fn10121
func Fn10121(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10151 github.com/goccy/pythonwasm2go/p0.Fn10151
func Fn10151(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn10153 github.com/goccy/pythonwasm2go/p0.Fn10153
func Fn10153(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn10160 github.com/goccy/pythonwasm2go/p0.Fn10160
func Fn10160(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10164 github.com/goccy/pythonwasm2go/p0.Fn10164
func Fn10164(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10168 github.com/goccy/pythonwasm2go/p0.Fn10168
func Fn10168(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10174 github.com/goccy/pythonwasm2go/p1.Fn10174
func Fn10174(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn10186 github.com/goccy/pythonwasm2go/p1.Fn10186
func Fn10186(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn10216 github.com/goccy/pythonwasm2go/p0.Fn10216
func Fn10216(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10226 github.com/goccy/pythonwasm2go/p1.Fn10226
func Fn10226(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn10230 github.com/goccy/pythonwasm2go/p1.Fn10230
func Fn10230(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10255 github.com/goccy/pythonwasm2go/p0.Fn10255
func Fn10255(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn10289 github.com/goccy/pythonwasm2go/p1.Fn10289
func Fn10289(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn10293 github.com/goccy/pythonwasm2go/p1.Fn10293
func Fn10293(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10295 github.com/goccy/pythonwasm2go/p0.Fn10295
func Fn10295(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
