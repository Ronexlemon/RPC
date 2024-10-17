package arithmeticserver

import (
	arithmeticproto "gorpc/arithmetic/proto"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)


func ArithClient(){
	conn,err := grpc.Dial("localhost:9090",grpc.WithInsecure())
	if err != nil{
		panic(err)
	}

	client := arithmeticproto.NewAddServiceClient(conn)
	g:= gin.Default()
	g.GET("/add/:a/:b",func(ctx *gin.Context) {
		a,err := strconv.ParseUint(ctx.Param("a"),10,64)
		if err != nil{
			ctx.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
			}
			b,err := strconv.ParseUint(ctx.Param("b"),10,64)
		if err != nil{
			ctx.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
			}
			req:= &arithmeticproto.Request{A:int64(a),B: int64(b)}
			if response,err := client.Add(ctx,req); err == nil{
				ctx.JSON(http.StatusOK,gin.H{"result":response.Result})
			}else{
				ctx.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
			}

	})
	g.GET("/mult/:a/:b",func(ctx *gin.Context) {
		a,err := strconv.ParseUint(ctx.Param("a"),10,64)
		if err != nil{
			ctx.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
			}
			b,err := strconv.ParseUint(ctx.Param("b"),10,64)
		if err != nil{
			ctx.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
			}
			req:= &arithmeticproto.Request{A:int64(a),B: int64(b)}
			if response,err := client.Multiply(ctx,req); err == nil{
				ctx.JSON(http.StatusOK,gin.H{"result":response.Result})
			}else{
				ctx.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
			}
		
	})
	if err := g.Run(":8080"); err !=nil{
		log.Fatal("Failed to run server",err)
	}

}