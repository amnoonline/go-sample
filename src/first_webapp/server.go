package main

import (
	"net/http"
)

func main() {
	// マルチプレクサを作成する
	/*
		マルチプレクサとは:
		要求されたURLを調べ、リクエストを所定のハンドラにリダイレクトする機能
		デフォルトのマルチプレクサを作成するにはNewSeveMux()を使う
	*/
	mux := http.NewServeMux()

	// ルートディレクトリを任意のディレクトリに指定する
	/*
		マルチプレクサはハンドラへのリダイレクトだけでなく、静的なファイルの返送にも使える
		FileServer()とDir()を使いルートディレクトリを指定するハンドラを作成できる
		'/public'を起点（ルート）とするようにサーバに指示する
	*/
	files := http.FileServer(http.Dir("/public"))

	// 静的なファイルの返送を設定する
	/*
		StripPrefix()でfilesに渡ってくるURLから'/static/'以前のURLを取り除く
		StripPrefix(prefix string, h Handler):
		prefixで指定した文字列で始まる全てのリクエストURLについて、URLから 文字列/(指定したprefix)/ を取り除く(prefix部分も取り除かれる)

		例)
		http://localhost/static/css/bootstrap.min.css
		というファイルへのリクエストを受けた場合、サーバは
		<アプリケーションのルート(今回の場合では'/public')>/css/bootstrap.min.css
		のファイルを探す
		ファイルが見つかればサーバは何も加工せずにそのままファイルを返信する
	*/
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// ルートURLをハンドラ関数にリダイレクトする
	/*
		任意のURLを指定したハンドラ関数にリダイレクトするにはHandleFunc()を使う
		HandleFuncは第一引数にURLをとり、第2引数にハンドラ関数名をとる
		'/(ルートURL)'を'index(というハンドラ関数)'にリダイレクトするようにしている
	*/
	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()

}
