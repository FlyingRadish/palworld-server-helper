package api

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"pal-server-helper/common"
	"pal-server-helper/pal"
	"pal-server-helper/pal/rcn"

	"github.com/gin-gonic/gin"
)

type rebootParam struct {
	NotifyReboot bool `json:"notifyReboot"`
}

type broadcastParam struct {
	Content string `json:"content"`
}

type rconParam struct {
	Command string `json:"command"`
}

type simpleParam struct {
	Data string `json:"data"`
}

func RunApiServer(port int) {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	router.Static("/panel", "./panel")
	router.Static("/assets", "./panel/assets")

	// 返回当前在线玩家信息
	router.GET("/api/players", func(c *gin.Context) {
		if pal.OnlinePlayers == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Players not ready"})
			return
		}
		c.JSON(http.StatusOK, pal.OnlinePlayers)
	})

	/** 游戏内广播
	**	content: 广播内容
	**/
	router.POST("/api/broadcast", func(c *gin.Context) {
		var json broadcastParam
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if len(json.Content) <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "empty content is not allowed"})
			return
		}
		rcn.GetRCNClient().Broadcast(json.Content)
		c.JSON(http.StatusNoContent, nil)
	})

	/** 转发RCON请求
	**	content: 广播内容
	**/
	router.POST("/api/rcon", func(c *gin.Context) {
		var json rconParam
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if len(json.Command) <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "empty command is not allowed"})
			return
		}
		response, err := rcn.GetRCNClient().ExecCommand(json.Command)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to forward rcon command"})
			return
		}
		c.JSON(http.StatusOK, simpleParam{Data: response})
	})

	/** 重启服务器
	**	notifyReboot: 是否在游戏内广播倒计时
	**/
	router.POST("/api/reboot", func(c *gin.Context) {
		var json rebootParam
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		pal.Reboot(json.NotifyReboot)
		c.JSON(http.StatusNoContent, nil)
	})

	// 返回当前内存占用
	router.GET("/api/memory", func(c *gin.Context) {
		memStatus, err := common.GetMemoryStats()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get memory status"})
			return
		}
		c.JSON(http.StatusOK, memStatus)
	})

	// Run the server on port 8080
	servePort := ":" + strconv.Itoa(port)
	fmt.Println("API server serve on " + servePort)
	router.Run(servePort)
}

func UpdatePanelApi(apiHost string, apiPort int) {
	filePath := "dist/env.js"

	// 删除文件
	if _, err := os.Stat(filePath); err == nil {
		err = os.Remove(filePath)
		if err != nil {
			fmt.Println("Error occurred while deleting file:", err)
			return
		}
		fmt.Println("File deleted successfully")
	}

	// 写入新内容
	apiURL := fmt.Sprintf("http://%s:%d", apiHost, apiPort)
	newContent := fmt.Sprintf(`window.palServerHelper={"apiUrl": "%s"}`, apiURL)

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error occurred while creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(newContent)
	if err != nil {
		fmt.Println("Error occurred while writing to file:", err)
		return
	}
	fmt.Println("Update panel api file successfully")
}
