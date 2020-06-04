package model

type Product struct {
	Id int
	Category
	SubProducts []SubProduct
	Sizes       []Size
}

type Category struct {
	Id          int
	Name        string
	Description string
}

type SubProduct struct {
	Id             int
	Name           string
	Article        string
	Price          int
	PhotoLinkShort string
	PhotoLinkLong  string
}

type Size struct {
	Id   int
	Size string
}

// categories
var kidsCategory = Category{1, "Кидсы", "<br>Удобные и модные детские Кидсы <br><br>Легкие, красивые, яркие <br><br>Отличное качество<br><br>Ваш малыш будет в восторге <br><br>Материал: полимер на основе мягкого  пластика <br><br>Страна производитель: Гонконг <br><br>В наличии все размеры: <br><br>24 размер, 2,5 года, стопа 14,6 см<br>25 размер, 2,5 — 3 года, стопа 15,3 см<br>26 размер, 3 — 4 года, стопа 16,5 см <br>27 размер, 4 — 4,5 года, стопа 17,3 см <br>28 размер, 4,5 — 5 лет, стопа 17,6 см <br>29 размер, 5-6 лет, стопа 18,5 см<br><br>Доставка: <br><br>— самовывоз по адресу г.Киев, пр. Бажана 10 (метро Осокорки) — бесплатно <br>— доставка курьером до любой станции метро (г. Киев) — 20 грн <br>— доставка курьером по указанному адресу (г.Киев) — 30 грн <br>— доставка Укрпочтой или Новой Почтой в любой город Украины — 25 грн <br><br>Оплата: <br><br>— наличными курьеру при получении товара <br>— предоплата на карточку Приватбанка <br>"}
var pajamasCategory = Category{2, "Пижамы", "<br>Детская пижама с длинным рукавом + длинные брюки (комплект) <br><br>Легкая, теплая, уютная, красивая и яркая<br><br>Отличное качество<br><br>Ваш малыш будет в восторге<br><br>Материал: 100% хлопок <br><br>Страна производитель: Гонконг <br><br>В наличии все размеры: <br><br>Размер 2Т (2 года) — Рост- 90 см<br>Размер 3Т (3 года) — Рост — 95 см<br>Размер 4т (4 года) — Рост — 100 см<br>Размер 5т (5 лет) — Рост — 110 см<br>Размер 6т (6 лет) — Рост — 120 см<br>Размер-7т (7 лет) — Рост — 130 см<br><br><div class = \\\"hidden-xs\\\"><br><img src=\"/assets/img/sizes/ironman.png\"><br><img src=\"/assets/img/sizes/superman.png\"><br><img src=\"/assets/img/sizes/spiderman.png\"><br><img src=\"/assets/img/sizes/hellokitty.png\"><br><br></div>Доставка:<br><br>— самовывоз по адресу г.Киев, пр. Бажана 10 (метро Осокорки) — бесплатно <br>— доставка курьером до любой станции метро (г. Киев) — 20 грн <br>— доставка курьером по указанному адресу (г.Киев) — 30 грн<br>— доставка Укрпочтой или Новой Почтой в любой город Украины — 25 грн<br><br>Оплата:<br><br>— наличными курьеру при получении товара<br>— предоплата на карточку Приватбанка<br>"}
var pantiesCategory = Category{3, "Трусики", "<br>Детские трусики-боксеры для мальчиков<br><br>В серии представлены трусы с различными принтами: Angry birds, Cars, Spiderman, SpongeBob, Mickey Mouse, Doreamon, Thomas<br><br>Задняя сторона трусиков однотонная<br><br>Ваш малыш будет в восторге<br><br>Отличное качество<br><br>Материал: 100% хлопок<br><br>Страна производитель: Гонконг<br><br>В наличии размеры:<br><br>Размер S, рост 100-110 см, талия 22 см * 2, 4-6 лет<br>Размер M, рост 110-120 см, талия 24 см * 2, 7-8 лет<br>Размер L, рост 120-130 см, талия 26 см * 2, 9-10 лет<br><br>Доставка:<br><br>— самовывоз по адресу г.Киев, пр. Бажана 10 (метро Осокорки) — бесплатно<br>— доставка курьером до любой станции метро (г. Киев) — 20 грн<br>— доставка курьером по указанному адресу (г.Киев) — 30 грн<br>— доставка Укрпочтой или Новой Почтой в любой город Украины — 25 грн<br><br>Оплата:<br><br>— наличными курьеру при получении товара<br>— предоплата на карточку Приватбанка<br>"}

// sub products
var blueKids = SubProduct{1, "Купить Кидсы (цвет голубой)", "300101", 415, "img/kids/preview/01blue185X185.png", "img/kids/01blue500X500.jpg"}
var whitePinkKids = SubProduct{2, "Купить Кидсы (цвет бело-розовый)", "300102", 415, "img/kids/preview/02white-pink185X185.png", "img/kids/02white-pink500X500.jpg"}
var yellowKids = SubProduct{3, "Купить Кидсы (цвет желтый)", "300103", 415, "img/kids/preview/03yellow185X185.png", "img/kids/03yellow500X500.jpg"}
var yellowGreenKids = SubProduct{4, "Купить Кидсы (цвет желто-зеленый)", "300104", 415, "img/kids/preview/04yellow-green185X185.png", "img/kids/04yellowgreen500X500.jpg"}
var whiteKids = SubProduct{5, "Купить Кидсы (цвет белый)", "300105", 415, "img/kids/preview/05white185X185.png", "img/kids/05white500X500.jpg"}
var yellowBlueKids = SubProduct{6, "Купить Кидсы (цвет желто-голубые)", "300106", 415, "img/kids/preview/06blue-yellow185X185.png", "img/kids/06blue-yellow500X500.jpg"}
var pinkKids = SubProduct{7, "Купить Кидсы (цвет розовый)", "300107", 415, "img/kids/preview/07pink185X185.png", "img/kids/07pink500X500.jpg"}
var lightGreenKids = SubProduct{8, "Купить Кидсы (цвет салатовый)", "300108", 415, "img/kids/preview/08green185X185.png", "img/kids/08green500X500.jpg"}

var ironManPajama = SubProduct{9, "'Купить детскую пижаму железный человек (ironman)", "400101", 350, "img/pajamas/preview/01ironman185X185.png", "img/pajamas/01ironman500X500.png"}
var spiderManPajama = SubProduct{10, "Купить детскую пижаму человек паук (spiderman)", "400102", 350, "img/pajamas/preview/02spiderman185X185.png", "img/pajamas/02spiderman500X500.png"}
var superManPajama = SubProduct{11, "Купить детскую пижаму супермена", "400103", 350, "img/pajamas/preview/03superman185X185.png", "img/pajamas/03superman500X500.png"}
var monsterPajama = SubProduct{12, "Купить детскую пижаму корпорация монстров", "400104", 350, "img/pajamas/preview/04monsters185X185.png", "img/pajamas/04monsters500X500.png"}
var helloKittyPajama = SubProduct{13, "Купить детскую пижаму Hello Kitty", "400105", 350, "img/pajamas/preview/05hellokitty185X185.png", "img/pajamas/05hellokitty500X500.png"}
var helloKittyJustDancePajama = SubProduct{14, "Купить детскую пижаму Hello Kitty Just Dance", "400106", 350, "img/pajamas/preview/06hellokitty185X185.png", "img/pajamas/06hellokitty500X500.png"}
var helloKittyOldNavyPajama = SubProduct{15, "Купить детскую пижаму Hello Kitty Old Navy", "400107", 350, "img/pajamas/preview/07hellokitty185X185.png", "img/pajamas/07hellokitty500X500.png"}
var oneMoreHelloKittyPajama = SubProduct{16, "Купить детскую пижаму Hello Kitty", "400108", 350, "img/pajamas/preview/08hellokitty185X185.png", "img/pajamas/08hellokitty500X500.png"}

var angryBirdsPanties = SubProduct{17, "Купить детские трусики Angry Birds", "500101", 150, "img/panties/preview/01angrybirds185X185.png", "img/panties/01angrybirds500X500.jpg"}
var carsPanties = SubProduct{18, "Купить детские трусики Cars (тачки)", "500102", 150, "img/panties/preview/02cars185X185.png", "img/panties/02cars500X500.jpg"}
var doraemonPanties = SubProduct{19, "Купить детские трусики Doraemon", "500103", 150, "img/panties/preview/03doreamon185X185.png", "img/panties/03doreamon500X500.jpg"}
var mickeyMousePanties = SubProduct{20, "Купить детские трусики Mickey Mouse", "500104", 150, "img/panties/preview/04mickey185X185.png", "img/panties/04mickey500X500.jpg"}
var spiderManPanties = SubProduct{21, "Купить детские трусики Spiderman", "500105", 150, "img/panties/preview/05spiderman185X185.png", "img/panties/05spiderman500X500.jpg"}
var thomasPanties = SubProduct{22, "Купить детские трусики Паровозик Томас", "500106", 150, "img/panties/preview/06thomas185X185.png", "img/panties/06thomas500X500.jpg"}
var spanchBobPanties = SubProduct{23, "Купить детские трусики Спанч Боб", "500107", 150, "img/panties/preview/07spanchbob185X185.png", "img/panties/07spanchbob500X500.jpg"}
var packPanties = SubProduct{24, "Купить детские трусики набор 4 штуки", "500108", 315, "img/panties/preview/08mix185X185.png", "img/panties/08mix500X500.jpg"}

// sizes
var kidsSize24 = Size{1, "[24 размер, 2,5 года, стопа 14,6 см]"}
var kidsSize25 = Size{2, "[25 размер, 2,5 - 3 года, стопа 15,3 см]"}
var kidsSize26 = Size{3, "[26 размер, 3 - 4 года, стопа 16,5 см]"}
var kidsSize27 = Size{4, "[27 размер, 4 - 4,5 года, стопа 17,3 см]"}
var kidsSize28 = Size{5, "[28 размер, 4,5 - 5 лет, стопа 17,6 см]"}
var kidsSize29 = Size{6, "[29 размер, 5-6 лет, стопа 18,5 см]"}

var pajamaSize2T = Size{7, "[2T, рост 90 см, 2 года]"}
var pajamaSize3T = Size{8, "[3T, рост 95 см, 3 года]"}
var pajamaSize4T = Size{9, "[4T, рост 100 см, 4 года]"}
var pajamaSize5T = Size{10, "[5T, рост 110 см, 5 лет]"}
var pajamaSize6T = Size{11, "[6T, рост 120 см, 6 лет]"}
var pajamaSize7T = Size{12, "[7T, рост 130 см, 7 лет]"}

var pantiesSizeS = Size{13, "[размер S, рост 100-110 см, талия 22 см * 2, 4-6 лет]"}
var pantiesSizeM = Size{14, "[размер M, рост 110-120 см, талия 24 см * 2, 7-8 лет]"}
var pantiesSizeL = Size{15, "[размер L, рост 120-130 см, талия 26 см * 2, 9-10 лет]"}

var Products = []Product{{
	Id:          1,
	Category:    kidsCategory,
	SubProducts: []SubProduct{blueKids, whitePinkKids, yellowKids, yellowGreenKids, whiteKids, yellowBlueKids, pinkKids, lightGreenKids},
	Sizes:       []Size{kidsSize24, kidsSize25, kidsSize26, kidsSize27, kidsSize28, kidsSize29},
}, {
	Id:          2,
	Category:    pajamasCategory,
	SubProducts: []SubProduct{ironManPajama, spiderManPajama, superManPajama, monsterPajama, helloKittyPajama, helloKittyJustDancePajama, helloKittyOldNavyPajama, oneMoreHelloKittyPajama},
	Sizes:       []Size{pajamaSize2T, pajamaSize3T, pajamaSize4T, pajamaSize5T, pajamaSize6T, pajamaSize7T},
}, {
	Id:          3,
	Category:    pantiesCategory,
	SubProducts: []SubProduct{angryBirdsPanties, carsPanties, doraemonPanties, mickeyMousePanties, spiderManPanties, thomasPanties, spanchBobPanties, packPanties},
	Sizes:       []Size{pantiesSizeS, pantiesSizeM, pantiesSizeL},
}}

var SubProducts = []SubProduct{blueKids, whitePinkKids, yellowKids, yellowGreenKids, whiteKids, yellowBlueKids, pinkKids, lightGreenKids, ironManPajama,
	spiderManPajama, superManPajama, monsterPajama, helloKittyPajama, helloKittyJustDancePajama, helloKittyOldNavyPajama, oneMoreHelloKittyPajama,
	angryBirdsPanties, carsPanties, doraemonPanties, mickeyMousePanties, spiderManPanties, thomasPanties, spanchBobPanties, packPanties}
