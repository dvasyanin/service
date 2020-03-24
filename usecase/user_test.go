package usecase_test

import (
	"encoding/json"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"modz/mocks"
	"modz/model"
	"modz/usecase"
)

var _ = Describe(".Usecase", func() {
	var (
		mockCtrl        *gomock.Controller
		mockRepo        *mocks.MockUserRepo
		mbService       usecase.User
		mbResp          []byte
		clientID        string
		userInfo        model.UserInformation
		customer        model.User
		balance         []model.Balances
		total           model.OrderStatistics
		userInfoInvalid model.UserInformation
		clientIDInvalid string
		mbRespInvalid   []byte
		body            string
		bodyInvalid     string
		url             string
	)
	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockUserRepo(mockCtrl)
		mbService = usecase.NewUserUsecase(mockRepo)

		url = "&operation=ClientSearch"

		clientID = "1"
		body = `{
  "customer": {
    "ids": {
      "storeezWebsiteId": "` + clientID + `"
    }
  }
}`
		customer = model.User{
			FirstName: "valid",
		}
		balance = append(balance, model.Balances{Total: 1})
		total = model.OrderStatistics{TotalPaidAmount: 1}
		userInfo = model.UserInformation{
			Customer:   customer,
			Balance:    balance,
			Statistics: total,
		}
		mbResp, _ = json.Marshal(userInfo)

		clientIDInvalid = "0"
		bodyInvalid = `{
  "customer": {
    "ids": {
      "storeezWebsiteId": "` + clientIDInvalid + `"
    }
  }
}`

		userInfoInvalid = model.UserInformation{
			Customer:   model.User{},
			Balance:    nil,
			Statistics: model.OrderStatistics{},
		}

		mbRespInvalid, _ = json.Marshal(userInfoInvalid)
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("User  information", func() {
		Context(".Info()", func() {
			BeforeEach(func() {

				mockRepo.EXPECT().Get(url, body).Return(mbResp, nil).AnyTimes()
			})
			It("should be a valid value response", func() {
				Expect(mbService.Info(clientID)).To(Equal(&userInfo))
			})

			BeforeEach(func() {
				mockRepo.EXPECT().Get(url, bodyInvalid).Return(mbRespInvalid, nil).AnyTimes()
			})
			It("should error", func() {
				_, err := mbService.Info(clientIDInvalid)
				Ω(err).Should(HaveOccurred())
			})
		})
	})
})
