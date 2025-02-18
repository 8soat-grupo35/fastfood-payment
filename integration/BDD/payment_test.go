package main_test

import (
	"net/http"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Pagamento", func() {
	Context("Dado que o usuário consulta o status do seu pagamento", func() {
		When("quando o pagamento existe", func() {
			BeforeEach(func() {
				req, _ := http.NewRequest(http.MethodPost, "http://localhost:8000/v1/payments/", strings.NewReader(`{"orderId": 99}`))
				req.Header.Set("Content-Type", "application/json")
				_, err := http.DefaultClient.Do(req)
				assert.NoError(GinkgoT(), err)
			})

			It("deve retornar os dados do pagamento corretamente", func() {
				req, _ := http.NewRequest(http.MethodGet, "http://localhost:8000/v1/payments/4/payment/status", nil)
				req.Header.Set("Content-Type", "application/json")
				res, err := http.DefaultClient.Do(req)

				assert.NoError(GinkgoT(), err)
				assert.Equal(GinkgoT(), http.StatusOK, res.StatusCode)
			})
		})
	})
})
