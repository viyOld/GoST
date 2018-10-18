package db

//_ "github.com/lib/pq"

func init() {
	// DB connect
	// host: teamo.work port: 5432 username: student password: 1509da4c70038a3ac38c89ddc0304274 database: test_db
	// Для работы создаете свою базу данных (ваша фамилия) *Чужие базы не трогаете!*

	//_, err = db.Exec("CREATE TABLE IF NOT EXISTS " + `voytenko_test("id" SERIAL PRIMARY KEY,` + `"name" varchar(10))`)
	//resultsql, err := db.Exec("\\list")
	//resultsql, err := db.Exec("\\dt")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(resultsql)
	//checkErr(err)
	//defer db.Close()

	// Данные для подключения к Minio:
	// Hostname: teamo.work:9000 AccessKey: 0A9VB38FYSAXMVUW39JV SecretKey: zLkDT/dbIn3Mo4XeU/myw/7bcu30bNAadMB6F1cX SQS ARNs:  arn:minio:sqs::1:nats
	// http://teamo.work:9000
	// *Вcе названия bucket`ов начинайте со своего префикса, например фамилии* *Чужие бакеты не трогаете!* (edited)
}
