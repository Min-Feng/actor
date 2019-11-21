actor模型主要說明  
請到[demo1資料夾](./cmd/demo1)查看

此處為嘗試 godoc example 功能的筆記

下載安裝godoc  
`go get -v -u golang.org/x/tools/cmd/godoc`

終端機開啟文件查看功能  
加上 -play 可線上執行  
但實際上無法引入第三方套件來進行example 的線上執行功能  
`godoc -http=:8888 -play`

到瀏覽器輸入 <http://localhost:8888/pkg/github.com/Min-Feng/actor/>  
進行線上查看

為什麼不能引入第三方套件的原因  
<https://stackoverflow.com/questions/41336341/cannot-find-package-error-on-godoc-playable-examples?rq=1>
