package Entities

type StatusCode struct{ value int }

var OK = StatusCode{200}
var Created = StatusCode{201}
var Accepted = StatusCode{202}
var Non_Authoritative_Information = StatusCode{203}
var No_Content = StatusCode{204}
var Reset_Content = StatusCode{205}
var Partial_Content = StatusCode{206}


