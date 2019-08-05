package app

import (
	"context"
	"log"
	"math/big"
	"net/http"
	"os"

	"github.com/bansatya/v4/contracts/quiz"
	"github.com/bansatya/v4/handler"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var myenv map[string]string

const envLoc = ".env"

// App has router and db instances
type App struct {
	Router *mux.Router
	//DB     *gorm.DB
}

// App initialize with predefined configuration
func (a *App) Initialize() {

	a.Router = mux.NewRouter()
	a.setRouters()

	//Init env
	loadEnv()
	ctx := context.Background()
	//   client, err := ethclient.Dial(os.Getenv("GATEWAY"))
	client, err := ethclient.Dial("/home/ubuntu/eth_node/node1/datadir/geth.ipc") //ethclient.Dial("https://rinkeby.infura.io/v3/6ae645967acc47fda0d4854e18e6b0ec");ethclient.Dial("/home/ubuntu/eth_node/node1/datadir/geth.ipc")
	if err != nil {
		log.Fatalf("could not connect to Ethereum gateway: %v\n", err)
	}
	defer client.Close()
	session := NewSession(ctx)
}

func loadEnv() {
	var err error
	if myenv, err = godotenv.Read(envLoc); err != nil {
		log.Printf("could not load env from %s: %v", envLoc, err)
	}
}

//Initialize session
func NewSession(ctx context.Context) (session quiz.QuizSession) {
	loadEnv()
	keystore, err := os.Open(myenv["KEYSTORE"])
	if err != nil {
		log.Printf(
			"could not load keystore from location %s: %v\n",
			myenv["KEYSTORE"],
			err,
		)
	}
	defer keystore.Close()

	keystorepass := myenv["KEYSTOREPASS"]
	auth, err := bind.NewTransactor(keystore, keystorepass)
	// *auth.GasLimit  = big.NewInt(3141592);
	if err != nil {
		log.Printf("%s\n", err)
	}

	// Return session without contract instance
	return quiz.QuizSession{
		// TransactOpts: *auth,
		TransactOpts: bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasLimit: 4700000,
			GasPrice: big.NewInt(1000000),
		},
		CallOpts: bind.CallOpts{
			From:    auth.From,
			Context: ctx,
		},
	}
}

// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/ping", a.Pong)
	a.Get("/accounts/{adddress}", a.GetAccount)
	a.Post("/accounts", a.CreateAccount)
	a.Post("/contracts", a.CreateContract)
	a.Post("/contracts/{address}/methods", a.Execute)
	a.Get("/contracts/{address}/methods", a.Call)
}

func (a *App) Pong(w http.ResponseWriter, r *http.Request) {
	handler.Pong(w, r)
}

// Handlers to manage Employee Data
func (a *App) GetAccount(w http.ResponseWriter, r *http.Request) {
	handler.GetAccount(w, r)
}

func (a *App) CreateAccount(w http.ResponseWriter, r *http.Request) {
	handler.CreateAccount(w, r)
}

func (a *App) CreateContract(w http.ResponseWriter, r *http.Request) {
	handler.CreateContract(w, r)
}

func (a *App) Execute(w http.ResponseWriter, r *http.Request) {
	handler.Execute(w, r)
}

func (a *App) Call(w http.ResponseWriter, r *http.Request) {
	handler.Call(w, r)
}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
