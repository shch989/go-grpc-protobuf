package main

import (
	"fmt" // 표준 출력 및 문자열 처리에 사용
	"log" // 로그 처리에 사용
	"my-protobuf/basic"
	"time" // 시간 관련 기능 제공
)

// logWrither 구조체 정의
// log.SetOutput에서 사용될 커스텀 로그 작성기 구조체
type logWrither struct {
}

// logWrither에 Write 메서드 정의
// log 패키지는 출력 대상에 Write(bytes []byte) (int, error) 메서드가 구현되어 있어야 합니다.
func (writer logWrither) Write(bytes []byte) (int, error) {
	// 현재 시간을 "시:분:초" 형식으로 포맷팅하고, 전달받은 로그 메시지를 추가로 출력
	return fmt.Println(time.Now().Format("15:04:05") + " " + string(bytes))
}

func main() {
	// 로그의 기본 플래그(날짜, 시간 등)를 제거
	log.SetFlags(0)

	// log 패키지가 사용하는 출력 대상을 logWrither로 설정
	log.SetOutput(new(logWrither))

	// basic.BasicHello()
	// basic.BasicUser()
	// basic.ProtoToJsonUser()
	// basic.JsonToProtoUser()
	// basic.BasicUserGroup()
	// jobsearch.JobSearchCandidate()
	// jobsearch.JobSearchSoftware()
	// basic.BasicUnmarshalAnyKnown()
	// basic.BasicUnmarshalAnyNotKnown()
	basic.BasicUnmarshalAnyIs()
}
