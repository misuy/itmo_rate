package DB

import (
	"itmo_rate/DB/entities"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open() (db *gorm.DB, err error) {
	dsn := os.Getenv("DB_URL")
	//dsn := "host=localhost user=itmo_rate dbname=itmo_rate_db port=5432"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	err = db.AutoMigrate(
		&entities.Faculty{},
		&entities.Subject{},
		&entities.Teacher{},
		&entities.User{},
		&entities.Review{},
		&entities.TeacherReviewRel{},
		&entities.SubjectTeacherRel{},
		&entities.CriteriaList{},
		&entities.Criteria{},
	)

	db.SetupJoinTable(&entities.Teacher{}, "Reviews", &entities.TeacherReviewRel{})
	db.SetupJoinTable(&entities.Subject{}, "Teachers", &entities.SubjectTeacherRel{})
	return
}

func SaveTestData(db *gorm.DB) {
	faculties := []entities.Faculty{
		entities.NewFaculty("ФПИиКТ"),
		entities.NewFaculty("ВТ"),
	}

	subjects := []entities.Subject{
		entities.NewSubject("ТПО"),
		entities.NewSubject("ОПД"),
		entities.NewSubject("Компьютерные сети"),
		entities.NewSubject("Функциональная схемотехника"),
		entities.NewSubject("Разработка компиляторов"),
		entities.NewSubject("Системы ввода-вывода"),
	}

	teachers := []entities.Teacher{
		entities.NewTeacher("Клименков Сергей Викторович"),
		entities.NewTeacher("Харитонова Анастасия Евгеньевна"),
		entities.NewTeacher("Осипов Святослав Владимирович"),
		entities.NewTeacher("Алиев Тауфик Измайлович"),
		entities.NewTeacher("Авксеньтьева Елена Юрьевна"),
		entities.NewTeacher("Кустарев Павел Велерьевич"),
		entities.NewTeacher("Васильев Сергей Евгеньевич"),
		entities.NewTeacher("Лаздин Артур Вячеславович"),
		entities.NewTeacher("Быковский Сергей Вячеславович"),
	}

	users := []entities.User{
		entities.NewUser("Ненов Владислав"),
		entities.NewUser("Передрий Михаил"),
	}

	reviews := []entities.Review{
		entities.NewReview(entities.NewDefaultCriteriaList(5, 5, 5, 5, 5).ToCriteriaList(), "На курс могут быть зачислены только студенты вузов, начиная с 3-го курса.\nУчастник должен быть готов к очной стажировке в Москве, Санкт-Петербурге или Иннополисе", true),
		entities.NewReview(entities.NewDefaultCriteriaList(7.8, 7.4, 2, 10, 4).ToCriteriaList(), "Мезоген — основная единица жидкого кристалла, которая порождает его структурный порядок.\nОбычно жидкокристаллическая молекула состоит из жесткой функциональной группы и одной или больше гибких частей. Жесткая часть выстраивает молекулы в одном направлении, в то время как гибкие части ответственны за текучее состояние жидкого кристалла. Эта жесткая часть и есть мезоген. Она обладает ключевой ролью в жидком кристалле, так как оптимальный баланс этих двух частей необходим для формирования жидкокристаллических материалов.", true),
		entities.NewReview(entities.NewDefaultCriteriaList(6, 8.8, 7, 10, 9).ToCriteriaList(), "В лингвистике термин «текст» используется в широком значении, включая и образцы устной речи. Восприятие текста изучается в рамках лингвистики текста и психолингвистики. Так, например, И. Р. Гальперин определяет текст следующим образом: «Это письменное сообщение, объективированное в виде письменного документа, состоящее из ряда высказываний, объединённых разными типами лексической, грамматической и логической связи, имеющее определённый модальный характер, прагматическую установку и соответственно литературно обработанное». ", true),
		entities.NewReview(entities.NewDefaultCriteriaList(7, 5, 7, 7, 8).ToCriteriaList(), "Здравствуйте, Передрий Михаил Сергеевич!\nИнформируем Вас, что Лекция по дисциплине Компьютерные сети, преподаватель Алиев Т.И., состоится 30.04.2024 18:40 в дистанционном формате для студентов, обучающихся на основании уведомления о дистанционном формате обучения, отправленного в Студенческий офис. \nПодключиться к видео - конференции Вы можете по ссылке, идентификатор конференции: 84280658133, пароль: 155233 \nЕсли вы не подавали уведомление о дистанционном формате обучения, но получили данное письмо, Вы можете воспользоваться ссылкой в случае если не можете присутствовать на занятии очно.", true),
		entities.NewReview(entities.NewDefaultCriteriaList(10, 9, 9.5, 10, 7).ToCriteriaList(), "Всем привет! Мы рады, что вы сумели пройти непростой отбор и попасть на этот курс. Совсем скоро мы начнем изучение Golang. Нас ждет много сложных, но интересных задач, которые помогут вам стать востребованным специалистом. Наша команда составила подробную программу, в которой вас ждут лекции, воркшопы, домашние задания, ревью кода и ответы на вопросы. Помимо Go мы разберем много полезных инструментов, таких как: Postgres, Kafka, Redis и многих других, без которых трудно представить современную разработку. После завершения курса лучших студентов мы будем ждать в рядах стажеров, которые помогут разрабатывать и улучшать проекты Ozon!", true),
		entities.NewReview(entities.NewDefaultCriteriaList(7.8, 7.4, 8, 10, 4).ToCriteriaList(), "Project presentation evaluation\nProject presentation should include:\n1. Introduction of the participants and their roles in the project\n2. General description of the idea (how it emerged)\n3. Process of development (issues, difficulties, breakthroughs, discoveries, etc.)\n4. Product overview for techies (if students choose to run a demo it should not exceed 2 minutes and must be commented on)\n5. Sales overview for end users (interface, application, utilities, benefits for end user = ' how it will make your life easier')\n6. If the software is only partially functioning or could be integrated into the bigger project - perspectives for the future development", true),
		entities.NewReview(entities.NewDefaultCriteriaList(9, 8, 8.8, 3, 3).ToCriteriaList(), "Газета выходила в Санкт-Петербурге с 1 июня 1862 по 2 июня 1863 года.\nРедактировал газету русский инженер, организатор электротехнического института Н. Г. Писаревский.\n«Современное слово» возникло в результате выделения в особое издание неофициальной части газеты «Русский инвалид». В 1862 году служило приложением к «Русскому инвалиду», в 1863 году выходило самостоятельно. С 1862 года редактором «Современного слова» являлся В. Н. Леонтьев, который, среди прочего, поместил в издании ряд своих статей о крестьянской реформе 1848 года. ", true),
		entities.NewReview(entities.NewDefaultCriteriaList(6, 6, 7, 7, 7).ToCriteriaList(), "Samsung Galaxy A32 — Android-смартфон среднего класса, разработанный и произведенный Samsung Electronics. Он является преемником Galaxy A31. Телефон похож на своего предшественника, но имеет улучшенную основную камеру на 64 Мп. Устройство также поставляется с вариантом 5G с урезанной камерой и экраном (основная камера 48 МП и TFT-дисплей PLS), но с более быстрой SoC.\nGalaxy A32 5G был первым, выпущенным 22 января 2021 года, а Galaxy A32 был выпущен позже, 17 февраля 2021 года. ", true),
		entities.NewReview(entities.NewDefaultCriteriaList(8, 9.5, 9.7, 7, 6).ToCriteriaList(), "Родился 4 декабря 1928 года в Котласе. После окончания средней школы в 1945 году переехал в Москву и поступил на библиографический факультет МГБИ, который он окончил в 1950 году, в том же году поступил на заочную аспирантуру МГБИ, которую он окончил в 1956 году. В 1950 году был принят на работу в Читинскую ОНБ имени А. С. Пушкина на должность библиографа и заведующего справочно-библиографическим отделом. В 1956 году окончательно переехал в Москву и был принят на работу в МГБИ, где он работал на кафедре библиографии в качестве ассистента, старшего преподавателя и доцента вплоть до смерти.", true),
		entities.NewReview(entities.NewDefaultCriteriaList(8, 7, 4, 8, 5).ToCriteriaList(), "Купе́ль — большой чашеобразный сосуд. Служит для проведения таинства крещения в христианской церкви. Купель может быть выполнена из различных материалов и играет важную роль в создании внутреннего убранства церковного помещения. Нередко купель является произведением искусства.\n«Купелью» иногда называют ещё «иордáнь» — ледяную прорубь, в которой купаются в праздник Крещения Господня. ", true),
		entities.NewReview(entities.NewDefaultCriteriaList(9.9, 5.5, 7.7, 9, 5).ToCriteriaList(), "Rock’s Backpages — веб-сайт, онлайн библиотека, содержащая материалы музыкальной прессы начиная с 1950 года по сегодняшний день. Ресурс был основан в 2000 году британским журналистом Барни Хоскинсом (Barney Hoskyns). По состоянию на январь 2008 года база данных содержит свыше 12 000 статей (интервью, характеристики, отзывы и т.п.) и охватывает широкий круг исполнителей популярной музыки (в том числе блюза и соула).[1]\nДоступ платный, однако часть материалов находится в свободном доступе (необходима регистрация)[2].\nБаза данных включает статьи из журналов Creem, Rolling Stone, New Musical Express, Melody Maker, Crawdaddy!, Mojo и других, содержит статьи более чем 300 журналистов, главным образом из Соединённых Штатов и Великобритании, в том числе Дэйв Марш (Dave Marsh), Ник Кент (Nick Kent), Чарльз Шаар Мюррей (Charles Shaar Murray), Ник Тосчез (Nick Tosches), Мик Фаррен (Mick Farren) и Ян Макдональд (Ian MacDonald). Все материалы в базе данных представлены при полном согласии и с разрешения владельцев авторских прав. ", true),
		entities.NewReview(entities.NewDefaultCriteriaList(8, 8, 3, 2, 7).ToCriteriaList(), "Knights and Merchants: The Shattered Kingdom (дословно с англ. — «Рыцари и купцы. Разрушенное королевство»; в России и странах СНГ была издана под названиями «Война и мир» и «Разделяй и властвуй») — компьютерная игра в жанре стратегии в реальном времени немецкой игровой школы, разработанная компанией Joymania Entertainment для ПК и портированная компанией e.p.i.c. Interactive Entertainment для Amiga-совместимой платформы Pegasos. 21 февраля 2003 года игра появилась в версии для MorphOS. Существует также версия, портированная для Mac OS.\nЧерез четыре года вышло продолжение игры Knights and Merchants: The Peasants Rebellion (в России и странах СНГ была издана под названием «Вторая корона») с небольшим рядом новшеств и изменений, а также с новой кампанией и новыми одиночными миссиями. ", true),
	}

	db.Create(&reviews)
	db.Create(&users)
	db.Create(&teachers)
	db.Create(&subjects)
	db.Create(&faculties)

	faculties[0].AddSubjects(db, subjects)
	faculties[1].AddSubjects(db, []entities.Subject{subjects[0], subjects[1], subjects[3], subjects[4]})

	subjects[0].AddTeacher(db, &teachers[0], "lecturer")
	subjects[0].AddTeacher(db, &teachers[1], "teacher")
	subjects[0].AddTeacher(db, &teachers[0], "teacher")

	subjects[0].AddReview(db, &reviews[0])
	teachers[0].AddReview(db, &reviews[0], "lecturer")
	teachers[1].AddReview(db, &reviews[0], "teacher")
	users[0].AddReview(db, &reviews[0])

	subjects[0].AddReview(db, &reviews[1])
	teachers[0].AddReview(db, &reviews[1], "lecturer")
	teachers[0].AddReview(db, &reviews[1], "teacher")
	users[0].AddReview(db, &reviews[1])

	subjects[1].AddTeacher(db, &teachers[0], "lecturer")
	subjects[1].AddTeacher(db, &teachers[2], "teacher")
	subjects[1].AddTeacher(db, &teachers[0], "teacher")

	subjects[1].AddReview(db, &reviews[2])
	teachers[0].AddReview(db, &reviews[2], "lecturer")
	teachers[2].AddReview(db, &reviews[2], "teacher")
	users[0].AddReview(db, &reviews[2])

	subjects[1].AddReview(db, &reviews[3])
	teachers[0].AddReview(db, &reviews[3], "lecturer")
	teachers[0].AddReview(db, &reviews[3], "teacher")
	users[0].AddReview(db, &reviews[3])

	subjects[2].AddTeacher(db, &teachers[3], "lecturer")
	subjects[2].AddTeacher(db, &teachers[3], "teacher")
	subjects[2].AddTeacher(db, &teachers[4], "teacher")

	subjects[2].AddReview(db, &reviews[4])
	teachers[3].AddReview(db, &reviews[4], "lecturer")
	teachers[3].AddReview(db, &reviews[4], "teacher")
	users[0].AddReview(db, &reviews[4])

	subjects[2].AddReview(db, &reviews[5])
	teachers[3].AddReview(db, &reviews[5], "lecturer")
	teachers[4].AddReview(db, &reviews[5], "teacher")
	users[0].AddReview(db, &reviews[5])

	subjects[3].AddTeacher(db, &teachers[5], "lecturer")
	subjects[3].AddTeacher(db, &teachers[5], "teacher")
	subjects[3].AddTeacher(db, &teachers[6], "teacher")

	subjects[3].AddReview(db, &reviews[6])
	teachers[5].AddReview(db, &reviews[6], "lecturer")
	teachers[5].AddReview(db, &reviews[6], "teacher")
	users[0].AddReview(db, &reviews[6])

	subjects[3].AddReview(db, &reviews[7])
	teachers[5].AddReview(db, &reviews[7], "lecturer")
	teachers[6].AddReview(db, &reviews[7], "teacher")
	users[0].AddReview(db, &reviews[7])

	subjects[4].AddTeacher(db, &teachers[7], "lecturer")
	subjects[4].AddTeacher(db, &teachers[7], "teacher")

	subjects[4].AddReview(db, &reviews[8])
	teachers[7].AddReview(db, &reviews[8], "lecturer")
	teachers[7].AddReview(db, &reviews[8], "teacher")
	users[0].AddReview(db, &reviews[8])

	subjects[4].AddReview(db, &reviews[9])
	teachers[7].AddReview(db, &reviews[9], "lecturer")
	teachers[7].AddReview(db, &reviews[9], "teacher")
	users[0].AddReview(db, &reviews[9])

	subjects[5].AddTeacher(db, &teachers[8], "lecturer")
	subjects[5].AddTeacher(db, &teachers[8], "teacher")

	subjects[5].AddReview(db, &reviews[10])
	teachers[8].AddReview(db, &reviews[10], "lecturer")
	teachers[8].AddReview(db, &reviews[10], "teacher")
	users[0].AddReview(db, &reviews[10])

	subjects[5].AddReview(db, &reviews[11])
	teachers[8].AddReview(db, &reviews[11], "lecturer")
	teachers[8].AddReview(db, &reviews[11], "teacher")
	users[0].AddReview(db, &reviews[11])

	db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&subjects)
	db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&reviews)
	db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&teachers)
}
