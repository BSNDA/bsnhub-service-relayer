package server

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"relayer/logging"
)

// HTTPService represents an HTTP service
type HTTPService struct {
	Router        *gin.Engine
	ChainManager  *ChainManager
}

// NewHTTPService constructs a new HTTPService instance
func NewHTTPService(
	chainManager *ChainManager,
) *HTTPService {
	srv := HTTPService{
		Router:        gin.Default(),
		ChainManager:  chainManager,
	}

	srv.createRouter()

	return &srv
}

func (srv *HTTPService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	srv.Router.ServeHTTP(w, r)
}

func (srv *HTTPService) createRouter() {
	r := gin.Default()

	api := r.Group("/api/v0")
	eth := api.Group("/eth")
	{
		eth.POST("/chains", srv.AddChain)
		eth.POST("/chains/:chainid/update", srv.UpdateChain)
		eth.POST("/chains/:chainid/delete", srv.DeleteChain)
		eth.POST("/chains/:chainid/start", srv.StartChain)
		eth.POST("/chains/:chainid/stop", srv.StopChain)
		eth.GET("/chains", srv.GetChains)
		eth.GET("/chains/:chainid/status", srv.GetChainStatus)
	}

	r.GET("/health", srv.ShowHealth)

	srv.Router = r
}

func (srv *HTTPService) AddChain(c *gin.Context) {
	var bodyBytes []byte
	// 从原有Request.Body读取
	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		c.JSON(http.StatusBadRequest, "invalid JSON payload")
		return
	}

	logging.Logger.Infof("AddChain Data is %s",string(bodyBytes))
	chainID, err := srv.ChainManager.AddChain(bodyBytes)
	if err != nil {
		onError(c, http.StatusInternalServerError, err.Error())
		return
	}

	onSuccess(c, AddChainResult{ChainID: chainID})
}

func (srv *HTTPService) DeleteChain(c *gin.Context) {
	chainID := c.Param("chainid")
	if err := ValidateChainID(chainID); err != nil {
		onError(c, http.StatusBadRequest, err.Error())
		return
	}

	err := srv.ChainManager.DeleteChain(chainID)
	if err != nil {
		onError(c, http.StatusInternalServerError, err.Error())
		return
	}

	onSuccess(c, nil)
}

func (srv *HTTPService) UpdateChain(c *gin.Context) {
	var bodyBytes []byte
	// 从原有Request.Body读取
	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		c.JSON(http.StatusBadRequest, "invalid JSON payload")
		return
	}
	chainID := c.Param("chainid")
	if err := ValidateChainID(chainID); err != nil {
		onError(c, http.StatusBadRequest, err.Error())
		return
	}

	err = srv.ChainManager.DeleteChain(chainID)
	if err != nil {
		onError(c, http.StatusInternalServerError, err.Error())
		return
	}
	logging.Logger.Infof("UpdateChain Data is %s",string(bodyBytes))

	chainID, err = srv.ChainManager.AddChain(bodyBytes)
	if err != nil {
		onError(c, http.StatusInternalServerError, err.Error())
		return
	}

	onSuccess(c, AddChainResult{ChainID: chainID})
}

func (srv *HTTPService) StartChain(c *gin.Context) {
	chainID := c.Param("chainid")
	if err := ValidateChainID(chainID); err != nil {
		onError(c, http.StatusBadRequest, err.Error())
		return
	}

	err := srv.ChainManager.StartChain(chainID)
	if err != nil {
		onError(c, http.StatusInternalServerError, err.Error())
		return
	}

	onSuccess(c, nil)
}

func (srv *HTTPService) StopChain(c *gin.Context) {
	chainID := c.Param("chainid")
	if err := ValidateChainID(chainID); err != nil {
		onError(c, http.StatusBadRequest, err.Error())
		return
	}

	err := srv.ChainManager.StopChain(chainID)
	if err != nil {
		onError(c, http.StatusInternalServerError, err.Error())
		return
	}

	onSuccess(c, nil)
}

func (srv *HTTPService) GetChains(c *gin.Context) {
	chains := srv.ChainManager.GetChains()

	onSuccess(c, chains)
}

func (srv *HTTPService) GetChainStatus(c *gin.Context) {
	chainID := c.Param("chainid")
	if err := ValidateChainID(chainID); err != nil {
		onError(c, http.StatusBadRequest, err.Error())
		return
	}

	state, height, err := srv.ChainManager.GetChainStatus(chainID)
	if err != nil {
		onError(c, http.StatusInternalServerError, err.Error())
		return
	}

	onSuccess(c, ChainStatus{State: state, Height: height})
}

// ShowHealth returns the health state
func (srv *HTTPService) ShowHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": true})
}

func onError(c *gin.Context, code int, msg string) {
	logging.Logger.Errorf(msg)

	c.JSON(code, ErrorResponse{
		Code:  CODE_ERROR,
		Error: msg,
	})
}

func onSuccess(c *gin.Context, result interface{}) {
	c.JSON(http.StatusOK, SuccessResponse{
		Code:   CODE_SUCCESS,
		Result: result,
	})
}
