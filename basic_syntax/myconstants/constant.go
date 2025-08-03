package myconstants

// 1. 首字母大写的全局变量（可导出，其他包可访问）
var FirstGlobalVar int = 123

// 2. 首字母小写的全局变量（不可导出，仅当前包内可见）
var secPrivateVar string = "qwe"
