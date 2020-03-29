package main

//https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/

func replaceSpace(s string) string {
    if len(s) == 0{
        return ""
    }

    var ssl = []byte(s)
    var res string
    for _,v := range ssl {
        if v == ' '{
            res += "%20"
        }else {
            res += string(v)
        }
    }
    return res
}
