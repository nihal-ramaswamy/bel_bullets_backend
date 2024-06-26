package clubs

import (
	"bel_bullets/internal/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetClubByIdHandler(clubId string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authToken := ctx.GetString("Authorization")
		page, ok := ctx.GetQuery("page")

		if !ok {
			page = "1"
		}

		clubDetails, err := utils.GetClubById(clubId, authToken, page)

		if nil != err {
			log.Fatal(err.Error())
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, clubDetails)
	}
}
