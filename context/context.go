package context

import (
	"context"
	"log"
	"net/http"
	"fmt"
)
func ContextMiddleware(w http.ResponseWriter, r *http.Request){
		// Extract the context from the request
		ctx := r.Context()
		log.Println(ctx)	
		context.Background()
		//var ctxs = context.Background()
		fmt.Println("context: ",ctx)
}
