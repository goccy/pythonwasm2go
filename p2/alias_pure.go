//go:build (!amd64 && !arm64) || purego

package p2

import (
	base "github.com/goccy/pythonwasm2go/base"
	_ "unsafe"
)

//go:linkname Fn91 github.com/goccy/pythonwasm2go/p0.Fn91
func Fn91(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn93 github.com/goccy/pythonwasm2go/p0.Fn93
func Fn93(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn95 github.com/goccy/pythonwasm2go/p0.Fn95
func Fn95(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn97 github.com/goccy/pythonwasm2go/p0.Fn97
func Fn97(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn100 github.com/goccy/pythonwasm2go/p0.Fn100
func Fn100(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn105 github.com/goccy/pythonwasm2go/p0.Fn105
func Fn105(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn106 github.com/goccy/pythonwasm2go/p0.Fn106
func Fn106(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn107 github.com/goccy/pythonwasm2go/p0.Fn107
func Fn107(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn108 github.com/goccy/pythonwasm2go/p0.Fn108
func Fn108(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn109 github.com/goccy/pythonwasm2go/p0.Fn109
func Fn109(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn112 github.com/goccy/pythonwasm2go/p0.Fn112
func Fn112(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn114 github.com/goccy/pythonwasm2go/p0.Fn114
func Fn114(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn116 github.com/goccy/pythonwasm2go/p0.Fn116
func Fn116(m *base.Module, l0 int32)

//go:linkname Fn118 github.com/goccy/pythonwasm2go/p0.Fn118
func Fn118(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn119 github.com/goccy/pythonwasm2go/p0.Fn119
func Fn119(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn120 github.com/goccy/pythonwasm2go/p0.Fn120
func Fn120(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn121 github.com/goccy/pythonwasm2go/p0.Fn121
func Fn121(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn123 github.com/goccy/pythonwasm2go/p0.Fn123
func Fn123(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn125 github.com/goccy/pythonwasm2go/p0.Fn125
func Fn125(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn137 github.com/goccy/pythonwasm2go/p1.Fn137
func Fn137(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn138 github.com/goccy/pythonwasm2go/p1.Fn138
func Fn138(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn139 github.com/goccy/pythonwasm2go/p1.Fn139
func Fn139(m *base.Module, l0 int32, l1 float64, l2 int32)

//go:linkname Fn143 github.com/goccy/pythonwasm2go/p0.Fn143
func Fn143(m *base.Module, l0 int32) int32

//go:linkname Fn147 github.com/goccy/pythonwasm2go/p0.Fn147
func Fn147(m *base.Module, l0 int32) int32

//go:linkname Fn149 github.com/goccy/pythonwasm2go/p0.Fn149
func Fn149(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn185 github.com/goccy/pythonwasm2go/p1.Fn185
func Fn185(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn190 github.com/goccy/pythonwasm2go/p1.Fn190
func Fn190(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn203 github.com/goccy/pythonwasm2go/p1.Fn203
func Fn203(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn204 github.com/goccy/pythonwasm2go/p1.Fn204
func Fn204(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn206 github.com/goccy/pythonwasm2go/p0.Fn206
func Fn206(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn209 github.com/goccy/pythonwasm2go/p0.Fn209
func Fn209(m *base.Module, l0 int32) int32

//go:linkname Fn210 github.com/goccy/pythonwasm2go/p0.Fn210
func Fn210(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn211 github.com/goccy/pythonwasm2go/p0.Fn211
func Fn211(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn215 github.com/goccy/pythonwasm2go/p0.Fn215
func Fn215(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn217 github.com/goccy/pythonwasm2go/p0.Fn217
func Fn217(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn218 github.com/goccy/pythonwasm2go/p0.Fn218
func Fn218(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn226 github.com/goccy/pythonwasm2go/p1.Fn226
func Fn226(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn233 github.com/goccy/pythonwasm2go/p1.Fn233
func Fn233(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn234 github.com/goccy/pythonwasm2go/p1.Fn234
func Fn234(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn236 github.com/goccy/pythonwasm2go/p1.Fn236
func Fn236(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn237 github.com/goccy/pythonwasm2go/p1.Fn237
func Fn237(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn238 github.com/goccy/pythonwasm2go/p1.Fn238
func Fn238(m *base.Module, l0 int32) int32

//go:linkname Fn311 github.com/goccy/pythonwasm2go/p0.Fn311
func Fn311(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn313 github.com/goccy/pythonwasm2go/p0.Fn313
func Fn313(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn314 github.com/goccy/pythonwasm2go/p0.Fn314
func Fn314(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn315 github.com/goccy/pythonwasm2go/p0.Fn315
func Fn315(m *base.Module, l0 int32) int32

//go:linkname Fn316 github.com/goccy/pythonwasm2go/p1.Fn316
func Fn316(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn318 github.com/goccy/pythonwasm2go/p1.Fn318
func Fn318(m *base.Module, l0 int32) int32

//go:linkname Fn334 github.com/goccy/pythonwasm2go/p1.Fn334
func Fn334(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn337 github.com/goccy/pythonwasm2go/p1.Fn337
func Fn337(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn376 github.com/goccy/pythonwasm2go/p1.Fn376
func Fn376(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn405 github.com/goccy/pythonwasm2go/p0.Fn405
func Fn405(m *base.Module)

//go:linkname Fn406 github.com/goccy/pythonwasm2go/p0.Fn406
func Fn406(m *base.Module, l0 int32) int32

//go:linkname Fn407 github.com/goccy/pythonwasm2go/p0.Fn407
func Fn407(m *base.Module, l0 int32) int32

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

//go:linkname Fn417 github.com/goccy/pythonwasm2go/p0.Fn417
func Fn417(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn422 github.com/goccy/pythonwasm2go/p0.Fn422
func Fn422(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn429 github.com/goccy/pythonwasm2go/p0.Fn429
func Fn429(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn430 github.com/goccy/pythonwasm2go/p1.Fn430
func Fn430(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn434 github.com/goccy/pythonwasm2go/p1.Fn434
func Fn434(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn450 github.com/goccy/pythonwasm2go/p1.Fn450
func Fn450(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn474 github.com/goccy/pythonwasm2go/p1.Fn474
func Fn474(m *base.Module, l0 int32) int32

//go:linkname Fn475 github.com/goccy/pythonwasm2go/p1.Fn475
func Fn475(m *base.Module, l0 int32) int32

//go:linkname Fn478 github.com/goccy/pythonwasm2go/p0.Fn478
func Fn478(m *base.Module, l0 int32) int32

//go:linkname Fn482 github.com/goccy/pythonwasm2go/p0.Fn482
func Fn482(m *base.Module, l0 int32) int32

//go:linkname Fn483 github.com/goccy/pythonwasm2go/p0.Fn483
func Fn483(m *base.Module, l0 int32) int32

//go:linkname Fn484 github.com/goccy/pythonwasm2go/p0.Fn484
func Fn484(m *base.Module, l0 int32) int32

//go:linkname Fn485 github.com/goccy/pythonwasm2go/p0.Fn485
func Fn485(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn486 github.com/goccy/pythonwasm2go/p0.Fn486
func Fn486(m *base.Module, l0 int32) int32

//go:linkname Fn487 github.com/goccy/pythonwasm2go/p0.Fn487
func Fn487(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn488 github.com/goccy/pythonwasm2go/p0.Fn488
func Fn488(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn490 github.com/goccy/pythonwasm2go/p0.Fn490
func Fn490(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn496 github.com/goccy/pythonwasm2go/p0.Fn496
func Fn496(m *base.Module, l0 int32) int32

//go:linkname Fn497 github.com/goccy/pythonwasm2go/p0.Fn497
func Fn497(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn500 github.com/goccy/pythonwasm2go/p0.Fn500
func Fn500(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn502 github.com/goccy/pythonwasm2go/p0.Fn502
func Fn502(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn503 github.com/goccy/pythonwasm2go/p0.Fn503
func Fn503(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn504 github.com/goccy/pythonwasm2go/p0.Fn504
func Fn504(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn506 github.com/goccy/pythonwasm2go/p0.Fn506
func Fn506(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn514 github.com/goccy/pythonwasm2go/p0.Fn514
func Fn514(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn515 github.com/goccy/pythonwasm2go/p0.Fn515
func Fn515(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn522 github.com/goccy/pythonwasm2go/p1.Fn522
func Fn522(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn524 github.com/goccy/pythonwasm2go/p1.Fn524
func Fn524(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn527 github.com/goccy/pythonwasm2go/p1.Fn527
func Fn527(m *base.Module, l0 int32) int32

//go:linkname Fn534 github.com/goccy/pythonwasm2go/p0.Fn534
func Fn534(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn535 github.com/goccy/pythonwasm2go/p0.Fn535
func Fn535(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn536 github.com/goccy/pythonwasm2go/p0.Fn536
func Fn536(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn538 github.com/goccy/pythonwasm2go/p0.Fn538
func Fn538(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn539 github.com/goccy/pythonwasm2go/p0.Fn539
func Fn539(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn543 github.com/goccy/pythonwasm2go/p0.Fn543
func Fn543(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn544 github.com/goccy/pythonwasm2go/p0.Fn544
func Fn544(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn549 github.com/goccy/pythonwasm2go/p0.Fn549
func Fn549(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn551 github.com/goccy/pythonwasm2go/p0.Fn551
func Fn551(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn552 github.com/goccy/pythonwasm2go/p0.Fn552
func Fn552(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn578 github.com/goccy/pythonwasm2go/p0.Fn578
func Fn578(m *base.Module, l0 int32) int32

//go:linkname Fn601 github.com/goccy/pythonwasm2go/p1.Fn601
func Fn601(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn623 github.com/goccy/pythonwasm2go/p1.Fn623
func Fn623(m *base.Module, l0 int32) int32

//go:linkname Fn626 github.com/goccy/pythonwasm2go/p0.Fn626
func Fn626(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn633 github.com/goccy/pythonwasm2go/p0.Fn633
func Fn633(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn644 github.com/goccy/pythonwasm2go/p0.Fn644
func Fn644(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn655 github.com/goccy/pythonwasm2go/p1.Fn655
func Fn655(m *base.Module, l0 int32)

//go:linkname Fn656 github.com/goccy/pythonwasm2go/p1.Fn656
func Fn656(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn668 github.com/goccy/pythonwasm2go/p1.Fn668
func Fn668(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn671 github.com/goccy/pythonwasm2go/p1.Fn671
func Fn671(m *base.Module, l0 int32) int32

//go:linkname Fn689 github.com/goccy/pythonwasm2go/p1.Fn689
func Fn689(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn694 github.com/goccy/pythonwasm2go/p1.Fn694
func Fn694(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn726 github.com/goccy/pythonwasm2go/p1.Fn726
func Fn726(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn728 github.com/goccy/pythonwasm2go/p1.Fn728
func Fn728(m *base.Module) int32

//go:linkname Fn729 github.com/goccy/pythonwasm2go/p0.Fn729
func Fn729(m *base.Module, l0 float64) int32

//go:linkname Fn733 github.com/goccy/pythonwasm2go/p1.Fn733
func Fn733(m *base.Module, l0 int32) int32

//go:linkname Fn735 github.com/goccy/pythonwasm2go/p0.Fn735
func Fn735(m *base.Module, l0 int32) float64

//go:linkname Fn744 github.com/goccy/pythonwasm2go/p1.Fn744
func Fn744(m *base.Module, l0 float64, l1 int32, l2 int32) int32

//go:linkname Fn745 github.com/goccy/pythonwasm2go/p1.Fn745
func Fn745(m *base.Module, l0 float64, l1 int32, l2 int32) int32

//go:linkname Fn746 github.com/goccy/pythonwasm2go/p1.Fn746
func Fn746(m *base.Module, l0 float64, l1 int32, l2 int32) int32

//go:linkname Fn748 github.com/goccy/pythonwasm2go/p1.Fn748
func Fn748(m *base.Module, l0 int32, l1 int32) float64

//go:linkname Fn749 github.com/goccy/pythonwasm2go/p1.Fn749
func Fn749(m *base.Module, l0 int32, l1 int32) float64

//go:linkname Fn778 github.com/goccy/pythonwasm2go/p0.Fn778
func Fn778(m *base.Module, l0 int32) int32

//go:linkname Fn793 github.com/goccy/pythonwasm2go/p0.Fn793
func Fn793(m *base.Module, l0 int32) int32

//go:linkname Fn794 github.com/goccy/pythonwasm2go/p0.Fn794
func Fn794(m *base.Module, l0 int32) int32

//go:linkname Fn795 github.com/goccy/pythonwasm2go/p0.Fn795
func Fn795(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn796 github.com/goccy/pythonwasm2go/p0.Fn796
func Fn796(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn801 github.com/goccy/pythonwasm2go/p0.Fn801
func Fn801(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn802 github.com/goccy/pythonwasm2go/p0.Fn802
func Fn802(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn803 github.com/goccy/pythonwasm2go/p0.Fn803
func Fn803(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn804 github.com/goccy/pythonwasm2go/p0.Fn804
func Fn804(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn805 github.com/goccy/pythonwasm2go/p0.Fn805
func Fn805(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn806 github.com/goccy/pythonwasm2go/p0.Fn806
func Fn806(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn807 github.com/goccy/pythonwasm2go/p0.Fn807
func Fn807(m *base.Module, l0 int32) int32

//go:linkname Fn808 github.com/goccy/pythonwasm2go/p0.Fn808
func Fn808(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn809 github.com/goccy/pythonwasm2go/p0.Fn809
func Fn809(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn812 github.com/goccy/pythonwasm2go/p0.Fn812
func Fn812(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn814 github.com/goccy/pythonwasm2go/p0.Fn814
func Fn814(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn817 github.com/goccy/pythonwasm2go/p0.Fn817
func Fn817(m *base.Module, l0 int32) int32

//go:linkname Fn820 github.com/goccy/pythonwasm2go/p0.Fn820
func Fn820(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn832 github.com/goccy/pythonwasm2go/p0.Fn832
func Fn832(m *base.Module, l0 int32) int32

//go:linkname Fn882 github.com/goccy/pythonwasm2go/p0.Fn882
func Fn882(m *base.Module, l0 int32) int32

//go:linkname Fn885 github.com/goccy/pythonwasm2go/p0.Fn885
func Fn885(m *base.Module, l0 int32) int32

//go:linkname Fn886 github.com/goccy/pythonwasm2go/p0.Fn886
func Fn886(m *base.Module, l0 int32) int32

//go:linkname Fn887 github.com/goccy/pythonwasm2go/p0.Fn887
func Fn887(m *base.Module, l0 int32) int32

//go:linkname Fn888 github.com/goccy/pythonwasm2go/p0.Fn888
func Fn888(m *base.Module, l0 int64) int32

//go:linkname Fn889 github.com/goccy/pythonwasm2go/p1.Fn889
func Fn889(m *base.Module, l0 float64) int32

//go:linkname Fn890 github.com/goccy/pythonwasm2go/p0.Fn890
func Fn890(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn891 github.com/goccy/pythonwasm2go/p0.Fn891
func Fn891(m *base.Module, l0 int32) int32

//go:linkname Fn892 github.com/goccy/pythonwasm2go/p0.Fn892
func Fn892(m *base.Module, l0 int32) int32

//go:linkname Fn893 github.com/goccy/pythonwasm2go/p0.Fn893
func Fn893(m *base.Module, l0 int32) int32

//go:linkname Fn896 github.com/goccy/pythonwasm2go/p0.Fn896
func Fn896(m *base.Module, l0 int32) int32

//go:linkname Fn901 github.com/goccy/pythonwasm2go/p1.Fn901
func Fn901(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn905 github.com/goccy/pythonwasm2go/p0.Fn905
func Fn905(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn906 github.com/goccy/pythonwasm2go/p1.Fn906
func Fn906(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn908 github.com/goccy/pythonwasm2go/p0.Fn908
func Fn908(m *base.Module, l0 int64) int32

//go:linkname Fn909 github.com/goccy/pythonwasm2go/p0.Fn909
func Fn909(m *base.Module, l0 int32) int64

//go:linkname Fn911 github.com/goccy/pythonwasm2go/p0.Fn911
func Fn911(m *base.Module, l0 int32) int64

//go:linkname Fn913 github.com/goccy/pythonwasm2go/p1.Fn913
func Fn913(m *base.Module, l0 int32, l1 int32) int64

//go:linkname Fn919 github.com/goccy/pythonwasm2go/p0.Fn919
func Fn919(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn920 github.com/goccy/pythonwasm2go/p1.Fn920
func Fn920(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn922 github.com/goccy/pythonwasm2go/p0.Fn922
func Fn922(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn927 github.com/goccy/pythonwasm2go/p1.Fn927
func Fn927(m *base.Module, l0 int32, l1 int32) float64

//go:linkname Fn930 github.com/goccy/pythonwasm2go/p0.Fn930
func Fn930(m *base.Module, l0 int32) float64

//go:linkname Fn934 github.com/goccy/pythonwasm2go/p1.Fn934
func Fn934(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn935 github.com/goccy/pythonwasm2go/p1.Fn935
func Fn935(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn938 github.com/goccy/pythonwasm2go/p1.Fn938
func Fn938(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn939 github.com/goccy/pythonwasm2go/p1.Fn939
func Fn939(m *base.Module, l0 int32, l1 int64) int32

//go:linkname Fn942 github.com/goccy/pythonwasm2go/p1.Fn942
func Fn942(m *base.Module, l0 int32, l1 int64) int32

//go:linkname Fn946 github.com/goccy/pythonwasm2go/p1.Fn946
func Fn946(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn950 github.com/goccy/pythonwasm2go/p1.Fn950
func Fn950(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn958 github.com/goccy/pythonwasm2go/p1.Fn958
func Fn958(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn974 github.com/goccy/pythonwasm2go/p1.Fn974
func Fn974(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn983 github.com/goccy/pythonwasm2go/p1.Fn983
func Fn983(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1011 github.com/goccy/pythonwasm2go/p1.Fn1011
func Fn1011(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1076 github.com/goccy/pythonwasm2go/p0.Fn1076
func Fn1076(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1082 github.com/goccy/pythonwasm2go/p1.Fn1082
func Fn1082(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1090 github.com/goccy/pythonwasm2go/p0.Fn1090
func Fn1090(m *base.Module) int32

//go:linkname Fn1091 github.com/goccy/pythonwasm2go/p0.Fn1091
func Fn1091(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1099 github.com/goccy/pythonwasm2go/p0.Fn1099
func Fn1099(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1102 github.com/goccy/pythonwasm2go/p0.Fn1102
func Fn1102(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1103 github.com/goccy/pythonwasm2go/p0.Fn1103
func Fn1103(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1108 github.com/goccy/pythonwasm2go/p1.Fn1108
func Fn1108(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1112 github.com/goccy/pythonwasm2go/p0.Fn1112
func Fn1112(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1113 github.com/goccy/pythonwasm2go/p0.Fn1113
func Fn1113(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1114 github.com/goccy/pythonwasm2go/p0.Fn1114
func Fn1114(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1115 github.com/goccy/pythonwasm2go/p0.Fn1115
func Fn1115(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1116 github.com/goccy/pythonwasm2go/p0.Fn1116
func Fn1116(m *base.Module, l0 int32) int32

//go:linkname Fn1118 github.com/goccy/pythonwasm2go/p0.Fn1118
func Fn1118(m *base.Module, l0 int32)

//go:linkname Fn1121 github.com/goccy/pythonwasm2go/p0.Fn1121
func Fn1121(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1122 github.com/goccy/pythonwasm2go/p0.Fn1122
func Fn1122(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1125 github.com/goccy/pythonwasm2go/p0.Fn1125
func Fn1125(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1126 github.com/goccy/pythonwasm2go/p0.Fn1126
func Fn1126(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1127 github.com/goccy/pythonwasm2go/p0.Fn1127
func Fn1127(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1129 github.com/goccy/pythonwasm2go/p0.Fn1129
func Fn1129(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1135 github.com/goccy/pythonwasm2go/p0.Fn1135
func Fn1135(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1136 github.com/goccy/pythonwasm2go/p0.Fn1136
func Fn1136(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1141 github.com/goccy/pythonwasm2go/p0.Fn1141
func Fn1141(m *base.Module, l0 int32)

//go:linkname Fn1142 github.com/goccy/pythonwasm2go/p0.Fn1142
func Fn1142(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1143 github.com/goccy/pythonwasm2go/p0.Fn1143
func Fn1143(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1145 github.com/goccy/pythonwasm2go/p1.Fn1145
func Fn1145(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1147 github.com/goccy/pythonwasm2go/p0.Fn1147
func Fn1147(m *base.Module, l0 int32) int32

//go:linkname Fn1149 github.com/goccy/pythonwasm2go/p1.Fn1149
func Fn1149(m *base.Module, l0 int32) int32

//go:linkname Fn1150 github.com/goccy/pythonwasm2go/p1.Fn1150
func Fn1150(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1151 github.com/goccy/pythonwasm2go/p0.Fn1151
func Fn1151(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1152 github.com/goccy/pythonwasm2go/p0.Fn1152
func Fn1152(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1153 github.com/goccy/pythonwasm2go/p0.Fn1153
func Fn1153(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1157 github.com/goccy/pythonwasm2go/p0.Fn1157
func Fn1157(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1159 github.com/goccy/pythonwasm2go/p0.Fn1159
func Fn1159(m *base.Module, l0 int32) int32

//go:linkname Fn1162 github.com/goccy/pythonwasm2go/p0.Fn1162
func Fn1162(m *base.Module, l0 int32) int32

//go:linkname Fn1163 github.com/goccy/pythonwasm2go/p0.Fn1163
func Fn1163(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1172 github.com/goccy/pythonwasm2go/p1.Fn1172
func Fn1172(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1179 github.com/goccy/pythonwasm2go/p0.Fn1179
func Fn1179(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1180 github.com/goccy/pythonwasm2go/p0.Fn1180
func Fn1180(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1205 github.com/goccy/pythonwasm2go/p0.Fn1205
func Fn1205(m *base.Module, l0 int32) int32

//go:linkname Fn1209 github.com/goccy/pythonwasm2go/p1.Fn1209
func Fn1209(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1221 github.com/goccy/pythonwasm2go/p0.Fn1221
func Fn1221(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1222 github.com/goccy/pythonwasm2go/p0.Fn1222
func Fn1222(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1249 github.com/goccy/pythonwasm2go/p0.Fn1249
func Fn1249(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1262 github.com/goccy/pythonwasm2go/p0.Fn1262
func Fn1262(m *base.Module) int32

//go:linkname Fn1263 github.com/goccy/pythonwasm2go/p0.Fn1263
func Fn1263(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1264 github.com/goccy/pythonwasm2go/p0.Fn1264
func Fn1264(m *base.Module, l0 int32) int32

//go:linkname Fn1269 github.com/goccy/pythonwasm2go/p0.Fn1269
func Fn1269(m *base.Module, l0 int32) int32

//go:linkname Fn1280 github.com/goccy/pythonwasm2go/p1.Fn1280
func Fn1280(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1291 github.com/goccy/pythonwasm2go/p1.Fn1291
func Fn1291(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1340 github.com/goccy/pythonwasm2go/p0.Fn1340
func Fn1340(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1368 github.com/goccy/pythonwasm2go/p1.Fn1368
func Fn1368(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1372 github.com/goccy/pythonwasm2go/p0.Fn1372
func Fn1372(m *base.Module, l0 int32) int32

//go:linkname Fn1378 github.com/goccy/pythonwasm2go/p0.Fn1378
func Fn1378(m *base.Module, l0 int32) int32

//go:linkname Fn1385 github.com/goccy/pythonwasm2go/p0.Fn1385
func Fn1385(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1421 github.com/goccy/pythonwasm2go/p1.Fn1421
func Fn1421(m *base.Module, l0 int32) int32

//go:linkname Fn1448 github.com/goccy/pythonwasm2go/p1.Fn1448
func Fn1448(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1449 github.com/goccy/pythonwasm2go/p1.Fn1449
func Fn1449(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn1451 github.com/goccy/pythonwasm2go/p1.Fn1451
func Fn1451(m *base.Module, l0 int32) int32

//go:linkname Fn1496 github.com/goccy/pythonwasm2go/p1.Fn1496
func Fn1496(m *base.Module, l0 int32) int32

//go:linkname Fn1499 github.com/goccy/pythonwasm2go/p1.Fn1499
func Fn1499(m *base.Module, l0 int32) int32

//go:linkname Fn1504 github.com/goccy/pythonwasm2go/p1.Fn1504
func Fn1504(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1532 github.com/goccy/pythonwasm2go/p1.Fn1532
func Fn1532(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1547 github.com/goccy/pythonwasm2go/p1.Fn1547
func Fn1547(m *base.Module, l0 int32) int32

//go:linkname Fn1557 github.com/goccy/pythonwasm2go/p1.Fn1557
func Fn1557(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1563 github.com/goccy/pythonwasm2go/p1.Fn1563
func Fn1563(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1573 github.com/goccy/pythonwasm2go/p1.Fn1573
func Fn1573(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1580 github.com/goccy/pythonwasm2go/p1.Fn1580
func Fn1580(m *base.Module, l0 int32) int32

//go:linkname Fn1582 github.com/goccy/pythonwasm2go/p1.Fn1582
func Fn1582(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1587 github.com/goccy/pythonwasm2go/p0.Fn1587
func Fn1587(m *base.Module, l0 int32) int32

//go:linkname Fn1588 github.com/goccy/pythonwasm2go/p1.Fn1588
func Fn1588(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1590 github.com/goccy/pythonwasm2go/p0.Fn1590
func Fn1590(m *base.Module, l0 int32) int32

//go:linkname Fn1597 github.com/goccy/pythonwasm2go/p1.Fn1597
func Fn1597(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1598 github.com/goccy/pythonwasm2go/p1.Fn1598
func Fn1598(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1599 github.com/goccy/pythonwasm2go/p1.Fn1599
func Fn1599(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1632 github.com/goccy/pythonwasm2go/p1.Fn1632
func Fn1632(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1636 github.com/goccy/pythonwasm2go/p1.Fn1636
func Fn1636(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1648 github.com/goccy/pythonwasm2go/p1.Fn1648
func Fn1648(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1650 github.com/goccy/pythonwasm2go/p1.Fn1650
func Fn1650(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1672 github.com/goccy/pythonwasm2go/p0.Fn1672
func Fn1672(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1673 github.com/goccy/pythonwasm2go/p0.Fn1673
func Fn1673(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1674 github.com/goccy/pythonwasm2go/p0.Fn1674
func Fn1674(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1675 github.com/goccy/pythonwasm2go/p0.Fn1675
func Fn1675(m *base.Module, l0 int32) int32

//go:linkname Fn1679 github.com/goccy/pythonwasm2go/p0.Fn1679
func Fn1679(m *base.Module, l0 int32) int32

//go:linkname Fn1702 github.com/goccy/pythonwasm2go/p0.Fn1702
func Fn1702(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn1706 github.com/goccy/pythonwasm2go/p1.Fn1706
func Fn1706(m *base.Module, l0 int32)

//go:linkname Fn1711 github.com/goccy/pythonwasm2go/p0.Fn1711
func Fn1711(m *base.Module, l0 int32) int32

//go:linkname Fn1715 github.com/goccy/pythonwasm2go/p0.Fn1715
func Fn1715(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1716 github.com/goccy/pythonwasm2go/p0.Fn1716
func Fn1716(m *base.Module, l0 int32) int32

//go:linkname Fn1717 github.com/goccy/pythonwasm2go/p0.Fn1717
func Fn1717(m *base.Module, l0 int32) int32

//go:linkname Fn1718 github.com/goccy/pythonwasm2go/p0.Fn1718
func Fn1718(m *base.Module, l0 int32) int32

//go:linkname Fn1724 github.com/goccy/pythonwasm2go/p0.Fn1724
func Fn1724(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1725 github.com/goccy/pythonwasm2go/p0.Fn1725
func Fn1725(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1727 github.com/goccy/pythonwasm2go/p0.Fn1727
func Fn1727(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1728 github.com/goccy/pythonwasm2go/p0.Fn1728
func Fn1728(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1730 github.com/goccy/pythonwasm2go/p0.Fn1730
func Fn1730(m *base.Module, l0 int32) int32

//go:linkname Fn1731 github.com/goccy/pythonwasm2go/p0.Fn1731
func Fn1731(m *base.Module, l0 int32) int32

//go:linkname Fn1732 github.com/goccy/pythonwasm2go/p0.Fn1732
func Fn1732(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1733 github.com/goccy/pythonwasm2go/p0.Fn1733
func Fn1733(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1738 github.com/goccy/pythonwasm2go/p0.Fn1738
func Fn1738(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1742 github.com/goccy/pythonwasm2go/p0.Fn1742
func Fn1742(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1747 github.com/goccy/pythonwasm2go/p1.Fn1747
func Fn1747(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1752 github.com/goccy/pythonwasm2go/p1.Fn1752
func Fn1752(m *base.Module, l0 int32) int32

//go:linkname Fn1773 github.com/goccy/pythonwasm2go/p1.Fn1773
func Fn1773(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1774 github.com/goccy/pythonwasm2go/p1.Fn1774
func Fn1774(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1779 github.com/goccy/pythonwasm2go/p1.Fn1779
func Fn1779(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1787 github.com/goccy/pythonwasm2go/p1.Fn1787
func Fn1787(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1793 github.com/goccy/pythonwasm2go/p1.Fn1793
func Fn1793(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1798 github.com/goccy/pythonwasm2go/p1.Fn1798
func Fn1798(m *base.Module, l0 int32) int32

//go:linkname Fn1801 github.com/goccy/pythonwasm2go/p1.Fn1801
func Fn1801(m *base.Module, l0 int32) int32

//go:linkname Fn1814 github.com/goccy/pythonwasm2go/p1.Fn1814
func Fn1814(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1822 github.com/goccy/pythonwasm2go/p1.Fn1822
func Fn1822(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1837 github.com/goccy/pythonwasm2go/p1.Fn1837
func Fn1837(m *base.Module, l0 int32)

//go:linkname Fn1871 github.com/goccy/pythonwasm2go/p1.Fn1871
func Fn1871(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1880 github.com/goccy/pythonwasm2go/p1.Fn1880
func Fn1880(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1893 github.com/goccy/pythonwasm2go/p1.Fn1893
func Fn1893(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn1897 github.com/goccy/pythonwasm2go/p1.Fn1897
func Fn1897(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32)

//go:linkname Fn1901 github.com/goccy/pythonwasm2go/p1.Fn1901
func Fn1901(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn1915 github.com/goccy/pythonwasm2go/p0.Fn1915
func Fn1915(m *base.Module, l0 int32)

//go:linkname Fn1917 github.com/goccy/pythonwasm2go/p0.Fn1917
func Fn1917(m *base.Module, l0 int32)

//go:linkname Fn1919 github.com/goccy/pythonwasm2go/p1.Fn1919
func Fn1919(m *base.Module, l0 int32)

//go:linkname Fn1921 github.com/goccy/pythonwasm2go/p0.Fn1921
func Fn1921(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1947 github.com/goccy/pythonwasm2go/p1.Fn1947
func Fn1947(m *base.Module, l0 int32)

//go:linkname Fn1948 github.com/goccy/pythonwasm2go/p1.Fn1948
func Fn1948(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1954 github.com/goccy/pythonwasm2go/p1.Fn1954
func Fn1954(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1956 github.com/goccy/pythonwasm2go/p1.Fn1956
func Fn1956(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn1981 github.com/goccy/pythonwasm2go/p1.Fn1981
func Fn1981(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1995 github.com/goccy/pythonwasm2go/p1.Fn1995
func Fn1995(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2007 github.com/goccy/pythonwasm2go/p0.Fn2007
func Fn2007(m *base.Module, l0 int32)

//go:linkname Fn2008 github.com/goccy/pythonwasm2go/p0.Fn2008
func Fn2008(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2014 github.com/goccy/pythonwasm2go/p0.Fn2014
func Fn2014(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2016 github.com/goccy/pythonwasm2go/p0.Fn2016
func Fn2016(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2027 github.com/goccy/pythonwasm2go/p0.Fn2027
func Fn2027(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2030 github.com/goccy/pythonwasm2go/p0.Fn2030
func Fn2030(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2036 github.com/goccy/pythonwasm2go/p0.Fn2036
func Fn2036(m *base.Module, l0 int32) int32

//go:linkname Fn2042 github.com/goccy/pythonwasm2go/p0.Fn2042
func Fn2042(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2046 github.com/goccy/pythonwasm2go/p0.Fn2046
func Fn2046(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2047 github.com/goccy/pythonwasm2go/p0.Fn2047
func Fn2047(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2048 github.com/goccy/pythonwasm2go/p0.Fn2048
func Fn2048(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2054 github.com/goccy/pythonwasm2go/p0.Fn2054
func Fn2054(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2055 github.com/goccy/pythonwasm2go/p0.Fn2055
func Fn2055(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2056 github.com/goccy/pythonwasm2go/p0.Fn2056
func Fn2056(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2060 github.com/goccy/pythonwasm2go/p0.Fn2060
func Fn2060(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2062 github.com/goccy/pythonwasm2go/p1.Fn2062
func Fn2062(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2064 github.com/goccy/pythonwasm2go/p1.Fn2064
func Fn2064(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2070 github.com/goccy/pythonwasm2go/p1.Fn2070
func Fn2070(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2072 github.com/goccy/pythonwasm2go/p1.Fn2072
func Fn2072(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2074 github.com/goccy/pythonwasm2go/p1.Fn2074
func Fn2074(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2099 github.com/goccy/pythonwasm2go/p1.Fn2099
func Fn2099(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2104 github.com/goccy/pythonwasm2go/p1.Fn2104
func Fn2104(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2114 github.com/goccy/pythonwasm2go/p0.Fn2114
func Fn2114(m *base.Module, l0 int32) int32

//go:linkname Fn2118 github.com/goccy/pythonwasm2go/p1.Fn2118
func Fn2118(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2127 github.com/goccy/pythonwasm2go/p1.Fn2127
func Fn2127(m *base.Module, l0 int32) int32

//go:linkname Fn2138 github.com/goccy/pythonwasm2go/p0.Fn2138
func Fn2138(m *base.Module, l0 int32) int32

//go:linkname Fn2139 github.com/goccy/pythonwasm2go/p0.Fn2139
func Fn2139(m *base.Module, l0 int32) int32

//go:linkname Fn2140 github.com/goccy/pythonwasm2go/p0.Fn2140
func Fn2140(m *base.Module, l0 int32) int32

//go:linkname Fn2143 github.com/goccy/pythonwasm2go/p0.Fn2143
func Fn2143(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2144 github.com/goccy/pythonwasm2go/p0.Fn2144
func Fn2144(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2146 github.com/goccy/pythonwasm2go/p0.Fn2146
func Fn2146(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2158 github.com/goccy/pythonwasm2go/p1.Fn2158
func Fn2158(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2189 github.com/goccy/pythonwasm2go/p0.Fn2189
func Fn2189(m *base.Module, l0 int32)

//go:linkname Fn2192 github.com/goccy/pythonwasm2go/p0.Fn2192
func Fn2192(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2193 github.com/goccy/pythonwasm2go/p0.Fn2193
func Fn2193(m *base.Module, l0 int32) int32

//go:linkname Fn2194 github.com/goccy/pythonwasm2go/p0.Fn2194
func Fn2194(m *base.Module, l0 int32) int32

//go:linkname Fn2195 github.com/goccy/pythonwasm2go/p0.Fn2195
func Fn2195(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2196 github.com/goccy/pythonwasm2go/p0.Fn2196
func Fn2196(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2201 github.com/goccy/pythonwasm2go/p0.Fn2201
func Fn2201(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2202 github.com/goccy/pythonwasm2go/p0.Fn2202
func Fn2202(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2203 github.com/goccy/pythonwasm2go/p0.Fn2203
func Fn2203(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2209 github.com/goccy/pythonwasm2go/p0.Fn2209
func Fn2209(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2213 github.com/goccy/pythonwasm2go/p1.Fn2213
func Fn2213(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2216 github.com/goccy/pythonwasm2go/p0.Fn2216
func Fn2216(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2222 github.com/goccy/pythonwasm2go/p0.Fn2222
func Fn2222(m *base.Module, l0 int32) int32

//go:linkname Fn2231 github.com/goccy/pythonwasm2go/p0.Fn2231
func Fn2231(m *base.Module, l0 int32) int32

//go:linkname Fn2237 github.com/goccy/pythonwasm2go/p1.Fn2237
func Fn2237(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2242 github.com/goccy/pythonwasm2go/p0.Fn2242
func Fn2242(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2245 github.com/goccy/pythonwasm2go/p0.Fn2245
func Fn2245(m *base.Module, l0 int32)

//go:linkname Fn2272 github.com/goccy/pythonwasm2go/p1.Fn2272
func Fn2272(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2273 github.com/goccy/pythonwasm2go/p0.Fn2273
func Fn2273(m *base.Module, l0 int32) int32

//go:linkname Fn2285 github.com/goccy/pythonwasm2go/p1.Fn2285
func Fn2285(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2289 github.com/goccy/pythonwasm2go/p0.Fn2289
func Fn2289(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2294 github.com/goccy/pythonwasm2go/p1.Fn2294
func Fn2294(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2295 github.com/goccy/pythonwasm2go/p1.Fn2295
func Fn2295(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2302 github.com/goccy/pythonwasm2go/p1.Fn2302
func Fn2302(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2304 github.com/goccy/pythonwasm2go/p0.Fn2304
func Fn2304(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2305 github.com/goccy/pythonwasm2go/p0.Fn2305
func Fn2305(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2307 github.com/goccy/pythonwasm2go/p1.Fn2307
func Fn2307(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2313 github.com/goccy/pythonwasm2go/p0.Fn2313
func Fn2313(m *base.Module, l0 int32) int32

//go:linkname Fn2319 github.com/goccy/pythonwasm2go/p1.Fn2319
func Fn2319(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2353 github.com/goccy/pythonwasm2go/p0.Fn2353
func Fn2353(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2363 github.com/goccy/pythonwasm2go/p1.Fn2363
func Fn2363(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2461 github.com/goccy/pythonwasm2go/p1.Fn2461
func Fn2461(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2471 github.com/goccy/pythonwasm2go/p1.Fn2471
func Fn2471(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32

//go:linkname Fn2484 github.com/goccy/pythonwasm2go/p1.Fn2484
func Fn2484(m *base.Module, l0 int32) int32

//go:linkname Fn2542 github.com/goccy/pythonwasm2go/p1.Fn2542
func Fn2542(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2546 github.com/goccy/pythonwasm2go/p0.Fn2546
func Fn2546(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2548 github.com/goccy/pythonwasm2go/p1.Fn2548
func Fn2548(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2553 github.com/goccy/pythonwasm2go/p0.Fn2553
func Fn2553(m *base.Module, l0 int32) int32

//go:linkname Fn2557 github.com/goccy/pythonwasm2go/p1.Fn2557
func Fn2557(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2561 github.com/goccy/pythonwasm2go/p0.Fn2561
func Fn2561(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2565 github.com/goccy/pythonwasm2go/p0.Fn2565
func Fn2565(m *base.Module, l0 int32, l1 int32, l2 int32) int32

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

//go:linkname Fn2576 github.com/goccy/pythonwasm2go/p0.Fn2576
func Fn2576(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2577 github.com/goccy/pythonwasm2go/p0.Fn2577
func Fn2577(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2579 github.com/goccy/pythonwasm2go/p0.Fn2579
func Fn2579(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2581 github.com/goccy/pythonwasm2go/p0.Fn2581
func Fn2581(m *base.Module, l0 int32) int32

//go:linkname Fn2583 github.com/goccy/pythonwasm2go/p0.Fn2583
func Fn2583(m *base.Module, l0 int32) int32

//go:linkname Fn2585 github.com/goccy/pythonwasm2go/p1.Fn2585
func Fn2585(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2587 github.com/goccy/pythonwasm2go/p0.Fn2587
func Fn2587(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2588 github.com/goccy/pythonwasm2go/p0.Fn2588
func Fn2588(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2589 github.com/goccy/pythonwasm2go/p0.Fn2589
func Fn2589(m *base.Module, l0 int32) int32

//go:linkname Fn2591 github.com/goccy/pythonwasm2go/p0.Fn2591
func Fn2591(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2592 github.com/goccy/pythonwasm2go/p0.Fn2592
func Fn2592(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2593 github.com/goccy/pythonwasm2go/p0.Fn2593
func Fn2593(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2597 github.com/goccy/pythonwasm2go/p0.Fn2597
func Fn2597(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2599 github.com/goccy/pythonwasm2go/p0.Fn2599
func Fn2599(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2601 github.com/goccy/pythonwasm2go/p0.Fn2601
func Fn2601(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2602 github.com/goccy/pythonwasm2go/p0.Fn2602
func Fn2602(m *base.Module, l0 int32) int32

//go:linkname Fn2605 github.com/goccy/pythonwasm2go/p0.Fn2605
func Fn2605(m *base.Module, l0 int32, l1 int32) int32

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

//go:linkname Fn2618 github.com/goccy/pythonwasm2go/p0.Fn2618
func Fn2618(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2619 github.com/goccy/pythonwasm2go/p0.Fn2619
func Fn2619(m *base.Module, l0 int32) int32

//go:linkname Fn2621 github.com/goccy/pythonwasm2go/p0.Fn2621
func Fn2621(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2626 github.com/goccy/pythonwasm2go/p0.Fn2626
func Fn2626(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2627 github.com/goccy/pythonwasm2go/p0.Fn2627
func Fn2627(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2628 github.com/goccy/pythonwasm2go/p0.Fn2628
func Fn2628(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2629 github.com/goccy/pythonwasm2go/p0.Fn2629
func Fn2629(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2630 github.com/goccy/pythonwasm2go/p0.Fn2630
func Fn2630(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2636 github.com/goccy/pythonwasm2go/p0.Fn2636
func Fn2636(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2637 github.com/goccy/pythonwasm2go/p0.Fn2637
func Fn2637(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2638 github.com/goccy/pythonwasm2go/p0.Fn2638
func Fn2638(m *base.Module, l0 int32) int32

//go:linkname Fn2639 github.com/goccy/pythonwasm2go/p0.Fn2639
func Fn2639(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2651 github.com/goccy/pythonwasm2go/p0.Fn2651
func Fn2651(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2652 github.com/goccy/pythonwasm2go/p0.Fn2652
func Fn2652(m *base.Module, l0 int32) int32

//go:linkname Fn2654 github.com/goccy/pythonwasm2go/p0.Fn2654
func Fn2654(m *base.Module, l0 int32) int32

//go:linkname Fn2657 github.com/goccy/pythonwasm2go/p1.Fn2657
func Fn2657(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2658 github.com/goccy/pythonwasm2go/p0.Fn2658
func Fn2658(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2659 github.com/goccy/pythonwasm2go/p0.Fn2659
func Fn2659(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2662 github.com/goccy/pythonwasm2go/p1.Fn2662
func Fn2662(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2663 github.com/goccy/pythonwasm2go/p0.Fn2663
func Fn2663(m *base.Module, l0 int32) int32

//go:linkname Fn2666 github.com/goccy/pythonwasm2go/p1.Fn2666
func Fn2666(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2668 github.com/goccy/pythonwasm2go/p1.Fn2668
func Fn2668(m *base.Module, l0 int32) int32

//go:linkname Fn2669 github.com/goccy/pythonwasm2go/p1.Fn2669
func Fn2669(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2671 github.com/goccy/pythonwasm2go/p1.Fn2671
func Fn2671(m *base.Module, l0 int32) int32

//go:linkname Fn2676 github.com/goccy/pythonwasm2go/p1.Fn2676
func Fn2676(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2677 github.com/goccy/pythonwasm2go/p1.Fn2677
func Fn2677(m *base.Module, l0 int32) int32

//go:linkname Fn2678 github.com/goccy/pythonwasm2go/p1.Fn2678
func Fn2678(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2683 github.com/goccy/pythonwasm2go/p0.Fn2683
func Fn2683(m *base.Module, l0 int32) int32

//go:linkname Fn2684 github.com/goccy/pythonwasm2go/p1.Fn2684
func Fn2684(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2691 github.com/goccy/pythonwasm2go/p1.Fn2691
func Fn2691(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2697 github.com/goccy/pythonwasm2go/p1.Fn2697
func Fn2697(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2701 github.com/goccy/pythonwasm2go/p1.Fn2701
func Fn2701(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2702 github.com/goccy/pythonwasm2go/p1.Fn2702
func Fn2702(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2703 github.com/goccy/pythonwasm2go/p1.Fn2703
func Fn2703(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2704 github.com/goccy/pythonwasm2go/p1.Fn2704
func Fn2704(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2707 github.com/goccy/pythonwasm2go/p0.Fn2707
func Fn2707(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2708 github.com/goccy/pythonwasm2go/p1.Fn2708
func Fn2708(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2710 github.com/goccy/pythonwasm2go/p1.Fn2710
func Fn2710(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2712 github.com/goccy/pythonwasm2go/p1.Fn2712
func Fn2712(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2714 github.com/goccy/pythonwasm2go/p0.Fn2714
func Fn2714(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2715 github.com/goccy/pythonwasm2go/p1.Fn2715
func Fn2715(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2718 github.com/goccy/pythonwasm2go/p1.Fn2718
func Fn2718(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2721 github.com/goccy/pythonwasm2go/p0.Fn2721
func Fn2721(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2723 github.com/goccy/pythonwasm2go/p1.Fn2723
func Fn2723(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2728 github.com/goccy/pythonwasm2go/p0.Fn2728
func Fn2728(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2729 github.com/goccy/pythonwasm2go/p0.Fn2729
func Fn2729(m *base.Module, l0 int32) int32

//go:linkname Fn2731 github.com/goccy/pythonwasm2go/p0.Fn2731
func Fn2731(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2734 github.com/goccy/pythonwasm2go/p0.Fn2734
func Fn2734(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2735 github.com/goccy/pythonwasm2go/p1.Fn2735
func Fn2735(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2736 github.com/goccy/pythonwasm2go/p1.Fn2736
func Fn2736(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2737 github.com/goccy/pythonwasm2go/p0.Fn2737
func Fn2737(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2749 github.com/goccy/pythonwasm2go/p0.Fn2749
func Fn2749(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2752 github.com/goccy/pythonwasm2go/p0.Fn2752
func Fn2752(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2769 github.com/goccy/pythonwasm2go/p0.Fn2769
func Fn2769(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2770 github.com/goccy/pythonwasm2go/p0.Fn2770
func Fn2770(m *base.Module, l0 int32) int32

//go:linkname Fn2782 github.com/goccy/pythonwasm2go/p1.Fn2782
func Fn2782(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2785 github.com/goccy/pythonwasm2go/p1.Fn2785
func Fn2785(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2786 github.com/goccy/pythonwasm2go/p1.Fn2786
func Fn2786(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2790 github.com/goccy/pythonwasm2go/p1.Fn2790
func Fn2790(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2791 github.com/goccy/pythonwasm2go/p1.Fn2791
func Fn2791(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2804 github.com/goccy/pythonwasm2go/p1.Fn2804
func Fn2804(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2853 github.com/goccy/pythonwasm2go/p1.Fn2853
func Fn2853(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2861 github.com/goccy/pythonwasm2go/p1.Fn2861
func Fn2861(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn2862 github.com/goccy/pythonwasm2go/p1.Fn2862
func Fn2862(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn2864 github.com/goccy/pythonwasm2go/p1.Fn2864
func Fn2864(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2896 github.com/goccy/pythonwasm2go/p1.Fn2896
func Fn2896(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2898 github.com/goccy/pythonwasm2go/p1.Fn2898
func Fn2898(m *base.Module, l0 int32) int32

//go:linkname Fn2899 github.com/goccy/pythonwasm2go/p1.Fn2899
func Fn2899(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2929 github.com/goccy/pythonwasm2go/p0.Fn2929
func Fn2929(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2940 github.com/goccy/pythonwasm2go/p0.Fn2940
func Fn2940(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2946 github.com/goccy/pythonwasm2go/p1.Fn2946
func Fn2946(m *base.Module, l0 int32)

//go:linkname Fn2994 github.com/goccy/pythonwasm2go/p0.Fn2994
func Fn2994(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2995 github.com/goccy/pythonwasm2go/p0.Fn2995
func Fn2995(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2998 github.com/goccy/pythonwasm2go/p0.Fn2998
func Fn2998(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3000 github.com/goccy/pythonwasm2go/p0.Fn3000
func Fn3000(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3003 github.com/goccy/pythonwasm2go/p0.Fn3003
func Fn3003(m *base.Module, l0 int32) int32

//go:linkname Fn3101 github.com/goccy/pythonwasm2go/p1.Fn3101
func Fn3101(m *base.Module, l0 int32) int32

//go:linkname Fn3106 github.com/goccy/pythonwasm2go/p0.Fn3106
func Fn3106(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3119 github.com/goccy/pythonwasm2go/p1.Fn3119
func Fn3119(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3153 github.com/goccy/pythonwasm2go/p0.Fn3153
func Fn3153(m *base.Module, l0 int32) int32

//go:linkname Fn3154 github.com/goccy/pythonwasm2go/p0.Fn3154
func Fn3154(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3155 github.com/goccy/pythonwasm2go/p0.Fn3155
func Fn3155(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3223 github.com/goccy/pythonwasm2go/p1.Fn3223
func Fn3223(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3240 github.com/goccy/pythonwasm2go/p1.Fn3240
func Fn3240(m *base.Module, l0 int32)

//go:linkname Fn3242 github.com/goccy/pythonwasm2go/p0.Fn3242
func Fn3242(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3244 github.com/goccy/pythonwasm2go/p1.Fn3244
func Fn3244(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3245 github.com/goccy/pythonwasm2go/p1.Fn3245
func Fn3245(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3247 github.com/goccy/pythonwasm2go/p0.Fn3247
func Fn3247(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3281 github.com/goccy/pythonwasm2go/p1.Fn3281
func Fn3281(m *base.Module) int32

//go:linkname Fn3285 github.com/goccy/pythonwasm2go/p0.Fn3285
func Fn3285(m *base.Module, l0 int32) int32

//go:linkname Fn3286 github.com/goccy/pythonwasm2go/p0.Fn3286
func Fn3286(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3292 github.com/goccy/pythonwasm2go/p0.Fn3292
func Fn3292(m *base.Module, l0 int32) int32

//go:linkname Fn3301 github.com/goccy/pythonwasm2go/p0.Fn3301
func Fn3301(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3304 github.com/goccy/pythonwasm2go/p0.Fn3304
func Fn3304(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3308 github.com/goccy/pythonwasm2go/p0.Fn3308
func Fn3308(m *base.Module, l0 int32) int32

//go:linkname Fn3309 github.com/goccy/pythonwasm2go/p0.Fn3309
func Fn3309(m *base.Module, l0 int32) int32

//go:linkname Fn3313 github.com/goccy/pythonwasm2go/p1.Fn3313
func Fn3313(m *base.Module, l0 int32) int32

//go:linkname Fn3315 github.com/goccy/pythonwasm2go/p1.Fn3315
func Fn3315(m *base.Module, l0 int32) int32

//go:linkname Fn3317 github.com/goccy/pythonwasm2go/p1.Fn3317
func Fn3317(m *base.Module, l0 int32) int32

//go:linkname Fn3319 github.com/goccy/pythonwasm2go/p1.Fn3319
func Fn3319(m *base.Module, l0 int32) int32

//go:linkname Fn3329 github.com/goccy/pythonwasm2go/p0.Fn3329
func Fn3329(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3346 github.com/goccy/pythonwasm2go/p1.Fn3346
func Fn3346(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3347 github.com/goccy/pythonwasm2go/p1.Fn3347
func Fn3347(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3349 github.com/goccy/pythonwasm2go/p0.Fn3349
func Fn3349(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3352 github.com/goccy/pythonwasm2go/p1.Fn3352
func Fn3352(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3378 github.com/goccy/pythonwasm2go/p0.Fn3378
func Fn3378(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3420 github.com/goccy/pythonwasm2go/p1.Fn3420
func Fn3420(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3425 github.com/goccy/pythonwasm2go/p1.Fn3425
func Fn3425(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3427 github.com/goccy/pythonwasm2go/p0.Fn3427
func Fn3427(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3428 github.com/goccy/pythonwasm2go/p0.Fn3428
func Fn3428(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3429 github.com/goccy/pythonwasm2go/p0.Fn3429
func Fn3429(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3445 github.com/goccy/pythonwasm2go/p1.Fn3445
func Fn3445(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3447 github.com/goccy/pythonwasm2go/p1.Fn3447
func Fn3447(m *base.Module, l0 int32)

//go:linkname Fn3464 github.com/goccy/pythonwasm2go/p1.Fn3464
func Fn3464(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3467 github.com/goccy/pythonwasm2go/p1.Fn3467
func Fn3467(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3471 github.com/goccy/pythonwasm2go/p1.Fn3471
func Fn3471(m *base.Module, l0 int32) int32

//go:linkname Fn3486 github.com/goccy/pythonwasm2go/p0.Fn3486
func Fn3486(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3487 github.com/goccy/pythonwasm2go/p1.Fn3487
func Fn3487(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3489 github.com/goccy/pythonwasm2go/p1.Fn3489
func Fn3489(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3541 github.com/goccy/pythonwasm2go/p1.Fn3541
func Fn3541(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3543 github.com/goccy/pythonwasm2go/p1.Fn3543
func Fn3543(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3579 github.com/goccy/pythonwasm2go/p0.Fn3579
func Fn3579(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3597 github.com/goccy/pythonwasm2go/p0.Fn3597
func Fn3597(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3607 github.com/goccy/pythonwasm2go/p0.Fn3607
func Fn3607(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3609 github.com/goccy/pythonwasm2go/p0.Fn3609
func Fn3609(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3613 github.com/goccy/pythonwasm2go/p0.Fn3613
func Fn3613(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3614 github.com/goccy/pythonwasm2go/p0.Fn3614
func Fn3614(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3618 github.com/goccy/pythonwasm2go/p0.Fn3618
func Fn3618(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3619 github.com/goccy/pythonwasm2go/p0.Fn3619
func Fn3619(m *base.Module, l0 int32)

//go:linkname Fn3622 github.com/goccy/pythonwasm2go/p0.Fn3622
func Fn3622(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3623 github.com/goccy/pythonwasm2go/p0.Fn3623
func Fn3623(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3625 github.com/goccy/pythonwasm2go/p0.Fn3625
func Fn3625(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3626 github.com/goccy/pythonwasm2go/p0.Fn3626
func Fn3626(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3627 github.com/goccy/pythonwasm2go/p0.Fn3627
func Fn3627(m *base.Module, l0 int32) int32

//go:linkname Fn3628 github.com/goccy/pythonwasm2go/p0.Fn3628
func Fn3628(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3640 github.com/goccy/pythonwasm2go/p0.Fn3640
func Fn3640(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3642 github.com/goccy/pythonwasm2go/p0.Fn3642
func Fn3642(m *base.Module)

//go:linkname Fn3643 github.com/goccy/pythonwasm2go/p0.Fn3643
func Fn3643(m *base.Module) int32

//go:linkname Fn3645 github.com/goccy/pythonwasm2go/p0.Fn3645
func Fn3645(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3647 github.com/goccy/pythonwasm2go/p0.Fn3647
func Fn3647(m *base.Module, l0 int32) int32

//go:linkname Fn3651 github.com/goccy/pythonwasm2go/p0.Fn3651
func Fn3651(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3653 github.com/goccy/pythonwasm2go/p1.Fn3653
func Fn3653(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3655 github.com/goccy/pythonwasm2go/p0.Fn3655
func Fn3655(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3656 github.com/goccy/pythonwasm2go/p0.Fn3656
func Fn3656(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3657 github.com/goccy/pythonwasm2go/p0.Fn3657
func Fn3657(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3658 github.com/goccy/pythonwasm2go/p1.Fn3658
func Fn3658(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn3662 github.com/goccy/pythonwasm2go/p1.Fn3662
func Fn3662(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3681 github.com/goccy/pythonwasm2go/p1.Fn3681
func Fn3681(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3689 github.com/goccy/pythonwasm2go/p1.Fn3689
func Fn3689(m *base.Module, l0 int32) int32

//go:linkname Fn3712 github.com/goccy/pythonwasm2go/p0.Fn3712
func Fn3712(m *base.Module, l0 int32) int32

//go:linkname Fn3713 github.com/goccy/pythonwasm2go/p1.Fn3713
func Fn3713(m *base.Module, l0 int32)

//go:linkname Fn3716 github.com/goccy/pythonwasm2go/p1.Fn3716
func Fn3716(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3725 github.com/goccy/pythonwasm2go/p0.Fn3725
func Fn3725(m *base.Module)

//go:linkname Fn3726 github.com/goccy/pythonwasm2go/p0.Fn3726
func Fn3726(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3735 github.com/goccy/pythonwasm2go/p0.Fn3735
func Fn3735(m *base.Module, l0 int32)

//go:linkname Fn3739 github.com/goccy/pythonwasm2go/p0.Fn3739
func Fn3739(m *base.Module, l0 int32) int32

//go:linkname Fn3741 github.com/goccy/pythonwasm2go/p0.Fn3741
func Fn3741(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3742 github.com/goccy/pythonwasm2go/p0.Fn3742
func Fn3742(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3749 github.com/goccy/pythonwasm2go/p0.Fn3749
func Fn3749(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3750 github.com/goccy/pythonwasm2go/p0.Fn3750
func Fn3750(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3755 github.com/goccy/pythonwasm2go/p0.Fn3755
func Fn3755(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3756 github.com/goccy/pythonwasm2go/p0.Fn3756
func Fn3756(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3760 github.com/goccy/pythonwasm2go/p1.Fn3760
func Fn3760(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3766 github.com/goccy/pythonwasm2go/p1.Fn3766
func Fn3766(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32

//go:linkname Fn3768 github.com/goccy/pythonwasm2go/p1.Fn3768
func Fn3768(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3781 github.com/goccy/pythonwasm2go/p0.Fn3781
func Fn3781(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3785 github.com/goccy/pythonwasm2go/p0.Fn3785
func Fn3785(m *base.Module, l0 int32)

//go:linkname Fn3786 github.com/goccy/pythonwasm2go/p0.Fn3786
func Fn3786(m *base.Module) int32

//go:linkname Fn3787 github.com/goccy/pythonwasm2go/p0.Fn3787
func Fn3787(m *base.Module, l0 int32)

//go:linkname Fn3794 github.com/goccy/pythonwasm2go/p0.Fn3794
func Fn3794(m *base.Module) int32

//go:linkname Fn3797 github.com/goccy/pythonwasm2go/p0.Fn3797
func Fn3797(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3806 github.com/goccy/pythonwasm2go/p1.Fn3806
func Fn3806(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3809 github.com/goccy/pythonwasm2go/p0.Fn3809
func Fn3809(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3810 github.com/goccy/pythonwasm2go/p0.Fn3810
func Fn3810(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3811 github.com/goccy/pythonwasm2go/p1.Fn3811
func Fn3811(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3813 github.com/goccy/pythonwasm2go/p1.Fn3813
func Fn3813(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3856 github.com/goccy/pythonwasm2go/p0.Fn3856
func Fn3856(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3858 github.com/goccy/pythonwasm2go/p0.Fn3858
func Fn3858(m *base.Module, l0 int32) int32

//go:linkname Fn3859 github.com/goccy/pythonwasm2go/p0.Fn3859
func Fn3859(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3874 github.com/goccy/pythonwasm2go/p1.Fn3874
func Fn3874(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3883 github.com/goccy/pythonwasm2go/p1.Fn3883
func Fn3883(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3887 github.com/goccy/pythonwasm2go/p1.Fn3887
func Fn3887(m *base.Module, l0 int32) int32

//go:linkname Fn3888 github.com/goccy/pythonwasm2go/p1.Fn3888
func Fn3888(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3893 github.com/goccy/pythonwasm2go/p0.Fn3893
func Fn3893(m *base.Module, l0 int32) int32

//go:linkname Fn3894 github.com/goccy/pythonwasm2go/p0.Fn3894
func Fn3894(m *base.Module, l0 int32) int32

//go:linkname Fn3895 github.com/goccy/pythonwasm2go/p0.Fn3895
func Fn3895(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3896 github.com/goccy/pythonwasm2go/p0.Fn3896
func Fn3896(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3897 github.com/goccy/pythonwasm2go/p1.Fn3897
func Fn3897(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3898 github.com/goccy/pythonwasm2go/p1.Fn3898
func Fn3898(m *base.Module, l0 int32) int32

//go:linkname Fn3901 github.com/goccy/pythonwasm2go/p0.Fn3901
func Fn3901(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3903 github.com/goccy/pythonwasm2go/p0.Fn3903
func Fn3903(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3907 github.com/goccy/pythonwasm2go/p1.Fn3907
func Fn3907(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3945 github.com/goccy/pythonwasm2go/p1.Fn3945
func Fn3945(m *base.Module, l0 int32)

//go:linkname Fn3947 github.com/goccy/pythonwasm2go/p1.Fn3947
func Fn3947(m *base.Module, l0 int32)

//go:linkname Fn3952 github.com/goccy/pythonwasm2go/p1.Fn3952
func Fn3952(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3981 github.com/goccy/pythonwasm2go/p1.Fn3981
func Fn3981(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3995 github.com/goccy/pythonwasm2go/p0.Fn3995
func Fn3995(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4008 github.com/goccy/pythonwasm2go/p1.Fn4008
func Fn4008(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4010 github.com/goccy/pythonwasm2go/p1.Fn4010
func Fn4010(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn4086 github.com/goccy/pythonwasm2go/p0.Fn4086
func Fn4086(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn4087 github.com/goccy/pythonwasm2go/p0.Fn4087
func Fn4087(m *base.Module, l0 int32) int32

//go:linkname Fn4090 github.com/goccy/pythonwasm2go/p0.Fn4090
func Fn4090(m *base.Module, l0 int32)

//go:linkname Fn4095 github.com/goccy/pythonwasm2go/p0.Fn4095
func Fn4095(m *base.Module, l0 int32)

//go:linkname Fn4096 github.com/goccy/pythonwasm2go/p0.Fn4096
func Fn4096(m *base.Module, l0 int32)

//go:linkname Fn4098 github.com/goccy/pythonwasm2go/p0.Fn4098
func Fn4098(m *base.Module, l0 int32) int32

//go:linkname Fn4099 github.com/goccy/pythonwasm2go/p0.Fn4099
func Fn4099(m *base.Module, l0 int32)

//go:linkname Fn4116 github.com/goccy/pythonwasm2go/p0.Fn4116
func Fn4116(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4118 github.com/goccy/pythonwasm2go/p1.Fn4118
func Fn4118(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4125 github.com/goccy/pythonwasm2go/p0.Fn4125
func Fn4125(m *base.Module, l0 int32) int32

//go:linkname Fn4140 github.com/goccy/pythonwasm2go/p0.Fn4140
func Fn4140(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4145 github.com/goccy/pythonwasm2go/p0.Fn4145
func Fn4145(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4161 github.com/goccy/pythonwasm2go/p0.Fn4161
func Fn4161(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn4164 github.com/goccy/pythonwasm2go/p0.Fn4164
func Fn4164(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4165 github.com/goccy/pythonwasm2go/p0.Fn4165
func Fn4165(m *base.Module, l0 int32)

//go:linkname Fn4184 github.com/goccy/pythonwasm2go/p1.Fn4184
func Fn4184(m *base.Module, l0 int32, l1 float64) int32

//go:linkname Fn4196 github.com/goccy/pythonwasm2go/p1.Fn4196
func Fn4196(m *base.Module, l0 int32)

//go:linkname Fn4201 github.com/goccy/pythonwasm2go/p0.Fn4201
func Fn4201(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4209 github.com/goccy/pythonwasm2go/p0.Fn4209
func Fn4209(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn4210 github.com/goccy/pythonwasm2go/p0.Fn4210
func Fn4210(m *base.Module) int32

//go:linkname Fn4212 github.com/goccy/pythonwasm2go/p0.Fn4212
func Fn4212(m *base.Module, l0 int32)

//go:linkname Fn4213 github.com/goccy/pythonwasm2go/p0.Fn4213
func Fn4213(m *base.Module, l0 int32)

//go:linkname Fn4218 github.com/goccy/pythonwasm2go/p0.Fn4218
func Fn4218(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4220 github.com/goccy/pythonwasm2go/p1.Fn4220
func Fn4220(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4225 github.com/goccy/pythonwasm2go/p0.Fn4225
func Fn4225(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4233 github.com/goccy/pythonwasm2go/p0.Fn4233
func Fn4233(m *base.Module, l0 int32)

//go:linkname Fn4234 github.com/goccy/pythonwasm2go/p0.Fn4234
func Fn4234(m *base.Module, l0 int32)

//go:linkname Fn4236 github.com/goccy/pythonwasm2go/p0.Fn4236
func Fn4236(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4246 github.com/goccy/pythonwasm2go/p0.Fn4246
func Fn4246(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4247 github.com/goccy/pythonwasm2go/p0.Fn4247
func Fn4247(m *base.Module, l0 int32) int32

//go:linkname Fn4261 github.com/goccy/pythonwasm2go/p0.Fn4261
func Fn4261(m *base.Module, l0 int32)

//go:linkname Fn4270 github.com/goccy/pythonwasm2go/p0.Fn4270
func Fn4270(m *base.Module, l0 int32)

//go:linkname Fn4272 github.com/goccy/pythonwasm2go/p0.Fn4272
func Fn4272(m *base.Module) int32

//go:linkname Fn4273 github.com/goccy/pythonwasm2go/p0.Fn4273
func Fn4273(m *base.Module, l0 int32)

//go:linkname Fn4275 github.com/goccy/pythonwasm2go/p0.Fn4275
func Fn4275(m *base.Module) int32

//go:linkname Fn4281 github.com/goccy/pythonwasm2go/p1.Fn4281
func Fn4281(m *base.Module, l0 int32)

//go:linkname Fn4283 github.com/goccy/pythonwasm2go/p1.Fn4283
func Fn4283(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4285 github.com/goccy/pythonwasm2go/p1.Fn4285
func Fn4285(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)

//go:linkname Fn4289 github.com/goccy/pythonwasm2go/p1.Fn4289
func Fn4289(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4301 github.com/goccy/pythonwasm2go/p1.Fn4301
func Fn4301(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4318 github.com/goccy/pythonwasm2go/p1.Fn4318
func Fn4318(m *base.Module, l0 int32) int32

//go:linkname Fn4320 github.com/goccy/pythonwasm2go/p1.Fn4320
func Fn4320(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4330 github.com/goccy/pythonwasm2go/p1.Fn4330
func Fn4330(m *base.Module, l0 int32) int32

//go:linkname Fn4333 github.com/goccy/pythonwasm2go/p0.Fn4333
func Fn4333(m *base.Module, l0 int32) int32

//go:linkname Fn4338 github.com/goccy/pythonwasm2go/p0.Fn4338
func Fn4338(m *base.Module, l0 int32) int32

//go:linkname Fn4382 github.com/goccy/pythonwasm2go/p0.Fn4382
func Fn4382(m *base.Module, l0 int32) int32

//go:linkname Fn4383 github.com/goccy/pythonwasm2go/p0.Fn4383
func Fn4383(m *base.Module, l0 int32) int32

//go:linkname Fn4384 github.com/goccy/pythonwasm2go/p0.Fn4384
func Fn4384(m *base.Module, l0 int32) int32

//go:linkname Fn4393 github.com/goccy/pythonwasm2go/p0.Fn4393
func Fn4393(m *base.Module, l0 int32) int32

//go:linkname Fn4417 github.com/goccy/pythonwasm2go/p0.Fn4417
func Fn4417(m *base.Module, l0 int32) int32

//go:linkname Fn4475 github.com/goccy/pythonwasm2go/p1.Fn4475
func Fn4475(m *base.Module, l0 int32) int32

//go:linkname Fn4493 github.com/goccy/pythonwasm2go/p1.Fn4493
func Fn4493(m *base.Module, l0 int32) int32

//go:linkname Fn4495 github.com/goccy/pythonwasm2go/p1.Fn4495
func Fn4495(m *base.Module, l0 int32) int32

//go:linkname Fn4496 github.com/goccy/pythonwasm2go/p1.Fn4496
func Fn4496(m *base.Module, l0 int32) int32

//go:linkname Fn4546 github.com/goccy/pythonwasm2go/p1.Fn4546
func Fn4546(m *base.Module, l0 int32) int32

//go:linkname Fn4556 github.com/goccy/pythonwasm2go/p1.Fn4556
func Fn4556(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4558 github.com/goccy/pythonwasm2go/p1.Fn4558
func Fn4558(m *base.Module, l0 int32) int32

//go:linkname Fn4560 github.com/goccy/pythonwasm2go/p1.Fn4560
func Fn4560(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn4562 github.com/goccy/pythonwasm2go/p1.Fn4562
func Fn4562(m *base.Module, l0 int32) int32

//go:linkname Fn4565 github.com/goccy/pythonwasm2go/p1.Fn4565
func Fn4565(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4571 github.com/goccy/pythonwasm2go/p1.Fn4571
func Fn4571(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn4576 github.com/goccy/pythonwasm2go/p0.Fn4576
func Fn4576(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4578 github.com/goccy/pythonwasm2go/p0.Fn4578
func Fn4578(m *base.Module, l0 int32)

//go:linkname Fn4583 github.com/goccy/pythonwasm2go/p0.Fn4583
func Fn4583(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4597 github.com/goccy/pythonwasm2go/p1.Fn4597
func Fn4597(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4601 github.com/goccy/pythonwasm2go/p0.Fn4601
func Fn4601(m *base.Module)

//go:linkname Fn4602 github.com/goccy/pythonwasm2go/p0.Fn4602
func Fn4602(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4604 github.com/goccy/pythonwasm2go/p1.Fn4604
func Fn4604(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4616 github.com/goccy/pythonwasm2go/p0.Fn4616
func Fn4616(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4617 github.com/goccy/pythonwasm2go/p0.Fn4617
func Fn4617(m *base.Module, l0 int32) int32

//go:linkname Fn4619 github.com/goccy/pythonwasm2go/p0.Fn4619
func Fn4619(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4620 github.com/goccy/pythonwasm2go/p0.Fn4620
func Fn4620(m *base.Module, l0 int32) int32

//go:linkname Fn4625 github.com/goccy/pythonwasm2go/p0.Fn4625
func Fn4625(m *base.Module, l0 int64) int64

//go:linkname Fn4627 github.com/goccy/pythonwasm2go/p1.Fn4627
func Fn4627(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4655 github.com/goccy/pythonwasm2go/p1.Fn4655
func Fn4655(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4656 github.com/goccy/pythonwasm2go/p1.Fn4656
func Fn4656(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4662 github.com/goccy/pythonwasm2go/p0.Fn4662
func Fn4662(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4666 github.com/goccy/pythonwasm2go/p1.Fn4666
func Fn4666(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4669 github.com/goccy/pythonwasm2go/p1.Fn4669
func Fn4669(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4695 github.com/goccy/pythonwasm2go/p1.Fn4695
func Fn4695(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4700 github.com/goccy/pythonwasm2go/p0.Fn4700
func Fn4700(m *base.Module, l0 int32) int32

//go:linkname Fn4702 github.com/goccy/pythonwasm2go/p0.Fn4702
func Fn4702(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4705 github.com/goccy/pythonwasm2go/p0.Fn4705
func Fn4705(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4708 github.com/goccy/pythonwasm2go/p0.Fn4708
func Fn4708(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4709 github.com/goccy/pythonwasm2go/p0.Fn4709
func Fn4709(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4710 github.com/goccy/pythonwasm2go/p0.Fn4710
func Fn4710(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4712 github.com/goccy/pythonwasm2go/p0.Fn4712
func Fn4712(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4714 github.com/goccy/pythonwasm2go/p1.Fn4714
func Fn4714(m *base.Module, l0 int32) int32

//go:linkname Fn4722 github.com/goccy/pythonwasm2go/p0.Fn4722
func Fn4722(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4794 github.com/goccy/pythonwasm2go/p0.Fn4794
func Fn4794(m *base.Module, l0 int32, l1 int64) int32

//go:linkname Fn4796 github.com/goccy/pythonwasm2go/p0.Fn4796
func Fn4796(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4803 github.com/goccy/pythonwasm2go/p0.Fn4803
func Fn4803(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4810 github.com/goccy/pythonwasm2go/p0.Fn4810
func Fn4810(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4815 github.com/goccy/pythonwasm2go/p0.Fn4815
func Fn4815(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4818 github.com/goccy/pythonwasm2go/p0.Fn4818
func Fn4818(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4834 github.com/goccy/pythonwasm2go/p1.Fn4834
func Fn4834(m *base.Module, l0 int32) int32

//go:linkname Fn4848 github.com/goccy/pythonwasm2go/p1.Fn4848
func Fn4848(m *base.Module) int32

//go:linkname Fn4849 github.com/goccy/pythonwasm2go/p0.Fn4849
func Fn4849(m *base.Module)

//go:linkname Fn4859 github.com/goccy/pythonwasm2go/p1.Fn4859
func Fn4859(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4874 github.com/goccy/pythonwasm2go/p1.Fn4874
func Fn4874(m *base.Module, l0 float64, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4876 github.com/goccy/pythonwasm2go/p1.Fn4876
func Fn4876(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4878 github.com/goccy/pythonwasm2go/p0.Fn4878
func Fn4878(m *base.Module, l0 int32, l1 int32) float64

//go:linkname Fn4884 github.com/goccy/pythonwasm2go/p1.Fn4884
func Fn4884(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4898 github.com/goccy/pythonwasm2go/p1.Fn4898
func Fn4898(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4900 github.com/goccy/pythonwasm2go/p1.Fn4900
func Fn4900(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn4907 github.com/goccy/pythonwasm2go/p1.Fn4907
func Fn4907(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4913 github.com/goccy/pythonwasm2go/p1.Fn4913
func Fn4913(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4916 github.com/goccy/pythonwasm2go/p1.Fn4916
func Fn4916(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4921 github.com/goccy/pythonwasm2go/p1.Fn4921
func Fn4921(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn4922 github.com/goccy/pythonwasm2go/p1.Fn4922
func Fn4922(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4924 github.com/goccy/pythonwasm2go/p1.Fn4924
func Fn4924(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn4937 github.com/goccy/pythonwasm2go/p0.Fn4937
func Fn4937(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4938 github.com/goccy/pythonwasm2go/p0.Fn4938
func Fn4938(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4943 github.com/goccy/pythonwasm2go/p1.Fn4943
func Fn4943(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4949 github.com/goccy/pythonwasm2go/p1.Fn4949
func Fn4949(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4950 github.com/goccy/pythonwasm2go/p1.Fn4950
func Fn4950(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn4987 github.com/goccy/pythonwasm2go/p1.Fn4987
func Fn4987(m *base.Module, l0 int32) int32

//go:linkname Fn4992 github.com/goccy/pythonwasm2go/p1.Fn4992
func Fn4992(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5082 github.com/goccy/pythonwasm2go/p1.Fn5082
func Fn5082(m *base.Module, l0 int32) int32

//go:linkname Fn5117 github.com/goccy/pythonwasm2go/p0.Fn5117
func Fn5117(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5121 github.com/goccy/pythonwasm2go/p1.Fn5121
func Fn5121(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5122 github.com/goccy/pythonwasm2go/p1.Fn5122
func Fn5122(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5125 github.com/goccy/pythonwasm2go/p1.Fn5125
func Fn5125(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5139 github.com/goccy/pythonwasm2go/p1.Fn5139
func Fn5139(m *base.Module, l0 int32) int32

//go:linkname Fn5144 github.com/goccy/pythonwasm2go/p1.Fn5144
func Fn5144(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5145 github.com/goccy/pythonwasm2go/p1.Fn5145
func Fn5145(m *base.Module, l0 int32)

//go:linkname Fn5228 github.com/goccy/pythonwasm2go/p1.Fn5228
func Fn5228(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5232 github.com/goccy/pythonwasm2go/p1.Fn5232
func Fn5232(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn5247 github.com/goccy/pythonwasm2go/p1.Fn5247
func Fn5247(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5254 github.com/goccy/pythonwasm2go/p1.Fn5254
func Fn5254(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5265 github.com/goccy/pythonwasm2go/p1.Fn5265
func Fn5265(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5272 github.com/goccy/pythonwasm2go/p0.Fn5272
func Fn5272(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5306 github.com/goccy/pythonwasm2go/p1.Fn5306
func Fn5306(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5313 github.com/goccy/pythonwasm2go/p1.Fn5313
func Fn5313(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5324 github.com/goccy/pythonwasm2go/p1.Fn5324
func Fn5324(m *base.Module, l0 int32)

//go:linkname Fn5331 github.com/goccy/pythonwasm2go/p1.Fn5331
func Fn5331(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5340 github.com/goccy/pythonwasm2go/p0.Fn5340
func Fn5340(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5344 github.com/goccy/pythonwasm2go/p1.Fn5344
func Fn5344(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5368 github.com/goccy/pythonwasm2go/p1.Fn5368
func Fn5368(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5379 github.com/goccy/pythonwasm2go/p1.Fn5379
func Fn5379(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5398 github.com/goccy/pythonwasm2go/p1.Fn5398
func Fn5398(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5408 github.com/goccy/pythonwasm2go/p1.Fn5408
func Fn5408(m *base.Module, l0 int32) int32

//go:linkname Fn5416 github.com/goccy/pythonwasm2go/p1.Fn5416
func Fn5416(m *base.Module, l0 int32) int32

//go:linkname Fn5429 github.com/goccy/pythonwasm2go/p1.Fn5429
func Fn5429(m *base.Module, l0 int32) int32

//go:linkname Fn5451 github.com/goccy/pythonwasm2go/p1.Fn5451
func Fn5451(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5465 github.com/goccy/pythonwasm2go/p1.Fn5465
func Fn5465(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5467 github.com/goccy/pythonwasm2go/p1.Fn5467
func Fn5467(m *base.Module, l0 int32) int32

//go:linkname Fn5481 github.com/goccy/pythonwasm2go/p1.Fn5481
func Fn5481(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5485 github.com/goccy/pythonwasm2go/p1.Fn5485
func Fn5485(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5489 github.com/goccy/pythonwasm2go/p1.Fn5489
func Fn5489(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5490 github.com/goccy/pythonwasm2go/p1.Fn5490
func Fn5490(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5491 github.com/goccy/pythonwasm2go/p1.Fn5491
func Fn5491(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5595 github.com/goccy/pythonwasm2go/p1.Fn5595
func Fn5595(m *base.Module, l0 int32) int32

//go:linkname Fn5603 github.com/goccy/pythonwasm2go/p1.Fn5603
func Fn5603(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5611 github.com/goccy/pythonwasm2go/p1.Fn5611
func Fn5611(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5613 github.com/goccy/pythonwasm2go/p0.Fn5613
func Fn5613(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5615 github.com/goccy/pythonwasm2go/p1.Fn5615
func Fn5615(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5714 github.com/goccy/pythonwasm2go/p1.Fn5714
func Fn5714(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5717 github.com/goccy/pythonwasm2go/p1.Fn5717
func Fn5717(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5720 github.com/goccy/pythonwasm2go/p1.Fn5720
func Fn5720(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5723 github.com/goccy/pythonwasm2go/p1.Fn5723
func Fn5723(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5731 github.com/goccy/pythonwasm2go/p1.Fn5731
func Fn5731(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5738 github.com/goccy/pythonwasm2go/p1.Fn5738
func Fn5738(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5742 github.com/goccy/pythonwasm2go/p1.Fn5742
func Fn5742(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5756 github.com/goccy/pythonwasm2go/p1.Fn5756
func Fn5756(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5757 github.com/goccy/pythonwasm2go/p1.Fn5757
func Fn5757(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5763 github.com/goccy/pythonwasm2go/p1.Fn5763
func Fn5763(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5764 github.com/goccy/pythonwasm2go/p1.Fn5764
func Fn5764(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5770 github.com/goccy/pythonwasm2go/p1.Fn5770
func Fn5770(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5810 github.com/goccy/pythonwasm2go/p1.Fn5810
func Fn5810(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5811 github.com/goccy/pythonwasm2go/p1.Fn5811
func Fn5811(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5818 github.com/goccy/pythonwasm2go/p1.Fn5818
func Fn5818(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5826 github.com/goccy/pythonwasm2go/p1.Fn5826
func Fn5826(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5837 github.com/goccy/pythonwasm2go/p1.Fn5837
func Fn5837(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5840 github.com/goccy/pythonwasm2go/p1.Fn5840
func Fn5840(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5842 github.com/goccy/pythonwasm2go/p1.Fn5842
func Fn5842(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5844 github.com/goccy/pythonwasm2go/p1.Fn5844
func Fn5844(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5849 github.com/goccy/pythonwasm2go/p1.Fn5849
func Fn5849(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn5860 github.com/goccy/pythonwasm2go/p1.Fn5860
func Fn5860(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn5865 github.com/goccy/pythonwasm2go/p1.Fn5865
func Fn5865(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5871 github.com/goccy/pythonwasm2go/p1.Fn5871
func Fn5871(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5877 github.com/goccy/pythonwasm2go/p1.Fn5877
func Fn5877(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5880 github.com/goccy/pythonwasm2go/p1.Fn5880
func Fn5880(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5881 github.com/goccy/pythonwasm2go/p1.Fn5881
func Fn5881(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5884 github.com/goccy/pythonwasm2go/p1.Fn5884
func Fn5884(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5891 github.com/goccy/pythonwasm2go/p1.Fn5891
func Fn5891(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5892 github.com/goccy/pythonwasm2go/p1.Fn5892
func Fn5892(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5907 github.com/goccy/pythonwasm2go/p1.Fn5907
func Fn5907(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5913 github.com/goccy/pythonwasm2go/p1.Fn5913
func Fn5913(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5929 github.com/goccy/pythonwasm2go/p1.Fn5929
func Fn5929(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5933 github.com/goccy/pythonwasm2go/p1.Fn5933
func Fn5933(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5935 github.com/goccy/pythonwasm2go/p1.Fn5935
func Fn5935(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5947 github.com/goccy/pythonwasm2go/p1.Fn5947
func Fn5947(m *base.Module, l0 int32) int32

//go:linkname Fn5954 github.com/goccy/pythonwasm2go/p1.Fn5954
func Fn5954(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32

//go:linkname Fn5982 github.com/goccy/pythonwasm2go/p1.Fn5982
func Fn5982(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5983 github.com/goccy/pythonwasm2go/p1.Fn5983
func Fn5983(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6003 github.com/goccy/pythonwasm2go/p0.Fn6003
func Fn6003(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6007 github.com/goccy/pythonwasm2go/p0.Fn6007
func Fn6007(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn6194 github.com/goccy/pythonwasm2go/p0.Fn6194
func Fn6194(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6197 github.com/goccy/pythonwasm2go/p1.Fn6197
func Fn6197(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6213 github.com/goccy/pythonwasm2go/p1.Fn6213
func Fn6213(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6216 github.com/goccy/pythonwasm2go/p1.Fn6216
func Fn6216(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6230 github.com/goccy/pythonwasm2go/p0.Fn6230
func Fn6230(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6232 github.com/goccy/pythonwasm2go/p1.Fn6232
func Fn6232(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6239 github.com/goccy/pythonwasm2go/p0.Fn6239
func Fn6239(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6240 github.com/goccy/pythonwasm2go/p1.Fn6240
func Fn6240(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6241 github.com/goccy/pythonwasm2go/p1.Fn6241
func Fn6241(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6245 github.com/goccy/pythonwasm2go/p1.Fn6245
func Fn6245(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6246 github.com/goccy/pythonwasm2go/p1.Fn6246
func Fn6246(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6250 github.com/goccy/pythonwasm2go/p1.Fn6250
func Fn6250(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6282 github.com/goccy/pythonwasm2go/p1.Fn6282
func Fn6282(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6284 github.com/goccy/pythonwasm2go/p1.Fn6284
func Fn6284(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6287 github.com/goccy/pythonwasm2go/p1.Fn6287
func Fn6287(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6317 github.com/goccy/pythonwasm2go/p1.Fn6317
func Fn6317(m *base.Module, l0 int32)

//go:linkname Fn6319 github.com/goccy/pythonwasm2go/p0.Fn6319
func Fn6319(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn6327 github.com/goccy/pythonwasm2go/p0.Fn6327
func Fn6327(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32)

//go:linkname Fn6342 github.com/goccy/pythonwasm2go/p1.Fn6342
func Fn6342(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6379 github.com/goccy/pythonwasm2go/p1.Fn6379
func Fn6379(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6381 github.com/goccy/pythonwasm2go/p1.Fn6381
func Fn6381(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6383 github.com/goccy/pythonwasm2go/p1.Fn6383
func Fn6383(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6385 github.com/goccy/pythonwasm2go/p1.Fn6385
func Fn6385(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6387 github.com/goccy/pythonwasm2go/p1.Fn6387
func Fn6387(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6389 github.com/goccy/pythonwasm2go/p1.Fn6389
func Fn6389(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6391 github.com/goccy/pythonwasm2go/p1.Fn6391
func Fn6391(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6393 github.com/goccy/pythonwasm2go/p1.Fn6393
func Fn6393(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6395 github.com/goccy/pythonwasm2go/p1.Fn6395
func Fn6395(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6397 github.com/goccy/pythonwasm2go/p1.Fn6397
func Fn6397(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6399 github.com/goccy/pythonwasm2go/p1.Fn6399
func Fn6399(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6401 github.com/goccy/pythonwasm2go/p1.Fn6401
func Fn6401(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6412 github.com/goccy/pythonwasm2go/p1.Fn6412
func Fn6412(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6418 github.com/goccy/pythonwasm2go/p1.Fn6418
func Fn6418(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6431 github.com/goccy/pythonwasm2go/p1.Fn6431
func Fn6431(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn6435 github.com/goccy/pythonwasm2go/p1.Fn6435
func Fn6435(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32

//go:linkname Fn6448 github.com/goccy/pythonwasm2go/p1.Fn6448
func Fn6448(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6450 github.com/goccy/pythonwasm2go/p1.Fn6450
func Fn6450(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6526 github.com/goccy/pythonwasm2go/p1.Fn6526
func Fn6526(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn6571 github.com/goccy/pythonwasm2go/p1.Fn6571
func Fn6571(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn6578 github.com/goccy/pythonwasm2go/p1.Fn6578
func Fn6578(m *base.Module, l0 int32)

//go:linkname Fn6586 github.com/goccy/pythonwasm2go/p1.Fn6586
func Fn6586(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6587 github.com/goccy/pythonwasm2go/p1.Fn6587
func Fn6587(m *base.Module, l0 int32) int32

//go:linkname Fn6589 github.com/goccy/pythonwasm2go/p1.Fn6589
func Fn6589(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6597 github.com/goccy/pythonwasm2go/p1.Fn6597
func Fn6597(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6630 github.com/goccy/pythonwasm2go/p1.Fn6630
func Fn6630(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6631 github.com/goccy/pythonwasm2go/p1.Fn6631
func Fn6631(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6645 github.com/goccy/pythonwasm2go/p1.Fn6645
func Fn6645(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6646 github.com/goccy/pythonwasm2go/p0.Fn6646
func Fn6646(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32

//go:linkname Fn6647 github.com/goccy/pythonwasm2go/p1.Fn6647
func Fn6647(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn6649 github.com/goccy/pythonwasm2go/p1.Fn6649
func Fn6649(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6658 github.com/goccy/pythonwasm2go/p1.Fn6658
func Fn6658(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn6660 github.com/goccy/pythonwasm2go/p1.Fn6660
func Fn6660(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn6671 github.com/goccy/pythonwasm2go/p0.Fn6671
func Fn6671(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn6673 github.com/goccy/pythonwasm2go/p1.Fn6673
func Fn6673(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6686 github.com/goccy/pythonwasm2go/p1.Fn6686
func Fn6686(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn6701 github.com/goccy/pythonwasm2go/p1.Fn6701
func Fn6701(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn6771 github.com/goccy/pythonwasm2go/p1.Fn6771
func Fn6771(m *base.Module, l0 int32) int32

//go:linkname Fn6789 github.com/goccy/pythonwasm2go/p1.Fn6789
func Fn6789(m *base.Module, l0 int32) int32

//go:linkname Fn6795 github.com/goccy/pythonwasm2go/p1.Fn6795
func Fn6795(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6797 github.com/goccy/pythonwasm2go/p1.Fn6797
func Fn6797(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6799 github.com/goccy/pythonwasm2go/p1.Fn6799
func Fn6799(m *base.Module, l0 int32) int32

//go:linkname Fn6809 github.com/goccy/pythonwasm2go/p1.Fn6809
func Fn6809(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn6847 github.com/goccy/pythonwasm2go/p1.Fn6847
func Fn6847(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6860 github.com/goccy/pythonwasm2go/p1.Fn6860
func Fn6860(m *base.Module, l0 int32) int32

//go:linkname Fn6881 github.com/goccy/pythonwasm2go/p1.Fn6881
func Fn6881(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn6883 github.com/goccy/pythonwasm2go/p1.Fn6883
func Fn6883(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6935 github.com/goccy/pythonwasm2go/p1.Fn6935
func Fn6935(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6978 github.com/goccy/pythonwasm2go/p1.Fn6978
func Fn6978(m *base.Module, l0 int32) int32

//go:linkname Fn6989 github.com/goccy/pythonwasm2go/p1.Fn6989
func Fn6989(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6999 github.com/goccy/pythonwasm2go/p1.Fn6999
func Fn6999(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7017 github.com/goccy/pythonwasm2go/p1.Fn7017
func Fn7017(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7039 github.com/goccy/pythonwasm2go/p1.Fn7039
func Fn7039(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7041 github.com/goccy/pythonwasm2go/p1.Fn7041
func Fn7041(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7080 github.com/goccy/pythonwasm2go/p1.Fn7080
func Fn7080(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7114 github.com/goccy/pythonwasm2go/p1.Fn7114
func Fn7114(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7115 github.com/goccy/pythonwasm2go/p1.Fn7115
func Fn7115(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7125 github.com/goccy/pythonwasm2go/p1.Fn7125
func Fn7125(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int64) int32

//go:linkname Fn7189 github.com/goccy/pythonwasm2go/p1.Fn7189
func Fn7189(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7339 github.com/goccy/pythonwasm2go/p1.Fn7339
func Fn7339(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7668 github.com/goccy/pythonwasm2go/p1.Fn7668
func Fn7668(m *base.Module, l0 int32) int32

//go:linkname Fn7718 github.com/goccy/pythonwasm2go/p1.Fn7718
func Fn7718(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8461 github.com/goccy/pythonwasm2go/p0.Fn8461
func Fn8461(m *base.Module, l0 int32)

//go:linkname Fn8463 github.com/goccy/pythonwasm2go/p0.Fn8463
func Fn8463(m *base.Module, l0 int32)

//go:linkname Fn8472 github.com/goccy/pythonwasm2go/p0.Fn8472
func Fn8472(m *base.Module)

//go:linkname Fn8477 github.com/goccy/pythonwasm2go/p1.Fn8477
func Fn8477(m *base.Module, l0 int32) int32

//go:linkname Fn8500 github.com/goccy/pythonwasm2go/p1.Fn8500
func Fn8500(m *base.Module, l0 int32) int32

//go:linkname Fn8503 github.com/goccy/pythonwasm2go/p1.Fn8503
func Fn8503(m *base.Module, l0 int32) int32

//go:linkname Fn8506 github.com/goccy/pythonwasm2go/p1.Fn8506
func Fn8506(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8558 github.com/goccy/pythonwasm2go/p1.Fn8558
func Fn8558(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8596 github.com/goccy/pythonwasm2go/p1.Fn8596
func Fn8596(m *base.Module) int32

//go:linkname Fn8614 github.com/goccy/pythonwasm2go/p1.Fn8614
func Fn8614(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8617 github.com/goccy/pythonwasm2go/p1.Fn8617
func Fn8617(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8619 github.com/goccy/pythonwasm2go/p1.Fn8619
func Fn8619(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8625 github.com/goccy/pythonwasm2go/p0.Fn8625
func Fn8625(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8626 github.com/goccy/pythonwasm2go/p0.Fn8626
func Fn8626(m *base.Module) int32

//go:linkname Fn8671 github.com/goccy/pythonwasm2go/p1.Fn8671
func Fn8671(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8678 github.com/goccy/pythonwasm2go/p1.Fn8678
func Fn8678(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8695 github.com/goccy/pythonwasm2go/p1.Fn8695
func Fn8695(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8699 github.com/goccy/pythonwasm2go/p1.Fn8699
func Fn8699(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8701 github.com/goccy/pythonwasm2go/p1.Fn8701
func Fn8701(m *base.Module, l0 int32) int32

//go:linkname Fn8737 github.com/goccy/pythonwasm2go/p1.Fn8737
func Fn8737(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8755 github.com/goccy/pythonwasm2go/p1.Fn8755
func Fn8755(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8781 github.com/goccy/pythonwasm2go/p1.Fn8781
func Fn8781(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn8836 github.com/goccy/pythonwasm2go/p1.Fn8836
func Fn8836(m *base.Module, l0 int32) int32

//go:linkname Fn8847 github.com/goccy/pythonwasm2go/p1.Fn8847
func Fn8847(m *base.Module, l0 int32)

//go:linkname Fn8903 github.com/goccy/pythonwasm2go/p1.Fn8903
func Fn8903(m *base.Module, l0 int32) int32

//go:linkname Fn8908 github.com/goccy/pythonwasm2go/p1.Fn8908
func Fn8908(m *base.Module, l0 int32) int32

//go:linkname Fn8921 github.com/goccy/pythonwasm2go/p1.Fn8921
func Fn8921(m *base.Module, l0 int32) int32

//go:linkname Fn8923 github.com/goccy/pythonwasm2go/p1.Fn8923
func Fn8923(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8947 github.com/goccy/pythonwasm2go/p1.Fn8947
func Fn8947(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8988 github.com/goccy/pythonwasm2go/p1.Fn8988
func Fn8988(m *base.Module, l0 int32) int32

//go:linkname Fn9039 github.com/goccy/pythonwasm2go/p1.Fn9039
func Fn9039(m *base.Module, l0 int32) int32

//go:linkname Fn9077 github.com/goccy/pythonwasm2go/p1.Fn9077
func Fn9077(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9108 github.com/goccy/pythonwasm2go/p1.Fn9108
func Fn9108(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9116 github.com/goccy/pythonwasm2go/p1.Fn9116
func Fn9116(m *base.Module, l0 int32, l1 int64, l2 int32) int64

//go:linkname Fn9117 github.com/goccy/pythonwasm2go/p1.Fn9117
func Fn9117(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9118 github.com/goccy/pythonwasm2go/p1.Fn9118
func Fn9118(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9165 github.com/goccy/pythonwasm2go/p1.Fn9165
func Fn9165(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9167 github.com/goccy/pythonwasm2go/p1.Fn9167
func Fn9167(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn9170 github.com/goccy/pythonwasm2go/p1.Fn9170
func Fn9170(m *base.Module, l0 int32) int32

//go:linkname Fn9174 github.com/goccy/pythonwasm2go/p1.Fn9174
func Fn9174(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9192 github.com/goccy/pythonwasm2go/p1.Fn9192
func Fn9192(m *base.Module, l0 int32) int32

//go:linkname Fn9238 github.com/goccy/pythonwasm2go/p1.Fn9238
func Fn9238(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9258 github.com/goccy/pythonwasm2go/p1.Fn9258
func Fn9258(m *base.Module, l0 int32) int32

//go:linkname Fn9361 github.com/goccy/pythonwasm2go/p1.Fn9361
func Fn9361(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn9363 github.com/goccy/pythonwasm2go/p1.Fn9363
func Fn9363(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9367 github.com/goccy/pythonwasm2go/p0.Fn9367
func Fn9367(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9369 github.com/goccy/pythonwasm2go/p1.Fn9369
func Fn9369(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn9380 github.com/goccy/pythonwasm2go/p0.Fn9380
func Fn9380(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9381 github.com/goccy/pythonwasm2go/p0.Fn9381
func Fn9381(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9382 github.com/goccy/pythonwasm2go/p0.Fn9382
func Fn9382(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9395 github.com/goccy/pythonwasm2go/p1.Fn9395
func Fn9395(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9396 github.com/goccy/pythonwasm2go/p1.Fn9396
func Fn9396(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9427 github.com/goccy/pythonwasm2go/p1.Fn9427
func Fn9427(m *base.Module, l0 int32) int32

//go:linkname Fn9471 github.com/goccy/pythonwasm2go/p1.Fn9471
func Fn9471(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9486 github.com/goccy/pythonwasm2go/p1.Fn9486
func Fn9486(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9505 github.com/goccy/pythonwasm2go/p1.Fn9505
func Fn9505(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9508 github.com/goccy/pythonwasm2go/p1.Fn9508
func Fn9508(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9509 github.com/goccy/pythonwasm2go/p1.Fn9509
func Fn9509(m *base.Module, l0 int32) int32

//go:linkname Fn9552 github.com/goccy/pythonwasm2go/p1.Fn9552
func Fn9552(m *base.Module, l0 int32) int32

//go:linkname Fn9573 github.com/goccy/pythonwasm2go/p1.Fn9573
func Fn9573(m *base.Module, l0 int32) int32

//go:linkname Fn9691 github.com/goccy/pythonwasm2go/p1.Fn9691
func Fn9691(m *base.Module, l0 int32)

//go:linkname Fn9744 github.com/goccy/pythonwasm2go/p1.Fn9744
func Fn9744(m *base.Module, l0 int32, l1 int64, l2 int32)

//go:linkname Fn9746 github.com/goccy/pythonwasm2go/p1.Fn9746
func Fn9746(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9774 github.com/goccy/pythonwasm2go/p1.Fn9774
func Fn9774(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9775 github.com/goccy/pythonwasm2go/p1.Fn9775
func Fn9775(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9781 github.com/goccy/pythonwasm2go/p1.Fn9781
func Fn9781(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn9801 github.com/goccy/pythonwasm2go/p1.Fn9801
func Fn9801(m *base.Module)

//go:linkname Fn9806 github.com/goccy/pythonwasm2go/p1.Fn9806
func Fn9806(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9811 github.com/goccy/pythonwasm2go/p1.Fn9811
func Fn9811(m *base.Module, l0 int32)

//go:linkname Fn9815 github.com/goccy/pythonwasm2go/p1.Fn9815
func Fn9815(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9820 github.com/goccy/pythonwasm2go/p1.Fn9820
func Fn9820(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn9834 github.com/goccy/pythonwasm2go/p1.Fn9834
func Fn9834(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9836 github.com/goccy/pythonwasm2go/p1.Fn9836
func Fn9836(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9843 github.com/goccy/pythonwasm2go/p1.Fn9843
func Fn9843(m *base.Module, l0 float64, l1 float64) float64

//go:linkname Fn9847 github.com/goccy/pythonwasm2go/p0.Fn9847
func Fn9847(m *base.Module, l0 float64, l1 int32) int32

//go:linkname Fn9858 github.com/goccy/pythonwasm2go/p1.Fn9858
func Fn9858(m *base.Module, l0 float64) float64

//go:linkname Fn9860 github.com/goccy/pythonwasm2go/p1.Fn9860
func Fn9860(m *base.Module, l0 float64) float64

//go:linkname Fn9862 github.com/goccy/pythonwasm2go/p1.Fn9862
func Fn9862(m *base.Module, l0 float64, l1 float64) float64

//go:linkname Fn9864 github.com/goccy/pythonwasm2go/p1.Fn9864
func Fn9864(m *base.Module, l0 float64, l1 float64) float64

//go:linkname Fn9867 github.com/goccy/pythonwasm2go/p1.Fn9867
func Fn9867(m *base.Module, l0 float64) float64

//go:linkname Fn9868 github.com/goccy/pythonwasm2go/p1.Fn9868
func Fn9868(m *base.Module, l0 float64) float64

//go:linkname Fn9869 github.com/goccy/pythonwasm2go/p1.Fn9869
func Fn9869(m *base.Module, l0 float64) float64

//go:linkname Fn9872 github.com/goccy/pythonwasm2go/p1.Fn9872
func Fn9872(m *base.Module, l0 float64, l1 float64) float64

//go:linkname Fn9878 github.com/goccy/pythonwasm2go/p1.Fn9878
func Fn9878(m *base.Module, l0 float64, l1 float64, l2 int32) float64

//go:linkname Fn9881 github.com/goccy/pythonwasm2go/p1.Fn9881
func Fn9881(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9884 github.com/goccy/pythonwasm2go/p1.Fn9884
func Fn9884(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9924 github.com/goccy/pythonwasm2go/p1.Fn9924
func Fn9924(m *base.Module, l0 int32, l1 int32, l2 int64) int64

//go:linkname Fn9929 github.com/goccy/pythonwasm2go/p1.Fn9929
func Fn9929(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9933 github.com/goccy/pythonwasm2go/p1.Fn9933
func Fn9933(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9935 github.com/goccy/pythonwasm2go/p1.Fn9935
func Fn9935(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9940 github.com/goccy/pythonwasm2go/p1.Fn9940
func Fn9940(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9942 github.com/goccy/pythonwasm2go/p1.Fn9942
func Fn9942(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9959 github.com/goccy/pythonwasm2go/p1.Fn9959
func Fn9959(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9963 github.com/goccy/pythonwasm2go/p0.Fn9963
func Fn9963(m *base.Module, l0 int32) int32

//go:linkname Fn9964 github.com/goccy/pythonwasm2go/p1.Fn9964
func Fn9964(m *base.Module, l0 int32)

//go:linkname Fn9966 github.com/goccy/pythonwasm2go/p1.Fn9966
func Fn9966(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9980 github.com/goccy/pythonwasm2go/p1.Fn9980
func Fn9980(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn9981 github.com/goccy/pythonwasm2go/p1.Fn9981
func Fn9981(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn9992 github.com/goccy/pythonwasm2go/p1.Fn9992
func Fn9992(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9994 github.com/goccy/pythonwasm2go/p1.Fn9994
func Fn9994(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10001 github.com/goccy/pythonwasm2go/p1.Fn10001
func Fn10001(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10004 github.com/goccy/pythonwasm2go/p1.Fn10004
func Fn10004(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10024 github.com/goccy/pythonwasm2go/p0.Fn10024
func Fn10024(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn10043 github.com/goccy/pythonwasm2go/p0.Fn10043
func Fn10043(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10053 github.com/goccy/pythonwasm2go/p0.Fn10053
func Fn10053(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10055 github.com/goccy/pythonwasm2go/p1.Fn10055
func Fn10055(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn10059 github.com/goccy/pythonwasm2go/p1.Fn10059
func Fn10059(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn10060 github.com/goccy/pythonwasm2go/p1.Fn10060
func Fn10060(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn10071 github.com/goccy/pythonwasm2go/p1.Fn10071
func Fn10071(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn10072 github.com/goccy/pythonwasm2go/p1.Fn10072
func Fn10072(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn10089 github.com/goccy/pythonwasm2go/p1.Fn10089
func Fn10089(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn10098 github.com/goccy/pythonwasm2go/p0.Fn10098
func Fn10098(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10107 github.com/goccy/pythonwasm2go/p0.Fn10107
func Fn10107(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10110 github.com/goccy/pythonwasm2go/p1.Fn10110
func Fn10110(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10111 github.com/goccy/pythonwasm2go/p0.Fn10111
func Fn10111(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10118 github.com/goccy/pythonwasm2go/p1.Fn10118
func Fn10118(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10148 github.com/goccy/pythonwasm2go/p0.Fn10148
func Fn10148(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn10168 github.com/goccy/pythonwasm2go/p1.Fn10168
func Fn10168(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10177 github.com/goccy/pythonwasm2go/p1.Fn10177
func Fn10177(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn10201 github.com/goccy/pythonwasm2go/p1.Fn10201
func Fn10201(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32

//go:linkname Fn10202 github.com/goccy/pythonwasm2go/p1.Fn10202
func Fn10202(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32

//go:linkname Fn10203 github.com/goccy/pythonwasm2go/p1.Fn10203
func Fn10203(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32

//go:linkname Fn10204 github.com/goccy/pythonwasm2go/p1.Fn10204
func Fn10204(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32

//go:linkname Fn10219 github.com/goccy/pythonwasm2go/p1.Fn10219
func Fn10219(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
