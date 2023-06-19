package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	forumDelivery "github.com/mmikhail2001/technopark_db_project/internal/app/delivery/forum"
	postDelivery "github.com/mmikhail2001/technopark_db_project/internal/app/delivery/post"
	serviceDelivery "github.com/mmikhail2001/technopark_db_project/internal/app/delivery/service"
	threadDelivery "github.com/mmikhail2001/technopark_db_project/internal/app/delivery/thread"
	userDelivery "github.com/mmikhail2001/technopark_db_project/internal/app/delivery/user"
	voteDelivery "github.com/mmikhail2001/technopark_db_project/internal/app/delivery/vote"

	forumRepo "github.com/mmikhail2001/technopark_db_project/internal/app/repository/forum"
	postRepo "github.com/mmikhail2001/technopark_db_project/internal/app/repository/post"
	serviceRepo "github.com/mmikhail2001/technopark_db_project/internal/app/repository/service"
	threadRepo "github.com/mmikhail2001/technopark_db_project/internal/app/repository/thread"
	userRepo "github.com/mmikhail2001/technopark_db_project/internal/app/repository/user"
	voteRepo "github.com/mmikhail2001/technopark_db_project/internal/app/repository/vote"

	forumUsecase "github.com/mmikhail2001/technopark_db_project/internal/app/usecase/forum"
	postUsecase "github.com/mmikhail2001/technopark_db_project/internal/app/usecase/post"
	serviceUsecase "github.com/mmikhail2001/technopark_db_project/internal/app/usecase/service"
	threadUsecase "github.com/mmikhail2001/technopark_db_project/internal/app/usecase/thread"
	userUsecase "github.com/mmikhail2001/technopark_db_project/internal/app/usecase/user"
	voteUsecase "github.com/mmikhail2001/technopark_db_project/internal/app/usecase/vote"
	"github.com/mmikhail2001/technopark_db_project/internal/pkg/config"
	"github.com/mmikhail2001/technopark_db_project/internal/pkg/sqltools"
)

func main() {
	if err := Run(); err != nil {
		log.Fatalf("%+v", err)
	}
	fmt.Println("success")
}

func Run() error {
	configPtr := flag.String("config", "config/config.yml", "path to config file")

	flag.Parse()

	if *configPtr == "" {
		return fmt.Errorf("needed to pass config file")
	}

	cfg, err := config.NewConfig(*configPtr)
	if err != nil {
		return fmt.Errorf("fail to parse config yml file: %w", err)
	}

	// log.SetOutput(ioutil.Discard)
	log.SetFlags(log.LstdFlags | log.Llongfile)

	client, err := sqltools.NewClientPostgres(cfg.Postgres)
	if err != nil {
		return err
	}

	forumRepo := forumRepo.NewRepository(client)
	userRepo := userRepo.NewRepository(client)
	threadRepo := threadRepo.NewRepository(client)
	postRepo := postRepo.NewRepository(client)
	serviceRepo := serviceRepo.NewRepository(client)
	voteRepo := voteRepo.NewRepository(client)

	forumUsecase := forumUsecase.NewUsecase(forumRepo, userRepo, threadRepo)
	threadUsecase := threadUsecase.NewUsecase(threadRepo, forumRepo, userRepo, postRepo)
	postUsecase := postUsecase.NewUsecase(postRepo, threadRepo, forumRepo, userRepo)
	serviceUsecase := serviceUsecase.NewUsecase(serviceRepo)
	userUsecase := userUsecase.NewUsecase(userRepo)
	voteUsecase := voteUsecase.NewUsecase(voteRepo, userRepo, threadRepo)

	forumDelivery := forumDelivery.NewHandler(forumUsecase)
	threadDelivery := threadDelivery.NewHandler(threadUsecase)
	postDelivery := postDelivery.NewHandler(postUsecase)
	serviceDelivery := serviceDelivery.NewHandler(serviceUsecase)
	userDelivery := userDelivery.NewHandler(userUsecase)
	voteDelivery := voteDelivery.NewHandler(voteUsecase)

	// Routes
	// mw := middleware.NewHTTPMiddleware()
	router := mux.NewRouter()
	// router.Use(mw.AccessLogMiddleware)

	// Forum
	router.HandleFunc("/api/forum/create", forumDelivery.ForumCreate).Methods("POST")
	router.HandleFunc("/api/forum/{slug}/details", forumDelivery.ForumDetails).Methods("GET")
	router.HandleFunc("/api/forum/{slug}/create", threadDelivery.ThreadCreate).Methods("POST")
	router.HandleFunc("/api/forum/{slug}/users", forumDelivery.ForumUsers).Methods("GET")
	router.HandleFunc("/api/forum/{slug}/threads", forumDelivery.ForumThreads).Methods("GET")
	router.HandleFunc("/api/post/{id}/details", postDelivery.PostDetails).Methods("GET")
	router.HandleFunc("/api/post/{id}/details", postDelivery.PostUpdateMessage).Methods("POST")
	router.HandleFunc("/api/service/clear", serviceDelivery.ServiceClear).Methods("POST")
	router.HandleFunc("/api/service/status", serviceDelivery.ServiceStatus).Methods("GET")
	router.HandleFunc("/api/thread/{slug_or_id}/create", postDelivery.PostsCreate).Methods("POST")
	router.HandleFunc("/api/thread/{slug_or_id}/details", threadDelivery.ThreadDetails).Methods("GET")
	router.HandleFunc("/api/thread/{slug_or_id}/details", threadDelivery.ThreadUpdate).Methods("POST")
	router.HandleFunc("/api/thread/{slug_or_id}/posts", threadDelivery.ThreadPosts).Methods("GET")
	router.HandleFunc("/api/thread/{slug_or_id}/vote", voteDelivery.Vote).Methods("POST")
	router.HandleFunc("/api/user/{nickname}/create", userDelivery.UserCreate).Methods("POST")
	router.HandleFunc("/api/user/{nickname}/profile", userDelivery.UserDetails).Methods("GET")
	router.HandleFunc("/api/user/{nickname}/profile", userDelivery.UserUpdate).Methods("POST")

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	addr := fmt.Sprintf("%s:%s", cfg.Server.BindIP, cfg.Server.Port)

	server := http.Server{
		Addr:              addr,
		Handler:           router,
		ReadHeaderTimeout: time.Second * time.Duration(cfg.Server.ReadHeaderTimeout),
		WriteTimeout:      time.Second * time.Duration(cfg.Server.WriteTimeout),
		ReadTimeout:       time.Second * time.Duration(cfg.Server.ReadTimeout),
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal("listen and server error: ", err)
		}
	}()
	log.Println("start listening on", addr)

	<-ctx.Done()

	log.Println("server shutdown")

	shutdownCtx, cancel := context.WithTimeout(context.Background(),
		time.Second*time.Duration(cfg.Server.ShutdownTimeout))
	defer cancel()

	err = server.Shutdown(shutdownCtx)
	if err != nil {
		return fmt.Errorf("shutdown: %w", err)
	}
	return nil
}
