package Handler

import (
	"First/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	customersWithMsg = map[int]model.CustomerWithMsg{}
	customers        = map[int]model.Customer{}
	id               = 1
)

type request struct {
	CName         string
	CAddress      string
	CTel          int
	CRegisterDate time.Time
}

func Create(c echo.Context) error {
	thisTime := time.Now().Format("2006-01-02")
	var req request
	msg := "success"
	if err := c.Bind(&req); err != nil {
		//return c.JSON(http.StatusBadRequest,)
		return echo.NewHTTPError(http.StatusBadRequest, err) //behtare khode err ro bar nagardoonim
	}
	m := model.CustomerWithMsg{
		CId:           id,
		CName:         req.CName,
		CTel:          req.CTel,
		CAddress:      req.CAddress,
		CRegisterDate: thisTime,
		CMsg:          msg,
	}
	customersWithMsg[id] = m
	m2 := model.Customer{
		CId:           id,
		CName:         req.CName,
		CTel:          req.CTel,
		CAddress:      req.CAddress,
		CRegisterDate: thisTime,
	}
	customers[id] = m2
	id++
	return c.JSON(http.StatusCreated, m)
}
func Edit(c echo.Context) error {
	thisTime := time.Now().Format("2006-01-02")
	ThisId, _ := strconv.Atoi(c.Param("id"))
	var req request
	var flag bool
	SuccessMsg := "success"
	ErrMsg := "Id not found!"
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err) //behtare khode err ro bar nagardoonim
	}
	for key := range customersWithMsg {
		if ThisId == customersWithMsg[key].CId {
			flag = true
			break
		} else {
			flag = false
		}
	}
	if flag {
		m := model.CustomerWithMsg{
			CId:           ThisId,
			CName:         req.CName,
			CTel:          req.CTel,
			CAddress:      req.CAddress,
			CRegisterDate: thisTime,
			CMsg:          SuccessMsg,
		}
		customersWithMsg[ThisId] = m
		m2 := model.Customer{
			CId:           ThisId,
			CName:         req.CName,
			CTel:          req.CTel,
			CAddress:      req.CAddress,
			CRegisterDate: thisTime,
		}
		customers[ThisId] = m2
		return c.JSON(http.StatusCreated, m)
	} else {
		m := model.Massage{
			Msg: ErrMsg,
		}
		return c.JSON(http.StatusCreated, m)
	}
}
func Delete(c echo.Context) error {
	ThisId, _ := strconv.Atoi(c.Param("id"))
	var flag bool
	ErrMsg := "Id not found!"
	for key := range customersWithMsg {
		if ThisId == customersWithMsg[key].CId {
			flag = true
			break
		} else {
			flag = false
		}
	}
	if flag {
		delete(customers, ThisId)
		delete(customersWithMsg, ThisId)
		return c.JSON(http.StatusOK, "deleted")
	} else {
		m := model.Massage{
			Msg: ErrMsg,
		}
		return c.JSON(http.StatusCreated, m)
	}
}
func Report(c echo.Context) error {
	ErrMsg := "Enter Month between 0 ,1, ... ,11"
	counter := 0
	Month, _ := strconv.Atoi(c.Param("month")) //int
	if Month >= 0 && Month < 12 {
		for key := range customers {
			monthOfCustomer := strings.Split(customers[key].CRegisterDate, "-")[1] //string
			monthOfCustomerNumber, _ := strconv.Atoi(monthOfCustomer)
			if monthOfCustomerNumber == Month+1 {
				counter++
			}
		}
		m := model.Report{
			TotalCustomers: counter,
			Period:         Month,
			Msg:            "success",
		}
		return c.JSON(http.StatusOK, m)
	} else {
		m := model.Massage{
			Msg: ErrMsg,
		}
		return c.JSON(http.StatusCreated, m)
	}
}
func Information(c echo.Context) error {
	ThisQueryName := c.QueryParam("cName")
	if ThisQueryName != "" {
		counter := 0
		ThisNameR := strings.Split(ThisQueryName, "{")[1]
		ThisName := strings.Split(ThisNameR, "}")[0]
		CustomersFound := map[int]model.Customer{}
		for key := range customers {
			if strings.HasPrefix(customers[key].CName, ThisName) {
				CustomersFound[counter] = customers[key]
			}
			counter++
		}
		return c.JSON(http.StatusOK, CustomersFound)
	} else {
		ErrMsg := "no customers found!"
		if len(customersWithMsg) == 0 {
			m := model.Massage{
				Msg: ErrMsg,
			}
			return c.JSON(http.StatusCreated, m)
		} else {
			m := model.Customers{
				Size:      len(customers),
				Customers: customers,
				CMsg:      "success",
			}
			return c.JSON(http.StatusOK, m)
		}
	}
}