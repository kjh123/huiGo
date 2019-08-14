package user

import (
    "net/http"
)

type parseFormMiddleWare struct {
}

func (m *parseFormMiddleWare) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
    if err := r.ParseForm(); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    next(w, r)
}
