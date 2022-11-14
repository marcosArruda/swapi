package httpservice

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/marcosArruda/swapi/pkg/services"

	"github.com/gin-gonic/gin"
)

type (
	httpServiceFinal struct {
		sm     services.ServiceManager
		router *gin.Engine
		srv    *http.Server
	}
)

func NewHttpService(ctx context.Context) services.HttpService {
	return &httpServiceFinal{}
}

func (n *httpServiceFinal) Start(ctx context.Context) error {
	gin.SetMode(gin.ReleaseMode)
	n.router = gin.Default()

	n.router.GET("/hello/:name", n.Hello)
	n.router.GET("/planet/:id", n.GetPlanetById)
	n.router.GET("/planet/n/:name", n.SearchPlanetsByName)
	n.router.DELETE("/planet/:id", n.RemovePlanetById)
	n.router.DELETE("/planet/n/:name", n.RemovePlanetByExactName)
	n.router.GET("/planets", n.ListAllPlanets)
	n.srv = &http.Server{
		Addr:    "localhost:8080",
		Handler: n.router,
	}

	return n.Close(ctx)
}

func (n *httpServiceFinal) Close(ctx context.Context) error {
	go func() {
		// service connections
		n.sm.LogsService().Info(ctx, "Http Server Listening!")
		if err := n.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err.Error())
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shuting down server ...")

	ctxT, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	if err := n.srv.Shutdown(ctxT); err != nil {
		log.Fatal("Something went wrong executing server shutdown: %s\n", err.Error())
	}

	select {
	case <-ctxT.Done():
		log.Println("timed out after 2 seconds.")
	}
	log.Println("server exiting")

	return nil
}

func (n *httpServiceFinal) Healthy(ctx context.Context) error {
	return nil
}

func (n *httpServiceFinal) WithServiceManager(sm services.ServiceManager) services.HttpService {
	n.sm = sm
	return n
}

func (n *httpServiceFinal) ServiceManager() services.ServiceManager {
	return n.sm
}

func (n *httpServiceFinal) GetPlanetById(c *gin.Context) {
	id := c.Param("id")
	nId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("Endpoint for GetPlanetById expects an integer on the last block of the url. Got a '%s' instead.\n", id)
	}

	p, err := n.sm.PlanetFinderService().GetPlanetById(c.Request.Context(), nId)
	if err != nil {
		//TODO: parse error and return
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Something went wrong"})
		return
	}
	c.IndentedJSON(http.StatusOK, p)

}

func (n *httpServiceFinal) SearchPlanetsByName(c *gin.Context) {
	partialName := c.Param("name")

	p, err := n.sm.PlanetFinderService().SearchPlanetsByName(c.Request.Context(), partialName)
	if err != nil {
		//TODO: parse error and return
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Something went wrong"})
		return
	}

	c.IndentedJSON(http.StatusOK, p)
}

func (n *httpServiceFinal) ListAllPlanets(c *gin.Context) {
	p, err := n.sm.PlanetFinderService().ListAllPlanets(c.Request.Context())
	if err != nil {
		//TODO: parse error and return
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Something went wrong"})
		return
	}

	c.IndentedJSON(http.StatusOK, p)
}

func (n *httpServiceFinal) RemovePlanetById(c *gin.Context) {
	id := c.Param("id")
	nId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("Endpoint for RemovePlanetById expects an integer on the last block of the url. Got a '%s' instead.\n", id)
	}

	if err := n.sm.PlanetFinderService().RemovePlanetById(c.Request.Context(), nId); err != nil {
		//TODO: parse error and return
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Something went wrong"})
		return
	}
	c.IndentedJSON(http.StatusAccepted, gin.H{"message": fmt.Sprintf("Removed Planet with ID %d", nId)})
}

func (n *httpServiceFinal) RemovePlanetByExactName(c *gin.Context) {
	exactName := c.Param("name")

	if err := n.sm.PlanetFinderService().RemovePlanetByExactName(c.Request.Context(), exactName); err != nil {
		//TODO: parse error and return
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Something went wrong"})
		return
	}
	c.IndentedJSON(http.StatusAccepted, gin.H{"message": fmt.Sprintf("Removed Planet with exact name '%s'", exactName)})
}

func (n *httpServiceFinal) Hello(c *gin.Context) {
	name := c.Param("name")
	fmt.Println("Hello, " + name + "!")
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Hello, " + name + "!"})
}
