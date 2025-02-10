# 이미지 바이너리 빌드 할때 사용할 고랭 이미지
FROM golang:1.23.0-alpine AS builder

# 컨테이너 내무 작업 디렉토리 설정
WORKDIR /app

# 고랭 모듈 파일 복사 및 설치
# COPY go.mod go.sum ./
COPY go.mod ./
RUN go mod download

# 소스코드 복사
COPY . .

# 고 바이너리 빌드 (arm64 아키텍처). amd64 아키텍처로 빌드하려면 GOARCH=amd64 로 변경
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=linux GOARCH=$TARGETARCH go build -o search-cron .

# 이미지 바이너리 빌드 할때 사용할 베이스 이미지. 가벼운 알파인 리눅스 사용
FROM alpine:latest

# 빌드 스탭에서 빌드한 바이너리 /app으로 파일 복사
COPY --from=builder /app/search-cron .

# 포트 오픈
EXPOSE 8085

# 바이너리 실행
CMD ["./search-cron"]