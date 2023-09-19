-- 检测是不是预期中的值
-- 如果是，删除，如果不是，返回一个值
if redis.Call("get", KEYS[1]) == ARGV[1] then
    return redis.Call("del", KEYS[1])
else
    -- 返回0 代表key不存在或值不对
    return 0
end