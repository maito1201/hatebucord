# hatebucord
はてブ1位(テクノロジー)をdiscordに投げてワイワイするきっかけを作りたい

# Usage

環境変数`DISCORD_WEBHOOK`を適切に設定して、なんらかのランタイムでコードを実行してください

```main.go
package main

import "github.com/maito1201/hatebucord"

func main() {
	hatebucord.PostHatebu()
}

```