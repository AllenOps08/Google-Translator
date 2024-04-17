package router

import (
	"net/http"

	"github.com/AllenOps08/TranslatorBot/constant"
	"github.com/AllenOps08/TranslatorBot/controller"
)

// This routes represent user , translate from one language to another
var translateRoutes = Routes{
    Route{"Translate from to language", http.MethodPost , constant.TranslatePath, controller.Translate},
}