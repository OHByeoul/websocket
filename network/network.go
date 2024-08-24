package network

import (
	"github.com/gin-contrib/gin"
	"github.com/gin-gonic/cors"
	"log"
)

type Network struct {
	engin *gin.Engine
}

func NewServer() *Network {
	n := &Network{engin: gin.New()}

	// 미들웨어 추가
	n.engin.Use(gin.Logger()) // 모든 http 요청에 대해 로그 남김
	n.engin.Use(gin.Recovery()) // 서버 죽을 경우 recovery
	n.engin.Use(cors.New(cors.Config{
		AllowWebSockets: true,
		AllowOrigins:	 []string{"*"}
		AllowMethods:	 []string{"GET","POST","PUT"}
		AllowHeaders:	 []string{"*"},
		AllowCredentials: true, // 쿠키 인증정보 허용
	}))


	return n
}

func (n *Network) StartServer() error { // 리시버  - 어떤 타입에 대해 호출될 것인지 정의 *Network
	log.Println("Starting server")
	return n.engin.Run(":8080")
}