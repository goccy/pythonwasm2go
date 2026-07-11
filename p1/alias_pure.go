//go:build (!amd64 && !arm64) || purego

package p1

import (
	base "github.com/goccy/pythonwasm2go/base"
	_ "unsafe"
)

//go:linkname Fn97 github.com/goccy/pythonwasm2go/p2.Fn97
func Fn97(m *base.Module, l0 int32) int32

//go:linkname Fn98 github.com/goccy/pythonwasm2go/p2.Fn98
func Fn98(m *base.Module, l0 int32) int64

//go:linkname Fn99 github.com/goccy/pythonwasm2go/p2.Fn99
func Fn99(m *base.Module, l0 int32)

//go:linkname Fn101 github.com/goccy/pythonwasm2go/p2.Fn101
func Fn101(m *base.Module, l0 int32) int64

//go:linkname Fn102 github.com/goccy/pythonwasm2go/p2.Fn102
func Fn102(m *base.Module, l0 int32)

//go:linkname Fn104 github.com/goccy/pythonwasm2go/p2.Fn104
func Fn104(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn105 github.com/goccy/pythonwasm2go/p2.Fn105
func Fn105(m *base.Module, l0 int32, l1 int64)

//go:linkname Fn108 github.com/goccy/pythonwasm2go/p2.Fn108
func Fn108(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn109 github.com/goccy/pythonwasm2go/p2.Fn109
func Fn109(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn110 github.com/goccy/pythonwasm2go/p2.Fn110
func Fn110(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn111 github.com/goccy/pythonwasm2go/p2.Fn111
func Fn111(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn121 github.com/goccy/pythonwasm2go/p2.Fn121
func Fn121(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn122 github.com/goccy/pythonwasm2go/p2.Fn122
func Fn122(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn123 github.com/goccy/pythonwasm2go/p2.Fn123
func Fn123(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn125 github.com/goccy/pythonwasm2go/p2.Fn125
func Fn125(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn126 github.com/goccy/pythonwasm2go/p2.Fn126
func Fn126(m *base.Module, l0 int32)

//go:linkname Fn128 github.com/goccy/pythonwasm2go/p2.Fn128
func Fn128(m *base.Module, l0 int32)

//go:linkname Fn129 github.com/goccy/pythonwasm2go/p2.Fn129
func Fn129(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn134 github.com/goccy/pythonwasm2go/p0.Fn134
func Fn134(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn135 github.com/goccy/pythonwasm2go/p2.Fn135
func Fn135(m *base.Module, l0 int32) int32

//go:linkname Fn138 github.com/goccy/pythonwasm2go/p0.Fn138
func Fn138(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn143 github.com/goccy/pythonwasm2go/p0.Fn143
func Fn143(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn146 github.com/goccy/pythonwasm2go/p2.Fn146
func Fn146(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn149 github.com/goccy/pythonwasm2go/p0.Fn149
func Fn149(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn151 github.com/goccy/pythonwasm2go/p0.Fn151
func Fn151(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn152 github.com/goccy/pythonwasm2go/p0.Fn152
func Fn152(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn154 github.com/goccy/pythonwasm2go/p2.Fn154
func Fn154(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn157 github.com/goccy/pythonwasm2go/p0.Fn157
func Fn157(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn160 github.com/goccy/pythonwasm2go/p2.Fn160
func Fn160(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn162 github.com/goccy/pythonwasm2go/p0.Fn162
func Fn162(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn164 github.com/goccy/pythonwasm2go/p0.Fn164
func Fn164(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn166 github.com/goccy/pythonwasm2go/p0.Fn166
func Fn166(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn168 github.com/goccy/pythonwasm2go/p0.Fn168
func Fn168(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn169 github.com/goccy/pythonwasm2go/p2.Fn169
func Fn169(m *base.Module, l0 int32) int32

//go:linkname Fn178 github.com/goccy/pythonwasm2go/p2.Fn178
func Fn178(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn182 github.com/goccy/pythonwasm2go/p2.Fn182
func Fn182(m *base.Module, l0 float64, l1 float64)

//go:linkname Fn183 github.com/goccy/pythonwasm2go/p2.Fn183
func Fn183(m *base.Module, l0 int32) float64

//go:linkname Fn184 github.com/goccy/pythonwasm2go/p0.Fn184
func Fn184(m *base.Module, l0 int32) int32

//go:linkname Fn186 github.com/goccy/pythonwasm2go/p2.Fn186
func Fn186(m *base.Module, l0 int32) float64

//go:linkname Fn187 github.com/goccy/pythonwasm2go/p2.Fn187
func Fn187(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn188 github.com/goccy/pythonwasm2go/p0.Fn188
func Fn188(m *base.Module, l0 int32) int32

//go:linkname Fn189 github.com/goccy/pythonwasm2go/p2.Fn189
func Fn189(m *base.Module, l0 int32) float64

//go:linkname Fn190 github.com/goccy/pythonwasm2go/p0.Fn190
func Fn190(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn194 github.com/goccy/pythonwasm2go/p2.Fn194
func Fn194(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn198 github.com/goccy/pythonwasm2go/p2.Fn198
func Fn198(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn202 github.com/goccy/pythonwasm2go/p2.Fn202
func Fn202(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn227 github.com/goccy/pythonwasm2go/p2.Fn227
func Fn227(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn228 github.com/goccy/pythonwasm2go/p2.Fn228
func Fn228(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn229 github.com/goccy/pythonwasm2go/p2.Fn229
func Fn229(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn230 github.com/goccy/pythonwasm2go/p2.Fn230
func Fn230(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn242 github.com/goccy/pythonwasm2go/p2.Fn242
func Fn242(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn243 github.com/goccy/pythonwasm2go/p2.Fn243
func Fn243(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn246 github.com/goccy/pythonwasm2go/p2.Fn246
func Fn246(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn247 github.com/goccy/pythonwasm2go/p0.Fn247
func Fn247(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn248 github.com/goccy/pythonwasm2go/p0.Fn248
func Fn248(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn250 github.com/goccy/pythonwasm2go/p0.Fn250
func Fn250(m *base.Module, l0 int32) int32

//go:linkname Fn251 github.com/goccy/pythonwasm2go/p0.Fn251
func Fn251(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn252 github.com/goccy/pythonwasm2go/p0.Fn252
func Fn252(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn254 github.com/goccy/pythonwasm2go/p2.Fn254
func Fn254(m *base.Module, l0 int32)

//go:linkname Fn255 github.com/goccy/pythonwasm2go/p2.Fn255
func Fn255(m *base.Module, l0 int32)

//go:linkname Fn256 github.com/goccy/pythonwasm2go/p0.Fn256
func Fn256(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn258 github.com/goccy/pythonwasm2go/p0.Fn258
func Fn258(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn260 github.com/goccy/pythonwasm2go/p2.Fn260
func Fn260(m *base.Module, l0 int32) int32

//go:linkname Fn262 github.com/goccy/pythonwasm2go/p2.Fn262
func Fn262(m *base.Module, l0 int32) int32

//go:linkname Fn263 github.com/goccy/pythonwasm2go/p2.Fn263
func Fn263(m *base.Module, l0 int32) int32

//go:linkname Fn264 github.com/goccy/pythonwasm2go/p2.Fn264
func Fn264(m *base.Module, l0 int32) int32

//go:linkname Fn266 github.com/goccy/pythonwasm2go/p0.Fn266
func Fn266(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn268 github.com/goccy/pythonwasm2go/p2.Fn268
func Fn268(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn269 github.com/goccy/pythonwasm2go/p2.Fn269
func Fn269(m *base.Module, l0 int32) int32

//go:linkname Fn270 github.com/goccy/pythonwasm2go/p2.Fn270
func Fn270(m *base.Module, l0 int32) int32

//go:linkname Fn271 github.com/goccy/pythonwasm2go/p2.Fn271
func Fn271(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn273 github.com/goccy/pythonwasm2go/p2.Fn273
func Fn273(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn275 github.com/goccy/pythonwasm2go/p2.Fn275
func Fn275(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn291 github.com/goccy/pythonwasm2go/p2.Fn291
func Fn291(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn293 github.com/goccy/pythonwasm2go/p2.Fn293
func Fn293(m *base.Module, l0 int32) int32

//go:linkname Fn303 github.com/goccy/pythonwasm2go/p2.Fn303
func Fn303(m *base.Module, l0 int32) int32

//go:linkname Fn304 github.com/goccy/pythonwasm2go/p2.Fn304
func Fn304(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn328 github.com/goccy/pythonwasm2go/p2.Fn328
func Fn328(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn329 github.com/goccy/pythonwasm2go/p2.Fn329
func Fn329(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn350 github.com/goccy/pythonwasm2go/p2.Fn350
func Fn350(m *base.Module, l0 int32) int32

//go:linkname Fn351 github.com/goccy/pythonwasm2go/p0.Fn351
func Fn351(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn352 github.com/goccy/pythonwasm2go/p2.Fn352
func Fn352(m *base.Module, l0 int32) int32

//go:linkname Fn353 github.com/goccy/pythonwasm2go/p0.Fn353
func Fn353(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn354 github.com/goccy/pythonwasm2go/p0.Fn354
func Fn354(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn355 github.com/goccy/pythonwasm2go/p0.Fn355
func Fn355(m *base.Module, l0 int32) int32

//go:linkname Fn363 github.com/goccy/pythonwasm2go/p2.Fn363
func Fn363(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn364 github.com/goccy/pythonwasm2go/p2.Fn364
func Fn364(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn388 github.com/goccy/pythonwasm2go/p2.Fn388
func Fn388(m *base.Module, l0 int32) int32

//go:linkname Fn389 github.com/goccy/pythonwasm2go/p2.Fn389
func Fn389(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn415 github.com/goccy/pythonwasm2go/p2.Fn415
func Fn415(m *base.Module, l0 int32) int32

//go:linkname Fn444 github.com/goccy/pythonwasm2go/p2.Fn444
func Fn444(m *base.Module, l0 int32) int32

//go:linkname Fn445 github.com/goccy/pythonwasm2go/p0.Fn445
func Fn445(m *base.Module)

//go:linkname Fn446 github.com/goccy/pythonwasm2go/p0.Fn446
func Fn446(m *base.Module, l0 int32) int32

//go:linkname Fn448 github.com/goccy/pythonwasm2go/p0.Fn448
func Fn448(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn450 github.com/goccy/pythonwasm2go/p0.Fn450
func Fn450(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn451 github.com/goccy/pythonwasm2go/p0.Fn451
func Fn451(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn452 github.com/goccy/pythonwasm2go/p0.Fn452
func Fn452(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn453 github.com/goccy/pythonwasm2go/p0.Fn453
func Fn453(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn454 github.com/goccy/pythonwasm2go/p0.Fn454
func Fn454(m *base.Module, l0 int32) int32

//go:linkname Fn455 github.com/goccy/pythonwasm2go/p0.Fn455
func Fn455(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn456 github.com/goccy/pythonwasm2go/p0.Fn456
func Fn456(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn457 github.com/goccy/pythonwasm2go/p2.Fn457
func Fn457(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn459 github.com/goccy/pythonwasm2go/p2.Fn459
func Fn459(m *base.Module, l0 int32) int32

//go:linkname Fn460 github.com/goccy/pythonwasm2go/p2.Fn460
func Fn460(m *base.Module, l0 int32)

//go:linkname Fn461 github.com/goccy/pythonwasm2go/p0.Fn461
func Fn461(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn463 github.com/goccy/pythonwasm2go/p2.Fn463
func Fn463(m *base.Module, l0 int32) int32

//go:linkname Fn464 github.com/goccy/pythonwasm2go/p2.Fn464
func Fn464(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn465 github.com/goccy/pythonwasm2go/p2.Fn465
func Fn465(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn468 github.com/goccy/pythonwasm2go/p0.Fn468
func Fn468(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn470 github.com/goccy/pythonwasm2go/p2.Fn470
func Fn470(m *base.Module, l0 int32) int32

//go:linkname Fn476 github.com/goccy/pythonwasm2go/p2.Fn476
func Fn476(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn479 github.com/goccy/pythonwasm2go/p2.Fn479
func Fn479(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn481 github.com/goccy/pythonwasm2go/p2.Fn481
func Fn481(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn482 github.com/goccy/pythonwasm2go/p2.Fn482
func Fn482(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn485 github.com/goccy/pythonwasm2go/p2.Fn485
func Fn485(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn486 github.com/goccy/pythonwasm2go/p2.Fn486
func Fn486(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn488 github.com/goccy/pythonwasm2go/p2.Fn488
func Fn488(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn505 github.com/goccy/pythonwasm2go/p2.Fn505
func Fn505(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn507 github.com/goccy/pythonwasm2go/p2.Fn507
func Fn507(m *base.Module, l0 int32) int32

//go:linkname Fn508 github.com/goccy/pythonwasm2go/p2.Fn508
func Fn508(m *base.Module, l0 int32) int32

//go:linkname Fn509 github.com/goccy/pythonwasm2go/p2.Fn509
func Fn509(m *base.Module, l0 int32) int32

//go:linkname Fn510 github.com/goccy/pythonwasm2go/p2.Fn510
func Fn510(m *base.Module, l0 int32) int32

//go:linkname Fn511 github.com/goccy/pythonwasm2go/p2.Fn511
func Fn511(m *base.Module, l0 int32) int32

//go:linkname Fn512 github.com/goccy/pythonwasm2go/p2.Fn512
func Fn512(m *base.Module, l0 int32) int32

//go:linkname Fn515 github.com/goccy/pythonwasm2go/p2.Fn515
func Fn515(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn516 github.com/goccy/pythonwasm2go/p2.Fn516
func Fn516(m *base.Module, l0 int32) int32

//go:linkname Fn517 github.com/goccy/pythonwasm2go/p0.Fn517
func Fn517(m *base.Module, l0 int32) int32

//go:linkname Fn518 github.com/goccy/pythonwasm2go/p2.Fn518
func Fn518(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn519 github.com/goccy/pythonwasm2go/p0.Fn519
func Fn519(m *base.Module, l0 int32) int32

//go:linkname Fn520 github.com/goccy/pythonwasm2go/p0.Fn520
func Fn520(m *base.Module, l0 int32) int32

//go:linkname Fn521 github.com/goccy/pythonwasm2go/p0.Fn521
func Fn521(m *base.Module, l0 int32) int32

//go:linkname Fn522 github.com/goccy/pythonwasm2go/p0.Fn522
func Fn522(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn523 github.com/goccy/pythonwasm2go/p0.Fn523
func Fn523(m *base.Module, l0 int32) int32

//go:linkname Fn524 github.com/goccy/pythonwasm2go/p0.Fn524
func Fn524(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn526 github.com/goccy/pythonwasm2go/p0.Fn526
func Fn526(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn527 github.com/goccy/pythonwasm2go/p2.Fn527
func Fn527(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn528 github.com/goccy/pythonwasm2go/p2.Fn528
func Fn528(m *base.Module, l0 int32) int32

//go:linkname Fn529 github.com/goccy/pythonwasm2go/p2.Fn529
func Fn529(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn530 github.com/goccy/pythonwasm2go/p2.Fn530
func Fn530(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn531 github.com/goccy/pythonwasm2go/p0.Fn531
func Fn531(m *base.Module, l0 int32) int32

//go:linkname Fn533 github.com/goccy/pythonwasm2go/p2.Fn533
func Fn533(m *base.Module, l0 int32) int32

//go:linkname Fn534 github.com/goccy/pythonwasm2go/p0.Fn534
func Fn534(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn538 github.com/goccy/pythonwasm2go/p0.Fn538
func Fn538(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn543 github.com/goccy/pythonwasm2go/p2.Fn543
func Fn543(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn544 github.com/goccy/pythonwasm2go/p2.Fn544
func Fn544(m *base.Module, l0 int32) int32

//go:linkname Fn546 github.com/goccy/pythonwasm2go/p2.Fn546
func Fn546(m *base.Module, l0 int32) int32

//go:linkname Fn547 github.com/goccy/pythonwasm2go/p0.Fn547
func Fn547(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn548 github.com/goccy/pythonwasm2go/p0.Fn548
func Fn548(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn549 github.com/goccy/pythonwasm2go/p2.Fn549
func Fn549(m *base.Module, l0 int32) int32

//go:linkname Fn550 github.com/goccy/pythonwasm2go/p2.Fn550
func Fn550(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn551 github.com/goccy/pythonwasm2go/p2.Fn551
func Fn551(m *base.Module, l0 int32) int32

//go:linkname Fn552 github.com/goccy/pythonwasm2go/p2.Fn552
func Fn552(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn553 github.com/goccy/pythonwasm2go/p2.Fn553
func Fn553(m *base.Module, l0 int32) int32

//go:linkname Fn554 github.com/goccy/pythonwasm2go/p2.Fn554
func Fn554(m *base.Module, l0 int32) int32

//go:linkname Fn555 github.com/goccy/pythonwasm2go/p2.Fn555
func Fn555(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn558 github.com/goccy/pythonwasm2go/p2.Fn558
func Fn558(m *base.Module, l0 int32)

//go:linkname Fn562 github.com/goccy/pythonwasm2go/p2.Fn562
func Fn562(m *base.Module, l0 int32) int32

//go:linkname Fn564 github.com/goccy/pythonwasm2go/p2.Fn564
func Fn564(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn565 github.com/goccy/pythonwasm2go/p2.Fn565
func Fn565(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn567 github.com/goccy/pythonwasm2go/p0.Fn567
func Fn567(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn568 github.com/goccy/pythonwasm2go/p0.Fn568
func Fn568(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn569 github.com/goccy/pythonwasm2go/p2.Fn569
func Fn569(m *base.Module, l0 int32) int32

//go:linkname Fn570 github.com/goccy/pythonwasm2go/p0.Fn570
func Fn570(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn571 github.com/goccy/pythonwasm2go/p0.Fn571
func Fn571(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn572 github.com/goccy/pythonwasm2go/p0.Fn572
func Fn572(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn573 github.com/goccy/pythonwasm2go/p0.Fn573
func Fn573(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn574 github.com/goccy/pythonwasm2go/p0.Fn574
func Fn574(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn575 github.com/goccy/pythonwasm2go/p0.Fn575
func Fn575(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn576 github.com/goccy/pythonwasm2go/p0.Fn576
func Fn576(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn577 github.com/goccy/pythonwasm2go/p0.Fn577
func Fn577(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn578 github.com/goccy/pythonwasm2go/p0.Fn578
func Fn578(m *base.Module, l0 int32) int32

//go:linkname Fn580 github.com/goccy/pythonwasm2go/p0.Fn580
func Fn580(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn606 github.com/goccy/pythonwasm2go/p0.Fn606
func Fn606(m *base.Module, l0 int32) int32

//go:linkname Fn628 github.com/goccy/pythonwasm2go/p2.Fn628
func Fn628(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn652 github.com/goccy/pythonwasm2go/p2.Fn652
func Fn652(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn685 github.com/goccy/pythonwasm2go/p2.Fn685
func Fn685(m *base.Module, l0 int32)

//go:linkname Fn686 github.com/goccy/pythonwasm2go/p2.Fn686
func Fn686(m *base.Module, l0 int32) int32

//go:linkname Fn688 github.com/goccy/pythonwasm2go/p2.Fn688
func Fn688(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn689 github.com/goccy/pythonwasm2go/p2.Fn689
func Fn689(m *base.Module, l0 int32) int32

//go:linkname Fn692 github.com/goccy/pythonwasm2go/p2.Fn692
func Fn692(m *base.Module, l0 int32)

//go:linkname Fn698 github.com/goccy/pythonwasm2go/p2.Fn698
func Fn698(m *base.Module, l0 int32) int32

//go:linkname Fn700 github.com/goccy/pythonwasm2go/p2.Fn700
func Fn700(m *base.Module, l0 int32) int32

//go:linkname Fn714 github.com/goccy/pythonwasm2go/p2.Fn714
func Fn714(m *base.Module, l0 int32) int32

//go:linkname Fn720 github.com/goccy/pythonwasm2go/p2.Fn720
func Fn720(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn721 github.com/goccy/pythonwasm2go/p2.Fn721
func Fn721(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn751 github.com/goccy/pythonwasm2go/p2.Fn751
func Fn751(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn756 github.com/goccy/pythonwasm2go/p0.Fn756
func Fn756(m *base.Module, l0 float64) int32

//go:linkname Fn762 github.com/goccy/pythonwasm2go/p0.Fn762
func Fn762(m *base.Module, l0 int32) float64

//go:linkname Fn763 github.com/goccy/pythonwasm2go/p2.Fn763
func Fn763(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn765 github.com/goccy/pythonwasm2go/p2.Fn765
func Fn765(m *base.Module, l0 int32) int32

//go:linkname Fn774 github.com/goccy/pythonwasm2go/p2.Fn774
func Fn774(m *base.Module, l0 int32, l1 int32) float64

//go:linkname Fn808 github.com/goccy/pythonwasm2go/p2.Fn808
func Fn808(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn820 github.com/goccy/pythonwasm2go/p0.Fn820
func Fn820(m *base.Module, l0 int32) int32

//go:linkname Fn821 github.com/goccy/pythonwasm2go/p0.Fn821
func Fn821(m *base.Module, l0 int32) int32

//go:linkname Fn822 github.com/goccy/pythonwasm2go/p0.Fn822
func Fn822(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn823 github.com/goccy/pythonwasm2go/p0.Fn823
func Fn823(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn824 github.com/goccy/pythonwasm2go/p2.Fn824
func Fn824(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn825 github.com/goccy/pythonwasm2go/p2.Fn825
func Fn825(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn826 github.com/goccy/pythonwasm2go/p2.Fn826
func Fn826(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn828 github.com/goccy/pythonwasm2go/p0.Fn828
func Fn828(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn830 github.com/goccy/pythonwasm2go/p0.Fn830
func Fn830(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn833 github.com/goccy/pythonwasm2go/p0.Fn833
func Fn833(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn834 github.com/goccy/pythonwasm2go/p0.Fn834
func Fn834(m *base.Module, l0 int32) int32

//go:linkname Fn835 github.com/goccy/pythonwasm2go/p0.Fn835
func Fn835(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn836 github.com/goccy/pythonwasm2go/p0.Fn836
func Fn836(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn840 github.com/goccy/pythonwasm2go/p2.Fn840
func Fn840(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn845 github.com/goccy/pythonwasm2go/p2.Fn845
func Fn845(m *base.Module, l0 int32)

//go:linkname Fn846 github.com/goccy/pythonwasm2go/p2.Fn846
func Fn846(m *base.Module, l0 int32) int32

//go:linkname Fn858 github.com/goccy/pythonwasm2go/p2.Fn858
func Fn858(m *base.Module, l0 int32) int32

//go:linkname Fn859 github.com/goccy/pythonwasm2go/p0.Fn859
func Fn859(m *base.Module, l0 int32) int32

//go:linkname Fn875 github.com/goccy/pythonwasm2go/p2.Fn875
func Fn875(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn883 github.com/goccy/pythonwasm2go/p2.Fn883
func Fn883(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn889 github.com/goccy/pythonwasm2go/p2.Fn889
func Fn889(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn909 github.com/goccy/pythonwasm2go/p0.Fn909
func Fn909(m *base.Module, l0 int32) int32

//go:linkname Fn911 github.com/goccy/pythonwasm2go/p2.Fn911
func Fn911(m *base.Module, l0 int32) int32

//go:linkname Fn912 github.com/goccy/pythonwasm2go/p0.Fn912
func Fn912(m *base.Module, l0 int32) int32

//go:linkname Fn914 github.com/goccy/pythonwasm2go/p0.Fn914
func Fn914(m *base.Module, l0 int32) int32

//go:linkname Fn915 github.com/goccy/pythonwasm2go/p0.Fn915
func Fn915(m *base.Module, l0 int64) int32

//go:linkname Fn917 github.com/goccy/pythonwasm2go/p0.Fn917
func Fn917(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn918 github.com/goccy/pythonwasm2go/p0.Fn918
func Fn918(m *base.Module, l0 int32) int32

//go:linkname Fn919 github.com/goccy/pythonwasm2go/p0.Fn919
func Fn919(m *base.Module, l0 int32) int32

//go:linkname Fn920 github.com/goccy/pythonwasm2go/p0.Fn920
func Fn920(m *base.Module, l0 int32) int32

//go:linkname Fn921 github.com/goccy/pythonwasm2go/p2.Fn921
func Fn921(m *base.Module, l0 int32) int32

//go:linkname Fn922 github.com/goccy/pythonwasm2go/p2.Fn922
func Fn922(m *base.Module, l0 int32) int32

//go:linkname Fn923 github.com/goccy/pythonwasm2go/p0.Fn923
func Fn923(m *base.Module, l0 int32) int32

//go:linkname Fn926 github.com/goccy/pythonwasm2go/p2.Fn926
func Fn926(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn927 github.com/goccy/pythonwasm2go/p2.Fn927
func Fn927(m *base.Module, l0 int32) int64

//go:linkname Fn929 github.com/goccy/pythonwasm2go/p2.Fn929
func Fn929(m *base.Module, l0 int32) int32

//go:linkname Fn930 github.com/goccy/pythonwasm2go/p2.Fn930
func Fn930(m *base.Module, l0 int32) int32

//go:linkname Fn931 github.com/goccy/pythonwasm2go/p2.Fn931
func Fn931(m *base.Module, l0 int32)

//go:linkname Fn932 github.com/goccy/pythonwasm2go/p0.Fn932
func Fn932(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn934 github.com/goccy/pythonwasm2go/p2.Fn934
func Fn934(m *base.Module, l0 int32) int32

//go:linkname Fn935 github.com/goccy/pythonwasm2go/p0.Fn935
func Fn935(m *base.Module, l0 int64) int32

//go:linkname Fn936 github.com/goccy/pythonwasm2go/p0.Fn936
func Fn936(m *base.Module, l0 int32) int64

//go:linkname Fn937 github.com/goccy/pythonwasm2go/p2.Fn937
func Fn937(m *base.Module, l0 int32) int64

//go:linkname Fn938 github.com/goccy/pythonwasm2go/p0.Fn938
func Fn938(m *base.Module, l0 int32) int64

//go:linkname Fn941 github.com/goccy/pythonwasm2go/p2.Fn941
func Fn941(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn942 github.com/goccy/pythonwasm2go/p2.Fn942
func Fn942(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn943 github.com/goccy/pythonwasm2go/p2.Fn943
func Fn943(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn944 github.com/goccy/pythonwasm2go/p2.Fn944
func Fn944(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn945 github.com/goccy/pythonwasm2go/p2.Fn945
func Fn945(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn948 github.com/goccy/pythonwasm2go/p2.Fn948
func Fn948(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn949 github.com/goccy/pythonwasm2go/p0.Fn949
func Fn949(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn950 github.com/goccy/pythonwasm2go/p2.Fn950
func Fn950(m *base.Module, l0 int32)

//go:linkname Fn951 github.com/goccy/pythonwasm2go/p2.Fn951
func Fn951(m *base.Module, l0 int64) int32

//go:linkname Fn952 github.com/goccy/pythonwasm2go/p2.Fn952
func Fn952(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn953 github.com/goccy/pythonwasm2go/p2.Fn953
func Fn953(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn955 github.com/goccy/pythonwasm2go/p2.Fn955
func Fn955(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn956 github.com/goccy/pythonwasm2go/p2.Fn956
func Fn956(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn957 github.com/goccy/pythonwasm2go/p0.Fn957
func Fn957(m *base.Module, l0 int32) float64

//go:linkname Fn960 github.com/goccy/pythonwasm2go/p2.Fn960
func Fn960(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn963 github.com/goccy/pythonwasm2go/p2.Fn963
func Fn963(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn967 github.com/goccy/pythonwasm2go/p2.Fn967
func Fn967(m *base.Module, l0 int32) int32

//go:linkname Fn968 github.com/goccy/pythonwasm2go/p2.Fn968
func Fn968(m *base.Module, l0 int32, l1 int64) int32

//go:linkname Fn971 github.com/goccy/pythonwasm2go/p2.Fn971
func Fn971(m *base.Module, l0 int32) int32

//go:linkname Fn972 github.com/goccy/pythonwasm2go/p2.Fn972
func Fn972(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn974 github.com/goccy/pythonwasm2go/p2.Fn974
func Fn974(m *base.Module, l0 int32) int32

//go:linkname Fn975 github.com/goccy/pythonwasm2go/p2.Fn975
func Fn975(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn979 github.com/goccy/pythonwasm2go/p2.Fn979
func Fn979(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn988 github.com/goccy/pythonwasm2go/p2.Fn988
func Fn988(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn989 github.com/goccy/pythonwasm2go/p2.Fn989
func Fn989(m *base.Module, l0 int32)

//go:linkname Fn990 github.com/goccy/pythonwasm2go/p2.Fn990
func Fn990(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn991 github.com/goccy/pythonwasm2go/p2.Fn991
func Fn991(m *base.Module, l0 int32) int32

//go:linkname Fn992 github.com/goccy/pythonwasm2go/p2.Fn992
func Fn992(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn993 github.com/goccy/pythonwasm2go/p2.Fn993
func Fn993(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn994 github.com/goccy/pythonwasm2go/p2.Fn994
func Fn994(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1001 github.com/goccy/pythonwasm2go/p0.Fn1001
func Fn1001(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1014 github.com/goccy/pythonwasm2go/p2.Fn1014
func Fn1014(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1016 github.com/goccy/pythonwasm2go/p2.Fn1016
func Fn1016(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1035 github.com/goccy/pythonwasm2go/p2.Fn1035
func Fn1035(m *base.Module, l0 int32) int32

//go:linkname Fn1046 github.com/goccy/pythonwasm2go/p2.Fn1046
func Fn1046(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1058 github.com/goccy/pythonwasm2go/p0.Fn1058
func Fn1058(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1078 github.com/goccy/pythonwasm2go/p2.Fn1078
func Fn1078(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1102 github.com/goccy/pythonwasm2go/p0.Fn1102
func Fn1102(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1116 github.com/goccy/pythonwasm2go/p0.Fn1116
func Fn1116(m *base.Module) int32

//go:linkname Fn1119 github.com/goccy/pythonwasm2go/p2.Fn1119
func Fn1119(m *base.Module, l0 int32)

//go:linkname Fn1125 github.com/goccy/pythonwasm2go/p0.Fn1125
func Fn1125(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1130 github.com/goccy/pythonwasm2go/p2.Fn1130
func Fn1130(m *base.Module, l0 int32)

//go:linkname Fn1131 github.com/goccy/pythonwasm2go/p2.Fn1131
func Fn1131(m *base.Module, l0 int32)

//go:linkname Fn1133 github.com/goccy/pythonwasm2go/p2.Fn1133
func Fn1133(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1138 github.com/goccy/pythonwasm2go/p0.Fn1138
func Fn1138(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1141 github.com/goccy/pythonwasm2go/p0.Fn1141
func Fn1141(m *base.Module, l0 int32) int32

//go:linkname Fn1143 github.com/goccy/pythonwasm2go/p0.Fn1143
func Fn1143(m *base.Module, l0 int32)

//go:linkname Fn1144 github.com/goccy/pythonwasm2go/p2.Fn1144
func Fn1144(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1145 github.com/goccy/pythonwasm2go/p0.Fn1145
func Fn1145(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1146 github.com/goccy/pythonwasm2go/p0.Fn1146
func Fn1146(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1147 github.com/goccy/pythonwasm2go/p0.Fn1147
func Fn1147(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1148 github.com/goccy/pythonwasm2go/p2.Fn1148
func Fn1148(m *base.Module, l0 int32) int32

//go:linkname Fn1151 github.com/goccy/pythonwasm2go/p0.Fn1151
func Fn1151(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1152 github.com/goccy/pythonwasm2go/p0.Fn1152
func Fn1152(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1153 github.com/goccy/pythonwasm2go/p2.Fn1153
func Fn1153(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1154 github.com/goccy/pythonwasm2go/p0.Fn1154
func Fn1154(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1155 github.com/goccy/pythonwasm2go/p2.Fn1155
func Fn1155(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1156 github.com/goccy/pythonwasm2go/p2.Fn1156
func Fn1156(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1159 github.com/goccy/pythonwasm2go/p2.Fn1159
func Fn1159(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1160 github.com/goccy/pythonwasm2go/p0.Fn1160
func Fn1160(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1161 github.com/goccy/pythonwasm2go/p0.Fn1161
func Fn1161(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1164 github.com/goccy/pythonwasm2go/p2.Fn1164
func Fn1164(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1165 github.com/goccy/pythonwasm2go/p2.Fn1165
func Fn1165(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1167 github.com/goccy/pythonwasm2go/p0.Fn1167
func Fn1167(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1168 github.com/goccy/pythonwasm2go/p0.Fn1168
func Fn1168(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1169 github.com/goccy/pythonwasm2go/p2.Fn1169
func Fn1169(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1171 github.com/goccy/pythonwasm2go/p0.Fn1171
func Fn1171(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1172 github.com/goccy/pythonwasm2go/p0.Fn1172
func Fn1172(m *base.Module, l0 int32) int32

//go:linkname Fn1173 github.com/goccy/pythonwasm2go/p2.Fn1173
func Fn1173(m *base.Module, l0 int32) int32

//go:linkname Fn1176 github.com/goccy/pythonwasm2go/p0.Fn1176
func Fn1176(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1180 github.com/goccy/pythonwasm2go/p0.Fn1180
func Fn1180(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1181 github.com/goccy/pythonwasm2go/p0.Fn1181
func Fn1181(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1182 github.com/goccy/pythonwasm2go/p0.Fn1182
func Fn1182(m *base.Module, l0 int32) int32

//go:linkname Fn1183 github.com/goccy/pythonwasm2go/p0.Fn1183
func Fn1183(m *base.Module, l0 int32) int32

//go:linkname Fn1185 github.com/goccy/pythonwasm2go/p0.Fn1185
func Fn1185(m *base.Module, l0 int32) int32

//go:linkname Fn1186 github.com/goccy/pythonwasm2go/p0.Fn1186
func Fn1186(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1189 github.com/goccy/pythonwasm2go/p2.Fn1189
func Fn1189(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1201 github.com/goccy/pythonwasm2go/p2.Fn1201
func Fn1201(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1202 github.com/goccy/pythonwasm2go/p0.Fn1202
func Fn1202(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1203 github.com/goccy/pythonwasm2go/p0.Fn1203
func Fn1203(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1210 github.com/goccy/pythonwasm2go/p2.Fn1210
func Fn1210(m *base.Module, l0 int32)

//go:linkname Fn1220 github.com/goccy/pythonwasm2go/p2.Fn1220
func Fn1220(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1226 github.com/goccy/pythonwasm2go/p0.Fn1226
func Fn1226(m *base.Module, l0 int32) int32

//go:linkname Fn1227 github.com/goccy/pythonwasm2go/p0.Fn1227
func Fn1227(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1229 github.com/goccy/pythonwasm2go/p2.Fn1229
func Fn1229(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1230 github.com/goccy/pythonwasm2go/p2.Fn1230
func Fn1230(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1231 github.com/goccy/pythonwasm2go/p2.Fn1231
func Fn1231(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1234 github.com/goccy/pythonwasm2go/p2.Fn1234
func Fn1234(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1236 github.com/goccy/pythonwasm2go/p2.Fn1236
func Fn1236(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1237 github.com/goccy/pythonwasm2go/p2.Fn1237
func Fn1237(m *base.Module, l0 int32)

//go:linkname Fn1238 github.com/goccy/pythonwasm2go/p2.Fn1238
func Fn1238(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1239 github.com/goccy/pythonwasm2go/p2.Fn1239
func Fn1239(m *base.Module, l0 int32)

//go:linkname Fn1271 github.com/goccy/pythonwasm2go/p2.Fn1271
func Fn1271(m *base.Module, l0 int32) int32

//go:linkname Fn1272 github.com/goccy/pythonwasm2go/p0.Fn1272
func Fn1272(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1284 github.com/goccy/pythonwasm2go/p2.Fn1284
func Fn1284(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1286 github.com/goccy/pythonwasm2go/p0.Fn1286
func Fn1286(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1291 github.com/goccy/pythonwasm2go/p2.Fn1291
func Fn1291(m *base.Module, l0 int32)

//go:linkname Fn1292 github.com/goccy/pythonwasm2go/p0.Fn1292
func Fn1292(m *base.Module, l0 int32) int32

//go:linkname Fn1293 github.com/goccy/pythonwasm2go/p2.Fn1293
func Fn1293(m *base.Module, l0 int32) int32

//go:linkname Fn1295 github.com/goccy/pythonwasm2go/p2.Fn1295
func Fn1295(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1296 github.com/goccy/pythonwasm2go/p2.Fn1296
func Fn1296(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1297 github.com/goccy/pythonwasm2go/p2.Fn1297
func Fn1297(m *base.Module, l0 int32)

//go:linkname Fn1298 github.com/goccy/pythonwasm2go/p2.Fn1298
func Fn1298(m *base.Module, l0 int32)

//go:linkname Fn1299 github.com/goccy/pythonwasm2go/p2.Fn1299
func Fn1299(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1300 github.com/goccy/pythonwasm2go/p2.Fn1300
func Fn1300(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1301 github.com/goccy/pythonwasm2go/p2.Fn1301
func Fn1301(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1308 github.com/goccy/pythonwasm2go/p2.Fn1308
func Fn1308(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1311 github.com/goccy/pythonwasm2go/p2.Fn1311
func Fn1311(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1312 github.com/goccy/pythonwasm2go/p2.Fn1312
func Fn1312(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1313 github.com/goccy/pythonwasm2go/p2.Fn1313
func Fn1313(m *base.Module) int32

//go:linkname Fn1315 github.com/goccy/pythonwasm2go/p2.Fn1315
func Fn1315(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32

//go:linkname Fn1316 github.com/goccy/pythonwasm2go/p2.Fn1316
func Fn1316(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32

//go:linkname Fn1317 github.com/goccy/pythonwasm2go/p2.Fn1317
func Fn1317(m *base.Module, l0 int32)

//go:linkname Fn1319 github.com/goccy/pythonwasm2go/p2.Fn1319
func Fn1319(m *base.Module, l0 int32) int32

//go:linkname Fn1322 github.com/goccy/pythonwasm2go/p2.Fn1322
func Fn1322(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1323 github.com/goccy/pythonwasm2go/p2.Fn1323
func Fn1323(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1325 github.com/goccy/pythonwasm2go/p2.Fn1325
func Fn1325(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)

//go:linkname Fn1326 github.com/goccy/pythonwasm2go/p2.Fn1326
func Fn1326(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1327 github.com/goccy/pythonwasm2go/p2.Fn1327
func Fn1327(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1330 github.com/goccy/pythonwasm2go/p2.Fn1330
func Fn1330(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1331 github.com/goccy/pythonwasm2go/p2.Fn1331
func Fn1331(m *base.Module, l0 int32) int32

//go:linkname Fn1332 github.com/goccy/pythonwasm2go/p2.Fn1332
func Fn1332(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1333 github.com/goccy/pythonwasm2go/p2.Fn1333
func Fn1333(m *base.Module, l0 int32) int32

//go:linkname Fn1337 github.com/goccy/pythonwasm2go/p2.Fn1337
func Fn1337(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1363 github.com/goccy/pythonwasm2go/p0.Fn1363
func Fn1363(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1383 github.com/goccy/pythonwasm2go/p2.Fn1383
func Fn1383(m *base.Module, l0 int32) int32

//go:linkname Fn1384 github.com/goccy/pythonwasm2go/p2.Fn1384
func Fn1384(m *base.Module, l0 int32) int32

//go:linkname Fn1385 github.com/goccy/pythonwasm2go/p2.Fn1385
func Fn1385(m *base.Module, l0 int32) int32

//go:linkname Fn1389 github.com/goccy/pythonwasm2go/p2.Fn1389
func Fn1389(m *base.Module, l0 int32) int32

//go:linkname Fn1390 github.com/goccy/pythonwasm2go/p2.Fn1390
func Fn1390(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1392 github.com/goccy/pythonwasm2go/p2.Fn1392
func Fn1392(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1393 github.com/goccy/pythonwasm2go/p2.Fn1393
func Fn1393(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1394 github.com/goccy/pythonwasm2go/p2.Fn1394
func Fn1394(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1395 github.com/goccy/pythonwasm2go/p0.Fn1395
func Fn1395(m *base.Module, l0 int32) int32

//go:linkname Fn1396 github.com/goccy/pythonwasm2go/p2.Fn1396
func Fn1396(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1397 github.com/goccy/pythonwasm2go/p2.Fn1397
func Fn1397(m *base.Module, l0 int32) int32

//go:linkname Fn1398 github.com/goccy/pythonwasm2go/p2.Fn1398
func Fn1398(m *base.Module, l0 int32) int32

//go:linkname Fn1399 github.com/goccy/pythonwasm2go/p2.Fn1399
func Fn1399(m *base.Module, l0 int32) int32

//go:linkname Fn1400 github.com/goccy/pythonwasm2go/p2.Fn1400
func Fn1400(m *base.Module, l0 int32) int32

//go:linkname Fn1401 github.com/goccy/pythonwasm2go/p0.Fn1401
func Fn1401(m *base.Module, l0 int32) int32

//go:linkname Fn1402 github.com/goccy/pythonwasm2go/p2.Fn1402
func Fn1402(m *base.Module, l0 int32) int32

//go:linkname Fn1404 github.com/goccy/pythonwasm2go/p2.Fn1404
func Fn1404(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1405 github.com/goccy/pythonwasm2go/p0.Fn1405
func Fn1405(m *base.Module, l0 int32) int32

//go:linkname Fn1418 github.com/goccy/pythonwasm2go/p2.Fn1418
func Fn1418(m *base.Module, l0 int32) int32

//go:linkname Fn1428 github.com/goccy/pythonwasm2go/p2.Fn1428
func Fn1428(m *base.Module, l0 int32) int32

//go:linkname Fn1443 github.com/goccy/pythonwasm2go/p2.Fn1443
func Fn1443(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn1445 github.com/goccy/pythonwasm2go/p2.Fn1445
func Fn1445(m *base.Module, l0 int32) int32

//go:linkname Fn1447 github.com/goccy/pythonwasm2go/p2.Fn1447
func Fn1447(m *base.Module, l0 int32) int32

//go:linkname Fn1448 github.com/goccy/pythonwasm2go/p2.Fn1448
func Fn1448(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1450 github.com/goccy/pythonwasm2go/p2.Fn1450
func Fn1450(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1454 github.com/goccy/pythonwasm2go/p2.Fn1454
func Fn1454(m *base.Module, l0 int32) int32

//go:linkname Fn1456 github.com/goccy/pythonwasm2go/p2.Fn1456
func Fn1456(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1457 github.com/goccy/pythonwasm2go/p2.Fn1457
func Fn1457(m *base.Module, l0 int32) int32

//go:linkname Fn1458 github.com/goccy/pythonwasm2go/p2.Fn1458
func Fn1458(m *base.Module, l0 int32) int32

//go:linkname Fn1463 github.com/goccy/pythonwasm2go/p2.Fn1463
func Fn1463(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1465 github.com/goccy/pythonwasm2go/p2.Fn1465
func Fn1465(m *base.Module, l0 int32) int32

//go:linkname Fn1469 github.com/goccy/pythonwasm2go/p2.Fn1469
func Fn1469(m *base.Module, l0 int32) int32

//go:linkname Fn1470 github.com/goccy/pythonwasm2go/p2.Fn1470
func Fn1470(m *base.Module, l0 int32) int32

//go:linkname Fn1473 github.com/goccy/pythonwasm2go/p2.Fn1473
func Fn1473(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1475 github.com/goccy/pythonwasm2go/p2.Fn1475
func Fn1475(m *base.Module, l0 int32) int32

//go:linkname Fn1483 github.com/goccy/pythonwasm2go/p2.Fn1483
func Fn1483(m *base.Module, l0 int32)

//go:linkname Fn1491 github.com/goccy/pythonwasm2go/p2.Fn1491
func Fn1491(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1498 github.com/goccy/pythonwasm2go/p2.Fn1498
func Fn1498(m *base.Module, l0 int32) int32

//go:linkname Fn1510 github.com/goccy/pythonwasm2go/p2.Fn1510
func Fn1510(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1513 github.com/goccy/pythonwasm2go/p2.Fn1513
func Fn1513(m *base.Module, l0 int32) int32

//go:linkname Fn1514 github.com/goccy/pythonwasm2go/p2.Fn1514
func Fn1514(m *base.Module, l0 int32) int32

//go:linkname Fn1516 github.com/goccy/pythonwasm2go/p2.Fn1516
func Fn1516(m *base.Module, l0 int32)

//go:linkname Fn1517 github.com/goccy/pythonwasm2go/p2.Fn1517
func Fn1517(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1530 github.com/goccy/pythonwasm2go/p2.Fn1530
func Fn1530(m *base.Module, l0 int32) int32

//go:linkname Fn1563 github.com/goccy/pythonwasm2go/p2.Fn1563
func Fn1563(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1564 github.com/goccy/pythonwasm2go/p2.Fn1564
func Fn1564(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1566 github.com/goccy/pythonwasm2go/p2.Fn1566
func Fn1566(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1567 github.com/goccy/pythonwasm2go/p2.Fn1567
func Fn1567(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1568 github.com/goccy/pythonwasm2go/p2.Fn1568
func Fn1568(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1569 github.com/goccy/pythonwasm2go/p2.Fn1569
func Fn1569(m *base.Module, l0 int32)

//go:linkname Fn1578 github.com/goccy/pythonwasm2go/p2.Fn1578
func Fn1578(m *base.Module) int32

//go:linkname Fn1579 github.com/goccy/pythonwasm2go/p2.Fn1579
func Fn1579(m *base.Module) int32

//go:linkname Fn1582 github.com/goccy/pythonwasm2go/p2.Fn1582
func Fn1582(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1584 github.com/goccy/pythonwasm2go/p2.Fn1584
func Fn1584(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1587 github.com/goccy/pythonwasm2go/p2.Fn1587
func Fn1587(m *base.Module, l0 int32) int32

//go:linkname Fn1590 github.com/goccy/pythonwasm2go/p2.Fn1590
func Fn1590(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1625 github.com/goccy/pythonwasm2go/p2.Fn1625
func Fn1625(m *base.Module, l0 int32)

//go:linkname Fn1629 github.com/goccy/pythonwasm2go/p2.Fn1629
func Fn1629(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1651 github.com/goccy/pythonwasm2go/p0.Fn1651
func Fn1651(m *base.Module, l0 int32) int32

//go:linkname Fn1653 github.com/goccy/pythonwasm2go/p2.Fn1653
func Fn1653(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1654 github.com/goccy/pythonwasm2go/p2.Fn1654
func Fn1654(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1655 github.com/goccy/pythonwasm2go/p2.Fn1655
func Fn1655(m *base.Module, l0 int32) int32

//go:linkname Fn1656 github.com/goccy/pythonwasm2go/p0.Fn1656
func Fn1656(m *base.Module, l0 int32) int32

//go:linkname Fn1698 github.com/goccy/pythonwasm2go/p2.Fn1698
func Fn1698(m *base.Module, l0 int32) int32

//go:linkname Fn1703 github.com/goccy/pythonwasm2go/p2.Fn1703
func Fn1703(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1707 github.com/goccy/pythonwasm2go/p2.Fn1707
func Fn1707(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1711 github.com/goccy/pythonwasm2go/p2.Fn1711
func Fn1711(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1725 github.com/goccy/pythonwasm2go/p2.Fn1725
func Fn1725(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1734 github.com/goccy/pythonwasm2go/p2.Fn1734
func Fn1734(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1735 github.com/goccy/pythonwasm2go/p2.Fn1735
func Fn1735(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1740 github.com/goccy/pythonwasm2go/p0.Fn1740
func Fn1740(m *base.Module, l0 int32) int32

//go:linkname Fn1741 github.com/goccy/pythonwasm2go/p0.Fn1741
func Fn1741(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1742 github.com/goccy/pythonwasm2go/p0.Fn1742
func Fn1742(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1743 github.com/goccy/pythonwasm2go/p0.Fn1743
func Fn1743(m *base.Module, l0 int32) int32

//go:linkname Fn1744 github.com/goccy/pythonwasm2go/p2.Fn1744
func Fn1744(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1747 github.com/goccy/pythonwasm2go/p0.Fn1747
func Fn1747(m *base.Module, l0 int32) int32

//go:linkname Fn1763 github.com/goccy/pythonwasm2go/p2.Fn1763
func Fn1763(m *base.Module, l0 int32) int32

//go:linkname Fn1766 github.com/goccy/pythonwasm2go/p2.Fn1766
func Fn1766(m *base.Module, l0 int32) int32

//go:linkname Fn1769 github.com/goccy/pythonwasm2go/p2.Fn1769
func Fn1769(m *base.Module, l0 int32) int32

//go:linkname Fn1774 github.com/goccy/pythonwasm2go/p2.Fn1774
func Fn1774(m *base.Module, l0 int32)

//go:linkname Fn1775 github.com/goccy/pythonwasm2go/p2.Fn1775
func Fn1775(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1776 github.com/goccy/pythonwasm2go/p2.Fn1776
func Fn1776(m *base.Module, l0 int32)

//go:linkname Fn1777 github.com/goccy/pythonwasm2go/p0.Fn1777
func Fn1777(m *base.Module, l0 int32) int32

//go:linkname Fn1778 github.com/goccy/pythonwasm2go/p2.Fn1778
func Fn1778(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1780 github.com/goccy/pythonwasm2go/p2.Fn1780
func Fn1780(m *base.Module, l0 int32) int32

//go:linkname Fn1782 github.com/goccy/pythonwasm2go/p0.Fn1782
func Fn1782(m *base.Module, l0 int32) int32

//go:linkname Fn1783 github.com/goccy/pythonwasm2go/p0.Fn1783
func Fn1783(m *base.Module, l0 int32) int32

//go:linkname Fn1784 github.com/goccy/pythonwasm2go/p0.Fn1784
func Fn1784(m *base.Module, l0 int32) int32

//go:linkname Fn1788 github.com/goccy/pythonwasm2go/p2.Fn1788
func Fn1788(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1790 github.com/goccy/pythonwasm2go/p0.Fn1790
func Fn1790(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1791 github.com/goccy/pythonwasm2go/p0.Fn1791
func Fn1791(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1792 github.com/goccy/pythonwasm2go/p2.Fn1792
func Fn1792(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1793 github.com/goccy/pythonwasm2go/p0.Fn1793
func Fn1793(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1794 github.com/goccy/pythonwasm2go/p0.Fn1794
func Fn1794(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1795 github.com/goccy/pythonwasm2go/p2.Fn1795
func Fn1795(m *base.Module, l0 int32) int32

//go:linkname Fn1797 github.com/goccy/pythonwasm2go/p0.Fn1797
func Fn1797(m *base.Module, l0 int32) int32

//go:linkname Fn1798 github.com/goccy/pythonwasm2go/p0.Fn1798
func Fn1798(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1799 github.com/goccy/pythonwasm2go/p0.Fn1799
func Fn1799(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1800 github.com/goccy/pythonwasm2go/p0.Fn1800
func Fn1800(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1801 github.com/goccy/pythonwasm2go/p2.Fn1801
func Fn1801(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1802 github.com/goccy/pythonwasm2go/p0.Fn1802
func Fn1802(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1803 github.com/goccy/pythonwasm2go/p2.Fn1803
func Fn1803(m *base.Module, l0 int32) int32

//go:linkname Fn1804 github.com/goccy/pythonwasm2go/p2.Fn1804
func Fn1804(m *base.Module, l0 int32) int32

//go:linkname Fn1805 github.com/goccy/pythonwasm2go/p2.Fn1805
func Fn1805(m *base.Module, l0 int32)

//go:linkname Fn1806 github.com/goccy/pythonwasm2go/p0.Fn1806
func Fn1806(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1814 github.com/goccy/pythonwasm2go/p2.Fn1814
func Fn1814(m *base.Module, l0 int32) int32

//go:linkname Fn1815 github.com/goccy/pythonwasm2go/p2.Fn1815
func Fn1815(m *base.Module, l0 int32) int32

//go:linkname Fn1817 github.com/goccy/pythonwasm2go/p2.Fn1817
func Fn1817(m *base.Module, l0 int32)

//go:linkname Fn1823 github.com/goccy/pythonwasm2go/p2.Fn1823
func Fn1823(m *base.Module, l0 int32)

//go:linkname Fn1824 github.com/goccy/pythonwasm2go/p2.Fn1824
func Fn1824(m *base.Module, l0 int32) int32

//go:linkname Fn1825 github.com/goccy/pythonwasm2go/p2.Fn1825
func Fn1825(m *base.Module, l0 int32)

//go:linkname Fn1826 github.com/goccy/pythonwasm2go/p2.Fn1826
func Fn1826(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1827 github.com/goccy/pythonwasm2go/p2.Fn1827
func Fn1827(m *base.Module, l0 int32) int32

//go:linkname Fn1828 github.com/goccy/pythonwasm2go/p2.Fn1828
func Fn1828(m *base.Module, l0 int32) int32

//go:linkname Fn1835 github.com/goccy/pythonwasm2go/p2.Fn1835
func Fn1835(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1836 github.com/goccy/pythonwasm2go/p2.Fn1836
func Fn1836(m *base.Module, l0 int32) int32

//go:linkname Fn1839 github.com/goccy/pythonwasm2go/p2.Fn1839
func Fn1839(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1840 github.com/goccy/pythonwasm2go/p2.Fn1840
func Fn1840(m *base.Module, l0 int32) int32

//go:linkname Fn1844 github.com/goccy/pythonwasm2go/p2.Fn1844
func Fn1844(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1847 github.com/goccy/pythonwasm2go/p2.Fn1847
func Fn1847(m *base.Module, l0 int32)

//go:linkname Fn1849 github.com/goccy/pythonwasm2go/p2.Fn1849
func Fn1849(m *base.Module, l0 int32) int32

//go:linkname Fn1853 github.com/goccy/pythonwasm2go/p2.Fn1853
func Fn1853(m *base.Module, l0 int32) int32

//go:linkname Fn1854 github.com/goccy/pythonwasm2go/p2.Fn1854
func Fn1854(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1855 github.com/goccy/pythonwasm2go/p2.Fn1855
func Fn1855(m *base.Module) int32

//go:linkname Fn1856 github.com/goccy/pythonwasm2go/p2.Fn1856
func Fn1856(m *base.Module) int32

//go:linkname Fn1862 github.com/goccy/pythonwasm2go/p2.Fn1862
func Fn1862(m *base.Module, l0 int32) int32

//go:linkname Fn1864 github.com/goccy/pythonwasm2go/p2.Fn1864
func Fn1864(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1865 github.com/goccy/pythonwasm2go/p2.Fn1865
func Fn1865(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1866 github.com/goccy/pythonwasm2go/p2.Fn1866
func Fn1866(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1867 github.com/goccy/pythonwasm2go/p2.Fn1867
func Fn1867(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1868 github.com/goccy/pythonwasm2go/p2.Fn1868
func Fn1868(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1869 github.com/goccy/pythonwasm2go/p2.Fn1869
func Fn1869(m *base.Module) int32

//go:linkname Fn1870 github.com/goccy/pythonwasm2go/p2.Fn1870
func Fn1870(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1871 github.com/goccy/pythonwasm2go/p2.Fn1871
func Fn1871(m *base.Module) int64

//go:linkname Fn1872 github.com/goccy/pythonwasm2go/p2.Fn1872
func Fn1872(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1874 github.com/goccy/pythonwasm2go/p2.Fn1874
func Fn1874(m *base.Module, l0 int32, l1 int64)

//go:linkname Fn1875 github.com/goccy/pythonwasm2go/p2.Fn1875
func Fn1875(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1877 github.com/goccy/pythonwasm2go/p2.Fn1877
func Fn1877(m *base.Module, l0 int32)

//go:linkname Fn1878 github.com/goccy/pythonwasm2go/p2.Fn1878
func Fn1878(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1879 github.com/goccy/pythonwasm2go/p2.Fn1879
func Fn1879(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1882 github.com/goccy/pythonwasm2go/p2.Fn1882
func Fn1882(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1883 github.com/goccy/pythonwasm2go/p2.Fn1883
func Fn1883(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1884 github.com/goccy/pythonwasm2go/p2.Fn1884
func Fn1884(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1886 github.com/goccy/pythonwasm2go/p2.Fn1886
func Fn1886(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1887 github.com/goccy/pythonwasm2go/p2.Fn1887
func Fn1887(m *base.Module, l0 int32)

//go:linkname Fn1888 github.com/goccy/pythonwasm2go/p2.Fn1888
func Fn1888(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1893 github.com/goccy/pythonwasm2go/p2.Fn1893
func Fn1893(m *base.Module)

//go:linkname Fn1894 github.com/goccy/pythonwasm2go/p2.Fn1894
func Fn1894(m *base.Module, l0 int32) int32

//go:linkname Fn1895 github.com/goccy/pythonwasm2go/p2.Fn1895
func Fn1895(m *base.Module, l0 int32)

//go:linkname Fn1899 github.com/goccy/pythonwasm2go/p2.Fn1899
func Fn1899(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1902 github.com/goccy/pythonwasm2go/p2.Fn1902
func Fn1902(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1903 github.com/goccy/pythonwasm2go/p2.Fn1903
func Fn1903(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1905 github.com/goccy/pythonwasm2go/p2.Fn1905
func Fn1905(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1906 github.com/goccy/pythonwasm2go/p2.Fn1906
func Fn1906(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1907 github.com/goccy/pythonwasm2go/p2.Fn1907
func Fn1907(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1908 github.com/goccy/pythonwasm2go/p2.Fn1908
func Fn1908(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1910 github.com/goccy/pythonwasm2go/p2.Fn1910
func Fn1910(m *base.Module, l0 int32)

//go:linkname Fn1916 github.com/goccy/pythonwasm2go/p2.Fn1916
func Fn1916(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1920 github.com/goccy/pythonwasm2go/p2.Fn1920
func Fn1920(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1924 github.com/goccy/pythonwasm2go/p2.Fn1924
func Fn1924(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1925 github.com/goccy/pythonwasm2go/p2.Fn1925
func Fn1925(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1926 github.com/goccy/pythonwasm2go/p2.Fn1926
func Fn1926(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1929 github.com/goccy/pythonwasm2go/p2.Fn1929
func Fn1929(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1935 github.com/goccy/pythonwasm2go/p2.Fn1935
func Fn1935(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1936 github.com/goccy/pythonwasm2go/p2.Fn1936
func Fn1936(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1937 github.com/goccy/pythonwasm2go/p2.Fn1937
func Fn1937(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1938 github.com/goccy/pythonwasm2go/p2.Fn1938
func Fn1938(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1940 github.com/goccy/pythonwasm2go/p2.Fn1940
func Fn1940(m *base.Module, l0 int32) int32

//go:linkname Fn1943 github.com/goccy/pythonwasm2go/p2.Fn1943
func Fn1943(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1944 github.com/goccy/pythonwasm2go/p2.Fn1944
func Fn1944(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1945 github.com/goccy/pythonwasm2go/p2.Fn1945
func Fn1945(m *base.Module, l0 int32) int32

//go:linkname Fn1946 github.com/goccy/pythonwasm2go/p2.Fn1946
func Fn1946(m *base.Module, l0 int32)

//go:linkname Fn1947 github.com/goccy/pythonwasm2go/p2.Fn1947
func Fn1947(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1950 github.com/goccy/pythonwasm2go/p2.Fn1950
func Fn1950(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1951 github.com/goccy/pythonwasm2go/p2.Fn1951
func Fn1951(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1961 github.com/goccy/pythonwasm2go/p2.Fn1961
func Fn1961(m *base.Module, l0 int32) int32

//go:linkname Fn1970 github.com/goccy/pythonwasm2go/p2.Fn1970
func Fn1970(m *base.Module, l0 int32) int32

//go:linkname Fn1971 github.com/goccy/pythonwasm2go/p2.Fn1971
func Fn1971(m *base.Module, l0 int32) int32

//go:linkname Fn1974 github.com/goccy/pythonwasm2go/p0.Fn1974
func Fn1974(m *base.Module, l0 int32)

//go:linkname Fn1976 github.com/goccy/pythonwasm2go/p0.Fn1976
func Fn1976(m *base.Module, l0 int32)

//go:linkname Fn1977 github.com/goccy/pythonwasm2go/p2.Fn1977
func Fn1977(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1979 github.com/goccy/pythonwasm2go/p2.Fn1979
func Fn1979(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1980 github.com/goccy/pythonwasm2go/p0.Fn1980
func Fn1980(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1983 github.com/goccy/pythonwasm2go/p2.Fn1983
func Fn1983(m *base.Module, l0 int32) int32

//go:linkname Fn1984 github.com/goccy/pythonwasm2go/p2.Fn1984
func Fn1984(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1985 github.com/goccy/pythonwasm2go/p2.Fn1985
func Fn1985(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1986 github.com/goccy/pythonwasm2go/p2.Fn1986
func Fn1986(m *base.Module, l0 int32)

//go:linkname Fn1987 github.com/goccy/pythonwasm2go/p2.Fn1987
func Fn1987(m *base.Module, l0 int32) int32

//go:linkname Fn1988 github.com/goccy/pythonwasm2go/p2.Fn1988
func Fn1988(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1989 github.com/goccy/pythonwasm2go/p2.Fn1989
func Fn1989(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1990 github.com/goccy/pythonwasm2go/p2.Fn1990
func Fn1990(m *base.Module, l0 int32)

//go:linkname Fn1991 github.com/goccy/pythonwasm2go/p2.Fn1991
func Fn1991(m *base.Module, l0 int32) int32

//go:linkname Fn1992 github.com/goccy/pythonwasm2go/p2.Fn1992
func Fn1992(m *base.Module, l0 int32) int32

//go:linkname Fn1993 github.com/goccy/pythonwasm2go/p2.Fn1993
func Fn1993(m *base.Module, l0 int32) int32

//go:linkname Fn1995 github.com/goccy/pythonwasm2go/p2.Fn1995
func Fn1995(m *base.Module, l0 int32) int32

//go:linkname Fn1996 github.com/goccy/pythonwasm2go/p2.Fn1996
func Fn1996(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1997 github.com/goccy/pythonwasm2go/p2.Fn1997
func Fn1997(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1998 github.com/goccy/pythonwasm2go/p2.Fn1998
func Fn1998(m *base.Module, l0 int32)

//go:linkname Fn2000 github.com/goccy/pythonwasm2go/p2.Fn2000
func Fn2000(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2001 github.com/goccy/pythonwasm2go/p2.Fn2001
func Fn2001(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2004 github.com/goccy/pythonwasm2go/p2.Fn2004
func Fn2004(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2005 github.com/goccy/pythonwasm2go/p2.Fn2005
func Fn2005(m *base.Module, l0 int32) int32

//go:linkname Fn2014 github.com/goccy/pythonwasm2go/p2.Fn2014
func Fn2014(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn2016 github.com/goccy/pythonwasm2go/p2.Fn2016
func Fn2016(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2017 github.com/goccy/pythonwasm2go/p2.Fn2017
func Fn2017(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2018 github.com/goccy/pythonwasm2go/p2.Fn2018
func Fn2018(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2019 github.com/goccy/pythonwasm2go/p2.Fn2019
func Fn2019(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2020 github.com/goccy/pythonwasm2go/p2.Fn2020
func Fn2020(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2021 github.com/goccy/pythonwasm2go/p2.Fn2021
func Fn2021(m *base.Module, l0 int32)

//go:linkname Fn2023 github.com/goccy/pythonwasm2go/p2.Fn2023
func Fn2023(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn2037 github.com/goccy/pythonwasm2go/p2.Fn2037
func Fn2037(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2038 github.com/goccy/pythonwasm2go/p2.Fn2038
func Fn2038(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2052 github.com/goccy/pythonwasm2go/p2.Fn2052
func Fn2052(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2069 github.com/goccy/pythonwasm2go/p0.Fn2069
func Fn2069(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2073 github.com/goccy/pythonwasm2go/p0.Fn2073
func Fn2073(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2075 github.com/goccy/pythonwasm2go/p0.Fn2075
func Fn2075(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2078 github.com/goccy/pythonwasm2go/p2.Fn2078
func Fn2078(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2080 github.com/goccy/pythonwasm2go/p2.Fn2080
func Fn2080(m *base.Module, l0 int32) int32

//go:linkname Fn2095 github.com/goccy/pythonwasm2go/p0.Fn2095
func Fn2095(m *base.Module, l0 int32) int32

//go:linkname Fn2096 github.com/goccy/pythonwasm2go/p2.Fn2096
func Fn2096(m *base.Module, l0 int32) int32

//go:linkname Fn2097 github.com/goccy/pythonwasm2go/p2.Fn2097
func Fn2097(m *base.Module, l0 int32) int32

//go:linkname Fn2099 github.com/goccy/pythonwasm2go/p2.Fn2099
func Fn2099(m *base.Module, l0 int32) int32

//go:linkname Fn2101 github.com/goccy/pythonwasm2go/p0.Fn2101
func Fn2101(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2102 github.com/goccy/pythonwasm2go/p2.Fn2102
func Fn2102(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2103 github.com/goccy/pythonwasm2go/p2.Fn2103
func Fn2103(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2104 github.com/goccy/pythonwasm2go/p2.Fn2104
func Fn2104(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2105 github.com/goccy/pythonwasm2go/p0.Fn2105
func Fn2105(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2107 github.com/goccy/pythonwasm2go/p0.Fn2107
func Fn2107(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2108 github.com/goccy/pythonwasm2go/p0.Fn2108
func Fn2108(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2109 github.com/goccy/pythonwasm2go/p0.Fn2109
func Fn2109(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2111 github.com/goccy/pythonwasm2go/p2.Fn2111
func Fn2111(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2126 github.com/goccy/pythonwasm2go/p2.Fn2126
func Fn2126(m *base.Module, l0 int32) int32

//go:linkname Fn2134 github.com/goccy/pythonwasm2go/p2.Fn2134
func Fn2134(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2135 github.com/goccy/pythonwasm2go/p2.Fn2135
func Fn2135(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2156 github.com/goccy/pythonwasm2go/p2.Fn2156
func Fn2156(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2159 github.com/goccy/pythonwasm2go/p2.Fn2159
func Fn2159(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2160 github.com/goccy/pythonwasm2go/p2.Fn2160
func Fn2160(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2162 github.com/goccy/pythonwasm2go/p2.Fn2162
func Fn2162(m *base.Module, l0 int32) int32

//go:linkname Fn2171 github.com/goccy/pythonwasm2go/p0.Fn2171
func Fn2171(m *base.Module, l0 int32) int32

//go:linkname Fn2172 github.com/goccy/pythonwasm2go/p0.Fn2172
func Fn2172(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2173 github.com/goccy/pythonwasm2go/p2.Fn2173
func Fn2173(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2174 github.com/goccy/pythonwasm2go/p2.Fn2174
func Fn2174(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2176 github.com/goccy/pythonwasm2go/p2.Fn2176
func Fn2176(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2177 github.com/goccy/pythonwasm2go/p2.Fn2177
func Fn2177(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2189 github.com/goccy/pythonwasm2go/p2.Fn2189
func Fn2189(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2195 github.com/goccy/pythonwasm2go/p0.Fn2195
func Fn2195(m *base.Module, l0 int32) int32

//go:linkname Fn2196 github.com/goccy/pythonwasm2go/p0.Fn2196
func Fn2196(m *base.Module, l0 int32) int32

//go:linkname Fn2197 github.com/goccy/pythonwasm2go/p0.Fn2197
func Fn2197(m *base.Module, l0 int32) int32

//go:linkname Fn2198 github.com/goccy/pythonwasm2go/p2.Fn2198
func Fn2198(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2200 github.com/goccy/pythonwasm2go/p0.Fn2200
func Fn2200(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2201 github.com/goccy/pythonwasm2go/p0.Fn2201
func Fn2201(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2204 github.com/goccy/pythonwasm2go/p2.Fn2204
func Fn2204(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2220 github.com/goccy/pythonwasm2go/p2.Fn2220
func Fn2220(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2226 github.com/goccy/pythonwasm2go/p2.Fn2226
func Fn2226(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2228 github.com/goccy/pythonwasm2go/p2.Fn2228
func Fn2228(m *base.Module, l0 int32) int32

//go:linkname Fn2229 github.com/goccy/pythonwasm2go/p2.Fn2229
func Fn2229(m *base.Module, l0 int32) int32

//go:linkname Fn2230 github.com/goccy/pythonwasm2go/p2.Fn2230
func Fn2230(m *base.Module, l0 int32) int32

//go:linkname Fn2231 github.com/goccy/pythonwasm2go/p2.Fn2231
func Fn2231(m *base.Module, l0 int32) int32

//go:linkname Fn2232 github.com/goccy/pythonwasm2go/p2.Fn2232
func Fn2232(m *base.Module, l0 int32) int32

//go:linkname Fn2233 github.com/goccy/pythonwasm2go/p2.Fn2233
func Fn2233(m *base.Module, l0 int32) int32

//go:linkname Fn2243 github.com/goccy/pythonwasm2go/p2.Fn2243
func Fn2243(m *base.Module) int32

//go:linkname Fn2244 github.com/goccy/pythonwasm2go/p2.Fn2244
func Fn2244(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2245 github.com/goccy/pythonwasm2go/p2.Fn2245
func Fn2245(m *base.Module, l0 int32)

//go:linkname Fn2246 github.com/goccy/pythonwasm2go/p0.Fn2246
func Fn2246(m *base.Module, l0 int32)

//go:linkname Fn2247 github.com/goccy/pythonwasm2go/p2.Fn2247
func Fn2247(m *base.Module, l0 int32) int32

//go:linkname Fn2249 github.com/goccy/pythonwasm2go/p0.Fn2249
func Fn2249(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2250 github.com/goccy/pythonwasm2go/p0.Fn2250
func Fn2250(m *base.Module, l0 int32) int32

//go:linkname Fn2252 github.com/goccy/pythonwasm2go/p0.Fn2252
func Fn2252(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2254 github.com/goccy/pythonwasm2go/p2.Fn2254
func Fn2254(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2255 github.com/goccy/pythonwasm2go/p2.Fn2255
func Fn2255(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2256 github.com/goccy/pythonwasm2go/p2.Fn2256
func Fn2256(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2257 github.com/goccy/pythonwasm2go/p2.Fn2257
func Fn2257(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2258 github.com/goccy/pythonwasm2go/p0.Fn2258
func Fn2258(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2259 github.com/goccy/pythonwasm2go/p0.Fn2259
func Fn2259(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2260 github.com/goccy/pythonwasm2go/p0.Fn2260
func Fn2260(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2261 github.com/goccy/pythonwasm2go/p0.Fn2261
func Fn2261(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2262 github.com/goccy/pythonwasm2go/p0.Fn2262
func Fn2262(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2271 github.com/goccy/pythonwasm2go/p2.Fn2271
func Fn2271(m *base.Module, l0 int32) int32

//go:linkname Fn2272 github.com/goccy/pythonwasm2go/p2.Fn2272
func Fn2272(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2273 github.com/goccy/pythonwasm2go/p0.Fn2273
func Fn2273(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2274 github.com/goccy/pythonwasm2go/p2.Fn2274
func Fn2274(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2275 github.com/goccy/pythonwasm2go/p2.Fn2275
func Fn2275(m *base.Module, l0 int32) int32

//go:linkname Fn2277 github.com/goccy/pythonwasm2go/p2.Fn2277
func Fn2277(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2279 github.com/goccy/pythonwasm2go/p0.Fn2279
func Fn2279(m *base.Module, l0 int32) int32

//go:linkname Fn2280 github.com/goccy/pythonwasm2go/p2.Fn2280
func Fn2280(m *base.Module, l0 int32) int32

//go:linkname Fn2282 github.com/goccy/pythonwasm2go/p2.Fn2282
func Fn2282(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2283 github.com/goccy/pythonwasm2go/p0.Fn2283
func Fn2283(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2284 github.com/goccy/pythonwasm2go/p2.Fn2284
func Fn2284(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2285 github.com/goccy/pythonwasm2go/p2.Fn2285
func Fn2285(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2286 github.com/goccy/pythonwasm2go/p2.Fn2286
func Fn2286(m *base.Module, l0 int32) int32

//go:linkname Fn2288 github.com/goccy/pythonwasm2go/p0.Fn2288
func Fn2288(m *base.Module, l0 int32) int32

//go:linkname Fn2289 github.com/goccy/pythonwasm2go/p2.Fn2289
func Fn2289(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2290 github.com/goccy/pythonwasm2go/p2.Fn2290
func Fn2290(m *base.Module, l0 int32) int32

//go:linkname Fn2291 github.com/goccy/pythonwasm2go/p2.Fn2291
func Fn2291(m *base.Module, l0 int32) int32

//go:linkname Fn2292 github.com/goccy/pythonwasm2go/p2.Fn2292
func Fn2292(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2293 github.com/goccy/pythonwasm2go/p2.Fn2293
func Fn2293(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2295 github.com/goccy/pythonwasm2go/p2.Fn2295
func Fn2295(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2296 github.com/goccy/pythonwasm2go/p2.Fn2296
func Fn2296(m *base.Module, l0 int32) int32

//go:linkname Fn2297 github.com/goccy/pythonwasm2go/p2.Fn2297
func Fn2297(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2298 github.com/goccy/pythonwasm2go/p2.Fn2298
func Fn2298(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2305 github.com/goccy/pythonwasm2go/p0.Fn2305
func Fn2305(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2310 github.com/goccy/pythonwasm2go/p2.Fn2310
func Fn2310(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2312 github.com/goccy/pythonwasm2go/p2.Fn2312
func Fn2312(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2330 github.com/goccy/pythonwasm2go/p0.Fn2330
func Fn2330(m *base.Module, l0 int32) int32

//go:linkname Fn2335 github.com/goccy/pythonwasm2go/p2.Fn2335
func Fn2335(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2337 github.com/goccy/pythonwasm2go/p2.Fn2337
func Fn2337(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2338 github.com/goccy/pythonwasm2go/p0.Fn2338
func Fn2338(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2339 github.com/goccy/pythonwasm2go/p0.Fn2339
func Fn2339(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2340 github.com/goccy/pythonwasm2go/p2.Fn2340
func Fn2340(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2341 github.com/goccy/pythonwasm2go/p0.Fn2341
func Fn2341(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2343 github.com/goccy/pythonwasm2go/p2.Fn2343
func Fn2343(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2345 github.com/goccy/pythonwasm2go/p2.Fn2345
func Fn2345(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2363 github.com/goccy/pythonwasm2go/p0.Fn2363
func Fn2363(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2380 github.com/goccy/pythonwasm2go/p2.Fn2380
func Fn2380(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2386 github.com/goccy/pythonwasm2go/p2.Fn2386
func Fn2386(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2400 github.com/goccy/pythonwasm2go/p2.Fn2400
func Fn2400(m *base.Module, l0 int32) int32

//go:linkname Fn2408 github.com/goccy/pythonwasm2go/p2.Fn2408
func Fn2408(m *base.Module) int32

//go:linkname Fn2410 github.com/goccy/pythonwasm2go/p0.Fn2410
func Fn2410(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2462 github.com/goccy/pythonwasm2go/p2.Fn2462
func Fn2462(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2520 github.com/goccy/pythonwasm2go/p2.Fn2520
func Fn2520(m *base.Module) int32

//go:linkname Fn2523 github.com/goccy/pythonwasm2go/p2.Fn2523
func Fn2523(m *base.Module, l0 int32) int32

//go:linkname Fn2530 github.com/goccy/pythonwasm2go/p2.Fn2530
func Fn2530(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn2535 github.com/goccy/pythonwasm2go/p2.Fn2535
func Fn2535(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2536 github.com/goccy/pythonwasm2go/p2.Fn2536
func Fn2536(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2537 github.com/goccy/pythonwasm2go/p2.Fn2537
func Fn2537(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2538 github.com/goccy/pythonwasm2go/p2.Fn2538
func Fn2538(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2550 github.com/goccy/pythonwasm2go/p2.Fn2550
func Fn2550(m *base.Module, l0 int32) int32

//go:linkname Fn2601 github.com/goccy/pythonwasm2go/p2.Fn2601
func Fn2601(m *base.Module, l0 int32) int32

//go:linkname Fn2602 github.com/goccy/pythonwasm2go/p2.Fn2602
func Fn2602(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2603 github.com/goccy/pythonwasm2go/p0.Fn2603
func Fn2603(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2604 github.com/goccy/pythonwasm2go/p2.Fn2604
func Fn2604(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn2606 github.com/goccy/pythonwasm2go/p2.Fn2606
func Fn2606(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2607 github.com/goccy/pythonwasm2go/p2.Fn2607
func Fn2607(m *base.Module, l0 int32) int32

//go:linkname Fn2608 github.com/goccy/pythonwasm2go/p2.Fn2608
func Fn2608(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2610 github.com/goccy/pythonwasm2go/p0.Fn2610
func Fn2610(m *base.Module, l0 int32) int32

//go:linkname Fn2611 github.com/goccy/pythonwasm2go/p2.Fn2611
func Fn2611(m *base.Module, l0 int32) int32

//go:linkname Fn2612 github.com/goccy/pythonwasm2go/p2.Fn2612
func Fn2612(m *base.Module, l0 int32) int32

//go:linkname Fn2614 github.com/goccy/pythonwasm2go/p0.Fn2614
func Fn2614(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2615 github.com/goccy/pythonwasm2go/p0.Fn2615
func Fn2615(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2616 github.com/goccy/pythonwasm2go/p2.Fn2616
func Fn2616(m *base.Module, l0 int32) int32

//go:linkname Fn2617 github.com/goccy/pythonwasm2go/p0.Fn2617
func Fn2617(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2620 github.com/goccy/pythonwasm2go/p2.Fn2620
func Fn2620(m *base.Module, l0 int32) int32

//go:linkname Fn2622 github.com/goccy/pythonwasm2go/p0.Fn2622
func Fn2622(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2624 github.com/goccy/pythonwasm2go/p0.Fn2624
func Fn2624(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2625 github.com/goccy/pythonwasm2go/p0.Fn2625
func Fn2625(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2627 github.com/goccy/pythonwasm2go/p0.Fn2627
func Fn2627(m *base.Module, l0 int32) int32

//go:linkname Fn2628 github.com/goccy/pythonwasm2go/p0.Fn2628
func Fn2628(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2630 github.com/goccy/pythonwasm2go/p0.Fn2630
func Fn2630(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2631 github.com/goccy/pythonwasm2go/p2.Fn2631
func Fn2631(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2632 github.com/goccy/pythonwasm2go/p0.Fn2632
func Fn2632(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2633 github.com/goccy/pythonwasm2go/p0.Fn2633
func Fn2633(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2634 github.com/goccy/pythonwasm2go/p0.Fn2634
func Fn2634(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2635 github.com/goccy/pythonwasm2go/p0.Fn2635
func Fn2635(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2638 github.com/goccy/pythonwasm2go/p2.Fn2638
func Fn2638(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2639 github.com/goccy/pythonwasm2go/p0.Fn2639
func Fn2639(m *base.Module, l0 int32) int32

//go:linkname Fn2640 github.com/goccy/pythonwasm2go/p2.Fn2640
func Fn2640(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2642 github.com/goccy/pythonwasm2go/p2.Fn2642
func Fn2642(m *base.Module, l0 int32) int32

//go:linkname Fn2643 github.com/goccy/pythonwasm2go/p0.Fn2643
func Fn2643(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2645 github.com/goccy/pythonwasm2go/p0.Fn2645
func Fn2645(m *base.Module, l0 int32) int32

//go:linkname Fn2646 github.com/goccy/pythonwasm2go/p2.Fn2646
func Fn2646(m *base.Module, l0 int32)

//go:linkname Fn2647 github.com/goccy/pythonwasm2go/p0.Fn2647
func Fn2647(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2648 github.com/goccy/pythonwasm2go/p0.Fn2648
func Fn2648(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2649 github.com/goccy/pythonwasm2go/p0.Fn2649
func Fn2649(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2653 github.com/goccy/pythonwasm2go/p0.Fn2653
func Fn2653(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2657 github.com/goccy/pythonwasm2go/p0.Fn2657
func Fn2657(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2658 github.com/goccy/pythonwasm2go/p0.Fn2658
func Fn2658(m *base.Module, l0 int32) int32

//go:linkname Fn2659 github.com/goccy/pythonwasm2go/p2.Fn2659
func Fn2659(m *base.Module, l0 int32) int32

//go:linkname Fn2661 github.com/goccy/pythonwasm2go/p0.Fn2661
func Fn2661(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2662 github.com/goccy/pythonwasm2go/p0.Fn2662
func Fn2662(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2664 github.com/goccy/pythonwasm2go/p0.Fn2664
func Fn2664(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2665 github.com/goccy/pythonwasm2go/p0.Fn2665
func Fn2665(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2666 github.com/goccy/pythonwasm2go/p0.Fn2666
func Fn2666(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2667 github.com/goccy/pythonwasm2go/p0.Fn2667
func Fn2667(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2668 github.com/goccy/pythonwasm2go/p0.Fn2668
func Fn2668(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2672 github.com/goccy/pythonwasm2go/p0.Fn2672
func Fn2672(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32

//go:linkname Fn2674 github.com/goccy/pythonwasm2go/p0.Fn2674
func Fn2674(m *base.Module, l0 int32) int32

//go:linkname Fn2676 github.com/goccy/pythonwasm2go/p0.Fn2676
func Fn2676(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2681 github.com/goccy/pythonwasm2go/p0.Fn2681
func Fn2681(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2682 github.com/goccy/pythonwasm2go/p0.Fn2682
func Fn2682(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2687 github.com/goccy/pythonwasm2go/p0.Fn2687
func Fn2687(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32

//go:linkname Fn2688 github.com/goccy/pythonwasm2go/p0.Fn2688
func Fn2688(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn2690 github.com/goccy/pythonwasm2go/p0.Fn2690
func Fn2690(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2692 github.com/goccy/pythonwasm2go/p0.Fn2692
func Fn2692(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2693 github.com/goccy/pythonwasm2go/p0.Fn2693
func Fn2693(m *base.Module, l0 int32) int32

//go:linkname Fn2694 github.com/goccy/pythonwasm2go/p0.Fn2694
func Fn2694(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2697 github.com/goccy/pythonwasm2go/p2.Fn2697
func Fn2697(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2699 github.com/goccy/pythonwasm2go/p2.Fn2699
func Fn2699(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2700 github.com/goccy/pythonwasm2go/p2.Fn2700
func Fn2700(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2701 github.com/goccy/pythonwasm2go/p2.Fn2701
func Fn2701(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2702 github.com/goccy/pythonwasm2go/p2.Fn2702
func Fn2702(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2703 github.com/goccy/pythonwasm2go/p2.Fn2703
func Fn2703(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2704 github.com/goccy/pythonwasm2go/p2.Fn2704
func Fn2704(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2705 github.com/goccy/pythonwasm2go/p2.Fn2705
func Fn2705(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2706 github.com/goccy/pythonwasm2go/p0.Fn2706
func Fn2706(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2707 github.com/goccy/pythonwasm2go/p0.Fn2707
func Fn2707(m *base.Module, l0 int32) int32

//go:linkname Fn2708 github.com/goccy/pythonwasm2go/p2.Fn2708
func Fn2708(m *base.Module, l0 int32) int32

//go:linkname Fn2709 github.com/goccy/pythonwasm2go/p0.Fn2709
func Fn2709(m *base.Module, l0 int32) int32

//go:linkname Fn2710 github.com/goccy/pythonwasm2go/p2.Fn2710
func Fn2710(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2711 github.com/goccy/pythonwasm2go/p0.Fn2711
func Fn2711(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2712 github.com/goccy/pythonwasm2go/p0.Fn2712
func Fn2712(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2718 github.com/goccy/pythonwasm2go/p2.Fn2718
func Fn2718(m *base.Module) int32

//go:linkname Fn2719 github.com/goccy/pythonwasm2go/p0.Fn2719
func Fn2719(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2723 github.com/goccy/pythonwasm2go/p2.Fn2723
func Fn2723(m *base.Module, l0 int32) int32

//go:linkname Fn2724 github.com/goccy/pythonwasm2go/p2.Fn2724
func Fn2724(m *base.Module, l0 int32) int32

//go:linkname Fn2726 github.com/goccy/pythonwasm2go/p2.Fn2726
func Fn2726(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2728 github.com/goccy/pythonwasm2go/p2.Fn2728
func Fn2728(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2729 github.com/goccy/pythonwasm2go/p0.Fn2729
func Fn2729(m *base.Module, l0 int32) int32

//go:linkname Fn2731 github.com/goccy/pythonwasm2go/p2.Fn2731
func Fn2731(m *base.Module, l0 int32) int32

//go:linkname Fn2732 github.com/goccy/pythonwasm2go/p2.Fn2732
func Fn2732(m *base.Module, l0 int32) int32

//go:linkname Fn2735 github.com/goccy/pythonwasm2go/p2.Fn2735
func Fn2735(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn2736 github.com/goccy/pythonwasm2go/p0.Fn2736
func Fn2736(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2737 github.com/goccy/pythonwasm2go/p2.Fn2737
func Fn2737(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2738 github.com/goccy/pythonwasm2go/p2.Fn2738
func Fn2738(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2739 github.com/goccy/pythonwasm2go/p2.Fn2739
func Fn2739(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2741 github.com/goccy/pythonwasm2go/p2.Fn2741
func Fn2741(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2742 github.com/goccy/pythonwasm2go/p2.Fn2742
func Fn2742(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2743 github.com/goccy/pythonwasm2go/p2.Fn2743
func Fn2743(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2748 github.com/goccy/pythonwasm2go/p2.Fn2748
func Fn2748(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2749 github.com/goccy/pythonwasm2go/p2.Fn2749
func Fn2749(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2752 github.com/goccy/pythonwasm2go/p2.Fn2752
func Fn2752(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2754 github.com/goccy/pythonwasm2go/p2.Fn2754
func Fn2754(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2756 github.com/goccy/pythonwasm2go/p2.Fn2756
func Fn2756(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2757 github.com/goccy/pythonwasm2go/p0.Fn2757
func Fn2757(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2759 github.com/goccy/pythonwasm2go/p2.Fn2759
func Fn2759(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2761 github.com/goccy/pythonwasm2go/p2.Fn2761
func Fn2761(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2763 github.com/goccy/pythonwasm2go/p0.Fn2763
func Fn2763(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2764 github.com/goccy/pythonwasm2go/p2.Fn2764
func Fn2764(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2766 github.com/goccy/pythonwasm2go/p2.Fn2766
func Fn2766(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2769 github.com/goccy/pythonwasm2go/p2.Fn2769
func Fn2769(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2770 github.com/goccy/pythonwasm2go/p0.Fn2770
func Fn2770(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2771 github.com/goccy/pythonwasm2go/p0.Fn2771
func Fn2771(m *base.Module, l0 int32) int32

//go:linkname Fn2772 github.com/goccy/pythonwasm2go/p2.Fn2772
func Fn2772(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2773 github.com/goccy/pythonwasm2go/p0.Fn2773
func Fn2773(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2775 github.com/goccy/pythonwasm2go/p0.Fn2775
func Fn2775(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2776 github.com/goccy/pythonwasm2go/p0.Fn2776
func Fn2776(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2777 github.com/goccy/pythonwasm2go/p2.Fn2777
func Fn2777(m *base.Module, l0 int32) int32

//go:linkname Fn2778 github.com/goccy/pythonwasm2go/p2.Fn2778
func Fn2778(m *base.Module, l0 int32)

//go:linkname Fn2780 github.com/goccy/pythonwasm2go/p2.Fn2780
func Fn2780(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2781 github.com/goccy/pythonwasm2go/p2.Fn2781
func Fn2781(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2782 github.com/goccy/pythonwasm2go/p2.Fn2782
func Fn2782(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2783 github.com/goccy/pythonwasm2go/p2.Fn2783
func Fn2783(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2784 github.com/goccy/pythonwasm2go/p2.Fn2784
func Fn2784(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2785 github.com/goccy/pythonwasm2go/p2.Fn2785
func Fn2785(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2786 github.com/goccy/pythonwasm2go/p2.Fn2786
func Fn2786(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2787 github.com/goccy/pythonwasm2go/p2.Fn2787
func Fn2787(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2788 github.com/goccy/pythonwasm2go/p0.Fn2788
func Fn2788(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2789 github.com/goccy/pythonwasm2go/p2.Fn2789
func Fn2789(m *base.Module, l0 int32) int32

//go:linkname Fn2801 github.com/goccy/pythonwasm2go/p2.Fn2801
func Fn2801(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2807 github.com/goccy/pythonwasm2go/p0.Fn2807
func Fn2807(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2808 github.com/goccy/pythonwasm2go/p0.Fn2808
func Fn2808(m *base.Module, l0 int32) int32

//go:linkname Fn2817 github.com/goccy/pythonwasm2go/p2.Fn2817
func Fn2817(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2818 github.com/goccy/pythonwasm2go/p2.Fn2818
func Fn2818(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2819 github.com/goccy/pythonwasm2go/p2.Fn2819
func Fn2819(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2821 github.com/goccy/pythonwasm2go/p2.Fn2821
func Fn2821(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2822 github.com/goccy/pythonwasm2go/p2.Fn2822
func Fn2822(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2825 github.com/goccy/pythonwasm2go/p2.Fn2825
func Fn2825(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2826 github.com/goccy/pythonwasm2go/p2.Fn2826
func Fn2826(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2827 github.com/goccy/pythonwasm2go/p2.Fn2827
func Fn2827(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2830 github.com/goccy/pythonwasm2go/p2.Fn2830
func Fn2830(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2833 github.com/goccy/pythonwasm2go/p2.Fn2833
func Fn2833(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2886 github.com/goccy/pythonwasm2go/p2.Fn2886
func Fn2886(m *base.Module, l0 int32) int32

//go:linkname Fn2901 github.com/goccy/pythonwasm2go/p2.Fn2901
func Fn2901(m *base.Module, l0 int32) int32

//go:linkname Fn2903 github.com/goccy/pythonwasm2go/p2.Fn2903
func Fn2903(m *base.Module, l0 int32) int32

//go:linkname Fn2912 github.com/goccy/pythonwasm2go/p2.Fn2912
func Fn2912(m *base.Module, l0 int32) int32

//go:linkname Fn2913 github.com/goccy/pythonwasm2go/p2.Fn2913
func Fn2913(m *base.Module, l0 int32) int32

//go:linkname Fn2915 github.com/goccy/pythonwasm2go/p2.Fn2915
func Fn2915(m *base.Module, l0 int32) int32

//go:linkname Fn2916 github.com/goccy/pythonwasm2go/p2.Fn2916
func Fn2916(m *base.Module, l0 int32) int32

//go:linkname Fn2923 github.com/goccy/pythonwasm2go/p2.Fn2923
func Fn2923(m *base.Module, l0 int32) int32

//go:linkname Fn2925 github.com/goccy/pythonwasm2go/p2.Fn2925
func Fn2925(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2932 github.com/goccy/pythonwasm2go/p2.Fn2932
func Fn2932(m *base.Module, l0 int32) int32

//go:linkname Fn2935 github.com/goccy/pythonwasm2go/p2.Fn2935
func Fn2935(m *base.Module, l0 int32)

//go:linkname Fn2938 github.com/goccy/pythonwasm2go/p2.Fn2938
func Fn2938(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2953 github.com/goccy/pythonwasm2go/p2.Fn2953
func Fn2953(m *base.Module, l0 int32) int32

//go:linkname Fn2954 github.com/goccy/pythonwasm2go/p2.Fn2954
func Fn2954(m *base.Module, l0 int32) int32

//go:linkname Fn2956 github.com/goccy/pythonwasm2go/p2.Fn2956
func Fn2956(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2971 github.com/goccy/pythonwasm2go/p2.Fn2971
func Fn2971(m *base.Module, l0 int32) int32

//go:linkname Fn2978 github.com/goccy/pythonwasm2go/p0.Fn2978
func Fn2978(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2980 github.com/goccy/pythonwasm2go/p2.Fn2980
func Fn2980(m *base.Module, l0 int32) int32

//go:linkname Fn2981 github.com/goccy/pythonwasm2go/p2.Fn2981
func Fn2981(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2982 github.com/goccy/pythonwasm2go/p2.Fn2982
func Fn2982(m *base.Module, l0 int32) int32

//go:linkname Fn2984 github.com/goccy/pythonwasm2go/p2.Fn2984
func Fn2984(m *base.Module, l0 int32) int32

//go:linkname Fn2985 github.com/goccy/pythonwasm2go/p2.Fn2985
func Fn2985(m *base.Module, l0 int32)

//go:linkname Fn2987 github.com/goccy/pythonwasm2go/p2.Fn2987
func Fn2987(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3031 github.com/goccy/pythonwasm2go/p0.Fn3031
func Fn3031(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3035 github.com/goccy/pythonwasm2go/p0.Fn3035
func Fn3035(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3036 github.com/goccy/pythonwasm2go/p0.Fn3036
func Fn3036(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3037 github.com/goccy/pythonwasm2go/p0.Fn3037
func Fn3037(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3038 github.com/goccy/pythonwasm2go/p0.Fn3038
func Fn3038(m *base.Module) int32

//go:linkname Fn3039 github.com/goccy/pythonwasm2go/p0.Fn3039
func Fn3039(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32

//go:linkname Fn3043 github.com/goccy/pythonwasm2go/p0.Fn3043
func Fn3043(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3045 github.com/goccy/pythonwasm2go/p2.Fn3045
func Fn3045(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3057 github.com/goccy/pythonwasm2go/p2.Fn3057
func Fn3057(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3087 github.com/goccy/pythonwasm2go/p2.Fn3087
func Fn3087(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3088 github.com/goccy/pythonwasm2go/p2.Fn3088
func Fn3088(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3089 github.com/goccy/pythonwasm2go/p2.Fn3089
func Fn3089(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3092 github.com/goccy/pythonwasm2go/p2.Fn3092
func Fn3092(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn3093 github.com/goccy/pythonwasm2go/p2.Fn3093
func Fn3093(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3096 github.com/goccy/pythonwasm2go/p2.Fn3096
func Fn3096(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3107 github.com/goccy/pythonwasm2go/p2.Fn3107
func Fn3107(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn3109 github.com/goccy/pythonwasm2go/p2.Fn3109
func Fn3109(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3111 github.com/goccy/pythonwasm2go/p2.Fn3111
func Fn3111(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3112 github.com/goccy/pythonwasm2go/p2.Fn3112
func Fn3112(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn3113 github.com/goccy/pythonwasm2go/p2.Fn3113
func Fn3113(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn3114 github.com/goccy/pythonwasm2go/p2.Fn3114
func Fn3114(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3115 github.com/goccy/pythonwasm2go/p2.Fn3115
func Fn3115(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3116 github.com/goccy/pythonwasm2go/p2.Fn3116
func Fn3116(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3117 github.com/goccy/pythonwasm2go/p2.Fn3117
func Fn3117(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3121 github.com/goccy/pythonwasm2go/p2.Fn3121
func Fn3121(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn3122 github.com/goccy/pythonwasm2go/p2.Fn3122
func Fn3122(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn3124 github.com/goccy/pythonwasm2go/p2.Fn3124
func Fn3124(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3125 github.com/goccy/pythonwasm2go/p2.Fn3125
func Fn3125(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3129 github.com/goccy/pythonwasm2go/p2.Fn3129
func Fn3129(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3139 github.com/goccy/pythonwasm2go/p2.Fn3139
func Fn3139(m *base.Module) int32

//go:linkname Fn3141 github.com/goccy/pythonwasm2go/p2.Fn3141
func Fn3141(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3143 github.com/goccy/pythonwasm2go/p0.Fn3143
func Fn3143(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3145 github.com/goccy/pythonwasm2go/p0.Fn3145
func Fn3145(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3147 github.com/goccy/pythonwasm2go/p2.Fn3147
func Fn3147(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3149 github.com/goccy/pythonwasm2go/p0.Fn3149
func Fn3149(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3150 github.com/goccy/pythonwasm2go/p2.Fn3150
func Fn3150(m *base.Module, l0 int32) int32

//go:linkname Fn3175 github.com/goccy/pythonwasm2go/p2.Fn3175
func Fn3175(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3189 github.com/goccy/pythonwasm2go/p2.Fn3189
func Fn3189(m *base.Module, l0 int32) int32

//go:linkname Fn3191 github.com/goccy/pythonwasm2go/p0.Fn3191
func Fn3191(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3200 github.com/goccy/pythonwasm2go/p2.Fn3200
func Fn3200(m *base.Module, l0 int32) int32

//go:linkname Fn3203 github.com/goccy/pythonwasm2go/p2.Fn3203
func Fn3203(m *base.Module, l0 int32) int32

//go:linkname Fn3207 github.com/goccy/pythonwasm2go/p2.Fn3207
func Fn3207(m *base.Module, l0 int32) int32

//go:linkname Fn3208 github.com/goccy/pythonwasm2go/p2.Fn3208
func Fn3208(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3252 github.com/goccy/pythonwasm2go/p2.Fn3252
func Fn3252(m *base.Module, l0 int32) int32

//go:linkname Fn3276 github.com/goccy/pythonwasm2go/p2.Fn3276
func Fn3276(m *base.Module, l0 int32) int32

//go:linkname Fn3278 github.com/goccy/pythonwasm2go/p2.Fn3278
func Fn3278(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3279 github.com/goccy/pythonwasm2go/p0.Fn3279
func Fn3279(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3283 github.com/goccy/pythonwasm2go/p0.Fn3283
func Fn3283(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3284 github.com/goccy/pythonwasm2go/p0.Fn3284
func Fn3284(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3290 github.com/goccy/pythonwasm2go/p2.Fn3290
func Fn3290(m *base.Module, l0 int32) int32

//go:linkname Fn3295 github.com/goccy/pythonwasm2go/p2.Fn3295
func Fn3295(m *base.Module, l0 int32) int32

//go:linkname Fn3307 github.com/goccy/pythonwasm2go/p2.Fn3307
func Fn3307(m *base.Module, l0 int32)

//go:linkname Fn3308 github.com/goccy/pythonwasm2go/p2.Fn3308
func Fn3308(m *base.Module, l0 int32)

//go:linkname Fn3309 github.com/goccy/pythonwasm2go/p2.Fn3309
func Fn3309(m *base.Module, l0 int32) int32

//go:linkname Fn3310 github.com/goccy/pythonwasm2go/p2.Fn3310
func Fn3310(m *base.Module, l0 int32) int32

//go:linkname Fn3311 github.com/goccy/pythonwasm2go/p2.Fn3311
func Fn3311(m *base.Module) int32

//go:linkname Fn3312 github.com/goccy/pythonwasm2go/p2.Fn3312
func Fn3312(m *base.Module) int32

//go:linkname Fn3313 github.com/goccy/pythonwasm2go/p2.Fn3313
func Fn3313(m *base.Module, l0 int32) int32

//go:linkname Fn3315 github.com/goccy/pythonwasm2go/p2.Fn3315
func Fn3315(m *base.Module, l0 int32) int32

//go:linkname Fn3316 github.com/goccy/pythonwasm2go/p2.Fn3316
func Fn3316(m *base.Module) int32

//go:linkname Fn3317 github.com/goccy/pythonwasm2go/p2.Fn3317
func Fn3317(m *base.Module, l0 int32) int32

//go:linkname Fn3319 github.com/goccy/pythonwasm2go/p2.Fn3319
func Fn3319(m *base.Module) int32

//go:linkname Fn3320 github.com/goccy/pythonwasm2go/p2.Fn3320
func Fn3320(m *base.Module, l0 int32) int32

//go:linkname Fn3321 github.com/goccy/pythonwasm2go/p2.Fn3321
func Fn3321(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3324 github.com/goccy/pythonwasm2go/p2.Fn3324
func Fn3324(m *base.Module, l0 int32) int32

//go:linkname Fn3325 github.com/goccy/pythonwasm2go/p2.Fn3325
func Fn3325(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3326 github.com/goccy/pythonwasm2go/p2.Fn3326
func Fn3326(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3327 github.com/goccy/pythonwasm2go/p2.Fn3327
func Fn3327(m *base.Module, l0 int32) int32

//go:linkname Fn3330 github.com/goccy/pythonwasm2go/p2.Fn3330
func Fn3330(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3334 github.com/goccy/pythonwasm2go/p2.Fn3334
func Fn3334(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3337 github.com/goccy/pythonwasm2go/p2.Fn3337
func Fn3337(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3339 github.com/goccy/pythonwasm2go/p0.Fn3339
func Fn3339(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3342 github.com/goccy/pythonwasm2go/p0.Fn3342
func Fn3342(m *base.Module, l0 int32) int32

//go:linkname Fn3343 github.com/goccy/pythonwasm2go/p0.Fn3343
func Fn3343(m *base.Module, l0 int32) int32

//go:linkname Fn3346 github.com/goccy/pythonwasm2go/p2.Fn3346
func Fn3346(m *base.Module, l0 int32)

//go:linkname Fn3348 github.com/goccy/pythonwasm2go/p2.Fn3348
func Fn3348(m *base.Module, l0 int32) int32

//go:linkname Fn3350 github.com/goccy/pythonwasm2go/p2.Fn3350
func Fn3350(m *base.Module, l0 int32) int32

//go:linkname Fn3382 github.com/goccy/pythonwasm2go/p0.Fn3382
func Fn3382(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3391 github.com/goccy/pythonwasm2go/p2.Fn3391
func Fn3391(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3402 github.com/goccy/pythonwasm2go/p2.Fn3402
func Fn3402(m *base.Module, l0 int32) int32

//go:linkname Fn3410 github.com/goccy/pythonwasm2go/p2.Fn3410
func Fn3410(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3415 github.com/goccy/pythonwasm2go/p2.Fn3415
func Fn3415(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3416 github.com/goccy/pythonwasm2go/p2.Fn3416
func Fn3416(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3432 github.com/goccy/pythonwasm2go/p2.Fn3432
func Fn3432(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3440 github.com/goccy/pythonwasm2go/p2.Fn3440
func Fn3440(m *base.Module, l0 int32) int32

//go:linkname Fn3449 github.com/goccy/pythonwasm2go/p0.Fn3449
func Fn3449(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn3454 github.com/goccy/pythonwasm2go/p2.Fn3454
func Fn3454(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3455 github.com/goccy/pythonwasm2go/p2.Fn3455
func Fn3455(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3456 github.com/goccy/pythonwasm2go/p2.Fn3456
func Fn3456(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3460 github.com/goccy/pythonwasm2go/p0.Fn3460
func Fn3460(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3461 github.com/goccy/pythonwasm2go/p0.Fn3461
func Fn3461(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3462 github.com/goccy/pythonwasm2go/p0.Fn3462
func Fn3462(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3475 github.com/goccy/pythonwasm2go/p2.Fn3475
func Fn3475(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3476 github.com/goccy/pythonwasm2go/p2.Fn3476
func Fn3476(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3481 github.com/goccy/pythonwasm2go/p2.Fn3481
func Fn3481(m *base.Module, l0 int32)

//go:linkname Fn3483 github.com/goccy/pythonwasm2go/p2.Fn3483
func Fn3483(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3485 github.com/goccy/pythonwasm2go/p2.Fn3485
func Fn3485(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3487 github.com/goccy/pythonwasm2go/p2.Fn3487
func Fn3487(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3489 github.com/goccy/pythonwasm2go/p2.Fn3489
func Fn3489(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3490 github.com/goccy/pythonwasm2go/p2.Fn3490
func Fn3490(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3491 github.com/goccy/pythonwasm2go/p2.Fn3491
func Fn3491(m *base.Module, l0 int32) int32

//go:linkname Fn3492 github.com/goccy/pythonwasm2go/p2.Fn3492
func Fn3492(m *base.Module, l0 int32) int32

//go:linkname Fn3493 github.com/goccy/pythonwasm2go/p2.Fn3493
func Fn3493(m *base.Module, l0 int32) int32

//go:linkname Fn3495 github.com/goccy/pythonwasm2go/p2.Fn3495
func Fn3495(m *base.Module, l0 int32) int32

//go:linkname Fn3496 github.com/goccy/pythonwasm2go/p2.Fn3496
func Fn3496(m *base.Module, l0 int32) int32

//go:linkname Fn3497 github.com/goccy/pythonwasm2go/p2.Fn3497
func Fn3497(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3499 github.com/goccy/pythonwasm2go/p2.Fn3499
func Fn3499(m *base.Module, l0 int32) int32

//go:linkname Fn3500 github.com/goccy/pythonwasm2go/p2.Fn3500
func Fn3500(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3502 github.com/goccy/pythonwasm2go/p2.Fn3502
func Fn3502(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3503 github.com/goccy/pythonwasm2go/p2.Fn3503
func Fn3503(m *base.Module, l0 int32)

//go:linkname Fn3504 github.com/goccy/pythonwasm2go/p2.Fn3504
func Fn3504(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3507 github.com/goccy/pythonwasm2go/p2.Fn3507
func Fn3507(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3510 github.com/goccy/pythonwasm2go/p2.Fn3510
func Fn3510(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3516 github.com/goccy/pythonwasm2go/p2.Fn3516
func Fn3516(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3518 github.com/goccy/pythonwasm2go/p2.Fn3518
func Fn3518(m *base.Module, l0 int32) int32

//go:linkname Fn3519 github.com/goccy/pythonwasm2go/p2.Fn3519
func Fn3519(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3520 github.com/goccy/pythonwasm2go/p2.Fn3520
func Fn3520(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3522 github.com/goccy/pythonwasm2go/p2.Fn3522
func Fn3522(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3524 github.com/goccy/pythonwasm2go/p2.Fn3524
func Fn3524(m *base.Module) int32

//go:linkname Fn3525 github.com/goccy/pythonwasm2go/p2.Fn3525
func Fn3525(m *base.Module, l0 int32)

//go:linkname Fn3526 github.com/goccy/pythonwasm2go/p2.Fn3526
func Fn3526(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn3527 github.com/goccy/pythonwasm2go/p2.Fn3527
func Fn3527(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3529 github.com/goccy/pythonwasm2go/p2.Fn3529
func Fn3529(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3532 github.com/goccy/pythonwasm2go/p2.Fn3532
func Fn3532(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3534 github.com/goccy/pythonwasm2go/p2.Fn3534
func Fn3534(m *base.Module, l0 int32) int32

//go:linkname Fn3546 github.com/goccy/pythonwasm2go/p2.Fn3546
func Fn3546(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3547 github.com/goccy/pythonwasm2go/p2.Fn3547
func Fn3547(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3553 github.com/goccy/pythonwasm2go/p2.Fn3553
func Fn3553(m *base.Module, l0 int32)

//go:linkname Fn3555 github.com/goccy/pythonwasm2go/p0.Fn3555
func Fn3555(m *base.Module) int32

//go:linkname Fn3556 github.com/goccy/pythonwasm2go/p2.Fn3556
func Fn3556(m *base.Module, l0 int32)

//go:linkname Fn3557 github.com/goccy/pythonwasm2go/p2.Fn3557
func Fn3557(m *base.Module, l0 int32)

//go:linkname Fn3558 github.com/goccy/pythonwasm2go/p2.Fn3558
func Fn3558(m *base.Module, l0 int32)

//go:linkname Fn3563 github.com/goccy/pythonwasm2go/p0.Fn3563
func Fn3563(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3565 github.com/goccy/pythonwasm2go/p2.Fn3565
func Fn3565(m *base.Module, l0 int32) int32

//go:linkname Fn3566 github.com/goccy/pythonwasm2go/p0.Fn3566
func Fn3566(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3567 github.com/goccy/pythonwasm2go/p2.Fn3567
func Fn3567(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3568 github.com/goccy/pythonwasm2go/p2.Fn3568
func Fn3568(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3569 github.com/goccy/pythonwasm2go/p2.Fn3569
func Fn3569(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3572 github.com/goccy/pythonwasm2go/p2.Fn3572
func Fn3572(m *base.Module, l0 int32)

//go:linkname Fn3573 github.com/goccy/pythonwasm2go/p2.Fn3573
func Fn3573(m *base.Module, l0 int32)

//go:linkname Fn3574 github.com/goccy/pythonwasm2go/p2.Fn3574
func Fn3574(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3577 github.com/goccy/pythonwasm2go/p0.Fn3577
func Fn3577(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3587 github.com/goccy/pythonwasm2go/p2.Fn3587
func Fn3587(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3591 github.com/goccy/pythonwasm2go/p0.Fn3591
func Fn3591(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3593 github.com/goccy/pythonwasm2go/p2.Fn3593
func Fn3593(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3596 github.com/goccy/pythonwasm2go/p0.Fn3596
func Fn3596(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3597 github.com/goccy/pythonwasm2go/p2.Fn3597
func Fn3597(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3598 github.com/goccy/pythonwasm2go/p2.Fn3598
func Fn3598(m *base.Module, l0 int32)

//go:linkname Fn3599 github.com/goccy/pythonwasm2go/p2.Fn3599
func Fn3599(m *base.Module, l0 int32) int32

//go:linkname Fn3600 github.com/goccy/pythonwasm2go/p0.Fn3600
func Fn3600(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3601 github.com/goccy/pythonwasm2go/p0.Fn3601
func Fn3601(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3602 github.com/goccy/pythonwasm2go/p2.Fn3602
func Fn3602(m *base.Module, l0 int32)

//go:linkname Fn3604 github.com/goccy/pythonwasm2go/p2.Fn3604
func Fn3604(m *base.Module, l0 int32) int32

//go:linkname Fn3605 github.com/goccy/pythonwasm2go/p0.Fn3605
func Fn3605(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3606 github.com/goccy/pythonwasm2go/p0.Fn3606
func Fn3606(m *base.Module, l0 int32)

//go:linkname Fn3607 github.com/goccy/pythonwasm2go/p2.Fn3607
func Fn3607(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3608 github.com/goccy/pythonwasm2go/p2.Fn3608
func Fn3608(m *base.Module, l0 int32)

//go:linkname Fn3609 github.com/goccy/pythonwasm2go/p0.Fn3609
func Fn3609(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3610 github.com/goccy/pythonwasm2go/p0.Fn3610
func Fn3610(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3611 github.com/goccy/pythonwasm2go/p2.Fn3611
func Fn3611(m *base.Module) int32

//go:linkname Fn3612 github.com/goccy/pythonwasm2go/p0.Fn3612
func Fn3612(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3613 github.com/goccy/pythonwasm2go/p0.Fn3613
func Fn3613(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3614 github.com/goccy/pythonwasm2go/p0.Fn3614
func Fn3614(m *base.Module, l0 int32) int32

//go:linkname Fn3617 github.com/goccy/pythonwasm2go/p2.Fn3617
func Fn3617(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3618 github.com/goccy/pythonwasm2go/p2.Fn3618
func Fn3618(m *base.Module) int32

//go:linkname Fn3619 github.com/goccy/pythonwasm2go/p2.Fn3619
func Fn3619(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3620 github.com/goccy/pythonwasm2go/p2.Fn3620
func Fn3620(m *base.Module)

//go:linkname Fn3622 github.com/goccy/pythonwasm2go/p2.Fn3622
func Fn3622(m *base.Module, l0 int32)

//go:linkname Fn3625 github.com/goccy/pythonwasm2go/p2.Fn3625
func Fn3625(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3626 github.com/goccy/pythonwasm2go/p0.Fn3626
func Fn3626(m *base.Module)

//go:linkname Fn3627 github.com/goccy/pythonwasm2go/p0.Fn3627
func Fn3627(m *base.Module) int32

//go:linkname Fn3628 github.com/goccy/pythonwasm2go/p2.Fn3628
func Fn3628(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3629 github.com/goccy/pythonwasm2go/p0.Fn3629
func Fn3629(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3630 github.com/goccy/pythonwasm2go/p2.Fn3630
func Fn3630(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3631 github.com/goccy/pythonwasm2go/p0.Fn3631
func Fn3631(m *base.Module, l0 int32) int32

//go:linkname Fn3633 github.com/goccy/pythonwasm2go/p2.Fn3633
func Fn3633(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3634 github.com/goccy/pythonwasm2go/p2.Fn3634
func Fn3634(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3635 github.com/goccy/pythonwasm2go/p0.Fn3635
func Fn3635(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3636 github.com/goccy/pythonwasm2go/p0.Fn3636
func Fn3636(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3639 github.com/goccy/pythonwasm2go/p0.Fn3639
func Fn3639(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3641 github.com/goccy/pythonwasm2go/p2.Fn3641
func Fn3641(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3642 github.com/goccy/pythonwasm2go/p2.Fn3642
func Fn3642(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn3646 github.com/goccy/pythonwasm2go/p2.Fn3646
func Fn3646(m *base.Module, l0 int32) int32

//go:linkname Fn3647 github.com/goccy/pythonwasm2go/p2.Fn3647
func Fn3647(m *base.Module, l0 int32)

//go:linkname Fn3648 github.com/goccy/pythonwasm2go/p2.Fn3648
func Fn3648(m *base.Module, l0 int32) int32

//go:linkname Fn3649 github.com/goccy/pythonwasm2go/p2.Fn3649
func Fn3649(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3651 github.com/goccy/pythonwasm2go/p0.Fn3651
func Fn3651(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3652 github.com/goccy/pythonwasm2go/p2.Fn3652
func Fn3652(m *base.Module, l0 int32) int32

//go:linkname Fn3656 github.com/goccy/pythonwasm2go/p2.Fn3656
func Fn3656(m *base.Module, l0 int32) int32

//go:linkname Fn3657 github.com/goccy/pythonwasm2go/p2.Fn3657
func Fn3657(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3660 github.com/goccy/pythonwasm2go/p2.Fn3660
func Fn3660(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3665 github.com/goccy/pythonwasm2go/p2.Fn3665
func Fn3665(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3666 github.com/goccy/pythonwasm2go/p2.Fn3666
func Fn3666(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3667 github.com/goccy/pythonwasm2go/p2.Fn3667
func Fn3667(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3673 github.com/goccy/pythonwasm2go/p2.Fn3673
func Fn3673(m *base.Module, l0 int32) int32

//go:linkname Fn3674 github.com/goccy/pythonwasm2go/p2.Fn3674
func Fn3674(m *base.Module, l0 int32, l1 int64, l2 int32)

//go:linkname Fn3677 github.com/goccy/pythonwasm2go/p2.Fn3677
func Fn3677(m *base.Module, l0 int32) int32

//go:linkname Fn3678 github.com/goccy/pythonwasm2go/p2.Fn3678
func Fn3678(m *base.Module, l0 int32) int32

//go:linkname Fn3679 github.com/goccy/pythonwasm2go/p2.Fn3679
func Fn3679(m *base.Module, l0 int32) int32

//go:linkname Fn3680 github.com/goccy/pythonwasm2go/p2.Fn3680
func Fn3680(m *base.Module, l0 int32) int32

//go:linkname Fn3683 github.com/goccy/pythonwasm2go/p0.Fn3683
func Fn3683(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3687 github.com/goccy/pythonwasm2go/p0.Fn3687
func Fn3687(m *base.Module, l0 int32) int32

//go:linkname Fn3688 github.com/goccy/pythonwasm2go/p2.Fn3688
func Fn3688(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3689 github.com/goccy/pythonwasm2go/p2.Fn3689
func Fn3689(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3694 github.com/goccy/pythonwasm2go/p0.Fn3694
func Fn3694(m *base.Module, l0 int32) int32

//go:linkname Fn3703 github.com/goccy/pythonwasm2go/p2.Fn3703
func Fn3703(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3716 github.com/goccy/pythonwasm2go/p0.Fn3716
func Fn3716(m *base.Module, l0 int32)

//go:linkname Fn3717 github.com/goccy/pythonwasm2go/p2.Fn3717
func Fn3717(m *base.Module, l0 int32)

//go:linkname Fn3719 github.com/goccy/pythonwasm2go/p0.Fn3719
func Fn3719(m *base.Module, l0 int32)

//go:linkname Fn3720 github.com/goccy/pythonwasm2go/p0.Fn3720
func Fn3720(m *base.Module, l0 int32) int32

//go:linkname Fn3721 github.com/goccy/pythonwasm2go/p2.Fn3721
func Fn3721(m *base.Module, l0 int32) int32

//go:linkname Fn3723 github.com/goccy/pythonwasm2go/p0.Fn3723
func Fn3723(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3724 github.com/goccy/pythonwasm2go/p2.Fn3724
func Fn3724(m *base.Module, l0 int32)

//go:linkname Fn3726 github.com/goccy/pythonwasm2go/p2.Fn3726
func Fn3726(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3728 github.com/goccy/pythonwasm2go/p0.Fn3728
func Fn3728(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3729 github.com/goccy/pythonwasm2go/p2.Fn3729
func Fn3729(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3730 github.com/goccy/pythonwasm2go/p2.Fn3730
func Fn3730(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3731 github.com/goccy/pythonwasm2go/p0.Fn3731
func Fn3731(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3732 github.com/goccy/pythonwasm2go/p0.Fn3732
func Fn3732(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn3733 github.com/goccy/pythonwasm2go/p0.Fn3733
func Fn3733(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3734 github.com/goccy/pythonwasm2go/p0.Fn3734
func Fn3734(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3735 github.com/goccy/pythonwasm2go/p2.Fn3735
func Fn3735(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3739 github.com/goccy/pythonwasm2go/p2.Fn3739
func Fn3739(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3740 github.com/goccy/pythonwasm2go/p2.Fn3740
func Fn3740(m *base.Module, l0 int32) int32

//go:linkname Fn3741 github.com/goccy/pythonwasm2go/p2.Fn3741
func Fn3741(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3747 github.com/goccy/pythonwasm2go/p2.Fn3747
func Fn3747(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3749 github.com/goccy/pythonwasm2go/p2.Fn3749
func Fn3749(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3750 github.com/goccy/pythonwasm2go/p2.Fn3750
func Fn3750(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3751 github.com/goccy/pythonwasm2go/p2.Fn3751
func Fn3751(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3759 github.com/goccy/pythonwasm2go/p0.Fn3759
func Fn3759(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3763 github.com/goccy/pythonwasm2go/p0.Fn3763
func Fn3763(m *base.Module, l0 int32)

//go:linkname Fn3764 github.com/goccy/pythonwasm2go/p0.Fn3764
func Fn3764(m *base.Module) int32

//go:linkname Fn3765 github.com/goccy/pythonwasm2go/p0.Fn3765
func Fn3765(m *base.Module, l0 int32)

//go:linkname Fn3767 github.com/goccy/pythonwasm2go/p2.Fn3767
func Fn3767(m *base.Module, l0 int32) int32

//go:linkname Fn3769 github.com/goccy/pythonwasm2go/p0.Fn3769
func Fn3769(m *base.Module, l0 int32) int32

//go:linkname Fn3771 github.com/goccy/pythonwasm2go/p0.Fn3771
func Fn3771(m *base.Module, l0 int32) int32

//go:linkname Fn3772 github.com/goccy/pythonwasm2go/p0.Fn3772
func Fn3772(m *base.Module) int32

//go:linkname Fn3783 github.com/goccy/pythonwasm2go/p2.Fn3783
func Fn3783(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3787 github.com/goccy/pythonwasm2go/p0.Fn3787
func Fn3787(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3788 github.com/goccy/pythonwasm2go/p0.Fn3788
func Fn3788(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3790 github.com/goccy/pythonwasm2go/p2.Fn3790
func Fn3790(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3823 github.com/goccy/pythonwasm2go/p2.Fn3823
func Fn3823(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3825 github.com/goccy/pythonwasm2go/p2.Fn3825
func Fn3825(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3826 github.com/goccy/pythonwasm2go/p2.Fn3826
func Fn3826(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3827 github.com/goccy/pythonwasm2go/p2.Fn3827
func Fn3827(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3828 github.com/goccy/pythonwasm2go/p2.Fn3828
func Fn3828(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3830 github.com/goccy/pythonwasm2go/p2.Fn3830
func Fn3830(m *base.Module, l0 int32) int32

//go:linkname Fn3831 github.com/goccy/pythonwasm2go/p2.Fn3831
func Fn3831(m *base.Module, l0 int32)

//go:linkname Fn3833 github.com/goccy/pythonwasm2go/p2.Fn3833
func Fn3833(m *base.Module, l0 int32)

//go:linkname Fn3835 github.com/goccy/pythonwasm2go/p0.Fn3835
func Fn3835(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3836 github.com/goccy/pythonwasm2go/p2.Fn3836
func Fn3836(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3837 github.com/goccy/pythonwasm2go/p0.Fn3837
func Fn3837(m *base.Module, l0 int32) int32

//go:linkname Fn3838 github.com/goccy/pythonwasm2go/p0.Fn3838
func Fn3838(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3842 github.com/goccy/pythonwasm2go/p2.Fn3842
func Fn3842(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3843 github.com/goccy/pythonwasm2go/p2.Fn3843
func Fn3843(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3844 github.com/goccy/pythonwasm2go/p2.Fn3844
func Fn3844(m *base.Module, l0 int32) int32

//go:linkname Fn3846 github.com/goccy/pythonwasm2go/p2.Fn3846
func Fn3846(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3847 github.com/goccy/pythonwasm2go/p2.Fn3847
func Fn3847(m *base.Module, l0 int32) int32

//go:linkname Fn3848 github.com/goccy/pythonwasm2go/p2.Fn3848
func Fn3848(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3849 github.com/goccy/pythonwasm2go/p2.Fn3849
func Fn3849(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3850 github.com/goccy/pythonwasm2go/p2.Fn3850
func Fn3850(m *base.Module, l0 int32) int32

//go:linkname Fn3853 github.com/goccy/pythonwasm2go/p2.Fn3853
func Fn3853(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3854 github.com/goccy/pythonwasm2go/p2.Fn3854
func Fn3854(m *base.Module)

//go:linkname Fn3855 github.com/goccy/pythonwasm2go/p2.Fn3855
func Fn3855(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3859 github.com/goccy/pythonwasm2go/p2.Fn3859
func Fn3859(m *base.Module, l0 int32)

//go:linkname Fn3860 github.com/goccy/pythonwasm2go/p2.Fn3860
func Fn3860(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3863 github.com/goccy/pythonwasm2go/p2.Fn3863
func Fn3863(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3864 github.com/goccy/pythonwasm2go/p2.Fn3864
func Fn3864(m *base.Module, l0 int32) int32

//go:linkname Fn3865 github.com/goccy/pythonwasm2go/p2.Fn3865
func Fn3865(m *base.Module) int32

//go:linkname Fn3866 github.com/goccy/pythonwasm2go/p0.Fn3866
func Fn3866(m *base.Module, l0 int32) int32

//go:linkname Fn3867 github.com/goccy/pythonwasm2go/p0.Fn3867
func Fn3867(m *base.Module, l0 int32) int32

//go:linkname Fn3873 github.com/goccy/pythonwasm2go/p0.Fn3873
func Fn3873(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3877 github.com/goccy/pythonwasm2go/p2.Fn3877
func Fn3877(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3900 github.com/goccy/pythonwasm2go/p2.Fn3900
func Fn3900(m *base.Module, l0 int32)

//go:linkname Fn3901 github.com/goccy/pythonwasm2go/p2.Fn3901
func Fn3901(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3905 github.com/goccy/pythonwasm2go/p2.Fn3905
func Fn3905(m *base.Module, l0 int32) int32

//go:linkname Fn3908 github.com/goccy/pythonwasm2go/p2.Fn3908
func Fn3908(m *base.Module, l0 int32) int32

//go:linkname Fn3909 github.com/goccy/pythonwasm2go/p2.Fn3909
func Fn3909(m *base.Module, l0 int32)

//go:linkname Fn3910 github.com/goccy/pythonwasm2go/p2.Fn3910
func Fn3910(m *base.Module, l0 int32)

//go:linkname Fn3913 github.com/goccy/pythonwasm2go/p2.Fn3913
func Fn3913(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3914 github.com/goccy/pythonwasm2go/p2.Fn3914
func Fn3914(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3916 github.com/goccy/pythonwasm2go/p2.Fn3916
func Fn3916(m *base.Module, l0 int32) int32

//go:linkname Fn3918 github.com/goccy/pythonwasm2go/p2.Fn3918
func Fn3918(m *base.Module, l0 int32)

//go:linkname Fn3922 github.com/goccy/pythonwasm2go/p2.Fn3922
func Fn3922(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3928 github.com/goccy/pythonwasm2go/p2.Fn3928
func Fn3928(m *base.Module, l0 int32)

//go:linkname Fn3944 github.com/goccy/pythonwasm2go/p2.Fn3944
func Fn3944(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3945 github.com/goccy/pythonwasm2go/p2.Fn3945
func Fn3945(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3947 github.com/goccy/pythonwasm2go/p2.Fn3947
func Fn3947(m *base.Module, l0 int32) int32

//go:linkname Fn3948 github.com/goccy/pythonwasm2go/p2.Fn3948
func Fn3948(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3949 github.com/goccy/pythonwasm2go/p2.Fn3949
func Fn3949(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3967 github.com/goccy/pythonwasm2go/p2.Fn3967
func Fn3967(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3974 github.com/goccy/pythonwasm2go/p2.Fn3974
func Fn3974(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3976 github.com/goccy/pythonwasm2go/p2.Fn3976
func Fn3976(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3977 github.com/goccy/pythonwasm2go/p2.Fn3977
func Fn3977(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3980 github.com/goccy/pythonwasm2go/p2.Fn3980
func Fn3980(m *base.Module, l0 int32) int32

//go:linkname Fn3994 github.com/goccy/pythonwasm2go/p2.Fn3994
func Fn3994(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3995 github.com/goccy/pythonwasm2go/p2.Fn3995
func Fn3995(m *base.Module, l0 int32) int32

//go:linkname Fn3996 github.com/goccy/pythonwasm2go/p2.Fn3996
func Fn3996(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3997 github.com/goccy/pythonwasm2go/p2.Fn3997
func Fn3997(m *base.Module, l0 int32) int32

//go:linkname Fn3998 github.com/goccy/pythonwasm2go/p2.Fn3998
func Fn3998(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3999 github.com/goccy/pythonwasm2go/p2.Fn3999
func Fn3999(m *base.Module, l0 int32) int32

//go:linkname Fn4001 github.com/goccy/pythonwasm2go/p2.Fn4001
func Fn4001(m *base.Module, l0 int32)

//go:linkname Fn4002 github.com/goccy/pythonwasm2go/p2.Fn4002
func Fn4002(m *base.Module) int32

//go:linkname Fn4028 github.com/goccy/pythonwasm2go/p2.Fn4028
func Fn4028(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4038 github.com/goccy/pythonwasm2go/p2.Fn4038
func Fn4038(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4039 github.com/goccy/pythonwasm2go/p2.Fn4039
func Fn4039(m *base.Module, l0 int32) int32

//go:linkname Fn4040 github.com/goccy/pythonwasm2go/p2.Fn4040
func Fn4040(m *base.Module, l0 int32) int32

//go:linkname Fn4055 github.com/goccy/pythonwasm2go/p0.Fn4055
func Fn4055(m *base.Module, l0 int32)

//go:linkname Fn4056 github.com/goccy/pythonwasm2go/p2.Fn4056
func Fn4056(m *base.Module, l0 int32)

//go:linkname Fn4059 github.com/goccy/pythonwasm2go/p2.Fn4059
func Fn4059(m *base.Module, l0 int32) int32

//go:linkname Fn4060 github.com/goccy/pythonwasm2go/p0.Fn4060
func Fn4060(m *base.Module, l0 int32)

//go:linkname Fn4061 github.com/goccy/pythonwasm2go/p0.Fn4061
func Fn4061(m *base.Module, l0 int32)

//go:linkname Fn4062 github.com/goccy/pythonwasm2go/p2.Fn4062
func Fn4062(m *base.Module, l0 int32)

//go:linkname Fn4063 github.com/goccy/pythonwasm2go/p0.Fn4063
func Fn4063(m *base.Module, l0 int32) int32

//go:linkname Fn4064 github.com/goccy/pythonwasm2go/p0.Fn4064
func Fn4064(m *base.Module, l0 int32)

//go:linkname Fn4068 github.com/goccy/pythonwasm2go/p2.Fn4068
func Fn4068(m *base.Module, l0 int32)

//go:linkname Fn4069 github.com/goccy/pythonwasm2go/p2.Fn4069
func Fn4069(m *base.Module, l0 int32)

//go:linkname Fn4071 github.com/goccy/pythonwasm2go/p2.Fn4071
func Fn4071(m *base.Module, l0 int32)

//go:linkname Fn4073 github.com/goccy/pythonwasm2go/p2.Fn4073
func Fn4073(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4074 github.com/goccy/pythonwasm2go/p0.Fn4074
func Fn4074(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4075 github.com/goccy/pythonwasm2go/p2.Fn4075
func Fn4075(m *base.Module, l0 int32)

//go:linkname Fn4082 github.com/goccy/pythonwasm2go/p0.Fn4082
func Fn4082(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4090 github.com/goccy/pythonwasm2go/p2.Fn4090
func Fn4090(m *base.Module, l0 int32) int32

//go:linkname Fn4105 github.com/goccy/pythonwasm2go/p2.Fn4105
func Fn4105(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4106 github.com/goccy/pythonwasm2go/p0.Fn4106
func Fn4106(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4114 github.com/goccy/pythonwasm2go/p2.Fn4114
func Fn4114(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4115 github.com/goccy/pythonwasm2go/p2.Fn4115
func Fn4115(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4117 github.com/goccy/pythonwasm2go/p2.Fn4117
func Fn4117(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4118 github.com/goccy/pythonwasm2go/p2.Fn4118
func Fn4118(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4119 github.com/goccy/pythonwasm2go/p2.Fn4119
func Fn4119(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4121 github.com/goccy/pythonwasm2go/p2.Fn4121
func Fn4121(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4127 github.com/goccy/pythonwasm2go/p0.Fn4127
func Fn4127(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn4135 github.com/goccy/pythonwasm2go/p2.Fn4135
func Fn4135(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4136 github.com/goccy/pythonwasm2go/p2.Fn4136
func Fn4136(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4138 github.com/goccy/pythonwasm2go/p2.Fn4138
func Fn4138(m *base.Module, l0 int32)

//go:linkname Fn4144 github.com/goccy/pythonwasm2go/p2.Fn4144
func Fn4144(m *base.Module) int32

//go:linkname Fn4147 github.com/goccy/pythonwasm2go/p2.Fn4147
func Fn4147(m *base.Module, l0 int32)

//go:linkname Fn4149 github.com/goccy/pythonwasm2go/p2.Fn4149
func Fn4149(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4151 github.com/goccy/pythonwasm2go/p2.Fn4151
func Fn4151(m *base.Module, l0 int32) int32

//go:linkname Fn4152 github.com/goccy/pythonwasm2go/p2.Fn4152
func Fn4152(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4157 github.com/goccy/pythonwasm2go/p2.Fn4157
func Fn4157(m *base.Module, l0 int32) int32

//go:linkname Fn4158 github.com/goccy/pythonwasm2go/p2.Fn4158
func Fn4158(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4171 github.com/goccy/pythonwasm2go/p0.Fn4171
func Fn4171(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4173 github.com/goccy/pythonwasm2go/p0.Fn4173
func Fn4173(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4174 github.com/goccy/pythonwasm2go/p2.Fn4174
func Fn4174(m *base.Module, l0 int32)

//go:linkname Fn4178 github.com/goccy/pythonwasm2go/p0.Fn4178
func Fn4178(m *base.Module, l0 int32)

//go:linkname Fn4179 github.com/goccy/pythonwasm2go/p0.Fn4179
func Fn4179(m *base.Module, l0 int32)

//go:linkname Fn4184 github.com/goccy/pythonwasm2go/p0.Fn4184
func Fn4184(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4185 github.com/goccy/pythonwasm2go/p2.Fn4185
func Fn4185(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4187 github.com/goccy/pythonwasm2go/p2.Fn4187
func Fn4187(m *base.Module) int32

//go:linkname Fn4188 github.com/goccy/pythonwasm2go/p2.Fn4188
func Fn4188(m *base.Module, l0 int32) int32

//go:linkname Fn4189 github.com/goccy/pythonwasm2go/p0.Fn4189
func Fn4189(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4192 github.com/goccy/pythonwasm2go/p2.Fn4192
func Fn4192(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4196 github.com/goccy/pythonwasm2go/p2.Fn4196
func Fn4196(m *base.Module)

//go:linkname Fn4198 github.com/goccy/pythonwasm2go/p2.Fn4198
func Fn4198(m *base.Module, l0 int32)

//go:linkname Fn4199 github.com/goccy/pythonwasm2go/p0.Fn4199
func Fn4199(m *base.Module, l0 int32)

//go:linkname Fn4200 github.com/goccy/pythonwasm2go/p0.Fn4200
func Fn4200(m *base.Module, l0 int32)

//go:linkname Fn4201 github.com/goccy/pythonwasm2go/p0.Fn4201
func Fn4201(m *base.Module, l0 int32)

//go:linkname Fn4204 github.com/goccy/pythonwasm2go/p0.Fn4204
func Fn4204(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4205 github.com/goccy/pythonwasm2go/p2.Fn4205
func Fn4205(m *base.Module, l0 int32)

//go:linkname Fn4206 github.com/goccy/pythonwasm2go/p2.Fn4206
func Fn4206(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4207 github.com/goccy/pythonwasm2go/p2.Fn4207
func Fn4207(m *base.Module, l0 int32) int32

//go:linkname Fn4209 github.com/goccy/pythonwasm2go/p2.Fn4209
func Fn4209(m *base.Module, l0 int32) int32

//go:linkname Fn4211 github.com/goccy/pythonwasm2go/p2.Fn4211
func Fn4211(m *base.Module, l0 int32) int64

//go:linkname Fn4212 github.com/goccy/pythonwasm2go/p0.Fn4212
func Fn4212(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4213 github.com/goccy/pythonwasm2go/p0.Fn4213
func Fn4213(m *base.Module, l0 int32) int32

//go:linkname Fn4217 github.com/goccy/pythonwasm2go/p0.Fn4217
func Fn4217(m *base.Module, l0 int32)

//go:linkname Fn4219 github.com/goccy/pythonwasm2go/p2.Fn4219
func Fn4219(m *base.Module) int32

//go:linkname Fn4221 github.com/goccy/pythonwasm2go/p2.Fn4221
func Fn4221(m *base.Module, l0 int32) int32

//go:linkname Fn4222 github.com/goccy/pythonwasm2go/p2.Fn4222
func Fn4222(m *base.Module, l0 int32) int32

//go:linkname Fn4223 github.com/goccy/pythonwasm2go/p0.Fn4223
func Fn4223(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4224 github.com/goccy/pythonwasm2go/p2.Fn4224
func Fn4224(m *base.Module, l0 int32) int32

//go:linkname Fn4225 github.com/goccy/pythonwasm2go/p2.Fn4225
func Fn4225(m *base.Module, l0 int32)

//go:linkname Fn4226 github.com/goccy/pythonwasm2go/p0.Fn4226
func Fn4226(m *base.Module, l0 int32)

//go:linkname Fn4227 github.com/goccy/pythonwasm2go/p0.Fn4227
func Fn4227(m *base.Module, l0 int32)

//go:linkname Fn4232 github.com/goccy/pythonwasm2go/p2.Fn4232
func Fn4232(m *base.Module) int32

//go:linkname Fn4233 github.com/goccy/pythonwasm2go/p0.Fn4233
func Fn4233(m *base.Module, l0 int32) int32

//go:linkname Fn4235 github.com/goccy/pythonwasm2go/p2.Fn4235
func Fn4235(m *base.Module) int32

//go:linkname Fn4236 github.com/goccy/pythonwasm2go/p0.Fn4236
func Fn4236(m *base.Module, l0 int32)

//go:linkname Fn4241 github.com/goccy/pythonwasm2go/p0.Fn4241
func Fn4241(m *base.Module) int32

//go:linkname Fn4242 github.com/goccy/pythonwasm2go/p2.Fn4242
func Fn4242(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4244 github.com/goccy/pythonwasm2go/p2.Fn4244
func Fn4244(m *base.Module, l0 int32) int32

//go:linkname Fn4248 github.com/goccy/pythonwasm2go/p2.Fn4248
func Fn4248(m *base.Module, l0 int32)

//go:linkname Fn4249 github.com/goccy/pythonwasm2go/p0.Fn4249
func Fn4249(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4250 github.com/goccy/pythonwasm2go/p2.Fn4250
func Fn4250(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)

//go:linkname Fn4252 github.com/goccy/pythonwasm2go/p0.Fn4252
func Fn4252(m *base.Module, l0 int32) int32

//go:linkname Fn4253 github.com/goccy/pythonwasm2go/p2.Fn4253
func Fn4253(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4254 github.com/goccy/pythonwasm2go/p2.Fn4254
func Fn4254(m *base.Module, l0 int32)

//go:linkname Fn4257 github.com/goccy/pythonwasm2go/p2.Fn4257
func Fn4257(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4258 github.com/goccy/pythonwasm2go/p2.Fn4258
func Fn4258(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4264 github.com/goccy/pythonwasm2go/p2.Fn4264
func Fn4264(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4265 github.com/goccy/pythonwasm2go/p2.Fn4265
func Fn4265(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4266 github.com/goccy/pythonwasm2go/p2.Fn4266
func Fn4266(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4269 github.com/goccy/pythonwasm2go/p2.Fn4269
func Fn4269(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4270 github.com/goccy/pythonwasm2go/p2.Fn4270
func Fn4270(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4272 github.com/goccy/pythonwasm2go/p2.Fn4272
func Fn4272(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4273 github.com/goccy/pythonwasm2go/p2.Fn4273
func Fn4273(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4276 github.com/goccy/pythonwasm2go/p2.Fn4276
func Fn4276(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4290 github.com/goccy/pythonwasm2go/p2.Fn4290
func Fn4290(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4293 github.com/goccy/pythonwasm2go/p0.Fn4293
func Fn4293(m *base.Module, l0 int32) int32

//go:linkname Fn4299 github.com/goccy/pythonwasm2go/p0.Fn4299
func Fn4299(m *base.Module, l0 int32) int32

//go:linkname Fn4304 github.com/goccy/pythonwasm2go/p0.Fn4304
func Fn4304(m *base.Module, l0 int32) int32

//go:linkname Fn4315 github.com/goccy/pythonwasm2go/p2.Fn4315
func Fn4315(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4329 github.com/goccy/pythonwasm2go/p0.Fn4329
func Fn4329(m *base.Module, l0 int32) int32

//go:linkname Fn4378 github.com/goccy/pythonwasm2go/p0.Fn4378
func Fn4378(m *base.Module, l0 int32) int32

//go:linkname Fn4382 github.com/goccy/pythonwasm2go/p0.Fn4382
func Fn4382(m *base.Module, l0 int32) int32

//go:linkname Fn4383 github.com/goccy/pythonwasm2go/p0.Fn4383
func Fn4383(m *base.Module, l0 int32) int32

//go:linkname Fn4386 github.com/goccy/pythonwasm2go/p2.Fn4386
func Fn4386(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4390 github.com/goccy/pythonwasm2go/p0.Fn4390
func Fn4390(m *base.Module, l0 int32) int32

//go:linkname Fn4428 github.com/goccy/pythonwasm2go/p0.Fn4428
func Fn4428(m *base.Module, l0 int32) int32

//go:linkname Fn4436 github.com/goccy/pythonwasm2go/p2.Fn4436
func Fn4436(m *base.Module, l0 int32) int32

//go:linkname Fn4450 github.com/goccy/pythonwasm2go/p0.Fn4450
func Fn4450(m *base.Module, l0 int32) int32

//go:linkname Fn4454 github.com/goccy/pythonwasm2go/p2.Fn4454
func Fn4454(m *base.Module, l0 int32) int32

//go:linkname Fn4470 github.com/goccy/pythonwasm2go/p2.Fn4470
func Fn4470(m *base.Module, l0 int32) int32

//go:linkname Fn4475 github.com/goccy/pythonwasm2go/p2.Fn4475
func Fn4475(m *base.Module, l0 int32) int32

//go:linkname Fn4476 github.com/goccy/pythonwasm2go/p2.Fn4476
func Fn4476(m *base.Module, l0 int32) int32

//go:linkname Fn4479 github.com/goccy/pythonwasm2go/p0.Fn4479
func Fn4479(m *base.Module, l0 int32) int32

//go:linkname Fn4480 github.com/goccy/pythonwasm2go/p0.Fn4480
func Fn4480(m *base.Module, l0 int32) int32

//go:linkname Fn4498 github.com/goccy/pythonwasm2go/p2.Fn4498
func Fn4498(m *base.Module, l0 int32) int32

//go:linkname Fn4499 github.com/goccy/pythonwasm2go/p2.Fn4499
func Fn4499(m *base.Module, l0 int32) int32

//go:linkname Fn4507 github.com/goccy/pythonwasm2go/p2.Fn4507
func Fn4507(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4508 github.com/goccy/pythonwasm2go/p2.Fn4508
func Fn4508(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4509 github.com/goccy/pythonwasm2go/p2.Fn4509
func Fn4509(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4511 github.com/goccy/pythonwasm2go/p2.Fn4511
func Fn4511(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4512 github.com/goccy/pythonwasm2go/p0.Fn4512
func Fn4512(m *base.Module, l0 int32) int32

//go:linkname Fn4513 github.com/goccy/pythonwasm2go/p2.Fn4513
func Fn4513(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4514 github.com/goccy/pythonwasm2go/p2.Fn4514
func Fn4514(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4515 github.com/goccy/pythonwasm2go/p2.Fn4515
func Fn4515(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4516 github.com/goccy/pythonwasm2go/p2.Fn4516
func Fn4516(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4518 github.com/goccy/pythonwasm2go/p2.Fn4518
func Fn4518(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4519 github.com/goccy/pythonwasm2go/p2.Fn4519
func Fn4519(m *base.Module, l0 int32) int32

//go:linkname Fn4520 github.com/goccy/pythonwasm2go/p2.Fn4520
func Fn4520(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4521 github.com/goccy/pythonwasm2go/p2.Fn4521
func Fn4521(m *base.Module, l0 int32) int32

//go:linkname Fn4523 github.com/goccy/pythonwasm2go/p2.Fn4523
func Fn4523(m *base.Module, l0 int32) int32

//go:linkname Fn4527 github.com/goccy/pythonwasm2go/p2.Fn4527
func Fn4527(m *base.Module, l0 int32)

//go:linkname Fn4530 github.com/goccy/pythonwasm2go/p2.Fn4530
func Fn4530(m *base.Module, l0 int32) int32

//go:linkname Fn4532 github.com/goccy/pythonwasm2go/p2.Fn4532
func Fn4532(m *base.Module, l0 int32)

//go:linkname Fn4535 github.com/goccy/pythonwasm2go/p2.Fn4535
func Fn4535(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4536 github.com/goccy/pythonwasm2go/p2.Fn4536
func Fn4536(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn4539 github.com/goccy/pythonwasm2go/p2.Fn4539
func Fn4539(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4540 github.com/goccy/pythonwasm2go/p2.Fn4540
func Fn4540(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4545 github.com/goccy/pythonwasm2go/p0.Fn4545
func Fn4545(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4550 github.com/goccy/pythonwasm2go/p2.Fn4550
func Fn4550(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4551 github.com/goccy/pythonwasm2go/p2.Fn4551
func Fn4551(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4556 github.com/goccy/pythonwasm2go/p2.Fn4556
func Fn4556(m *base.Module, l0 int32, l1 int64) int32

//go:linkname Fn4557 github.com/goccy/pythonwasm2go/p2.Fn4557
func Fn4557(m *base.Module, l0 int32) int64

//go:linkname Fn4558 github.com/goccy/pythonwasm2go/p2.Fn4558
func Fn4558(m *base.Module)

//go:linkname Fn4560 github.com/goccy/pythonwasm2go/p2.Fn4560
func Fn4560(m *base.Module, l0 int32) int32

//go:linkname Fn4561 github.com/goccy/pythonwasm2go/p2.Fn4561
func Fn4561(m *base.Module, l0 float64, l1 int32) float64

//go:linkname Fn4562 github.com/goccy/pythonwasm2go/p2.Fn4562
func Fn4562(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4564 github.com/goccy/pythonwasm2go/p2.Fn4564
func Fn4564(m *base.Module, l0 int32) int64

//go:linkname Fn4565 github.com/goccy/pythonwasm2go/p0.Fn4565
func Fn4565(m *base.Module)

//go:linkname Fn4567 github.com/goccy/pythonwasm2go/p2.Fn4567
func Fn4567(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4569 github.com/goccy/pythonwasm2go/p2.Fn4569
func Fn4569(m *base.Module, l0 int64) float64

//go:linkname Fn4570 github.com/goccy/pythonwasm2go/p2.Fn4570
func Fn4570(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn4572 github.com/goccy/pythonwasm2go/p2.Fn4572
func Fn4572(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn4573 github.com/goccy/pythonwasm2go/p2.Fn4573
func Fn4573(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn4576 github.com/goccy/pythonwasm2go/p2.Fn4576
func Fn4576(m *base.Module, l0 int32) int32

//go:linkname Fn4577 github.com/goccy/pythonwasm2go/p0.Fn4577
func Fn4577(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4579 github.com/goccy/pythonwasm2go/p2.Fn4579
func Fn4579(m *base.Module, l0 int32) int32

//go:linkname Fn4581 github.com/goccy/pythonwasm2go/p0.Fn4581
func Fn4581(m *base.Module, l0 int32) int32

//go:linkname Fn4582 github.com/goccy/pythonwasm2go/p2.Fn4582
func Fn4582(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4583 github.com/goccy/pythonwasm2go/p2.Fn4583
func Fn4583(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn4585 github.com/goccy/pythonwasm2go/p2.Fn4585
func Fn4585(m *base.Module, l0 int64) int64

//go:linkname Fn4586 github.com/goccy/pythonwasm2go/p0.Fn4586
func Fn4586(m *base.Module, l0 int64) int64

//go:linkname Fn4587 github.com/goccy/pythonwasm2go/p2.Fn4587
func Fn4587(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4590 github.com/goccy/pythonwasm2go/p2.Fn4590
func Fn4590(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4595 github.com/goccy/pythonwasm2go/p2.Fn4595
func Fn4595(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4621 github.com/goccy/pythonwasm2go/p2.Fn4621
func Fn4621(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4622 github.com/goccy/pythonwasm2go/p0.Fn4622
func Fn4622(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4623 github.com/goccy/pythonwasm2go/p0.Fn4623
func Fn4623(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4624 github.com/goccy/pythonwasm2go/p2.Fn4624
func Fn4624(m *base.Module, l0 int32) int32

//go:linkname Fn4625 github.com/goccy/pythonwasm2go/p0.Fn4625
func Fn4625(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn4626 github.com/goccy/pythonwasm2go/p2.Fn4626
func Fn4626(m *base.Module, l0 int32)

//go:linkname Fn4629 github.com/goccy/pythonwasm2go/p2.Fn4629
func Fn4629(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4650 github.com/goccy/pythonwasm2go/p2.Fn4650
func Fn4650(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4651 github.com/goccy/pythonwasm2go/p2.Fn4651
func Fn4651(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4653 github.com/goccy/pythonwasm2go/p2.Fn4653
func Fn4653(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4654 github.com/goccy/pythonwasm2go/p2.Fn4654
func Fn4654(m *base.Module, l0 int32) int32

//go:linkname Fn4655 github.com/goccy/pythonwasm2go/p2.Fn4655
func Fn4655(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4659 github.com/goccy/pythonwasm2go/p2.Fn4659
func Fn4659(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4660 github.com/goccy/pythonwasm2go/p2.Fn4660
func Fn4660(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4661 github.com/goccy/pythonwasm2go/p0.Fn4661
func Fn4661(m *base.Module, l0 int32) int32

//go:linkname Fn4662 github.com/goccy/pythonwasm2go/p2.Fn4662
func Fn4662(m *base.Module, l0 int32) int32

//go:linkname Fn4663 github.com/goccy/pythonwasm2go/p0.Fn4663
func Fn4663(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4669 github.com/goccy/pythonwasm2go/p0.Fn4669
func Fn4669(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4670 github.com/goccy/pythonwasm2go/p0.Fn4670
func Fn4670(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4673 github.com/goccy/pythonwasm2go/p0.Fn4673
func Fn4673(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4674 github.com/goccy/pythonwasm2go/p2.Fn4674
func Fn4674(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4677 github.com/goccy/pythonwasm2go/p2.Fn4677
func Fn4677(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4691 github.com/goccy/pythonwasm2go/p0.Fn4691
func Fn4691(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4743 github.com/goccy/pythonwasm2go/p2.Fn4743
func Fn4743(m *base.Module)

//go:linkname Fn4744 github.com/goccy/pythonwasm2go/p2.Fn4744
func Fn4744(m *base.Module) int32

//go:linkname Fn4746 github.com/goccy/pythonwasm2go/p2.Fn4746
func Fn4746(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4749 github.com/goccy/pythonwasm2go/p2.Fn4749
func Fn4749(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4750 github.com/goccy/pythonwasm2go/p2.Fn4750
func Fn4750(m *base.Module) int64

//go:linkname Fn4751 github.com/goccy/pythonwasm2go/p2.Fn4751
func Fn4751(m *base.Module) int32

//go:linkname Fn4752 github.com/goccy/pythonwasm2go/p2.Fn4752
func Fn4752(m *base.Module, l0 int32)

//go:linkname Fn4754 github.com/goccy/pythonwasm2go/p2.Fn4754
func Fn4754(m *base.Module, l0 int32)

//go:linkname Fn4755 github.com/goccy/pythonwasm2go/p0.Fn4755
func Fn4755(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4756 github.com/goccy/pythonwasm2go/p2.Fn4756
func Fn4756(m *base.Module, l0 int32) int32

//go:linkname Fn4757 github.com/goccy/pythonwasm2go/p2.Fn4757
func Fn4757(m *base.Module, l0 int32)

//go:linkname Fn4758 github.com/goccy/pythonwasm2go/p2.Fn4758
func Fn4758(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4759 github.com/goccy/pythonwasm2go/p2.Fn4759
func Fn4759(m *base.Module, l0 int32) int32

//go:linkname Fn4761 github.com/goccy/pythonwasm2go/p0.Fn4761
func Fn4761(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4763 github.com/goccy/pythonwasm2go/p2.Fn4763
func Fn4763(m *base.Module, l0 int32)

//go:linkname Fn4777 github.com/goccy/pythonwasm2go/p2.Fn4777
func Fn4777(m *base.Module, l0 int32)

//go:linkname Fn4784 github.com/goccy/pythonwasm2go/p2.Fn4784
func Fn4784(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4788 github.com/goccy/pythonwasm2go/p2.Fn4788
func Fn4788(m *base.Module) int32

//go:linkname Fn4789 github.com/goccy/pythonwasm2go/p2.Fn4789
func Fn4789(m *base.Module, l0 int32) int32

//go:linkname Fn4793 github.com/goccy/pythonwasm2go/p2.Fn4793
func Fn4793(m *base.Module) int32

//go:linkname Fn4809 github.com/goccy/pythonwasm2go/p2.Fn4809
func Fn4809(m *base.Module, l0 int32)

//go:linkname Fn4810 github.com/goccy/pythonwasm2go/p0.Fn4810
func Fn4810(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4818 github.com/goccy/pythonwasm2go/p2.Fn4818
func Fn4818(m *base.Module, l0 int32) int32

//go:linkname Fn4824 github.com/goccy/pythonwasm2go/p2.Fn4824
func Fn4824(m *base.Module)

//go:linkname Fn4826 github.com/goccy/pythonwasm2go/p2.Fn4826
func Fn4826(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4827 github.com/goccy/pythonwasm2go/p2.Fn4827
func Fn4827(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4828 github.com/goccy/pythonwasm2go/p2.Fn4828
func Fn4828(m *base.Module, l0 int32, l1 int32) float64

//go:linkname Fn4835 github.com/goccy/pythonwasm2go/p2.Fn4835
func Fn4835(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4836 github.com/goccy/pythonwasm2go/p0.Fn4836
func Fn4836(m *base.Module, l0 int32, l1 int32) float64

//go:linkname Fn4837 github.com/goccy/pythonwasm2go/p2.Fn4837
func Fn4837(m *base.Module, l0 int32) int32

//go:linkname Fn4843 github.com/goccy/pythonwasm2go/p2.Fn4843
func Fn4843(m *base.Module, l0 int32)

//go:linkname Fn4846 github.com/goccy/pythonwasm2go/p2.Fn4846
func Fn4846(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4850 github.com/goccy/pythonwasm2go/p2.Fn4850
func Fn4850(m *base.Module, l0 int32)

//go:linkname Fn4851 github.com/goccy/pythonwasm2go/p0.Fn4851
func Fn4851(m *base.Module, l0 float64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn4857 github.com/goccy/pythonwasm2go/p2.Fn4857
func Fn4857(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4859 github.com/goccy/pythonwasm2go/p2.Fn4859
func Fn4859(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn4860 github.com/goccy/pythonwasm2go/p2.Fn4860
func Fn4860(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn4861 github.com/goccy/pythonwasm2go/p2.Fn4861
func Fn4861(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4862 github.com/goccy/pythonwasm2go/p2.Fn4862
func Fn4862(m *base.Module, l0 int32) int32

//go:linkname Fn4863 github.com/goccy/pythonwasm2go/p2.Fn4863
func Fn4863(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4864 github.com/goccy/pythonwasm2go/p2.Fn4864
func Fn4864(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4866 github.com/goccy/pythonwasm2go/p2.Fn4866
func Fn4866(m *base.Module, l0 int32) int32

//go:linkname Fn4870 github.com/goccy/pythonwasm2go/p2.Fn4870
func Fn4870(m *base.Module, l0 int32)

//go:linkname Fn4872 github.com/goccy/pythonwasm2go/p2.Fn4872
func Fn4872(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn4875 github.com/goccy/pythonwasm2go/p2.Fn4875
func Fn4875(m *base.Module) int32

//go:linkname Fn4877 github.com/goccy/pythonwasm2go/p2.Fn4877
func Fn4877(m *base.Module) int32

//go:linkname Fn4878 github.com/goccy/pythonwasm2go/p2.Fn4878
func Fn4878(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4881 github.com/goccy/pythonwasm2go/p2.Fn4881
func Fn4881(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4883 github.com/goccy/pythonwasm2go/p2.Fn4883
func Fn4883(m *base.Module, l0 int32) int32

//go:linkname Fn4884 github.com/goccy/pythonwasm2go/p2.Fn4884
func Fn4884(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn4885 github.com/goccy/pythonwasm2go/p2.Fn4885
func Fn4885(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4886 github.com/goccy/pythonwasm2go/p2.Fn4886
func Fn4886(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4889 github.com/goccy/pythonwasm2go/p2.Fn4889
func Fn4889(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4890 github.com/goccy/pythonwasm2go/p2.Fn4890
func Fn4890(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4893 github.com/goccy/pythonwasm2go/p2.Fn4893
func Fn4893(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4896 github.com/goccy/pythonwasm2go/p0.Fn4896
func Fn4896(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4897 github.com/goccy/pythonwasm2go/p2.Fn4897
func Fn4897(m *base.Module, l0 int32) int32

//go:linkname Fn4899 github.com/goccy/pythonwasm2go/p2.Fn4899
func Fn4899(m *base.Module, l0 int32) int32

//go:linkname Fn4900 github.com/goccy/pythonwasm2go/p2.Fn4900
func Fn4900(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4902 github.com/goccy/pythonwasm2go/p2.Fn4902
func Fn4902(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4904 github.com/goccy/pythonwasm2go/p2.Fn4904
func Fn4904(m *base.Module, l0 int32) int32

//go:linkname Fn4905 github.com/goccy/pythonwasm2go/p2.Fn4905
func Fn4905(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4906 github.com/goccy/pythonwasm2go/p2.Fn4906
func Fn4906(m *base.Module, l0 int32) int32

//go:linkname Fn4909 github.com/goccy/pythonwasm2go/p2.Fn4909
func Fn4909(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4948 github.com/goccy/pythonwasm2go/p2.Fn4948
func Fn4948(m *base.Module, l0 int32) int32

//go:linkname Fn4949 github.com/goccy/pythonwasm2go/p2.Fn4949
func Fn4949(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4983 github.com/goccy/pythonwasm2go/p2.Fn4983
func Fn4983(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4984 github.com/goccy/pythonwasm2go/p2.Fn4984
func Fn4984(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4986 github.com/goccy/pythonwasm2go/p2.Fn4986
func Fn4986(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4989 github.com/goccy/pythonwasm2go/p2.Fn4989
func Fn4989(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4994 github.com/goccy/pythonwasm2go/p2.Fn4994
func Fn4994(m *base.Module, l0 int32) int32

//go:linkname Fn4999 github.com/goccy/pythonwasm2go/p2.Fn4999
func Fn4999(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5006 github.com/goccy/pythonwasm2go/p2.Fn5006
func Fn5006(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5024 github.com/goccy/pythonwasm2go/p2.Fn5024
func Fn5024(m *base.Module, l0 int32) int32

//go:linkname Fn5043 github.com/goccy/pythonwasm2go/p2.Fn5043
func Fn5043(m *base.Module) int32

//go:linkname Fn5045 github.com/goccy/pythonwasm2go/p2.Fn5045
func Fn5045(m *base.Module, l0 int32) int32

//go:linkname Fn5050 github.com/goccy/pythonwasm2go/p2.Fn5050
func Fn5050(m *base.Module, l0 int32)

//go:linkname Fn5053 github.com/goccy/pythonwasm2go/p2.Fn5053
func Fn5053(m *base.Module, l0 int32)

//go:linkname Fn5060 github.com/goccy/pythonwasm2go/p2.Fn5060
func Fn5060(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5062 github.com/goccy/pythonwasm2go/p2.Fn5062
func Fn5062(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5063 github.com/goccy/pythonwasm2go/p2.Fn5063
func Fn5063(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5067 github.com/goccy/pythonwasm2go/p2.Fn5067
func Fn5067(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5071 github.com/goccy/pythonwasm2go/p2.Fn5071
func Fn5071(m *base.Module, l0 int32) int32

//go:linkname Fn5072 github.com/goccy/pythonwasm2go/p2.Fn5072
func Fn5072(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5075 github.com/goccy/pythonwasm2go/p0.Fn5075
func Fn5075(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5081 github.com/goccy/pythonwasm2go/p2.Fn5081
func Fn5081(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5084 github.com/goccy/pythonwasm2go/p2.Fn5084
func Fn5084(m *base.Module, l0 int32) int32

//go:linkname Fn5085 github.com/goccy/pythonwasm2go/p2.Fn5085
func Fn5085(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5088 github.com/goccy/pythonwasm2go/p2.Fn5088
func Fn5088(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5131 github.com/goccy/pythonwasm2go/p2.Fn5131
func Fn5131(m *base.Module, l0 int32)

//go:linkname Fn5156 github.com/goccy/pythonwasm2go/p2.Fn5156
func Fn5156(m *base.Module, l0 int32) int32

//go:linkname Fn5160 github.com/goccy/pythonwasm2go/p2.Fn5160
func Fn5160(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5177 github.com/goccy/pythonwasm2go/p2.Fn5177
func Fn5177(m *base.Module, l0 int32) int32

//go:linkname Fn5182 github.com/goccy/pythonwasm2go/p2.Fn5182
func Fn5182(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5188 github.com/goccy/pythonwasm2go/p2.Fn5188
func Fn5188(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5195 github.com/goccy/pythonwasm2go/p2.Fn5195
func Fn5195(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5225 github.com/goccy/pythonwasm2go/p2.Fn5225
func Fn5225(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5237 github.com/goccy/pythonwasm2go/p0.Fn5237
func Fn5237(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5253 github.com/goccy/pythonwasm2go/p2.Fn5253
func Fn5253(m *base.Module, l0 int32)

//go:linkname Fn5258 github.com/goccy/pythonwasm2go/p2.Fn5258
func Fn5258(m *base.Module, l0 int32) int32

//go:linkname Fn5275 github.com/goccy/pythonwasm2go/p2.Fn5275
func Fn5275(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5276 github.com/goccy/pythonwasm2go/p2.Fn5276
func Fn5276(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5278 github.com/goccy/pythonwasm2go/p2.Fn5278
func Fn5278(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5285 github.com/goccy/pythonwasm2go/p2.Fn5285
func Fn5285(m *base.Module, l0 int32) int32

//go:linkname Fn5286 github.com/goccy/pythonwasm2go/p2.Fn5286
func Fn5286(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5287 github.com/goccy/pythonwasm2go/p2.Fn5287
func Fn5287(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5288 github.com/goccy/pythonwasm2go/p2.Fn5288
func Fn5288(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5290 github.com/goccy/pythonwasm2go/p2.Fn5290
func Fn5290(m *base.Module, l0 int32) int32

//go:linkname Fn5292 github.com/goccy/pythonwasm2go/p2.Fn5292
func Fn5292(m *base.Module, l0 int32) int32

//go:linkname Fn5294 github.com/goccy/pythonwasm2go/p2.Fn5294
func Fn5294(m *base.Module, l0 int32) int32

//go:linkname Fn5296 github.com/goccy/pythonwasm2go/p2.Fn5296
func Fn5296(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5297 github.com/goccy/pythonwasm2go/p2.Fn5297
func Fn5297(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5298 github.com/goccy/pythonwasm2go/p0.Fn5298
func Fn5298(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5300 github.com/goccy/pythonwasm2go/p2.Fn5300
func Fn5300(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5301 github.com/goccy/pythonwasm2go/p2.Fn5301
func Fn5301(m *base.Module) int32

//go:linkname Fn5303 github.com/goccy/pythonwasm2go/p0.Fn5303
func Fn5303(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5304 github.com/goccy/pythonwasm2go/p2.Fn5304
func Fn5304(m *base.Module, l0 int32)

//go:linkname Fn5305 github.com/goccy/pythonwasm2go/p2.Fn5305
func Fn5305(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5322 github.com/goccy/pythonwasm2go/p2.Fn5322
func Fn5322(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5323 github.com/goccy/pythonwasm2go/p2.Fn5323
func Fn5323(m *base.Module, l0 int32) int32

//go:linkname Fn5325 github.com/goccy/pythonwasm2go/p2.Fn5325
func Fn5325(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5327 github.com/goccy/pythonwasm2go/p2.Fn5327
func Fn5327(m *base.Module, l0 int32) int32

//go:linkname Fn5328 github.com/goccy/pythonwasm2go/p2.Fn5328
func Fn5328(m *base.Module, l0 int32) int32

//go:linkname Fn5329 github.com/goccy/pythonwasm2go/p2.Fn5329
func Fn5329(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5330 github.com/goccy/pythonwasm2go/p2.Fn5330
func Fn5330(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5331 github.com/goccy/pythonwasm2go/p2.Fn5331
func Fn5331(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5338 github.com/goccy/pythonwasm2go/p2.Fn5338
func Fn5338(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5343 github.com/goccy/pythonwasm2go/p2.Fn5343
func Fn5343(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5345 github.com/goccy/pythonwasm2go/p2.Fn5345
func Fn5345(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5347 github.com/goccy/pythonwasm2go/p2.Fn5347
func Fn5347(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5354 github.com/goccy/pythonwasm2go/p2.Fn5354
func Fn5354(m *base.Module, l0 int32)

//go:linkname Fn5355 github.com/goccy/pythonwasm2go/p2.Fn5355
func Fn5355(m *base.Module, l0 int32) int32

//go:linkname Fn5357 github.com/goccy/pythonwasm2go/p2.Fn5357
func Fn5357(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5358 github.com/goccy/pythonwasm2go/p2.Fn5358
func Fn5358(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5369 github.com/goccy/pythonwasm2go/p2.Fn5369
func Fn5369(m *base.Module, l0 int32)

//go:linkname Fn5377 github.com/goccy/pythonwasm2go/p2.Fn5377
func Fn5377(m *base.Module, l0 int32)

//go:linkname Fn5384 github.com/goccy/pythonwasm2go/p2.Fn5384
func Fn5384(m *base.Module, l0 int32) int32

//go:linkname Fn5415 github.com/goccy/pythonwasm2go/p2.Fn5415
func Fn5415(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5437 github.com/goccy/pythonwasm2go/p2.Fn5437
func Fn5437(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5456 github.com/goccy/pythonwasm2go/p2.Fn5456
func Fn5456(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5457 github.com/goccy/pythonwasm2go/p2.Fn5457
func Fn5457(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5547 github.com/goccy/pythonwasm2go/p2.Fn5547
func Fn5547(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5549 github.com/goccy/pythonwasm2go/p2.Fn5549
func Fn5549(m *base.Module, l0 int32)

//go:linkname Fn5550 github.com/goccy/pythonwasm2go/p2.Fn5550
func Fn5550(m *base.Module, l0 int32)

//go:linkname Fn5554 github.com/goccy/pythonwasm2go/p2.Fn5554
func Fn5554(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5557 github.com/goccy/pythonwasm2go/p2.Fn5557
func Fn5557(m *base.Module) int32

//go:linkname Fn5563 github.com/goccy/pythonwasm2go/p2.Fn5563
func Fn5563(m *base.Module, l0 int32)

//go:linkname Fn5565 github.com/goccy/pythonwasm2go/p2.Fn5565
func Fn5565(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5566 github.com/goccy/pythonwasm2go/p2.Fn5566
func Fn5566(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5567 github.com/goccy/pythonwasm2go/p2.Fn5567
func Fn5567(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5571 github.com/goccy/pythonwasm2go/p0.Fn5571
func Fn5571(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5578 github.com/goccy/pythonwasm2go/p2.Fn5578
func Fn5578(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn5589 github.com/goccy/pythonwasm2go/p2.Fn5589
func Fn5589(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5616 github.com/goccy/pythonwasm2go/p2.Fn5616
func Fn5616(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5620 github.com/goccy/pythonwasm2go/p2.Fn5620
func Fn5620(m *base.Module, l0 float64, l1 int32) int32

//go:linkname Fn5632 github.com/goccy/pythonwasm2go/p2.Fn5632
func Fn5632(m *base.Module, l0 int64) int32

//go:linkname Fn5641 github.com/goccy/pythonwasm2go/p2.Fn5641
func Fn5641(m *base.Module, l0 float64) float64

//go:linkname Fn5657 github.com/goccy/pythonwasm2go/p2.Fn5657
func Fn5657(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5666 github.com/goccy/pythonwasm2go/p2.Fn5666
func Fn5666(m *base.Module, l0 float64) float64

//go:linkname Fn5667 github.com/goccy/pythonwasm2go/p2.Fn5667
func Fn5667(m *base.Module, l0 float64) float64

//go:linkname Fn5671 github.com/goccy/pythonwasm2go/p2.Fn5671
func Fn5671(m *base.Module, l0 float64) int32

//go:linkname Fn5690 github.com/goccy/pythonwasm2go/p2.Fn5690
func Fn5690(m *base.Module)

//go:linkname Fn5705 github.com/goccy/pythonwasm2go/p2.Fn5705
func Fn5705(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5706 github.com/goccy/pythonwasm2go/p2.Fn5706
func Fn5706(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5707 github.com/goccy/pythonwasm2go/p2.Fn5707
func Fn5707(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5708 github.com/goccy/pythonwasm2go/p2.Fn5708
func Fn5708(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5709 github.com/goccy/pythonwasm2go/p2.Fn5709
func Fn5709(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5710 github.com/goccy/pythonwasm2go/p2.Fn5710
func Fn5710(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5711 github.com/goccy/pythonwasm2go/p2.Fn5711
func Fn5711(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5713 github.com/goccy/pythonwasm2go/p2.Fn5713
func Fn5713(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5719 github.com/goccy/pythonwasm2go/p2.Fn5719
func Fn5719(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5723 github.com/goccy/pythonwasm2go/p2.Fn5723
func Fn5723(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5724 github.com/goccy/pythonwasm2go/p2.Fn5724
func Fn5724(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5725 github.com/goccy/pythonwasm2go/p2.Fn5725
func Fn5725(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5726 github.com/goccy/pythonwasm2go/p2.Fn5726
func Fn5726(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5727 github.com/goccy/pythonwasm2go/p2.Fn5727
func Fn5727(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5730 github.com/goccy/pythonwasm2go/p2.Fn5730
func Fn5730(m *base.Module, l0 int32) int32

//go:linkname Fn5739 github.com/goccy/pythonwasm2go/p2.Fn5739
func Fn5739(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5740 github.com/goccy/pythonwasm2go/p2.Fn5740
func Fn5740(m *base.Module, l0 int32) int32

//go:linkname Fn5741 github.com/goccy/pythonwasm2go/p2.Fn5741
func Fn5741(m *base.Module, l0 int32) int32

//go:linkname Fn5742 github.com/goccy/pythonwasm2go/p2.Fn5742
func Fn5742(m *base.Module, l0 int32) int32

//go:linkname Fn5744 github.com/goccy/pythonwasm2go/p2.Fn5744
func Fn5744(m *base.Module, l0 int32) int32

//go:linkname Fn5745 github.com/goccy/pythonwasm2go/p2.Fn5745
func Fn5745(m *base.Module, l0 int32) int32

//go:linkname Fn5746 github.com/goccy/pythonwasm2go/p2.Fn5746
func Fn5746(m *base.Module, l0 int32) int32

//go:linkname Fn5747 github.com/goccy/pythonwasm2go/p2.Fn5747
func Fn5747(m *base.Module, l0 int32) int32

//go:linkname Fn5748 github.com/goccy/pythonwasm2go/p2.Fn5748
func Fn5748(m *base.Module, l0 int32) int32

//go:linkname Fn5749 github.com/goccy/pythonwasm2go/p2.Fn5749
func Fn5749(m *base.Module, l0 int32) int32

//go:linkname Fn5750 github.com/goccy/pythonwasm2go/p2.Fn5750
func Fn5750(m *base.Module, l0 int32) int32

//go:linkname Fn5751 github.com/goccy/pythonwasm2go/p2.Fn5751
func Fn5751(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5752 github.com/goccy/pythonwasm2go/p2.Fn5752
func Fn5752(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5753 github.com/goccy/pythonwasm2go/p2.Fn5753
func Fn5753(m *base.Module, l0 int32)

//go:linkname Fn5754 github.com/goccy/pythonwasm2go/p2.Fn5754
func Fn5754(m *base.Module, l0 int32)

//go:linkname Fn5756 github.com/goccy/pythonwasm2go/p2.Fn5756
func Fn5756(m *base.Module, l0 int32)

//go:linkname Fn5757 github.com/goccy/pythonwasm2go/p2.Fn5757
func Fn5757(m *base.Module, l0 int32)

//go:linkname Fn5758 github.com/goccy/pythonwasm2go/p2.Fn5758
func Fn5758(m *base.Module, l0 int32)

//go:linkname Fn5759 github.com/goccy/pythonwasm2go/p2.Fn5759
func Fn5759(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5760 github.com/goccy/pythonwasm2go/p2.Fn5760
func Fn5760(m *base.Module, l0 int32) int32

//go:linkname Fn5761 github.com/goccy/pythonwasm2go/p2.Fn5761
func Fn5761(m *base.Module, l0 int32) int32

//go:linkname Fn5762 github.com/goccy/pythonwasm2go/p2.Fn5762
func Fn5762(m *base.Module, l0 int32) int32

//go:linkname Fn5763 github.com/goccy/pythonwasm2go/p2.Fn5763
func Fn5763(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5764 github.com/goccy/pythonwasm2go/p2.Fn5764
func Fn5764(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5765 github.com/goccy/pythonwasm2go/p2.Fn5765
func Fn5765(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5766 github.com/goccy/pythonwasm2go/p2.Fn5766
func Fn5766(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5770 github.com/goccy/pythonwasm2go/p2.Fn5770
func Fn5770(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5771 github.com/goccy/pythonwasm2go/p2.Fn5771
func Fn5771(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5772 github.com/goccy/pythonwasm2go/p2.Fn5772
func Fn5772(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5773 github.com/goccy/pythonwasm2go/p2.Fn5773
func Fn5773(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5775 github.com/goccy/pythonwasm2go/p2.Fn5775
func Fn5775(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5777 github.com/goccy/pythonwasm2go/p2.Fn5777
func Fn5777(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5778 github.com/goccy/pythonwasm2go/p2.Fn5778
func Fn5778(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5779 github.com/goccy/pythonwasm2go/p2.Fn5779
func Fn5779(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5780 github.com/goccy/pythonwasm2go/p2.Fn5780
func Fn5780(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5781 github.com/goccy/pythonwasm2go/p2.Fn5781
func Fn5781(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5782 github.com/goccy/pythonwasm2go/p2.Fn5782
func Fn5782(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5783 github.com/goccy/pythonwasm2go/p2.Fn5783
func Fn5783(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5785 github.com/goccy/pythonwasm2go/p2.Fn5785
func Fn5785(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5786 github.com/goccy/pythonwasm2go/p2.Fn5786
func Fn5786(m *base.Module, l0 int32) int32

//go:linkname Fn5789 github.com/goccy/pythonwasm2go/p2.Fn5789
func Fn5789(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5790 github.com/goccy/pythonwasm2go/p2.Fn5790
func Fn5790(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5792 github.com/goccy/pythonwasm2go/p2.Fn5792
func Fn5792(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5793 github.com/goccy/pythonwasm2go/p2.Fn5793
func Fn5793(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5794 github.com/goccy/pythonwasm2go/p2.Fn5794
func Fn5794(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5796 github.com/goccy/pythonwasm2go/p2.Fn5796
func Fn5796(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5797 github.com/goccy/pythonwasm2go/p2.Fn5797
func Fn5797(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5803 github.com/goccy/pythonwasm2go/p2.Fn5803
func Fn5803(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5806 github.com/goccy/pythonwasm2go/p2.Fn5806
func Fn5806(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5809 github.com/goccy/pythonwasm2go/p2.Fn5809
func Fn5809(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5814 github.com/goccy/pythonwasm2go/p2.Fn5814
func Fn5814(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5816 github.com/goccy/pythonwasm2go/p2.Fn5816
func Fn5816(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5819 github.com/goccy/pythonwasm2go/p0.Fn5819
func Fn5819(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5825 github.com/goccy/pythonwasm2go/p2.Fn5825
func Fn5825(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5826 github.com/goccy/pythonwasm2go/p2.Fn5826
func Fn5826(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5827 github.com/goccy/pythonwasm2go/p2.Fn5827
func Fn5827(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn5828 github.com/goccy/pythonwasm2go/p2.Fn5828
func Fn5828(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn5830 github.com/goccy/pythonwasm2go/p2.Fn5830
func Fn5830(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5831 github.com/goccy/pythonwasm2go/p2.Fn5831
func Fn5831(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5836 github.com/goccy/pythonwasm2go/p2.Fn5836
func Fn5836(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5837 github.com/goccy/pythonwasm2go/p2.Fn5837
func Fn5837(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5840 github.com/goccy/pythonwasm2go/p2.Fn5840
func Fn5840(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5841 github.com/goccy/pythonwasm2go/p2.Fn5841
func Fn5841(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5843 github.com/goccy/pythonwasm2go/p2.Fn5843
func Fn5843(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5844 github.com/goccy/pythonwasm2go/p2.Fn5844
func Fn5844(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5846 github.com/goccy/pythonwasm2go/p2.Fn5846
func Fn5846(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5847 github.com/goccy/pythonwasm2go/p2.Fn5847
func Fn5847(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5848 github.com/goccy/pythonwasm2go/p2.Fn5848
func Fn5848(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5851 github.com/goccy/pythonwasm2go/p2.Fn5851
func Fn5851(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5852 github.com/goccy/pythonwasm2go/p0.Fn5852
func Fn5852(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5854 github.com/goccy/pythonwasm2go/p2.Fn5854
func Fn5854(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5856 github.com/goccy/pythonwasm2go/p2.Fn5856
func Fn5856(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5860 github.com/goccy/pythonwasm2go/p2.Fn5860
func Fn5860(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5863 github.com/goccy/pythonwasm2go/p2.Fn5863
func Fn5863(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5864 github.com/goccy/pythonwasm2go/p2.Fn5864
func Fn5864(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5867 github.com/goccy/pythonwasm2go/p2.Fn5867
func Fn5867(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5873 github.com/goccy/pythonwasm2go/p2.Fn5873
func Fn5873(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5875 github.com/goccy/pythonwasm2go/p2.Fn5875
func Fn5875(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5876 github.com/goccy/pythonwasm2go/p2.Fn5876
func Fn5876(m *base.Module, l0 int32) int32

//go:linkname Fn5877 github.com/goccy/pythonwasm2go/p2.Fn5877
func Fn5877(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5879 github.com/goccy/pythonwasm2go/p2.Fn5879
func Fn5879(m *base.Module) int32

//go:linkname Fn5880 github.com/goccy/pythonwasm2go/p2.Fn5880
func Fn5880(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5881 github.com/goccy/pythonwasm2go/p2.Fn5881
func Fn5881(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5883 github.com/goccy/pythonwasm2go/p2.Fn5883
func Fn5883(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5884 github.com/goccy/pythonwasm2go/p2.Fn5884
func Fn5884(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5885 github.com/goccy/pythonwasm2go/p2.Fn5885
func Fn5885(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5887 github.com/goccy/pythonwasm2go/p2.Fn5887
func Fn5887(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5895 github.com/goccy/pythonwasm2go/p2.Fn5895
func Fn5895(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5896 github.com/goccy/pythonwasm2go/p2.Fn5896
func Fn5896(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5903 github.com/goccy/pythonwasm2go/p2.Fn5903
func Fn5903(m *base.Module, l0 int32) int32

//go:linkname Fn5905 github.com/goccy/pythonwasm2go/p2.Fn5905
func Fn5905(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5909 github.com/goccy/pythonwasm2go/p2.Fn5909
func Fn5909(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5910 github.com/goccy/pythonwasm2go/p2.Fn5910
func Fn5910(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5911 github.com/goccy/pythonwasm2go/p2.Fn5911
func Fn5911(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5912 github.com/goccy/pythonwasm2go/p2.Fn5912
func Fn5912(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5913 github.com/goccy/pythonwasm2go/p2.Fn5913
func Fn5913(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5914 github.com/goccy/pythonwasm2go/p2.Fn5914
func Fn5914(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5915 github.com/goccy/pythonwasm2go/p2.Fn5915
func Fn5915(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5916 github.com/goccy/pythonwasm2go/p2.Fn5916
func Fn5916(m *base.Module, l0 int32)

//go:linkname Fn5917 github.com/goccy/pythonwasm2go/p2.Fn5917
func Fn5917(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5918 github.com/goccy/pythonwasm2go/p2.Fn5918
func Fn5918(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5919 github.com/goccy/pythonwasm2go/p2.Fn5919
func Fn5919(m *base.Module, l0 int32)

//go:linkname Fn5920 github.com/goccy/pythonwasm2go/p2.Fn5920
func Fn5920(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5928 github.com/goccy/pythonwasm2go/p2.Fn5928
func Fn5928(m *base.Module, l0 int32) int32

//go:linkname Fn5933 github.com/goccy/pythonwasm2go/p2.Fn5933
func Fn5933(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5934 github.com/goccy/pythonwasm2go/p2.Fn5934
func Fn5934(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5937 github.com/goccy/pythonwasm2go/p2.Fn5937
func Fn5937(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5939 github.com/goccy/pythonwasm2go/p2.Fn5939
func Fn5939(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5941 github.com/goccy/pythonwasm2go/p2.Fn5941
func Fn5941(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5942 github.com/goccy/pythonwasm2go/p2.Fn5942
func Fn5942(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5945 github.com/goccy/pythonwasm2go/p2.Fn5945
func Fn5945(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5956 github.com/goccy/pythonwasm2go/p0.Fn5956
func Fn5956(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5960 github.com/goccy/pythonwasm2go/p0.Fn5960
func Fn5960(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn6025 github.com/goccy/pythonwasm2go/p2.Fn6025
func Fn6025(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6078 github.com/goccy/pythonwasm2go/p2.Fn6078
func Fn6078(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6104 github.com/goccy/pythonwasm2go/p2.Fn6104
func Fn6104(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6106 github.com/goccy/pythonwasm2go/p2.Fn6106
func Fn6106(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6107 github.com/goccy/pythonwasm2go/p2.Fn6107
func Fn6107(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6132 github.com/goccy/pythonwasm2go/p2.Fn6132
func Fn6132(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6144 github.com/goccy/pythonwasm2go/p2.Fn6144
func Fn6144(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6146 github.com/goccy/pythonwasm2go/p2.Fn6146
func Fn6146(m *base.Module, l0 int32)

//go:linkname Fn6147 github.com/goccy/pythonwasm2go/p2.Fn6147
func Fn6147(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6148 github.com/goccy/pythonwasm2go/p0.Fn6148
func Fn6148(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6149 github.com/goccy/pythonwasm2go/p2.Fn6149
func Fn6149(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32)

//go:linkname Fn6150 github.com/goccy/pythonwasm2go/p2.Fn6150
func Fn6150(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6155 github.com/goccy/pythonwasm2go/p2.Fn6155
func Fn6155(m *base.Module, l0 int32) int32

//go:linkname Fn6165 github.com/goccy/pythonwasm2go/p2.Fn6165
func Fn6165(m *base.Module, l0 int32)

//go:linkname Fn6166 github.com/goccy/pythonwasm2go/p2.Fn6166
func Fn6166(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6168 github.com/goccy/pythonwasm2go/p2.Fn6168
func Fn6168(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32)

//go:linkname Fn6169 github.com/goccy/pythonwasm2go/p2.Fn6169
func Fn6169(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6182 github.com/goccy/pythonwasm2go/p2.Fn6182
func Fn6182(m *base.Module, l0 int32)

//go:linkname Fn6183 github.com/goccy/pythonwasm2go/p2.Fn6183
func Fn6183(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6185 github.com/goccy/pythonwasm2go/p2.Fn6185
func Fn6185(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6187 github.com/goccy/pythonwasm2go/p2.Fn6187
func Fn6187(m *base.Module, l0 int32)

//go:linkname Fn6188 github.com/goccy/pythonwasm2go/p2.Fn6188
func Fn6188(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6189 github.com/goccy/pythonwasm2go/p2.Fn6189
func Fn6189(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6190 github.com/goccy/pythonwasm2go/p2.Fn6190
func Fn6190(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6191 github.com/goccy/pythonwasm2go/p2.Fn6191
func Fn6191(m *base.Module, l0 int32)

//go:linkname Fn6192 github.com/goccy/pythonwasm2go/p2.Fn6192
func Fn6192(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6193 github.com/goccy/pythonwasm2go/p0.Fn6193
func Fn6193(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6196 github.com/goccy/pythonwasm2go/p2.Fn6196
func Fn6196(m *base.Module, l0 int32)

//go:linkname Fn6197 github.com/goccy/pythonwasm2go/p2.Fn6197
func Fn6197(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6198 github.com/goccy/pythonwasm2go/p2.Fn6198
func Fn6198(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6202 github.com/goccy/pythonwasm2go/p2.Fn6202
func Fn6202(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6203 github.com/goccy/pythonwasm2go/p2.Fn6203
func Fn6203(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6206 github.com/goccy/pythonwasm2go/p2.Fn6206
func Fn6206(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6207 github.com/goccy/pythonwasm2go/p2.Fn6207
func Fn6207(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6213 github.com/goccy/pythonwasm2go/p2.Fn6213
func Fn6213(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6214 github.com/goccy/pythonwasm2go/p2.Fn6214
func Fn6214(m *base.Module, l0 int32) int32

//go:linkname Fn6216 github.com/goccy/pythonwasm2go/p2.Fn6216
func Fn6216(m *base.Module, l0 int32) int32

//go:linkname Fn6218 github.com/goccy/pythonwasm2go/p2.Fn6218
func Fn6218(m *base.Module, l0 int32) int32

//go:linkname Fn6220 github.com/goccy/pythonwasm2go/p2.Fn6220
func Fn6220(m *base.Module, l0 int32) int32

//go:linkname Fn6233 github.com/goccy/pythonwasm2go/p2.Fn6233
func Fn6233(m *base.Module, l0 int32)

//go:linkname Fn6234 github.com/goccy/pythonwasm2go/p2.Fn6234
func Fn6234(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6235 github.com/goccy/pythonwasm2go/p2.Fn6235
func Fn6235(m *base.Module, l0 int32) int32

//go:linkname Fn6240 github.com/goccy/pythonwasm2go/p2.Fn6240
func Fn6240(m *base.Module, l0 int32) int32

//go:linkname Fn6256 github.com/goccy/pythonwasm2go/p2.Fn6256
func Fn6256(m *base.Module, l0 int32) int32

//go:linkname Fn6270 github.com/goccy/pythonwasm2go/p2.Fn6270
func Fn6270(m *base.Module, l0 int32, l1 int64)

//go:linkname Fn6272 github.com/goccy/pythonwasm2go/p2.Fn6272
func Fn6272(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn6274 github.com/goccy/pythonwasm2go/p2.Fn6274
func Fn6274(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn6275 github.com/goccy/pythonwasm2go/p2.Fn6275
func Fn6275(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6279 github.com/goccy/pythonwasm2go/p2.Fn6279
func Fn6279(m *base.Module, l0 int32)

//go:linkname Fn6280 github.com/goccy/pythonwasm2go/p2.Fn6280
func Fn6280(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32)

//go:linkname Fn6282 github.com/goccy/pythonwasm2go/p2.Fn6282
func Fn6282(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32, l6 int32)

//go:linkname Fn6283 github.com/goccy/pythonwasm2go/p2.Fn6283
func Fn6283(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6285 github.com/goccy/pythonwasm2go/p2.Fn6285
func Fn6285(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6292 github.com/goccy/pythonwasm2go/p2.Fn6292
func Fn6292(m *base.Module, l0 int32) int32

//go:linkname Fn6301 github.com/goccy/pythonwasm2go/p2.Fn6301
func Fn6301(m *base.Module, l0 int32) int32

//go:linkname Fn6302 github.com/goccy/pythonwasm2go/p2.Fn6302
func Fn6302(m *base.Module, l0 int32) int32

//go:linkname Fn6303 github.com/goccy/pythonwasm2go/p2.Fn6303
func Fn6303(m *base.Module, l0 int32) int32

//go:linkname Fn6305 github.com/goccy/pythonwasm2go/p2.Fn6305
func Fn6305(m *base.Module, l0 int32)

//go:linkname Fn6311 github.com/goccy/pythonwasm2go/p2.Fn6311
func Fn6311(m *base.Module, l0 int32)

//go:linkname Fn6329 github.com/goccy/pythonwasm2go/p2.Fn6329
func Fn6329(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6330 github.com/goccy/pythonwasm2go/p2.Fn6330
func Fn6330(m *base.Module, l0 int32) int32

//go:linkname Fn6359 github.com/goccy/pythonwasm2go/p2.Fn6359
func Fn6359(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6371 github.com/goccy/pythonwasm2go/p2.Fn6371
func Fn6371(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6373 github.com/goccy/pythonwasm2go/p2.Fn6373
func Fn6373(m *base.Module, l0 int32) int32

//go:linkname Fn6379 github.com/goccy/pythonwasm2go/p2.Fn6379
func Fn6379(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6387 github.com/goccy/pythonwasm2go/p2.Fn6387
func Fn6387(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32

//go:linkname Fn6391 github.com/goccy/pythonwasm2go/p2.Fn6391
func Fn6391(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6395 github.com/goccy/pythonwasm2go/p2.Fn6395
func Fn6395(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32

//go:linkname Fn6397 github.com/goccy/pythonwasm2go/p2.Fn6397
func Fn6397(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6418 github.com/goccy/pythonwasm2go/p2.Fn6418
func Fn6418(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6429 github.com/goccy/pythonwasm2go/p2.Fn6429
func Fn6429(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6450 github.com/goccy/pythonwasm2go/p2.Fn6450
func Fn6450(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6473 github.com/goccy/pythonwasm2go/p2.Fn6473
func Fn6473(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6524 github.com/goccy/pythonwasm2go/p2.Fn6524
func Fn6524(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6526 github.com/goccy/pythonwasm2go/p2.Fn6526
func Fn6526(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6527 github.com/goccy/pythonwasm2go/p2.Fn6527
func Fn6527(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn6528 github.com/goccy/pythonwasm2go/p2.Fn6528
func Fn6528(m *base.Module, l0 int32) int32

//go:linkname Fn6529 github.com/goccy/pythonwasm2go/p2.Fn6529
func Fn6529(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int64, l5 int32)

//go:linkname Fn6530 github.com/goccy/pythonwasm2go/p2.Fn6530
func Fn6530(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6531 github.com/goccy/pythonwasm2go/p2.Fn6531
func Fn6531(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6533 github.com/goccy/pythonwasm2go/p2.Fn6533
func Fn6533(m *base.Module, l0 int32)

//go:linkname Fn6534 github.com/goccy/pythonwasm2go/p2.Fn6534
func Fn6534(m *base.Module, l0 int32) int32

//go:linkname Fn6535 github.com/goccy/pythonwasm2go/p2.Fn6535
func Fn6535(m *base.Module, l0 int32)

//go:linkname Fn6537 github.com/goccy/pythonwasm2go/p2.Fn6537
func Fn6537(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6538 github.com/goccy/pythonwasm2go/p2.Fn6538
func Fn6538(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6539 github.com/goccy/pythonwasm2go/p2.Fn6539
func Fn6539(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6547 github.com/goccy/pythonwasm2go/p2.Fn6547
func Fn6547(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6548 github.com/goccy/pythonwasm2go/p2.Fn6548
func Fn6548(m *base.Module, l0 int32)

//go:linkname Fn6553 github.com/goccy/pythonwasm2go/p2.Fn6553
func Fn6553(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6574 github.com/goccy/pythonwasm2go/p2.Fn6574
func Fn6574(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6581 github.com/goccy/pythonwasm2go/p2.Fn6581
func Fn6581(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6582 github.com/goccy/pythonwasm2go/p2.Fn6582
func Fn6582(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6583 github.com/goccy/pythonwasm2go/p2.Fn6583
func Fn6583(m *base.Module, l0 int32) int32

//go:linkname Fn6590 github.com/goccy/pythonwasm2go/p2.Fn6590
func Fn6590(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6591 github.com/goccy/pythonwasm2go/p2.Fn6591
func Fn6591(m *base.Module, l0 int32) int32

//go:linkname Fn6597 github.com/goccy/pythonwasm2go/p2.Fn6597
func Fn6597(m *base.Module, l0 int32) int32

//go:linkname Fn6600 github.com/goccy/pythonwasm2go/p0.Fn6600
func Fn6600(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32

//go:linkname Fn6602 github.com/goccy/pythonwasm2go/p2.Fn6602
func Fn6602(m *base.Module, l0 int32)

//go:linkname Fn6604 github.com/goccy/pythonwasm2go/p2.Fn6604
func Fn6604(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6609 github.com/goccy/pythonwasm2go/p2.Fn6609
func Fn6609(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6613 github.com/goccy/pythonwasm2go/p2.Fn6613
func Fn6613(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6616 github.com/goccy/pythonwasm2go/p2.Fn6616
func Fn6616(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6618 github.com/goccy/pythonwasm2go/p2.Fn6618
func Fn6618(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6619 github.com/goccy/pythonwasm2go/p2.Fn6619
func Fn6619(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6620 github.com/goccy/pythonwasm2go/p2.Fn6620
func Fn6620(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6621 github.com/goccy/pythonwasm2go/p2.Fn6621
func Fn6621(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6622 github.com/goccy/pythonwasm2go/p2.Fn6622
func Fn6622(m *base.Module, l0 int32) float32

//go:linkname Fn6623 github.com/goccy/pythonwasm2go/p2.Fn6623
func Fn6623(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6624 github.com/goccy/pythonwasm2go/p2.Fn6624
func Fn6624(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6625 github.com/goccy/pythonwasm2go/p0.Fn6625
func Fn6625(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn6636 github.com/goccy/pythonwasm2go/p2.Fn6636
func Fn6636(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6642 github.com/goccy/pythonwasm2go/p2.Fn6642
func Fn6642(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6644 github.com/goccy/pythonwasm2go/p2.Fn6644
func Fn6644(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6648 github.com/goccy/pythonwasm2go/p2.Fn6648
func Fn6648(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6650 github.com/goccy/pythonwasm2go/p2.Fn6650
func Fn6650(m *base.Module, l0 int32) int32

//go:linkname Fn6652 github.com/goccy/pythonwasm2go/p2.Fn6652
func Fn6652(m *base.Module, l0 int32) int32

//go:linkname Fn6653 github.com/goccy/pythonwasm2go/p2.Fn6653
func Fn6653(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6654 github.com/goccy/pythonwasm2go/p2.Fn6654
func Fn6654(m *base.Module, l0 int32)

//go:linkname Fn6669 github.com/goccy/pythonwasm2go/p2.Fn6669
func Fn6669(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6678 github.com/goccy/pythonwasm2go/p2.Fn6678
func Fn6678(m *base.Module, l0 int32) int32

//go:linkname Fn6685 github.com/goccy/pythonwasm2go/p2.Fn6685
func Fn6685(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6691 github.com/goccy/pythonwasm2go/p2.Fn6691
func Fn6691(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6730 github.com/goccy/pythonwasm2go/p2.Fn6730
func Fn6730(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6731 github.com/goccy/pythonwasm2go/p2.Fn6731
func Fn6731(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6733 github.com/goccy/pythonwasm2go/p2.Fn6733
func Fn6733(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6734 github.com/goccy/pythonwasm2go/p2.Fn6734
func Fn6734(m *base.Module, l0 int32)

//go:linkname Fn6735 github.com/goccy/pythonwasm2go/p2.Fn6735
func Fn6735(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6740 github.com/goccy/pythonwasm2go/p2.Fn6740
func Fn6740(m *base.Module, l0 int32) int32

//go:linkname Fn6741 github.com/goccy/pythonwasm2go/p2.Fn6741
func Fn6741(m *base.Module, l0 int32) int32

//go:linkname Fn6742 github.com/goccy/pythonwasm2go/p2.Fn6742
func Fn6742(m *base.Module, l0 int32) int32

//go:linkname Fn6748 github.com/goccy/pythonwasm2go/p2.Fn6748
func Fn6748(m *base.Module, l0 int32) int32

//go:linkname Fn6760 github.com/goccy/pythonwasm2go/p2.Fn6760
func Fn6760(m *base.Module, l0 int32) int32

//go:linkname Fn6761 github.com/goccy/pythonwasm2go/p2.Fn6761
func Fn6761(m *base.Module, l0 int32) int32

//go:linkname Fn6762 github.com/goccy/pythonwasm2go/p2.Fn6762
func Fn6762(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6763 github.com/goccy/pythonwasm2go/p2.Fn6763
func Fn6763(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6770 github.com/goccy/pythonwasm2go/p2.Fn6770
func Fn6770(m *base.Module, l0 int32)

//go:linkname Fn6772 github.com/goccy/pythonwasm2go/p2.Fn6772
func Fn6772(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6775 github.com/goccy/pythonwasm2go/p2.Fn6775
func Fn6775(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6780 github.com/goccy/pythonwasm2go/p2.Fn6780
func Fn6780(m *base.Module, l0 int32)

//go:linkname Fn6794 github.com/goccy/pythonwasm2go/p2.Fn6794
func Fn6794(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6818 github.com/goccy/pythonwasm2go/p2.Fn6818
func Fn6818(m *base.Module, l0 int32) int32

//go:linkname Fn6830 github.com/goccy/pythonwasm2go/p2.Fn6830
func Fn6830(m *base.Module, l0 int32) int32

//go:linkname Fn6831 github.com/goccy/pythonwasm2go/p2.Fn6831
func Fn6831(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn6894 github.com/goccy/pythonwasm2go/p2.Fn6894
func Fn6894(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6914 github.com/goccy/pythonwasm2go/p2.Fn6914
func Fn6914(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6940 github.com/goccy/pythonwasm2go/p2.Fn6940
func Fn6940(m *base.Module, l0 int32) int32

//go:linkname Fn6942 github.com/goccy/pythonwasm2go/p2.Fn6942
func Fn6942(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6946 github.com/goccy/pythonwasm2go/p2.Fn6946
func Fn6946(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6947 github.com/goccy/pythonwasm2go/p2.Fn6947
func Fn6947(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6952 github.com/goccy/pythonwasm2go/p2.Fn6952
func Fn6952(m *base.Module, l0 int32) int32

//go:linkname Fn6963 github.com/goccy/pythonwasm2go/p2.Fn6963
func Fn6963(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6964 github.com/goccy/pythonwasm2go/p2.Fn6964
func Fn6964(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6987 github.com/goccy/pythonwasm2go/p2.Fn6987
func Fn6987(m *base.Module, l0 int32) int32

//go:linkname Fn6998 github.com/goccy/pythonwasm2go/p2.Fn6998
func Fn6998(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7002 github.com/goccy/pythonwasm2go/p2.Fn7002
func Fn7002(m *base.Module, l0 int32) int32

//go:linkname Fn7003 github.com/goccy/pythonwasm2go/p2.Fn7003
func Fn7003(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7004 github.com/goccy/pythonwasm2go/p2.Fn7004
func Fn7004(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn7005 github.com/goccy/pythonwasm2go/p2.Fn7005
func Fn7005(m *base.Module, l0 int32) int32

//go:linkname Fn7006 github.com/goccy/pythonwasm2go/p2.Fn7006
func Fn7006(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7007 github.com/goccy/pythonwasm2go/p2.Fn7007
func Fn7007(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn7008 github.com/goccy/pythonwasm2go/p2.Fn7008
func Fn7008(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7021 github.com/goccy/pythonwasm2go/p2.Fn7021
func Fn7021(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7022 github.com/goccy/pythonwasm2go/p2.Fn7022
func Fn7022(m *base.Module, l0 int32)

//go:linkname Fn7038 github.com/goccy/pythonwasm2go/p2.Fn7038
func Fn7038(m *base.Module) int32

//go:linkname Fn7048 github.com/goccy/pythonwasm2go/p2.Fn7048
func Fn7048(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7049 github.com/goccy/pythonwasm2go/p2.Fn7049
func Fn7049(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7061 github.com/goccy/pythonwasm2go/p2.Fn7061
func Fn7061(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7089 github.com/goccy/pythonwasm2go/p2.Fn7089
func Fn7089(m *base.Module) int64

//go:linkname Fn7104 github.com/goccy/pythonwasm2go/p2.Fn7104
func Fn7104(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7121 github.com/goccy/pythonwasm2go/p2.Fn7121
func Fn7121(m *base.Module, l0 int32) int32

//go:linkname Fn7125 github.com/goccy/pythonwasm2go/p2.Fn7125
func Fn7125(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7131 github.com/goccy/pythonwasm2go/p2.Fn7131
func Fn7131(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7132 github.com/goccy/pythonwasm2go/p2.Fn7132
func Fn7132(m *base.Module, l0 int32)

//go:linkname Fn7133 github.com/goccy/pythonwasm2go/p0.Fn7133
func Fn7133(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn7134 github.com/goccy/pythonwasm2go/p2.Fn7134
func Fn7134(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7138 github.com/goccy/pythonwasm2go/p2.Fn7138
func Fn7138(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7141 github.com/goccy/pythonwasm2go/p2.Fn7141
func Fn7141(m *base.Module, l0 int32)

//go:linkname Fn7146 github.com/goccy/pythonwasm2go/p2.Fn7146
func Fn7146(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7175 github.com/goccy/pythonwasm2go/p2.Fn7175
func Fn7175(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7176 github.com/goccy/pythonwasm2go/p2.Fn7176
func Fn7176(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7177 github.com/goccy/pythonwasm2go/p2.Fn7177
func Fn7177(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7180 github.com/goccy/pythonwasm2go/p2.Fn7180
func Fn7180(m *base.Module, l0 int32) int32

//go:linkname Fn7199 github.com/goccy/pythonwasm2go/p2.Fn7199
func Fn7199(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7200 github.com/goccy/pythonwasm2go/p2.Fn7200
func Fn7200(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7211 github.com/goccy/pythonwasm2go/p2.Fn7211
func Fn7211(m *base.Module, l0 int32) int32

//go:linkname Fn7214 github.com/goccy/pythonwasm2go/p2.Fn7214
func Fn7214(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7225 github.com/goccy/pythonwasm2go/p2.Fn7225
func Fn7225(m *base.Module, l0 int32) int32

//go:linkname Fn7227 github.com/goccy/pythonwasm2go/p2.Fn7227
func Fn7227(m *base.Module, l0 int32) int32

//go:linkname Fn7228 github.com/goccy/pythonwasm2go/p2.Fn7228
func Fn7228(m *base.Module, l0 int32) int32

//go:linkname Fn7229 github.com/goccy/pythonwasm2go/p2.Fn7229
func Fn7229(m *base.Module, l0 int32) int32

//go:linkname Fn7245 github.com/goccy/pythonwasm2go/p2.Fn7245
func Fn7245(m *base.Module, l0 int32) int32

//go:linkname Fn7282 github.com/goccy/pythonwasm2go/p2.Fn7282
func Fn7282(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7284 github.com/goccy/pythonwasm2go/p2.Fn7284
func Fn7284(m *base.Module, l0 int32) int32

//go:linkname Fn7300 github.com/goccy/pythonwasm2go/p2.Fn7300
func Fn7300(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7329 github.com/goccy/pythonwasm2go/p2.Fn7329
func Fn7329(m *base.Module, l0 int32) int32

//go:linkname Fn7336 github.com/goccy/pythonwasm2go/p2.Fn7336
func Fn7336(m *base.Module, l0 int32)

//go:linkname Fn7337 github.com/goccy/pythonwasm2go/p2.Fn7337
func Fn7337(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7350 github.com/goccy/pythonwasm2go/p2.Fn7350
func Fn7350(m *base.Module, l0 int32) int32

//go:linkname Fn7402 github.com/goccy/pythonwasm2go/p2.Fn7402
func Fn7402(m *base.Module, l0 int32) int32

//go:linkname Fn7412 github.com/goccy/pythonwasm2go/p2.Fn7412
func Fn7412(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7449 github.com/goccy/pythonwasm2go/p2.Fn7449
func Fn7449(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7476 github.com/goccy/pythonwasm2go/p2.Fn7476
func Fn7476(m *base.Module, l0 int32) int32

//go:linkname Fn7477 github.com/goccy/pythonwasm2go/p2.Fn7477
func Fn7477(m *base.Module, l0 int32) int32

//go:linkname Fn7487 github.com/goccy/pythonwasm2go/p2.Fn7487
func Fn7487(m *base.Module, l0 int32) int32

//go:linkname Fn7511 github.com/goccy/pythonwasm2go/p2.Fn7511
func Fn7511(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7518 github.com/goccy/pythonwasm2go/p2.Fn7518
func Fn7518(m *base.Module, l0 int32) int32

//go:linkname Fn7540 github.com/goccy/pythonwasm2go/p2.Fn7540
func Fn7540(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7758 github.com/goccy/pythonwasm2go/p2.Fn7758
func Fn7758(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7759 github.com/goccy/pythonwasm2go/p2.Fn7759
func Fn7759(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7822 github.com/goccy/pythonwasm2go/p2.Fn7822
func Fn7822(m *base.Module, l0 int32)

//go:linkname Fn7851 github.com/goccy/pythonwasm2go/p2.Fn7851
func Fn7851(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7866 github.com/goccy/pythonwasm2go/p2.Fn7866
func Fn7866(m *base.Module, l0 int32) int32

//go:linkname Fn7909 github.com/goccy/pythonwasm2go/p2.Fn7909
func Fn7909(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7981 github.com/goccy/pythonwasm2go/p0.Fn7981
func Fn7981(m *base.Module)

//go:linkname Fn7984 github.com/goccy/pythonwasm2go/p2.Fn7984
func Fn7984(m *base.Module, l0 int32)

//go:linkname Fn7987 github.com/goccy/pythonwasm2go/p2.Fn7987
func Fn7987(m *base.Module) int32

//go:linkname Fn7997 github.com/goccy/pythonwasm2go/p2.Fn7997
func Fn7997(m *base.Module, l0 int32) int32

//go:linkname Fn8007 github.com/goccy/pythonwasm2go/p2.Fn8007
func Fn8007(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8008 github.com/goccy/pythonwasm2go/p2.Fn8008
func Fn8008(m *base.Module, l0 int32) int32

//go:linkname Fn8016 github.com/goccy/pythonwasm2go/p2.Fn8016
func Fn8016(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8018 github.com/goccy/pythonwasm2go/p2.Fn8018
func Fn8018(m *base.Module, l0 int32)

//go:linkname Fn8021 github.com/goccy/pythonwasm2go/p2.Fn8021
func Fn8021(m *base.Module, l0 int32) int32

//go:linkname Fn8023 github.com/goccy/pythonwasm2go/p2.Fn8023
func Fn8023(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8024 github.com/goccy/pythonwasm2go/p2.Fn8024
func Fn8024(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8029 github.com/goccy/pythonwasm2go/p2.Fn8029
func Fn8029(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8034 github.com/goccy/pythonwasm2go/p0.Fn8034
func Fn8034(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32) int32

//go:linkname Fn8045 github.com/goccy/pythonwasm2go/p2.Fn8045
func Fn8045(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8046 github.com/goccy/pythonwasm2go/p2.Fn8046
func Fn8046(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8047 github.com/goccy/pythonwasm2go/p2.Fn8047
func Fn8047(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8048 github.com/goccy/pythonwasm2go/p2.Fn8048
func Fn8048(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8060 github.com/goccy/pythonwasm2go/p2.Fn8060
func Fn8060(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8071 github.com/goccy/pythonwasm2go/p2.Fn8071
func Fn8071(m *base.Module, l0 int32, l1 int64) int32

//go:linkname Fn8099 github.com/goccy/pythonwasm2go/p2.Fn8099
func Fn8099(m *base.Module, l0 int32) int32

//go:linkname Fn8106 github.com/goccy/pythonwasm2go/p2.Fn8106
func Fn8106(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8107 github.com/goccy/pythonwasm2go/p2.Fn8107
func Fn8107(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int64, l6 int32) int32

//go:linkname Fn8111 github.com/goccy/pythonwasm2go/p2.Fn8111
func Fn8111(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8117 github.com/goccy/pythonwasm2go/p2.Fn8117
func Fn8117(m *base.Module, l0 int32)

//go:linkname Fn8130 github.com/goccy/pythonwasm2go/p2.Fn8130
func Fn8130(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8131 github.com/goccy/pythonwasm2go/p2.Fn8131
func Fn8131(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8134 github.com/goccy/pythonwasm2go/p0.Fn8134
func Fn8134(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8135 github.com/goccy/pythonwasm2go/p0.Fn8135
func Fn8135(m *base.Module) int32

//go:linkname Fn8136 github.com/goccy/pythonwasm2go/p2.Fn8136
func Fn8136(m *base.Module, l0 int32) int32

//go:linkname Fn8137 github.com/goccy/pythonwasm2go/p0.Fn8137
func Fn8137(m *base.Module, l0 int32) int32

//go:linkname Fn8162 github.com/goccy/pythonwasm2go/p2.Fn8162
func Fn8162(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8163 github.com/goccy/pythonwasm2go/p2.Fn8163
func Fn8163(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn8164 github.com/goccy/pythonwasm2go/p2.Fn8164
func Fn8164(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32

//go:linkname Fn8166 github.com/goccy/pythonwasm2go/p2.Fn8166
func Fn8166(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8169 github.com/goccy/pythonwasm2go/p2.Fn8169
func Fn8169(m *base.Module, l0 int32) int32

//go:linkname Fn8177 github.com/goccy/pythonwasm2go/p2.Fn8177
func Fn8177(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8178 github.com/goccy/pythonwasm2go/p2.Fn8178
func Fn8178(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8179 github.com/goccy/pythonwasm2go/p2.Fn8179
func Fn8179(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8181 github.com/goccy/pythonwasm2go/p2.Fn8181
func Fn8181(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8182 github.com/goccy/pythonwasm2go/p2.Fn8182
func Fn8182(m *base.Module, l0 int32) int32

//go:linkname Fn8184 github.com/goccy/pythonwasm2go/p2.Fn8184
func Fn8184(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8185 github.com/goccy/pythonwasm2go/p2.Fn8185
func Fn8185(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8186 github.com/goccy/pythonwasm2go/p2.Fn8186
func Fn8186(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32

//go:linkname Fn8188 github.com/goccy/pythonwasm2go/p2.Fn8188
func Fn8188(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32

//go:linkname Fn8189 github.com/goccy/pythonwasm2go/p2.Fn8189
func Fn8189(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8190 github.com/goccy/pythonwasm2go/p2.Fn8190
func Fn8190(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8193 github.com/goccy/pythonwasm2go/p2.Fn8193
func Fn8193(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8199 github.com/goccy/pythonwasm2go/p2.Fn8199
func Fn8199(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8200 github.com/goccy/pythonwasm2go/p2.Fn8200
func Fn8200(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8202 github.com/goccy/pythonwasm2go/p2.Fn8202
func Fn8202(m *base.Module, l0 int32) int32

//go:linkname Fn8212 github.com/goccy/pythonwasm2go/p2.Fn8212
func Fn8212(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8213 github.com/goccy/pythonwasm2go/p2.Fn8213
func Fn8213(m *base.Module, l0 int32) int32

//go:linkname Fn8217 github.com/goccy/pythonwasm2go/p2.Fn8217
func Fn8217(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8219 github.com/goccy/pythonwasm2go/p2.Fn8219
func Fn8219(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8221 github.com/goccy/pythonwasm2go/p2.Fn8221
func Fn8221(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8228 github.com/goccy/pythonwasm2go/p2.Fn8228
func Fn8228(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8232 github.com/goccy/pythonwasm2go/p2.Fn8232
func Fn8232(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8240 github.com/goccy/pythonwasm2go/p2.Fn8240
func Fn8240(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn8249 github.com/goccy/pythonwasm2go/p2.Fn8249
func Fn8249(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn8251 github.com/goccy/pythonwasm2go/p2.Fn8251
func Fn8251(m *base.Module, l0 int32) int32

//go:linkname Fn8252 github.com/goccy/pythonwasm2go/p2.Fn8252
func Fn8252(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8260 github.com/goccy/pythonwasm2go/p2.Fn8260
func Fn8260(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8261 github.com/goccy/pythonwasm2go/p2.Fn8261
func Fn8261(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8262 github.com/goccy/pythonwasm2go/p2.Fn8262
func Fn8262(m *base.Module, l0 int32) int32

//go:linkname Fn8263 github.com/goccy/pythonwasm2go/p2.Fn8263
func Fn8263(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8270 github.com/goccy/pythonwasm2go/p2.Fn8270
func Fn8270(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8271 github.com/goccy/pythonwasm2go/p2.Fn8271
func Fn8271(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8273 github.com/goccy/pythonwasm2go/p2.Fn8273
func Fn8273(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8277 github.com/goccy/pythonwasm2go/p2.Fn8277
func Fn8277(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8286 github.com/goccy/pythonwasm2go/p2.Fn8286
func Fn8286(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8302 github.com/goccy/pythonwasm2go/p2.Fn8302
func Fn8302(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8306 github.com/goccy/pythonwasm2go/p2.Fn8306
func Fn8306(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8324 github.com/goccy/pythonwasm2go/p2.Fn8324
func Fn8324(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int64

//go:linkname Fn8330 github.com/goccy/pythonwasm2go/p2.Fn8330
func Fn8330(m *base.Module, l0 int64) int32

//go:linkname Fn8375 github.com/goccy/pythonwasm2go/p2.Fn8375
func Fn8375(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8418 github.com/goccy/pythonwasm2go/p2.Fn8418
func Fn8418(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8421 github.com/goccy/pythonwasm2go/p2.Fn8421
func Fn8421(m *base.Module, l0 int32) int32

//go:linkname Fn8422 github.com/goccy/pythonwasm2go/p2.Fn8422
func Fn8422(m *base.Module, l0 int32) int32

//go:linkname Fn8424 github.com/goccy/pythonwasm2go/p2.Fn8424
func Fn8424(m *base.Module, l0 int32) int32

//go:linkname Fn8427 github.com/goccy/pythonwasm2go/p2.Fn8427
func Fn8427(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8428 github.com/goccy/pythonwasm2go/p2.Fn8428
func Fn8428(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8435 github.com/goccy/pythonwasm2go/p2.Fn8435
func Fn8435(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8441 github.com/goccy/pythonwasm2go/p2.Fn8441
func Fn8441(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8446 github.com/goccy/pythonwasm2go/p2.Fn8446
func Fn8446(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8448 github.com/goccy/pythonwasm2go/p2.Fn8448
func Fn8448(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8494 github.com/goccy/pythonwasm2go/p2.Fn8494
func Fn8494(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8495 github.com/goccy/pythonwasm2go/p2.Fn8495
func Fn8495(m *base.Module, l0 int32, l1 int32) int64

//go:linkname Fn8504 github.com/goccy/pythonwasm2go/p2.Fn8504
func Fn8504(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8505 github.com/goccy/pythonwasm2go/p2.Fn8505
func Fn8505(m *base.Module, l0 int32) int32

//go:linkname Fn8506 github.com/goccy/pythonwasm2go/p2.Fn8506
func Fn8506(m *base.Module, l0 int32) int32

//go:linkname Fn8507 github.com/goccy/pythonwasm2go/p2.Fn8507
func Fn8507(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8508 github.com/goccy/pythonwasm2go/p2.Fn8508
func Fn8508(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8509 github.com/goccy/pythonwasm2go/p2.Fn8509
func Fn8509(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8510 github.com/goccy/pythonwasm2go/p2.Fn8510
func Fn8510(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8522 github.com/goccy/pythonwasm2go/p2.Fn8522
func Fn8522(m *base.Module, l0 int32) int32

//go:linkname Fn8536 github.com/goccy/pythonwasm2go/p2.Fn8536
func Fn8536(m *base.Module, l0 int32) int32

//go:linkname Fn8540 github.com/goccy/pythonwasm2go/p2.Fn8540
func Fn8540(m *base.Module, l0 int32) int32

//go:linkname Fn8542 github.com/goccy/pythonwasm2go/p2.Fn8542
func Fn8542(m *base.Module, l0 int32) int32

//go:linkname Fn8543 github.com/goccy/pythonwasm2go/p2.Fn8543
func Fn8543(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8546 github.com/goccy/pythonwasm2go/p2.Fn8546
func Fn8546(m *base.Module)

//go:linkname Fn8547 github.com/goccy/pythonwasm2go/p2.Fn8547
func Fn8547(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8556 github.com/goccy/pythonwasm2go/p2.Fn8556
func Fn8556(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8572 github.com/goccy/pythonwasm2go/p2.Fn8572
func Fn8572(m *base.Module, l0 int32) int32

//go:linkname Fn8573 github.com/goccy/pythonwasm2go/p2.Fn8573
func Fn8573(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8576 github.com/goccy/pythonwasm2go/p2.Fn8576
func Fn8576(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8580 github.com/goccy/pythonwasm2go/p2.Fn8580
func Fn8580(m *base.Module, l0 int32) int32

//go:linkname Fn8597 github.com/goccy/pythonwasm2go/p2.Fn8597
func Fn8597(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8601 github.com/goccy/pythonwasm2go/p2.Fn8601
func Fn8601(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8605 github.com/goccy/pythonwasm2go/p2.Fn8605
func Fn8605(m *base.Module) int32

//go:linkname Fn8619 github.com/goccy/pythonwasm2go/p2.Fn8619
func Fn8619(m *base.Module, l0 int32) int32

//go:linkname Fn8620 github.com/goccy/pythonwasm2go/p2.Fn8620
func Fn8620(m *base.Module, l0 int32) int32

//go:linkname Fn8621 github.com/goccy/pythonwasm2go/p2.Fn8621
func Fn8621(m *base.Module, l0 int32) int32

//go:linkname Fn8622 github.com/goccy/pythonwasm2go/p2.Fn8622
func Fn8622(m *base.Module, l0 int32) int32

//go:linkname Fn8623 github.com/goccy/pythonwasm2go/p2.Fn8623
func Fn8623(m *base.Module, l0 int32) int32

//go:linkname Fn8628 github.com/goccy/pythonwasm2go/p2.Fn8628
func Fn8628(m *base.Module, l0 int32)

//go:linkname Fn8632 github.com/goccy/pythonwasm2go/p2.Fn8632
func Fn8632(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8639 github.com/goccy/pythonwasm2go/p2.Fn8639
func Fn8639(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8646 github.com/goccy/pythonwasm2go/p2.Fn8646
func Fn8646(m *base.Module, l0 int32) int64

//go:linkname Fn8675 github.com/goccy/pythonwasm2go/p2.Fn8675
func Fn8675(m *base.Module, l0 int32) int32

//go:linkname Fn8677 github.com/goccy/pythonwasm2go/p2.Fn8677
func Fn8677(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8685 github.com/goccy/pythonwasm2go/p2.Fn8685
func Fn8685(m *base.Module, l0 int32) int32

//go:linkname Fn8686 github.com/goccy/pythonwasm2go/p2.Fn8686
func Fn8686(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8689 github.com/goccy/pythonwasm2go/p2.Fn8689
func Fn8689(m *base.Module, l0 int32) int32

//go:linkname Fn8700 github.com/goccy/pythonwasm2go/p2.Fn8700
func Fn8700(m *base.Module, l0 int32) int32

//go:linkname Fn8703 github.com/goccy/pythonwasm2go/p2.Fn8703
func Fn8703(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8706 github.com/goccy/pythonwasm2go/p2.Fn8706
func Fn8706(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8712 github.com/goccy/pythonwasm2go/p2.Fn8712
func Fn8712(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8722 github.com/goccy/pythonwasm2go/p2.Fn8722
func Fn8722(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8723 github.com/goccy/pythonwasm2go/p2.Fn8723
func Fn8723(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8743 github.com/goccy/pythonwasm2go/p2.Fn8743
func Fn8743(m *base.Module, l0 int32) int32

//go:linkname Fn8744 github.com/goccy/pythonwasm2go/p2.Fn8744
func Fn8744(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8746 github.com/goccy/pythonwasm2go/p2.Fn8746
func Fn8746(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8753 github.com/goccy/pythonwasm2go/p2.Fn8753
func Fn8753(m *base.Module, l0 int32) int32

//go:linkname Fn8770 github.com/goccy/pythonwasm2go/p2.Fn8770
func Fn8770(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8771 github.com/goccy/pythonwasm2go/p2.Fn8771
func Fn8771(m *base.Module, l0 int32) int32

//go:linkname Fn8772 github.com/goccy/pythonwasm2go/p2.Fn8772
func Fn8772(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8779 github.com/goccy/pythonwasm2go/p2.Fn8779
func Fn8779(m *base.Module, l0 int32) int32

//go:linkname Fn8812 github.com/goccy/pythonwasm2go/p2.Fn8812
func Fn8812(m *base.Module, l0 int32) int32

//go:linkname Fn8813 github.com/goccy/pythonwasm2go/p2.Fn8813
func Fn8813(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8854 github.com/goccy/pythonwasm2go/p2.Fn8854
func Fn8854(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8871 github.com/goccy/pythonwasm2go/p2.Fn8871
func Fn8871(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8873 github.com/goccy/pythonwasm2go/p2.Fn8873
func Fn8873(m *base.Module, l0 int32)

//go:linkname Fn8876 github.com/goccy/pythonwasm2go/p0.Fn8876
func Fn8876(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8881 github.com/goccy/pythonwasm2go/p2.Fn8881
func Fn8881(m *base.Module, l0 int32)

//go:linkname Fn8882 github.com/goccy/pythonwasm2go/p2.Fn8882
func Fn8882(m *base.Module, l0 int32)

//go:linkname Fn8883 github.com/goccy/pythonwasm2go/p2.Fn8883
func Fn8883(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8884 github.com/goccy/pythonwasm2go/p2.Fn8884
func Fn8884(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8887 github.com/goccy/pythonwasm2go/p2.Fn8887
func Fn8887(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8906 github.com/goccy/pythonwasm2go/p2.Fn8906
func Fn8906(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8913 github.com/goccy/pythonwasm2go/p2.Fn8913
func Fn8913(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8934 github.com/goccy/pythonwasm2go/p2.Fn8934
func Fn8934(m *base.Module, l0 int32)

//go:linkname Fn8939 github.com/goccy/pythonwasm2go/p2.Fn8939
func Fn8939(m *base.Module) int32

//go:linkname Fn8941 github.com/goccy/pythonwasm2go/p2.Fn8941
func Fn8941(m *base.Module, l0 int32)

//go:linkname Fn8942 github.com/goccy/pythonwasm2go/p2.Fn8942
func Fn8942(m *base.Module, l0 int32) int64

//go:linkname Fn8944 github.com/goccy/pythonwasm2go/p2.Fn8944
func Fn8944(m *base.Module, l0 int32) int32

//go:linkname Fn8960 github.com/goccy/pythonwasm2go/p2.Fn8960
func Fn8960(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8962 github.com/goccy/pythonwasm2go/p2.Fn8962
func Fn8962(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9000 github.com/goccy/pythonwasm2go/p2.Fn9000
func Fn9000(m *base.Module, l0 int32) int32

//go:linkname Fn9027 github.com/goccy/pythonwasm2go/p2.Fn9027
func Fn9027(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9031 github.com/goccy/pythonwasm2go/p2.Fn9031
func Fn9031(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9046 github.com/goccy/pythonwasm2go/p2.Fn9046
func Fn9046(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9048 github.com/goccy/pythonwasm2go/p2.Fn9048
func Fn9048(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9054 github.com/goccy/pythonwasm2go/p2.Fn9054
func Fn9054(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9066 github.com/goccy/pythonwasm2go/p2.Fn9066
func Fn9066(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn9070 github.com/goccy/pythonwasm2go/p2.Fn9070
func Fn9070(m *base.Module, l0 int32) int32

//go:linkname Fn9077 github.com/goccy/pythonwasm2go/p2.Fn9077
func Fn9077(m *base.Module, l0 int32)

//go:linkname Fn9085 github.com/goccy/pythonwasm2go/p2.Fn9085
func Fn9085(m *base.Module, l0 int32)

//go:linkname Fn9178 github.com/goccy/pythonwasm2go/p2.Fn9178
func Fn9178(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9188 github.com/goccy/pythonwasm2go/p2.Fn9188
func Fn9188(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9189 github.com/goccy/pythonwasm2go/p2.Fn9189
func Fn9189(m *base.Module, l0 int32) int32

//go:linkname Fn9209 github.com/goccy/pythonwasm2go/p2.Fn9209
func Fn9209(m *base.Module, l0 int32) int32

//go:linkname Fn9241 github.com/goccy/pythonwasm2go/p2.Fn9241
func Fn9241(m *base.Module)

//go:linkname Fn9243 github.com/goccy/pythonwasm2go/p2.Fn9243
func Fn9243(m *base.Module, l0 int32)

//go:linkname Fn9244 github.com/goccy/pythonwasm2go/p2.Fn9244
func Fn9244(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9246 github.com/goccy/pythonwasm2go/p2.Fn9246
func Fn9246(m *base.Module, l0 int32) int32

//go:linkname Fn9247 github.com/goccy/pythonwasm2go/p2.Fn9247
func Fn9247(m *base.Module, l0 int32)

//go:linkname Fn9249 github.com/goccy/pythonwasm2go/p2.Fn9249
func Fn9249(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9251 github.com/goccy/pythonwasm2go/p2.Fn9251
func Fn9251(m *base.Module) int32

//go:linkname Fn9252 github.com/goccy/pythonwasm2go/p2.Fn9252
func Fn9252(m *base.Module, l0 int32)

//go:linkname Fn9259 github.com/goccy/pythonwasm2go/p2.Fn9259
func Fn9259(m *base.Module) int64

//go:linkname Fn9262 github.com/goccy/pythonwasm2go/p2.Fn9262
func Fn9262(m *base.Module, l0 int32, l1 int64) int32

//go:linkname Fn9263 github.com/goccy/pythonwasm2go/p2.Fn9263
func Fn9263(m *base.Module, l0 int32, l1 int64, l2 int32) int64

//go:linkname Fn9264 github.com/goccy/pythonwasm2go/p2.Fn9264
func Fn9264(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9265 github.com/goccy/pythonwasm2go/p2.Fn9265
func Fn9265(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9266 github.com/goccy/pythonwasm2go/p2.Fn9266
func Fn9266(m *base.Module, l0 int32)

//go:linkname Fn9268 github.com/goccy/pythonwasm2go/p2.Fn9268
func Fn9268(m *base.Module, l0 int32) int32

//go:linkname Fn9269 github.com/goccy/pythonwasm2go/p2.Fn9269
func Fn9269(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn9270 github.com/goccy/pythonwasm2go/p2.Fn9270
func Fn9270(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9271 github.com/goccy/pythonwasm2go/p2.Fn9271
func Fn9271(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9273 github.com/goccy/pythonwasm2go/p2.Fn9273
func Fn9273(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32

//go:linkname Fn9275 github.com/goccy/pythonwasm2go/p2.Fn9275
func Fn9275(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9276 github.com/goccy/pythonwasm2go/p2.Fn9276
func Fn9276(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9277 github.com/goccy/pythonwasm2go/p2.Fn9277
func Fn9277(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9278 github.com/goccy/pythonwasm2go/p2.Fn9278
func Fn9278(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9279 github.com/goccy/pythonwasm2go/p2.Fn9279
func Fn9279(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9280 github.com/goccy/pythonwasm2go/p2.Fn9280
func Fn9280(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9281 github.com/goccy/pythonwasm2go/p2.Fn9281
func Fn9281(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9282 github.com/goccy/pythonwasm2go/p2.Fn9282
func Fn9282(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9285 github.com/goccy/pythonwasm2go/p2.Fn9285
func Fn9285(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9288 github.com/goccy/pythonwasm2go/p2.Fn9288
func Fn9288(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9289 github.com/goccy/pythonwasm2go/p2.Fn9289
func Fn9289(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9291 github.com/goccy/pythonwasm2go/p2.Fn9291
func Fn9291(m *base.Module)

//go:linkname Fn9292 github.com/goccy/pythonwasm2go/p2.Fn9292
func Fn9292(m *base.Module)

//go:linkname Fn9294 github.com/goccy/pythonwasm2go/p2.Fn9294
func Fn9294(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9295 github.com/goccy/pythonwasm2go/p2.Fn9295
func Fn9295(m *base.Module, l0 int32) int32

//go:linkname Fn9299 github.com/goccy/pythonwasm2go/p2.Fn9299
func Fn9299(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9300 github.com/goccy/pythonwasm2go/p2.Fn9300
func Fn9300(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9301 github.com/goccy/pythonwasm2go/p2.Fn9301
func Fn9301(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9304 github.com/goccy/pythonwasm2go/p2.Fn9304
func Fn9304(m *base.Module, l0 int32) int32

//go:linkname Fn9305 github.com/goccy/pythonwasm2go/p2.Fn9305
func Fn9305(m *base.Module, l0 int32) int32

//go:linkname Fn9306 github.com/goccy/pythonwasm2go/p2.Fn9306
func Fn9306(m *base.Module, l0 int32) int32

//go:linkname Fn9307 github.com/goccy/pythonwasm2go/p2.Fn9307
func Fn9307(m *base.Module, l0 int32) int32

//go:linkname Fn9308 github.com/goccy/pythonwasm2go/p2.Fn9308
func Fn9308(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9309 github.com/goccy/pythonwasm2go/p2.Fn9309
func Fn9309(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9312 github.com/goccy/pythonwasm2go/p2.Fn9312
func Fn9312(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn9313 github.com/goccy/pythonwasm2go/p2.Fn9313
func Fn9313(m *base.Module, l0 int32) int32

//go:linkname Fn9314 github.com/goccy/pythonwasm2go/p2.Fn9314
func Fn9314(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9318 github.com/goccy/pythonwasm2go/p2.Fn9318
func Fn9318(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9321 github.com/goccy/pythonwasm2go/p2.Fn9321
func Fn9321(m *base.Module, l0 int32) int32

//go:linkname Fn9322 github.com/goccy/pythonwasm2go/p2.Fn9322
func Fn9322(m *base.Module, l0 int32) int32

//go:linkname Fn9325 github.com/goccy/pythonwasm2go/p2.Fn9325
func Fn9325(m *base.Module, l0 int32) int32

//go:linkname Fn9326 github.com/goccy/pythonwasm2go/p2.Fn9326
func Fn9326(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9327 github.com/goccy/pythonwasm2go/p2.Fn9327
func Fn9327(m *base.Module, l0 int32) int32

//go:linkname Fn9328 github.com/goccy/pythonwasm2go/p2.Fn9328
func Fn9328(m *base.Module, l0 int32) int32

//go:linkname Fn9330 github.com/goccy/pythonwasm2go/p2.Fn9330
func Fn9330(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9332 github.com/goccy/pythonwasm2go/p0.Fn9332
func Fn9332(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn9334 github.com/goccy/pythonwasm2go/p0.Fn9334
func Fn9334(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn9337 github.com/goccy/pythonwasm2go/p2.Fn9337
func Fn9337(m *base.Module, l0 int32) int32

//go:linkname Fn9338 github.com/goccy/pythonwasm2go/p2.Fn9338
func Fn9338(m *base.Module, l0 int32) int32

//go:linkname Fn9342 github.com/goccy/pythonwasm2go/p2.Fn9342
func Fn9342(m *base.Module, l0 int32) int32

//go:linkname Fn9346 github.com/goccy/pythonwasm2go/p2.Fn9346
func Fn9346(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9350 github.com/goccy/pythonwasm2go/p2.Fn9350
func Fn9350(m *base.Module, l0 float64) float64

//go:linkname Fn9358 github.com/goccy/pythonwasm2go/p2.Fn9358
func Fn9358(m *base.Module, l0 float64) float64

//go:linkname Fn9360 github.com/goccy/pythonwasm2go/p2.Fn9360
func Fn9360(m *base.Module, l0 float64) float64

//go:linkname Fn9365 github.com/goccy/pythonwasm2go/p2.Fn9365
func Fn9365(m *base.Module, l0 int32) float64

//go:linkname Fn9366 github.com/goccy/pythonwasm2go/p2.Fn9366
func Fn9366(m *base.Module, l0 int32) float64

//go:linkname Fn9372 github.com/goccy/pythonwasm2go/p2.Fn9372
func Fn9372(m *base.Module, l0 float64, l1 int32) float64

//go:linkname Fn9374 github.com/goccy/pythonwasm2go/p2.Fn9374
func Fn9374(m *base.Module, l0 float64, l1 int32) float64

//go:linkname Fn9375 github.com/goccy/pythonwasm2go/p2.Fn9375
func Fn9375(m *base.Module, l0 float64) float64

//go:linkname Fn9379 github.com/goccy/pythonwasm2go/p2.Fn9379
func Fn9379(m *base.Module, l0 float64, l1 int32) float64

//go:linkname Fn9380 github.com/goccy/pythonwasm2go/p2.Fn9380
func Fn9380(m *base.Module, l0 float64, l1 float64) float64

//go:linkname Fn9382 github.com/goccy/pythonwasm2go/p2.Fn9382
func Fn9382(m *base.Module, l0 int64) int32

//go:linkname Fn9383 github.com/goccy/pythonwasm2go/p2.Fn9383
func Fn9383(m *base.Module, l0 float64) float64

//go:linkname Fn9384 github.com/goccy/pythonwasm2go/p2.Fn9384
func Fn9384(m *base.Module, l0 float64, l1 int32) float64

//go:linkname Fn9385 github.com/goccy/pythonwasm2go/p2.Fn9385
func Fn9385(m *base.Module, l0 float64) float64

//go:linkname Fn9386 github.com/goccy/pythonwasm2go/p2.Fn9386
func Fn9386(m *base.Module, l0 float64) float64

//go:linkname Fn9388 github.com/goccy/pythonwasm2go/p2.Fn9388
func Fn9388(m *base.Module, l0 float64) float64

//go:linkname Fn9389 github.com/goccy/pythonwasm2go/p2.Fn9389
func Fn9389(m *base.Module, l0 float64) float64

//go:linkname Fn9397 github.com/goccy/pythonwasm2go/p2.Fn9397
func Fn9397(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9399 github.com/goccy/pythonwasm2go/p2.Fn9399
func Fn9399(m *base.Module, l0 int32)

//go:linkname Fn9400 github.com/goccy/pythonwasm2go/p2.Fn9400
func Fn9400(m *base.Module, l0 int32)

//go:linkname Fn9403 github.com/goccy/pythonwasm2go/p2.Fn9403
func Fn9403(m *base.Module, l0 int32) int32

//go:linkname Fn9404 github.com/goccy/pythonwasm2go/p2.Fn9404
func Fn9404(m *base.Module, l0 int32) int32

//go:linkname Fn9405 github.com/goccy/pythonwasm2go/p2.Fn9405
func Fn9405(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9410 github.com/goccy/pythonwasm2go/p2.Fn9410
func Fn9410(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9411 github.com/goccy/pythonwasm2go/p2.Fn9411
func Fn9411(m *base.Module, l0 int32) int32

//go:linkname Fn9413 github.com/goccy/pythonwasm2go/p2.Fn9413
func Fn9413(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9414 github.com/goccy/pythonwasm2go/p2.Fn9414
func Fn9414(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9415 github.com/goccy/pythonwasm2go/p2.Fn9415
func Fn9415(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9416 github.com/goccy/pythonwasm2go/p2.Fn9416
func Fn9416(m *base.Module, l0 int32) int32

//go:linkname Fn9418 github.com/goccy/pythonwasm2go/p2.Fn9418
func Fn9418(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9420 github.com/goccy/pythonwasm2go/p2.Fn9420
func Fn9420(m *base.Module, l0 int32) int32

//go:linkname Fn9421 github.com/goccy/pythonwasm2go/p2.Fn9421
func Fn9421(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9424 github.com/goccy/pythonwasm2go/p2.Fn9424
func Fn9424(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9425 github.com/goccy/pythonwasm2go/p2.Fn9425
func Fn9425(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9426 github.com/goccy/pythonwasm2go/p2.Fn9426
func Fn9426(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9431 github.com/goccy/pythonwasm2go/p2.Fn9431
func Fn9431(m *base.Module, l0 int32)

//go:linkname Fn9432 github.com/goccy/pythonwasm2go/p2.Fn9432
func Fn9432(m *base.Module, l0 int32) int32

//go:linkname Fn9434 github.com/goccy/pythonwasm2go/p2.Fn9434
func Fn9434(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9435 github.com/goccy/pythonwasm2go/p2.Fn9435
func Fn9435(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9436 github.com/goccy/pythonwasm2go/p2.Fn9436
func Fn9436(m *base.Module, l0 int32) int32

//go:linkname Fn9439 github.com/goccy/pythonwasm2go/p2.Fn9439
func Fn9439(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9440 github.com/goccy/pythonwasm2go/p2.Fn9440
func Fn9440(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9441 github.com/goccy/pythonwasm2go/p2.Fn9441
func Fn9441(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9443 github.com/goccy/pythonwasm2go/p2.Fn9443
func Fn9443(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9445 github.com/goccy/pythonwasm2go/p2.Fn9445
func Fn9445(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9446 github.com/goccy/pythonwasm2go/p2.Fn9446
func Fn9446(m *base.Module, l0 int32) int32

//go:linkname Fn9447 github.com/goccy/pythonwasm2go/p2.Fn9447
func Fn9447(m *base.Module, l0 int32) int32

//go:linkname Fn9448 github.com/goccy/pythonwasm2go/p2.Fn9448
func Fn9448(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9450 github.com/goccy/pythonwasm2go/p2.Fn9450
func Fn9450(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9452 github.com/goccy/pythonwasm2go/p2.Fn9452
func Fn9452(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9453 github.com/goccy/pythonwasm2go/p2.Fn9453
func Fn9453(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9454 github.com/goccy/pythonwasm2go/p2.Fn9454
func Fn9454(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9455 github.com/goccy/pythonwasm2go/p2.Fn9455
func Fn9455(m *base.Module, l0 int32) int32

//go:linkname Fn9456 github.com/goccy/pythonwasm2go/p2.Fn9456
func Fn9456(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9457 github.com/goccy/pythonwasm2go/p2.Fn9457
func Fn9457(m *base.Module, l0 int32) int32

//go:linkname Fn9460 github.com/goccy/pythonwasm2go/p2.Fn9460
func Fn9460(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9463 github.com/goccy/pythonwasm2go/p2.Fn9463
func Fn9463(m *base.Module, l0 int32) int32

//go:linkname Fn9464 github.com/goccy/pythonwasm2go/p2.Fn9464
func Fn9464(m *base.Module, l0 int32) int32

//go:linkname Fn9469 github.com/goccy/pythonwasm2go/p0.Fn9469
func Fn9469(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn9472 github.com/goccy/pythonwasm2go/p0.Fn9472
func Fn9472(m *base.Module, l0 int32) int32

//go:linkname Fn9474 github.com/goccy/pythonwasm2go/p2.Fn9474
func Fn9474(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9477 github.com/goccy/pythonwasm2go/p2.Fn9477
func Fn9477(m *base.Module, l0 int32, l1 int64, l2 int64)
