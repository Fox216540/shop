package api

import "net/http"

func refreshCookie(token string, maxAge int) *http.Cookie {
	return &http.Cookie{
		Name:     messages.RefreshCookieName,
		Value:    token,
		Path:     "/",
		MaxAge:   maxAge,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}
}

func expiredRefreshCookie() *http.Cookie {
	return refreshCookie("", -1)
}
