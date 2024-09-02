package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenOK(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=1&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки

	// Запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое
	require.Equal(t, http.StatusOK, responseRecorder.Code)

	body := responseRecorder.Body.String()
	require.NotEmpty(t, body)
}

func TestMainHandlerWhenWrongCity(t *testing.T) {
	expectedBody := "wrong city value"
	req := httptest.NewRequest("GET", "/cafe?count=1&city=volgograd", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки

	// Город, который передаётся в параметре city, не поддерживается. Сервис возвращает код ответа 400 и ошибку wrong city value в теле ответа.
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	body := responseRecorder.Body.String()
	require.NotEmpty(t, body)
	require.Equal(t, expectedBody, body)

}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	expectedCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки

	//Если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе.
	require.Equal(t, http.StatusOK, responseRecorder.Code)

	body := responseRecorder.Body.String()
	require.NotEmpty(t, body)

	list := strings.Split(body, ",")
	assert.Len(t, list, expectedCount)

}
