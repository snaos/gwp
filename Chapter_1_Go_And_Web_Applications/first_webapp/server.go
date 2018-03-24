package main // package  선언
// 실행 가능한 프로그램이 되려면 패키지 이름은 항상 'main'이 되어야 한다.
// ruby, python, java와 같은 언어와 차이는 앱을 server에 deploy하는 과정이 필요 없다.
import (
	"fmt"      // I/O를 formatting하기 위한 라이브러리
	"net/http" // HTTP와 상호작용을 위한 메인 패키지
)

func handler(writer http.ResponseWriter, request *http.Request) { //handler는 이벤트 발생시 연이어 호출되는 콜백 함수를 의미.
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)     // /에 요청이 오면 handler 함수로 처리하도록 트리거를 추가
	http.ListenAndServe(":8080", nil) // 8080포트에서 동작하는 서버.
}

/*
 go install directory

 이 명령어를 사용하기 위해서는 GOPATH 환경변수가 설정되어 있어야 한다.
 위 명령어를 이용하면 $GOPATH/bin 디렉터리에 directory라고 명명된 실행 가능한 바이너리 파일이 만들어진다.
*/
