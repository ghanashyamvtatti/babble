package DAL

import "ds-project/config"

type TokenDB struct {
	Tokens        map[string]string
}

func SetAccessToken(ctx context.Context, kv clientv3.KV, username string, token string) bool {
	bt := RAFT.GetKey(ctx,kv,"tokens")

	var result TokenDB
    err:= json.Unmarshal(bt, &result)
    if err != nil {
        panic(err)
    }

    result.Tokens[username] = token
	marshalledToken, err := json.Marshal(result)
	RAFT.PutKey(ctx,kv,"tokens",marshalledToken)
	return true
}

func GetAccessToken(ctx context.Context, kv clientv3.KV, username string) (string, bool) {
	bt := RAFT.GetKey(ctx,kv,"tokens")

	var result TokenDB
    err:= json.Unmarshal(bt, &result)
    if err != nil {
        panic(err)
    }

	token, ok := result.Tokens[username]
	return token, ok
}

func DeleteAccessToken(ctx context.Context, kv clientv3.KV, username string) bool {
	bt := RAFT.GetKey(ctx,kv,"tokens")

	var result TokenDB
    err:= json.Unmarshal(bt, &result)
    if err != nil {
        panic(err)
    }

   	delete(result.Tokens, username)

	marshalledToken, err := json.Marshal(result)
	RAFT.PutKey(ctx,kv,"tokens",marshalledToken)
	return true
}
