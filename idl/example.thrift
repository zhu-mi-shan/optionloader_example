namespace go example

struct Req {
    1: required i64 id
}

struct Resp {
    1: string code
    2: string msg
}

service testService{
    Resp Test(1: Req req)
}
