package factory

import (
	"fmt"
	"ktmall/app/models"
)

func factoryGoodsInfo() {
	dropAndCreateTable(&models.GoodsInfo{})

	goods := []*models.GoodsInfo{
		{CategoryId: 15, Desc: "Apple MacBook Air 13.3英寸笔记本电脑 银色(Core i5 处理器/8GB内存/128GB SSD闪存 MMGF2CH/A)", DefaultIcon: "https://img11.360buyimg.com/n7/jfs/t2968/143/2485546147/238650/70df281e/57b12a31N8f4f75a3.jpg", DefaultPrice: "1", Banner: "https://img11.360buyimg.com/n1/s450x450_jfs/t2968/143/2485546147/238650/70df281e/57b12a31N8f4f75a3.jpg,https://img13.360buyimg.com/n1/s450x450_jfs/t2632/157/4193453761/92922/2adb5ebc/57ad88f0Nb286ec7a.jpg,https://img11.360buyimg.com/n1/s450x450_jfs/t2977/86/2412624329/68019/dbe32c1f/57ad8846N64ac3c79.jpg", DetailOne: "https://img20.360buyimg.com/vc/jfs/t3034/151/748569500/226790/d6cd86a2/57b15612N81dc489d.jpg", DetailTwo: "https://img20.360buyimg.com/vc/jfs/t2683/60/4222930118/169462/233c7678/57b15616N1e285f09.jpg", SalesCount: 5000, StockCount: 10000, Code: "10000000001", DefaultSku: "1.6GHz i5处理器,GB内存/128GB SSD,1件"},
		{CategoryId: 15, Desc: "Apple MacBook Pro 13.3英寸笔记本电脑 银色(Core i5 处理器/8GB内存/128GB SSD闪存/Retina屏 MF839CH/A)", DefaultIcon: "https://img13.360buyimg.com/n7/jfs/t2590/338/4222589387/237425/25e40fb/57b12ac6N61374a75.jpg", DefaultPrice: "2", Banner: "https://img13.360buyimg.com/n1/s450x450_jfs/t2590/338/4222589387/237425/25e40fb/57b12ac6N61374a75.jpg,https://img13.360buyimg.com/n1/s450x450_jfs/t2632/157/4193453761/92922/2adb5ebc/57ad88f0Nb286ec7a.jpg,https://img11.360buyimg.com/n1/s450x450_jfs/t2977/86/2412624329/68019/dbe32c1f/57ad8846N64ac3c79.jpg", DetailOne: "https://img20.360buyimg.com/vc/jfs/t2218/160/1600609450/672652/12ef3837/56600724N28581935.jpg", DetailTwo: "https://img20.360buyimg.com/vc/jfs/t2305/205/1525153851/648458/79706530/56600727N1022ecd8.jpg", SalesCount: 13530, StockCount: 800, Code: "10000000002", DefaultSku: "13.3英寸/双核 i5/8G/128G闪存,1件"},
		{CategoryId: 15, Desc: "Apple MacBook Pro 15.4英寸笔记本电脑 深空灰色（Multi-Touch Bar/Core i7/16GB/256GB MLH32CH/A）", DefaultIcon: "https://img10.360buyimg.com/n7/jfs/t3499/165/739574790/179345/251c51d4/58126465Na27a9bf0.jpg", DefaultPrice: "1", Banner: "https://img13.360buyimg.com/n1/s450x450_jfs/t2590/338/4222589387/237425/25e40fb/57b12ac6N61374a75.jpg,https://img13.360buyimg.com/n1/s450x450_jfs/t2632/157/4193453761/92922/2adb5ebc/57ad88f0Nb286ec7a.jpg,https://img11.360buyimg.com/n1/s450x450_jfs/t2977/86/2412624329/68019/dbe32c1f/57ad8846N64ac3c79.jpg", DetailOne: "https://img12.360buyimg.com/cms/jfs/t3760/42/740411238/146541/be42da9e/581266b4N7360a7e4.jpg", DetailTwo: "https://img13.360buyimg.com/cms/jfs/t3367/226/740508746/107595/414c65e6/581266b7Nea97b88f.jpg", SalesCount: 1350, StockCount: 8200, Code: "10000000003", DefaultSku: "银色,Core i5/8G内存/256G闪存,1件"},
		{CategoryId: 15, Desc: "Apple 苹果 MacBook Air 13.3英寸笔记本电脑 MMGF2CH/A银色 Core i5/8G内存/128G闪存", DefaultIcon: "https://img13.360buyimg.com/n1/jfs/t5821/151/4295371754/246251/29179c11/594a3eabNbf7fceec.jpg", DefaultPrice: "3.00", Banner: "https://img13.360buyimg.com/n1/jfs/t5821/151/4295371754/246251/29179c11/594a3eabNbf7fceec.jpg,https://img13.360buyimg.com/n1/jfs/t3157/218/9378001544/103166/c8bb88da/58d0e987Nb3a34bfa.jpg,https://img13.360buyimg.com/n1/jfs/t4639/313/514186457/88273/6d83eae9/58d0e988N0009a10c.jpg", DetailOne: "https://img30.360buyimg.com/popWaterMark/jfs/t5746/329/4801800135/218275/4a67f0ef/5954a40fNf471b55d.jpg", DetailTwo: "https://img30.360buyimg.com/popWareDetail/jfs/t6532/96/2593310/189936/1c5b2d94/5937a66aN594f4851.jpg", SalesCount: 2350, StockCount: 3400, Code: "10000000005", DefaultSku: "MMGF2CH/A银色,Core i5/8G内存/128G闪存,1件"},
		{CategoryId: 15, Desc: "Apple 苹果 MacBook Pro 笔记本电脑 16年新款 15英寸 Multi-Touch Bar 256G 深空灰色", DefaultIcon: "https://img14.360buyimg.com/n5/s450x450_jfs/t3397/55/762020838/184157/7e507d32/58131c17Nb108ca54.jpg", DefaultPrice: "8.00", Banner: "https://img14.360buyimg.com/n5/s450x450_jfs/t3397/55/762020838/184157/7e507d32/58131c17Nb108ca54.jpg,https://img14.360buyimg.com/n5/s450x450_jfs/t3427/1/761535388/238533/b605b9f2/58130206Ncf90e695.jpg,https://img14.360buyimg.com/n5/s450x450_jfs/t3490/164/628752282/118644/88a1c360/58130288Ne2b8683e.jpg,https://img14.360buyimg.com/n5/s450x450_jfs/t3715/86/798415057/312693/b777b279/58131c16Ncba81f8b.jpg", DetailOne: "https://img10.360buyimg.com/imgzone/jfs/t3082/24/2402653567/223529/1a6b04b/57e0d2f8N10eb5602.jpg", DetailTwo: "https://img30.360buyimg.com/popWaterMark/jfs/t6241/302/1026912919/130228/685bcc83/5949c952N88858337.jpg", SalesCount: 150, StockCount: 900, Code: "10000000006", DefaultSku: "13英寸 Core i5 8G内存 256G闪存,深空灰色,1件"},
		{CategoryId: 15, Desc: "Apple MacBook Air 13.3英寸笔记本电脑 银色(Core i5 处理器/8GB内存/128GB SSD闪存 MMGF2CH/A)", DefaultIcon: "https://img11.360buyimg.com/n7/jfs/t2968/143/2485546147/238650/70df281e/57b12a31N8f4f75a3.jpg", DefaultPrice: "1", Banner: "https://img11.360buyimg.com/n1/s450x450_jfs/t2968/143/2485546147/238650/70df281e/57b12a31N8f4f75a3.jpg,https://img13.360buyimg.com/n1/s450x450_jfs/t2632/157/4193453761/92922/2adb5ebc/57ad88f0Nb286ec7a.jpg,https://img11.360buyimg.com/n1/s450x450_jfs/t2977/86/2412624329/68019/dbe32c1f/57ad8846N64ac3c79.jpg", DetailOne: "https://img20.360buyimg.com/vc/jfs/t3034/151/748569500/226790/d6cd86a2/57b15612N81dc489d.jpg", DetailTwo: "https://img20.360buyimg.com/vc/jfs/t2683/60/4222930118/169462/233c7678/57b15616N1e285f09.jpg", SalesCount: 5000, StockCount: 10000, Code: "10000000001", DefaultSku: "1.6GHz i5处理器,GB内存/128GB SSD,1件"},
		{CategoryId: 15, Desc: "Apple 苹果 MacBook Air 13.3英寸笔记本电脑 MMGF2CH/A银色 Core i5/8G内存/128G闪存", DefaultIcon: "https://img13.360buyimg.com/n1/jfs/t5821/151/4295371754/246251/29179c11/594a3eabNbf7fceec.jpg", DefaultPrice: "3.00", Banner: "https://img13.360buyimg.com/n1/jfs/t5821/151/4295371754/246251/29179c11/594a3eabNbf7fceec.jpg,https://img13.360buyimg.com/n1/jfs/t3157/218/9378001544/103166/c8bb88da/58d0e987Nb3a34bfa.jpg,https://img13.360buyimg.com/n1/jfs/t4639/313/514186457/88273/6d83eae9/58d0e988N0009a10c.jpg", DetailOne: "https://img30.360buyimg.com/popWaterMark/jfs/t5746/329/4801800135/218275/4a67f0ef/5954a40fNf471b55d.jpg", DetailTwo: "https://img30.360buyimg.com/popWareDetail/jfs/t6532/96/2593310/189936/1c5b2d94/5937a66aN594f4851.jpg", SalesCount: 2350, StockCount: 3400, Code: "10000000005", DefaultSku: "MMGF2CH/A银色,Core i5/8G内存/128G闪存,1件"},
		{CategoryId: 15, Desc: "国行Apple/苹果 MacBook Pro MF839CH/A 13.3英寸 商务笔记本电脑", DefaultIcon: "https://img.alicdn.com/imgextra/i3/1669409267/TB2bMSccXXXXXbVXXXXXXXXXXXX_!!1669409267.jpg_430x430q90.jpg", DefaultPrice: "5.00", Banner: "https://img.alicdn.com/imgextra/i3/1669409267/TB2bMSccXXXXXbVXXXXXXXXXXXX_!!1669409267.jpg_430x430q90.jpg,https://img.alicdn.com/imgextra/i2/1669409267/TB2WUa4qpXXXXaQXpXXXXXXXXXX_!!1669409267.jpg_430x430q90.jpg,https://img.alicdn.com/imgextra/i2/1669409267/TB2ZnUysXXXXXXOXXXXXXXXXXXX_!!1669409267.jpg_430x430q90.jpg", DetailOne: "https://gdp.alicdn.com/imgextra/i1/1669409267/TB25Yfnmr_0UKFjy1XaXXbKfXXa_!!1669409267.jpg", DetailTwo: "https://img.alicdn.com/imgextra/i4/1669409267/TB28SC_vdFopuFjSZFHXXbSlXXa_!!1669409267.jpg", SalesCount: 1660, StockCount: 1200, Code: "10000000007", DefaultSku: "8G 128G,银白色,1件"},
		{CategoryId: 25, Desc: "Apple iPhone 6s Plus (A1699) 32G 金色 移动联通电信4G手机", DefaultIcon: "https://img14.360buyimg.com/n1/s450x450_jfs/t3268/124/2646283367/114153/f5704b88/57e4a358N9ccc6645.jpg", DefaultPrice: "1.00", Banner: "https://img14.360buyimg.com/n1/s450x450_jfs/t3268/124/2646283367/114153/f5704b88/57e4a358N9ccc6645.jpg,https://img14.360buyimg.com/n1/s450x450_jfs/t3055/170/2635484631/101619/8000b30/57e4a359N218a4200.jpg,https://img12.360buyimg.com/n1/s450x450_jfs/t3274/10/2586151054/109903/711d3edd/57e4a35aNc4a0b203.jpg", DetailOne: "https://img30.360buyimg.com/jgsq-productsoa/jfs/t6337/310/2148869366/61744/dca36a9c/595dda76N64984138.jpg", DetailTwo: "https://img30.360buyimg.com/jgsq-productsoa/jfs/t6175/244/140312464/88326/30e3b943/593a4888N3187bea7.jpg", SalesCount: 9890, StockCount: 1200, Code: "10000000010", DefaultSku: "金色,32GB,1件"},
		{CategoryId: 25, Desc: "Apple iPhone 7 Plus (A1661) 32G 亮黑色 移动联通电信4G手机", DefaultIcon: "https://img14.360buyimg.com/n1/s450x450_jfs/t3064/188/1693292264/115570/e891640b/57d11bfaN2e8acade.jpg", DefaultPrice: "2.00", Banner: "https://img14.360buyimg.com/n1/s450x450_jfs/t3064/188/1693292264/115570/e891640b/57d11bfaN2e8acade.jpg,https://img14.360buyimg.com/n1/s450x450_jfs/t3202/139/1655646299/94486/f66e4dae/57d11bfbN147f8c9a.jpg,https://img14.360buyimg.com/n1/s450x450_jfs/t2998/150/2060445093/29614/e9c3eab5/57d11bfbN5a03e02d.jpg", DetailOne: "https://img30.360buyimg.com/jgsq-productsoa/jfs/t6337/310/2148869366/61744/dca36a9c/595dda76N64984138.jpg", DetailTwo: "https://img30.360buyimg.com/jgsq-productsoa/jfs/t6175/244/140312464/88326/30e3b943/593a4888N3187bea7.jpg", SalesCount: 12313, StockCount: 12566, Code: "10000000011", DefaultSku: "亮黑色,32GB,1件"},
		{CategoryId: 25, Desc: "苹果 Apple iPhone 6 4G手机 灰色 公开全网通(32G)标配", DefaultIcon: "https://img14.360buyimg.com/n1/s450x450_jfs/t5581/37/6255093126/256134/d8ae5c1d/59688753N177cfc26.jpg", DefaultPrice: "1.00", Banner: "https://img14.360buyimg.com/n1/s450x450_jfs/t5581/37/6255093126/256134/d8ae5c1d/59688753N177cfc26.jpg,https://img14.360buyimg.com/n1/s450x450_jfs/t5593/357/6209896656/70558/1cc69840/59688755Nd4a0c527.jpg,https://img14.360buyimg.com/n1/s450x450_jfs/t5848/309/6166315010/59102/2a85c081/59688756Ndce52d99.jpg", DetailOne: "https://img30.360buyimg.com/popWaterMark/jfs/t3016/136/650115357/171206/c66b2965/57ad8acaN57ed9319.jpg", DetailTwo: "https://img30.360buyimg.com/popWaterMark/jfs/t2935/12/2469495182/65696/a31dc72e/57ad8acaNef46eb26.jpg", SalesCount: 56667, StockCount: 8434, Code: "10000000012", DefaultSku: "灰色,32GB,1件"},
		{CategoryId: 16, Desc: "联想（ThinkPad）E460（20ETA063CD）14英寸笔记本电脑（i7-6498DU 8G 1T 2G独显 Win10）", DefaultIcon: "https://img14.360buyimg.com/n1/s450x450_jfs/t5446/184/2581563799/286954/9ebfad09/591bddbcNc724d048.jpg", DefaultPrice: "1.00", Banner: "https://img14.360buyimg.com/n1/s450x450_jfs/t5446/184/2581563799/286954/9ebfad09/591bddbcNc724d048.jpg,https://img14.360buyimg.com/n1/s450x450_jfs/t5209/149/1168851360/329572/9bdcbaa5/590c421eN49bc2d11.jpg,https://img14.360buyimg.com/n1/s450x450_jfs/t5437/240/2545062154/189407/7ab57b0b/591bddc2N6c52f1ad.jpg,https://img14.360buyimg.com/n1/s450x450_jfs/t5326/216/1147917773/304909/e711dfc6/590c421fNc5def6a7.jpg", DetailOne: "https://img20.360buyimg.com/vc/jfs/t5563/279/850602969/1094377/272f0258/5907fc23N3a312831.jpg", DetailTwo: "https://img20.360buyimg.com/vc/jfs/t5563/279/850602969/1094377/272f0258/5907fc23N3a312831.jpg", SalesCount: 56667, StockCount: 8434, Code: "10000000015", DefaultSku: "【i7豪华版】8G 1T,1件"},
		{CategoryId: 16, Desc: "ThinkPad New S2（20GU0000CD）13.3英寸轻薄笔记本电脑（i5-6200U 4G 240GB SSD FHD IPS Win10 银色）", DefaultIcon: "https://img12.360buyimg.com/n1/s450x450_jfs/t3439/192/665209614/82997/ce6c3ad8/58107973Na856a8e1.jpg", DefaultPrice: "2.00", Banner: "https://img12.360buyimg.com/n1/s450x450_jfs/t3439/192/665209614/82997/ce6c3ad8/58107973Na856a8e1.jpg,https://img12.360buyimg.com/n1/s450x450_jfs/t3313/27/610873927/76897/44bc7560/58107978N6e64b6ce.jpg,https://img12.360buyimg.com/n1/s450x450_jfs/t3766/313/640498039/58537/b7e06633/5810797dN0ec052bd.jpg,https://img12.360buyimg.com/n1/s450x450_jfs/t3379/29/613387717/62645/cad23f55/58107982Neca08199.jpg", DetailOne: "https://img20.360buyimg.com/vc/jfs/t3172/316/316003976/446142/2e681194/57b168fcNfc72b6c2.jpg", DetailTwo: "https://img20.360buyimg.com/vc/jfs/t3172/316/316003976/446142/2e681194/57b168fcNfc72b6c2.jpg", SalesCount: 53453, StockCount: 5623, Code: "10000000016", DefaultSku: "【轻薄特价】180打开i5标配,1件"},
		{CategoryId: 16, Desc: "联想（ThinkPad）黑侠E570 GTX（1NCD）游戏笔记本（i5-7200U 4G 500G+128G SSD GTX950M 2G独显 FHD Win10）", DefaultIcon: "https://img13.360buyimg.com/n1/s450x450_jfs/t3589/5/1216566713/222593/43da0a93/5822ed83N8f0984b9.jpg", DefaultPrice: "3.00", Banner: "https://img13.360buyimg.com/n1/s450x450_jfs/t3589/5/1216566713/222593/43da0a93/5822ed83N8f0984b9.jpg,https://img13.360buyimg.com/n1/s450x450_jfs/t3514/49/1286403745/215402/2c20d256/5822ed87N71a050ee.jpg,https://img13.360buyimg.com/n1/s450x450_jfs/t3406/2/1342234371/410922/dbedf7f8/5822ed8bN792d72cd.jpg,https://img13.360buyimg.com/n1/s450x450_jfs/t3496/362/1334197333/221885/678bc709/5822ed93Nd409d7f0.jpg", DetailOne: "https://img20.360buyimg.com/vc/jfs/t3850/230/1078650780/1998827/c8083d77/5822d122N3d828e6d.jpg", DetailTwo: "https://img20.360buyimg.com/vc/jfs/t3850/230/1078650780/1998827/c8083d77/5822d122N3d828e6d.jpg", SalesCount: 32423, StockCount: 7345, Code: "10000000017", DefaultSku: "【黑侠游戏本特价款】GTX950,1件"},
		{CategoryId: 16, Desc: "ThinkPad X1 Carbon 2017（20HRA007CD）14英寸轻薄笔记本电脑（i5-7200U 8G 256GSSD FHD Win10）", DefaultIcon: "https://img11.360buyimg.com/n1/s450x450_jfs/t4804/97/1759098535/205293/61ec47d0/58f428f6N5b35ea21.jpg", DefaultPrice: "3.00", Banner: "https://img11.360buyimg.com/n1/s450x450_jfs/t4804/97/1759098535/205293/61ec47d0/58f428f6N5b35ea21.jpg,https://img11.360buyimg.com/n1/s450x450_jfs/t4753/151/1078896686/306158/d1133be7/58ec5f8aN9074a72b.jpg,https://img11.360buyimg.com/n1/s450x450_jfs/t5065/111/1077695223/143927/6b0de1f/58ec5f90N0954eca1.jpg,https://img11.360buyimg.com/n1/s450x450_jfs/t4618/299/2175767160/96803/f14f7185/58ec5f94Nc2fbb336.jpg", DetailOne: "https://img20.360buyimg.com/vc/jfs/t3850/230/1078650780/1998827/c8083d77/5822d122N3d828e6d.jpg", DetailTwo: "https://img20.360buyimg.com/vc/jfs/t3850/230/1078650780/1998827/c8083d77/5822d122N3d828e6d.jpg", SalesCount: 9894, StockCount: 10000, Code: "10000000018", DefaultSku: "2017新X1 8G 256G,1件"},
	}

	for _, g := range goods {
		if err := models.DB().Create(&g).Error; err != nil {
			fmt.Printf("mock goods_info error： %v\n", err)
		}
	}
}
