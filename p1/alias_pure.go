//go:build (!amd64 && !arm64) || purego

package p1

import (
	base "github.com/goccy/pythonwasm2go/base"
	_ "unsafe"
)

//go:linkname Fn54 github.com/goccy/pythonwasm2go/p2.Fn54
func Fn54(m *base.Module, l0 int32) int32

//go:linkname Fn55 github.com/goccy/pythonwasm2go/p2.Fn55
func Fn55(m *base.Module, l0 int32) int64

//go:linkname Fn56 github.com/goccy/pythonwasm2go/p2.Fn56
func Fn56(m *base.Module, l0 int32)

//go:linkname Fn58 github.com/goccy/pythonwasm2go/p2.Fn58
func Fn58(m *base.Module, l0 int32) int64

//go:linkname Fn59 github.com/goccy/pythonwasm2go/p2.Fn59
func Fn59(m *base.Module, l0 int32)

//go:linkname Fn61 github.com/goccy/pythonwasm2go/p2.Fn61
func Fn61(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn62 github.com/goccy/pythonwasm2go/p2.Fn62
func Fn62(m *base.Module, l0 int32, l1 int64)

//go:linkname Fn65 github.com/goccy/pythonwasm2go/p2.Fn65
func Fn65(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn66 github.com/goccy/pythonwasm2go/p2.Fn66
func Fn66(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn67 github.com/goccy/pythonwasm2go/p2.Fn67
func Fn67(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn68 github.com/goccy/pythonwasm2go/p2.Fn68
func Fn68(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn78 github.com/goccy/pythonwasm2go/p2.Fn78
func Fn78(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn79 github.com/goccy/pythonwasm2go/p2.Fn79
func Fn79(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn80 github.com/goccy/pythonwasm2go/p2.Fn80
func Fn80(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn82 github.com/goccy/pythonwasm2go/p2.Fn82
func Fn82(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn83 github.com/goccy/pythonwasm2go/p2.Fn83
func Fn83(m *base.Module, l0 int32)

//go:linkname Fn85 github.com/goccy/pythonwasm2go/p2.Fn85
func Fn85(m *base.Module, l0 int32)

//go:linkname Fn86 github.com/goccy/pythonwasm2go/p2.Fn86
func Fn86(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn91 github.com/goccy/pythonwasm2go/p0.Fn91
func Fn91(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn92 github.com/goccy/pythonwasm2go/p2.Fn92
func Fn92(m *base.Module, l0 int32) int32

//go:linkname Fn95 github.com/goccy/pythonwasm2go/p0.Fn95
func Fn95(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn100 github.com/goccy/pythonwasm2go/p0.Fn100
func Fn100(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn103 github.com/goccy/pythonwasm2go/p2.Fn103
func Fn103(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn106 github.com/goccy/pythonwasm2go/p0.Fn106
func Fn106(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn108 github.com/goccy/pythonwasm2go/p0.Fn108
func Fn108(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn109 github.com/goccy/pythonwasm2go/p0.Fn109
func Fn109(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn111 github.com/goccy/pythonwasm2go/p2.Fn111
func Fn111(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn114 github.com/goccy/pythonwasm2go/p0.Fn114
func Fn114(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn117 github.com/goccy/pythonwasm2go/p2.Fn117
func Fn117(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn119 github.com/goccy/pythonwasm2go/p0.Fn119
func Fn119(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn121 github.com/goccy/pythonwasm2go/p0.Fn121
func Fn121(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn123 github.com/goccy/pythonwasm2go/p0.Fn123
func Fn123(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn125 github.com/goccy/pythonwasm2go/p0.Fn125
func Fn125(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn126 github.com/goccy/pythonwasm2go/p2.Fn126
func Fn126(m *base.Module, l0 int32) int32

//go:linkname Fn135 github.com/goccy/pythonwasm2go/p2.Fn135
func Fn135(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn136 github.com/goccy/pythonwasm2go/p2.Fn136
func Fn136(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn140 github.com/goccy/pythonwasm2go/p2.Fn140
func Fn140(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn141 github.com/goccy/pythonwasm2go/p2.Fn141
func Fn141(m *base.Module, l0 float64, l1 float64)

//go:linkname Fn142 github.com/goccy/pythonwasm2go/p2.Fn142
func Fn142(m *base.Module, l0 int32) float64

//go:linkname Fn143 github.com/goccy/pythonwasm2go/p0.Fn143
func Fn143(m *base.Module, l0 int32) int32

//go:linkname Fn145 github.com/goccy/pythonwasm2go/p2.Fn145
func Fn145(m *base.Module, l0 int32) float64

//go:linkname Fn146 github.com/goccy/pythonwasm2go/p2.Fn146
func Fn146(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn147 github.com/goccy/pythonwasm2go/p0.Fn147
func Fn147(m *base.Module, l0 int32) int32

//go:linkname Fn148 github.com/goccy/pythonwasm2go/p2.Fn148
func Fn148(m *base.Module, l0 int32) float64

//go:linkname Fn149 github.com/goccy/pythonwasm2go/p0.Fn149
func Fn149(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn153 github.com/goccy/pythonwasm2go/p2.Fn153
func Fn153(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn157 github.com/goccy/pythonwasm2go/p2.Fn157
func Fn157(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn161 github.com/goccy/pythonwasm2go/p2.Fn161
func Fn161(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn186 github.com/goccy/pythonwasm2go/p2.Fn186
func Fn186(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn187 github.com/goccy/pythonwasm2go/p2.Fn187
func Fn187(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn188 github.com/goccy/pythonwasm2go/p2.Fn188
func Fn188(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn189 github.com/goccy/pythonwasm2go/p2.Fn189
func Fn189(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn201 github.com/goccy/pythonwasm2go/p2.Fn201
func Fn201(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn202 github.com/goccy/pythonwasm2go/p2.Fn202
func Fn202(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn205 github.com/goccy/pythonwasm2go/p2.Fn205
func Fn205(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn206 github.com/goccy/pythonwasm2go/p0.Fn206
func Fn206(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn207 github.com/goccy/pythonwasm2go/p0.Fn207
func Fn207(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn209 github.com/goccy/pythonwasm2go/p0.Fn209
func Fn209(m *base.Module, l0 int32) int32

//go:linkname Fn210 github.com/goccy/pythonwasm2go/p0.Fn210
func Fn210(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn211 github.com/goccy/pythonwasm2go/p0.Fn211
func Fn211(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn213 github.com/goccy/pythonwasm2go/p2.Fn213
func Fn213(m *base.Module, l0 int32)

//go:linkname Fn214 github.com/goccy/pythonwasm2go/p2.Fn214
func Fn214(m *base.Module, l0 int32)

//go:linkname Fn215 github.com/goccy/pythonwasm2go/p0.Fn215
func Fn215(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn217 github.com/goccy/pythonwasm2go/p0.Fn217
func Fn217(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn219 github.com/goccy/pythonwasm2go/p2.Fn219
func Fn219(m *base.Module, l0 int32) int32

//go:linkname Fn221 github.com/goccy/pythonwasm2go/p2.Fn221
func Fn221(m *base.Module, l0 int32) int32

//go:linkname Fn222 github.com/goccy/pythonwasm2go/p2.Fn222
func Fn222(m *base.Module, l0 int32) int32

//go:linkname Fn223 github.com/goccy/pythonwasm2go/p2.Fn223
func Fn223(m *base.Module, l0 int32) int32

//go:linkname Fn225 github.com/goccy/pythonwasm2go/p0.Fn225
func Fn225(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn227 github.com/goccy/pythonwasm2go/p2.Fn227
func Fn227(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn228 github.com/goccy/pythonwasm2go/p2.Fn228
func Fn228(m *base.Module, l0 int32) int32

//go:linkname Fn229 github.com/goccy/pythonwasm2go/p2.Fn229
func Fn229(m *base.Module, l0 int32) int32

//go:linkname Fn230 github.com/goccy/pythonwasm2go/p2.Fn230
func Fn230(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn232 github.com/goccy/pythonwasm2go/p2.Fn232
func Fn232(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn235 github.com/goccy/pythonwasm2go/p2.Fn235
func Fn235(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn253 github.com/goccy/pythonwasm2go/p2.Fn253
func Fn253(m *base.Module, l0 int32) int32

//go:linkname Fn263 github.com/goccy/pythonwasm2go/p2.Fn263
func Fn263(m *base.Module, l0 int32) int32

//go:linkname Fn264 github.com/goccy/pythonwasm2go/p2.Fn264
func Fn264(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn288 github.com/goccy/pythonwasm2go/p2.Fn288
func Fn288(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn289 github.com/goccy/pythonwasm2go/p2.Fn289
func Fn289(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn310 github.com/goccy/pythonwasm2go/p2.Fn310
func Fn310(m *base.Module, l0 int32) int32

//go:linkname Fn311 github.com/goccy/pythonwasm2go/p0.Fn311
func Fn311(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn312 github.com/goccy/pythonwasm2go/p2.Fn312
func Fn312(m *base.Module, l0 int32) int32

//go:linkname Fn313 github.com/goccy/pythonwasm2go/p0.Fn313
func Fn313(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn314 github.com/goccy/pythonwasm2go/p0.Fn314
func Fn314(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn315 github.com/goccy/pythonwasm2go/p0.Fn315
func Fn315(m *base.Module, l0 int32) int32

//go:linkname Fn323 github.com/goccy/pythonwasm2go/p2.Fn323
func Fn323(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn324 github.com/goccy/pythonwasm2go/p2.Fn324
func Fn324(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn331 github.com/goccy/pythonwasm2go/p2.Fn331
func Fn331(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn348 github.com/goccy/pythonwasm2go/p2.Fn348
func Fn348(m *base.Module, l0 int32) int32

//go:linkname Fn349 github.com/goccy/pythonwasm2go/p2.Fn349
func Fn349(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn375 github.com/goccy/pythonwasm2go/p2.Fn375
func Fn375(m *base.Module, l0 int32) int32

//go:linkname Fn380 github.com/goccy/pythonwasm2go/p2.Fn380
func Fn380(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn381 github.com/goccy/pythonwasm2go/p2.Fn381
func Fn381(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn404 github.com/goccy/pythonwasm2go/p2.Fn404
func Fn404(m *base.Module, l0 int32) int32

//go:linkname Fn405 github.com/goccy/pythonwasm2go/p0.Fn405
func Fn405(m *base.Module)

//go:linkname Fn406 github.com/goccy/pythonwasm2go/p0.Fn406
func Fn406(m *base.Module, l0 int32) int32

//go:linkname Fn408 github.com/goccy/pythonwasm2go/p0.Fn408
func Fn408(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn410 github.com/goccy/pythonwasm2go/p0.Fn410
func Fn410(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn411 github.com/goccy/pythonwasm2go/p0.Fn411
func Fn411(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn412 github.com/goccy/pythonwasm2go/p0.Fn412
func Fn412(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn413 github.com/goccy/pythonwasm2go/p0.Fn413
func Fn413(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn414 github.com/goccy/pythonwasm2go/p0.Fn414
func Fn414(m *base.Module, l0 int32) int32

//go:linkname Fn415 github.com/goccy/pythonwasm2go/p0.Fn415
func Fn415(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn416 github.com/goccy/pythonwasm2go/p0.Fn416
func Fn416(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn418 github.com/goccy/pythonwasm2go/p2.Fn418
func Fn418(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn420 github.com/goccy/pythonwasm2go/p2.Fn420
func Fn420(m *base.Module, l0 int32) int32

//go:linkname Fn421 github.com/goccy/pythonwasm2go/p2.Fn421
func Fn421(m *base.Module, l0 int32)

//go:linkname Fn422 github.com/goccy/pythonwasm2go/p0.Fn422
func Fn422(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn423 github.com/goccy/pythonwasm2go/p2.Fn423
func Fn423(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn425 github.com/goccy/pythonwasm2go/p2.Fn425
func Fn425(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn426 github.com/goccy/pythonwasm2go/p2.Fn426
func Fn426(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn429 github.com/goccy/pythonwasm2go/p0.Fn429
func Fn429(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn431 github.com/goccy/pythonwasm2go/p2.Fn431
func Fn431(m *base.Module, l0 int32) int32

//go:linkname Fn437 github.com/goccy/pythonwasm2go/p2.Fn437
func Fn437(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn438 github.com/goccy/pythonwasm2go/p2.Fn438
func Fn438(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn439 github.com/goccy/pythonwasm2go/p2.Fn439
func Fn439(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn440 github.com/goccy/pythonwasm2go/p2.Fn440
func Fn440(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn442 github.com/goccy/pythonwasm2go/p2.Fn442
func Fn442(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn443 github.com/goccy/pythonwasm2go/p2.Fn443
func Fn443(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn446 github.com/goccy/pythonwasm2go/p2.Fn446
func Fn446(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn447 github.com/goccy/pythonwasm2go/p2.Fn447
func Fn447(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn449 github.com/goccy/pythonwasm2go/p2.Fn449
func Fn449(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn454 github.com/goccy/pythonwasm2go/p2.Fn454
func Fn454(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn466 github.com/goccy/pythonwasm2go/p2.Fn466
func Fn466(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn468 github.com/goccy/pythonwasm2go/p2.Fn468
func Fn468(m *base.Module, l0 int32) int32

//go:linkname Fn471 github.com/goccy/pythonwasm2go/p2.Fn471
func Fn471(m *base.Module, l0 int32) int32

//go:linkname Fn472 github.com/goccy/pythonwasm2go/p2.Fn472
func Fn472(m *base.Module, l0 int32) int32

//go:linkname Fn473 github.com/goccy/pythonwasm2go/p2.Fn473
func Fn473(m *base.Module, l0 int32) int32

//go:linkname Fn476 github.com/goccy/pythonwasm2go/p2.Fn476
func Fn476(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn477 github.com/goccy/pythonwasm2go/p2.Fn477
func Fn477(m *base.Module, l0 int32) int32

//go:linkname Fn478 github.com/goccy/pythonwasm2go/p0.Fn478
func Fn478(m *base.Module, l0 int32) int32

//go:linkname Fn479 github.com/goccy/pythonwasm2go/p2.Fn479
func Fn479(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn480 github.com/goccy/pythonwasm2go/p2.Fn480
func Fn480(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn482 github.com/goccy/pythonwasm2go/p0.Fn482
func Fn482(m *base.Module, l0 int32) int32

//go:linkname Fn483 github.com/goccy/pythonwasm2go/p0.Fn483
func Fn483(m *base.Module, l0 int32) int32

//go:linkname Fn484 github.com/goccy/pythonwasm2go/p0.Fn484
func Fn484(m *base.Module, l0 int32) int32

//go:linkname Fn486 github.com/goccy/pythonwasm2go/p0.Fn486
func Fn486(m *base.Module, l0 int32) int32

//go:linkname Fn487 github.com/goccy/pythonwasm2go/p0.Fn487
func Fn487(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn490 github.com/goccy/pythonwasm2go/p0.Fn490
func Fn490(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn491 github.com/goccy/pythonwasm2go/p2.Fn491
func Fn491(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn492 github.com/goccy/pythonwasm2go/p2.Fn492
func Fn492(m *base.Module, l0 int32) int32

//go:linkname Fn493 github.com/goccy/pythonwasm2go/p2.Fn493
func Fn493(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn495 github.com/goccy/pythonwasm2go/p2.Fn495
func Fn495(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn496 github.com/goccy/pythonwasm2go/p0.Fn496
func Fn496(m *base.Module, l0 int32) int32

//go:linkname Fn498 github.com/goccy/pythonwasm2go/p2.Fn498
func Fn498(m *base.Module, l0 int32) int32

//go:linkname Fn500 github.com/goccy/pythonwasm2go/p0.Fn500
func Fn500(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn504 github.com/goccy/pythonwasm2go/p0.Fn504
func Fn504(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn509 github.com/goccy/pythonwasm2go/p2.Fn509
func Fn509(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn510 github.com/goccy/pythonwasm2go/p2.Fn510
func Fn510(m *base.Module, l0 int32) int32

//go:linkname Fn511 github.com/goccy/pythonwasm2go/p2.Fn511
func Fn511(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn513 github.com/goccy/pythonwasm2go/p2.Fn513
func Fn513(m *base.Module, l0 int32) int32

//go:linkname Fn514 github.com/goccy/pythonwasm2go/p0.Fn514
func Fn514(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn515 github.com/goccy/pythonwasm2go/p0.Fn515
func Fn515(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn516 github.com/goccy/pythonwasm2go/p2.Fn516
func Fn516(m *base.Module, l0 int32) int32

//go:linkname Fn517 github.com/goccy/pythonwasm2go/p2.Fn517
func Fn517(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn518 github.com/goccy/pythonwasm2go/p2.Fn518
func Fn518(m *base.Module, l0 int32) int32

//go:linkname Fn519 github.com/goccy/pythonwasm2go/p2.Fn519
func Fn519(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn520 github.com/goccy/pythonwasm2go/p2.Fn520
func Fn520(m *base.Module, l0 int32) int32

//go:linkname Fn521 github.com/goccy/pythonwasm2go/p2.Fn521
func Fn521(m *base.Module, l0 int32) int32

//go:linkname Fn523 github.com/goccy/pythonwasm2go/p2.Fn523
func Fn523(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn526 github.com/goccy/pythonwasm2go/p2.Fn526
func Fn526(m *base.Module, l0 int32)

//go:linkname Fn530 github.com/goccy/pythonwasm2go/p2.Fn530
func Fn530(m *base.Module, l0 int32) int32

//go:linkname Fn532 github.com/goccy/pythonwasm2go/p2.Fn532
func Fn532(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn533 github.com/goccy/pythonwasm2go/p2.Fn533
func Fn533(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn535 github.com/goccy/pythonwasm2go/p0.Fn535
func Fn535(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn536 github.com/goccy/pythonwasm2go/p0.Fn536
func Fn536(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn537 github.com/goccy/pythonwasm2go/p2.Fn537
func Fn537(m *base.Module, l0 int32) int32

//go:linkname Fn538 github.com/goccy/pythonwasm2go/p0.Fn538
func Fn538(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn539 github.com/goccy/pythonwasm2go/p0.Fn539
func Fn539(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn540 github.com/goccy/pythonwasm2go/p2.Fn540
func Fn540(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn542 github.com/goccy/pythonwasm2go/p2.Fn542
func Fn542(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn543 github.com/goccy/pythonwasm2go/p0.Fn543
func Fn543(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn544 github.com/goccy/pythonwasm2go/p0.Fn544
func Fn544(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn545 github.com/goccy/pythonwasm2go/p2.Fn545
func Fn545(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn546 github.com/goccy/pythonwasm2go/p0.Fn546
func Fn546(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn547 github.com/goccy/pythonwasm2go/p0.Fn547
func Fn547(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn548 github.com/goccy/pythonwasm2go/p0.Fn548
func Fn548(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn549 github.com/goccy/pythonwasm2go/p0.Fn549
func Fn549(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn550 github.com/goccy/pythonwasm2go/p0.Fn550
func Fn550(m *base.Module, l0 int32) int32

//go:linkname Fn552 github.com/goccy/pythonwasm2go/p0.Fn552
func Fn552(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn578 github.com/goccy/pythonwasm2go/p0.Fn578
func Fn578(m *base.Module, l0 int32) int32

//go:linkname Fn600 github.com/goccy/pythonwasm2go/p2.Fn600
func Fn600(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn624 github.com/goccy/pythonwasm2go/p2.Fn624
func Fn624(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn657 github.com/goccy/pythonwasm2go/p2.Fn657
func Fn657(m *base.Module, l0 int32)

//go:linkname Fn658 github.com/goccy/pythonwasm2go/p2.Fn658
func Fn658(m *base.Module, l0 int32) int32

//go:linkname Fn660 github.com/goccy/pythonwasm2go/p2.Fn660
func Fn660(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn661 github.com/goccy/pythonwasm2go/p2.Fn661
func Fn661(m *base.Module, l0 int32) int32

//go:linkname Fn664 github.com/goccy/pythonwasm2go/p2.Fn664
func Fn664(m *base.Module, l0 int32)

//go:linkname Fn670 github.com/goccy/pythonwasm2go/p2.Fn670
func Fn670(m *base.Module, l0 int32) int32

//go:linkname Fn672 github.com/goccy/pythonwasm2go/p2.Fn672
func Fn672(m *base.Module, l0 int32) int32

//go:linkname Fn686 github.com/goccy/pythonwasm2go/p2.Fn686
func Fn686(m *base.Module, l0 int32) int32

//go:linkname Fn692 github.com/goccy/pythonwasm2go/p2.Fn692
func Fn692(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn693 github.com/goccy/pythonwasm2go/p2.Fn693
func Fn693(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn723 github.com/goccy/pythonwasm2go/p2.Fn723
func Fn723(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn729 github.com/goccy/pythonwasm2go/p0.Fn729
func Fn729(m *base.Module, l0 float64) int32

//go:linkname Fn735 github.com/goccy/pythonwasm2go/p0.Fn735
func Fn735(m *base.Module, l0 int32) float64

//go:linkname Fn736 github.com/goccy/pythonwasm2go/p2.Fn736
func Fn736(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn738 github.com/goccy/pythonwasm2go/p2.Fn738
func Fn738(m *base.Module, l0 int32) int32

//go:linkname Fn747 github.com/goccy/pythonwasm2go/p2.Fn747
func Fn747(m *base.Module, l0 int32, l1 int32) float64

//go:linkname Fn781 github.com/goccy/pythonwasm2go/p2.Fn781
func Fn781(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn793 github.com/goccy/pythonwasm2go/p0.Fn793
func Fn793(m *base.Module, l0 int32) int32

//go:linkname Fn794 github.com/goccy/pythonwasm2go/p0.Fn794
func Fn794(m *base.Module, l0 int32) int32

//go:linkname Fn795 github.com/goccy/pythonwasm2go/p0.Fn795
func Fn795(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn796 github.com/goccy/pythonwasm2go/p0.Fn796
func Fn796(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn797 github.com/goccy/pythonwasm2go/p2.Fn797
func Fn797(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn798 github.com/goccy/pythonwasm2go/p2.Fn798
func Fn798(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn799 github.com/goccy/pythonwasm2go/p2.Fn799
func Fn799(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn801 github.com/goccy/pythonwasm2go/p0.Fn801
func Fn801(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn803 github.com/goccy/pythonwasm2go/p0.Fn803
func Fn803(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn806 github.com/goccy/pythonwasm2go/p0.Fn806
func Fn806(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn807 github.com/goccy/pythonwasm2go/p0.Fn807
func Fn807(m *base.Module, l0 int32) int32

//go:linkname Fn808 github.com/goccy/pythonwasm2go/p0.Fn808
func Fn808(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn809 github.com/goccy/pythonwasm2go/p0.Fn809
func Fn809(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn813 github.com/goccy/pythonwasm2go/p2.Fn813
func Fn813(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn818 github.com/goccy/pythonwasm2go/p2.Fn818
func Fn818(m *base.Module, l0 int32)

//go:linkname Fn819 github.com/goccy/pythonwasm2go/p2.Fn819
func Fn819(m *base.Module, l0 int32) int32

//go:linkname Fn831 github.com/goccy/pythonwasm2go/p2.Fn831
func Fn831(m *base.Module, l0 int32) int32

//go:linkname Fn832 github.com/goccy/pythonwasm2go/p0.Fn832
func Fn832(m *base.Module, l0 int32) int32

//go:linkname Fn848 github.com/goccy/pythonwasm2go/p2.Fn848
func Fn848(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn856 github.com/goccy/pythonwasm2go/p2.Fn856
func Fn856(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn862 github.com/goccy/pythonwasm2go/p2.Fn862
func Fn862(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn882 github.com/goccy/pythonwasm2go/p0.Fn882
func Fn882(m *base.Module, l0 int32) int32

//go:linkname Fn884 github.com/goccy/pythonwasm2go/p2.Fn884
func Fn884(m *base.Module, l0 int32) int32

//go:linkname Fn885 github.com/goccy/pythonwasm2go/p0.Fn885
func Fn885(m *base.Module, l0 int32) int32

//go:linkname Fn887 github.com/goccy/pythonwasm2go/p0.Fn887
func Fn887(m *base.Module, l0 int32) int32

//go:linkname Fn888 github.com/goccy/pythonwasm2go/p0.Fn888
func Fn888(m *base.Module, l0 int64) int32

//go:linkname Fn890 github.com/goccy/pythonwasm2go/p0.Fn890
func Fn890(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn891 github.com/goccy/pythonwasm2go/p0.Fn891
func Fn891(m *base.Module, l0 int32) int32

//go:linkname Fn892 github.com/goccy/pythonwasm2go/p0.Fn892
func Fn892(m *base.Module, l0 int32) int32

//go:linkname Fn893 github.com/goccy/pythonwasm2go/p0.Fn893
func Fn893(m *base.Module, l0 int32) int32

//go:linkname Fn894 github.com/goccy/pythonwasm2go/p2.Fn894
func Fn894(m *base.Module, l0 int32) int32

//go:linkname Fn895 github.com/goccy/pythonwasm2go/p2.Fn895
func Fn895(m *base.Module, l0 int32) int32

//go:linkname Fn896 github.com/goccy/pythonwasm2go/p0.Fn896
func Fn896(m *base.Module, l0 int32) int32

//go:linkname Fn899 github.com/goccy/pythonwasm2go/p2.Fn899
func Fn899(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn900 github.com/goccy/pythonwasm2go/p2.Fn900
func Fn900(m *base.Module, l0 int32) int64

//go:linkname Fn902 github.com/goccy/pythonwasm2go/p2.Fn902
func Fn902(m *base.Module, l0 int32) int32

//go:linkname Fn903 github.com/goccy/pythonwasm2go/p2.Fn903
func Fn903(m *base.Module, l0 int32) int32

//go:linkname Fn904 github.com/goccy/pythonwasm2go/p2.Fn904
func Fn904(m *base.Module, l0 int32)

//go:linkname Fn905 github.com/goccy/pythonwasm2go/p0.Fn905
func Fn905(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn907 github.com/goccy/pythonwasm2go/p2.Fn907
func Fn907(m *base.Module, l0 int32) int32

//go:linkname Fn908 github.com/goccy/pythonwasm2go/p0.Fn908
func Fn908(m *base.Module, l0 int64) int32

//go:linkname Fn909 github.com/goccy/pythonwasm2go/p0.Fn909
func Fn909(m *base.Module, l0 int32) int64

//go:linkname Fn910 github.com/goccy/pythonwasm2go/p2.Fn910
func Fn910(m *base.Module, l0 int32) int64

//go:linkname Fn911 github.com/goccy/pythonwasm2go/p0.Fn911
func Fn911(m *base.Module, l0 int32) int64

//go:linkname Fn914 github.com/goccy/pythonwasm2go/p2.Fn914
func Fn914(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn915 github.com/goccy/pythonwasm2go/p2.Fn915
func Fn915(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn916 github.com/goccy/pythonwasm2go/p2.Fn916
func Fn916(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn917 github.com/goccy/pythonwasm2go/p2.Fn917
func Fn917(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn918 github.com/goccy/pythonwasm2go/p2.Fn918
func Fn918(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn921 github.com/goccy/pythonwasm2go/p2.Fn921
func Fn921(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn922 github.com/goccy/pythonwasm2go/p0.Fn922
func Fn922(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn923 github.com/goccy/pythonwasm2go/p2.Fn923
func Fn923(m *base.Module, l0 int32)

//go:linkname Fn924 github.com/goccy/pythonwasm2go/p2.Fn924
func Fn924(m *base.Module, l0 int64) int32

//go:linkname Fn925 github.com/goccy/pythonwasm2go/p2.Fn925
func Fn925(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn926 github.com/goccy/pythonwasm2go/p2.Fn926
func Fn926(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn928 github.com/goccy/pythonwasm2go/p2.Fn928
func Fn928(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn929 github.com/goccy/pythonwasm2go/p2.Fn929
func Fn929(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn930 github.com/goccy/pythonwasm2go/p0.Fn930
func Fn930(m *base.Module, l0 int32) float64

//go:linkname Fn933 github.com/goccy/pythonwasm2go/p2.Fn933
func Fn933(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn936 github.com/goccy/pythonwasm2go/p2.Fn936
func Fn936(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn940 github.com/goccy/pythonwasm2go/p2.Fn940
func Fn940(m *base.Module, l0 int32) int32

//go:linkname Fn941 github.com/goccy/pythonwasm2go/p2.Fn941
func Fn941(m *base.Module, l0 int32, l1 int64) int32

//go:linkname Fn944 github.com/goccy/pythonwasm2go/p2.Fn944
func Fn944(m *base.Module, l0 int32) int32

//go:linkname Fn945 github.com/goccy/pythonwasm2go/p2.Fn945
func Fn945(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn947 github.com/goccy/pythonwasm2go/p2.Fn947
func Fn947(m *base.Module, l0 int32) int32

//go:linkname Fn948 github.com/goccy/pythonwasm2go/p2.Fn948
func Fn948(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn952 github.com/goccy/pythonwasm2go/p2.Fn952
func Fn952(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn962 github.com/goccy/pythonwasm2go/p2.Fn962
func Fn962(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn963 github.com/goccy/pythonwasm2go/p2.Fn963
func Fn963(m *base.Module, l0 int32)

//go:linkname Fn964 github.com/goccy/pythonwasm2go/p2.Fn964
func Fn964(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn965 github.com/goccy/pythonwasm2go/p2.Fn965
func Fn965(m *base.Module, l0 int32) int32

//go:linkname Fn966 github.com/goccy/pythonwasm2go/p2.Fn966
func Fn966(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn967 github.com/goccy/pythonwasm2go/p2.Fn967
func Fn967(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn968 github.com/goccy/pythonwasm2go/p2.Fn968
func Fn968(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn975 github.com/goccy/pythonwasm2go/p0.Fn975
func Fn975(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn988 github.com/goccy/pythonwasm2go/p2.Fn988
func Fn988(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn990 github.com/goccy/pythonwasm2go/p2.Fn990
func Fn990(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1009 github.com/goccy/pythonwasm2go/p2.Fn1009
func Fn1009(m *base.Module, l0 int32) int32

//go:linkname Fn1020 github.com/goccy/pythonwasm2go/p2.Fn1020
func Fn1020(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1032 github.com/goccy/pythonwasm2go/p0.Fn1032
func Fn1032(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1052 github.com/goccy/pythonwasm2go/p2.Fn1052
func Fn1052(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1076 github.com/goccy/pythonwasm2go/p0.Fn1076
func Fn1076(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1090 github.com/goccy/pythonwasm2go/p0.Fn1090
func Fn1090(m *base.Module) int32

//go:linkname Fn1099 github.com/goccy/pythonwasm2go/p0.Fn1099
func Fn1099(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1104 github.com/goccy/pythonwasm2go/p2.Fn1104
func Fn1104(m *base.Module, l0 int32)

//go:linkname Fn1105 github.com/goccy/pythonwasm2go/p2.Fn1105
func Fn1105(m *base.Module, l0 int32)

//go:linkname Fn1107 github.com/goccy/pythonwasm2go/p2.Fn1107
func Fn1107(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1112 github.com/goccy/pythonwasm2go/p0.Fn1112
func Fn1112(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1116 github.com/goccy/pythonwasm2go/p0.Fn1116
func Fn1116(m *base.Module, l0 int32) int32

//go:linkname Fn1118 github.com/goccy/pythonwasm2go/p0.Fn1118
func Fn1118(m *base.Module, l0 int32)

//go:linkname Fn1119 github.com/goccy/pythonwasm2go/p2.Fn1119
func Fn1119(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1120 github.com/goccy/pythonwasm2go/p0.Fn1120
func Fn1120(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1121 github.com/goccy/pythonwasm2go/p0.Fn1121
func Fn1121(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1122 github.com/goccy/pythonwasm2go/p0.Fn1122
func Fn1122(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1123 github.com/goccy/pythonwasm2go/p2.Fn1123
func Fn1123(m *base.Module, l0 int32) int32

//go:linkname Fn1126 github.com/goccy/pythonwasm2go/p0.Fn1126
func Fn1126(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1127 github.com/goccy/pythonwasm2go/p0.Fn1127
func Fn1127(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1128 github.com/goccy/pythonwasm2go/p2.Fn1128
func Fn1128(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1129 github.com/goccy/pythonwasm2go/p0.Fn1129
func Fn1129(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1130 github.com/goccy/pythonwasm2go/p2.Fn1130
func Fn1130(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1131 github.com/goccy/pythonwasm2go/p2.Fn1131
func Fn1131(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1134 github.com/goccy/pythonwasm2go/p2.Fn1134
func Fn1134(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1135 github.com/goccy/pythonwasm2go/p0.Fn1135
func Fn1135(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1136 github.com/goccy/pythonwasm2go/p0.Fn1136
func Fn1136(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1139 github.com/goccy/pythonwasm2go/p2.Fn1139
func Fn1139(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1140 github.com/goccy/pythonwasm2go/p2.Fn1140
func Fn1140(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1142 github.com/goccy/pythonwasm2go/p0.Fn1142
func Fn1142(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1143 github.com/goccy/pythonwasm2go/p0.Fn1143
func Fn1143(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1144 github.com/goccy/pythonwasm2go/p2.Fn1144
func Fn1144(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1146 github.com/goccy/pythonwasm2go/p0.Fn1146
func Fn1146(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1147 github.com/goccy/pythonwasm2go/p0.Fn1147
func Fn1147(m *base.Module, l0 int32) int32

//go:linkname Fn1148 github.com/goccy/pythonwasm2go/p2.Fn1148
func Fn1148(m *base.Module, l0 int32) int32

//go:linkname Fn1151 github.com/goccy/pythonwasm2go/p0.Fn1151
func Fn1151(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1152 github.com/goccy/pythonwasm2go/p0.Fn1152
func Fn1152(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1156 github.com/goccy/pythonwasm2go/p0.Fn1156
func Fn1156(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1157 github.com/goccy/pythonwasm2go/p0.Fn1157
func Fn1157(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1158 github.com/goccy/pythonwasm2go/p2.Fn1158
func Fn1158(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1159 github.com/goccy/pythonwasm2go/p0.Fn1159
func Fn1159(m *base.Module, l0 int32) int32

//go:linkname Fn1160 github.com/goccy/pythonwasm2go/p0.Fn1160
func Fn1160(m *base.Module, l0 int32) int32

//go:linkname Fn1162 github.com/goccy/pythonwasm2go/p0.Fn1162
func Fn1162(m *base.Module, l0 int32) int32

//go:linkname Fn1163 github.com/goccy/pythonwasm2go/p0.Fn1163
func Fn1163(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1166 github.com/goccy/pythonwasm2go/p2.Fn1166
func Fn1166(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1178 github.com/goccy/pythonwasm2go/p2.Fn1178
func Fn1178(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1179 github.com/goccy/pythonwasm2go/p0.Fn1179
func Fn1179(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1180 github.com/goccy/pythonwasm2go/p0.Fn1180
func Fn1180(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1181 github.com/goccy/pythonwasm2go/p2.Fn1181
func Fn1181(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1187 github.com/goccy/pythonwasm2go/p2.Fn1187
func Fn1187(m *base.Module, l0 int32)

//go:linkname Fn1197 github.com/goccy/pythonwasm2go/p2.Fn1197
func Fn1197(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1203 github.com/goccy/pythonwasm2go/p0.Fn1203
func Fn1203(m *base.Module, l0 int32) int32

//go:linkname Fn1204 github.com/goccy/pythonwasm2go/p0.Fn1204
func Fn1204(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1206 github.com/goccy/pythonwasm2go/p2.Fn1206
func Fn1206(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1207 github.com/goccy/pythonwasm2go/p2.Fn1207
func Fn1207(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1208 github.com/goccy/pythonwasm2go/p2.Fn1208
func Fn1208(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1211 github.com/goccy/pythonwasm2go/p2.Fn1211
func Fn1211(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1213 github.com/goccy/pythonwasm2go/p2.Fn1213
func Fn1213(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1214 github.com/goccy/pythonwasm2go/p2.Fn1214
func Fn1214(m *base.Module, l0 int32)

//go:linkname Fn1215 github.com/goccy/pythonwasm2go/p2.Fn1215
func Fn1215(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1216 github.com/goccy/pythonwasm2go/p2.Fn1216
func Fn1216(m *base.Module, l0 int32)

//go:linkname Fn1248 github.com/goccy/pythonwasm2go/p2.Fn1248
func Fn1248(m *base.Module, l0 int32) int32

//go:linkname Fn1249 github.com/goccy/pythonwasm2go/p0.Fn1249
func Fn1249(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1261 github.com/goccy/pythonwasm2go/p2.Fn1261
func Fn1261(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1263 github.com/goccy/pythonwasm2go/p0.Fn1263
func Fn1263(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1268 github.com/goccy/pythonwasm2go/p2.Fn1268
func Fn1268(m *base.Module, l0 int32)

//go:linkname Fn1269 github.com/goccy/pythonwasm2go/p0.Fn1269
func Fn1269(m *base.Module, l0 int32) int32

//go:linkname Fn1270 github.com/goccy/pythonwasm2go/p2.Fn1270
func Fn1270(m *base.Module, l0 int32) int32

//go:linkname Fn1272 github.com/goccy/pythonwasm2go/p2.Fn1272
func Fn1272(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1273 github.com/goccy/pythonwasm2go/p2.Fn1273
func Fn1273(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1274 github.com/goccy/pythonwasm2go/p2.Fn1274
func Fn1274(m *base.Module, l0 int32)

//go:linkname Fn1275 github.com/goccy/pythonwasm2go/p2.Fn1275
func Fn1275(m *base.Module, l0 int32)

//go:linkname Fn1276 github.com/goccy/pythonwasm2go/p2.Fn1276
func Fn1276(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1277 github.com/goccy/pythonwasm2go/p2.Fn1277
func Fn1277(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1278 github.com/goccy/pythonwasm2go/p2.Fn1278
func Fn1278(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1285 github.com/goccy/pythonwasm2go/p2.Fn1285
func Fn1285(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1288 github.com/goccy/pythonwasm2go/p2.Fn1288
func Fn1288(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1289 github.com/goccy/pythonwasm2go/p2.Fn1289
func Fn1289(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1290 github.com/goccy/pythonwasm2go/p2.Fn1290
func Fn1290(m *base.Module) int32

//go:linkname Fn1292 github.com/goccy/pythonwasm2go/p2.Fn1292
func Fn1292(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32

//go:linkname Fn1293 github.com/goccy/pythonwasm2go/p2.Fn1293
func Fn1293(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32

//go:linkname Fn1294 github.com/goccy/pythonwasm2go/p2.Fn1294
func Fn1294(m *base.Module, l0 int32)

//go:linkname Fn1296 github.com/goccy/pythonwasm2go/p2.Fn1296
func Fn1296(m *base.Module, l0 int32) int32

//go:linkname Fn1299 github.com/goccy/pythonwasm2go/p2.Fn1299
func Fn1299(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1300 github.com/goccy/pythonwasm2go/p2.Fn1300
func Fn1300(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1302 github.com/goccy/pythonwasm2go/p2.Fn1302
func Fn1302(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)

//go:linkname Fn1303 github.com/goccy/pythonwasm2go/p2.Fn1303
func Fn1303(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1304 github.com/goccy/pythonwasm2go/p2.Fn1304
func Fn1304(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1307 github.com/goccy/pythonwasm2go/p2.Fn1307
func Fn1307(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1308 github.com/goccy/pythonwasm2go/p2.Fn1308
func Fn1308(m *base.Module, l0 int32) int32

//go:linkname Fn1309 github.com/goccy/pythonwasm2go/p2.Fn1309
func Fn1309(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1310 github.com/goccy/pythonwasm2go/p2.Fn1310
func Fn1310(m *base.Module, l0 int32) int32

//go:linkname Fn1314 github.com/goccy/pythonwasm2go/p2.Fn1314
func Fn1314(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1340 github.com/goccy/pythonwasm2go/p0.Fn1340
func Fn1340(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1360 github.com/goccy/pythonwasm2go/p2.Fn1360
func Fn1360(m *base.Module, l0 int32) int32

//go:linkname Fn1361 github.com/goccy/pythonwasm2go/p2.Fn1361
func Fn1361(m *base.Module, l0 int32) int32

//go:linkname Fn1362 github.com/goccy/pythonwasm2go/p2.Fn1362
func Fn1362(m *base.Module, l0 int32) int32

//go:linkname Fn1366 github.com/goccy/pythonwasm2go/p2.Fn1366
func Fn1366(m *base.Module, l0 int32) int32

//go:linkname Fn1367 github.com/goccy/pythonwasm2go/p2.Fn1367
func Fn1367(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1369 github.com/goccy/pythonwasm2go/p2.Fn1369
func Fn1369(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1370 github.com/goccy/pythonwasm2go/p2.Fn1370
func Fn1370(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1371 github.com/goccy/pythonwasm2go/p2.Fn1371
func Fn1371(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1372 github.com/goccy/pythonwasm2go/p0.Fn1372
func Fn1372(m *base.Module, l0 int32) int32

//go:linkname Fn1373 github.com/goccy/pythonwasm2go/p2.Fn1373
func Fn1373(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1374 github.com/goccy/pythonwasm2go/p2.Fn1374
func Fn1374(m *base.Module, l0 int32) int32

//go:linkname Fn1375 github.com/goccy/pythonwasm2go/p2.Fn1375
func Fn1375(m *base.Module, l0 int32) int32

//go:linkname Fn1376 github.com/goccy/pythonwasm2go/p2.Fn1376
func Fn1376(m *base.Module, l0 int32) int32

//go:linkname Fn1377 github.com/goccy/pythonwasm2go/p2.Fn1377
func Fn1377(m *base.Module, l0 int32) int32

//go:linkname Fn1378 github.com/goccy/pythonwasm2go/p0.Fn1378
func Fn1378(m *base.Module, l0 int32) int32

//go:linkname Fn1379 github.com/goccy/pythonwasm2go/p2.Fn1379
func Fn1379(m *base.Module, l0 int32) int32

//go:linkname Fn1381 github.com/goccy/pythonwasm2go/p2.Fn1381
func Fn1381(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1382 github.com/goccy/pythonwasm2go/p0.Fn1382
func Fn1382(m *base.Module, l0 int32) int32

//go:linkname Fn1395 github.com/goccy/pythonwasm2go/p2.Fn1395
func Fn1395(m *base.Module, l0 int32) int32

//go:linkname Fn1405 github.com/goccy/pythonwasm2go/p2.Fn1405
func Fn1405(m *base.Module, l0 int32) int32

//go:linkname Fn1420 github.com/goccy/pythonwasm2go/p2.Fn1420
func Fn1420(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn1422 github.com/goccy/pythonwasm2go/p2.Fn1422
func Fn1422(m *base.Module, l0 int32) int32

//go:linkname Fn1424 github.com/goccy/pythonwasm2go/p2.Fn1424
func Fn1424(m *base.Module, l0 int32) int32

//go:linkname Fn1425 github.com/goccy/pythonwasm2go/p2.Fn1425
func Fn1425(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1427 github.com/goccy/pythonwasm2go/p2.Fn1427
func Fn1427(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1431 github.com/goccy/pythonwasm2go/p2.Fn1431
func Fn1431(m *base.Module, l0 int32) int32

//go:linkname Fn1433 github.com/goccy/pythonwasm2go/p2.Fn1433
func Fn1433(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1434 github.com/goccy/pythonwasm2go/p2.Fn1434
func Fn1434(m *base.Module, l0 int32) int32

//go:linkname Fn1435 github.com/goccy/pythonwasm2go/p2.Fn1435
func Fn1435(m *base.Module, l0 int32) int32

//go:linkname Fn1440 github.com/goccy/pythonwasm2go/p2.Fn1440
func Fn1440(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1441 github.com/goccy/pythonwasm2go/p2.Fn1441
func Fn1441(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1442 github.com/goccy/pythonwasm2go/p2.Fn1442
func Fn1442(m *base.Module, l0 int32) int32

//go:linkname Fn1445 github.com/goccy/pythonwasm2go/p2.Fn1445
func Fn1445(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1446 github.com/goccy/pythonwasm2go/p2.Fn1446
func Fn1446(m *base.Module, l0 int32) int32

//go:linkname Fn1447 github.com/goccy/pythonwasm2go/p2.Fn1447
func Fn1447(m *base.Module, l0 int32) int32

//go:linkname Fn1450 github.com/goccy/pythonwasm2go/p2.Fn1450
func Fn1450(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1452 github.com/goccy/pythonwasm2go/p2.Fn1452
func Fn1452(m *base.Module, l0 int32) int32

//go:linkname Fn1460 github.com/goccy/pythonwasm2go/p2.Fn1460
func Fn1460(m *base.Module, l0 int32)

//go:linkname Fn1468 github.com/goccy/pythonwasm2go/p2.Fn1468
func Fn1468(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1475 github.com/goccy/pythonwasm2go/p2.Fn1475
func Fn1475(m *base.Module, l0 int32) int32

//go:linkname Fn1487 github.com/goccy/pythonwasm2go/p2.Fn1487
func Fn1487(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1490 github.com/goccy/pythonwasm2go/p2.Fn1490
func Fn1490(m *base.Module, l0 int32) int32

//go:linkname Fn1491 github.com/goccy/pythonwasm2go/p2.Fn1491
func Fn1491(m *base.Module, l0 int32) int32

//go:linkname Fn1493 github.com/goccy/pythonwasm2go/p2.Fn1493
func Fn1493(m *base.Module, l0 int32)

//go:linkname Fn1494 github.com/goccy/pythonwasm2go/p2.Fn1494
func Fn1494(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1507 github.com/goccy/pythonwasm2go/p2.Fn1507
func Fn1507(m *base.Module, l0 int32) int32

//go:linkname Fn1540 github.com/goccy/pythonwasm2go/p2.Fn1540
func Fn1540(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1541 github.com/goccy/pythonwasm2go/p2.Fn1541
func Fn1541(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1543 github.com/goccy/pythonwasm2go/p2.Fn1543
func Fn1543(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1544 github.com/goccy/pythonwasm2go/p2.Fn1544
func Fn1544(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1545 github.com/goccy/pythonwasm2go/p2.Fn1545
func Fn1545(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1546 github.com/goccy/pythonwasm2go/p2.Fn1546
func Fn1546(m *base.Module, l0 int32)

//go:linkname Fn1555 github.com/goccy/pythonwasm2go/p2.Fn1555
func Fn1555(m *base.Module, l0 int32)

//go:linkname Fn1559 github.com/goccy/pythonwasm2go/p2.Fn1559
func Fn1559(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1581 github.com/goccy/pythonwasm2go/p0.Fn1581
func Fn1581(m *base.Module, l0 int32) int32

//go:linkname Fn1584 github.com/goccy/pythonwasm2go/p2.Fn1584
func Fn1584(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1585 github.com/goccy/pythonwasm2go/p2.Fn1585
func Fn1585(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1586 github.com/goccy/pythonwasm2go/p2.Fn1586
func Fn1586(m *base.Module, l0 int32) int32

//go:linkname Fn1587 github.com/goccy/pythonwasm2go/p0.Fn1587
func Fn1587(m *base.Module, l0 int32) int32

//go:linkname Fn1630 github.com/goccy/pythonwasm2go/p2.Fn1630
func Fn1630(m *base.Module, l0 int32) int32

//go:linkname Fn1635 github.com/goccy/pythonwasm2go/p2.Fn1635
func Fn1635(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1638 github.com/goccy/pythonwasm2go/p2.Fn1638
func Fn1638(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1639 github.com/goccy/pythonwasm2go/p2.Fn1639
func Fn1639(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1643 github.com/goccy/pythonwasm2go/p2.Fn1643
func Fn1643(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1657 github.com/goccy/pythonwasm2go/p2.Fn1657
func Fn1657(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1666 github.com/goccy/pythonwasm2go/p2.Fn1666
func Fn1666(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1672 github.com/goccy/pythonwasm2go/p0.Fn1672
func Fn1672(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1673 github.com/goccy/pythonwasm2go/p0.Fn1673
func Fn1673(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1674 github.com/goccy/pythonwasm2go/p0.Fn1674
func Fn1674(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1675 github.com/goccy/pythonwasm2go/p0.Fn1675
func Fn1675(m *base.Module, l0 int32) int32

//go:linkname Fn1676 github.com/goccy/pythonwasm2go/p2.Fn1676
func Fn1676(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1679 github.com/goccy/pythonwasm2go/p0.Fn1679
func Fn1679(m *base.Module, l0 int32) int32

//go:linkname Fn1698 github.com/goccy/pythonwasm2go/p2.Fn1698
func Fn1698(m *base.Module, l0 int32) int32

//go:linkname Fn1701 github.com/goccy/pythonwasm2go/p2.Fn1701
func Fn1701(m *base.Module, l0 int32) int32

//go:linkname Fn1707 github.com/goccy/pythonwasm2go/p2.Fn1707
func Fn1707(m *base.Module, l0 int32)

//go:linkname Fn1708 github.com/goccy/pythonwasm2go/p2.Fn1708
func Fn1708(m *base.Module, l0 int32)

//go:linkname Fn1709 github.com/goccy/pythonwasm2go/p2.Fn1709
func Fn1709(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1710 github.com/goccy/pythonwasm2go/p2.Fn1710
func Fn1710(m *base.Module, l0 int32)

//go:linkname Fn1711 github.com/goccy/pythonwasm2go/p0.Fn1711
func Fn1711(m *base.Module, l0 int32) int32

//go:linkname Fn1712 github.com/goccy/pythonwasm2go/p2.Fn1712
func Fn1712(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1714 github.com/goccy/pythonwasm2go/p2.Fn1714
func Fn1714(m *base.Module, l0 int32) int32

//go:linkname Fn1716 github.com/goccy/pythonwasm2go/p0.Fn1716
func Fn1716(m *base.Module, l0 int32) int32

//go:linkname Fn1717 github.com/goccy/pythonwasm2go/p0.Fn1717
func Fn1717(m *base.Module, l0 int32) int32

//go:linkname Fn1718 github.com/goccy/pythonwasm2go/p0.Fn1718
func Fn1718(m *base.Module, l0 int32) int32

//go:linkname Fn1719 github.com/goccy/pythonwasm2go/p2.Fn1719
func Fn1719(m *base.Module, l0 int32) int32

//go:linkname Fn1722 github.com/goccy/pythonwasm2go/p2.Fn1722
func Fn1722(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1724 github.com/goccy/pythonwasm2go/p0.Fn1724
func Fn1724(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1725 github.com/goccy/pythonwasm2go/p0.Fn1725
func Fn1725(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1726 github.com/goccy/pythonwasm2go/p2.Fn1726
func Fn1726(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1727 github.com/goccy/pythonwasm2go/p0.Fn1727
func Fn1727(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1728 github.com/goccy/pythonwasm2go/p0.Fn1728
func Fn1728(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1729 github.com/goccy/pythonwasm2go/p2.Fn1729
func Fn1729(m *base.Module, l0 int32) int32

//go:linkname Fn1731 github.com/goccy/pythonwasm2go/p0.Fn1731
func Fn1731(m *base.Module, l0 int32) int32

//go:linkname Fn1732 github.com/goccy/pythonwasm2go/p0.Fn1732
func Fn1732(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1733 github.com/goccy/pythonwasm2go/p0.Fn1733
func Fn1733(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1734 github.com/goccy/pythonwasm2go/p0.Fn1734
func Fn1734(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1737 github.com/goccy/pythonwasm2go/p2.Fn1737
func Fn1737(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1738 github.com/goccy/pythonwasm2go/p0.Fn1738
func Fn1738(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1739 github.com/goccy/pythonwasm2go/p2.Fn1739
func Fn1739(m *base.Module, l0 int32) int32

//go:linkname Fn1740 github.com/goccy/pythonwasm2go/p2.Fn1740
func Fn1740(m *base.Module, l0 int32) int32

//go:linkname Fn1741 github.com/goccy/pythonwasm2go/p2.Fn1741
func Fn1741(m *base.Module, l0 int32)

//go:linkname Fn1742 github.com/goccy/pythonwasm2go/p0.Fn1742
func Fn1742(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1744 github.com/goccy/pythonwasm2go/p2.Fn1744
func Fn1744(m *base.Module, l0 int32) int32

//go:linkname Fn1750 github.com/goccy/pythonwasm2go/p2.Fn1750
func Fn1750(m *base.Module, l0 int32) int32

//go:linkname Fn1751 github.com/goccy/pythonwasm2go/p2.Fn1751
func Fn1751(m *base.Module, l0 int32) int32

//go:linkname Fn1753 github.com/goccy/pythonwasm2go/p2.Fn1753
func Fn1753(m *base.Module, l0 int32)

//go:linkname Fn1759 github.com/goccy/pythonwasm2go/p2.Fn1759
func Fn1759(m *base.Module, l0 int32)

//go:linkname Fn1762 github.com/goccy/pythonwasm2go/p2.Fn1762
func Fn1762(m *base.Module, l0 int32) int32

//go:linkname Fn1763 github.com/goccy/pythonwasm2go/p2.Fn1763
func Fn1763(m *base.Module, l0 int32)

//go:linkname Fn1764 github.com/goccy/pythonwasm2go/p2.Fn1764
func Fn1764(m *base.Module, l0 int32) int32

//go:linkname Fn1765 github.com/goccy/pythonwasm2go/p2.Fn1765
func Fn1765(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1766 github.com/goccy/pythonwasm2go/p2.Fn1766
func Fn1766(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1767 github.com/goccy/pythonwasm2go/p2.Fn1767
func Fn1767(m *base.Module, l0 int32) int32

//go:linkname Fn1768 github.com/goccy/pythonwasm2go/p2.Fn1768
func Fn1768(m *base.Module, l0 int32) int32

//go:linkname Fn1769 github.com/goccy/pythonwasm2go/p2.Fn1769
func Fn1769(m *base.Module, l0 int32) int32

//go:linkname Fn1776 github.com/goccy/pythonwasm2go/p2.Fn1776
func Fn1776(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1777 github.com/goccy/pythonwasm2go/p2.Fn1777
func Fn1777(m *base.Module, l0 int32) int32

//go:linkname Fn1780 github.com/goccy/pythonwasm2go/p2.Fn1780
func Fn1780(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1781 github.com/goccy/pythonwasm2go/p2.Fn1781
func Fn1781(m *base.Module, l0 int32) int32

//go:linkname Fn1785 github.com/goccy/pythonwasm2go/p2.Fn1785
func Fn1785(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1788 github.com/goccy/pythonwasm2go/p2.Fn1788
func Fn1788(m *base.Module, l0 int32)

//go:linkname Fn1790 github.com/goccy/pythonwasm2go/p2.Fn1790
func Fn1790(m *base.Module, l0 int32) int32

//go:linkname Fn1794 github.com/goccy/pythonwasm2go/p2.Fn1794
func Fn1794(m *base.Module, l0 int32) int32

//go:linkname Fn1795 github.com/goccy/pythonwasm2go/p2.Fn1795
func Fn1795(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1796 github.com/goccy/pythonwasm2go/p2.Fn1796
func Fn1796(m *base.Module) int32

//go:linkname Fn1797 github.com/goccy/pythonwasm2go/p2.Fn1797
func Fn1797(m *base.Module) int32

//go:linkname Fn1803 github.com/goccy/pythonwasm2go/p2.Fn1803
func Fn1803(m *base.Module, l0 int32) int32

//go:linkname Fn1804 github.com/goccy/pythonwasm2go/p2.Fn1804
func Fn1804(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn1805 github.com/goccy/pythonwasm2go/p2.Fn1805
func Fn1805(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1806 github.com/goccy/pythonwasm2go/p2.Fn1806
func Fn1806(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1807 github.com/goccy/pythonwasm2go/p2.Fn1807
func Fn1807(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1808 github.com/goccy/pythonwasm2go/p2.Fn1808
func Fn1808(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1809 github.com/goccy/pythonwasm2go/p2.Fn1809
func Fn1809(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1810 github.com/goccy/pythonwasm2go/p2.Fn1810
func Fn1810(m *base.Module) int32

//go:linkname Fn1811 github.com/goccy/pythonwasm2go/p2.Fn1811
func Fn1811(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1812 github.com/goccy/pythonwasm2go/p2.Fn1812
func Fn1812(m *base.Module) int64

//go:linkname Fn1813 github.com/goccy/pythonwasm2go/p2.Fn1813
func Fn1813(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1815 github.com/goccy/pythonwasm2go/p2.Fn1815
func Fn1815(m *base.Module, l0 int32, l1 int64)

//go:linkname Fn1816 github.com/goccy/pythonwasm2go/p2.Fn1816
func Fn1816(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1818 github.com/goccy/pythonwasm2go/p2.Fn1818
func Fn1818(m *base.Module, l0 int32)

//go:linkname Fn1819 github.com/goccy/pythonwasm2go/p2.Fn1819
func Fn1819(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1820 github.com/goccy/pythonwasm2go/p2.Fn1820
func Fn1820(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1823 github.com/goccy/pythonwasm2go/p2.Fn1823
func Fn1823(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1824 github.com/goccy/pythonwasm2go/p2.Fn1824
func Fn1824(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1825 github.com/goccy/pythonwasm2go/p2.Fn1825
func Fn1825(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1827 github.com/goccy/pythonwasm2go/p2.Fn1827
func Fn1827(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1828 github.com/goccy/pythonwasm2go/p2.Fn1828
func Fn1828(m *base.Module, l0 int32)

//go:linkname Fn1829 github.com/goccy/pythonwasm2go/p2.Fn1829
func Fn1829(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1834 github.com/goccy/pythonwasm2go/p2.Fn1834
func Fn1834(m *base.Module)

//go:linkname Fn1835 github.com/goccy/pythonwasm2go/p2.Fn1835
func Fn1835(m *base.Module, l0 int32) int32

//go:linkname Fn1836 github.com/goccy/pythonwasm2go/p2.Fn1836
func Fn1836(m *base.Module, l0 int32)

//go:linkname Fn1840 github.com/goccy/pythonwasm2go/p2.Fn1840
func Fn1840(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1843 github.com/goccy/pythonwasm2go/p2.Fn1843
func Fn1843(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1844 github.com/goccy/pythonwasm2go/p2.Fn1844
func Fn1844(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1846 github.com/goccy/pythonwasm2go/p2.Fn1846
func Fn1846(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1847 github.com/goccy/pythonwasm2go/p2.Fn1847
func Fn1847(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1848 github.com/goccy/pythonwasm2go/p2.Fn1848
func Fn1848(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1849 github.com/goccy/pythonwasm2go/p2.Fn1849
func Fn1849(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1851 github.com/goccy/pythonwasm2go/p2.Fn1851
func Fn1851(m *base.Module, l0 int32)

//go:linkname Fn1857 github.com/goccy/pythonwasm2go/p2.Fn1857
func Fn1857(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1861 github.com/goccy/pythonwasm2go/p2.Fn1861
func Fn1861(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1865 github.com/goccy/pythonwasm2go/p2.Fn1865
func Fn1865(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1866 github.com/goccy/pythonwasm2go/p2.Fn1866
func Fn1866(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1867 github.com/goccy/pythonwasm2go/p2.Fn1867
func Fn1867(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1870 github.com/goccy/pythonwasm2go/p2.Fn1870
func Fn1870(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1876 github.com/goccy/pythonwasm2go/p2.Fn1876
func Fn1876(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1877 github.com/goccy/pythonwasm2go/p2.Fn1877
func Fn1877(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1878 github.com/goccy/pythonwasm2go/p2.Fn1878
func Fn1878(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1879 github.com/goccy/pythonwasm2go/p2.Fn1879
func Fn1879(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1881 github.com/goccy/pythonwasm2go/p2.Fn1881
func Fn1881(m *base.Module, l0 int32) int32

//go:linkname Fn1883 github.com/goccy/pythonwasm2go/p2.Fn1883
func Fn1883(m *base.Module, l0 int32) int32

//go:linkname Fn1884 github.com/goccy/pythonwasm2go/p2.Fn1884
func Fn1884(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1885 github.com/goccy/pythonwasm2go/p2.Fn1885
func Fn1885(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1886 github.com/goccy/pythonwasm2go/p2.Fn1886
func Fn1886(m *base.Module, l0 int32) int32

//go:linkname Fn1887 github.com/goccy/pythonwasm2go/p2.Fn1887
func Fn1887(m *base.Module, l0 int32)

//go:linkname Fn1888 github.com/goccy/pythonwasm2go/p2.Fn1888
func Fn1888(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1891 github.com/goccy/pythonwasm2go/p2.Fn1891
func Fn1891(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1892 github.com/goccy/pythonwasm2go/p2.Fn1892
func Fn1892(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1911 github.com/goccy/pythonwasm2go/p2.Fn1911
func Fn1911(m *base.Module, l0 int32) int32

//go:linkname Fn1912 github.com/goccy/pythonwasm2go/p2.Fn1912
func Fn1912(m *base.Module, l0 int32) int32

//go:linkname Fn1915 github.com/goccy/pythonwasm2go/p0.Fn1915
func Fn1915(m *base.Module, l0 int32)

//go:linkname Fn1917 github.com/goccy/pythonwasm2go/p0.Fn1917
func Fn1917(m *base.Module, l0 int32)

//go:linkname Fn1918 github.com/goccy/pythonwasm2go/p2.Fn1918
func Fn1918(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1920 github.com/goccy/pythonwasm2go/p2.Fn1920
func Fn1920(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1921 github.com/goccy/pythonwasm2go/p0.Fn1921
func Fn1921(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1924 github.com/goccy/pythonwasm2go/p2.Fn1924
func Fn1924(m *base.Module, l0 int32) int32

//go:linkname Fn1925 github.com/goccy/pythonwasm2go/p2.Fn1925
func Fn1925(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1926 github.com/goccy/pythonwasm2go/p2.Fn1926
func Fn1926(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1927 github.com/goccy/pythonwasm2go/p2.Fn1927
func Fn1927(m *base.Module, l0 int32)

//go:linkname Fn1928 github.com/goccy/pythonwasm2go/p2.Fn1928
func Fn1928(m *base.Module, l0 int32) int32

//go:linkname Fn1929 github.com/goccy/pythonwasm2go/p2.Fn1929
func Fn1929(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1930 github.com/goccy/pythonwasm2go/p2.Fn1930
func Fn1930(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1931 github.com/goccy/pythonwasm2go/p2.Fn1931
func Fn1931(m *base.Module, l0 int32)

//go:linkname Fn1932 github.com/goccy/pythonwasm2go/p2.Fn1932
func Fn1932(m *base.Module, l0 int32) int32

//go:linkname Fn1933 github.com/goccy/pythonwasm2go/p2.Fn1933
func Fn1933(m *base.Module, l0 int32) int32

//go:linkname Fn1934 github.com/goccy/pythonwasm2go/p2.Fn1934
func Fn1934(m *base.Module, l0 int32) int32

//go:linkname Fn1936 github.com/goccy/pythonwasm2go/p2.Fn1936
func Fn1936(m *base.Module, l0 int32) int32

//go:linkname Fn1937 github.com/goccy/pythonwasm2go/p2.Fn1937
func Fn1937(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1938 github.com/goccy/pythonwasm2go/p2.Fn1938
func Fn1938(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1939 github.com/goccy/pythonwasm2go/p2.Fn1939
func Fn1939(m *base.Module, l0 int32)

//go:linkname Fn1941 github.com/goccy/pythonwasm2go/p2.Fn1941
func Fn1941(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1942 github.com/goccy/pythonwasm2go/p2.Fn1942
func Fn1942(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1945 github.com/goccy/pythonwasm2go/p2.Fn1945
func Fn1945(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1946 github.com/goccy/pythonwasm2go/p2.Fn1946
func Fn1946(m *base.Module, l0 int32) int32

//go:linkname Fn1955 github.com/goccy/pythonwasm2go/p2.Fn1955
func Fn1955(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn1957 github.com/goccy/pythonwasm2go/p2.Fn1957
func Fn1957(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1958 github.com/goccy/pythonwasm2go/p2.Fn1958
func Fn1958(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1959 github.com/goccy/pythonwasm2go/p2.Fn1959
func Fn1959(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1960 github.com/goccy/pythonwasm2go/p2.Fn1960
func Fn1960(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1961 github.com/goccy/pythonwasm2go/p2.Fn1961
func Fn1961(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1962 github.com/goccy/pythonwasm2go/p2.Fn1962
func Fn1962(m *base.Module, l0 int32)

//go:linkname Fn1964 github.com/goccy/pythonwasm2go/p2.Fn1964
func Fn1964(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1978 github.com/goccy/pythonwasm2go/p2.Fn1978
func Fn1978(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1979 github.com/goccy/pythonwasm2go/p2.Fn1979
func Fn1979(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1990 github.com/goccy/pythonwasm2go/p2.Fn1990
func Fn1990(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1993 github.com/goccy/pythonwasm2go/p2.Fn1993
func Fn1993(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2010 github.com/goccy/pythonwasm2go/p0.Fn2010
func Fn2010(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2014 github.com/goccy/pythonwasm2go/p0.Fn2014
func Fn2014(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2016 github.com/goccy/pythonwasm2go/p0.Fn2016
func Fn2016(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2019 github.com/goccy/pythonwasm2go/p2.Fn2019
func Fn2019(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2021 github.com/goccy/pythonwasm2go/p2.Fn2021
func Fn2021(m *base.Module, l0 int32) int32

//go:linkname Fn2036 github.com/goccy/pythonwasm2go/p0.Fn2036
func Fn2036(m *base.Module, l0 int32) int32

//go:linkname Fn2037 github.com/goccy/pythonwasm2go/p2.Fn2037
func Fn2037(m *base.Module, l0 int32) int32

//go:linkname Fn2038 github.com/goccy/pythonwasm2go/p2.Fn2038
func Fn2038(m *base.Module, l0 int32) int32

//go:linkname Fn2040 github.com/goccy/pythonwasm2go/p2.Fn2040
func Fn2040(m *base.Module, l0 int32) int32

//go:linkname Fn2042 github.com/goccy/pythonwasm2go/p0.Fn2042
func Fn2042(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2043 github.com/goccy/pythonwasm2go/p2.Fn2043
func Fn2043(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2044 github.com/goccy/pythonwasm2go/p2.Fn2044
func Fn2044(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2045 github.com/goccy/pythonwasm2go/p2.Fn2045
func Fn2045(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2046 github.com/goccy/pythonwasm2go/p0.Fn2046
func Fn2046(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2048 github.com/goccy/pythonwasm2go/p0.Fn2048
func Fn2048(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2049 github.com/goccy/pythonwasm2go/p0.Fn2049
func Fn2049(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2050 github.com/goccy/pythonwasm2go/p0.Fn2050
func Fn2050(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2053 github.com/goccy/pythonwasm2go/p2.Fn2053
func Fn2053(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2068 github.com/goccy/pythonwasm2go/p2.Fn2068
func Fn2068(m *base.Module, l0 int32) int32

//go:linkname Fn2076 github.com/goccy/pythonwasm2go/p2.Fn2076
func Fn2076(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2077 github.com/goccy/pythonwasm2go/p2.Fn2077
func Fn2077(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2098 github.com/goccy/pythonwasm2go/p2.Fn2098
func Fn2098(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2101 github.com/goccy/pythonwasm2go/p2.Fn2101
func Fn2101(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2102 github.com/goccy/pythonwasm2go/p2.Fn2102
func Fn2102(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2103 github.com/goccy/pythonwasm2go/p2.Fn2103
func Fn2103(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2105 github.com/goccy/pythonwasm2go/p2.Fn2105
func Fn2105(m *base.Module, l0 int32) int32

//go:linkname Fn2114 github.com/goccy/pythonwasm2go/p0.Fn2114
func Fn2114(m *base.Module, l0 int32) int32

//go:linkname Fn2115 github.com/goccy/pythonwasm2go/p0.Fn2115
func Fn2115(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2116 github.com/goccy/pythonwasm2go/p2.Fn2116
func Fn2116(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2117 github.com/goccy/pythonwasm2go/p2.Fn2117
func Fn2117(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2119 github.com/goccy/pythonwasm2go/p2.Fn2119
func Fn2119(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2120 github.com/goccy/pythonwasm2go/p2.Fn2120
func Fn2120(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2132 github.com/goccy/pythonwasm2go/p2.Fn2132
func Fn2132(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2138 github.com/goccy/pythonwasm2go/p0.Fn2138
func Fn2138(m *base.Module, l0 int32) int32

//go:linkname Fn2139 github.com/goccy/pythonwasm2go/p0.Fn2139
func Fn2139(m *base.Module, l0 int32) int32

//go:linkname Fn2140 github.com/goccy/pythonwasm2go/p0.Fn2140
func Fn2140(m *base.Module, l0 int32) int32

//go:linkname Fn2141 github.com/goccy/pythonwasm2go/p2.Fn2141
func Fn2141(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2142 github.com/goccy/pythonwasm2go/p2.Fn2142
func Fn2142(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2143 github.com/goccy/pythonwasm2go/p0.Fn2143
func Fn2143(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2144 github.com/goccy/pythonwasm2go/p0.Fn2144
func Fn2144(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2147 github.com/goccy/pythonwasm2go/p2.Fn2147
func Fn2147(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2163 github.com/goccy/pythonwasm2go/p2.Fn2163
func Fn2163(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2169 github.com/goccy/pythonwasm2go/p2.Fn2169
func Fn2169(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2171 github.com/goccy/pythonwasm2go/p2.Fn2171
func Fn2171(m *base.Module, l0 int32) int32

//go:linkname Fn2172 github.com/goccy/pythonwasm2go/p2.Fn2172
func Fn2172(m *base.Module, l0 int32) int32

//go:linkname Fn2173 github.com/goccy/pythonwasm2go/p2.Fn2173
func Fn2173(m *base.Module, l0 int32) int32

//go:linkname Fn2174 github.com/goccy/pythonwasm2go/p2.Fn2174
func Fn2174(m *base.Module, l0 int32) int32

//go:linkname Fn2175 github.com/goccy/pythonwasm2go/p2.Fn2175
func Fn2175(m *base.Module, l0 int32) int32

//go:linkname Fn2176 github.com/goccy/pythonwasm2go/p2.Fn2176
func Fn2176(m *base.Module, l0 int32) int32

//go:linkname Fn2186 github.com/goccy/pythonwasm2go/p2.Fn2186
func Fn2186(m *base.Module) int32

//go:linkname Fn2187 github.com/goccy/pythonwasm2go/p2.Fn2187
func Fn2187(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2188 github.com/goccy/pythonwasm2go/p2.Fn2188
func Fn2188(m *base.Module, l0 int32)

//go:linkname Fn2189 github.com/goccy/pythonwasm2go/p0.Fn2189
func Fn2189(m *base.Module, l0 int32)

//go:linkname Fn2190 github.com/goccy/pythonwasm2go/p2.Fn2190
func Fn2190(m *base.Module, l0 int32) int32

//go:linkname Fn2193 github.com/goccy/pythonwasm2go/p0.Fn2193
func Fn2193(m *base.Module, l0 int32) int32

//go:linkname Fn2195 github.com/goccy/pythonwasm2go/p0.Fn2195
func Fn2195(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2197 github.com/goccy/pythonwasm2go/p2.Fn2197
func Fn2197(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2198 github.com/goccy/pythonwasm2go/p2.Fn2198
func Fn2198(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2199 github.com/goccy/pythonwasm2go/p2.Fn2199
func Fn2199(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2200 github.com/goccy/pythonwasm2go/p2.Fn2200
func Fn2200(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2201 github.com/goccy/pythonwasm2go/p0.Fn2201
func Fn2201(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2202 github.com/goccy/pythonwasm2go/p0.Fn2202
func Fn2202(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2203 github.com/goccy/pythonwasm2go/p0.Fn2203
func Fn2203(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2204 github.com/goccy/pythonwasm2go/p0.Fn2204
func Fn2204(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2205 github.com/goccy/pythonwasm2go/p0.Fn2205
func Fn2205(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2214 github.com/goccy/pythonwasm2go/p2.Fn2214
func Fn2214(m *base.Module, l0 int32) int32

//go:linkname Fn2215 github.com/goccy/pythonwasm2go/p2.Fn2215
func Fn2215(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2216 github.com/goccy/pythonwasm2go/p0.Fn2216
func Fn2216(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2218 github.com/goccy/pythonwasm2go/p2.Fn2218
func Fn2218(m *base.Module, l0 int32) int32

//go:linkname Fn2220 github.com/goccy/pythonwasm2go/p2.Fn2220
func Fn2220(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2222 github.com/goccy/pythonwasm2go/p0.Fn2222
func Fn2222(m *base.Module, l0 int32) int32

//go:linkname Fn2223 github.com/goccy/pythonwasm2go/p2.Fn2223
func Fn2223(m *base.Module, l0 int32) int32

//go:linkname Fn2225 github.com/goccy/pythonwasm2go/p2.Fn2225
func Fn2225(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2226 github.com/goccy/pythonwasm2go/p0.Fn2226
func Fn2226(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2227 github.com/goccy/pythonwasm2go/p2.Fn2227
func Fn2227(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2228 github.com/goccy/pythonwasm2go/p2.Fn2228
func Fn2228(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2229 github.com/goccy/pythonwasm2go/p2.Fn2229
func Fn2229(m *base.Module, l0 int32) int32

//go:linkname Fn2231 github.com/goccy/pythonwasm2go/p0.Fn2231
func Fn2231(m *base.Module, l0 int32) int32

//go:linkname Fn2233 github.com/goccy/pythonwasm2go/p2.Fn2233
func Fn2233(m *base.Module, l0 int32) int32

//go:linkname Fn2234 github.com/goccy/pythonwasm2go/p2.Fn2234
func Fn2234(m *base.Module, l0 int32) int32

//go:linkname Fn2235 github.com/goccy/pythonwasm2go/p2.Fn2235
func Fn2235(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2236 github.com/goccy/pythonwasm2go/p2.Fn2236
func Fn2236(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2238 github.com/goccy/pythonwasm2go/p2.Fn2238
func Fn2238(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2239 github.com/goccy/pythonwasm2go/p2.Fn2239
func Fn2239(m *base.Module, l0 int32) int32

//go:linkname Fn2240 github.com/goccy/pythonwasm2go/p2.Fn2240
func Fn2240(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2241 github.com/goccy/pythonwasm2go/p2.Fn2241
func Fn2241(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2248 github.com/goccy/pythonwasm2go/p0.Fn2248
func Fn2248(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2253 github.com/goccy/pythonwasm2go/p2.Fn2253
func Fn2253(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2255 github.com/goccy/pythonwasm2go/p2.Fn2255
func Fn2255(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2273 github.com/goccy/pythonwasm2go/p0.Fn2273
func Fn2273(m *base.Module, l0 int32) int32

//go:linkname Fn2278 github.com/goccy/pythonwasm2go/p2.Fn2278
func Fn2278(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2280 github.com/goccy/pythonwasm2go/p2.Fn2280
func Fn2280(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2281 github.com/goccy/pythonwasm2go/p0.Fn2281
func Fn2281(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2282 github.com/goccy/pythonwasm2go/p0.Fn2282
func Fn2282(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2283 github.com/goccy/pythonwasm2go/p2.Fn2283
func Fn2283(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2284 github.com/goccy/pythonwasm2go/p0.Fn2284
func Fn2284(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2286 github.com/goccy/pythonwasm2go/p2.Fn2286
func Fn2286(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2288 github.com/goccy/pythonwasm2go/p2.Fn2288
func Fn2288(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2306 github.com/goccy/pythonwasm2go/p0.Fn2306
func Fn2306(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2323 github.com/goccy/pythonwasm2go/p2.Fn2323
func Fn2323(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2329 github.com/goccy/pythonwasm2go/p2.Fn2329
func Fn2329(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2343 github.com/goccy/pythonwasm2go/p2.Fn2343
func Fn2343(m *base.Module, l0 int32) int32

//go:linkname Fn2351 github.com/goccy/pythonwasm2go/p2.Fn2351
func Fn2351(m *base.Module) int32

//go:linkname Fn2353 github.com/goccy/pythonwasm2go/p0.Fn2353
func Fn2353(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2405 github.com/goccy/pythonwasm2go/p2.Fn2405
func Fn2405(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2463 github.com/goccy/pythonwasm2go/p2.Fn2463
func Fn2463(m *base.Module) int32

//go:linkname Fn2466 github.com/goccy/pythonwasm2go/p2.Fn2466
func Fn2466(m *base.Module, l0 int32) int32

//go:linkname Fn2473 github.com/goccy/pythonwasm2go/p2.Fn2473
func Fn2473(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn2478 github.com/goccy/pythonwasm2go/p2.Fn2478
func Fn2478(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2479 github.com/goccy/pythonwasm2go/p2.Fn2479
func Fn2479(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2480 github.com/goccy/pythonwasm2go/p2.Fn2480
func Fn2480(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2481 github.com/goccy/pythonwasm2go/p2.Fn2481
func Fn2481(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2493 github.com/goccy/pythonwasm2go/p2.Fn2493
func Fn2493(m *base.Module, l0 int32) int32

//go:linkname Fn2544 github.com/goccy/pythonwasm2go/p2.Fn2544
func Fn2544(m *base.Module, l0 int32) int32

//go:linkname Fn2545 github.com/goccy/pythonwasm2go/p2.Fn2545
func Fn2545(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2546 github.com/goccy/pythonwasm2go/p0.Fn2546
func Fn2546(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2547 github.com/goccy/pythonwasm2go/p2.Fn2547
func Fn2547(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn2549 github.com/goccy/pythonwasm2go/p2.Fn2549
func Fn2549(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2550 github.com/goccy/pythonwasm2go/p2.Fn2550
func Fn2550(m *base.Module, l0 int32) int32

//go:linkname Fn2551 github.com/goccy/pythonwasm2go/p2.Fn2551
func Fn2551(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2553 github.com/goccy/pythonwasm2go/p0.Fn2553
func Fn2553(m *base.Module, l0 int32) int32

//go:linkname Fn2554 github.com/goccy/pythonwasm2go/p2.Fn2554
func Fn2554(m *base.Module, l0 int32) int32

//go:linkname Fn2555 github.com/goccy/pythonwasm2go/p2.Fn2555
func Fn2555(m *base.Module, l0 int32) int32

//go:linkname Fn2556 github.com/goccy/pythonwasm2go/p2.Fn2556
func Fn2556(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2558 github.com/goccy/pythonwasm2go/p0.Fn2558
func Fn2558(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2559 github.com/goccy/pythonwasm2go/p0.Fn2559
func Fn2559(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2560 github.com/goccy/pythonwasm2go/p2.Fn2560
func Fn2560(m *base.Module, l0 int32) int32

//go:linkname Fn2561 github.com/goccy/pythonwasm2go/p0.Fn2561
func Fn2561(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2564 github.com/goccy/pythonwasm2go/p2.Fn2564
func Fn2564(m *base.Module, l0 int32) int32

//go:linkname Fn2566 github.com/goccy/pythonwasm2go/p0.Fn2566
func Fn2566(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2568 github.com/goccy/pythonwasm2go/p0.Fn2568
func Fn2568(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2569 github.com/goccy/pythonwasm2go/p0.Fn2569
func Fn2569(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2571 github.com/goccy/pythonwasm2go/p0.Fn2571
func Fn2571(m *base.Module, l0 int32) int32

//go:linkname Fn2572 github.com/goccy/pythonwasm2go/p0.Fn2572
func Fn2572(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2574 github.com/goccy/pythonwasm2go/p0.Fn2574
func Fn2574(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2575 github.com/goccy/pythonwasm2go/p2.Fn2575
func Fn2575(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2576 github.com/goccy/pythonwasm2go/p0.Fn2576
func Fn2576(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2577 github.com/goccy/pythonwasm2go/p0.Fn2577
func Fn2577(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2578 github.com/goccy/pythonwasm2go/p0.Fn2578
func Fn2578(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2579 github.com/goccy/pythonwasm2go/p0.Fn2579
func Fn2579(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2582 github.com/goccy/pythonwasm2go/p2.Fn2582
func Fn2582(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2583 github.com/goccy/pythonwasm2go/p0.Fn2583
func Fn2583(m *base.Module, l0 int32) int32

//go:linkname Fn2584 github.com/goccy/pythonwasm2go/p2.Fn2584
func Fn2584(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2586 github.com/goccy/pythonwasm2go/p2.Fn2586
func Fn2586(m *base.Module, l0 int32) int32

//go:linkname Fn2587 github.com/goccy/pythonwasm2go/p0.Fn2587
func Fn2587(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2589 github.com/goccy/pythonwasm2go/p0.Fn2589
func Fn2589(m *base.Module, l0 int32) int32

//go:linkname Fn2590 github.com/goccy/pythonwasm2go/p2.Fn2590
func Fn2590(m *base.Module, l0 int32)

//go:linkname Fn2591 github.com/goccy/pythonwasm2go/p0.Fn2591
func Fn2591(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2592 github.com/goccy/pythonwasm2go/p0.Fn2592
func Fn2592(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2593 github.com/goccy/pythonwasm2go/p0.Fn2593
func Fn2593(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2597 github.com/goccy/pythonwasm2go/p0.Fn2597
func Fn2597(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2601 github.com/goccy/pythonwasm2go/p0.Fn2601
func Fn2601(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2602 github.com/goccy/pythonwasm2go/p0.Fn2602
func Fn2602(m *base.Module, l0 int32) int32

//go:linkname Fn2603 github.com/goccy/pythonwasm2go/p2.Fn2603
func Fn2603(m *base.Module, l0 int32) int32

//go:linkname Fn2604 github.com/goccy/pythonwasm2go/p2.Fn2604
func Fn2604(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2606 github.com/goccy/pythonwasm2go/p0.Fn2606
func Fn2606(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2608 github.com/goccy/pythonwasm2go/p0.Fn2608
func Fn2608(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2609 github.com/goccy/pythonwasm2go/p0.Fn2609
func Fn2609(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2610 github.com/goccy/pythonwasm2go/p0.Fn2610
func Fn2610(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2611 github.com/goccy/pythonwasm2go/p0.Fn2611
func Fn2611(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2612 github.com/goccy/pythonwasm2go/p0.Fn2612
func Fn2612(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2616 github.com/goccy/pythonwasm2go/p0.Fn2616
func Fn2616(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32

//go:linkname Fn2619 github.com/goccy/pythonwasm2go/p0.Fn2619
func Fn2619(m *base.Module, l0 int32) int32

//go:linkname Fn2621 github.com/goccy/pythonwasm2go/p0.Fn2621
func Fn2621(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2632 github.com/goccy/pythonwasm2go/p0.Fn2632
func Fn2632(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32

//go:linkname Fn2633 github.com/goccy/pythonwasm2go/p0.Fn2633
func Fn2633(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn2635 github.com/goccy/pythonwasm2go/p0.Fn2635
func Fn2635(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2637 github.com/goccy/pythonwasm2go/p0.Fn2637
func Fn2637(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2638 github.com/goccy/pythonwasm2go/p0.Fn2638
func Fn2638(m *base.Module, l0 int32) int32

//go:linkname Fn2639 github.com/goccy/pythonwasm2go/p0.Fn2639
func Fn2639(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2642 github.com/goccy/pythonwasm2go/p2.Fn2642
func Fn2642(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2644 github.com/goccy/pythonwasm2go/p2.Fn2644
func Fn2644(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2645 github.com/goccy/pythonwasm2go/p2.Fn2645
func Fn2645(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2646 github.com/goccy/pythonwasm2go/p2.Fn2646
func Fn2646(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2647 github.com/goccy/pythonwasm2go/p2.Fn2647
func Fn2647(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2648 github.com/goccy/pythonwasm2go/p2.Fn2648
func Fn2648(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2649 github.com/goccy/pythonwasm2go/p2.Fn2649
func Fn2649(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2650 github.com/goccy/pythonwasm2go/p2.Fn2650
func Fn2650(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2651 github.com/goccy/pythonwasm2go/p0.Fn2651
func Fn2651(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2652 github.com/goccy/pythonwasm2go/p0.Fn2652
func Fn2652(m *base.Module, l0 int32) int32

//go:linkname Fn2653 github.com/goccy/pythonwasm2go/p2.Fn2653
func Fn2653(m *base.Module, l0 int32) int32

//go:linkname Fn2654 github.com/goccy/pythonwasm2go/p0.Fn2654
func Fn2654(m *base.Module, l0 int32) int32

//go:linkname Fn2655 github.com/goccy/pythonwasm2go/p2.Fn2655
func Fn2655(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2656 github.com/goccy/pythonwasm2go/p2.Fn2656
func Fn2656(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2658 github.com/goccy/pythonwasm2go/p0.Fn2658
func Fn2658(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2659 github.com/goccy/pythonwasm2go/p0.Fn2659
func Fn2659(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2665 github.com/goccy/pythonwasm2go/p2.Fn2665
func Fn2665(m *base.Module) int32

//go:linkname Fn2667 github.com/goccy/pythonwasm2go/p2.Fn2667
func Fn2667(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2674 github.com/goccy/pythonwasm2go/p2.Fn2674
func Fn2674(m *base.Module, l0 int32) int32

//go:linkname Fn2675 github.com/goccy/pythonwasm2go/p2.Fn2675
func Fn2675(m *base.Module, l0 int32) int32

//go:linkname Fn2680 github.com/goccy/pythonwasm2go/p2.Fn2680
func Fn2680(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2682 github.com/goccy/pythonwasm2go/p2.Fn2682
func Fn2682(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2683 github.com/goccy/pythonwasm2go/p0.Fn2683
func Fn2683(m *base.Module, l0 int32) int32

//go:linkname Fn2686 github.com/goccy/pythonwasm2go/p2.Fn2686
func Fn2686(m *base.Module, l0 int32) int32

//go:linkname Fn2687 github.com/goccy/pythonwasm2go/p2.Fn2687
func Fn2687(m *base.Module, l0 int32) int32

//go:linkname Fn2690 github.com/goccy/pythonwasm2go/p2.Fn2690
func Fn2690(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn2692 github.com/goccy/pythonwasm2go/p0.Fn2692
func Fn2692(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2693 github.com/goccy/pythonwasm2go/p2.Fn2693
func Fn2693(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2694 github.com/goccy/pythonwasm2go/p2.Fn2694
func Fn2694(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2695 github.com/goccy/pythonwasm2go/p2.Fn2695
func Fn2695(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2698 github.com/goccy/pythonwasm2go/p2.Fn2698
func Fn2698(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2699 github.com/goccy/pythonwasm2go/p2.Fn2699
func Fn2699(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2700 github.com/goccy/pythonwasm2go/p2.Fn2700
func Fn2700(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2705 github.com/goccy/pythonwasm2go/p2.Fn2705
func Fn2705(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2706 github.com/goccy/pythonwasm2go/p2.Fn2706
func Fn2706(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2709 github.com/goccy/pythonwasm2go/p2.Fn2709
func Fn2709(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2711 github.com/goccy/pythonwasm2go/p2.Fn2711
func Fn2711(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2713 github.com/goccy/pythonwasm2go/p2.Fn2713
func Fn2713(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2714 github.com/goccy/pythonwasm2go/p0.Fn2714
func Fn2714(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2716 github.com/goccy/pythonwasm2go/p2.Fn2716
func Fn2716(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2719 github.com/goccy/pythonwasm2go/p2.Fn2719
func Fn2719(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2721 github.com/goccy/pythonwasm2go/p0.Fn2721
func Fn2721(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2722 github.com/goccy/pythonwasm2go/p2.Fn2722
func Fn2722(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2724 github.com/goccy/pythonwasm2go/p2.Fn2724
func Fn2724(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2727 github.com/goccy/pythonwasm2go/p2.Fn2727
func Fn2727(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2728 github.com/goccy/pythonwasm2go/p0.Fn2728
func Fn2728(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2729 github.com/goccy/pythonwasm2go/p0.Fn2729
func Fn2729(m *base.Module, l0 int32) int32

//go:linkname Fn2730 github.com/goccy/pythonwasm2go/p2.Fn2730
func Fn2730(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2731 github.com/goccy/pythonwasm2go/p0.Fn2731
func Fn2731(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2734 github.com/goccy/pythonwasm2go/p0.Fn2734
func Fn2734(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2737 github.com/goccy/pythonwasm2go/p0.Fn2737
func Fn2737(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2738 github.com/goccy/pythonwasm2go/p2.Fn2738
func Fn2738(m *base.Module, l0 int32) int32

//go:linkname Fn2739 github.com/goccy/pythonwasm2go/p2.Fn2739
func Fn2739(m *base.Module, l0 int32)

//go:linkname Fn2741 github.com/goccy/pythonwasm2go/p2.Fn2741
func Fn2741(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2742 github.com/goccy/pythonwasm2go/p2.Fn2742
func Fn2742(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2743 github.com/goccy/pythonwasm2go/p2.Fn2743
func Fn2743(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2744 github.com/goccy/pythonwasm2go/p2.Fn2744
func Fn2744(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2745 github.com/goccy/pythonwasm2go/p2.Fn2745
func Fn2745(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2746 github.com/goccy/pythonwasm2go/p2.Fn2746
func Fn2746(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2747 github.com/goccy/pythonwasm2go/p2.Fn2747
func Fn2747(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2748 github.com/goccy/pythonwasm2go/p2.Fn2748
func Fn2748(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2749 github.com/goccy/pythonwasm2go/p0.Fn2749
func Fn2749(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2750 github.com/goccy/pythonwasm2go/p2.Fn2750
func Fn2750(m *base.Module, l0 int32) int32

//go:linkname Fn2763 github.com/goccy/pythonwasm2go/p2.Fn2763
func Fn2763(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2769 github.com/goccy/pythonwasm2go/p0.Fn2769
func Fn2769(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2770 github.com/goccy/pythonwasm2go/p0.Fn2770
func Fn2770(m *base.Module, l0 int32) int32

//go:linkname Fn2779 github.com/goccy/pythonwasm2go/p2.Fn2779
func Fn2779(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2780 github.com/goccy/pythonwasm2go/p2.Fn2780
func Fn2780(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2781 github.com/goccy/pythonwasm2go/p2.Fn2781
func Fn2781(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2783 github.com/goccy/pythonwasm2go/p2.Fn2783
func Fn2783(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2784 github.com/goccy/pythonwasm2go/p2.Fn2784
func Fn2784(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2787 github.com/goccy/pythonwasm2go/p2.Fn2787
func Fn2787(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2788 github.com/goccy/pythonwasm2go/p2.Fn2788
func Fn2788(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2789 github.com/goccy/pythonwasm2go/p2.Fn2789
func Fn2789(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2792 github.com/goccy/pythonwasm2go/p2.Fn2792
func Fn2792(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2795 github.com/goccy/pythonwasm2go/p2.Fn2795
func Fn2795(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2848 github.com/goccy/pythonwasm2go/p2.Fn2848
func Fn2848(m *base.Module, l0 int32) int32

//go:linkname Fn2863 github.com/goccy/pythonwasm2go/p2.Fn2863
func Fn2863(m *base.Module, l0 int32) int32

//go:linkname Fn2865 github.com/goccy/pythonwasm2go/p2.Fn2865
func Fn2865(m *base.Module, l0 int32) int32

//go:linkname Fn2874 github.com/goccy/pythonwasm2go/p2.Fn2874
func Fn2874(m *base.Module, l0 int32) int32

//go:linkname Fn2877 github.com/goccy/pythonwasm2go/p2.Fn2877
func Fn2877(m *base.Module, l0 int32) int32

//go:linkname Fn2878 github.com/goccy/pythonwasm2go/p2.Fn2878
func Fn2878(m *base.Module, l0 int32) int32

//go:linkname Fn2885 github.com/goccy/pythonwasm2go/p2.Fn2885
func Fn2885(m *base.Module, l0 int32) int32

//go:linkname Fn2894 github.com/goccy/pythonwasm2go/p2.Fn2894
func Fn2894(m *base.Module, l0 int32) int32

//go:linkname Fn2897 github.com/goccy/pythonwasm2go/p2.Fn2897
func Fn2897(m *base.Module, l0 int32)

//go:linkname Fn2900 github.com/goccy/pythonwasm2go/p2.Fn2900
func Fn2900(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2915 github.com/goccy/pythonwasm2go/p2.Fn2915
func Fn2915(m *base.Module, l0 int32) int32

//go:linkname Fn2916 github.com/goccy/pythonwasm2go/p2.Fn2916
func Fn2916(m *base.Module, l0 int32) int32

//go:linkname Fn2918 github.com/goccy/pythonwasm2go/p2.Fn2918
func Fn2918(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2933 github.com/goccy/pythonwasm2go/p2.Fn2933
func Fn2933(m *base.Module, l0 int32) int32

//go:linkname Fn2940 github.com/goccy/pythonwasm2go/p0.Fn2940
func Fn2940(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2943 github.com/goccy/pythonwasm2go/p2.Fn2943
func Fn2943(m *base.Module, l0 int32) int32

//go:linkname Fn2944 github.com/goccy/pythonwasm2go/p2.Fn2944
func Fn2944(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2945 github.com/goccy/pythonwasm2go/p2.Fn2945
func Fn2945(m *base.Module, l0 int32) int32

//go:linkname Fn2947 github.com/goccy/pythonwasm2go/p2.Fn2947
func Fn2947(m *base.Module, l0 int32) int32

//go:linkname Fn2948 github.com/goccy/pythonwasm2go/p2.Fn2948
func Fn2948(m *base.Module, l0 int32)

//go:linkname Fn2950 github.com/goccy/pythonwasm2go/p2.Fn2950
func Fn2950(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2994 github.com/goccy/pythonwasm2go/p0.Fn2994
func Fn2994(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2998 github.com/goccy/pythonwasm2go/p0.Fn2998
func Fn2998(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2999 github.com/goccy/pythonwasm2go/p0.Fn2999
func Fn2999(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3000 github.com/goccy/pythonwasm2go/p0.Fn3000
func Fn3000(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3001 github.com/goccy/pythonwasm2go/p0.Fn3001
func Fn3001(m *base.Module) int32

//go:linkname Fn3002 github.com/goccy/pythonwasm2go/p0.Fn3002
func Fn3002(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32

//go:linkname Fn3006 github.com/goccy/pythonwasm2go/p0.Fn3006
func Fn3006(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3008 github.com/goccy/pythonwasm2go/p2.Fn3008
func Fn3008(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3020 github.com/goccy/pythonwasm2go/p2.Fn3020
func Fn3020(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3055 github.com/goccy/pythonwasm2go/p2.Fn3055
func Fn3055(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn3056 github.com/goccy/pythonwasm2go/p2.Fn3056
func Fn3056(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3072 github.com/goccy/pythonwasm2go/p2.Fn3072
func Fn3072(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3074 github.com/goccy/pythonwasm2go/p2.Fn3074
func Fn3074(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3075 github.com/goccy/pythonwasm2go/p2.Fn3075
func Fn3075(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn3076 github.com/goccy/pythonwasm2go/p2.Fn3076
func Fn3076(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn3077 github.com/goccy/pythonwasm2go/p2.Fn3077
func Fn3077(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3078 github.com/goccy/pythonwasm2go/p2.Fn3078
func Fn3078(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3079 github.com/goccy/pythonwasm2go/p2.Fn3079
func Fn3079(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3080 github.com/goccy/pythonwasm2go/p2.Fn3080
func Fn3080(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3084 github.com/goccy/pythonwasm2go/p2.Fn3084
func Fn3084(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn3085 github.com/goccy/pythonwasm2go/p2.Fn3085
func Fn3085(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn3087 github.com/goccy/pythonwasm2go/p2.Fn3087
func Fn3087(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3088 github.com/goccy/pythonwasm2go/p2.Fn3088
func Fn3088(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3092 github.com/goccy/pythonwasm2go/p2.Fn3092
func Fn3092(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3097 github.com/goccy/pythonwasm2go/p2.Fn3097
func Fn3097(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn3098 github.com/goccy/pythonwasm2go/p2.Fn3098
func Fn3098(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3099 github.com/goccy/pythonwasm2go/p2.Fn3099
func Fn3099(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3102 github.com/goccy/pythonwasm2go/p2.Fn3102
func Fn3102(m *base.Module) int32

//go:linkname Fn3104 github.com/goccy/pythonwasm2go/p2.Fn3104
func Fn3104(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3106 github.com/goccy/pythonwasm2go/p0.Fn3106
func Fn3106(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3108 github.com/goccy/pythonwasm2go/p0.Fn3108
func Fn3108(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3110 github.com/goccy/pythonwasm2go/p2.Fn3110
func Fn3110(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3112 github.com/goccy/pythonwasm2go/p0.Fn3112
func Fn3112(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3113 github.com/goccy/pythonwasm2go/p2.Fn3113
func Fn3113(m *base.Module, l0 int32) int32

//go:linkname Fn3138 github.com/goccy/pythonwasm2go/p2.Fn3138
func Fn3138(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3152 github.com/goccy/pythonwasm2go/p2.Fn3152
func Fn3152(m *base.Module, l0 int32) int32

//go:linkname Fn3154 github.com/goccy/pythonwasm2go/p0.Fn3154
func Fn3154(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3163 github.com/goccy/pythonwasm2go/p2.Fn3163
func Fn3163(m *base.Module, l0 int32) int32

//go:linkname Fn3166 github.com/goccy/pythonwasm2go/p2.Fn3166
func Fn3166(m *base.Module, l0 int32) int32

//go:linkname Fn3170 github.com/goccy/pythonwasm2go/p2.Fn3170
func Fn3170(m *base.Module, l0 int32) int32

//go:linkname Fn3171 github.com/goccy/pythonwasm2go/p2.Fn3171
func Fn3171(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3215 github.com/goccy/pythonwasm2go/p2.Fn3215
func Fn3215(m *base.Module, l0 int32) int32

//go:linkname Fn3239 github.com/goccy/pythonwasm2go/p2.Fn3239
func Fn3239(m *base.Module, l0 int32) int32

//go:linkname Fn3241 github.com/goccy/pythonwasm2go/p2.Fn3241
func Fn3241(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3242 github.com/goccy/pythonwasm2go/p0.Fn3242
func Fn3242(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3246 github.com/goccy/pythonwasm2go/p0.Fn3246
func Fn3246(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3247 github.com/goccy/pythonwasm2go/p0.Fn3247
func Fn3247(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3253 github.com/goccy/pythonwasm2go/p2.Fn3253
func Fn3253(m *base.Module, l0 int32) int32

//go:linkname Fn3258 github.com/goccy/pythonwasm2go/p2.Fn3258
func Fn3258(m *base.Module, l0 int32) int32

//go:linkname Fn3270 github.com/goccy/pythonwasm2go/p2.Fn3270
func Fn3270(m *base.Module, l0 int32)

//go:linkname Fn3271 github.com/goccy/pythonwasm2go/p2.Fn3271
func Fn3271(m *base.Module, l0 int32)

//go:linkname Fn3272 github.com/goccy/pythonwasm2go/p2.Fn3272
func Fn3272(m *base.Module, l0 int32) int32

//go:linkname Fn3273 github.com/goccy/pythonwasm2go/p2.Fn3273
func Fn3273(m *base.Module, l0 int32) int32

//go:linkname Fn3274 github.com/goccy/pythonwasm2go/p2.Fn3274
func Fn3274(m *base.Module) int32

//go:linkname Fn3275 github.com/goccy/pythonwasm2go/p2.Fn3275
func Fn3275(m *base.Module) int32

//go:linkname Fn3276 github.com/goccy/pythonwasm2go/p2.Fn3276
func Fn3276(m *base.Module, l0 int32) int32

//go:linkname Fn3278 github.com/goccy/pythonwasm2go/p2.Fn3278
func Fn3278(m *base.Module, l0 int32) int32

//go:linkname Fn3279 github.com/goccy/pythonwasm2go/p2.Fn3279
func Fn3279(m *base.Module) int32

//go:linkname Fn3280 github.com/goccy/pythonwasm2go/p2.Fn3280
func Fn3280(m *base.Module, l0 int32) int32

//go:linkname Fn3282 github.com/goccy/pythonwasm2go/p2.Fn3282
func Fn3282(m *base.Module) int32

//go:linkname Fn3283 github.com/goccy/pythonwasm2go/p2.Fn3283
func Fn3283(m *base.Module, l0 int32) int32

//go:linkname Fn3284 github.com/goccy/pythonwasm2go/p2.Fn3284
func Fn3284(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3287 github.com/goccy/pythonwasm2go/p2.Fn3287
func Fn3287(m *base.Module, l0 int32) int32

//go:linkname Fn3288 github.com/goccy/pythonwasm2go/p2.Fn3288
func Fn3288(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3289 github.com/goccy/pythonwasm2go/p2.Fn3289
func Fn3289(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3290 github.com/goccy/pythonwasm2go/p2.Fn3290
func Fn3290(m *base.Module, l0 int32) int32

//go:linkname Fn3293 github.com/goccy/pythonwasm2go/p2.Fn3293
func Fn3293(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3300 github.com/goccy/pythonwasm2go/p2.Fn3300
func Fn3300(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3303 github.com/goccy/pythonwasm2go/p2.Fn3303
func Fn3303(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3305 github.com/goccy/pythonwasm2go/p0.Fn3305
func Fn3305(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3308 github.com/goccy/pythonwasm2go/p0.Fn3308
func Fn3308(m *base.Module, l0 int32) int32

//go:linkname Fn3309 github.com/goccy/pythonwasm2go/p0.Fn3309
func Fn3309(m *base.Module, l0 int32) int32

//go:linkname Fn3312 github.com/goccy/pythonwasm2go/p2.Fn3312
func Fn3312(m *base.Module, l0 int32)

//go:linkname Fn3314 github.com/goccy/pythonwasm2go/p2.Fn3314
func Fn3314(m *base.Module, l0 int32) int32

//go:linkname Fn3316 github.com/goccy/pythonwasm2go/p2.Fn3316
func Fn3316(m *base.Module, l0 int32) int32

//go:linkname Fn3349 github.com/goccy/pythonwasm2go/p0.Fn3349
func Fn3349(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3358 github.com/goccy/pythonwasm2go/p2.Fn3358
func Fn3358(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3377 github.com/goccy/pythonwasm2go/p2.Fn3377
func Fn3377(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3382 github.com/goccy/pythonwasm2go/p2.Fn3382
func Fn3382(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3383 github.com/goccy/pythonwasm2go/p2.Fn3383
func Fn3383(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3399 github.com/goccy/pythonwasm2go/p2.Fn3399
func Fn3399(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3416 github.com/goccy/pythonwasm2go/p0.Fn3416
func Fn3416(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn3421 github.com/goccy/pythonwasm2go/p2.Fn3421
func Fn3421(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3422 github.com/goccy/pythonwasm2go/p2.Fn3422
func Fn3422(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3423 github.com/goccy/pythonwasm2go/p2.Fn3423
func Fn3423(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3426 github.com/goccy/pythonwasm2go/p2.Fn3426
func Fn3426(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3429 github.com/goccy/pythonwasm2go/p0.Fn3429
func Fn3429(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3442 github.com/goccy/pythonwasm2go/p2.Fn3442
func Fn3442(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3443 github.com/goccy/pythonwasm2go/p2.Fn3443
func Fn3443(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3448 github.com/goccy/pythonwasm2go/p2.Fn3448
func Fn3448(m *base.Module, l0 int32)

//go:linkname Fn3450 github.com/goccy/pythonwasm2go/p2.Fn3450
func Fn3450(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3452 github.com/goccy/pythonwasm2go/p2.Fn3452
func Fn3452(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3454 github.com/goccy/pythonwasm2go/p2.Fn3454
func Fn3454(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3457 github.com/goccy/pythonwasm2go/p2.Fn3457
func Fn3457(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3458 github.com/goccy/pythonwasm2go/p2.Fn3458
func Fn3458(m *base.Module, l0 int32) int32

//go:linkname Fn3459 github.com/goccy/pythonwasm2go/p2.Fn3459
func Fn3459(m *base.Module, l0 int32) int32

//go:linkname Fn3461 github.com/goccy/pythonwasm2go/p2.Fn3461
func Fn3461(m *base.Module, l0 int32) int32

//go:linkname Fn3462 github.com/goccy/pythonwasm2go/p2.Fn3462
func Fn3462(m *base.Module, l0 int32) int32

//go:linkname Fn3463 github.com/goccy/pythonwasm2go/p2.Fn3463
func Fn3463(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3465 github.com/goccy/pythonwasm2go/p2.Fn3465
func Fn3465(m *base.Module, l0 int32) int32

//go:linkname Fn3466 github.com/goccy/pythonwasm2go/p2.Fn3466
func Fn3466(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3468 github.com/goccy/pythonwasm2go/p2.Fn3468
func Fn3468(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3469 github.com/goccy/pythonwasm2go/p2.Fn3469
func Fn3469(m *base.Module, l0 int32)

//go:linkname Fn3470 github.com/goccy/pythonwasm2go/p2.Fn3470
func Fn3470(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3476 github.com/goccy/pythonwasm2go/p2.Fn3476
func Fn3476(m *base.Module) int32

//go:linkname Fn3477 github.com/goccy/pythonwasm2go/p2.Fn3477
func Fn3477(m *base.Module) int32

//go:linkname Fn3480 github.com/goccy/pythonwasm2go/p2.Fn3480
func Fn3480(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3482 github.com/goccy/pythonwasm2go/p2.Fn3482
func Fn3482(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3488 github.com/goccy/pythonwasm2go/p2.Fn3488
func Fn3488(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3520 github.com/goccy/pythonwasm2go/p2.Fn3520
func Fn3520(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3523 github.com/goccy/pythonwasm2go/p2.Fn3523
func Fn3523(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3529 github.com/goccy/pythonwasm2go/p2.Fn3529
func Fn3529(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3531 github.com/goccy/pythonwasm2go/p2.Fn3531
func Fn3531(m *base.Module, l0 int32) int32

//go:linkname Fn3532 github.com/goccy/pythonwasm2go/p2.Fn3532
func Fn3532(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3533 github.com/goccy/pythonwasm2go/p2.Fn3533
func Fn3533(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3535 github.com/goccy/pythonwasm2go/p2.Fn3535
func Fn3535(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3537 github.com/goccy/pythonwasm2go/p2.Fn3537
func Fn3537(m *base.Module) int32

//go:linkname Fn3538 github.com/goccy/pythonwasm2go/p2.Fn3538
func Fn3538(m *base.Module, l0 int32)

//go:linkname Fn3540 github.com/goccy/pythonwasm2go/p2.Fn3540
func Fn3540(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3542 github.com/goccy/pythonwasm2go/p2.Fn3542
func Fn3542(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3545 github.com/goccy/pythonwasm2go/p2.Fn3545
func Fn3545(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3547 github.com/goccy/pythonwasm2go/p2.Fn3547
func Fn3547(m *base.Module, l0 int32) int32

//go:linkname Fn3559 github.com/goccy/pythonwasm2go/p2.Fn3559
func Fn3559(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3560 github.com/goccy/pythonwasm2go/p2.Fn3560
func Fn3560(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3568 github.com/goccy/pythonwasm2go/p0.Fn3568
func Fn3568(m *base.Module) int32

//go:linkname Fn3569 github.com/goccy/pythonwasm2go/p2.Fn3569
func Fn3569(m *base.Module, l0 int32)

//go:linkname Fn3570 github.com/goccy/pythonwasm2go/p2.Fn3570
func Fn3570(m *base.Module, l0 int32)

//go:linkname Fn3571 github.com/goccy/pythonwasm2go/p2.Fn3571
func Fn3571(m *base.Module, l0 int32)

//go:linkname Fn3576 github.com/goccy/pythonwasm2go/p0.Fn3576
func Fn3576(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3578 github.com/goccy/pythonwasm2go/p2.Fn3578
func Fn3578(m *base.Module, l0 int32) int32

//go:linkname Fn3579 github.com/goccy/pythonwasm2go/p0.Fn3579
func Fn3579(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3580 github.com/goccy/pythonwasm2go/p2.Fn3580
func Fn3580(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3581 github.com/goccy/pythonwasm2go/p2.Fn3581
func Fn3581(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3582 github.com/goccy/pythonwasm2go/p2.Fn3582
func Fn3582(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3583 github.com/goccy/pythonwasm2go/p0.Fn3583
func Fn3583(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3585 github.com/goccy/pythonwasm2go/p2.Fn3585
func Fn3585(m *base.Module, l0 int32)

//go:linkname Fn3586 github.com/goccy/pythonwasm2go/p2.Fn3586
func Fn3586(m *base.Module, l0 int32)

//go:linkname Fn3587 github.com/goccy/pythonwasm2go/p2.Fn3587
func Fn3587(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3590 github.com/goccy/pythonwasm2go/p0.Fn3590
func Fn3590(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3600 github.com/goccy/pythonwasm2go/p2.Fn3600
func Fn3600(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3604 github.com/goccy/pythonwasm2go/p0.Fn3604
func Fn3604(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3606 github.com/goccy/pythonwasm2go/p2.Fn3606
func Fn3606(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3609 github.com/goccy/pythonwasm2go/p0.Fn3609
func Fn3609(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3610 github.com/goccy/pythonwasm2go/p2.Fn3610
func Fn3610(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3611 github.com/goccy/pythonwasm2go/p2.Fn3611
func Fn3611(m *base.Module, l0 int32)

//go:linkname Fn3612 github.com/goccy/pythonwasm2go/p2.Fn3612
func Fn3612(m *base.Module, l0 int32) int32

//go:linkname Fn3613 github.com/goccy/pythonwasm2go/p0.Fn3613
func Fn3613(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3614 github.com/goccy/pythonwasm2go/p0.Fn3614
func Fn3614(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3615 github.com/goccy/pythonwasm2go/p2.Fn3615
func Fn3615(m *base.Module, l0 int32)

//go:linkname Fn3617 github.com/goccy/pythonwasm2go/p2.Fn3617
func Fn3617(m *base.Module, l0 int32) int32

//go:linkname Fn3618 github.com/goccy/pythonwasm2go/p0.Fn3618
func Fn3618(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3619 github.com/goccy/pythonwasm2go/p0.Fn3619
func Fn3619(m *base.Module, l0 int32)

//go:linkname Fn3620 github.com/goccy/pythonwasm2go/p2.Fn3620
func Fn3620(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3621 github.com/goccy/pythonwasm2go/p2.Fn3621
func Fn3621(m *base.Module, l0 int32)

//go:linkname Fn3622 github.com/goccy/pythonwasm2go/p0.Fn3622
func Fn3622(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3623 github.com/goccy/pythonwasm2go/p0.Fn3623
func Fn3623(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3624 github.com/goccy/pythonwasm2go/p2.Fn3624
func Fn3624(m *base.Module) int32

//go:linkname Fn3625 github.com/goccy/pythonwasm2go/p0.Fn3625
func Fn3625(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3626 github.com/goccy/pythonwasm2go/p0.Fn3626
func Fn3626(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3627 github.com/goccy/pythonwasm2go/p0.Fn3627
func Fn3627(m *base.Module, l0 int32) int32

//go:linkname Fn3630 github.com/goccy/pythonwasm2go/p2.Fn3630
func Fn3630(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3631 github.com/goccy/pythonwasm2go/p2.Fn3631
func Fn3631(m *base.Module) int32

//go:linkname Fn3632 github.com/goccy/pythonwasm2go/p2.Fn3632
func Fn3632(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3633 github.com/goccy/pythonwasm2go/p2.Fn3633
func Fn3633(m *base.Module)

//go:linkname Fn3634 github.com/goccy/pythonwasm2go/p2.Fn3634
func Fn3634(m *base.Module, l0 int32) int32

//go:linkname Fn3635 github.com/goccy/pythonwasm2go/p2.Fn3635
func Fn3635(m *base.Module, l0 int32) int32

//go:linkname Fn3636 github.com/goccy/pythonwasm2go/p2.Fn3636
func Fn3636(m *base.Module, l0 int32)

//go:linkname Fn3638 github.com/goccy/pythonwasm2go/p2.Fn3638
func Fn3638(m *base.Module, l0 int32)

//go:linkname Fn3641 github.com/goccy/pythonwasm2go/p2.Fn3641
func Fn3641(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3642 github.com/goccy/pythonwasm2go/p0.Fn3642
func Fn3642(m *base.Module)

//go:linkname Fn3643 github.com/goccy/pythonwasm2go/p0.Fn3643
func Fn3643(m *base.Module) int32

//go:linkname Fn3644 github.com/goccy/pythonwasm2go/p2.Fn3644
func Fn3644(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3645 github.com/goccy/pythonwasm2go/p0.Fn3645
func Fn3645(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3646 github.com/goccy/pythonwasm2go/p2.Fn3646
func Fn3646(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3647 github.com/goccy/pythonwasm2go/p0.Fn3647
func Fn3647(m *base.Module, l0 int32) int32

//go:linkname Fn3650 github.com/goccy/pythonwasm2go/p2.Fn3650
func Fn3650(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3651 github.com/goccy/pythonwasm2go/p0.Fn3651
func Fn3651(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3652 github.com/goccy/pythonwasm2go/p0.Fn3652
func Fn3652(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3654 github.com/goccy/pythonwasm2go/p2.Fn3654
func Fn3654(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3656 github.com/goccy/pythonwasm2go/p0.Fn3656
func Fn3656(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3659 github.com/goccy/pythonwasm2go/p2.Fn3659
func Fn3659(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3660 github.com/goccy/pythonwasm2go/p2.Fn3660
func Fn3660(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn3664 github.com/goccy/pythonwasm2go/p2.Fn3664
func Fn3664(m *base.Module, l0 int32) int32

//go:linkname Fn3665 github.com/goccy/pythonwasm2go/p2.Fn3665
func Fn3665(m *base.Module, l0 int32)

//go:linkname Fn3666 github.com/goccy/pythonwasm2go/p2.Fn3666
func Fn3666(m *base.Module, l0 int32) int32

//go:linkname Fn3667 github.com/goccy/pythonwasm2go/p2.Fn3667
func Fn3667(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3669 github.com/goccy/pythonwasm2go/p0.Fn3669
func Fn3669(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3670 github.com/goccy/pythonwasm2go/p2.Fn3670
func Fn3670(m *base.Module, l0 int32) int32

//go:linkname Fn3674 github.com/goccy/pythonwasm2go/p2.Fn3674
func Fn3674(m *base.Module, l0 int32) int32

//go:linkname Fn3675 github.com/goccy/pythonwasm2go/p2.Fn3675
func Fn3675(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3691 github.com/goccy/pythonwasm2go/p2.Fn3691
func Fn3691(m *base.Module, l0 int32) int32

//go:linkname Fn3692 github.com/goccy/pythonwasm2go/p2.Fn3692
func Fn3692(m *base.Module, l0 int32, l1 int64, l2 int32)

//go:linkname Fn3695 github.com/goccy/pythonwasm2go/p2.Fn3695
func Fn3695(m *base.Module, l0 int32) int32

//go:linkname Fn3696 github.com/goccy/pythonwasm2go/p2.Fn3696
func Fn3696(m *base.Module, l0 int32) int32

//go:linkname Fn3697 github.com/goccy/pythonwasm2go/p2.Fn3697
func Fn3697(m *base.Module, l0 int32) int32

//go:linkname Fn3698 github.com/goccy/pythonwasm2go/p2.Fn3698
func Fn3698(m *base.Module, l0 int32) int32

//go:linkname Fn3701 github.com/goccy/pythonwasm2go/p0.Fn3701
func Fn3701(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3705 github.com/goccy/pythonwasm2go/p0.Fn3705
func Fn3705(m *base.Module, l0 int32) int32

//go:linkname Fn3706 github.com/goccy/pythonwasm2go/p2.Fn3706
func Fn3706(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3707 github.com/goccy/pythonwasm2go/p2.Fn3707
func Fn3707(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3712 github.com/goccy/pythonwasm2go/p0.Fn3712
func Fn3712(m *base.Module, l0 int32) int32

//go:linkname Fn3721 github.com/goccy/pythonwasm2go/p2.Fn3721
func Fn3721(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3735 github.com/goccy/pythonwasm2go/p0.Fn3735
func Fn3735(m *base.Module, l0 int32)

//go:linkname Fn3736 github.com/goccy/pythonwasm2go/p2.Fn3736
func Fn3736(m *base.Module, l0 int32)

//go:linkname Fn3738 github.com/goccy/pythonwasm2go/p0.Fn3738
func Fn3738(m *base.Module, l0 int32)

//go:linkname Fn3739 github.com/goccy/pythonwasm2go/p0.Fn3739
func Fn3739(m *base.Module, l0 int32) int32

//go:linkname Fn3740 github.com/goccy/pythonwasm2go/p2.Fn3740
func Fn3740(m *base.Module, l0 int32) int32

//go:linkname Fn3742 github.com/goccy/pythonwasm2go/p0.Fn3742
func Fn3742(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3743 github.com/goccy/pythonwasm2go/p2.Fn3743
func Fn3743(m *base.Module, l0 int32)

//go:linkname Fn3744 github.com/goccy/pythonwasm2go/p2.Fn3744
func Fn3744(m *base.Module, l0 int32) int32

//go:linkname Fn3748 github.com/goccy/pythonwasm2go/p2.Fn3748
func Fn3748(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3750 github.com/goccy/pythonwasm2go/p0.Fn3750
func Fn3750(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3751 github.com/goccy/pythonwasm2go/p2.Fn3751
func Fn3751(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3752 github.com/goccy/pythonwasm2go/p2.Fn3752
func Fn3752(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3753 github.com/goccy/pythonwasm2go/p0.Fn3753
func Fn3753(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3754 github.com/goccy/pythonwasm2go/p0.Fn3754
func Fn3754(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn3755 github.com/goccy/pythonwasm2go/p0.Fn3755
func Fn3755(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3756 github.com/goccy/pythonwasm2go/p0.Fn3756
func Fn3756(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3757 github.com/goccy/pythonwasm2go/p2.Fn3757
func Fn3757(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3761 github.com/goccy/pythonwasm2go/p2.Fn3761
func Fn3761(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3762 github.com/goccy/pythonwasm2go/p2.Fn3762
func Fn3762(m *base.Module, l0 int32) int32

//go:linkname Fn3763 github.com/goccy/pythonwasm2go/p2.Fn3763
func Fn3763(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3769 github.com/goccy/pythonwasm2go/p2.Fn3769
func Fn3769(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3772 github.com/goccy/pythonwasm2go/p2.Fn3772
func Fn3772(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3773 github.com/goccy/pythonwasm2go/p2.Fn3773
func Fn3773(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3786 github.com/goccy/pythonwasm2go/p0.Fn3786
func Fn3786(m *base.Module) int32

//go:linkname Fn3787 github.com/goccy/pythonwasm2go/p0.Fn3787
func Fn3787(m *base.Module, l0 int32)

//go:linkname Fn3788 github.com/goccy/pythonwasm2go/p2.Fn3788
func Fn3788(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3789 github.com/goccy/pythonwasm2go/p2.Fn3789
func Fn3789(m *base.Module, l0 int32) int32

//go:linkname Fn3791 github.com/goccy/pythonwasm2go/p0.Fn3791
func Fn3791(m *base.Module, l0 int32) int32

//go:linkname Fn3793 github.com/goccy/pythonwasm2go/p0.Fn3793
func Fn3793(m *base.Module, l0 int32) int32

//go:linkname Fn3794 github.com/goccy/pythonwasm2go/p0.Fn3794
func Fn3794(m *base.Module) int32

//go:linkname Fn3801 github.com/goccy/pythonwasm2go/p2.Fn3801
func Fn3801(m *base.Module, l0 int32) int32

//go:linkname Fn3802 github.com/goccy/pythonwasm2go/p0.Fn3802
func Fn3802(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3803 github.com/goccy/pythonwasm2go/p2.Fn3803
func Fn3803(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3804 github.com/goccy/pythonwasm2go/p2.Fn3804
func Fn3804(m *base.Module, l0 int32) int32

//go:linkname Fn3805 github.com/goccy/pythonwasm2go/p2.Fn3805
func Fn3805(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3808 github.com/goccy/pythonwasm2go/p2.Fn3808
func Fn3808(m *base.Module, l0 int32) int32

//go:linkname Fn3809 github.com/goccy/pythonwasm2go/p0.Fn3809
func Fn3809(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3810 github.com/goccy/pythonwasm2go/p0.Fn3810
func Fn3810(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3812 github.com/goccy/pythonwasm2go/p2.Fn3812
func Fn3812(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3844 github.com/goccy/pythonwasm2go/p2.Fn3844
func Fn3844(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3846 github.com/goccy/pythonwasm2go/p2.Fn3846
func Fn3846(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3847 github.com/goccy/pythonwasm2go/p2.Fn3847
func Fn3847(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3848 github.com/goccy/pythonwasm2go/p2.Fn3848
func Fn3848(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3849 github.com/goccy/pythonwasm2go/p2.Fn3849
func Fn3849(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3851 github.com/goccy/pythonwasm2go/p2.Fn3851
func Fn3851(m *base.Module, l0 int32) int32

//go:linkname Fn3852 github.com/goccy/pythonwasm2go/p2.Fn3852
func Fn3852(m *base.Module, l0 int32)

//go:linkname Fn3854 github.com/goccy/pythonwasm2go/p2.Fn3854
func Fn3854(m *base.Module, l0 int32)

//go:linkname Fn3856 github.com/goccy/pythonwasm2go/p0.Fn3856
func Fn3856(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3857 github.com/goccy/pythonwasm2go/p2.Fn3857
func Fn3857(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3858 github.com/goccy/pythonwasm2go/p0.Fn3858
func Fn3858(m *base.Module, l0 int32) int32

//go:linkname Fn3863 github.com/goccy/pythonwasm2go/p2.Fn3863
func Fn3863(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3865 github.com/goccy/pythonwasm2go/p2.Fn3865
func Fn3865(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3866 github.com/goccy/pythonwasm2go/p2.Fn3866
func Fn3866(m *base.Module, l0 int32) int32

//go:linkname Fn3867 github.com/goccy/pythonwasm2go/p2.Fn3867
func Fn3867(m *base.Module, l0 int32) int32

//go:linkname Fn3868 github.com/goccy/pythonwasm2go/p2.Fn3868
func Fn3868(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3869 github.com/goccy/pythonwasm2go/p2.Fn3869
func Fn3869(m *base.Module, l0 int32) int32

//go:linkname Fn3870 github.com/goccy/pythonwasm2go/p2.Fn3870
func Fn3870(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3871 github.com/goccy/pythonwasm2go/p2.Fn3871
func Fn3871(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3872 github.com/goccy/pythonwasm2go/p2.Fn3872
func Fn3872(m *base.Module, l0 int32) int32

//go:linkname Fn3875 github.com/goccy/pythonwasm2go/p2.Fn3875
func Fn3875(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3876 github.com/goccy/pythonwasm2go/p2.Fn3876
func Fn3876(m *base.Module)

//go:linkname Fn3877 github.com/goccy/pythonwasm2go/p2.Fn3877
func Fn3877(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3881 github.com/goccy/pythonwasm2go/p2.Fn3881
func Fn3881(m *base.Module, l0 int32)

//go:linkname Fn3884 github.com/goccy/pythonwasm2go/p2.Fn3884
func Fn3884(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3885 github.com/goccy/pythonwasm2go/p2.Fn3885
func Fn3885(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3889 github.com/goccy/pythonwasm2go/p2.Fn3889
func Fn3889(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3890 github.com/goccy/pythonwasm2go/p2.Fn3890
func Fn3890(m *base.Module, l0 int32) int32

//go:linkname Fn3891 github.com/goccy/pythonwasm2go/p2.Fn3891
func Fn3891(m *base.Module) int32

//go:linkname Fn3893 github.com/goccy/pythonwasm2go/p0.Fn3893
func Fn3893(m *base.Module, l0 int32) int32

//go:linkname Fn3894 github.com/goccy/pythonwasm2go/p0.Fn3894
func Fn3894(m *base.Module, l0 int32) int32

//go:linkname Fn3901 github.com/goccy/pythonwasm2go/p0.Fn3901
func Fn3901(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3905 github.com/goccy/pythonwasm2go/p2.Fn3905
func Fn3905(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3928 github.com/goccy/pythonwasm2go/p2.Fn3928
func Fn3928(m *base.Module, l0 int32)

//go:linkname Fn3929 github.com/goccy/pythonwasm2go/p2.Fn3929
func Fn3929(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3933 github.com/goccy/pythonwasm2go/p2.Fn3933
func Fn3933(m *base.Module, l0 int32) int32

//go:linkname Fn3936 github.com/goccy/pythonwasm2go/p2.Fn3936
func Fn3936(m *base.Module, l0 int32) int32

//go:linkname Fn3937 github.com/goccy/pythonwasm2go/p2.Fn3937
func Fn3937(m *base.Module, l0 int32)

//go:linkname Fn3938 github.com/goccy/pythonwasm2go/p2.Fn3938
func Fn3938(m *base.Module, l0 int32)

//go:linkname Fn3941 github.com/goccy/pythonwasm2go/p2.Fn3941
func Fn3941(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3942 github.com/goccy/pythonwasm2go/p2.Fn3942
func Fn3942(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3944 github.com/goccy/pythonwasm2go/p2.Fn3944
func Fn3944(m *base.Module, l0 int32) int32

//go:linkname Fn3946 github.com/goccy/pythonwasm2go/p2.Fn3946
func Fn3946(m *base.Module, l0 int32)

//go:linkname Fn3950 github.com/goccy/pythonwasm2go/p2.Fn3950
func Fn3950(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3957 github.com/goccy/pythonwasm2go/p2.Fn3957
func Fn3957(m *base.Module, l0 int32)

//go:linkname Fn3971 github.com/goccy/pythonwasm2go/p2.Fn3971
func Fn3971(m *base.Module, l0 int32) int32

//go:linkname Fn3972 github.com/goccy/pythonwasm2go/p2.Fn3972
func Fn3972(m *base.Module, l0 int32) int32

//go:linkname Fn3973 github.com/goccy/pythonwasm2go/p2.Fn3973
func Fn3973(m *base.Module, l0 int32) int32

//go:linkname Fn3974 github.com/goccy/pythonwasm2go/p2.Fn3974
func Fn3974(m *base.Module, l0 int32)

//go:linkname Fn3979 github.com/goccy/pythonwasm2go/p2.Fn3979
func Fn3979(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3980 github.com/goccy/pythonwasm2go/p2.Fn3980
func Fn3980(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3982 github.com/goccy/pythonwasm2go/p2.Fn3982
func Fn3982(m *base.Module, l0 int32) int32

//go:linkname Fn3983 github.com/goccy/pythonwasm2go/p2.Fn3983
func Fn3983(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3984 github.com/goccy/pythonwasm2go/p2.Fn3984
func Fn3984(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4002 github.com/goccy/pythonwasm2go/p2.Fn4002
func Fn4002(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4009 github.com/goccy/pythonwasm2go/p2.Fn4009
func Fn4009(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4011 github.com/goccy/pythonwasm2go/p2.Fn4011
func Fn4011(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4012 github.com/goccy/pythonwasm2go/p2.Fn4012
func Fn4012(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4015 github.com/goccy/pythonwasm2go/p2.Fn4015
func Fn4015(m *base.Module, l0 int32) int32

//go:linkname Fn4029 github.com/goccy/pythonwasm2go/p2.Fn4029
func Fn4029(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4030 github.com/goccy/pythonwasm2go/p2.Fn4030
func Fn4030(m *base.Module, l0 int32) int32

//go:linkname Fn4031 github.com/goccy/pythonwasm2go/p2.Fn4031
func Fn4031(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4032 github.com/goccy/pythonwasm2go/p2.Fn4032
func Fn4032(m *base.Module, l0 int32) int32

//go:linkname Fn4033 github.com/goccy/pythonwasm2go/p2.Fn4033
func Fn4033(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4034 github.com/goccy/pythonwasm2go/p2.Fn4034
func Fn4034(m *base.Module, l0 int32) int32

//go:linkname Fn4036 github.com/goccy/pythonwasm2go/p2.Fn4036
func Fn4036(m *base.Module, l0 int32)

//go:linkname Fn4037 github.com/goccy/pythonwasm2go/p2.Fn4037
func Fn4037(m *base.Module) int32

//go:linkname Fn4063 github.com/goccy/pythonwasm2go/p2.Fn4063
func Fn4063(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4073 github.com/goccy/pythonwasm2go/p2.Fn4073
func Fn4073(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4074 github.com/goccy/pythonwasm2go/p2.Fn4074
func Fn4074(m *base.Module, l0 int32) int32

//go:linkname Fn4075 github.com/goccy/pythonwasm2go/p2.Fn4075
func Fn4075(m *base.Module, l0 int32) int32

//go:linkname Fn4090 github.com/goccy/pythonwasm2go/p0.Fn4090
func Fn4090(m *base.Module, l0 int32)

//go:linkname Fn4091 github.com/goccy/pythonwasm2go/p2.Fn4091
func Fn4091(m *base.Module, l0 int32)

//go:linkname Fn4094 github.com/goccy/pythonwasm2go/p2.Fn4094
func Fn4094(m *base.Module, l0 int32) int32

//go:linkname Fn4095 github.com/goccy/pythonwasm2go/p0.Fn4095
func Fn4095(m *base.Module, l0 int32)

//go:linkname Fn4096 github.com/goccy/pythonwasm2go/p0.Fn4096
func Fn4096(m *base.Module, l0 int32)

//go:linkname Fn4097 github.com/goccy/pythonwasm2go/p2.Fn4097
func Fn4097(m *base.Module, l0 int32)

//go:linkname Fn4098 github.com/goccy/pythonwasm2go/p0.Fn4098
func Fn4098(m *base.Module, l0 int32) int32

//go:linkname Fn4099 github.com/goccy/pythonwasm2go/p0.Fn4099
func Fn4099(m *base.Module, l0 int32)

//go:linkname Fn4103 github.com/goccy/pythonwasm2go/p2.Fn4103
func Fn4103(m *base.Module, l0 int32)

//go:linkname Fn4104 github.com/goccy/pythonwasm2go/p2.Fn4104
func Fn4104(m *base.Module, l0 int32)

//go:linkname Fn4106 github.com/goccy/pythonwasm2go/p2.Fn4106
func Fn4106(m *base.Module, l0 int32)

//go:linkname Fn4108 github.com/goccy/pythonwasm2go/p2.Fn4108
func Fn4108(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4109 github.com/goccy/pythonwasm2go/p0.Fn4109
func Fn4109(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4110 github.com/goccy/pythonwasm2go/p2.Fn4110
func Fn4110(m *base.Module, l0 int32)

//go:linkname Fn4116 github.com/goccy/pythonwasm2go/p0.Fn4116
func Fn4116(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4124 github.com/goccy/pythonwasm2go/p2.Fn4124
func Fn4124(m *base.Module, l0 int32) int32

//go:linkname Fn4139 github.com/goccy/pythonwasm2go/p2.Fn4139
func Fn4139(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4140 github.com/goccy/pythonwasm2go/p0.Fn4140
func Fn4140(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4148 github.com/goccy/pythonwasm2go/p2.Fn4148
func Fn4148(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4149 github.com/goccy/pythonwasm2go/p2.Fn4149
func Fn4149(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4151 github.com/goccy/pythonwasm2go/p2.Fn4151
func Fn4151(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4152 github.com/goccy/pythonwasm2go/p2.Fn4152
func Fn4152(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4153 github.com/goccy/pythonwasm2go/p2.Fn4153
func Fn4153(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4155 github.com/goccy/pythonwasm2go/p2.Fn4155
func Fn4155(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4161 github.com/goccy/pythonwasm2go/p0.Fn4161
func Fn4161(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn4169 github.com/goccy/pythonwasm2go/p2.Fn4169
func Fn4169(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4170 github.com/goccy/pythonwasm2go/p2.Fn4170
func Fn4170(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4172 github.com/goccy/pythonwasm2go/p2.Fn4172
func Fn4172(m *base.Module, l0 int32)

//go:linkname Fn4178 github.com/goccy/pythonwasm2go/p2.Fn4178
func Fn4178(m *base.Module) int32

//go:linkname Fn4181 github.com/goccy/pythonwasm2go/p2.Fn4181
func Fn4181(m *base.Module, l0 int32)

//go:linkname Fn4183 github.com/goccy/pythonwasm2go/p2.Fn4183
func Fn4183(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4185 github.com/goccy/pythonwasm2go/p2.Fn4185
func Fn4185(m *base.Module, l0 int32) int32

//go:linkname Fn4186 github.com/goccy/pythonwasm2go/p2.Fn4186
func Fn4186(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4191 github.com/goccy/pythonwasm2go/p2.Fn4191
func Fn4191(m *base.Module, l0 int32) int32

//go:linkname Fn4192 github.com/goccy/pythonwasm2go/p2.Fn4192
func Fn4192(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4203 github.com/goccy/pythonwasm2go/p2.Fn4203
func Fn4203(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4205 github.com/goccy/pythonwasm2go/p0.Fn4205
func Fn4205(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4207 github.com/goccy/pythonwasm2go/p0.Fn4207
func Fn4207(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4208 github.com/goccy/pythonwasm2go/p2.Fn4208
func Fn4208(m *base.Module, l0 int32)

//go:linkname Fn4212 github.com/goccy/pythonwasm2go/p0.Fn4212
func Fn4212(m *base.Module, l0 int32)

//go:linkname Fn4213 github.com/goccy/pythonwasm2go/p0.Fn4213
func Fn4213(m *base.Module, l0 int32)

//go:linkname Fn4218 github.com/goccy/pythonwasm2go/p0.Fn4218
func Fn4218(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4219 github.com/goccy/pythonwasm2go/p2.Fn4219
func Fn4219(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4221 github.com/goccy/pythonwasm2go/p2.Fn4221
func Fn4221(m *base.Module) int32

//go:linkname Fn4222 github.com/goccy/pythonwasm2go/p2.Fn4222
func Fn4222(m *base.Module, l0 int32) int32

//go:linkname Fn4223 github.com/goccy/pythonwasm2go/p0.Fn4223
func Fn4223(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4226 github.com/goccy/pythonwasm2go/p2.Fn4226
func Fn4226(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4230 github.com/goccy/pythonwasm2go/p2.Fn4230
func Fn4230(m *base.Module)

//go:linkname Fn4232 github.com/goccy/pythonwasm2go/p2.Fn4232
func Fn4232(m *base.Module, l0 int32)

//go:linkname Fn4233 github.com/goccy/pythonwasm2go/p0.Fn4233
func Fn4233(m *base.Module, l0 int32)

//go:linkname Fn4234 github.com/goccy/pythonwasm2go/p0.Fn4234
func Fn4234(m *base.Module, l0 int32)

//go:linkname Fn4235 github.com/goccy/pythonwasm2go/p0.Fn4235
func Fn4235(m *base.Module, l0 int32)

//go:linkname Fn4238 github.com/goccy/pythonwasm2go/p0.Fn4238
func Fn4238(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4240 github.com/goccy/pythonwasm2go/p2.Fn4240
func Fn4240(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4241 github.com/goccy/pythonwasm2go/p2.Fn4241
func Fn4241(m *base.Module, l0 int32) int32

//go:linkname Fn4242 github.com/goccy/pythonwasm2go/p2.Fn4242
func Fn4242(m *base.Module, l0 int32) int32

//go:linkname Fn4243 github.com/goccy/pythonwasm2go/p2.Fn4243
func Fn4243(m *base.Module, l0 int32) int32

//go:linkname Fn4245 github.com/goccy/pythonwasm2go/p2.Fn4245
func Fn4245(m *base.Module, l0 int32) int64

//go:linkname Fn4247 github.com/goccy/pythonwasm2go/p0.Fn4247
func Fn4247(m *base.Module, l0 int32) int32

//go:linkname Fn4251 github.com/goccy/pythonwasm2go/p0.Fn4251
func Fn4251(m *base.Module, l0 int32)

//go:linkname Fn4253 github.com/goccy/pythonwasm2go/p2.Fn4253
func Fn4253(m *base.Module) int32

//go:linkname Fn4256 github.com/goccy/pythonwasm2go/p2.Fn4256
func Fn4256(m *base.Module, l0 int32) int32

//go:linkname Fn4257 github.com/goccy/pythonwasm2go/p0.Fn4257
func Fn4257(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4258 github.com/goccy/pythonwasm2go/p2.Fn4258
func Fn4258(m *base.Module, l0 int32) int32

//go:linkname Fn4259 github.com/goccy/pythonwasm2go/p2.Fn4259
func Fn4259(m *base.Module, l0 int32)

//go:linkname Fn4260 github.com/goccy/pythonwasm2go/p0.Fn4260
func Fn4260(m *base.Module, l0 int32)

//go:linkname Fn4266 github.com/goccy/pythonwasm2go/p2.Fn4266
func Fn4266(m *base.Module) int32

//go:linkname Fn4267 github.com/goccy/pythonwasm2go/p0.Fn4267
func Fn4267(m *base.Module, l0 int32) int32

//go:linkname Fn4269 github.com/goccy/pythonwasm2go/p2.Fn4269
func Fn4269(m *base.Module) int32

//go:linkname Fn4270 github.com/goccy/pythonwasm2go/p0.Fn4270
func Fn4270(m *base.Module, l0 int32)

//go:linkname Fn4275 github.com/goccy/pythonwasm2go/p0.Fn4275
func Fn4275(m *base.Module) int32

//go:linkname Fn4276 github.com/goccy/pythonwasm2go/p2.Fn4276
func Fn4276(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4282 github.com/goccy/pythonwasm2go/p2.Fn4282
func Fn4282(m *base.Module, l0 int32)

//go:linkname Fn4284 github.com/goccy/pythonwasm2go/p2.Fn4284
func Fn4284(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)

//go:linkname Fn4287 github.com/goccy/pythonwasm2go/p2.Fn4287
func Fn4287(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4288 github.com/goccy/pythonwasm2go/p2.Fn4288
func Fn4288(m *base.Module, l0 int32)

//go:linkname Fn4291 github.com/goccy/pythonwasm2go/p2.Fn4291
func Fn4291(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4292 github.com/goccy/pythonwasm2go/p2.Fn4292
func Fn4292(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4298 github.com/goccy/pythonwasm2go/p2.Fn4298
func Fn4298(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4299 github.com/goccy/pythonwasm2go/p2.Fn4299
func Fn4299(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4300 github.com/goccy/pythonwasm2go/p2.Fn4300
func Fn4300(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4303 github.com/goccy/pythonwasm2go/p2.Fn4303
func Fn4303(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4304 github.com/goccy/pythonwasm2go/p2.Fn4304
func Fn4304(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4306 github.com/goccy/pythonwasm2go/p2.Fn4306
func Fn4306(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4307 github.com/goccy/pythonwasm2go/p2.Fn4307
func Fn4307(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4310 github.com/goccy/pythonwasm2go/p2.Fn4310
func Fn4310(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4327 github.com/goccy/pythonwasm2go/p0.Fn4327
func Fn4327(m *base.Module, l0 int32) int32

//go:linkname Fn4333 github.com/goccy/pythonwasm2go/p0.Fn4333
func Fn4333(m *base.Module, l0 int32) int32

//go:linkname Fn4338 github.com/goccy/pythonwasm2go/p0.Fn4338
func Fn4338(m *base.Module, l0 int32) int32

//go:linkname Fn4349 github.com/goccy/pythonwasm2go/p2.Fn4349
func Fn4349(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4363 github.com/goccy/pythonwasm2go/p0.Fn4363
func Fn4363(m *base.Module, l0 int32) int32

//go:linkname Fn4412 github.com/goccy/pythonwasm2go/p0.Fn4412
func Fn4412(m *base.Module, l0 int32) int32

//go:linkname Fn4416 github.com/goccy/pythonwasm2go/p0.Fn4416
func Fn4416(m *base.Module, l0 int32) int32

//go:linkname Fn4417 github.com/goccy/pythonwasm2go/p0.Fn4417
func Fn4417(m *base.Module, l0 int32) int32

//go:linkname Fn4420 github.com/goccy/pythonwasm2go/p2.Fn4420
func Fn4420(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4424 github.com/goccy/pythonwasm2go/p0.Fn4424
func Fn4424(m *base.Module, l0 int32) int32

//go:linkname Fn4462 github.com/goccy/pythonwasm2go/p0.Fn4462
func Fn4462(m *base.Module, l0 int32) int32

//go:linkname Fn4470 github.com/goccy/pythonwasm2go/p2.Fn4470
func Fn4470(m *base.Module, l0 int32) int32

//go:linkname Fn4471 github.com/goccy/pythonwasm2go/p2.Fn4471
func Fn4471(m *base.Module, l0 int32) int32

//go:linkname Fn4484 github.com/goccy/pythonwasm2go/p0.Fn4484
func Fn4484(m *base.Module, l0 int32) int32

//go:linkname Fn4488 github.com/goccy/pythonwasm2go/p2.Fn4488
func Fn4488(m *base.Module, l0 int32) int32

//go:linkname Fn4489 github.com/goccy/pythonwasm2go/p2.Fn4489
func Fn4489(m *base.Module, l0 int32) int32

//go:linkname Fn4491 github.com/goccy/pythonwasm2go/p2.Fn4491
func Fn4491(m *base.Module, l0 int32) int32

//go:linkname Fn4504 github.com/goccy/pythonwasm2go/p2.Fn4504
func Fn4504(m *base.Module, l0 int32) int32

//go:linkname Fn4509 github.com/goccy/pythonwasm2go/p2.Fn4509
func Fn4509(m *base.Module, l0 int32) int32

//go:linkname Fn4510 github.com/goccy/pythonwasm2go/p2.Fn4510
func Fn4510(m *base.Module, l0 int32) int32

//go:linkname Fn4513 github.com/goccy/pythonwasm2go/p0.Fn4513
func Fn4513(m *base.Module, l0 int32) int32

//go:linkname Fn4514 github.com/goccy/pythonwasm2go/p0.Fn4514
func Fn4514(m *base.Module, l0 int32) int32

//go:linkname Fn4532 github.com/goccy/pythonwasm2go/p2.Fn4532
func Fn4532(m *base.Module, l0 int32) int32

//go:linkname Fn4533 github.com/goccy/pythonwasm2go/p2.Fn4533
func Fn4533(m *base.Module, l0 int32) int32

//go:linkname Fn4541 github.com/goccy/pythonwasm2go/p2.Fn4541
func Fn4541(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4542 github.com/goccy/pythonwasm2go/p2.Fn4542
func Fn4542(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4543 github.com/goccy/pythonwasm2go/p2.Fn4543
func Fn4543(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4545 github.com/goccy/pythonwasm2go/p2.Fn4545
func Fn4545(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4547 github.com/goccy/pythonwasm2go/p2.Fn4547
func Fn4547(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4548 github.com/goccy/pythonwasm2go/p2.Fn4548
func Fn4548(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4549 github.com/goccy/pythonwasm2go/p2.Fn4549
func Fn4549(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4550 github.com/goccy/pythonwasm2go/p2.Fn4550
func Fn4550(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4553 github.com/goccy/pythonwasm2go/p2.Fn4553
func Fn4553(m *base.Module, l0 int32) int32

//go:linkname Fn4554 github.com/goccy/pythonwasm2go/p2.Fn4554
func Fn4554(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4555 github.com/goccy/pythonwasm2go/p2.Fn4555
func Fn4555(m *base.Module, l0 int32) int32

//go:linkname Fn4557 github.com/goccy/pythonwasm2go/p2.Fn4557
func Fn4557(m *base.Module, l0 int32) int32

//go:linkname Fn4561 github.com/goccy/pythonwasm2go/p2.Fn4561
func Fn4561(m *base.Module, l0 int32)

//go:linkname Fn4564 github.com/goccy/pythonwasm2go/p2.Fn4564
func Fn4564(m *base.Module, l0 int32) int32

//go:linkname Fn4566 github.com/goccy/pythonwasm2go/p2.Fn4566
func Fn4566(m *base.Module, l0 int32)

//go:linkname Fn4569 github.com/goccy/pythonwasm2go/p2.Fn4569
func Fn4569(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4570 github.com/goccy/pythonwasm2go/p2.Fn4570
func Fn4570(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn4573 github.com/goccy/pythonwasm2go/p2.Fn4573
func Fn4573(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4574 github.com/goccy/pythonwasm2go/p2.Fn4574
func Fn4574(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4579 github.com/goccy/pythonwasm2go/p0.Fn4579
func Fn4579(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4584 github.com/goccy/pythonwasm2go/p2.Fn4584
func Fn4584(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4585 github.com/goccy/pythonwasm2go/p2.Fn4585
func Fn4585(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4590 github.com/goccy/pythonwasm2go/p2.Fn4590
func Fn4590(m *base.Module, l0 int32, l1 int64) int32

//go:linkname Fn4591 github.com/goccy/pythonwasm2go/p2.Fn4591
func Fn4591(m *base.Module, l0 int32) int64

//go:linkname Fn4592 github.com/goccy/pythonwasm2go/p2.Fn4592
func Fn4592(m *base.Module)

//go:linkname Fn4594 github.com/goccy/pythonwasm2go/p2.Fn4594
func Fn4594(m *base.Module, l0 int32) int32

//go:linkname Fn4595 github.com/goccy/pythonwasm2go/p2.Fn4595
func Fn4595(m *base.Module, l0 float64, l1 int32) float64

//go:linkname Fn4596 github.com/goccy/pythonwasm2go/p2.Fn4596
func Fn4596(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4599 github.com/goccy/pythonwasm2go/p2.Fn4599
func Fn4599(m *base.Module, l0 int32) int64

//go:linkname Fn4600 github.com/goccy/pythonwasm2go/p2.Fn4600
func Fn4600(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4601 github.com/goccy/pythonwasm2go/p0.Fn4601
func Fn4601(m *base.Module)

//go:linkname Fn4603 github.com/goccy/pythonwasm2go/p2.Fn4603
func Fn4603(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4605 github.com/goccy/pythonwasm2go/p2.Fn4605
func Fn4605(m *base.Module, l0 int64) float64

//go:linkname Fn4606 github.com/goccy/pythonwasm2go/p2.Fn4606
func Fn4606(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn4608 github.com/goccy/pythonwasm2go/p2.Fn4608
func Fn4608(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn4609 github.com/goccy/pythonwasm2go/p2.Fn4609
func Fn4609(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn4612 github.com/goccy/pythonwasm2go/p2.Fn4612
func Fn4612(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn4615 github.com/goccy/pythonwasm2go/p2.Fn4615
func Fn4615(m *base.Module, l0 int32) int32

//go:linkname Fn4616 github.com/goccy/pythonwasm2go/p0.Fn4616
func Fn4616(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4618 github.com/goccy/pythonwasm2go/p2.Fn4618
func Fn4618(m *base.Module, l0 int32) int32

//go:linkname Fn4620 github.com/goccy/pythonwasm2go/p0.Fn4620
func Fn4620(m *base.Module, l0 int32) int32

//go:linkname Fn4621 github.com/goccy/pythonwasm2go/p2.Fn4621
func Fn4621(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4622 github.com/goccy/pythonwasm2go/p2.Fn4622
func Fn4622(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn4624 github.com/goccy/pythonwasm2go/p2.Fn4624
func Fn4624(m *base.Module, l0 int64) int64

//go:linkname Fn4625 github.com/goccy/pythonwasm2go/p0.Fn4625
func Fn4625(m *base.Module, l0 int64) int64

//go:linkname Fn4626 github.com/goccy/pythonwasm2go/p2.Fn4626
func Fn4626(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4629 github.com/goccy/pythonwasm2go/p2.Fn4629
func Fn4629(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4634 github.com/goccy/pythonwasm2go/p2.Fn4634
func Fn4634(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4660 github.com/goccy/pythonwasm2go/p2.Fn4660
func Fn4660(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4661 github.com/goccy/pythonwasm2go/p0.Fn4661
func Fn4661(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4662 github.com/goccy/pythonwasm2go/p0.Fn4662
func Fn4662(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4663 github.com/goccy/pythonwasm2go/p2.Fn4663
func Fn4663(m *base.Module, l0 int32) int32

//go:linkname Fn4664 github.com/goccy/pythonwasm2go/p0.Fn4664
func Fn4664(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn4665 github.com/goccy/pythonwasm2go/p2.Fn4665
func Fn4665(m *base.Module, l0 int32)

//go:linkname Fn4668 github.com/goccy/pythonwasm2go/p2.Fn4668
func Fn4668(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4689 github.com/goccy/pythonwasm2go/p2.Fn4689
func Fn4689(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4690 github.com/goccy/pythonwasm2go/p2.Fn4690
func Fn4690(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4692 github.com/goccy/pythonwasm2go/p2.Fn4692
func Fn4692(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4693 github.com/goccy/pythonwasm2go/p2.Fn4693
func Fn4693(m *base.Module, l0 int32) int32

//go:linkname Fn4694 github.com/goccy/pythonwasm2go/p2.Fn4694
func Fn4694(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4698 github.com/goccy/pythonwasm2go/p2.Fn4698
func Fn4698(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4699 github.com/goccy/pythonwasm2go/p2.Fn4699
func Fn4699(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4700 github.com/goccy/pythonwasm2go/p0.Fn4700
func Fn4700(m *base.Module, l0 int32) int32

//go:linkname Fn4701 github.com/goccy/pythonwasm2go/p2.Fn4701
func Fn4701(m *base.Module, l0 int32) int32

//go:linkname Fn4702 github.com/goccy/pythonwasm2go/p0.Fn4702
func Fn4702(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4706 github.com/goccy/pythonwasm2go/p2.Fn4706
func Fn4706(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4708 github.com/goccy/pythonwasm2go/p0.Fn4708
func Fn4708(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4709 github.com/goccy/pythonwasm2go/p0.Fn4709
func Fn4709(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4712 github.com/goccy/pythonwasm2go/p0.Fn4712
func Fn4712(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4713 github.com/goccy/pythonwasm2go/p2.Fn4713
func Fn4713(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4716 github.com/goccy/pythonwasm2go/p2.Fn4716
func Fn4716(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4717 github.com/goccy/pythonwasm2go/p2.Fn4717
func Fn4717(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4723 github.com/goccy/pythonwasm2go/p2.Fn4723
func Fn4723(m *base.Module, l0 int32) int32

//go:linkname Fn4732 github.com/goccy/pythonwasm2go/p0.Fn4732
func Fn4732(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4784 github.com/goccy/pythonwasm2go/p2.Fn4784
func Fn4784(m *base.Module)

//go:linkname Fn4785 github.com/goccy/pythonwasm2go/p2.Fn4785
func Fn4785(m *base.Module) int32

//go:linkname Fn4787 github.com/goccy/pythonwasm2go/p2.Fn4787
func Fn4787(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4790 github.com/goccy/pythonwasm2go/p2.Fn4790
func Fn4790(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4791 github.com/goccy/pythonwasm2go/p2.Fn4791
func Fn4791(m *base.Module) int64

//go:linkname Fn4792 github.com/goccy/pythonwasm2go/p2.Fn4792
func Fn4792(m *base.Module) int32

//go:linkname Fn4793 github.com/goccy/pythonwasm2go/p2.Fn4793
func Fn4793(m *base.Module, l0 int32)

//go:linkname Fn4795 github.com/goccy/pythonwasm2go/p2.Fn4795
func Fn4795(m *base.Module, l0 int32)

//go:linkname Fn4796 github.com/goccy/pythonwasm2go/p0.Fn4796
func Fn4796(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4797 github.com/goccy/pythonwasm2go/p2.Fn4797
func Fn4797(m *base.Module, l0 int32) int32

//go:linkname Fn4798 github.com/goccy/pythonwasm2go/p2.Fn4798
func Fn4798(m *base.Module, l0 int32)

//go:linkname Fn4799 github.com/goccy/pythonwasm2go/p2.Fn4799
func Fn4799(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4800 github.com/goccy/pythonwasm2go/p2.Fn4800
func Fn4800(m *base.Module, l0 int32) int32

//go:linkname Fn4802 github.com/goccy/pythonwasm2go/p0.Fn4802
func Fn4802(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4804 github.com/goccy/pythonwasm2go/p2.Fn4804
func Fn4804(m *base.Module, l0 int32)

//go:linkname Fn4819 github.com/goccy/pythonwasm2go/p2.Fn4819
func Fn4819(m *base.Module, l0 int32)

//go:linkname Fn4826 github.com/goccy/pythonwasm2go/p2.Fn4826
func Fn4826(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4830 github.com/goccy/pythonwasm2go/p2.Fn4830
func Fn4830(m *base.Module) int32

//go:linkname Fn4831 github.com/goccy/pythonwasm2go/p2.Fn4831
func Fn4831(m *base.Module, l0 int32) int32

//go:linkname Fn4835 github.com/goccy/pythonwasm2go/p2.Fn4835
func Fn4835(m *base.Module) int32

//go:linkname Fn4851 github.com/goccy/pythonwasm2go/p2.Fn4851
func Fn4851(m *base.Module, l0 int32)

//go:linkname Fn4852 github.com/goccy/pythonwasm2go/p0.Fn4852
func Fn4852(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4860 github.com/goccy/pythonwasm2go/p2.Fn4860
func Fn4860(m *base.Module, l0 int32) int32

//go:linkname Fn4866 github.com/goccy/pythonwasm2go/p2.Fn4866
func Fn4866(m *base.Module)

//go:linkname Fn4868 github.com/goccy/pythonwasm2go/p2.Fn4868
func Fn4868(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4869 github.com/goccy/pythonwasm2go/p2.Fn4869
func Fn4869(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4870 github.com/goccy/pythonwasm2go/p2.Fn4870
func Fn4870(m *base.Module, l0 int32, l1 int32) float64

//go:linkname Fn4872 github.com/goccy/pythonwasm2go/p2.Fn4872
func Fn4872(m *base.Module, l0 int32, l1 int32, l2 int32) float64

//go:linkname Fn4873 github.com/goccy/pythonwasm2go/p2.Fn4873
func Fn4873(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn4877 github.com/goccy/pythonwasm2go/p2.Fn4877
func Fn4877(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4878 github.com/goccy/pythonwasm2go/p0.Fn4878
func Fn4878(m *base.Module, l0 int32, l1 int32) float64

//go:linkname Fn4879 github.com/goccy/pythonwasm2go/p2.Fn4879
func Fn4879(m *base.Module, l0 int32) int32

//go:linkname Fn4885 github.com/goccy/pythonwasm2go/p2.Fn4885
func Fn4885(m *base.Module, l0 int32)

//go:linkname Fn4888 github.com/goccy/pythonwasm2go/p2.Fn4888
func Fn4888(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4892 github.com/goccy/pythonwasm2go/p2.Fn4892
func Fn4892(m *base.Module, l0 int32)

//go:linkname Fn4893 github.com/goccy/pythonwasm2go/p0.Fn4893
func Fn4893(m *base.Module, l0 float64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn4899 github.com/goccy/pythonwasm2go/p2.Fn4899
func Fn4899(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4901 github.com/goccy/pythonwasm2go/p2.Fn4901
func Fn4901(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn4902 github.com/goccy/pythonwasm2go/p2.Fn4902
func Fn4902(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn4903 github.com/goccy/pythonwasm2go/p2.Fn4903
func Fn4903(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4904 github.com/goccy/pythonwasm2go/p2.Fn4904
func Fn4904(m *base.Module, l0 int32) int32

//go:linkname Fn4905 github.com/goccy/pythonwasm2go/p2.Fn4905
func Fn4905(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4906 github.com/goccy/pythonwasm2go/p2.Fn4906
func Fn4906(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4908 github.com/goccy/pythonwasm2go/p2.Fn4908
func Fn4908(m *base.Module, l0 int32) int32

//go:linkname Fn4912 github.com/goccy/pythonwasm2go/p2.Fn4912
func Fn4912(m *base.Module, l0 int32)

//go:linkname Fn4914 github.com/goccy/pythonwasm2go/p2.Fn4914
func Fn4914(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn4917 github.com/goccy/pythonwasm2go/p2.Fn4917
func Fn4917(m *base.Module) int32

//go:linkname Fn4919 github.com/goccy/pythonwasm2go/p2.Fn4919
func Fn4919(m *base.Module) int32

//go:linkname Fn4920 github.com/goccy/pythonwasm2go/p2.Fn4920
func Fn4920(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4923 github.com/goccy/pythonwasm2go/p2.Fn4923
func Fn4923(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4925 github.com/goccy/pythonwasm2go/p2.Fn4925
func Fn4925(m *base.Module, l0 int32) int32

//go:linkname Fn4927 github.com/goccy/pythonwasm2go/p2.Fn4927
func Fn4927(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4928 github.com/goccy/pythonwasm2go/p2.Fn4928
func Fn4928(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4931 github.com/goccy/pythonwasm2go/p2.Fn4931
func Fn4931(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4932 github.com/goccy/pythonwasm2go/p2.Fn4932
func Fn4932(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4934 github.com/goccy/pythonwasm2go/p2.Fn4934
func Fn4934(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4935 github.com/goccy/pythonwasm2go/p2.Fn4935
func Fn4935(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4938 github.com/goccy/pythonwasm2go/p0.Fn4938
func Fn4938(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4939 github.com/goccy/pythonwasm2go/p2.Fn4939
func Fn4939(m *base.Module, l0 int32) int32

//go:linkname Fn4942 github.com/goccy/pythonwasm2go/p2.Fn4942
func Fn4942(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4944 github.com/goccy/pythonwasm2go/p2.Fn4944
func Fn4944(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4946 github.com/goccy/pythonwasm2go/p2.Fn4946
func Fn4946(m *base.Module, l0 int32) int32

//go:linkname Fn4947 github.com/goccy/pythonwasm2go/p2.Fn4947
func Fn4947(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4948 github.com/goccy/pythonwasm2go/p2.Fn4948
func Fn4948(m *base.Module, l0 int32) int32

//go:linkname Fn4951 github.com/goccy/pythonwasm2go/p2.Fn4951
func Fn4951(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4990 github.com/goccy/pythonwasm2go/p2.Fn4990
func Fn4990(m *base.Module, l0 int32) int32

//go:linkname Fn4991 github.com/goccy/pythonwasm2go/p2.Fn4991
func Fn4991(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5025 github.com/goccy/pythonwasm2go/p2.Fn5025
func Fn5025(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5026 github.com/goccy/pythonwasm2go/p2.Fn5026
func Fn5026(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5028 github.com/goccy/pythonwasm2go/p2.Fn5028
func Fn5028(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5029 github.com/goccy/pythonwasm2go/p2.Fn5029
func Fn5029(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5031 github.com/goccy/pythonwasm2go/p2.Fn5031
func Fn5031(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5036 github.com/goccy/pythonwasm2go/p2.Fn5036
func Fn5036(m *base.Module, l0 int32) int32

//go:linkname Fn5041 github.com/goccy/pythonwasm2go/p2.Fn5041
func Fn5041(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5048 github.com/goccy/pythonwasm2go/p2.Fn5048
func Fn5048(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5066 github.com/goccy/pythonwasm2go/p2.Fn5066
func Fn5066(m *base.Module, l0 int32) int32

//go:linkname Fn5085 github.com/goccy/pythonwasm2go/p2.Fn5085
func Fn5085(m *base.Module) int32

//go:linkname Fn5087 github.com/goccy/pythonwasm2go/p2.Fn5087
func Fn5087(m *base.Module, l0 int32) int32

//go:linkname Fn5092 github.com/goccy/pythonwasm2go/p2.Fn5092
func Fn5092(m *base.Module, l0 int32)

//go:linkname Fn5095 github.com/goccy/pythonwasm2go/p2.Fn5095
func Fn5095(m *base.Module, l0 int32)

//go:linkname Fn5102 github.com/goccy/pythonwasm2go/p2.Fn5102
func Fn5102(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5104 github.com/goccy/pythonwasm2go/p2.Fn5104
func Fn5104(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5105 github.com/goccy/pythonwasm2go/p2.Fn5105
func Fn5105(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5109 github.com/goccy/pythonwasm2go/p2.Fn5109
func Fn5109(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5113 github.com/goccy/pythonwasm2go/p2.Fn5113
func Fn5113(m *base.Module, l0 int32) int32

//go:linkname Fn5114 github.com/goccy/pythonwasm2go/p2.Fn5114
func Fn5114(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5117 github.com/goccy/pythonwasm2go/p0.Fn5117
func Fn5117(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5123 github.com/goccy/pythonwasm2go/p2.Fn5123
func Fn5123(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5126 github.com/goccy/pythonwasm2go/p2.Fn5126
func Fn5126(m *base.Module, l0 int32) int32

//go:linkname Fn5127 github.com/goccy/pythonwasm2go/p2.Fn5127
func Fn5127(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5129 github.com/goccy/pythonwasm2go/p2.Fn5129
func Fn5129(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5173 github.com/goccy/pythonwasm2go/p2.Fn5173
func Fn5173(m *base.Module, l0 int32)

//go:linkname Fn5198 github.com/goccy/pythonwasm2go/p2.Fn5198
func Fn5198(m *base.Module, l0 int32) int32

//go:linkname Fn5202 github.com/goccy/pythonwasm2go/p2.Fn5202
func Fn5202(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5219 github.com/goccy/pythonwasm2go/p2.Fn5219
func Fn5219(m *base.Module, l0 int32) int32

//go:linkname Fn5224 github.com/goccy/pythonwasm2go/p2.Fn5224
func Fn5224(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5229 github.com/goccy/pythonwasm2go/p2.Fn5229
func Fn5229(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5230 github.com/goccy/pythonwasm2go/p2.Fn5230
func Fn5230(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5237 github.com/goccy/pythonwasm2go/p2.Fn5237
func Fn5237(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5267 github.com/goccy/pythonwasm2go/p2.Fn5267
func Fn5267(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5279 github.com/goccy/pythonwasm2go/p0.Fn5279
func Fn5279(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5295 github.com/goccy/pythonwasm2go/p2.Fn5295
func Fn5295(m *base.Module, l0 int32)

//go:linkname Fn5300 github.com/goccy/pythonwasm2go/p2.Fn5300
func Fn5300(m *base.Module, l0 int32) int32

//go:linkname Fn5317 github.com/goccy/pythonwasm2go/p2.Fn5317
func Fn5317(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5318 github.com/goccy/pythonwasm2go/p2.Fn5318
func Fn5318(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5320 github.com/goccy/pythonwasm2go/p2.Fn5320
func Fn5320(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5327 github.com/goccy/pythonwasm2go/p2.Fn5327
func Fn5327(m *base.Module, l0 int32) int32

//go:linkname Fn5328 github.com/goccy/pythonwasm2go/p2.Fn5328
func Fn5328(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5329 github.com/goccy/pythonwasm2go/p2.Fn5329
func Fn5329(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5330 github.com/goccy/pythonwasm2go/p2.Fn5330
func Fn5330(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5332 github.com/goccy/pythonwasm2go/p2.Fn5332
func Fn5332(m *base.Module, l0 int32) int32

//go:linkname Fn5334 github.com/goccy/pythonwasm2go/p2.Fn5334
func Fn5334(m *base.Module, l0 int32) int32

//go:linkname Fn5336 github.com/goccy/pythonwasm2go/p2.Fn5336
func Fn5336(m *base.Module, l0 int32) int32

//go:linkname Fn5338 github.com/goccy/pythonwasm2go/p2.Fn5338
func Fn5338(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5339 github.com/goccy/pythonwasm2go/p2.Fn5339
func Fn5339(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5340 github.com/goccy/pythonwasm2go/p0.Fn5340
func Fn5340(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5342 github.com/goccy/pythonwasm2go/p2.Fn5342
func Fn5342(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5343 github.com/goccy/pythonwasm2go/p2.Fn5343
func Fn5343(m *base.Module) int32

//go:linkname Fn5345 github.com/goccy/pythonwasm2go/p0.Fn5345
func Fn5345(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5346 github.com/goccy/pythonwasm2go/p2.Fn5346
func Fn5346(m *base.Module, l0 int32)

//go:linkname Fn5347 github.com/goccy/pythonwasm2go/p2.Fn5347
func Fn5347(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5364 github.com/goccy/pythonwasm2go/p2.Fn5364
func Fn5364(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5365 github.com/goccy/pythonwasm2go/p2.Fn5365
func Fn5365(m *base.Module, l0 int32) int32

//go:linkname Fn5367 github.com/goccy/pythonwasm2go/p2.Fn5367
func Fn5367(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5369 github.com/goccy/pythonwasm2go/p2.Fn5369
func Fn5369(m *base.Module, l0 int32) int32

//go:linkname Fn5370 github.com/goccy/pythonwasm2go/p2.Fn5370
func Fn5370(m *base.Module, l0 int32) int32

//go:linkname Fn5371 github.com/goccy/pythonwasm2go/p2.Fn5371
func Fn5371(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5372 github.com/goccy/pythonwasm2go/p2.Fn5372
func Fn5372(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5373 github.com/goccy/pythonwasm2go/p2.Fn5373
func Fn5373(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5380 github.com/goccy/pythonwasm2go/p2.Fn5380
func Fn5380(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5385 github.com/goccy/pythonwasm2go/p2.Fn5385
func Fn5385(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5387 github.com/goccy/pythonwasm2go/p2.Fn5387
func Fn5387(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5389 github.com/goccy/pythonwasm2go/p2.Fn5389
func Fn5389(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5392 github.com/goccy/pythonwasm2go/p2.Fn5392
func Fn5392(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5396 github.com/goccy/pythonwasm2go/p2.Fn5396
func Fn5396(m *base.Module, l0 int32)

//go:linkname Fn5397 github.com/goccy/pythonwasm2go/p2.Fn5397
func Fn5397(m *base.Module, l0 int32) int32

//go:linkname Fn5399 github.com/goccy/pythonwasm2go/p2.Fn5399
func Fn5399(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5400 github.com/goccy/pythonwasm2go/p2.Fn5400
func Fn5400(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5411 github.com/goccy/pythonwasm2go/p2.Fn5411
func Fn5411(m *base.Module, l0 int32)

//go:linkname Fn5419 github.com/goccy/pythonwasm2go/p2.Fn5419
func Fn5419(m *base.Module, l0 int32)

//go:linkname Fn5426 github.com/goccy/pythonwasm2go/p2.Fn5426
func Fn5426(m *base.Module, l0 int32) int32

//go:linkname Fn5457 github.com/goccy/pythonwasm2go/p2.Fn5457
func Fn5457(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5479 github.com/goccy/pythonwasm2go/p2.Fn5479
func Fn5479(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5498 github.com/goccy/pythonwasm2go/p2.Fn5498
func Fn5498(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5499 github.com/goccy/pythonwasm2go/p2.Fn5499
func Fn5499(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5589 github.com/goccy/pythonwasm2go/p2.Fn5589
func Fn5589(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5591 github.com/goccy/pythonwasm2go/p2.Fn5591
func Fn5591(m *base.Module, l0 int32)

//go:linkname Fn5592 github.com/goccy/pythonwasm2go/p2.Fn5592
func Fn5592(m *base.Module, l0 int32)

//go:linkname Fn5596 github.com/goccy/pythonwasm2go/p2.Fn5596
func Fn5596(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5599 github.com/goccy/pythonwasm2go/p2.Fn5599
func Fn5599(m *base.Module) int32

//go:linkname Fn5605 github.com/goccy/pythonwasm2go/p2.Fn5605
func Fn5605(m *base.Module, l0 int32)

//go:linkname Fn5607 github.com/goccy/pythonwasm2go/p2.Fn5607
func Fn5607(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5608 github.com/goccy/pythonwasm2go/p2.Fn5608
func Fn5608(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5609 github.com/goccy/pythonwasm2go/p2.Fn5609
func Fn5609(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5613 github.com/goccy/pythonwasm2go/p0.Fn5613
func Fn5613(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5620 github.com/goccy/pythonwasm2go/p2.Fn5620
func Fn5620(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn5631 github.com/goccy/pythonwasm2go/p2.Fn5631
func Fn5631(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5658 github.com/goccy/pythonwasm2go/p2.Fn5658
func Fn5658(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5662 github.com/goccy/pythonwasm2go/p2.Fn5662
func Fn5662(m *base.Module, l0 float64, l1 int32) int32

//go:linkname Fn5674 github.com/goccy/pythonwasm2go/p2.Fn5674
func Fn5674(m *base.Module, l0 int64) int32

//go:linkname Fn5683 github.com/goccy/pythonwasm2go/p2.Fn5683
func Fn5683(m *base.Module, l0 float64) float64

//go:linkname Fn5699 github.com/goccy/pythonwasm2go/p2.Fn5699
func Fn5699(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5708 github.com/goccy/pythonwasm2go/p2.Fn5708
func Fn5708(m *base.Module, l0 float64) float64

//go:linkname Fn5709 github.com/goccy/pythonwasm2go/p2.Fn5709
func Fn5709(m *base.Module, l0 float64) float64

//go:linkname Fn5713 github.com/goccy/pythonwasm2go/p2.Fn5713
func Fn5713(m *base.Module, l0 float64) int32

//go:linkname Fn5732 github.com/goccy/pythonwasm2go/p2.Fn5732
func Fn5732(m *base.Module)

//go:linkname Fn5747 github.com/goccy/pythonwasm2go/p2.Fn5747
func Fn5747(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5748 github.com/goccy/pythonwasm2go/p2.Fn5748
func Fn5748(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5749 github.com/goccy/pythonwasm2go/p2.Fn5749
func Fn5749(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5750 github.com/goccy/pythonwasm2go/p2.Fn5750
func Fn5750(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5751 github.com/goccy/pythonwasm2go/p2.Fn5751
func Fn5751(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5752 github.com/goccy/pythonwasm2go/p2.Fn5752
func Fn5752(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5753 github.com/goccy/pythonwasm2go/p2.Fn5753
func Fn5753(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5755 github.com/goccy/pythonwasm2go/p2.Fn5755
func Fn5755(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5761 github.com/goccy/pythonwasm2go/p2.Fn5761
func Fn5761(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5765 github.com/goccy/pythonwasm2go/p2.Fn5765
func Fn5765(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5766 github.com/goccy/pythonwasm2go/p2.Fn5766
func Fn5766(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5767 github.com/goccy/pythonwasm2go/p2.Fn5767
func Fn5767(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5768 github.com/goccy/pythonwasm2go/p2.Fn5768
func Fn5768(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5769 github.com/goccy/pythonwasm2go/p2.Fn5769
func Fn5769(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5772 github.com/goccy/pythonwasm2go/p2.Fn5772
func Fn5772(m *base.Module, l0 int32) int32

//go:linkname Fn5777 github.com/goccy/pythonwasm2go/p2.Fn5777
func Fn5777(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5781 github.com/goccy/pythonwasm2go/p2.Fn5781
func Fn5781(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5782 github.com/goccy/pythonwasm2go/p2.Fn5782
func Fn5782(m *base.Module, l0 int32) int32

//go:linkname Fn5783 github.com/goccy/pythonwasm2go/p2.Fn5783
func Fn5783(m *base.Module, l0 int32) int32

//go:linkname Fn5784 github.com/goccy/pythonwasm2go/p2.Fn5784
func Fn5784(m *base.Module, l0 int32) int32

//go:linkname Fn5786 github.com/goccy/pythonwasm2go/p2.Fn5786
func Fn5786(m *base.Module, l0 int32) int32

//go:linkname Fn5787 github.com/goccy/pythonwasm2go/p2.Fn5787
func Fn5787(m *base.Module, l0 int32) int32

//go:linkname Fn5788 github.com/goccy/pythonwasm2go/p2.Fn5788
func Fn5788(m *base.Module, l0 int32) int32

//go:linkname Fn5789 github.com/goccy/pythonwasm2go/p2.Fn5789
func Fn5789(m *base.Module, l0 int32) int32

//go:linkname Fn5790 github.com/goccy/pythonwasm2go/p2.Fn5790
func Fn5790(m *base.Module, l0 int32) int32

//go:linkname Fn5791 github.com/goccy/pythonwasm2go/p2.Fn5791
func Fn5791(m *base.Module, l0 int32) int32

//go:linkname Fn5792 github.com/goccy/pythonwasm2go/p2.Fn5792
func Fn5792(m *base.Module, l0 int32) int32

//go:linkname Fn5795 github.com/goccy/pythonwasm2go/p2.Fn5795
func Fn5795(m *base.Module, l0 int32)

//go:linkname Fn5796 github.com/goccy/pythonwasm2go/p2.Fn5796
func Fn5796(m *base.Module, l0 int32)

//go:linkname Fn5798 github.com/goccy/pythonwasm2go/p2.Fn5798
func Fn5798(m *base.Module, l0 int32)

//go:linkname Fn5799 github.com/goccy/pythonwasm2go/p2.Fn5799
func Fn5799(m *base.Module, l0 int32)

//go:linkname Fn5800 github.com/goccy/pythonwasm2go/p2.Fn5800
func Fn5800(m *base.Module, l0 int32)

//go:linkname Fn5801 github.com/goccy/pythonwasm2go/p2.Fn5801
func Fn5801(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5802 github.com/goccy/pythonwasm2go/p2.Fn5802
func Fn5802(m *base.Module, l0 int32) int32

//go:linkname Fn5803 github.com/goccy/pythonwasm2go/p2.Fn5803
func Fn5803(m *base.Module, l0 int32) int32

//go:linkname Fn5804 github.com/goccy/pythonwasm2go/p2.Fn5804
func Fn5804(m *base.Module, l0 int32) int32

//go:linkname Fn5805 github.com/goccy/pythonwasm2go/p2.Fn5805
func Fn5805(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5806 github.com/goccy/pythonwasm2go/p2.Fn5806
func Fn5806(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5807 github.com/goccy/pythonwasm2go/p2.Fn5807
func Fn5807(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5808 github.com/goccy/pythonwasm2go/p2.Fn5808
func Fn5808(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5809 github.com/goccy/pythonwasm2go/p2.Fn5809
func Fn5809(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5812 github.com/goccy/pythonwasm2go/p2.Fn5812
func Fn5812(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5813 github.com/goccy/pythonwasm2go/p2.Fn5813
func Fn5813(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5814 github.com/goccy/pythonwasm2go/p2.Fn5814
func Fn5814(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5815 github.com/goccy/pythonwasm2go/p2.Fn5815
func Fn5815(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5816 github.com/goccy/pythonwasm2go/p2.Fn5816
func Fn5816(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5817 github.com/goccy/pythonwasm2go/p2.Fn5817
func Fn5817(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5819 github.com/goccy/pythonwasm2go/p2.Fn5819
func Fn5819(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5820 github.com/goccy/pythonwasm2go/p2.Fn5820
func Fn5820(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5821 github.com/goccy/pythonwasm2go/p2.Fn5821
func Fn5821(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5822 github.com/goccy/pythonwasm2go/p2.Fn5822
func Fn5822(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5823 github.com/goccy/pythonwasm2go/p2.Fn5823
func Fn5823(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5824 github.com/goccy/pythonwasm2go/p2.Fn5824
func Fn5824(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5825 github.com/goccy/pythonwasm2go/p2.Fn5825
func Fn5825(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5827 github.com/goccy/pythonwasm2go/p2.Fn5827
func Fn5827(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5828 github.com/goccy/pythonwasm2go/p2.Fn5828
func Fn5828(m *base.Module, l0 int32) int32

//go:linkname Fn5831 github.com/goccy/pythonwasm2go/p2.Fn5831
func Fn5831(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5832 github.com/goccy/pythonwasm2go/p2.Fn5832
func Fn5832(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5833 github.com/goccy/pythonwasm2go/p2.Fn5833
func Fn5833(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5835 github.com/goccy/pythonwasm2go/p2.Fn5835
func Fn5835(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5836 github.com/goccy/pythonwasm2go/p2.Fn5836
func Fn5836(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5838 github.com/goccy/pythonwasm2go/p2.Fn5838
func Fn5838(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5839 github.com/goccy/pythonwasm2go/p2.Fn5839
func Fn5839(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5845 github.com/goccy/pythonwasm2go/p2.Fn5845
func Fn5845(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5848 github.com/goccy/pythonwasm2go/p2.Fn5848
func Fn5848(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5850 github.com/goccy/pythonwasm2go/p2.Fn5850
func Fn5850(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5851 github.com/goccy/pythonwasm2go/p2.Fn5851
func Fn5851(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5856 github.com/goccy/pythonwasm2go/p2.Fn5856
func Fn5856(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5858 github.com/goccy/pythonwasm2go/p2.Fn5858
func Fn5858(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5867 github.com/goccy/pythonwasm2go/p2.Fn5867
func Fn5867(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5868 github.com/goccy/pythonwasm2go/p2.Fn5868
func Fn5868(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5869 github.com/goccy/pythonwasm2go/p2.Fn5869
func Fn5869(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn5870 github.com/goccy/pythonwasm2go/p2.Fn5870
func Fn5870(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn5872 github.com/goccy/pythonwasm2go/p2.Fn5872
func Fn5872(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5873 github.com/goccy/pythonwasm2go/p2.Fn5873
func Fn5873(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5878 github.com/goccy/pythonwasm2go/p2.Fn5878
func Fn5878(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5879 github.com/goccy/pythonwasm2go/p2.Fn5879
func Fn5879(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5882 github.com/goccy/pythonwasm2go/p2.Fn5882
func Fn5882(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5883 github.com/goccy/pythonwasm2go/p2.Fn5883
func Fn5883(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5885 github.com/goccy/pythonwasm2go/p2.Fn5885
func Fn5885(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5886 github.com/goccy/pythonwasm2go/p2.Fn5886
func Fn5886(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5888 github.com/goccy/pythonwasm2go/p2.Fn5888
func Fn5888(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5889 github.com/goccy/pythonwasm2go/p2.Fn5889
func Fn5889(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5890 github.com/goccy/pythonwasm2go/p2.Fn5890
func Fn5890(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5893 github.com/goccy/pythonwasm2go/p2.Fn5893
func Fn5893(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5894 github.com/goccy/pythonwasm2go/p0.Fn5894
func Fn5894(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5896 github.com/goccy/pythonwasm2go/p2.Fn5896
func Fn5896(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5898 github.com/goccy/pythonwasm2go/p2.Fn5898
func Fn5898(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5900 github.com/goccy/pythonwasm2go/p2.Fn5900
func Fn5900(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5902 github.com/goccy/pythonwasm2go/p2.Fn5902
func Fn5902(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5905 github.com/goccy/pythonwasm2go/p2.Fn5905
func Fn5905(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5906 github.com/goccy/pythonwasm2go/p2.Fn5906
func Fn5906(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5909 github.com/goccy/pythonwasm2go/p2.Fn5909
func Fn5909(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5914 github.com/goccy/pythonwasm2go/p2.Fn5914
func Fn5914(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5915 github.com/goccy/pythonwasm2go/p2.Fn5915
func Fn5915(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5917 github.com/goccy/pythonwasm2go/p2.Fn5917
func Fn5917(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5918 github.com/goccy/pythonwasm2go/p2.Fn5918
func Fn5918(m *base.Module, l0 int32) int32

//go:linkname Fn5919 github.com/goccy/pythonwasm2go/p2.Fn5919
func Fn5919(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5921 github.com/goccy/pythonwasm2go/p2.Fn5921
func Fn5921(m *base.Module) int32

//go:linkname Fn5922 github.com/goccy/pythonwasm2go/p2.Fn5922
func Fn5922(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5923 github.com/goccy/pythonwasm2go/p2.Fn5923
func Fn5923(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5925 github.com/goccy/pythonwasm2go/p2.Fn5925
func Fn5925(m *base.Module, l0 int32)

//go:linkname Fn5926 github.com/goccy/pythonwasm2go/p2.Fn5926
func Fn5926(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5927 github.com/goccy/pythonwasm2go/p2.Fn5927
func Fn5927(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5928 github.com/goccy/pythonwasm2go/p2.Fn5928
func Fn5928(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5930 github.com/goccy/pythonwasm2go/p2.Fn5930
func Fn5930(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5931 github.com/goccy/pythonwasm2go/p2.Fn5931
func Fn5931(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5932 github.com/goccy/pythonwasm2go/p2.Fn5932
func Fn5932(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5934 github.com/goccy/pythonwasm2go/p2.Fn5934
func Fn5934(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5942 github.com/goccy/pythonwasm2go/p2.Fn5942
func Fn5942(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5943 github.com/goccy/pythonwasm2go/p2.Fn5943
func Fn5943(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5950 github.com/goccy/pythonwasm2go/p2.Fn5950
func Fn5950(m *base.Module, l0 int32) int32

//go:linkname Fn5952 github.com/goccy/pythonwasm2go/p2.Fn5952
func Fn5952(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5956 github.com/goccy/pythonwasm2go/p2.Fn5956
func Fn5956(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5957 github.com/goccy/pythonwasm2go/p2.Fn5957
func Fn5957(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5958 github.com/goccy/pythonwasm2go/p2.Fn5958
func Fn5958(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5959 github.com/goccy/pythonwasm2go/p2.Fn5959
func Fn5959(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5960 github.com/goccy/pythonwasm2go/p2.Fn5960
func Fn5960(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5961 github.com/goccy/pythonwasm2go/p2.Fn5961
func Fn5961(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5962 github.com/goccy/pythonwasm2go/p2.Fn5962
func Fn5962(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5963 github.com/goccy/pythonwasm2go/p2.Fn5963
func Fn5963(m *base.Module, l0 int32)

//go:linkname Fn5964 github.com/goccy/pythonwasm2go/p2.Fn5964
func Fn5964(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5965 github.com/goccy/pythonwasm2go/p2.Fn5965
func Fn5965(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5966 github.com/goccy/pythonwasm2go/p2.Fn5966
func Fn5966(m *base.Module, l0 int32)

//go:linkname Fn5967 github.com/goccy/pythonwasm2go/p2.Fn5967
func Fn5967(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5975 github.com/goccy/pythonwasm2go/p2.Fn5975
func Fn5975(m *base.Module, l0 int32) int32

//go:linkname Fn5980 github.com/goccy/pythonwasm2go/p2.Fn5980
func Fn5980(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5981 github.com/goccy/pythonwasm2go/p2.Fn5981
func Fn5981(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5984 github.com/goccy/pythonwasm2go/p2.Fn5984
func Fn5984(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5986 github.com/goccy/pythonwasm2go/p2.Fn5986
func Fn5986(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5987 github.com/goccy/pythonwasm2go/p2.Fn5987
func Fn5987(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5988 github.com/goccy/pythonwasm2go/p2.Fn5988
func Fn5988(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5989 github.com/goccy/pythonwasm2go/p2.Fn5989
func Fn5989(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5992 github.com/goccy/pythonwasm2go/p2.Fn5992
func Fn5992(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6003 github.com/goccy/pythonwasm2go/p0.Fn6003
func Fn6003(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6007 github.com/goccy/pythonwasm2go/p0.Fn6007
func Fn6007(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn6125 github.com/goccy/pythonwasm2go/p2.Fn6125
func Fn6125(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6151 github.com/goccy/pythonwasm2go/p2.Fn6151
func Fn6151(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6153 github.com/goccy/pythonwasm2go/p2.Fn6153
func Fn6153(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6154 github.com/goccy/pythonwasm2go/p2.Fn6154
func Fn6154(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6178 github.com/goccy/pythonwasm2go/p2.Fn6178
func Fn6178(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6190 github.com/goccy/pythonwasm2go/p2.Fn6190
func Fn6190(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6192 github.com/goccy/pythonwasm2go/p2.Fn6192
func Fn6192(m *base.Module, l0 int32)

//go:linkname Fn6193 github.com/goccy/pythonwasm2go/p2.Fn6193
func Fn6193(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6194 github.com/goccy/pythonwasm2go/p0.Fn6194
func Fn6194(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6195 github.com/goccy/pythonwasm2go/p2.Fn6195
func Fn6195(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32)

//go:linkname Fn6196 github.com/goccy/pythonwasm2go/p2.Fn6196
func Fn6196(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6201 github.com/goccy/pythonwasm2go/p2.Fn6201
func Fn6201(m *base.Module, l0 int32) int32

//go:linkname Fn6211 github.com/goccy/pythonwasm2go/p2.Fn6211
func Fn6211(m *base.Module, l0 int32)

//go:linkname Fn6212 github.com/goccy/pythonwasm2go/p2.Fn6212
func Fn6212(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6214 github.com/goccy/pythonwasm2go/p2.Fn6214
func Fn6214(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32)

//go:linkname Fn6215 github.com/goccy/pythonwasm2go/p2.Fn6215
func Fn6215(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6228 github.com/goccy/pythonwasm2go/p2.Fn6228
func Fn6228(m *base.Module, l0 int32)

//go:linkname Fn6229 github.com/goccy/pythonwasm2go/p2.Fn6229
func Fn6229(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6231 github.com/goccy/pythonwasm2go/p2.Fn6231
func Fn6231(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6233 github.com/goccy/pythonwasm2go/p2.Fn6233
func Fn6233(m *base.Module, l0 int32)

//go:linkname Fn6234 github.com/goccy/pythonwasm2go/p2.Fn6234
func Fn6234(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6235 github.com/goccy/pythonwasm2go/p2.Fn6235
func Fn6235(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6236 github.com/goccy/pythonwasm2go/p2.Fn6236
func Fn6236(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6237 github.com/goccy/pythonwasm2go/p2.Fn6237
func Fn6237(m *base.Module, l0 int32)

//go:linkname Fn6238 github.com/goccy/pythonwasm2go/p2.Fn6238
func Fn6238(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6239 github.com/goccy/pythonwasm2go/p0.Fn6239
func Fn6239(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6242 github.com/goccy/pythonwasm2go/p2.Fn6242
func Fn6242(m *base.Module, l0 int32)

//go:linkname Fn6243 github.com/goccy/pythonwasm2go/p2.Fn6243
func Fn6243(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6244 github.com/goccy/pythonwasm2go/p2.Fn6244
func Fn6244(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6248 github.com/goccy/pythonwasm2go/p2.Fn6248
func Fn6248(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6249 github.com/goccy/pythonwasm2go/p2.Fn6249
func Fn6249(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6252 github.com/goccy/pythonwasm2go/p2.Fn6252
func Fn6252(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6253 github.com/goccy/pythonwasm2go/p2.Fn6253
func Fn6253(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6259 github.com/goccy/pythonwasm2go/p2.Fn6259
func Fn6259(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6260 github.com/goccy/pythonwasm2go/p2.Fn6260
func Fn6260(m *base.Module, l0 int32) int32

//go:linkname Fn6262 github.com/goccy/pythonwasm2go/p2.Fn6262
func Fn6262(m *base.Module, l0 int32) int32

//go:linkname Fn6264 github.com/goccy/pythonwasm2go/p2.Fn6264
func Fn6264(m *base.Module, l0 int32) int32

//go:linkname Fn6266 github.com/goccy/pythonwasm2go/p2.Fn6266
func Fn6266(m *base.Module, l0 int32) int32

//go:linkname Fn6279 github.com/goccy/pythonwasm2go/p2.Fn6279
func Fn6279(m *base.Module, l0 int32)

//go:linkname Fn6280 github.com/goccy/pythonwasm2go/p2.Fn6280
func Fn6280(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6281 github.com/goccy/pythonwasm2go/p2.Fn6281
func Fn6281(m *base.Module, l0 int32) int32

//go:linkname Fn6286 github.com/goccy/pythonwasm2go/p2.Fn6286
func Fn6286(m *base.Module, l0 int32) int32

//go:linkname Fn6302 github.com/goccy/pythonwasm2go/p2.Fn6302
func Fn6302(m *base.Module, l0 int32) int32

//go:linkname Fn6316 github.com/goccy/pythonwasm2go/p2.Fn6316
func Fn6316(m *base.Module, l0 int32, l1 int64)

//go:linkname Fn6318 github.com/goccy/pythonwasm2go/p2.Fn6318
func Fn6318(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn6320 github.com/goccy/pythonwasm2go/p2.Fn6320
func Fn6320(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn6321 github.com/goccy/pythonwasm2go/p2.Fn6321
func Fn6321(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6325 github.com/goccy/pythonwasm2go/p2.Fn6325
func Fn6325(m *base.Module, l0 int32)

//go:linkname Fn6326 github.com/goccy/pythonwasm2go/p2.Fn6326
func Fn6326(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32)

//go:linkname Fn6328 github.com/goccy/pythonwasm2go/p2.Fn6328
func Fn6328(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32, l6 int32)

//go:linkname Fn6329 github.com/goccy/pythonwasm2go/p2.Fn6329
func Fn6329(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6331 github.com/goccy/pythonwasm2go/p2.Fn6331
func Fn6331(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6338 github.com/goccy/pythonwasm2go/p2.Fn6338
func Fn6338(m *base.Module, l0 int32) int32

//go:linkname Fn6347 github.com/goccy/pythonwasm2go/p2.Fn6347
func Fn6347(m *base.Module, l0 int32) int32

//go:linkname Fn6348 github.com/goccy/pythonwasm2go/p2.Fn6348
func Fn6348(m *base.Module, l0 int32) int32

//go:linkname Fn6349 github.com/goccy/pythonwasm2go/p2.Fn6349
func Fn6349(m *base.Module, l0 int32) int32

//go:linkname Fn6351 github.com/goccy/pythonwasm2go/p2.Fn6351
func Fn6351(m *base.Module, l0 int32)

//go:linkname Fn6357 github.com/goccy/pythonwasm2go/p2.Fn6357
func Fn6357(m *base.Module, l0 int32)

//go:linkname Fn6375 github.com/goccy/pythonwasm2go/p2.Fn6375
func Fn6375(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6376 github.com/goccy/pythonwasm2go/p2.Fn6376
func Fn6376(m *base.Module, l0 int32) int32

//go:linkname Fn6405 github.com/goccy/pythonwasm2go/p2.Fn6405
func Fn6405(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6417 github.com/goccy/pythonwasm2go/p2.Fn6417
func Fn6417(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6419 github.com/goccy/pythonwasm2go/p2.Fn6419
func Fn6419(m *base.Module, l0 int32) int32

//go:linkname Fn6425 github.com/goccy/pythonwasm2go/p2.Fn6425
func Fn6425(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6433 github.com/goccy/pythonwasm2go/p2.Fn6433
func Fn6433(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32

//go:linkname Fn6437 github.com/goccy/pythonwasm2go/p2.Fn6437
func Fn6437(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6441 github.com/goccy/pythonwasm2go/p2.Fn6441
func Fn6441(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32

//go:linkname Fn6443 github.com/goccy/pythonwasm2go/p2.Fn6443
func Fn6443(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6464 github.com/goccy/pythonwasm2go/p2.Fn6464
func Fn6464(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6475 github.com/goccy/pythonwasm2go/p2.Fn6475
func Fn6475(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6496 github.com/goccy/pythonwasm2go/p2.Fn6496
func Fn6496(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6519 github.com/goccy/pythonwasm2go/p2.Fn6519
func Fn6519(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6570 github.com/goccy/pythonwasm2go/p2.Fn6570
func Fn6570(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6572 github.com/goccy/pythonwasm2go/p2.Fn6572
func Fn6572(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6573 github.com/goccy/pythonwasm2go/p2.Fn6573
func Fn6573(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn6574 github.com/goccy/pythonwasm2go/p2.Fn6574
func Fn6574(m *base.Module, l0 int32) int32

//go:linkname Fn6575 github.com/goccy/pythonwasm2go/p2.Fn6575
func Fn6575(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int64, l5 int32)

//go:linkname Fn6576 github.com/goccy/pythonwasm2go/p2.Fn6576
func Fn6576(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6577 github.com/goccy/pythonwasm2go/p2.Fn6577
func Fn6577(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6579 github.com/goccy/pythonwasm2go/p2.Fn6579
func Fn6579(m *base.Module, l0 int32)

//go:linkname Fn6580 github.com/goccy/pythonwasm2go/p2.Fn6580
func Fn6580(m *base.Module, l0 int32) int32

//go:linkname Fn6581 github.com/goccy/pythonwasm2go/p2.Fn6581
func Fn6581(m *base.Module, l0 int32)

//go:linkname Fn6583 github.com/goccy/pythonwasm2go/p2.Fn6583
func Fn6583(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6584 github.com/goccy/pythonwasm2go/p2.Fn6584
func Fn6584(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6585 github.com/goccy/pythonwasm2go/p2.Fn6585
func Fn6585(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6593 github.com/goccy/pythonwasm2go/p2.Fn6593
func Fn6593(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6594 github.com/goccy/pythonwasm2go/p2.Fn6594
func Fn6594(m *base.Module, l0 int32)

//go:linkname Fn6599 github.com/goccy/pythonwasm2go/p2.Fn6599
func Fn6599(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6620 github.com/goccy/pythonwasm2go/p2.Fn6620
func Fn6620(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6627 github.com/goccy/pythonwasm2go/p2.Fn6627
func Fn6627(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6628 github.com/goccy/pythonwasm2go/p2.Fn6628
func Fn6628(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6629 github.com/goccy/pythonwasm2go/p2.Fn6629
func Fn6629(m *base.Module, l0 int32) int32

//go:linkname Fn6636 github.com/goccy/pythonwasm2go/p2.Fn6636
func Fn6636(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6637 github.com/goccy/pythonwasm2go/p2.Fn6637
func Fn6637(m *base.Module, l0 int32) int32

//go:linkname Fn6643 github.com/goccy/pythonwasm2go/p2.Fn6643
func Fn6643(m *base.Module, l0 int32) int32

//go:linkname Fn6646 github.com/goccy/pythonwasm2go/p0.Fn6646
func Fn6646(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32

//go:linkname Fn6648 github.com/goccy/pythonwasm2go/p2.Fn6648
func Fn6648(m *base.Module, l0 int32)

//go:linkname Fn6650 github.com/goccy/pythonwasm2go/p2.Fn6650
func Fn6650(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6655 github.com/goccy/pythonwasm2go/p2.Fn6655
func Fn6655(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6659 github.com/goccy/pythonwasm2go/p2.Fn6659
func Fn6659(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6662 github.com/goccy/pythonwasm2go/p2.Fn6662
func Fn6662(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6664 github.com/goccy/pythonwasm2go/p2.Fn6664
func Fn6664(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6665 github.com/goccy/pythonwasm2go/p2.Fn6665
func Fn6665(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6666 github.com/goccy/pythonwasm2go/p2.Fn6666
func Fn6666(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6667 github.com/goccy/pythonwasm2go/p2.Fn6667
func Fn6667(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6668 github.com/goccy/pythonwasm2go/p2.Fn6668
func Fn6668(m *base.Module, l0 int32) float32

//go:linkname Fn6669 github.com/goccy/pythonwasm2go/p2.Fn6669
func Fn6669(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6670 github.com/goccy/pythonwasm2go/p2.Fn6670
func Fn6670(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6671 github.com/goccy/pythonwasm2go/p0.Fn6671
func Fn6671(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn6690 github.com/goccy/pythonwasm2go/p2.Fn6690
func Fn6690(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6694 github.com/goccy/pythonwasm2go/p2.Fn6694
func Fn6694(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6696 github.com/goccy/pythonwasm2go/p2.Fn6696
func Fn6696(m *base.Module, l0 int32) int32

//go:linkname Fn6698 github.com/goccy/pythonwasm2go/p2.Fn6698
func Fn6698(m *base.Module, l0 int32) int32

//go:linkname Fn6699 github.com/goccy/pythonwasm2go/p2.Fn6699
func Fn6699(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6700 github.com/goccy/pythonwasm2go/p2.Fn6700
func Fn6700(m *base.Module, l0 int32)

//go:linkname Fn6715 github.com/goccy/pythonwasm2go/p2.Fn6715
func Fn6715(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6731 github.com/goccy/pythonwasm2go/p2.Fn6731
func Fn6731(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6737 github.com/goccy/pythonwasm2go/p2.Fn6737
func Fn6737(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6774 github.com/goccy/pythonwasm2go/p2.Fn6774
func Fn6774(m *base.Module, l0 int32) int32

//go:linkname Fn6775 github.com/goccy/pythonwasm2go/p2.Fn6775
func Fn6775(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6776 github.com/goccy/pythonwasm2go/p2.Fn6776
func Fn6776(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6778 github.com/goccy/pythonwasm2go/p2.Fn6778
func Fn6778(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6779 github.com/goccy/pythonwasm2go/p2.Fn6779
func Fn6779(m *base.Module, l0 int32)

//go:linkname Fn6780 github.com/goccy/pythonwasm2go/p2.Fn6780
func Fn6780(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6785 github.com/goccy/pythonwasm2go/p2.Fn6785
func Fn6785(m *base.Module, l0 int32) int32

//go:linkname Fn6786 github.com/goccy/pythonwasm2go/p2.Fn6786
func Fn6786(m *base.Module, l0 int32) int32

//go:linkname Fn6787 github.com/goccy/pythonwasm2go/p2.Fn6787
func Fn6787(m *base.Module, l0 int32) int32

//go:linkname Fn6793 github.com/goccy/pythonwasm2go/p2.Fn6793
func Fn6793(m *base.Module, l0 int32) int32

//go:linkname Fn6803 github.com/goccy/pythonwasm2go/p2.Fn6803
func Fn6803(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6805 github.com/goccy/pythonwasm2go/p2.Fn6805
func Fn6805(m *base.Module, l0 int32) int32

//go:linkname Fn6806 github.com/goccy/pythonwasm2go/p2.Fn6806
func Fn6806(m *base.Module, l0 int32) int32

//go:linkname Fn6807 github.com/goccy/pythonwasm2go/p2.Fn6807
func Fn6807(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6808 github.com/goccy/pythonwasm2go/p2.Fn6808
func Fn6808(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6815 github.com/goccy/pythonwasm2go/p2.Fn6815
func Fn6815(m *base.Module, l0 int32)

//go:linkname Fn6817 github.com/goccy/pythonwasm2go/p2.Fn6817
func Fn6817(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6820 github.com/goccy/pythonwasm2go/p2.Fn6820
func Fn6820(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6821 github.com/goccy/pythonwasm2go/p2.Fn6821
func Fn6821(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6825 github.com/goccy/pythonwasm2go/p2.Fn6825
func Fn6825(m *base.Module, l0 int32)

//go:linkname Fn6863 github.com/goccy/pythonwasm2go/p2.Fn6863
func Fn6863(m *base.Module, l0 int32) int32

//go:linkname Fn6875 github.com/goccy/pythonwasm2go/p2.Fn6875
func Fn6875(m *base.Module, l0 int32) int32

//go:linkname Fn6876 github.com/goccy/pythonwasm2go/p2.Fn6876
func Fn6876(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn6939 github.com/goccy/pythonwasm2go/p2.Fn6939
func Fn6939(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6959 github.com/goccy/pythonwasm2go/p2.Fn6959
func Fn6959(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6985 github.com/goccy/pythonwasm2go/p2.Fn6985
func Fn6985(m *base.Module, l0 int32) int32

//go:linkname Fn6987 github.com/goccy/pythonwasm2go/p2.Fn6987
func Fn6987(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6991 github.com/goccy/pythonwasm2go/p2.Fn6991
func Fn6991(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6992 github.com/goccy/pythonwasm2go/p2.Fn6992
func Fn6992(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6997 github.com/goccy/pythonwasm2go/p2.Fn6997
func Fn6997(m *base.Module, l0 int32) int32

//go:linkname Fn7008 github.com/goccy/pythonwasm2go/p2.Fn7008
func Fn7008(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7009 github.com/goccy/pythonwasm2go/p2.Fn7009
func Fn7009(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7032 github.com/goccy/pythonwasm2go/p2.Fn7032
func Fn7032(m *base.Module, l0 int32) int32

//go:linkname Fn7043 github.com/goccy/pythonwasm2go/p2.Fn7043
func Fn7043(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7047 github.com/goccy/pythonwasm2go/p2.Fn7047
func Fn7047(m *base.Module, l0 int32) int32

//go:linkname Fn7048 github.com/goccy/pythonwasm2go/p2.Fn7048
func Fn7048(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7049 github.com/goccy/pythonwasm2go/p2.Fn7049
func Fn7049(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn7050 github.com/goccy/pythonwasm2go/p2.Fn7050
func Fn7050(m *base.Module, l0 int32) int32

//go:linkname Fn7051 github.com/goccy/pythonwasm2go/p2.Fn7051
func Fn7051(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7052 github.com/goccy/pythonwasm2go/p2.Fn7052
func Fn7052(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn7053 github.com/goccy/pythonwasm2go/p2.Fn7053
func Fn7053(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7066 github.com/goccy/pythonwasm2go/p2.Fn7066
func Fn7066(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7067 github.com/goccy/pythonwasm2go/p2.Fn7067
func Fn7067(m *base.Module, l0 int32)

//go:linkname Fn7083 github.com/goccy/pythonwasm2go/p2.Fn7083
func Fn7083(m *base.Module) int32

//go:linkname Fn7093 github.com/goccy/pythonwasm2go/p2.Fn7093
func Fn7093(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7094 github.com/goccy/pythonwasm2go/p2.Fn7094
func Fn7094(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7106 github.com/goccy/pythonwasm2go/p2.Fn7106
func Fn7106(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7134 github.com/goccy/pythonwasm2go/p2.Fn7134
func Fn7134(m *base.Module) int64

//go:linkname Fn7149 github.com/goccy/pythonwasm2go/p2.Fn7149
func Fn7149(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7166 github.com/goccy/pythonwasm2go/p2.Fn7166
func Fn7166(m *base.Module, l0 int32) int32

//go:linkname Fn7170 github.com/goccy/pythonwasm2go/p2.Fn7170
func Fn7170(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7176 github.com/goccy/pythonwasm2go/p2.Fn7176
func Fn7176(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7177 github.com/goccy/pythonwasm2go/p2.Fn7177
func Fn7177(m *base.Module, l0 int32)

//go:linkname Fn7178 github.com/goccy/pythonwasm2go/p0.Fn7178
func Fn7178(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn7179 github.com/goccy/pythonwasm2go/p2.Fn7179
func Fn7179(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7183 github.com/goccy/pythonwasm2go/p2.Fn7183
func Fn7183(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7185 github.com/goccy/pythonwasm2go/p2.Fn7185
func Fn7185(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7186 github.com/goccy/pythonwasm2go/p2.Fn7186
func Fn7186(m *base.Module, l0 int32)

//go:linkname Fn7191 github.com/goccy/pythonwasm2go/p2.Fn7191
func Fn7191(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7220 github.com/goccy/pythonwasm2go/p2.Fn7220
func Fn7220(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7221 github.com/goccy/pythonwasm2go/p2.Fn7221
func Fn7221(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7222 github.com/goccy/pythonwasm2go/p2.Fn7222
func Fn7222(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7244 github.com/goccy/pythonwasm2go/p2.Fn7244
func Fn7244(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7245 github.com/goccy/pythonwasm2go/p2.Fn7245
func Fn7245(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7256 github.com/goccy/pythonwasm2go/p2.Fn7256
func Fn7256(m *base.Module, l0 int32) int32

//go:linkname Fn7259 github.com/goccy/pythonwasm2go/p2.Fn7259
func Fn7259(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7270 github.com/goccy/pythonwasm2go/p2.Fn7270
func Fn7270(m *base.Module, l0 int32) int32

//go:linkname Fn7272 github.com/goccy/pythonwasm2go/p2.Fn7272
func Fn7272(m *base.Module, l0 int32) int32

//go:linkname Fn7273 github.com/goccy/pythonwasm2go/p2.Fn7273
func Fn7273(m *base.Module, l0 int32) int32

//go:linkname Fn7274 github.com/goccy/pythonwasm2go/p2.Fn7274
func Fn7274(m *base.Module, l0 int32) int32

//go:linkname Fn7290 github.com/goccy/pythonwasm2go/p2.Fn7290
func Fn7290(m *base.Module, l0 int32) int32

//go:linkname Fn7327 github.com/goccy/pythonwasm2go/p2.Fn7327
func Fn7327(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7345 github.com/goccy/pythonwasm2go/p2.Fn7345
func Fn7345(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7386 github.com/goccy/pythonwasm2go/p2.Fn7386
func Fn7386(m *base.Module, l0 int32) int32

//go:linkname Fn7393 github.com/goccy/pythonwasm2go/p2.Fn7393
func Fn7393(m *base.Module, l0 int32)

//go:linkname Fn7394 github.com/goccy/pythonwasm2go/p2.Fn7394
func Fn7394(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7459 github.com/goccy/pythonwasm2go/p2.Fn7459
func Fn7459(m *base.Module, l0 int32) int32

//go:linkname Fn7469 github.com/goccy/pythonwasm2go/p2.Fn7469
func Fn7469(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7506 github.com/goccy/pythonwasm2go/p2.Fn7506
func Fn7506(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7543 github.com/goccy/pythonwasm2go/p2.Fn7543
func Fn7543(m *base.Module, l0 int32) int32

//go:linkname Fn7544 github.com/goccy/pythonwasm2go/p2.Fn7544
func Fn7544(m *base.Module, l0 int32) int32

//go:linkname Fn7554 github.com/goccy/pythonwasm2go/p2.Fn7554
func Fn7554(m *base.Module, l0 int32) int32

//go:linkname Fn7580 github.com/goccy/pythonwasm2go/p2.Fn7580
func Fn7580(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7587 github.com/goccy/pythonwasm2go/p2.Fn7587
func Fn7587(m *base.Module, l0 int32) int32

//go:linkname Fn7950 github.com/goccy/pythonwasm2go/p2.Fn7950
func Fn7950(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7951 github.com/goccy/pythonwasm2go/p2.Fn7951
func Fn7951(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8125 github.com/goccy/pythonwasm2go/p2.Fn8125
func Fn8125(m *base.Module, l0 int32)

//go:linkname Fn8174 github.com/goccy/pythonwasm2go/p2.Fn8174
func Fn8174(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8176 github.com/goccy/pythonwasm2go/p2.Fn8176
func Fn8176(m *base.Module, l0 int32) int32

//go:linkname Fn8182 github.com/goccy/pythonwasm2go/p2.Fn8182
func Fn8182(m *base.Module, l0 int32) int32

//go:linkname Fn8245 github.com/goccy/pythonwasm2go/p2.Fn8245
func Fn8245(m *base.Module, l0 int32) int32

//go:linkname Fn8247 github.com/goccy/pythonwasm2go/p2.Fn8247
func Fn8247(m *base.Module, l0 int32) int32

//go:linkname Fn8357 github.com/goccy/pythonwasm2go/p2.Fn8357
func Fn8357(m *base.Module, l0 int32) int32

//go:linkname Fn8472 github.com/goccy/pythonwasm2go/p0.Fn8472
func Fn8472(m *base.Module)

//go:linkname Fn8475 github.com/goccy/pythonwasm2go/p2.Fn8475
func Fn8475(m *base.Module, l0 int32)

//go:linkname Fn8478 github.com/goccy/pythonwasm2go/p2.Fn8478
func Fn8478(m *base.Module) int32

//go:linkname Fn8488 github.com/goccy/pythonwasm2go/p2.Fn8488
func Fn8488(m *base.Module, l0 int32) int32

//go:linkname Fn8498 github.com/goccy/pythonwasm2go/p2.Fn8498
func Fn8498(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8499 github.com/goccy/pythonwasm2go/p2.Fn8499
func Fn8499(m *base.Module, l0 int32) int32

//go:linkname Fn8507 github.com/goccy/pythonwasm2go/p2.Fn8507
func Fn8507(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8509 github.com/goccy/pythonwasm2go/p2.Fn8509
func Fn8509(m *base.Module, l0 int32)

//go:linkname Fn8512 github.com/goccy/pythonwasm2go/p2.Fn8512
func Fn8512(m *base.Module, l0 int32) int32

//go:linkname Fn8514 github.com/goccy/pythonwasm2go/p2.Fn8514
func Fn8514(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8515 github.com/goccy/pythonwasm2go/p2.Fn8515
func Fn8515(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8520 github.com/goccy/pythonwasm2go/p2.Fn8520
func Fn8520(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8525 github.com/goccy/pythonwasm2go/p0.Fn8525
func Fn8525(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32) int32

//go:linkname Fn8536 github.com/goccy/pythonwasm2go/p2.Fn8536
func Fn8536(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8537 github.com/goccy/pythonwasm2go/p2.Fn8537
func Fn8537(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8538 github.com/goccy/pythonwasm2go/p2.Fn8538
func Fn8538(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8539 github.com/goccy/pythonwasm2go/p2.Fn8539
func Fn8539(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8551 github.com/goccy/pythonwasm2go/p2.Fn8551
func Fn8551(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8562 github.com/goccy/pythonwasm2go/p2.Fn8562
func Fn8562(m *base.Module, l0 int32, l1 int64) int32

//go:linkname Fn8590 github.com/goccy/pythonwasm2go/p2.Fn8590
func Fn8590(m *base.Module, l0 int32) int32

//go:linkname Fn8597 github.com/goccy/pythonwasm2go/p2.Fn8597
func Fn8597(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8598 github.com/goccy/pythonwasm2go/p2.Fn8598
func Fn8598(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int64, l6 int32) int32

//go:linkname Fn8602 github.com/goccy/pythonwasm2go/p2.Fn8602
func Fn8602(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8608 github.com/goccy/pythonwasm2go/p2.Fn8608
func Fn8608(m *base.Module, l0 int32)

//go:linkname Fn8621 github.com/goccy/pythonwasm2go/p2.Fn8621
func Fn8621(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8622 github.com/goccy/pythonwasm2go/p2.Fn8622
func Fn8622(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8625 github.com/goccy/pythonwasm2go/p0.Fn8625
func Fn8625(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8626 github.com/goccy/pythonwasm2go/p0.Fn8626
func Fn8626(m *base.Module) int32

//go:linkname Fn8627 github.com/goccy/pythonwasm2go/p2.Fn8627
func Fn8627(m *base.Module, l0 int32) int32

//go:linkname Fn8628 github.com/goccy/pythonwasm2go/p0.Fn8628
func Fn8628(m *base.Module, l0 int32) int32

//go:linkname Fn8652 github.com/goccy/pythonwasm2go/p2.Fn8652
func Fn8652(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8653 github.com/goccy/pythonwasm2go/p2.Fn8653
func Fn8653(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8654 github.com/goccy/pythonwasm2go/p2.Fn8654
func Fn8654(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn8655 github.com/goccy/pythonwasm2go/p2.Fn8655
func Fn8655(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32

//go:linkname Fn8657 github.com/goccy/pythonwasm2go/p2.Fn8657
func Fn8657(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8660 github.com/goccy/pythonwasm2go/p2.Fn8660
func Fn8660(m *base.Module, l0 int32) int32

//go:linkname Fn8668 github.com/goccy/pythonwasm2go/p2.Fn8668
func Fn8668(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8669 github.com/goccy/pythonwasm2go/p2.Fn8669
func Fn8669(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8670 github.com/goccy/pythonwasm2go/p2.Fn8670
func Fn8670(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8672 github.com/goccy/pythonwasm2go/p2.Fn8672
func Fn8672(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8673 github.com/goccy/pythonwasm2go/p2.Fn8673
func Fn8673(m *base.Module, l0 int32) int32

//go:linkname Fn8675 github.com/goccy/pythonwasm2go/p2.Fn8675
func Fn8675(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8676 github.com/goccy/pythonwasm2go/p2.Fn8676
func Fn8676(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8677 github.com/goccy/pythonwasm2go/p2.Fn8677
func Fn8677(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32

//go:linkname Fn8679 github.com/goccy/pythonwasm2go/p2.Fn8679
func Fn8679(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32

//go:linkname Fn8680 github.com/goccy/pythonwasm2go/p2.Fn8680
func Fn8680(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8681 github.com/goccy/pythonwasm2go/p2.Fn8681
func Fn8681(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8684 github.com/goccy/pythonwasm2go/p2.Fn8684
func Fn8684(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8690 github.com/goccy/pythonwasm2go/p2.Fn8690
func Fn8690(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8691 github.com/goccy/pythonwasm2go/p2.Fn8691
func Fn8691(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8693 github.com/goccy/pythonwasm2go/p2.Fn8693
func Fn8693(m *base.Module, l0 int32) int32

//go:linkname Fn8703 github.com/goccy/pythonwasm2go/p2.Fn8703
func Fn8703(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8704 github.com/goccy/pythonwasm2go/p2.Fn8704
func Fn8704(m *base.Module, l0 int32) int32

//go:linkname Fn8708 github.com/goccy/pythonwasm2go/p2.Fn8708
func Fn8708(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8710 github.com/goccy/pythonwasm2go/p2.Fn8710
func Fn8710(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8711 github.com/goccy/pythonwasm2go/p2.Fn8711
func Fn8711(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8712 github.com/goccy/pythonwasm2go/p2.Fn8712
func Fn8712(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8719 github.com/goccy/pythonwasm2go/p2.Fn8719
func Fn8719(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8723 github.com/goccy/pythonwasm2go/p2.Fn8723
func Fn8723(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8729 github.com/goccy/pythonwasm2go/p2.Fn8729
func Fn8729(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8731 github.com/goccy/pythonwasm2go/p2.Fn8731
func Fn8731(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn8740 github.com/goccy/pythonwasm2go/p2.Fn8740
func Fn8740(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn8751 github.com/goccy/pythonwasm2go/p2.Fn8751
func Fn8751(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8752 github.com/goccy/pythonwasm2go/p2.Fn8752
func Fn8752(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8753 github.com/goccy/pythonwasm2go/p2.Fn8753
func Fn8753(m *base.Module, l0 int32) int32

//go:linkname Fn8754 github.com/goccy/pythonwasm2go/p2.Fn8754
func Fn8754(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8761 github.com/goccy/pythonwasm2go/p2.Fn8761
func Fn8761(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8762 github.com/goccy/pythonwasm2go/p2.Fn8762
func Fn8762(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8764 github.com/goccy/pythonwasm2go/p2.Fn8764
func Fn8764(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8768 github.com/goccy/pythonwasm2go/p2.Fn8768
func Fn8768(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8776 github.com/goccy/pythonwasm2go/p2.Fn8776
func Fn8776(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn8777 github.com/goccy/pythonwasm2go/p2.Fn8777
func Fn8777(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8793 github.com/goccy/pythonwasm2go/p2.Fn8793
func Fn8793(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8797 github.com/goccy/pythonwasm2go/p2.Fn8797
func Fn8797(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8815 github.com/goccy/pythonwasm2go/p2.Fn8815
func Fn8815(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int64

//go:linkname Fn8821 github.com/goccy/pythonwasm2go/p2.Fn8821
func Fn8821(m *base.Module, l0 int64) int32

//go:linkname Fn8866 github.com/goccy/pythonwasm2go/p2.Fn8866
func Fn8866(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8909 github.com/goccy/pythonwasm2go/p2.Fn8909
func Fn8909(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8912 github.com/goccy/pythonwasm2go/p2.Fn8912
func Fn8912(m *base.Module, l0 int32) int32

//go:linkname Fn8913 github.com/goccy/pythonwasm2go/p2.Fn8913
func Fn8913(m *base.Module, l0 int32) int32

//go:linkname Fn8915 github.com/goccy/pythonwasm2go/p2.Fn8915
func Fn8915(m *base.Module, l0 int32) int32

//go:linkname Fn8918 github.com/goccy/pythonwasm2go/p2.Fn8918
func Fn8918(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8919 github.com/goccy/pythonwasm2go/p2.Fn8919
func Fn8919(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8926 github.com/goccy/pythonwasm2go/p2.Fn8926
func Fn8926(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8932 github.com/goccy/pythonwasm2go/p2.Fn8932
func Fn8932(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8937 github.com/goccy/pythonwasm2go/p2.Fn8937
func Fn8937(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8939 github.com/goccy/pythonwasm2go/p2.Fn8939
func Fn8939(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8985 github.com/goccy/pythonwasm2go/p2.Fn8985
func Fn8985(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8986 github.com/goccy/pythonwasm2go/p2.Fn8986
func Fn8986(m *base.Module, l0 int32, l1 int32) int64

//go:linkname Fn8995 github.com/goccy/pythonwasm2go/p2.Fn8995
func Fn8995(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8996 github.com/goccy/pythonwasm2go/p2.Fn8996
func Fn8996(m *base.Module, l0 int32) int32

//go:linkname Fn8997 github.com/goccy/pythonwasm2go/p2.Fn8997
func Fn8997(m *base.Module, l0 int32) int32

//go:linkname Fn8998 github.com/goccy/pythonwasm2go/p2.Fn8998
func Fn8998(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8999 github.com/goccy/pythonwasm2go/p2.Fn8999
func Fn8999(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9000 github.com/goccy/pythonwasm2go/p2.Fn9000
func Fn9000(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9001 github.com/goccy/pythonwasm2go/p2.Fn9001
func Fn9001(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9013 github.com/goccy/pythonwasm2go/p2.Fn9013
func Fn9013(m *base.Module, l0 int32) int32

//go:linkname Fn9027 github.com/goccy/pythonwasm2go/p2.Fn9027
func Fn9027(m *base.Module, l0 int32) int32

//go:linkname Fn9031 github.com/goccy/pythonwasm2go/p2.Fn9031
func Fn9031(m *base.Module, l0 int32) int32

//go:linkname Fn9033 github.com/goccy/pythonwasm2go/p2.Fn9033
func Fn9033(m *base.Module, l0 int32) int32

//go:linkname Fn9034 github.com/goccy/pythonwasm2go/p2.Fn9034
func Fn9034(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9037 github.com/goccy/pythonwasm2go/p2.Fn9037
func Fn9037(m *base.Module)

//go:linkname Fn9038 github.com/goccy/pythonwasm2go/p2.Fn9038
func Fn9038(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9047 github.com/goccy/pythonwasm2go/p2.Fn9047
func Fn9047(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9063 github.com/goccy/pythonwasm2go/p2.Fn9063
func Fn9063(m *base.Module, l0 int32) int32

//go:linkname Fn9064 github.com/goccy/pythonwasm2go/p2.Fn9064
func Fn9064(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9067 github.com/goccy/pythonwasm2go/p2.Fn9067
func Fn9067(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9071 github.com/goccy/pythonwasm2go/p2.Fn9071
func Fn9071(m *base.Module, l0 int32) int32

//go:linkname Fn9088 github.com/goccy/pythonwasm2go/p2.Fn9088
func Fn9088(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9092 github.com/goccy/pythonwasm2go/p2.Fn9092
func Fn9092(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9096 github.com/goccy/pythonwasm2go/p2.Fn9096
func Fn9096(m *base.Module) int32

//go:linkname Fn9110 github.com/goccy/pythonwasm2go/p2.Fn9110
func Fn9110(m *base.Module, l0 int32) int32

//go:linkname Fn9111 github.com/goccy/pythonwasm2go/p2.Fn9111
func Fn9111(m *base.Module, l0 int32) int32

//go:linkname Fn9112 github.com/goccy/pythonwasm2go/p2.Fn9112
func Fn9112(m *base.Module, l0 int32) int32

//go:linkname Fn9113 github.com/goccy/pythonwasm2go/p2.Fn9113
func Fn9113(m *base.Module, l0 int32) int32

//go:linkname Fn9114 github.com/goccy/pythonwasm2go/p2.Fn9114
func Fn9114(m *base.Module, l0 int32) int32

//go:linkname Fn9115 github.com/goccy/pythonwasm2go/p2.Fn9115
func Fn9115(m *base.Module, l0 int32) int32

//go:linkname Fn9119 github.com/goccy/pythonwasm2go/p2.Fn9119
func Fn9119(m *base.Module, l0 int32)

//go:linkname Fn9123 github.com/goccy/pythonwasm2go/p2.Fn9123
func Fn9123(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9130 github.com/goccy/pythonwasm2go/p2.Fn9130
func Fn9130(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9137 github.com/goccy/pythonwasm2go/p2.Fn9137
func Fn9137(m *base.Module, l0 int32) int64

//go:linkname Fn9166 github.com/goccy/pythonwasm2go/p2.Fn9166
func Fn9166(m *base.Module, l0 int32) int32

//go:linkname Fn9168 github.com/goccy/pythonwasm2go/p2.Fn9168
func Fn9168(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9176 github.com/goccy/pythonwasm2go/p2.Fn9176
func Fn9176(m *base.Module, l0 int32) int32

//go:linkname Fn9177 github.com/goccy/pythonwasm2go/p2.Fn9177
func Fn9177(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9180 github.com/goccy/pythonwasm2go/p2.Fn9180
func Fn9180(m *base.Module, l0 int32) int32

//go:linkname Fn9191 github.com/goccy/pythonwasm2go/p2.Fn9191
func Fn9191(m *base.Module, l0 int32) int32

//go:linkname Fn9194 github.com/goccy/pythonwasm2go/p2.Fn9194
func Fn9194(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9197 github.com/goccy/pythonwasm2go/p2.Fn9197
func Fn9197(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9203 github.com/goccy/pythonwasm2go/p2.Fn9203
func Fn9203(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9213 github.com/goccy/pythonwasm2go/p2.Fn9213
func Fn9213(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9214 github.com/goccy/pythonwasm2go/p2.Fn9214
func Fn9214(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9216 github.com/goccy/pythonwasm2go/p2.Fn9216
func Fn9216(m *base.Module, l0 int32) int32

//go:linkname Fn9234 github.com/goccy/pythonwasm2go/p2.Fn9234
func Fn9234(m *base.Module, l0 int32) int32

//go:linkname Fn9235 github.com/goccy/pythonwasm2go/p2.Fn9235
func Fn9235(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9237 github.com/goccy/pythonwasm2go/p2.Fn9237
func Fn9237(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9244 github.com/goccy/pythonwasm2go/p2.Fn9244
func Fn9244(m *base.Module, l0 int32) int32

//go:linkname Fn9261 github.com/goccy/pythonwasm2go/p2.Fn9261
func Fn9261(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9262 github.com/goccy/pythonwasm2go/p2.Fn9262
func Fn9262(m *base.Module, l0 int32) int32

//go:linkname Fn9263 github.com/goccy/pythonwasm2go/p2.Fn9263
func Fn9263(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9270 github.com/goccy/pythonwasm2go/p2.Fn9270
func Fn9270(m *base.Module, l0 int32) int32

//go:linkname Fn9303 github.com/goccy/pythonwasm2go/p2.Fn9303
func Fn9303(m *base.Module, l0 int32) int32

//go:linkname Fn9304 github.com/goccy/pythonwasm2go/p2.Fn9304
func Fn9304(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9345 github.com/goccy/pythonwasm2go/p2.Fn9345
func Fn9345(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn9353 github.com/goccy/pythonwasm2go/p2.Fn9353
func Fn9353(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9362 github.com/goccy/pythonwasm2go/p2.Fn9362
func Fn9362(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9364 github.com/goccy/pythonwasm2go/p2.Fn9364
func Fn9364(m *base.Module, l0 int32)

//go:linkname Fn9367 github.com/goccy/pythonwasm2go/p0.Fn9367
func Fn9367(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9372 github.com/goccy/pythonwasm2go/p2.Fn9372
func Fn9372(m *base.Module, l0 int32)

//go:linkname Fn9373 github.com/goccy/pythonwasm2go/p2.Fn9373
func Fn9373(m *base.Module, l0 int32)

//go:linkname Fn9374 github.com/goccy/pythonwasm2go/p2.Fn9374
func Fn9374(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9375 github.com/goccy/pythonwasm2go/p2.Fn9375
func Fn9375(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn9378 github.com/goccy/pythonwasm2go/p2.Fn9378
func Fn9378(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn9397 github.com/goccy/pythonwasm2go/p2.Fn9397
func Fn9397(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9404 github.com/goccy/pythonwasm2go/p2.Fn9404
func Fn9404(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9425 github.com/goccy/pythonwasm2go/p2.Fn9425
func Fn9425(m *base.Module, l0 int32)

//go:linkname Fn9430 github.com/goccy/pythonwasm2go/p2.Fn9430
func Fn9430(m *base.Module) int32

//go:linkname Fn9432 github.com/goccy/pythonwasm2go/p2.Fn9432
func Fn9432(m *base.Module, l0 int32)

//go:linkname Fn9433 github.com/goccy/pythonwasm2go/p2.Fn9433
func Fn9433(m *base.Module, l0 int32) int64

//go:linkname Fn9435 github.com/goccy/pythonwasm2go/p2.Fn9435
func Fn9435(m *base.Module, l0 int32) int32

//go:linkname Fn9451 github.com/goccy/pythonwasm2go/p2.Fn9451
func Fn9451(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9453 github.com/goccy/pythonwasm2go/p2.Fn9453
func Fn9453(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9491 github.com/goccy/pythonwasm2go/p2.Fn9491
func Fn9491(m *base.Module, l0 int32) int32

//go:linkname Fn9518 github.com/goccy/pythonwasm2go/p2.Fn9518
func Fn9518(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9522 github.com/goccy/pythonwasm2go/p2.Fn9522
func Fn9522(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9537 github.com/goccy/pythonwasm2go/p2.Fn9537
func Fn9537(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9539 github.com/goccy/pythonwasm2go/p2.Fn9539
func Fn9539(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9543 github.com/goccy/pythonwasm2go/p2.Fn9543
func Fn9543(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9545 github.com/goccy/pythonwasm2go/p2.Fn9545
func Fn9545(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9557 github.com/goccy/pythonwasm2go/p2.Fn9557
func Fn9557(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn9561 github.com/goccy/pythonwasm2go/p2.Fn9561
func Fn9561(m *base.Module, l0 int32) int32

//go:linkname Fn9568 github.com/goccy/pythonwasm2go/p2.Fn9568
func Fn9568(m *base.Module, l0 int32)

//go:linkname Fn9576 github.com/goccy/pythonwasm2go/p2.Fn9576
func Fn9576(m *base.Module, l0 int32)

//go:linkname Fn9669 github.com/goccy/pythonwasm2go/p2.Fn9669
func Fn9669(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9679 github.com/goccy/pythonwasm2go/p2.Fn9679
func Fn9679(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9680 github.com/goccy/pythonwasm2go/p2.Fn9680
func Fn9680(m *base.Module, l0 int32) int32

//go:linkname Fn9700 github.com/goccy/pythonwasm2go/p2.Fn9700
func Fn9700(m *base.Module, l0 int32) int32

//go:linkname Fn9732 github.com/goccy/pythonwasm2go/p2.Fn9732
func Fn9732(m *base.Module)

//go:linkname Fn9734 github.com/goccy/pythonwasm2go/p2.Fn9734
func Fn9734(m *base.Module, l0 int32)

//go:linkname Fn9735 github.com/goccy/pythonwasm2go/p2.Fn9735
func Fn9735(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9737 github.com/goccy/pythonwasm2go/p2.Fn9737
func Fn9737(m *base.Module, l0 int32) int32

//go:linkname Fn9738 github.com/goccy/pythonwasm2go/p2.Fn9738
func Fn9738(m *base.Module, l0 int32)

//go:linkname Fn9740 github.com/goccy/pythonwasm2go/p2.Fn9740
func Fn9740(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9742 github.com/goccy/pythonwasm2go/p2.Fn9742
func Fn9742(m *base.Module) int32

//go:linkname Fn9743 github.com/goccy/pythonwasm2go/p2.Fn9743
func Fn9743(m *base.Module, l0 int32)

//go:linkname Fn9750 github.com/goccy/pythonwasm2go/p2.Fn9750
func Fn9750(m *base.Module) int64

//go:linkname Fn9753 github.com/goccy/pythonwasm2go/p2.Fn9753
func Fn9753(m *base.Module, l0 int32, l1 int64) int32

//go:linkname Fn9754 github.com/goccy/pythonwasm2go/p2.Fn9754
func Fn9754(m *base.Module, l0 int32, l1 int64, l2 int32) int64

//go:linkname Fn9755 github.com/goccy/pythonwasm2go/p2.Fn9755
func Fn9755(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9756 github.com/goccy/pythonwasm2go/p2.Fn9756
func Fn9756(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9757 github.com/goccy/pythonwasm2go/p2.Fn9757
func Fn9757(m *base.Module, l0 int32)

//go:linkname Fn9759 github.com/goccy/pythonwasm2go/p2.Fn9759
func Fn9759(m *base.Module, l0 int32) int32

//go:linkname Fn9760 github.com/goccy/pythonwasm2go/p2.Fn9760
func Fn9760(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn9761 github.com/goccy/pythonwasm2go/p2.Fn9761
func Fn9761(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9764 github.com/goccy/pythonwasm2go/p2.Fn9764
func Fn9764(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32

//go:linkname Fn9766 github.com/goccy/pythonwasm2go/p2.Fn9766
func Fn9766(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9767 github.com/goccy/pythonwasm2go/p2.Fn9767
func Fn9767(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9768 github.com/goccy/pythonwasm2go/p2.Fn9768
func Fn9768(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9769 github.com/goccy/pythonwasm2go/p2.Fn9769
func Fn9769(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9770 github.com/goccy/pythonwasm2go/p2.Fn9770
func Fn9770(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9771 github.com/goccy/pythonwasm2go/p2.Fn9771
func Fn9771(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9772 github.com/goccy/pythonwasm2go/p2.Fn9772
func Fn9772(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9773 github.com/goccy/pythonwasm2go/p2.Fn9773
func Fn9773(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9776 github.com/goccy/pythonwasm2go/p2.Fn9776
func Fn9776(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9779 github.com/goccy/pythonwasm2go/p2.Fn9779
func Fn9779(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9780 github.com/goccy/pythonwasm2go/p2.Fn9780
func Fn9780(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9782 github.com/goccy/pythonwasm2go/p2.Fn9782
func Fn9782(m *base.Module)

//go:linkname Fn9783 github.com/goccy/pythonwasm2go/p2.Fn9783
func Fn9783(m *base.Module)

//go:linkname Fn9785 github.com/goccy/pythonwasm2go/p2.Fn9785
func Fn9785(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9786 github.com/goccy/pythonwasm2go/p2.Fn9786
func Fn9786(m *base.Module, l0 int32) int32

//go:linkname Fn9790 github.com/goccy/pythonwasm2go/p2.Fn9790
func Fn9790(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9791 github.com/goccy/pythonwasm2go/p2.Fn9791
func Fn9791(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9792 github.com/goccy/pythonwasm2go/p2.Fn9792
func Fn9792(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9795 github.com/goccy/pythonwasm2go/p2.Fn9795
func Fn9795(m *base.Module, l0 int32) int32

//go:linkname Fn9797 github.com/goccy/pythonwasm2go/p2.Fn9797
func Fn9797(m *base.Module, l0 int32) int32

//go:linkname Fn9798 github.com/goccy/pythonwasm2go/p2.Fn9798
func Fn9798(m *base.Module, l0 int32) int32

//go:linkname Fn9799 github.com/goccy/pythonwasm2go/p2.Fn9799
func Fn9799(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9800 github.com/goccy/pythonwasm2go/p2.Fn9800
func Fn9800(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9803 github.com/goccy/pythonwasm2go/p2.Fn9803
func Fn9803(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn9804 github.com/goccy/pythonwasm2go/p2.Fn9804
func Fn9804(m *base.Module, l0 int32) int32

//go:linkname Fn9805 github.com/goccy/pythonwasm2go/p2.Fn9805
func Fn9805(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9809 github.com/goccy/pythonwasm2go/p2.Fn9809
func Fn9809(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9812 github.com/goccy/pythonwasm2go/p2.Fn9812
func Fn9812(m *base.Module, l0 int32) int32

//go:linkname Fn9813 github.com/goccy/pythonwasm2go/p2.Fn9813
func Fn9813(m *base.Module, l0 int32) int32

//go:linkname Fn9816 github.com/goccy/pythonwasm2go/p2.Fn9816
func Fn9816(m *base.Module, l0 int32) int32

//go:linkname Fn9817 github.com/goccy/pythonwasm2go/p2.Fn9817
func Fn9817(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9818 github.com/goccy/pythonwasm2go/p2.Fn9818
func Fn9818(m *base.Module, l0 int32) int32

//go:linkname Fn9819 github.com/goccy/pythonwasm2go/p2.Fn9819
func Fn9819(m *base.Module, l0 int32) int32

//go:linkname Fn9821 github.com/goccy/pythonwasm2go/p2.Fn9821
func Fn9821(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9824 github.com/goccy/pythonwasm2go/p2.Fn9824
func Fn9824(m *base.Module, l0 int32) int32

//go:linkname Fn9828 github.com/goccy/pythonwasm2go/p2.Fn9828
func Fn9828(m *base.Module, l0 int32) int32

//go:linkname Fn9829 github.com/goccy/pythonwasm2go/p2.Fn9829
func Fn9829(m *base.Module, l0 int32) int32

//go:linkname Fn9832 github.com/goccy/pythonwasm2go/p2.Fn9832
func Fn9832(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9833 github.com/goccy/pythonwasm2go/p2.Fn9833
func Fn9833(m *base.Module, l0 int32) int32

//go:linkname Fn9837 github.com/goccy/pythonwasm2go/p2.Fn9837
func Fn9837(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9841 github.com/goccy/pythonwasm2go/p2.Fn9841
func Fn9841(m *base.Module, l0 float64) float64

//go:linkname Fn9849 github.com/goccy/pythonwasm2go/p2.Fn9849
func Fn9849(m *base.Module, l0 float64) float64

//go:linkname Fn9851 github.com/goccy/pythonwasm2go/p2.Fn9851
func Fn9851(m *base.Module, l0 float64) float64

//go:linkname Fn9856 github.com/goccy/pythonwasm2go/p2.Fn9856
func Fn9856(m *base.Module, l0 int32) float64

//go:linkname Fn9857 github.com/goccy/pythonwasm2go/p2.Fn9857
func Fn9857(m *base.Module, l0 int32) float64

//go:linkname Fn9863 github.com/goccy/pythonwasm2go/p2.Fn9863
func Fn9863(m *base.Module, l0 float64, l1 int32) float64

//go:linkname Fn9865 github.com/goccy/pythonwasm2go/p2.Fn9865
func Fn9865(m *base.Module, l0 float64, l1 int32) float64

//go:linkname Fn9866 github.com/goccy/pythonwasm2go/p2.Fn9866
func Fn9866(m *base.Module, l0 float64) float64

//go:linkname Fn9870 github.com/goccy/pythonwasm2go/p2.Fn9870
func Fn9870(m *base.Module, l0 float64, l1 int32) float64

//go:linkname Fn9871 github.com/goccy/pythonwasm2go/p2.Fn9871
func Fn9871(m *base.Module, l0 float64, l1 float64) float64

//go:linkname Fn9873 github.com/goccy/pythonwasm2go/p2.Fn9873
func Fn9873(m *base.Module, l0 int64) int32

//go:linkname Fn9874 github.com/goccy/pythonwasm2go/p2.Fn9874
func Fn9874(m *base.Module, l0 float64) float64

//go:linkname Fn9875 github.com/goccy/pythonwasm2go/p2.Fn9875
func Fn9875(m *base.Module, l0 float64, l1 int32) float64

//go:linkname Fn9876 github.com/goccy/pythonwasm2go/p2.Fn9876
func Fn9876(m *base.Module, l0 float64) float64

//go:linkname Fn9877 github.com/goccy/pythonwasm2go/p2.Fn9877
func Fn9877(m *base.Module, l0 float64) float64

//go:linkname Fn9879 github.com/goccy/pythonwasm2go/p2.Fn9879
func Fn9879(m *base.Module, l0 float64) float64

//go:linkname Fn9880 github.com/goccy/pythonwasm2go/p2.Fn9880
func Fn9880(m *base.Module, l0 float64) float64

//go:linkname Fn9890 github.com/goccy/pythonwasm2go/p2.Fn9890
func Fn9890(m *base.Module, l0 int32)

//go:linkname Fn9891 github.com/goccy/pythonwasm2go/p2.Fn9891
func Fn9891(m *base.Module, l0 int32)

//go:linkname Fn9892 github.com/goccy/pythonwasm2go/p2.Fn9892
func Fn9892(m *base.Module, l0 int32) int32

//go:linkname Fn9894 github.com/goccy/pythonwasm2go/p2.Fn9894
func Fn9894(m *base.Module, l0 int32) int32

//go:linkname Fn9895 github.com/goccy/pythonwasm2go/p2.Fn9895
func Fn9895(m *base.Module, l0 int32) int32

//go:linkname Fn9896 github.com/goccy/pythonwasm2go/p2.Fn9896
func Fn9896(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9901 github.com/goccy/pythonwasm2go/p2.Fn9901
func Fn9901(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9902 github.com/goccy/pythonwasm2go/p2.Fn9902
func Fn9902(m *base.Module, l0 int32) int32

//go:linkname Fn9904 github.com/goccy/pythonwasm2go/p2.Fn9904
func Fn9904(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9905 github.com/goccy/pythonwasm2go/p2.Fn9905
func Fn9905(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9906 github.com/goccy/pythonwasm2go/p2.Fn9906
func Fn9906(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9907 github.com/goccy/pythonwasm2go/p2.Fn9907
func Fn9907(m *base.Module, l0 int32) int32

//go:linkname Fn9909 github.com/goccy/pythonwasm2go/p2.Fn9909
func Fn9909(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9911 github.com/goccy/pythonwasm2go/p2.Fn9911
func Fn9911(m *base.Module, l0 int32) int32

//go:linkname Fn9912 github.com/goccy/pythonwasm2go/p2.Fn9912
func Fn9912(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9915 github.com/goccy/pythonwasm2go/p2.Fn9915
func Fn9915(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9916 github.com/goccy/pythonwasm2go/p2.Fn9916
func Fn9916(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9917 github.com/goccy/pythonwasm2go/p2.Fn9917
func Fn9917(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9922 github.com/goccy/pythonwasm2go/p2.Fn9922
func Fn9922(m *base.Module, l0 int32)

//go:linkname Fn9923 github.com/goccy/pythonwasm2go/p2.Fn9923
func Fn9923(m *base.Module, l0 int32) int32

//go:linkname Fn9925 github.com/goccy/pythonwasm2go/p2.Fn9925
func Fn9925(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9926 github.com/goccy/pythonwasm2go/p2.Fn9926
func Fn9926(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9927 github.com/goccy/pythonwasm2go/p2.Fn9927
func Fn9927(m *base.Module, l0 int32) int32

//go:linkname Fn9930 github.com/goccy/pythonwasm2go/p2.Fn9930
func Fn9930(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9931 github.com/goccy/pythonwasm2go/p2.Fn9931
func Fn9931(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9932 github.com/goccy/pythonwasm2go/p2.Fn9932
func Fn9932(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9934 github.com/goccy/pythonwasm2go/p2.Fn9934
func Fn9934(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9936 github.com/goccy/pythonwasm2go/p2.Fn9936
func Fn9936(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9937 github.com/goccy/pythonwasm2go/p2.Fn9937
func Fn9937(m *base.Module, l0 int32) int32

//go:linkname Fn9938 github.com/goccy/pythonwasm2go/p2.Fn9938
func Fn9938(m *base.Module, l0 int32) int32

//go:linkname Fn9939 github.com/goccy/pythonwasm2go/p2.Fn9939
func Fn9939(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9941 github.com/goccy/pythonwasm2go/p2.Fn9941
func Fn9941(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9943 github.com/goccy/pythonwasm2go/p2.Fn9943
func Fn9943(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9944 github.com/goccy/pythonwasm2go/p2.Fn9944
func Fn9944(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9945 github.com/goccy/pythonwasm2go/p2.Fn9945
func Fn9945(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9946 github.com/goccy/pythonwasm2go/p2.Fn9946
func Fn9946(m *base.Module, l0 int32) int32

//go:linkname Fn9947 github.com/goccy/pythonwasm2go/p2.Fn9947
func Fn9947(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9948 github.com/goccy/pythonwasm2go/p2.Fn9948
func Fn9948(m *base.Module, l0 int32) int32

//go:linkname Fn9951 github.com/goccy/pythonwasm2go/p2.Fn9951
func Fn9951(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9960 github.com/goccy/pythonwasm2go/p0.Fn9960
func Fn9960(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn9963 github.com/goccy/pythonwasm2go/p0.Fn9963
func Fn9963(m *base.Module, l0 int32) int32

//go:linkname Fn9965 github.com/goccy/pythonwasm2go/p2.Fn9965
func Fn9965(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9968 github.com/goccy/pythonwasm2go/p2.Fn9968
func Fn9968(m *base.Module, l0 int32, l1 int64, l2 int64)
