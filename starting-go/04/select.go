package main

import (
	"fmt"
)

func main() {
	//e1 := <-ch1 // 受信待ちでgoroutineが停止

	// ↑で処理が停止しているので下記にはたどり着かない
	//e2 := <-ch2

	/* goroutineを停止させずに複数のチャネルをコントロールする 方法 */

	// 複数条件のcaseが成り立つ場合、goのランタイムはどのcase節を実行するかランダムに決定する
	select {
		case e1 := <-ch1
			// ch1から受信した場合の処理
		case e2 := <-ch2
			// ch2からの受信が成功した場合の処理
		default:
			// case節の条件が成立しなかった場合の処理
	}
}
